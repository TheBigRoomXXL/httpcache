package otter

import (
	"context"

	"github.com/maypok86/otter/v2"
)

// otterStorage is an implementation of httpcache/core.Storage that caches responses in memory with Otter
type otterStorage struct {
	otter *otter.Cache[string, []byte]
}

// Get returns the response corresponding to key if present.
func (c otterStorage) Get(_ context.Context, key string) (resp []byte, ok bool) {
	return c.otter.GetIfPresent(key)
}

// Set saves a response to the storage.
func (c otterStorage) Set(_ context.Context, key string, resp []byte) {
	c.otter.Set(key, resp)
}

// Delete removes the response from the storage.
func (c otterStorage) Delete(_ context.Context, key string) {
	c.otter.Invalidate(key)
}

// New returns a httpcache/core.Storage implementation using the provided Otter as underlying storage.
func New(ottterCache *otter.Cache[string, []byte]) *otterStorage {
	return &otterStorage{ottterCache}
}
