package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler        struct{}
	HealthResponse struct {
		Status string `json:"status"`
	}
)

// NewHandler Creates a handler for responding to health checks
func NewHandler() Handler {
	return Handler{}
}

// GetHealth - GET - Responds with the current status of the api server
func (h *Handler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{Status: "OK"})
}
