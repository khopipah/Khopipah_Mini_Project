package main

import (
	"khopipah_mini_project/config"
	"khopipah_mini_project/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
