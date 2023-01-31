package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"time"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/jung-kurt/gofpdf"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTaskMappingDigitalFile(ctx context.Context, req *pb.GetTaskMappingDigitalFileRequest) (*httpbody.HttpBody, error) {
	result := &httpbody.HttpBody{}

	reqPB := &pb.GetTaskMappingDigitalRequest{
		Limit:  req.Limit,
		Page:   req.Page,
		Sort:   req.Sort,
		Dir:    req.Dir,
		Filter: req.Filter,
		Query:  req.Query,
	}

	resPB, err := s.GetTaskMappingDigital(ctx, reqPB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	file := GetTaskMappingDigitalFileGenerator(resPB)
	if req.FileFormat.String() == "pdf" {
		return file.TaskMappingDigitalToPDFv2(ctx)
	}
	if req.FileFormat.String() == "csv" {
		return file.TaskMappingDigitalToCsv(ctx)
	}

	if req.FileFormat.String() == "xls" {
		return file.TaskMappingDigitalToXls(ctx)
	}

	return result, nil

}

type TaskMappingDigitalFile struct {
	res *pb.GetTaskMappingDigitalResponse
}

func GetTaskMappingDigitalFileGenerator(res *pb.GetTaskMappingDigitalResponse) *TaskMappingDigitalFile {
	return &TaskMappingDigitalFile{res: res}
}

func (file *TaskMappingDigitalFile) TaskMappingDigitalToPDFv2(ctx context.Context) (*httpbody.HttpBody, error) {

	var buf bytes.Buffer
	var err error

	var y float64
	var x float64

	const (
		colCount = 3
		colWd    = 60.0
		marginH  = 10.0
		lineHt   = 5.5
		cellGap  = 1.0
	)
	type cellType struct {
		str  string
		list [][]byte
		ht   float64
	}

	pdf := gofpdf.New("L", "mm", "Letter", "")

	fields := []string{"No", "Company", "Third Party", "Beneficiary", "Date Created", "Date Modified", "Status"}
	widths := []float64{8, 30, 35, 20, 20, 28, 20}
	align := []string{"TL", "TL", "TL", "TL", "TL", "TL", "TC"}

	var (
		cellList [7]cellType
		cell     cellType
	)

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	exportedAt := time.Now().In(location).Format("02/01/2006")

	pdf.SetTopMargin(30)

	pdf.SetHeaderFuncMode(func() {
		pdf.Image("assets/bricams.png", 12.5, 12, 40, 0, false, "", 0, "")
		pdf.Image("assets/bri.png", 240, 12, 20, 0, false, "", 0, "")
		pdf.SetFont("Times", "B", 9.5)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetFillColor(2, 75, 140)
		pdf.SetX(marginH)
		for i, header := range fields {
			pdf.CellFormat(widths[i], lineHt, header, "1", 0, align[i], true, 0, "")
		}
		pdf.Ln(-1)
		_, y = pdf.GetXY()
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Times", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb} - %v", pdf.PageNo(), exportedAt),
			"", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")

	pdf.AddPage()
	_, pageh := pdf.GetPageSize()

	pdf.SetX(marginH)
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 9.5)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(255, 255, 255)

	y = pdf.GetY()
	x = marginH

	for index, v := range file.res.Data {

		company := string(v.Company.CompanyName)

		thirdPartyName := v.Data[0].ThirdPartyName

		beneficiaryName := v.Data[0].BeneficiaryName

		curYear, _, _ := time.Now().Date()
		dateCreated := ""
		dateModified := ""

		err = v.Task.CreatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.CreatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateCreated = v.Task.CreatedAt.AsTime().In(location).Format("02/01/2006")
			}
		}

		err = v.Task.UpdatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.UpdatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateModified = v.Task.UpdatedAt.AsTime().In(location).Format("02/01/2006")
			}
		}

		status := v.Task.Status

		maxHt := lineHt
		vals := []string{
			fmt.Sprintf("%d", index+1),
			company,
			thirdPartyName,
			beneficiaryName,
			dateCreated,
			dateModified,
			status,
		}
		// Cell height calculation loop
		for colJ := 0; colJ < len(vals); colJ++ {
			cell.str = vals[colJ]
			cell.list = pdf.SplitLines([]byte(cell.str), widths[colJ])
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[colJ] = cell
		}

		if y >= pageh-((marginH*2)+maxHt+lineHt) {
			pdf.AddPage()
			pdf.SetX(marginH)
			pdf.Ln(-1)
			_, y = pdf.GetXY()
		}

		// Cell render loop
		x = marginH
		for colJ := 0; colJ < len(vals); colJ++ {
			pdf.Rect(x, y, widths[colJ], maxHt+cellGap+cellGap, "D")
			cell = cellList[colJ]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x, cellY)
				pdf.CellFormat(widths[colJ], lineHt, string(cell.list[splitJ]), "", 0,
					align[colJ], false, 0, "")
				cellY += lineHt

			}
			x += widths[colJ]
		}
		y += maxHt + cellGap + cellGap

		if y >= pageh-((marginH*2)+lineHt) {
			pdf.AddPage()
			pdf.SetX(marginH)
			pdf.Ln(-1)
			_, y = pdf.GetXY()
		}

	}

	err = pdf.Output(&buf)
	if err == nil {
		logrus.Println("Length of buffer: %d\n", buf.Len())
		// return nil, status.Errorf(codes.Internal, "Server error")
	} else {
		logrus.Errorf("Error generating PDF: %s\n", err)
		return nil, status.Errorf(codes.Internal, "Server error")
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"BG.pdf\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/pdf",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}

func (file *TaskMappingDigitalFile) TaskMappingDigitalToCsv(ctx context.Context) (*httpbody.HttpBody, error) {
	var buf bytes.Buffer

	w := csv.NewWriter(&buf)

	fields := []string{"No", "Company", "Third Party", "Beneficiary", "Date Created", "Date Modified", "Status"}

	_ = w.Write(fields)

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	for index, v := range file.res.Data {

		company := string(v.Company.CompanyName)

		thirdPartyName := v.Data[0].ThirdPartyName

		beneficiaryName := v.Data[0].BeneficiaryName

		curYear, _, _ := time.Now().Date()
		dateCreated := ""
		dateModified := ""

		err := v.Task.CreatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.CreatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateCreated = v.Task.CreatedAt.AsTime().In(location).Format("02/01/2006")
			}
		}

		err = v.Task.UpdatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.UpdatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateModified = v.Task.UpdatedAt.AsTime().In(location).Format("02/01/2006")
			}
		}

		status := v.Task.Status
		row := []string{
			fmt.Sprintf("%d", index+1),
			company,
			thirdPartyName,
			beneficiaryName,
			dateCreated,
			dateModified,
			status,
		}
		_ = w.Write(row)
	}

	w.Flush()
	err = w.Error()
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"BG.csv\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/csv",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil

}

func (file *TaskMappingDigitalFile) TaskMappingDigitalToXls(ctx context.Context) (*httpbody.HttpBody, error) {
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Company")
	_ = f.SetCellValue("Sheet1", "C1", "Third Party")
	_ = f.SetCellValue("Sheet1", "C1", "Beneficiary")
	_ = f.SetCellValue("Sheet1", "D1", "Date Created")
	_ = f.SetCellValue("Sheet1", "E1", "Date Modified")
	_ = f.SetCellValue("Sheet1", "F1", "Status")

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	for k, v := range file.res.Data {

		company := string(v.Company.CompanyName)

		thirdPartyName := v.Data[0].ThirdPartyName

		beneficiaryName := v.Data[0].BeneficiaryName

		curYear, _, _ := time.Now().Date()
		dateCreated := ""
		dateModified := ""

		err := v.Task.CreatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.CreatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateCreated = v.Task.CreatedAt.AsTime().In(location).Format("02/01/2006")
			}
		}

		err = v.Task.UpdatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.UpdatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateModified = v.Task.UpdatedAt.AsTime().Format("02/01/2006")
			}
		}

		status := v.Task.Status
		rowNumber := k + 2
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowNumber), fmt.Sprintf("%d", k+1))
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNumber), company)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNumber), thirdPartyName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNumber), beneficiaryName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNumber), dateCreated)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNumber), dateModified)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowNumber), status)

	}

	f.SetActiveSheet(sheet)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"BG.xlsx\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}
