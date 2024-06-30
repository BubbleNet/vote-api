package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Split off a testable chunk
func CreateServer() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "yooooooo")
	})
	return e
}

func Run() {
	e := CreateServer()
	e.Logger.Fatal(e.Start(":8080"))
}
