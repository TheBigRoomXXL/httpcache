package diskcache

import (
	"os"
	"testing"

	"github.com/peterbourgon/diskv/v3"
	"pkg.lovergne.dev/httpcache/test"
)

func TestDiskCache(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "httpcache")
	if err != nil {
		t.Fatalf("TempDir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	kv := diskv.New(diskv.Options{
		BasePath:     tempDir,
		CacheSizeMax: 100 * 1024 * 1024, // 100MB
	})
	test.Cache(t, NewWithDiskv(kv))
}
