package main

import (
	m "static/middleware"
	"static/routes"
)

func main() {
	app := routes.App()

	m.LogMiddleware(app)

	app.Logger.Fatal(app.Start(":8080"))
}
