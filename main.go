package main

import (
	"os"
	"prakerja/configs"
	"prakerja/routes"
)

func main() {
	configs.InitDB()
	e := routes.InitRoute()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
