package routes

import (
	"github.com/Ryan-Dev-learning/payment-iran/server/handlers"
	"github.com/labstack/echo/v5"
)

func setupNovinoPayRoutes(group *echo.Group, h *handlers.Handlers) {
	group.POST("/novinopay/init", h.InitTransactionHandler)
	group.POST("/novinopay/verify", h.VerifyTransactionHandler)

	// CallBack
	group.GET("/novinopay/callback", h.InitTransactionCallbackHandler)
}
