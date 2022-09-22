package main

import (
	"hexagonal-architecture/internal/app/route"
	m "hexagonal-architecture/internal/middlewares"
)

func main() {
	app := route.Init()

	m.LogMiddleware(app)

	app.Start(":8080")
}
