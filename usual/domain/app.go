package domain

import "fmt"

type App struct {
	userFetcher     UserFetcher
	locationFetcher LocationFetcher
}

type UserFetcher interface {
	Fetch(ID string) User
}

type LocationFetcher interface {
	Fetch(ID string) Location
}

func NewApp(userFetcher UserFetcher, locationFetcher LocationFetcher) App {
	var app App

	app.userFetcher = userFetcher
	app.locationFetcher = locationFetcher

	return app
}

func (a *App) DoStuff() {
	fmt.Println("fetching an user")
	u := a.userFetcher.Fetch("38")
	fmt.Printf("here is the user: %#v\n", u)

	fmt.Println("fetching a location")
	l := a.locationFetcher.Fetch("45")
	fmt.Printf("here is the location: %#v\n", l)
}
