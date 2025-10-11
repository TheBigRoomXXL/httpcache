package test_test

import (
	"testing"

	"pkg.lovergne.dev/httpcache"
	"pkg.lovergne.dev/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
