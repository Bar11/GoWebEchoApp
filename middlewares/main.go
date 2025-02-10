package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var validToken = "valid-token"

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		autHeader := c.Request().Header.Get("Authorization")
		if autHeader == "" {
			//c.Redirect(http.StatusMovedPermanently, setting.Settings.LoginUrl)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error":  "Missing Authorization header",
				"status": http.StatusUnauthorized,
			})
		}

		tokenParts := strings.Split(autHeader, " ")
		fmt.Println(tokenParts[0], len(tokenParts))
		if tokenParts[0] != validToken {
			//c.Redirect(http.StatusMovedPermanently, setting.Settings.LoginUrl)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error":  "Invalid Authorization header",
				"status": http.StatusUnauthorized,
			})

		}

		return next(c)

	}
}
