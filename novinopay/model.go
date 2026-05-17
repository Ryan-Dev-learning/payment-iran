package novinopay

type PaymentStatus string

const (
	OK  PaymentStatus = "OK"
	NOK PaymentStatus = "NOK"
)

var invoiceStorage = make(map[string]Invoice)

type Invoice struct {
	InvoiceID   string `json:"invoice_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	Name        string `json:"name"`
	CardPan     string `json:"card_pan"`
}

type InitTransactionRequest struct {
	MerchantID     string `json:"merchant_id" example:"test"`
	Amount         int    `json:"amount" example:"10000"`
	CallbackUrl    string `json:"callback_url" example:"http://localhost:9091/api/v1/novinopay/callback"`
	CallbackMethod string `json:"callback_method" example:"GET"`
	InvoiceID      string `json:"invoice_id" example:"order_123456"`
	Description    string `json:"description" example:"NovinoPay Test"`
	Email          string `json:"email" example:"a@b.com"`
	Mobile         string `json:"mobile" example:"09121234567"`
	Name           string `json:"name" example:"ali yari"`
	CardPan        string `json:"card_pan" example:"6037861810449774"`
}

type InitTransactionResponse struct {
	Status  string               `json:"status" example:"100"`
	Message string               `json:"message" example:"عملیات موفقیت آمیز"`
	Data    *InitTransactionData `json:"data"`
	Errors  any                  `json:"errors" example:"null"`
}

type CallbackData struct {
	Status    PaymentStatus `json:"status"`
	Authority string        `json:"authority"`
	InvoiceID string        `json:"invoice_id"`
}

type InitTransactionData struct {
	Wage       int    `json:"wage" example:"0"`
	WagePayer  string `json:"wage_payer" example:"merchant"`
	Authority  string `json:"authority" example:"812F739E41057BAC22331918CD5B41C2"`
	TransID    int    `json:"trans_id" example:"337811"`
	PaymentUrl string `json:"payment_url" example:"https://ipg.novinopay.com/StartPay/812F739E41057BAC22331918CD5B41C2"`
}

type VerifyTransactionRequest struct {
	MerchantID string `json:"merchant_id" example:"test"`
	Amount     int    `json:"amount" example:"10000"`
	Authority  string `json:"authority" example:"812F739E41057BAC22331918CD5B41C2"`
}

type VerifyTransactionResponse struct {
	Status  string                 `json:"status" example:"100"`
	Message string                 `json:"message" example:"عمليات موفق"`
	Data    *VerifyTransactionData `json:"data"`
	Errors  any                    `json:"errors" example:"null"`
}

type VerifyTransactionData struct {
	TransID     int    `json:"trans_id" example:"337811"`
	RefID       string `json:"ref_id" example:"223003535268"`
	Authority   string `json:"authority" example:"812F739E41057BAC22331918CD5B41C2"`
	CardPan     string `json:"card_pan" example:"504172******0613"`
	Amount      int    `json:"amount" example:"1000"`
	InvoiceID   any    `json:"invoice_id" example:"null"`
	BuyerIP     string `json:"buyer_ip" example:"5.113.185.222"`
	PaymentTime int64  `json:"payment_time" example:"1663376322"`
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Err) Error() string {
	return e.Message
}

func addInvoice(invoice Invoice) {
	invoiceStorage[invoice.InvoiceID] = invoice
}
