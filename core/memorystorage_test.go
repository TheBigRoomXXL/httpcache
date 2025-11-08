package httpcache_test

import (
	"testing"

	httpcache "pkg.lovergne.dev/httpcache/core"
	"pkg.lovergne.dev/httpcache/storagetest"
)

func TestMemoryCache(t *testing.T) {
	storagetest.StorageLifecycle(t, httpcache.NewInMemoryStorage())
}
