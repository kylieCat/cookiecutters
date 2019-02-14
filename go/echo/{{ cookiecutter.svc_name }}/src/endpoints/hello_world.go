package endpoints

import (
	"net/http"
	"github.com/labstack/echo"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/logger"
)


// HelloWorld - Accepts one query parameter `name` that defaults to `"World"` if not passed in.
// Returns a greeting.
func HelloWorld(c echo.Context) error {
	logger.Info("Got request!")
	name := c.QueryParam("name")
	if name == "" {
		name = "World"
	}
	resp := Response{
		Message: "Hello, " + name,
		Status: http.StatusOK,
	}
	return c.JSON(resp.Status, resp)
}
