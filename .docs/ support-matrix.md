
# HTTP Cache Support Matrix

In the page you can find in details which part of RFC9111 and related specifications are currenlty supported by this livrary. 

At the moment the biggest "gap" in support for public (aka shared) cache but it's support is on the the roadmap. If you need public cache [Souin](https://github.com/darkweak/souin) might a better fit for you.

Specification links point to precise section in the specification, not just the specification itself. 

## Cache-Control — Request Directives

| Directive        | Supported | Specification |
| ---------------- | :-------: | ------------- |
| `max-age`        |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-max-age)|
| `max-stale`      |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-max-stale)|
| `min-fresh`      |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-min-fresh)|
| `no-cache`       |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-no-cache) |
| `no-store`       |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-no-store) |
| `no-transform`   |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-no-transform) |
| `only-if-cached` |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-only-if-cached) |

## Cache-Control — Response Directives

| Directive                | Supported | Specification |
| ------------------------ | :-------: | ------------- |
| `max-age`                |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-max-age-2) |
| `s-maxage`               |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-s-maxage) |
| `no-cache`               |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-no-cache-2) |
| `no-store`               |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-no-store-2) |
| `no-transform`           |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-no-transform-2) |
| `must-revalidate`        |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-must-revalidate) |
| `proxy-revalidate`       |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-proxy-revalidate) |
| `must-understand`        |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-must-understand) |
| `private`                |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-private) |
| `public`                 |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-public) |
| `immutable`              |    🔴     | [rfc8246](https://www.rfc-editor.org/rfc/rfc8246.html) |
| `stale-while-revalidate` |    🔴     | [rfc5861](https://www.rfc-editor.org/rfc/rfc5861.html#section-3)|
| `stale-if-error`         |    🟢     | [rfc5861](https://www.rfc-editor.org/rfc/rfc5861.html#section-4) |

## ETag Support

| Directive        | Supported | Specification |
| ---------------- | :-------: | ------------- |
| ETag strong      |    🟢     | [rfc9110](https://www.rfc-editor.org/rfc/rfc9110#name-etag), [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-validation) |
| ETag weak (`W/`) |    🔴     | [rfc9110](https://www.rfc-editor.org/rfc/rfc9110#name-etag), [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-validation) |
| If-None-Match    |    🟢     | [rfc9110](https://www.rfc-editor.org/rfc/rfc9110#name-if-none-match), [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-validation) |

## Other Headers

| Directive         | Supported | Specification |
| ----------------- | :-------: | ------------- |
| If-Modified-Since |    🟢     | [rfc9110](https://www.rfc-editor.org/rfc/rfc9110#name-if-modified-since), [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-validation) |
| Expires           |    🟢     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-expires) |
| Range             |    🟢     | [rfc9110](https://www.rfc-editor.org/rfc/rfc9110#name-range-requests), [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-storing-incomplete-response)  |
| Pragma            |    🔴     | [rfc9111](https://www.rfc-editor.org/rfc/rfc9111.html#name-pragma) |
| Clear-Site-Data   |    🔴     | [w3](https://www.w3.org/TR/clear-site-data/) |
| Cache-Status      |    🔴     | [rfc9211](https://www.rfc-editor.org/rfc/rfc9211.html) |

