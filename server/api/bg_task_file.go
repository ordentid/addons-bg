package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
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

func (s *Server) GetTaskMappingFile(ctx context.Context, req *pb.GetTaskMappingFileRequest) (*httpbody.HttpBody, error) {
	result := &httpbody.HttpBody{}

	reqPB := &pb.GetTaskMappingRequest{
		Limit:  req.Limit,
		Page:   req.Page,
		Sort:   req.Sort,
		Dir:    req.Dir,
		Filter: req.Filter,
		Query:  req.Query,
	}

	resPB, err := s.GetTaskMapping(ctx, reqPB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	file := GetTaskMappingFileGenerator(resPB)
	if req.FileFormat.String() == "pdf" {
		return file.ToPDFv2(ctx)
	}
	if req.FileFormat.String() == "csv" {
		return file.ToCsv(ctx)
	}

	if req.FileFormat.String() == "xls" {
		return file.ToXls(ctx)
	}

	return result, nil

}

type TaskMappingFile struct {
	res *pb.GetTaskMappingResponse
}

func GetTaskMappingFileGenerator(res *pb.GetTaskMappingResponse) *TaskMappingFile {
	return &TaskMappingFile{res: res}
}

func (file *TaskMappingFile) ToPDFv2(ctx context.Context) (*httpbody.HttpBody, error) {
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

	fields := []string{"No", "Company", "Third Party Count", "Date Created", "Date Modified", "Status"}
	widths := []float64{8, 30, 35, 20, 28, 20}
	align := []string{"TL", "TL", "TL", "TL", "TL", "TC"}

	var (
		cellList [9]cellType
		cell     cellType
	)

	pdf.AddPage()
	_, pageh := pdf.GetPageSize()

	pdf.SetFont("Times", "B", 9.5)
	pdf.SetFillColor(240, 240, 240)
	pdf.SetX(marginH)
	for i, header := range fields {
		pdf.CellFormat(widths[i], lineHt, header, "1", 0, align[i], true, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 9.5)
	pdf.SetFillColor(255, 255, 255)

	y := pdf.GetY()
	x := marginH

	for index, v := range file.res.Data {

		company := string(v.Company.CompanyName)

		thirdPartyCount := len(v.Data)

		curYear, _, _ := time.Now().Date()
		dateCreated := ""
		dateModified := ""

		err := v.Task.CreatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.CreatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateCreated = v.Task.CreatedAt.AsTime().Format("02/01/2006")
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

		maxHt := lineHt
		vals := []string{
			fmt.Sprintf("%d", index+1),
			company,
			strconv.FormatUint(uint64(thirdPartyCount), 10),
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
			x, y = pdf.GetXY()
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
			x, y = pdf.GetXY()
		}

	}

	var buf bytes.Buffer
	var err error

	err = pdf.Output(&buf)
	if err == nil {
		logrus.Println("Length of buffer: %d\n", buf.Len())
		// return nil, status.Errorf(codes.Internal, "Server error")
	} else {
		logrus.Errorf("Error generating PDF: %s\n", err)
		return nil, status.Errorf(codes.Internal, "Server error")
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"Account.pdf\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/pdf",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}

func (file *TaskMappingFile) ToCsv(ctx context.Context) (*httpbody.HttpBody, error) {
	var buf bytes.Buffer

	w := csv.NewWriter(&buf)

	fields := []string{"No", "Company", "Third Party Count", "Date Created", "Date Modified", "Status"}

	_ = w.Write(fields)

	for index, v := range file.res.Data {

		company := string(v.Company.CompanyName)

		thirdPartyCount := len(v.Data)

		curYear, _, _ := time.Now().Date()
		dateCreated := ""
		dateModified := ""

		err := v.Task.CreatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.CreatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateCreated = v.Task.CreatedAt.AsTime().Format("02/01/2006")
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
		row := []string{
			fmt.Sprintf("%d", index+1),
			company,
			strconv.FormatUint(uint64(thirdPartyCount), 10),
			dateCreated,
			dateModified,
			status,
		}
		_ = w.Write(row)
	}

	w.Flush()
	err := w.Error()
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"Account.csv\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/csv",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil

}

func (file *TaskMappingFile) ToXls(ctx context.Context) (*httpbody.HttpBody, error) {
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Company")
	_ = f.SetCellValue("Sheet1", "C1", "Third Party Count")
	_ = f.SetCellValue("Sheet1", "D1", "Date Created")
	_ = f.SetCellValue("Sheet1", "E1", "Date Modified")
	_ = f.SetCellValue("Sheet1", "F1", "Status")

	for k, v := range file.res.Data {

		company := string(v.Company.CompanyName)

		thirdPartyCount := len(v.Data)

		curYear, _, _ := time.Now().Date()
		dateCreated := ""
		dateModified := ""

		err := v.Task.CreatedAt.CheckValid()
		if err == nil {
			year, _, _ := v.Task.CreatedAt.AsTime().Date()
			yearDiff := curYear - year
			if yearDiff < 10 && yearDiff > -10 {
				dateCreated = v.Task.CreatedAt.AsTime().Format("02/01/2006")
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
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNumber), strconv.FormatUint(uint64(thirdPartyCount), 10))
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNumber), dateCreated)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNumber), dateModified)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNumber), status)

	}

	f.SetActiveSheet(sheet)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("file-download", ""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"Account.xlsx\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil
}
