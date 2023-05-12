package main

import (
	"khopipah_mini_project/config"
	"khopipah_mini_project/middleware"
	"khopipah_mini_project/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
