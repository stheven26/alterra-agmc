package main

import "alterra-agmc-day2/routes"

func main() {
	app := routes.Init()

	app.Start(":8080")
}
