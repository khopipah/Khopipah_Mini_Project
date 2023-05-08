package main

import (
	"khopipah_mini_project-1/config"
	//m "khopipah_mini_project_structure/middlewares"
	"khopipah_mini_project-1/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	//m.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":8000"))
}
