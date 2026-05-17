package handlers

import (
	"net/http"
	"time"

	"github.com/Ryan-Dev-learning/payment-iran/novinopay"
	"github.com/labstack/echo/v5"
)

// InitTransactionHandler godoc
// @Summary Init Transaction
// @Description Initialize a new payment transaction with NovinoPay
// @Tags novinopay
// @Accept  json
// @Produce  json
// @Param request body novinopay.InitTransactionRequest true "Transaction Request"
// @Success 200 {object} novinopay.InitTransactionResponse
// @Failure 400 {object} novinopay.Err
// @Failure 500 {object} novinopay.Err
// @Router /novinopay/init [post]
func (h *Handlers) InitTransactionHandler(c *echo.Context) error {
	var request novinopay.InitTransactionRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	response, err := h.novinoPaySvc.InitTransaction(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// VerifyTransactionHandler godoc
// @Summary Verify Transaction
// @Description Verify a payment transaction with NovinoPay
// @Tags novinopay
// @Accept  json
// @Produce  json
// @Param request body novinopay.VerifyTransactionRequest true "Verification Request"
// @Success 200 {object} novinopay.VerifyTransactionResponse
// @Failure 400 {object} novinopay.Err
// @Failure 500 {object} novinopay.Err
// @Router /novinopay/verify [post]
func (h *Handlers) VerifyTransactionHandler(c *echo.Context) error {
	var request novinopay.VerifyTransactionRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	resp, err := h.novinoPaySvc.VerifyTransaction(request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handlers) InitTransactionCallbackHandler(c *echo.Context) error {
	var call novinopay.CallbackData
	authority := c.QueryParam("Authority")
	if authority != "" {
		call.Authority = authority
	}

	invoiceID := c.QueryParam("InvoiceID")
	if invoiceID != "" {
		call.InvoiceID = invoiceID
	}

	status := c.QueryParam("PaymentStatus")
	if status != "" {
		call.Status = novinopay.PaymentStatus(status)
	}

	invoice, err := h.novinoPaySvc.GetInvoice(call.InvoiceID)
	if err != nil {
		return err
	}

	data := map[string]any{
		"Callback":    call,
		"Invoice":     invoice,
		"CurrentTime": time.Now().Format("2006-01-02 15:04:05"),
	}

	return c.Render(http.StatusOK, "callback.html", data)
}
