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

// Storage is an implementation of httpcache/core.Storage that use Diskv for persistant storage
type diskvStorage struct {
	diskv *diskv.Diskv
}

// Get returns the response corresponding to key if present
func (s *diskvStorage) Get(_ context.Context, key string) (resp []byte, ok bool) {
	key = keyToFilename(key)
	resp, err := s.diskv.Read(key)
	if err != nil {
		return []byte{}, false
	}
	return resp, true
}

// Set saves a response to the storage
func (s *diskvStorage) Set(_ context.Context, key string, resp []byte) {
	key = keyToFilename(key)
	s.diskv.WriteStream(key, bytes.NewReader(resp), true)
}

// Delete removes the response from the storage
func (s *diskvStorage) Delete(_ context.Context, key string) {
	key = keyToFilename(key)
	s.diskv.Erase(key)
}

// URLs don't make good filename so we must hash them.
func keyToFilename(key string) string {
	h := md5.New()
	io.WriteString(h, key)
	return hex.EncodeToString(h.Sum(nil))
}

// New returns a new implementation of httpcache/core.Storage using the provided Diskv as underlying storage.
func New(d *diskv.Diskv) *diskvStorage {
	return &diskvStorage{d}
}
