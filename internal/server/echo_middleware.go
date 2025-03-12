package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// getRequestIdFromHeader will read a header and get the request ID
func getRequestIdFromHeader(req *http.Request, res *echo.Response) string {
	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}
	return id
}

// LoggerMiddleware is a middleware that logs HTTP requests via logrus.
func loggerMiddleware(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			var errorMessage string
			if err = next(c); err != nil {
				c.Error(err)
				errorMessage = err.Error()
			}
			stop := time.Now()

			id := getRequestIdFromHeader(c.Request(), c.Response())
			reqSize := req.Header.Get(echo.HeaderContentLength)
			if reqSize == "" {
				reqSize = "0"
			}
			attrs := []slog.Attr{
				slog.String("id", id),
				slog.String("remote-ip", c.RealIP()),
				slog.String("host", req.Host),
				slog.String("method", req.Method),
				slog.String("uri", req.RequestURI),
				slog.Int("status", res.Status),
				slog.String("error", errorMessage),
				slog.String("size", reqSize),
				slog.String("latency", stop.Sub(start).String()),
				slog.String("referer", req.Referer()),
				slog.String("user-agent", req.UserAgent()),
			}

			n := res.Status
			switch {
			case n >= 500:
				logger.LogAttrs(nil, slog.LevelError, "request failed", attrs...)
			case n >= 400:
				logger.LogAttrs(nil, slog.LevelWarn, "client error", attrs...)
			case n >= 300:
				logger.LogAttrs(nil, slog.LevelInfo, "redirection", attrs...)
			default:
				logger.LogAttrs(nil, slog.LevelInfo, "request handled", attrs...)
			}

			return err
		}
	}
}
