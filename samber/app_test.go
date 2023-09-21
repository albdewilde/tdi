package samber

import (
	"testing"

	flf "github.com/albdewilde/tdi/samber/locationfetcher/fake"
	fuf "github.com/albdewilde/tdi/samber/userfetcher/fake"
	"github.com/samber/do"
)

func TestOurApp(t *testing.T) {
	app := NewApp(
		do.MustInvoke[*fuf.FakeUserFetch](nil),
		do.MustInvoke[*flf.FakeLocationFetch](nil),
	)

	app.DoStuff()
}
