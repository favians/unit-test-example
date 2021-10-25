package controllers

import (
	"mvc/lib/database"
	"mvc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserByIDControllers(c echo.Context) error {
	id := c.QueryParam("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	users, err := database.GetSingleUser(idInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserControllers(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func InsertUserControllers(c echo.Context) error {
	user := models.Users{}

	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = database.InsertUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
