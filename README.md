# OVERVIEW
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/squidr.svg)](https://pkg.go.dev/paepcke.de/squidr) [![Go Report Card](https://goreportcard.com/badge/paepcke.de/squidr)](https://goreportcard.com/report/paepcke.de/squidr) [![Go Build](https://github.com/paepckehh/squidr/actions/workflows/golang.yml/badge.svg)](https://github.com/paepckehh/squidr/actions/workflows/golang.yml)

[paepche.de/squidr](https://paepcke.de/squidr/)

Squid Cache Proxy [squid-cache.org](https://squid-cache.org/) Infosec Companion! 

Protect your Corporate / Privacy Squid (MiTM) Proxy!

Evaluates and analyzes FIRST-IN-LINE every connect attempt via a customized
TLS stack for upstream problems (TLS trace, fingerprints, presented trustchains)
before the squid/openssl/os stack is allowed to fetch.

Automatically documents full tls handshake traces and presented certificates,
dns states, presented target urls and many more into immutable incident snapshots. 

## Adds this new additional http api endpoints to squid

* /tls/{domain}
  take & archive snapshot of tls stack handshake, and store detailed tls traces

* /dns/{domain}
  take & archive snapshot dns env view from several differend upstream resolver

* /url/{target}
  fetch site, scrape urls & archive snapshot (incl. script/hidden) target uri/urls

* /sss/{keywords}
  search keywords via differned privacy first search-engines


## Example incident reports

* [example-tls-report](https://paepckehh.github.io/squidr-examples/tls.html)
* [example-dns-report](https://paepckehh.github.io/squidr-examples/dns.html)
* [example-url-report](https://paepckehh.github.io/squidr-examples/url.html)


## optional

 * Respond to social media requests (twitter, reddit, ...) via
   selected, verified or randomized alternative instances of (nitter, teddit, ...)

 * Respond to (unsecure) search engines request (google, bing, ...) via
   selected, verified or randomized alternative instances (searX, searXNG, ...)

# Anything else?

Yes, its build-time-configuration, no prebuild binaries!

# DOCS

[pkg.go.dev/paepcke.de/squidr](https://pkg.go.dev/paepcke.de/squidr)

# CONTRIBUTION

Yes, Please! PRs Welcome! 

