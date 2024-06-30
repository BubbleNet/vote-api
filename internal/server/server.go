package server

import (
	"github.com/BubbleNet/vote-api/internal/health"
	"github.com/labstack/echo/v4"
)

// CreateServer creates the server object, adds middleware, and routes.
func CreateServer() *echo.Echo {
	e := echo.New()

	// Handlers
	healthHandler := health.NewHandler()

	// Routes
	e.GET("/health", healthHandler.GetHealth)
	return e
}

// Run starts the server.
func Run() {
	e := CreateServer()
	e.Logger.Fatal(e.Start(":8080"))
}
