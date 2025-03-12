package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// echoServer is a wrapper struct for *echo.Echo
type echoServer struct {
	Server *echo.Echo
	logger *slog.Logger
}

var (
	_ HTTPServer = (*echoServer)(nil)
)

// New returns a new instance of an echo HTTP server
func New(logger *slog.Logger) HTTPServer {
	return &echoServer{
		Server: echo.New(),
		logger: logger,
	}
}

// Instance returns a singleton echo instance
func (s *echoServer) Instance() *echo.Echo {
	return s.Server
}

// Setup the web server
func (s *echoServer) Start(port string) {
	e := s.Instance()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(loggerMiddleware(s.logger))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	s.logger.Info("Starting API server", slog.String("address", ":8080"))

	e.Logger.Fatal(e.Start(":" + port))
}

// Shutdown the web server
func (s *echoServer) Shutdown() {
	s.Instance().Shutdown(context.Background())
}
