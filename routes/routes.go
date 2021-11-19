package routes

import (
	"mvc/constant"
	"mvc/controllers"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUsersController)

	auth := e.Group("/users")
	auth.Use(echoMiddleware.JWT([]byte(constant.JWTSecret)))
	auth.GET("", controllers.GetUserByIDControllers)
	auth.GET("/all", controllers.GetUserControllers)
	auth.POST("", controllers.InsertUserControllers)
	auth.PUT("", controllers.UpdateUserControllers)

	return e
}
