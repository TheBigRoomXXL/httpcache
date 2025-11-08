package otter_test

import (
	"testing"

	"github.com/maypok86/otter/v2"
	storageotter "pkg.lovergne.dev/httpcache/otter"
	"pkg.lovergne.dev/httpcache/storagetest"
)

func TestOtterCache(t *testing.T) {
	cache := otter.Must(&otter.Options[string, []byte]{
		MaximumSize:     10_000,
		InitialCapacity: 1_000,
	})
	storagetest.StorageLifecycle(t, storageotter.New(cache))
}
