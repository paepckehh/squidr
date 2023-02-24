# OVERVIEW
[![Go Reference](https://pkg.go.dev/badge/paepcke.de/squidr.svg)](https://pkg.go.dev/paepcke.de/squidr) [![Go Report Card](https://goreportcard.com/badge/paepcke.de/squidr)](https://goreportcard.com/report/paepcke.de/squidr) [![Go Build](https://github.com/paepckehh/squidr/actions/workflows/golang.yml/badge.svg)](https://github.com/paepckehh/squidr/actions/workflows/golang.yml)

[paepcke.de/squidr](https://paepcke.de/squidr/)

Squid Cache Proxy [squid-cache.org](https://squid-cache.org/) Infosec Companion! 

Protect your corporate privacy with a Squid MiTM (Man-in-the-Middle) proxy!

Evaluate and analyze every connection attempt with a customized TLS stack for
upstream problems (TLS trace, fingerprints, presented trust chains) before the 
Squid/OpenSSL/OS stack is allowed to fetch. Automatically document full TLS 
handshake traces and presented certificates, DNS states, presented target URLs,
and more into immutable incident snapshots. 

## Adds this new additional http api endpoints to squid

* /tls/{domain}
  take & archive snapshot of tls stack handshake, and store detailed tls traces

* /dns/{domain}
  take & archive snapshot dns env view from several differend upstream resolver

* /url/{target}
  fetch site, scrape urls & archive snapshot (incl. script/hidden) target uri/urls

* /sss/{keywords}
  search keywords via differned privacy first search-engines


## Example incident reports.

* [example-tls-report](https://paepckehh.github.io/squidr-examples/tls.html)
* [example-dns-report](https://paepckehh.github.io/squidr-examples/dns.html)
* [example-url-report](https://paepckehh.github.io/squidr-examples/url.html)


## Optional!

 * Respond to social media requests (twitter, reddit, ...) via
   selected, verified or randomized alternative instances of (nitter, teddit, ...)

 * Respond to (unsecure) search engines request (google, bing, ...) via
   selected, verified or randomized alternative instances (searX, searXNG, ...)

## Anything else?

Yes, its build-time-configuration, no prebuild binaries!

# DOCS

[pkg.go.dev/paepcke.de/squidr](https://pkg.go.dev/paepcke.de/squidr)

# ARTWORK

Generated by OpenAI. 

* Some additional ai created samples from the App Logo creation process.
* No [Squids](https://www.schneier.com/tag/squid/) or Cats are harmed in the creation process!

![ai_generated_squid_gets_eaten_by_a_cute_cat](https://github.com/paepckehh/paepckehh/raw/main/artwork/squidr.png)
![ai_generated_squid_gets_eaten_by_a_cute_cat](https://github.com/paepckehh/paepckehh/raw/main/artwork/squidr2.png)
![ai_generated_squid_gets_eaten_by_a_cute_cat](https://github.com/paepckehh/paepckehh/raw/main/artwork/squidr4.png)

# CONTRIBUTION

Yes, Please! PRs Welcome! 
