package redis_test

import (
	"testing"

	"github.com/gomodule/redigo/redis"
	storageedis "pkg.lovergne.dev/httpcache/redis"
	"pkg.lovergne.dev/httpcache/storagetest"
)

func TestRedisCache(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("no server running at localhost:6379")
	}
	conn.Do("FLUSHALL")

	storagetest.StorageLifecycle(t, storageedis.New(conn, "httpcache"))
}

func TestRedisCachePBT(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatalf("no server running at localhost:6379")
	}
	conn.Do("FLUSHALL")

	storagetest.StorageLifecyclePBT(t, storageedis.New(conn, "httpcache"))
}
