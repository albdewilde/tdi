package main

import (
	"fmt"

	"github.com/samber/do"

	usualdomain "github.com/albdewilde/tdi/usual/domain"
	usualinfra "github.com/albdewilde/tdi/usual/infrastructure"

	"github.com/albdewilde/tdi/samber"
)

func main() {
	// Usual way to initialize out app
	userFetch := usualinfra.NewUserFetch("example.com", "key")
	locFetch := usualinfra.NewLocationFetch("location.com", "key")

	usualApp := usualdomain.NewApp(userFetch, locFetch)

	// Do stuff with our app
	usualApp.DoStuff()

	fmt.Println("---")

	// Initialize the app with samber/do
	samberApp := do.MustInvoke[*samber.App](nil)

	// Do stuff with our app
	samberApp.DoStuff()
}
