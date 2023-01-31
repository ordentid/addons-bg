package servicelogger

import "github.com/sirupsen/logrus"

type TransactionData struct {
	UserID           uint64 `json:"user_id"`
	Username         string `json:"username"`
	CompanyID        uint64 `json:"company_id"`
	CompanyName      string `json:"company_name"`
	AccountDebit     string `json:"account_debit"`
	AccountCredit    string `json:"account_credit"`
	Currency         string `json:"currency"`
	AmountRate       string `json:"amount_rate"`
	AmountTrx        string `json:"amount_trx"`
	TotalFee         string `json:"total_fee"`
	TransactionID    string `json:"transaction_id"`
	TransactionType  string `json:"transaction_type"`
	IsRequestSucceed bool   `json:"is_request_Succeed"`
	RequestURL       string `json:"request_url"`
	HeaderReq        string `json:"req_header"`
	BodyReq          string `json:"req_body"`
	ResStatusCode    int    `json:"res_status_code"`
	BodyResp         string `json:"res_body"`
}

type TransactionDataTag struct {
	log     *logrus.Logger
	TagName string
	Data    *TransactionData
}

func NewDataTag(data *TransactionData, tagName string, log *logrus.Logger) *TransactionDataTag {

	if tagName == "" {
		tagName = getEnv("DATA_TAG1", "")
	}

	if data == nil {
		data = &TransactionData{
			UserID:           0,
			Username:         "",
			CompanyID:        0,
			CompanyName:      "",
			AccountDebit:     "",
			AccountCredit:    "",
			Currency:         "",
			AmountRate:       "",
			AmountTrx:        "",
			TotalFee:         "",
			TransactionID:    "",
			TransactionType:  "",
			IsRequestSucceed: false,
			RequestURL:       "",
			HeaderReq:        "",
			BodyReq:          "",
			ResStatusCode:    0,
			BodyResp:         "",
		}
	}

	return &TransactionDataTag{
		log:     log,
		TagName: tagName,
		Data:    data,
	}
}

func (dl *TransactionDataTag) SendData(textinfo string) {
	dl.log.WithField("data_tag", dl.TagName).WithField("transaction", dl.Data).Info(textinfo)
}

func (dl *TransactionDataTag) SetData(data *TransactionData) {
	dl.Data = data
}

func (dl *TransactionDataTag) SetTagName(tagName string) {
	dl.TagName = tagName
}

func (dl *TransactionDataTag) SetLog(log *logrus.Logger) {
	dl.log = log
}

func (dl *TransactionDataTag) SetUserID(x uint64) {
	dl.Data.UserID = x
}
func (dl *TransactionDataTag) SetUsername(x string) {
	dl.Data.Username = x
}
func (dl *TransactionDataTag) SetCompanyID(x uint64) {
	dl.Data.CompanyID = x
}
func (dl *TransactionDataTag) SetCompanyName(x string) {
	dl.Data.CompanyName = x
}
func (dl *TransactionDataTag) SetAccountDebit(x string) {
	dl.Data.AccountDebit = x
}
func (dl *TransactionDataTag) SetAccountCredit(x string) {
	dl.Data.AccountCredit = x
}
func (dl *TransactionDataTag) SetCurrency(x string) {
	dl.Data.Currency = x
}
func (dl *TransactionDataTag) SetAmountRate(x string) {
	dl.Data.AmountRate = x
}
func (dl *TransactionDataTag) SetAmountTrx(x string) {
	dl.Data.AmountTrx = x
}
func (dl *TransactionDataTag) SetTotalFee(x string) {
	dl.Data.TotalFee = x
}
func (dl *TransactionDataTag) SetTransactionID(x string) {
	dl.Data.TransactionID = x
}
func (dl *TransactionDataTag) SetTransactionType(x string) {
	dl.Data.TransactionType = x
}
func (dl *TransactionDataTag) SetIsRequestSucceed(x bool) {
	dl.Data.IsRequestSucceed = x
}
func (dl *TransactionDataTag) SetRequestURL(x string) {
	dl.Data.RequestURL = x
}
func (dl *TransactionDataTag) SetHeaderReq(x string) {
	dl.Data.HeaderReq = x
}
func (dl *TransactionDataTag) SetBodyReq(x string) {
	dl.Data.BodyReq = x
}
func (dl *TransactionDataTag) SetResStatusCode(x int) {
	dl.Data.ResStatusCode = x
}
func (dl *TransactionDataTag) SetBodyResp(x string) {
	dl.Data.BodyResp = x
}
