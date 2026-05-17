package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/Ryan-Dev-learning/payment-iran/config"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(c *echo.Context, w io.Writer, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Server struct {
	Router *echo.Echo
	config *config.ServerConfig
}

func New(config *config.ServerConfig) *Server {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}
	e.Renderer = t

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
