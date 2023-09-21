package fake

import (
	"fmt"

	"github.com/albdewilde/tdi/samber/locationfetcher"
	"github.com/samber/do"
)

func init() {
	do.Provide(
		nil,
		func(i *do.Injector) (*FakeLocationFetch, error) {
			return &FakeLocationFetch{}, nil
		},
	)
}

type FakeLocationFetch struct{}

func (flf *FakeLocationFetch) Fetch(ID string) locationfetcher.Location {
	return locationfetcher.Location{ID: fmt.Sprintf("fake %s", ID)}
}
