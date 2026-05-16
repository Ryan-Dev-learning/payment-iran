package app

import (
	"net/http"

	"github.com/Ryan-Dev-learning/payment-iran/config"
	"github.com/Ryan-Dev-learning/payment-iran/novinopay"
	"github.com/Ryan-Dev-learning/payment-iran/server"
	"github.com/Ryan-Dev-learning/payment-iran/server/handlers"
	"github.com/Ryan-Dev-learning/payment-iran/server/routes"
)

type App struct {
	cfg *config.Config
	srv *server.Server
}

func New() *App {
	cfg, err := config.Load("config.toml")
	if err != nil {
		panic(err)
	}

	return &App{
		cfg: &cfg,
	}
}

func (a *App) Setup() {
	//Setup Server
	a.srv = server.New(&a.cfg.Server)
	//Setup Services
	client := &http.Client{}
	novinopaySvc := novinopay.NewSvc(&a.cfg.NovinoPay, client)

	//Setup Handlers
	handler := handlers.NewHandlers(novinopaySvc)

	//Setup Routes
	routes.Setup(a.srv.Router, handler)

}
func (a *App) Run() error {
	return a.srv.Start()
}
