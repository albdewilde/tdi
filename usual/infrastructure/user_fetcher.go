package infrastructure

import "github.com/albdewilde/tdi/usual/domain"

// UserFetch implement the usual.UserFetcher interface to be used by the usual.App structure
type UserFetch struct {
	host string
	key  string
}

func NewUserFetch(host, key string) *UserFetch {
	var userFetch UserFetch
	userFetch.host = host
	userFetch.key = key
	return &userFetch
}

func (uf *UserFetch) Fetch(ID string) domain.User {
	return domain.User{ID: ID}
}
