package endpoints

import (
	"github.com/labstack/echo"
	"net/http"
)

// Health Check - Health check endpoint
func HC(c echo.Context) error {
	resp := Response{Message: "OK!", Status: http.StatusOK}
	return c.JSON(resp.Status, resp)
}

