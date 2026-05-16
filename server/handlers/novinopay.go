package handlers

import (
	"net/http"

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
