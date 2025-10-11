package test

import (
	"bytes"
	"context"
	"testing"

	"pkg.lovergne.dev/httpcache"
)

// Cache excercises a httpcache.Cache implementation.
func Cache(t *testing.T, cache httpcache.Cache) {
	CacheWithHook(t, cache, func() {})
}

// Same as Cache with hook added to sync eventually concistant cache
func CacheWithHook(t *testing.T, cache httpcache.Cache, afterSet func()) {
	key := "testKey"
	_, ok := cache.Get(context.Background(), key)
	if ok {
		t.Fatal("retrieved key before adding it")
	}

	val := []byte("some bytes")
	cache.Set(context.Background(), key, val)

	// Hack to that we are sure not to fail reading on eventually concistant cache
	afterSet()

	retVal, ok := cache.Get(context.Background(), key)
	if !ok {
		t.Fatal("could not retrieve an element we just added")
	}
	if !bytes.Equal(retVal, val) {
		t.Fatal("retrieved a different value than what we put in")
	}

	cache.Delete(context.Background(), key)

	_, ok = cache.Get(context.Background(), key)
	if ok {
		t.Fatal("deleted key still present")
	}
}
