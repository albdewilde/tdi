package infrastructure

import "github.com/albdewilde/tdi/usual/domain"

// LocationFetch implement the usual.LocationFetcher interface to be used by the usual.App structure
type LocationFetch struct {
	host string
	key  string
}

func NewLocationFetch(host, key string) *LocationFetch {
	var userFetch LocationFetch
	userFetch.host = host
	userFetch.key = key
	return &userFetch
}

func (uf *LocationFetch) Fetch(ID string) domain.Location {
	return domain.Location{}
}
