package main

import (
	"olvido/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
	// get notification to send
	go handlers.scraper()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
