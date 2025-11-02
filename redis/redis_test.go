package redis

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	"pkg.lovergne.dev/httpcache/core/test"
)

func TestRedisCache(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("no server running at localhost:6379")
	}
	conn.Do("FLUSHALL")
	test.Cache(t, New(conn, "httpcache"))
}
