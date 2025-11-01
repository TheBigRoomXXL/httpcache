package test_test

import (
	"testing"

	httpcache "github.com/TheBigRoomXXL/httpcache/core"
	"github.com/TheBigRoomXXL/httpcache/core/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
