package main

import (
	"github.com/ramadhanalfarisi/go-simple-dockerizing/app"
)

func main() {
	var a app.App
	a.CreateConnection()
	// a.Migrate()
	a.CreateRoutes()
	a.Run()
}