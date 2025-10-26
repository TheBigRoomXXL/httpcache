// Package diskv provides an implementation of httpcache.Cache that uses the diskv package
// Thank to diskv the cache can benefit from in-memory buffer and compression.
//
// This cache ignore context as diskv does not support it.

package diskv

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"

	"github.com/peterbourgon/diskv/v3"
)

// Cache is an implementation of httpcache.Cache that supplements the in-memory map with persistent storage
type Cache struct {
	d *diskv.Diskv
}

// Get returns the response corresponding to key if present
func (c *Cache) Get(_ context.Context, key string) (resp []byte, ok bool) {
	key = keyToFilename(key)
	resp, err := c.d.Read(key)
	if err != nil {
		return []byte{}, false
	}
	return resp, true
}

// Set saves a response to the cache as key
func (c *Cache) Set(_ context.Context, key string, resp []byte) {
	key = keyToFilename(key)
	c.d.WriteStream(key, bytes.NewReader(resp), true)
}

// Delete removes the response with key from the cache
func (c *Cache) Delete(_ context.Context, key string) {
	key = keyToFilename(key)
	c.d.Erase(key)
}

func keyToFilename(key string) string {
	h := md5.New()
	io.WriteString(h, key)
	return hex.EncodeToString(h.Sum(nil))
}

// New returns a new Cache using the provided Diskv as underlying storage.
func New(d *diskv.Diskv) *Cache {
	return &Cache{d}
}
