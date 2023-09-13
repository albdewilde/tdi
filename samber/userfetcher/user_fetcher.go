package userfetcher

import (
	"github.com/samber/do"
)

func init() {
	do.Provide(
		nil,
		func(i *do.Injector) (*UserFetch, error) {
			return NewUserFetch("example.com", "key"), nil
		},
	)
}

// UserFetch implement the samber.UserFetcher interface to be used by the samber.App structure
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

func (uf *UserFetch) Fetch(ID string) User {
	return User{ID: ID}
}
