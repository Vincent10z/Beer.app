package middleware

import (
	"Beer.app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "missing or invalid token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token format"})
		}

		token := parts[1]
		userID, err := utils.VerifyJWT(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid token"})
		}

		c.Set("userID", userID)
		return next(c)
	}
}
