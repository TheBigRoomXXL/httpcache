# Welcome 

TODO: add banner
![banner](img/banner.svg)

**httpcache** implements HTTP caching semantics ([RFC 9111](https://www.rfc-editor.org/rfc/rfc9111.html)) for Go’s `http.RoundTripper` and `http.Client`.

The core package `pkg.lovergne.dev/httpcache/core` transparently adds caching behavior to standard Go HTTP clients, enabling efficient reuse of responses according to HTTP cache rules.

**httpcache** supports multiple backends for storing cached responses, including:

* [Otter](https://maypok86.github.io/otter/)
* [Diskv](https://github.com/peterbourgon/diskv)
* [Redis](https://redis.io/docs/latest/)
* [Memcached](https://memcached.org/)

and can be extended easily with custom storage implementations.

This library aims to be largely compliant with [RFC 9111](https://www.rfc-editor.org/rfc/rfc9111.html), while remaining practical and lightweight. Currently, is is only suitable for use as a "private" cache (i.e. for a web-browser or an API-client and not for a shared proxy).


## Acknowledgement

This project is a revival of the awesome library [httpcache by gregjones](https://github.com/gregjones/httpcache) which implemented support for most RFC 7234 directive. This library is really well written and quite simple in it's architecture. Most of the core of this package is still taken from that library and we simply update it where needed.


## Feedback, help or bug report

If you need anything related to this project, whether it's just giving feedback, requesting a feature, or simply asking for help to understand something, open an [issue]((https://github.com/TheBigRoomXXL/httpcache/issues)) on the official [repository](https://github.com/TheBigRoomXXL/httpcache/).

*If you want to contribute code, please open an issue first so that we can collaborate efficiently.*

## Getting started

To get started go the [Installation](/installation) page and then proceed to the [Usage](/usage) page.
