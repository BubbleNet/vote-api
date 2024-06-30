package server

import (
	"github.com/BubbleNet/vote-api/pkg/health"
	"github.com/labstack/echo/v4"
)

func CreateServer() *echo.Echo {
	e := echo.New()

	// Handlers
	healthHandler := health.NewHandler()

	// Routes
	e.GET("/health", healthHandler.GetHealth)
	return e
}

func Run() {
	e := CreateServer()
	e.Logger.Fatal(e.Start(":8080"))
}
