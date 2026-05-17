package novinopay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ryan-Dev-learning/payment-iran/config"
)

type Service interface {
	InitTransaction(request InitTransactionRequest) (InitTransactionResponse, error)
	VerifyTransaction(request VerifyTransactionRequest) (VerifyTransactionResponse, error)
	GetInvoice(invoiceID string) (Invoice, error)
}

type Svc struct {
	cfg    *config.Config
	client *http.Client
}

func NewSvc(cfg *config.Config, client *http.Client) *Svc {
	return &Svc{
		cfg:    cfg,
		client: client,
	}
}

func (s *Svc) InitTransaction(request InitTransactionRequest) (InitTransactionResponse, error) {
	if request.CallbackUrl == "" {
		request.CallbackUrl = fmt.Sprintf("http://%s:%d%s", s.cfg.Server.Host, s.cfg.Server.Port, s.cfg.NovinoPay.CallbackUrl)
	}
	if request.MerchantID == "" {
		request.MerchantID = s.cfg.NovinoPay.MerchantID
	}

	invoice := Invoice{
		InvoiceID:   request.InvoiceID,
		Amount:      request.Amount,
		Description: request.Description,
		Email:       request.Email,
		Mobile:      request.Mobile,
		Name:        request.Name,
		CardPan:     request.CardPan,
	}
	addInvoice(invoice)

	body, err := json.Marshal(request)
	if err != nil {
		return InitTransactionResponse{}, &Err{Message: err.Error()}
	}

	req, err := http.NewRequest("POST", s.cfg.NovinoPay.InitTransactionUrl, bytes.NewReader(body))
	if err != nil {
		return InitTransactionResponse{}, &Err{Message: err.Error()}
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return InitTransactionResponse{}, &Err{Message: err.Error()}
	}
	defer resp.Body.Close()

	var response InitTransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return InitTransactionResponse{}, &Err{Message: err.Error()}
	}
	return response, nil
}

func (s *Svc) VerifyTransaction(request VerifyTransactionRequest) (VerifyTransactionResponse, error) {
	if request.MerchantID == "" {
		request.MerchantID = s.cfg.NovinoPay.MerchantID
	}

	body, err := json.Marshal(request)
	if err != nil {
		return VerifyTransactionResponse{}, &Err{Message: err.Error()}
	}
	req, err := http.NewRequest("POST", s.cfg.NovinoPay.VerifyTransactionUrl, bytes.NewReader(body))
	if err != nil {
		return VerifyTransactionResponse{}, &Err{Message: err.Error()}
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return VerifyTransactionResponse{}, &Err{Message: err.Error()}
	}

	defer resp.Body.Close()
	
	var response VerifyTransactionResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return VerifyTransactionResponse{}, &Err{Message: err.Error()}
	}

	return response, nil
}

func (s *Svc) GetInvoice(invoiceID string) (Invoice, error) {
	if invoiceID == "" {
		return Invoice{}, &Err{Message: "InvoiceID is required"}
	}
	if _, ok := invoiceStorage[invoiceID]; !ok {
		return Invoice{}, &Err{Message: "Invoice not found"}
	}

	return invoiceStorage[invoiceID], nil
}
