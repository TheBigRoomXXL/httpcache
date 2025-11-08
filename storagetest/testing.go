package storagetest

import (
	"bytes"
	"context"
	"testing"

	"pgregory.net/rapid"
)

// We copy paste the storage interface here so that we don't have dependencies to the core
// This simplify versioning and release.
type Storage interface {
	Get(ctx context.Context, key string) (responseBytes []byte, ok bool)
	Set(ctx context.Context, key string, responseBytes []byte)
	Delete(ctx context.Context, key string)
}

// StorageLifecycle excercises all basic operations of httpcache.Storage implementation.
func StorageLifecycle(t *testing.T, cache Storage) {
	key := "index.html"
	// key := "https://lovergne.dev/index.html"
	_, ok := cache.Get(context.Background(), key)
	if ok {
		t.Fatal("retrieved key before adding it")
	}

	val := []byte("some bytes")
	cache.Set(context.Background(), key, val)

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

func StorageLifecyclePBT(t *testing.T, cache Storage) {
	rapid.Check(t, func(t *rapid.T) {
		key := rapid.String().Draw(t, "key")
		val := rapid.SliceOf(rapid.Byte()).Draw(t, "value")

		_, ok := cache.Get(context.Background(), key)
		if ok {
			t.Fatal("retrieved key before adding it")
		}

		cache.Set(context.Background(), key, val)

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
	})
}
