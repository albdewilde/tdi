package fake

import (
	"fmt"

	"github.com/albdewilde/tdi/samber/userfetcher"
	"github.com/samber/do"
)

func init() {
	do.Provide(
		nil,
		func(i *do.Injector) (*FakeUserFetch, error) {
			return &FakeUserFetch{}, nil
		},
	)
}

type FakeUserFetch struct{}

func (fuf *FakeUserFetch) Fetch(ID string) userfetcher.User {
	return userfetcher.User{ID: fmt.Sprintf("fake %s", ID)}
}
