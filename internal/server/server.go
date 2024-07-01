package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

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

// Run starts the server and handles server shutdown
func Run(ctx context.Context, getenv func(string) string) {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	e := CreateServer()

	go func() {
		if err := e.Start(getenv("PORT")); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
