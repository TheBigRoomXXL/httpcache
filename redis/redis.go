// Package redis provides a redis interface for http caching.
package redis

import (
	"context"

	httpcache "github.com/TheBigRoomXXL/httpcache/core"
	"github.com/gomodule/redigo/redis"
)

// cache is an implementation of httpcache.Cache that caches responses in a
// redis server.
type cache struct {
	connection redis.Conn
	namespace  string // prefix every key to avoid collision
}

// cacheKey modifies an httpcache key for use in redis. Specifically, it
// prefixes keys to avoid collision with other data stored in redis.
func cacheKey(namespace string, key string) string {
	return namespace + key
}

// Get returns the response corresponding to key if present.
func (c cache) Get(ctx context.Context, key string) (resp []byte, ok bool) {
	item, err := redis.Bytes(redis.DoContext(c.connection, ctx, "GET", cacheKey(c.namespace, key)))
	if err != nil {
		return nil, false
	}
	return item, true
}

// Set saves a response to the cache as key.
func (c cache) Set(ctx context.Context, key string, resp []byte) {
	redis.DoContext(c.connection, ctx, "SET", cacheKey(c.namespace, key), resp)
}

// Delete removes the response with key from the cache.
func (c cache) Delete(ctx context.Context, key string) {
	redis.DoContext(c.connection, ctx, "DEL", cacheKey(c.namespace, key))
}

// New returns a new Cache with the given redis connection.
func New(client redis.Conn, namespace string) httpcache.Cache {
	return cache{client, namespace}
}
