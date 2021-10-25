package controllers

import (
	"mvc/lib/database"
	"mvc/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginUsersController(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}
