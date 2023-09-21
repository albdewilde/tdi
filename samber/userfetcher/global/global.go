package global

import (
	"github.com/albdewilde/tdi/samber/userfetcher"
	"github.com/samber/do"
)

func init() {
	do.Provide(
		nil,
		func(i *do.Injector) (*userfetcher.UserFetch, error) {
			return userfetcher.NewUserFetch("example.com", "key"), nil
		},
	)
}
