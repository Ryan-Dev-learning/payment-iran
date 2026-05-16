package novinopay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ryan-Dev-learning/payment-iran/config"
)

type Service interface {
	InitTransaction(request InitTransactionRequest) (InitTransactionResponse, error)
	VerifyTransaction(request VerifyTransactionRequest) (VerifyTransactionResponse, error)
}

type Svc struct {
	cfg    *config.NovinoPayConfig
	client *http.Client
}

func NewSvc(cfg *config.NovinoPayConfig, client *http.Client) *Svc {
	return &Svc{
		cfg:    cfg,
		client: client,
	}
}

func (s *Svc) InitTransaction(request InitTransactionRequest) (InitTransactionResponse, error) {
	if request.CallbackUrl == "" {
		request.CallbackUrl = fmt.Sprintf("http://localhost%s", s.cfg.CallbackUrl)
	}
	if request.MerchantID == "" {
		request.MerchantID = s.cfg.MerchantID
	}
	log.Println(request)

	body, err := json.Marshal(request)
	if err != nil {
		return InitTransactionResponse{}, &Err{Message: err.Error()}
	}

	req, err := http.NewRequest("POST", s.cfg.InitTransactionUrl, bytes.NewReader(body))
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
		request.MerchantID = s.cfg.MerchantID
	}

	body, err := json.Marshal(request)
	if err != nil {
		return VerifyTransactionResponse{}, &Err{Message: err.Error()}
	}
	req, err := http.NewRequest("POST", s.cfg.VerifyTransactionUrl, bytes.NewReader(body))
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
