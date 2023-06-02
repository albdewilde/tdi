package main

import (
	"github.com/albdewilde/tdi/usual/domain"
	"github.com/albdewilde/tdi/usual/infrastructure"
)

func main() {
	// Usual way to initialize out app
	userFetch := infrastructure.NewUserFetch("example.com", "key")
	locFetch := infrastructure.NewLocationFetch("locaiton.com", "key")

	app := domain.NewApp(userFetch, locFetch)

	// Do stuff with our app
	app.DoStuff()

	// Initialize the app with samber/do
}
