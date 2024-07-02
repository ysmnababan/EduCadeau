package helper

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ResponseRecorder wraps echo.ResponseWriter to capture the status code
type ResponseRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *ResponseRecorder) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

// Logging function to log request details
func Logging(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	// Use response recorder to get the status code
	rr, ok := c.Response().Writer.(*ResponseRecorder)
	if !ok {
		rr = &ResponseRecorder{
			ResponseWriter: c.Response().Writer,
			Status:         http.StatusOK, // default to 200
		}
	}

	return logrus.WithFields(logrus.Fields{
		"method": c.Request().Method,
		"status": rr.Status,
		"uri":    c.Request().URL.String(),
	})
}
