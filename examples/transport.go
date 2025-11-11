package main

import (
	"fmt"
	"net/http"

	httpcache "pkg.lovergne.dev/httpcache/core"
)

func exampleTransport() {
	// Instanciate a transport
	storage := httpcache.NewInMemoryStorage()
	transport := httpcache.NewCachedTransport(storage)

	// Use it like a normal http.Transport object
	request, err := http.NewRequest("GET", "https://lovergne.dev/", nil)
	if err != nil {
		panic(err)
	}
	resp, err := transport.RoundTrip(request)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
}
