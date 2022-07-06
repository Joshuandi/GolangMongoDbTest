package router

import (
	"GolangMongoDbTest/controller"

	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo) {
	e.POST("/users", controller.InsertUser)
	e.GET("/users/:id", controller.GetUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.GET("/users", controller.GetAllUser)
}
