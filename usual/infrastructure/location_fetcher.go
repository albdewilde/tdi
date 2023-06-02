package infrastructure

import "github.com/albdewilde/tdi/usual/domain"

// LocationFetch implement the usual.LocationFetcher interface to be used by the usual.App structure
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

func (uf *LocationFetch) Fetch(ID string) domain.Location {
	return domain.Location{ID: ID}
}
