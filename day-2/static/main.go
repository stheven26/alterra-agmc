package main

import (
	"static/routes"
)

func main() {
	app := routes.App()

	app.Start(":8080")
}
