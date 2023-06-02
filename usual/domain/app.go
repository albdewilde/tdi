package domain

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
