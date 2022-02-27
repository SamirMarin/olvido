package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	resp := HealthCheckResponse{
		Message: "Checked and verified!",
	}
	return c.JSON(http.StatusOK, resp)
}
