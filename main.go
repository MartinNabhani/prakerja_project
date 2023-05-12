package main

import (
	"prakerja/configs"
	"prakerja/routes"
)

func main() {
	configs.InitDB()
	e := routes.InitRoute()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
