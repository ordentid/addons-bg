package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
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

func (s *Server) GetTransactionFile(ctx context.Context, req *pb.GetTransactionFileRequest) (*httpbody.HttpBody, error) {
	result := &httpbody.HttpBody{}

	reqPB := &pb.GetTransactionRequest{
		Transaction: req.Transaction,
		Limit:       req.Limit,
		Page:        req.Page,
		Sort:        req.Sort,
		Dir:         req.Dir,
		Filter:      req.Filter,
		Query:       req.Query,
	}

	resPB, err := s.GetTransaction(ctx, reqPB)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	file := GetTransactionFileGenerator(resPB)
	if req.FileFormat.String() == "pdf" {
		return file.TransactionToPDFv2(ctx)
	}
	if req.FileFormat.String() == "csv" {
		return file.TransactionToCsv(ctx)
	}

	if req.FileFormat.String() == "xls" {
		return file.TransactionToXls(ctx)
	}

	return result, nil

}

type TransactionFile struct {
	res *pb.GetTransactionResponse
}

func GetTransactionFileGenerator(res *pb.GetTransactionResponse) *TransactionFile {
	return &TransactionFile{res: res}
}

func (file *TransactionFile) TransactionToPDFv2(ctx context.Context) (*httpbody.HttpBody, error) {

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

	fields := []string{"No", "Reference Number", "Registration Number", "Third Party", "Applicant", "Beneficiary", "Date Issued", "Effective Date", "Maturity Date", "Claim Period", "BG Type", "Amount", "BG Status"}
	widths := []float64{8, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20}
	align := []string{"TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TL", "TC"}

	var (
		cellList [13]cellType
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

		issueDateArr := strings.Split(v.IssueDate, "-")
		issueDate := issueDateArr[2] + "/" + issueDateArr[1] + "/" + issueDateArr[0]

		effectiveDateArr := strings.Split(v.EffectiveDate, "-")
		effectiveDate := effectiveDateArr[2] + "/" + effectiveDateArr[1] + "/" + effectiveDateArr[0]

		expiryDateArr := strings.Split(v.ExpiryDate, "-")
		expiryDate := expiryDateArr[2] + "/" + expiryDateArr[1] + "/" + expiryDateArr[0]

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
		bgType := bgTypeStrings[v.TransactionTypeID.Number()]

		ac := accounting.Accounting{Symbol: v.Currency, Precision: 2}
		amount := ac.FormatMoney(v.Amount)

		maxHt := lineHt
		vals := []string{
			fmt.Sprintf("%d", index+1),
			v.ReferenceNo,
			v.RegistrationNo,
			v.ThirdPartyName,
			v.ApplicantName,
			v.BeneficiaryName,
			issueDate,
			effectiveDate,
			expiryDate,
			strconv.FormatUint(uint64(v.ClaimPeriod), 10) + " day(s)",
			bgType,
			amount,
			v.Status,
		}
		// Cell height calculation loop
		for colJ := 0; colJ < len(vals); colJ++ {
			logrus.Print("Column Index: ", colJ)
			logrus.Print("Length: ", len(widths))
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

func (file *TransactionFile) TransactionToCsv(ctx context.Context) (*httpbody.HttpBody, error) {
	var buf bytes.Buffer

	w := csv.NewWriter(&buf)

	fields := []string{"No", "Reference Number", "Registration Number", "Third Party", "Applicant", "Beneficiary", "Date Issued", "Effective Date", "Maturity Date", "Claim Period", "BG Type", "Amount", "BG Status"}

	_ = w.Write(fields)

	for index, v := range file.res.Data {

		issueDateArr := strings.Split(v.IssueDate, "-")
		issueDate := issueDateArr[2] + "/" + issueDateArr[1] + "/" + issueDateArr[0]

		effectiveDateArr := strings.Split(v.EffectiveDate, "-")
		effectiveDate := effectiveDateArr[2] + "/" + effectiveDateArr[1] + "/" + effectiveDateArr[0]

		expiryDateArr := strings.Split(v.ExpiryDate, "-")
		expiryDate := expiryDateArr[2] + "/" + expiryDateArr[1] + "/" + expiryDateArr[0]

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
		bgType := bgTypeStrings[v.TransactionTypeID.Number()]

		ac := accounting.Accounting{Symbol: v.Currency, Precision: 2}
		amount := ac.FormatMoney(v.Amount)

		row := []string{
			fmt.Sprintf("%d", index+1),
			v.ReferenceNo,
			v.RegistrationNo,
			v.ThirdPartyName,
			v.ApplicantName,
			v.BeneficiaryName,
			issueDate,
			effectiveDate,
			expiryDate,
			strconv.FormatUint(uint64(v.ClaimPeriod), 10) + " day(s)",
			bgType,
			amount,
			v.Status,
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

func (file *TransactionFile) TransactionToXls(ctx context.Context) (*httpbody.HttpBody, error) {
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	_ = f.SetCellValue("Sheet1", "A1", "No")
	_ = f.SetCellValue("Sheet1", "B1", "Reference Number")
	_ = f.SetCellValue("Sheet1", "C1", "Registration Number")
	_ = f.SetCellValue("Sheet1", "D1", "Third Party")
	_ = f.SetCellValue("Sheet1", "E1", "Applicant")
	_ = f.SetCellValue("Sheet1", "F1", "Beneficiary")
	_ = f.SetCellValue("Sheet1", "G1", "Date Issued")
	_ = f.SetCellValue("Sheet1", "H1", "Effective Date")
	_ = f.SetCellValue("Sheet1", "I1", "Maturity Date")
	_ = f.SetCellValue("Sheet1", "J1", "Claim Period")
	_ = f.SetCellValue("Sheet1", "K1", "BG Type")
	_ = f.SetCellValue("Sheet1", "L1", "Amount")
	_ = f.SetCellValue("Sheet1", "M1", "BG Status")

	for k, v := range file.res.Data {

		issueDateArr := strings.Split(v.IssueDate, "-")
		issueDate := issueDateArr[2] + "/" + issueDateArr[1] + "/" + issueDateArr[0]

		effectiveDateArr := strings.Split(v.EffectiveDate, "-")
		effectiveDate := effectiveDateArr[2] + "/" + effectiveDateArr[1] + "/" + effectiveDateArr[0]

		expiryDateArr := strings.Split(v.ExpiryDate, "-")
		expiryDate := expiryDateArr[2] + "/" + expiryDateArr[1] + "/" + expiryDateArr[0]

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
		bgType := bgTypeStrings[v.TransactionTypeID.Number()]

		ac := accounting.Accounting{Symbol: v.Currency, Precision: 2}
		amount := ac.FormatMoney(v.Amount)

		rowNumber := k + 2
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowNumber), fmt.Sprintf("%d", k+1))
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNumber), v.ReferenceNo)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNumber), v.RegistrationNo)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNumber), v.ThirdPartyName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("E%d", rowNumber), v.ApplicantName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("F%d", rowNumber), v.BeneficiaryName)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("G%d", rowNumber), issueDate)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("H%d", rowNumber), effectiveDate)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("I%d", rowNumber), expiryDate)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("J%d", rowNumber), strconv.FormatUint(uint64(v.ClaimPeriod), 10)+" day(s)")
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("K%d", rowNumber), bgType)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("L%d", rowNumber), amount)
		_ = f.SetCellValue("Sheet1", fmt.Sprintf("M%d", rowNumber), v.Status)

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
