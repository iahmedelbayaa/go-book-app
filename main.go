package main

import (
	"net/http"

	"github.com/iahmedelbayaa/go-book-app/config"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}