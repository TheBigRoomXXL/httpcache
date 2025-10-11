package ristretto

import (
	"testing"

	"github.com/dgraph-io/ristretto/v2"
	"pkg.lovergne.dev/httpcache/test"
)

func TestRedisCache(t *testing.T) {
	cache, err := ristretto.NewCache(&ristretto.Config[string, []byte]{
		NumCounters: 1e6,               // number of keys to track frequency of (1M).
		MaxCost:     100 * 1024 * 1024, // 100MB
		BufferItems: 64,                // number of keys per Get buffer.
	})
	if err != nil {
		t.Fatal(err)
	}

	sync := func() {
		cache.Wait()
	}
	test.CacheWithHook(t, New(cache), sync)
}
