package statuserr_test

import (
	"errors"
	"testing"

	"github.com/sheki/statuserr"
)

type statusCode interface {
	Status() int
}

func TestBasic(t *testing.T) {
	t.Parallel()
	err := statuserr.New(400, "broken")
	if sc, ok := err.(statusCode); ok {
		if sc.Status() != 400 {
			t.Errorf("status code interface not implemented")
		}
	}
}

func TestWrap(t *testing.T) {
	t.Parallel()
	err := statuserr.BadGateway
	if err.Error() != "Bad Gateway" {
		t.Fatal("bad gatway does not have correct error string")
	}
	newErr := err.Wrap(errors.New("not bad gateway"))
	if newErr.Error() != "not bad gateway" {
		t.Fatalf("wrap has %s\n", newErr.Error())

	}
}
