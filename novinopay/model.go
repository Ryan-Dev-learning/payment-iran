package novinopay

type PaymentStatus string

const (
	OK  PaymentStatus = "OK"
	NOK PaymentStatus = "NOK"
)

type InitTransactionRequest struct {
	MerchantID     string `json:"merchant_id"`
	Amount         int    `json:"amount"`
	CallbackUrl    string `json:"callback_url"`
	CallbackMethod string `json:"callback_method"`
	InvoiceID      string `json:"invoice_id"`
	Description    string `json:"description"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Name           string `json:"name"`
	CardPan        string `json:"card_pan"`
}

type InitTransactionResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    *InitTransactionData `json:"data"`
	Errors  any                  `json:"errors"`
}

type InitTransactionData struct {
	Wage       int    `json:"wage"`
	WagePayer  string `json:"wage_payer"`
	Authority  string `json:"authority"`
	TransID    int    `json:"trans_id"`
	PaymentUrl string `json:"payment_url"`
}

type VerifyTransactionRequest struct {
	MerchantID string `json:"merchant_id"`
	Amount     int    `json:"amount"`
	Authority  string `json:"authority"`
}

type VerifyTransactionResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    *VerifyTransactionData `json:"data"`
	Errors  any                    `json:"errors"`
}

type VerifyTransactionData struct {
	TransID     int    `json:"trans_id"`
	RefID       string `json:"ref_id"`
	Authority   string `json:"authority"`
	CardPan     string `json:"card_pan"`
	Amount      int    `json:"amount"`
	InvoiceID   any    `json:"invoice_id"`
	BuyerIP     string `json:"buyer_ip"`
	PaymentTime int64  `json:"payment_time"`
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Err) Error() string {
	return e.Message
}
