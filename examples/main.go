package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var examples = map[string]func(){
	"basic":           exampleBasic,
	"disable-marking": exampleDisableMarking,
	"transport":       exampleTransport,
	"transport-wrap":  exampleTransportWrap,
	"otter":           exampleOtter,
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Bad Usage: you need to pass the example name.")
	}
	exampleFunc, ok := examples[os.Args[1]]
	if !ok {
		fmt.Println("Bad Usage: example not found.")
	}

	exampleFunc()
}

func getAndLogResponse(client *http.Client, url string) {
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	ellapsed := time.Since(start)

	// ⚠ if you don't reach EOF on the body, the request won't be cached.
	io.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println("  Cache-Control ", resp.Header.Get("Cache-Control"))
	fmt.Println("  X-From-Cache  ", resp.Header.Get("X-From-Cache"))
	fmt.Println("  Ellapsed time ", ellapsed)
}
