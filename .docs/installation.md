# Installation

To install the core package simply do:
```bash
go get pkg.lovergne.dev/httpcache/core
```

This will provide you with an implementations of `http.RoundTripper` and `http.Client` that support HTTP caching sementics.

The core package only come with the basic `InMemoryStorage`. If you need more storage options that depend on third-party library you need to install them separetly:

!!! info
    Having to install every storage separetly is a bit less pratical than including them in the core package but it allow you to minize the number of dependencies your project has. 


## Diskv Storage
```bash
go get pkg.lovergne.dev/httpcache/diskv
```

## Otter Storage
```bash
go get pkg.lovergne.dev/httpcache/otter
```

## Memcache Storage
```bash
go get pkg.lovergne.dev/httpcache/memcache
```

## Redis Storage
```bash
go get pkg.lovergne.dev/httpcache/diskv
```

You are now ready for you first example, continue to the [usage section](/usage).
