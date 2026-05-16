package server

import (
	"fmt"
	"net/http"

	"github.com/Ryan-Dev-learning/payment-iran/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Server struct {
	Router *echo.Echo
	config *config.ServerConfig
}

func New(config *config.ServerConfig) *Server {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLogger())

	return &Server{
		config: config,
		Router: e,
	}
}

func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	return s.Router.Start(address)
}
