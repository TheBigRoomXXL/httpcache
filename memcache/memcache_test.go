package memcache

import (
	"net"
	"testing"

	"github.com/bradfitz/gomemcache/memcache"

	"pkg.lovergne.dev/httpcache/test"
)

const testServer = "localhost:11211"

func TestMemCache(t *testing.T) {
	conn, err := net.Dial("tcp", testServer)
	if err != nil {
		t.Skipf("skipping test; no server running at %s", testServer)
	}
	conn.Write([]byte("flush_all\r\n")) // flush memcache
	conn.Close()

	cache := memcache.New(testServer)
	test.Cache(t, New(cache))
}
