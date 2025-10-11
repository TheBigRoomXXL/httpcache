// Package redis provides a redis interface for http caching.
package ristretto

import (
	"context"

	"github.com/dgraph-io/ristretto/v2"
	"pkg.lovergne.dev/httpcache"
)

// cache is an implementation of httpcache.Cache that caches responses in memory with ristretto
type cache struct {
	ristretto *ristretto.Cache[string, []byte]
}

// Get returns the response corresponding to key if present.
func (c cache) Get(_ context.Context, key string) (resp []byte, ok bool) {
	return c.ristretto.Get(key)
}

// Set saves a response to the cache as key.
func (c cache) Set(_ context.Context, key string, resp []byte) {
	c.ristretto.Set(key, resp, int64(len(resp)))
}

// Delete removes the response with key from the cache.
func (c cache) Delete(_ context.Context, key string) {
	c.ristretto.Del(key)
}

// New returns a new Cache with the given ristretto cache
func New(ris *ristretto.Cache[string, []byte]) httpcache.Cache {
	return cache{ris}
}
