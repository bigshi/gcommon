package gutil

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"testing"
)

func TestULID(t *testing.T) {
	id := ulid.Make()
	fmt.Println(id.String())
}
