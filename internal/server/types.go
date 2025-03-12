package server

import (
	"github.com/labstack/echo/v4"
)

// HTTPServer is an interface which provides method signatures for an HTTP server
type HTTPServer interface {
	Instance() *echo.Echo
	Start(port string)
	Shutdown()
}
