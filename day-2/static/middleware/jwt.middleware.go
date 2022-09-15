package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(id string) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"userId":     id,
		"exp":        time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func ExtractToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(int)
		return id
	}
	return 0
}
