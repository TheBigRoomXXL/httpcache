package memcache_test

import (
	"net"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"

	storagememcache "pkg.lovergne.dev/httpcache/memcache"
	"pkg.lovergne.dev/httpcache/storagetest"
)

const testServer = "localhost:11211"

func TestMemCache(t *testing.T) {
	conn, err := net.Dial("tcp", testServer)
	if err != nil {
		t.Fatalf("no server running at %s", testServer)
	}
	conn.Write([]byte("flush_all\r\n")) // flush memcache
	conn.Close()

	cache := memcache.New(testServer)
	storagetest.StorageLifecycle(t, storagememcache.New(cache))
}

func TestMemCachePBT(t *testing.T) {
	conn, err := net.Dial("tcp", testServer)
	if err != nil {
		t.Fatalf("no server running at %s", testServer)
	}
	conn.Write([]byte("flush_all\r\n")) // flush memcache
	conn.Close()

	cache := memcache.New(testServer)
	storagetest.StorageLifecyclePBT(t, storagememcache.New(cache))
}
