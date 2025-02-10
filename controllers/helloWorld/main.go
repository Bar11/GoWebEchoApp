package helloWorld

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SayHello(c echo.Context) error {
	name := c.Param("name")
	hello := "Hello, " + name
	return c.String(http.StatusOK, hello)
}
