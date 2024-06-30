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

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{Status: "OK"})
}
