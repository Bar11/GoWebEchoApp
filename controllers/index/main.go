package index

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
