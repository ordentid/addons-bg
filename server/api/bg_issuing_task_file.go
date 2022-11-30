package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"time"

	"bitbucket.bri.co.id/scm/addons/addons-bg-service/server/pb"
	"github.com/jung-kurt/gofpdf"
	"github.com/leekchan/accounting"
	"github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *Server) GetTaskIssuingFile(ctx context.Context, req *pb.GetTaskIssuingFileRequest) (*httpbody.HttpBody, error) {

	result := &httpbody.HttpBody{}

	reqPB := &pb.GetTaskIssuingRequest{
		Limit:  req.Limit,
		Page:   req.Page,
		Sort:   req.Sort,
		Dir:    req.Dir,
		Filter: req.Filter,
		Query:  req.Query,
	}

	resPB, err := s.GetTaskIssuing(ctx, reqPB)
	if err != nil {
		logrus.Errorln("[api][func: GetTaskIssuingFile] Unable to Get Task Issuing:", err.Error())
		return nil, err
	}

	file := GetTaskIssuingFileGenerator(resPB)
	if req.FileFormat.String() == "pdf" {
		return file.TaskIssuingToPDFv2(ctx)
	}
	if req.FileFormat.String() == "csv" {
		return file.TaskIssuingToCsv(ctx)
	}
	if req.FileFormat.String() == "xls" {
		return file.TaskIssuingToXls(ctx)
	}

	return result, nil

}

type TaskIssuingFile struct {
	res *pb.GetTaskIssuingResponse
}

func GetTaskIssuingFileGenerator(res *pb.GetTaskIssuingResponse) *TaskIssuingFile {
	return &TaskIssuingFile{res: res}
}

func (file *TaskIssuingFile) TaskIssuingToPDFv2(ctx context.Context) (*httpbody.HttpBody, error) {

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

	fields := []string{"No", "Reference Number", "Registration Number", "Applicant Name", "Beneficiary Name", "Created By", "Reviewed By", "BG Type", "Amount", "Workflow Status"}
	widths := []float64{8, 30, 30, 30, 25, 20, 20, 25, 20, 35}
	align := []string{"TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL"}

	var (
		cellList [10]cellType
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

		bgTypeStrings := []string{
			"Bid Bond",
			"Advance Payment",
			"Performance Bond",
			"Goverment Payment Guarantee",
			"Maintenance Bond",
			"Procurement Bond",
			"Transaction Risk Bond",
			"Customs Bond",
		}

		bgType := bgTypeStrings[v.Data.Publishing.BgType.Number()]

		ac := accounting.Accounting{Symbol: v.Data.Project.BgCurrency, Precision: 2}
		bgAmount := ac.FormatMoney(v.Data.Project.BgAmount)

		status := v.Task.Status

		referenceNumber := v.Data.ReferenceNo
		registrationNo := v.Data.RegistrationNo

		if referenceNumber == "" {
			referenceNumber = "TBA"
		}

		if registrationNo == "" {
			registrationNo = "TBA"
		}

		maxHt := lineHt
		vals := []string{
			fmt.Sprintf("%d", index+1),
			referenceNumber,
			registrationNo,
			v.Data.Applicant.Name,
			v.Data.Applicant.BeneficiaryName,
			v.Task.CreatedByName,
			v.Task.LastApprovedByName,
			bgType,
			bgAmount,
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
			pdf.SetX(marginH)
			pdf.Ln(-1)
			x, y = pdf.GetXY()
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

func (file *TaskIssuingFile) TaskIssuingToCsv(ctx context.Context) (*httpbody.HttpBody, error) {

	var buf bytes.Buffer

	w := csv.NewWriter(&buf)

	fields := []string{"No", "Reference Number", "Registration Number", "Applicant Name", "Beneficiary Name", "Created By", "Reviewed By", "BG Type", "Amount", "Workflow Status"}

	_ = w.Write(fields)

	for index, v := range file.res.Data {

		bgTypeStrings := []string{
			"Bid Bond",
			"Advance Payment",
			"Performance Bond",
			"Goverment Payment Guarantee",
			"Maintenance Bond",
			"Procurement Bond",
			"Transaction Risk Bond",
			"Customs Bond",
		}

		bgType := bgTypeStrings[v.Data.Publishing.BgType.Number()]
		ac := accounting.Accounting{Symbol: v.Data.Project.BgCurrency, Precision: 2}
		bgAmount := ac.FormatMoney(v.Data.Project.BgAmount)
		status := v.Task.Status
		referenceNumber := v.Data.ReferenceNo
		registrationNo := v.Data.RegistrationNo

		if referenceNumber == "" {
			referenceNumber = "TBA"
		}

		if registrationNo == "" {
			registrationNo = "TBA"
		}

		row := []string{
			fmt.Sprintf("%d", index+1),
			referenceNumber,
			registrationNo,
			v.Data.Applicant.Name,
			v.Data.Applicant.BeneficiaryName,
			v.Task.CreatedByName,
			v.Task.LastApprovedByName,
			bgType,
			bgAmount,
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
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Disposition", "attachment; filename=\"BG.csv\""))
	_ = grpc.SetHeader(ctx, metadata.Pairs("Content-Length", fmt.Sprintf("%v", buf.Len())))

	return &httpbody.HttpBody{
		ContentType: "application/csv",
		Data:        buf.Bytes(),
		Extensions:  nil,
	}, nil

}

func (file *TaskIssuingFile) TaskIssuingToXls(ctx context.Context) (*httpbody.HttpBody, error) {

	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Reference Number")
	_ = f.SetCellValue("Sheet1", "C1", "Registration Number")
	_ = f.SetCellValue("Sheet1", "D1", "Applicant Name")
	_ = f.SetCellValue("Sheet1", "E1", "Beneficiary Name")
	_ = f.SetCellValue("Sheet1", "F1", "Created By")
	_ = f.SetCellValue("Sheet1", "G1", "Reviewed By")
	_ = f.SetCellValue("Sheet1", "H1", "BG Type")
	_ = f.SetCellValue("Sheet1", "I1", "Amount")
	_ = f.SetCellValue("Sheet1", "J1", "Workkflow Status")

	for k, v := range file.res.Data {

		bgTypeStrings := []string{
			"Bid Bond",
			"Advance Payment",
			"Performance Bond",
			"Goverment Payment Guarantee",
			"Maintenance Bond",
			"Procurement Bond",
			"Transaction Risk Bond",
			"Customs Bond",
		}

		bgType := bgTypeStrings[v.Data.Publishing.BgType.Number()]
		ac := accounting.Accounting{Symbol: v.Data.Project.BgCurrency, Precision: 2}
		bgAmount := ac.FormatMoney(v.Data.Project.BgAmount)
		status := v.Task.Status
		referenceNumber := v.Data.ReferenceNo
		registrationNo := v.Data.RegistrationNo

		if referenceNumber == "" {
			referenceNumber = "TBA"
		}

		if registrationNo == "" {
			registrationNo = "TBA"
		}

		rowNumber := k + 2
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowNumber), fmt.Sprintf("%d", k+1))
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNumber), referenceNumber)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNumber), registrationNo)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNumber), v.Data.Applicant.Name)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNumber), v.Data.Applicant.BeneficiaryName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNumber), v.Task.CreatedByName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowNumber), v.Task.LastApprovedByName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", rowNumber), bgType)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("I%d", rowNumber), bgAmount)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("J%d", rowNumber), status)
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
