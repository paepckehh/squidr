package squidr

import "time"

const (
	_proxyURL                      = "http://127.0.0.80:8080"
	_dnsSrv                        = "127.0.0.53:53"
	_dnsTimeout                    = time.Second * 4
	_blockDomain                   = "access.deny"
	_blockPage                     = ""
	_doNotValidateRandomMapTargets = false
	_uaApp                         = "Mozilla/5.0 (X11; CrOS aarch64 13597.84.0) "
	_uaFramework                   = "AppleWebKit/537.36 (KHTML, like Gecko) "
	_uaOS                          = "Chrome/104.0.5112.105 Safari/537.36"
	_userAgent                     = _uaApp + _uaFramework + _uaOS
)

var (
	blocklist = []string{}

	allowPOST = _gitrepos

	rewrite = map[string][]string{
		"paepcke.de": {"git.paepcke.de"},
	}

	redirect302 = map[string][]string{
		"mobile.twitter.com": {"twitter.com", "www.twitter.com", "nitter.net"},
		"searx.com":          {"searx.net", "searxng.net", "searx.org", "searxng.org", "search"},
	}

	redirect301 = map[string][]string{
		"whooglesearch.ml": {"google.com", "www.google.com", "search.google.com"},
	}

	redirect301randomMap = map[string][]string{
		"twitter": {"mobile.twitter.com"},
		"reddit":  {"reddit.com", "www.reddit.com", "old.reddit.com"},
		"searx":   {"searx.com"},
	}

	randomMap = map[string]rMap{
		"reddit": {
			doNotValidate: false,
			headerOnly:    true,
			minSize:       1024 * 12,
			validateURL:   "",
			targets:       _teddit,
		},
		"twitter": {
			doNotValidate: false,
			headerOnly:    false,
			minSize:       1024 * 12,
			validateURL:   "/twitter",
			targets:       _nitter,
		},
		"searx": {
			doNotValidate: false,
			headerOnly:    true,
			minSize:       1024 * 3,
			validateURL:   "",
			targets:       _searx,
		},
	}
)
