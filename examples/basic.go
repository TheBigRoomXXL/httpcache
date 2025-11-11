package main

import (
	"fmt"
	"io"
	"time"

	httpcache "pkg.lovergne.dev/httpcache/core"
)

func basic() {
	storage := httpcache.NewInMemoryStorage()
	clientWithCache := httpcache.NewCachedClient(storage)
	url := "https://lovergne.dev"

	start := time.Now()
	resp, err := clientWithCache.Get(url)
	if err != nil {
		panic(err)
	}
	ellapsed := time.Since(start)

	// ⚠ if you don't reach EOF on the body, the request won't be cached.
	io.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println("First request:")
	fmt.Println("  Cache-Control ", resp.Header.Get("Cache-Control"))
	fmt.Println("  X-From-Cache  ", resp.Header.Get("X-From-Cache"))
	fmt.Println("  Ellapsed time ", ellapsed)

	start = time.Now()
	resp, err = clientWithCache.Get(url)
	if err != nil {
		panic(err)
	}
	ellapsed = time.Since(start)

	fmt.Println("\nSecond request:")
	fmt.Println("  Cache-Control ", resp.Header.Get("Cache-Control"))
	fmt.Println("  X-From-Cache  ", resp.Header.Get("X-From-Cache"))
	fmt.Println("  Ellapsed time ", ellapsed)
}
