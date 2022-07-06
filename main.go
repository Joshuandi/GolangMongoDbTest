package main

import (
	"GolangMongoDbTest/config"
	"GolangMongoDbTest/router"

	"github.com/labstack/echo/v4"
)

func main() {
	//use framework echo
	e := echo.New()
	//run database
	config.Connect()

	//router
	router.UserRouter(e)

	e.Logger.Fatal(e.Start(":8088"))
}
