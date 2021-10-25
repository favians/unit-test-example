package middleware

import (
	"mvc/config"
	"mvc/models"

	"github.com/labstack/echo"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.Users

	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}

	return true, nil
}
