package main

import "hexagonal-architecture/pkg/http/rest"

func main() {

	app := rest.InitHandlers()

	app.Start(":8080")
}
