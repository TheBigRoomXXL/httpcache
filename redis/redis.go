// Package redis provides a redis interface for http caching.
package redis

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

// cache is an implementation of httpcache/core.Storage that store responses in a redis server.
type redisStorage struct {
	connection redis.Conn
	namespace  string // prefix every key to avoid collision
}

// cacheKey modifies an httpcache key for use in redis. Specifically, it
// prefixes keys to avoid collision with other data stored in redis.
func cacheKey(namespace string, key string) string {
	return namespace + key
}

// Get returns the response corresponding to key if present.
func (s redisStorage) Get(ctx context.Context, key string) (resp []byte, ok bool) {
	item, err := redis.Bytes(redis.DoContext(s.connection, ctx, "GET", cacheKey(s.namespace, key)))
	if err != nil {
		return nil, false
	}
	return item, true
}

// Set saves a response to the storage.
func (s redisStorage) Set(ctx context.Context, key string, resp []byte) {
	redis.DoContext(s.connection, ctx, "SET", cacheKey(s.namespace, key), resp)
}

// Delete removes the response from the storage.
func (s redisStorage) Delete(ctx context.Context, key string) {
	redis.DoContext(s.connection, ctx, "DEL", cacheKey(s.namespace, key))
}

// New returns a httpcache/core.Storage implementation using the provided redis as underlying storage.
func New(client redis.Conn, namespace string) *redisStorage {
	return &redisStorage{client, namespace}
}
