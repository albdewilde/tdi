package locationfetcher

import (
	"github.com/samber/do"
)

func init() {
	do.Provide(
		nil,
		func(i *do.Injector) (*LocationFetch, error) {
			return NewLocationFetch("location.com", "key"), nil
		},
	)
}

// LocationFetch implement the samber.LocationFetcher interface to be used by the samber.App structure
type LocationFetch struct {
	host string
	key  string
}

func NewLocationFetch(host, key string) *LocationFetch {
	var locationFetcher LocationFetch
	locationFetcher.host = host
	locationFetcher.key = key
	return &locationFetcher
}

func (uf *LocationFetch) Fetch(ID string) Location {
	return Location{ID: ID}
}
