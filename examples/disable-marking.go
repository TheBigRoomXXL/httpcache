package main

import (
	"fmt"
	"net/http"

	httpcache "pkg.lovergne.dev/httpcache/core"
)

func exampleDisableMarking() {
	storage := httpcache.NewInMemoryStorage()
	clientWithCache := &http.Client{
		Transport: &httpcache.Transport{
			Storage:             storage,
			MarkCachedResponses: false,
		},
	}

	fmt.Println("First request:")
	getAndLogResponse(clientWithCache, "https://lovergne.dev")
	fmt.Println("\nSecond request:")
	getAndLogResponse(clientWithCache, "https://lovergne.dev")
}
