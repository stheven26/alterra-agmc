package main

import (
	m "alterra-agmc-day2/middlewares"
	"alterra-agmc-day2/routes"
)

func main() {
	app := routes.Init()

	m.LogMiddleware(app)

	app.Start(":8080")
}
