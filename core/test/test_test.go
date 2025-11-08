package test_test

import (
	"testing"

	httpcache "pkg.lovergne.dev/httpcache/core"
	"pkg.lovergne.dev/httpcache/core/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewInMemoryStorage())
}
