# Usage

## TODO
- range caching
- Explaination on the importance of no store for infinite stream


## Basic example

Let's start with a minimal example to demonstrate how **httpcache** works and how to use it. 

Here are the steps necessary to use the cache:

1.  instanciate a storage
2. instanciate `http.Client` with **httpcache** transport.
3. Make a request
4. Read it's body
5. Make another request with the same url. 

```go
func exampleBasic() {
	storage := httpcache.NewInMemoryStorage()
	clientWithCache := httpcache.NewCachedClient(storage)

	fmt.Println("First request:")
	getAndLogResponse(clientWithCache, "https://lovergne.dev")
	fmt.Println("\nSecond request:")
	getAndLogResponse(clientWithCache, "https://lovergne.dev")
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

	// Request coming from the cache will have the X-From-Cache header set to 1
	fmt.Println("  Cache-Control ", resp.Header.Get("Cache-Control"))
	fmt.Println("  X-From-Cache  ", resp.Header.Get("X-From-Cache"))
	fmt.Println("  Ellapsed time ", ellapsed)
}
```

I you run the following script, you should see an outut similar to that:
```txt
First request:
  Cache-Control  max-age=600
  X-From-Cache   
  Ellapsed time  310.794791ms

Second request:
  Cache-Control  max-age=600
  X-From-Cache   1
  Ellapsed time  173.154µs
```

As you can clearly see, the first request comes directly from the origin, that's why it take hundreds of milisecond does not have any `X-From-Cache` header. Because it has the  `max-age=600` directive in the `Cache-Control` header, when we make the second request the response is served directly from the cache, not the origin, so it takes only a few microseconds and has the `X-From-Cache` header set to one.

!!! info "Why do I have to read the full body to cache the response?"
    The caching mecanism is mostly lazy, the response won't be written to the storage until you have read the body and reach end of file. This is a bit counter-intuitive but it avoid issues with infinit stream blocking the transport in an infinit loop. See [PR 71 of gregjones/httpcache ](https://github.com/gregjones/httpcache/pull/71)

## Check if response is cached or has been re-evaluated

By default requests coming from transport will have the following headers: 
- `X-From-Cache header: 1 ` if they come from the cache storage.
- `X-Revalidated`:  if they come from the cache  and have been re-evaluated against the origin.

## Using other storage

`InMemoryStorage` is very fast and good enough for basic use case but it has one major flaw: it does not have an size limit. Because of that it will grow until you program reach an out-of-memory error and crash. For that reason **you should not use it for any kind of long-runner process.** Instead you will want to use of the other storage provided by **httpcache**: otter, diskv, memcache and redis. 

If you need help deciding which one is the best fit you use case, there is a [didicated page in the documentation](/how-to-choose-a-storage).

For this example we will use the awesome [otter](https://maypok86.github.io/otter/) as it also an in memery storage but with a size limit. Storage implementation are in separated package, so must install them separatly:
```bash
go get pkg.lovergne.dev/httpcache/otter
```
```go
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
```

This should produce the following output
```txt
First request:
  Cache-Control  max-age=600
  X-From-Cache   
  Ellapsed time  610.071785ms

Second request:
  Cache-Control  max-age=600
  X-From-Cache   1
  Ellapsed time  192.38µs
```

## Disable the marking of cached response

If you don't want the response comming from the cache to be marked as such with a `X-From-Cache` or `X-Revalidated` header you need to instanciate the transport struct directly with the `MarkCachedResponses` flag set to false

```go
func ExampleDisableMarking() {
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
```

This is the expected output:
```txt
First request:
  Cache-Control  max-age=600
  X-From-Cache   
  Ellapsed time  365.340115ms

Second request:
  Cache-Control  max-age=600
  X-From-Cache   
  Ellapsed time  137.387µs
```

As you can see the response is still served by the cache because it only takes 137µs but it is not maked with a header.

## Using the transport directly

While we provide a `NewCachedClient` which return `http.Client` for ease of use, the actual caching logic is implemented in a custom `Transport` struct that implemement the `http.RoundTripper` interface.

If you need you can use the Transport directly:
```go
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
```

This should produce the following output:
```txt
Response status: 200 OK
```

## Wrapping another transport

By default the the library will wraps `http.DefaultTransport` but it can also wrap any struct that support the `http.RoundTripper` interface:

```go

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
```

This should produce the following output:
```txt
Request took 352.391971ms
Response status: 200 OK
```
