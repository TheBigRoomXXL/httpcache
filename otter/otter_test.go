package otter

import (
	"testing"

	"github.com/TheBigRoomXXL/httpcache/core/test"
	"github.com/maypok86/otter/v2"
)

func TestOtterCache(t *testing.T) {
	cache := otter.Must(&otter.Options[string, []byte]{
		MaximumSize:     10_000,
		InitialCapacity: 1_000,
	})
	test.Cache(t, New(cache))
}
