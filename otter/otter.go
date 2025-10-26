package otter

import (
	"context"

	"github.com/maypok86/otter/v2"
	"pkg.lovergne.dev/httpcache"
)

// cache is an implementation of httpcache.Cache that caches responses in memory with Otter
type cache struct {
	otter *otter.Cache[string, []byte]
}

// Get returns the response corresponding to key if present.
func (c cache) Get(_ context.Context, key string) (resp []byte, ok bool) {
	return c.otter.GetIfPresent(key)
}

// Set saves a response to the cache as key.
func (c cache) Set(_ context.Context, key string, resp []byte) {
	c.otter.Set(key, resp)
}

// Delete removes the response with key from the cache.
func (c cache) Delete(_ context.Context, key string) {
	c.otter.Invalidate(key)
}

// New returns a new Cache with the given Otter cache
func New(ottterCache *otter.Cache[string, []byte]) httpcache.Cache {
	return cache{ottterCache}
}
