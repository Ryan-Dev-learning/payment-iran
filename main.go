package main

import (
	"github.com/Ryan-Dev-learning/payment-iran/app"
	_ "github.com/Ryan-Dev-learning/payment-iran/docs"
)

// @title Payment Iran API
// @version 1.0
// @description Payment integration server for Iran gateways.
// @host localhost:9091
// @BasePath /api/v1
func main() {
	appInstance := app.New()
	appInstance.Setup()
	appInstance.Run()
}
