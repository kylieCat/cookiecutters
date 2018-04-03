package endpoints

import (
	"github.com/labstack/echo"
	"net/http"
)

type Endpoint struct {
	Url string
	Method string
	Handler echo.HandlerFunc
}

type Response struct {
	Message string
	Status  int
}

func InitEndpoints() []Endpoint {
	endpoints := []Endpoint {
		{"/hc", http.MethodGet, HC},
		{"/hc", http.MethodGet, HelloWorld},
	}
	return endpoints
}
