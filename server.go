package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	go scrape() // start a goroutine
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Heilo, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func scrape() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("looep")
	}
}
