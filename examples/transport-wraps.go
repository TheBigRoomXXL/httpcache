package main

import (
	"fmt"
	"net/http"
	"time"

	httpcache "pkg.lovergne.dev/httpcache/core"
)

type customRoundTripper struct{}

func (r *customRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := http.DefaultTransport.RoundTrip(request)
	fmt.Println("Request took", time.Since(start))
	return resp, err
}

func exampleTransportWrap() {
	storage := httpcache.NewInMemoryStorage()
	client := &http.Client{
		Transport: &httpcache.Transport{
			Transport:           &customRoundTripper{},
			Storage:             storage,
			MarkCachedResponses: false,
		},
	}

	resp, err := client.Get("https://lovergne.dev/")
	if err != nil {
		panic(err)
	}
	fmt.Println("Response status:", resp.Status)
}
