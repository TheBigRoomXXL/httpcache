package main

import (
	"fmt"

	"github.com/maypok86/otter/v2"
	httpcache "pkg.lovergne.dev/httpcache/core"
	httpcacheotter "pkg.lovergne.dev/httpcache/otter"
)

func exampleOtter() {
	o := otter.Must(&otter.Options[string, []byte]{
		MaximumSize:     10_000,
		InitialCapacity: 1_000,
	})
	storage := httpcacheotter.New(o)
	clientWithCache := httpcache.NewCachedClient(storage)

	fmt.Println("First request:")
	getAndLogResponse(clientWithCache, "https://lovergne.dev")
	fmt.Println("\nSecond request:")
	getAndLogResponse(clientWithCache, "https://lovergne.dev")
}
