package routes

import (
	"github.com/Ryan-Dev-learning/payment-iran/server/handlers"
	"github.com/labstack/echo/v5"
	echoSwagger "github.com/swaggo/echo-swagger/v2"
)

func Setup(router *echo.Echo, handler *handlers.Handlers) {
	setupV1Routes(router, handler)
}
func setupV1Routes(router *echo.Echo, h *handlers.Handlers) {
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := router.Group("/api/v1")
	v1.GET("/greeting", h.GreetingHandler)
	setupNovinoPayRoutes(v1, h)
}
