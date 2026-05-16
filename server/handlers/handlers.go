package handlers

import (
	"github.com/Ryan-Dev-learning/payment-iran/novinopay"
	"github.com/labstack/echo/v5"
)

type Handlers struct {
	novinoPaySvc novinopay.Service
}

func NewHandlers(novinoPaySvc novinopay.Service) *Handlers {
	return &Handlers{
		novinoPaySvc: novinoPaySvc,
	}
}

// GreetingHandler godoc
// @Summary Greeting
// @Description get greeting message
// @Tags greeting
// @Accept  json
// @Produce  plain
// @Success 200 {string} string "Hello, World!"
// @Router /greeting [get]
func (h *Handlers) GreetingHandler(c *echo.Context) (err error) {
	return c.String(200, "Hello, World!")
}
