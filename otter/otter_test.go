package otter

import (
	"testing"

	"github.com/maypok86/otter/v2"
	"pkg.lovergne.dev/httpcache/core/test"
)

func TestRedisCache(t *testing.T) {
	cache := otter.Must(&otter.Options[string, []byte]{
		MaximumSize:     10_000,
		InitialCapacity: 1_000,
	})
	test.Cache(t, New(cache))
}
