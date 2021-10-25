package middleware

import (
	"mvc/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constant.JWTSecret))
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(int)
		return userId
	}

	return 0
}
