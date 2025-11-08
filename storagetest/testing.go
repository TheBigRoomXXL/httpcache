package storagetest

import (
	"bytes"
	"context"
	"testing"
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
	key := "testKey"
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
