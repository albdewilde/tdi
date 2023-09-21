package domain

import (
	"fmt"
	"testing"
)

type fakeUserFetcher struct{}

func (fuf *fakeUserFetcher) Fetch(ID string) User {
	return User{ID: fmt.Sprintf("fake %s", ID)}
}

type fakeLocationFetcher struct{}

func (fuf *fakeLocationFetcher) Fetch(ID string) Location {
	return Location{ID: fmt.Sprintf("fake %s", ID)}
}

func TestOurApp(t *testing.T) {
	app := NewApp(
		&fakeUserFetcher{},
		&fakeLocationFetcher{},
	)

	app.DoStuff()
}
