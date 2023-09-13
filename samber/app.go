package samber

import (
	"fmt"

	"github.com/samber/do"

	"github.com/albdewilde/tdi/samber/locationfetcher"
	"github.com/albdewilde/tdi/samber/userfetcher"
)

func init() {
	uf := do.MustInvoke[*userfetcher.UserFetch](nil)
	lf := do.MustInvoke[*locationfetcher.LocationFetch](nil)

	do.Provide(
		nil,
		func(i *do.Injector) (*App, error) {
			return NewApp(uf, lf), nil
		},
	)
}

type App struct {
	userFetcher     UserFetcher
	locationFetcher LocationFetcher
}

type UserFetcher interface {
	Fetch(ID string) userfetcher.User
}

type LocationFetcher interface {
	Fetch(ID string) locationfetcher.Location
}

func NewApp(userFetcher UserFetcher, locationFetcher LocationFetcher) *App {
	var app App

	app.userFetcher = userFetcher
	app.locationFetcher = locationFetcher

	return &app
}

func (a *App) DoStuff() {
	fmt.Println("fetching an user")
	u := a.userFetcher.Fetch("28")
	fmt.Printf("here is the user: %#v\n", u)

	fmt.Println("fetching a location")
	l := a.locationFetcher.Fetch("37")
	fmt.Printf("here is the location: %#v\n", l)
}
