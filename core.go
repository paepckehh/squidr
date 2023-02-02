package squidr

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"paepcke.de/dnscache"
)

const (
	_ok           = "ERR"
	_err          = "OK rewrite-url=\"https://127.0.0.80:9292/err.html"
	_dnserr       = "OK rewrite-url=\"https://127.0.0.80:9292/dns.html"
	_errRoot      = "OK rewrite-url=\"https://127.0.0.80:9292/"
	_errTLS       = "OK rewrite-url=\"https://127.0.0.80:9292/tls/"
	_errDNS       = "OK rewrite-url=\"https://127.0.0.80:9292/dns/"
	_errURL       = "OK rewrite-url=\"https://127.0.0.80:9292/url/"
	_errSSS       = "OK rewrite-url=\"https://127.0.0.80:9292/sss/"
	_rewrite      = "OK rewrite-url=\""
	_redirect301  = "OK status=+301 url=\""
	_redirect302  = "OK status=+302 url=\""
	_end          = "\""
	_sep          = " "
	_slashfwd     = "/"
	_linefeed     = "\n"
	_empty        = ""
	_doubleqoute  = "\""
	_http         = "http://"
	_https        = "https://"
	_url          = "url"
	_tls          = "tls"
	_dns          = "dns"
	_sss          = "sss"
	_errFile      = "unable to createfile "
	_errTrustExit = "unable to init trust store, exit "

	// HTTP METHODS
	_CONNECT = "CONNECT"
	_GET     = "GET"
	_HEAD    = "HEAD"
	_POST    = "POST"
	_PUT     = "PUT"
	_PUSH    = "PUSH"
	_PATCH   = "PATCH"
	_DELETE  = "DELETE"
	_OPTIONS = "OPTIONS"
	_TRACE   = "TRACE"
)

type rMap struct {
	doNotValidate bool     // do not validate if random target is responsive
	headerOnly    bool     // only validate header fields
	minSize       int64    // minimum body response size requirded
	validateURL   string   // target specific test url suffix
	targets       []string // targets
}

type search struct {
	target    string
	blocked   bool
	rewrite   bool
	temporary bool
	change    bool
}

var (
	urlCache     sync.Map
	debugLogChan = make(chan string, 1024)
	outChan      = make(chan string, 1024)
)

func itoa(in int) string { return strconv.Itoa(in) }

func action(line string) {
	status, valid := _redirect302, false
	debugLogChan <- _ioDebugIN + line
	s := strings.Fields(strings.TrimSpace(line))
	if len(s) > 3 && s[4] == _CONNECT {
		if len(s[1]) == 3 {
			switch s[1] {
			case _url, _tls, _dns, _sss:
				outChan <- s[0] + _sep + _err
				return
			}
		}
		outChan <- s[0] + _sep + _ok
		return
	}
	if len(s) < 1 || len(s[1]) < 8 {
		debugLogChan <- "FATAL INPUT: " + s[0] + _sep + line
		return
	}
	domain, path, _ := strings.Cut(s[1][8:], "/")
	_, cached := urlCache.Load(domain + path)
	switch {
	case cached:
		outChan <- s[0] + _sep + _ok
		return
	case len(s) < 3:
		debugLogChan <- _broken + line
		outChan <- s[0] + _sep + _err
		return
	case s[4] == http.MethodConnect:
		if _, err := dnscache.LookupHost(domain); err != nil { // fail early, fail cheap
			debugLogChan <- _dnsterminated + _sep + domain + _sep + line
			outChan <- s[0] + _sep + _dnserr
			return
		}
		outChan <- s[0] + _sep + _ok
		return
	case len(s[1]) < 10:
		debugLogChan <- _broken + line
		outChan <- s[0] + _sep + _err
		return
	case s[1][:7] == _http:
		debugLogChan <- _redirectHTTPS + line
		outChan <- s[0] + _sep + _redirect301 + _https + s[1][7:] + _end
		return
	case domain == _tls:
		err := isTlsChainValid(_tls, path)
		id := err.Error()
		if err != nil {
			debugLogChan <- _tls + _point + id + _point + line
			outChan <- s[0] + _sep + _errTLS + id
			return
		}
	case domain == _url:
		err := reportURL(path)
		id := err.Error()
		if err != nil {
			debugLogChan <- _url + _point + id + _point + line
			outChan <- s[0] + _sep + _errURL + id
			return
		}
	case domain == _dns:
		err := reportDNS(path)
		id := err.Error()
		if err != nil {
			debugLogChan <- _dns + _point + id + _point + line
			outChan <- s[0] + _sep + _errDNS + id
			return
		}
	case domain == _sss:
		err := reportSEARX(path)
		id := err.Error()
		if err != nil {
			debugLogChan <- _sss + _point + id + _point + line
			outChan <- s[0] + _sep + _errSSS + id
			return
		}
	case s[1][:8] == _https:
		if _, err := dnscache.LookupHost(domain); err != nil { // fail early, fail cheap
			debugLogChan <- _dnsterminated + _sep + domain + _sep + line
			outChan <- s[0] + _sep + _dnserr
			return
		}
		valid = true
		switch s[4] {
		case http.MethodGet, http.MethodHead:
		case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions:
			if !_gitRepos[domain] && !_allowPost[domain] && !_mapsearx[domain] {
				debugLogChan <- _nopost + line
				outChan <- s[0] + _sep + _err
				return
			}
		}
	}
	if !valid {
		debugLogChan <- _terminated + line
		outChan <- s[0] + _sep + _err
		return
	}
	search := searchDomain(domain)
	if search.change {
		domain = search.target
		switch {
		case search.blocked:
			status = _rewrite
			path = _blockPage
		case search.rewrite:
			status = _rewrite
		case search.temporary:
			status = _redirect301
		}
		outChan <- s[0] + _sep + status + _https + domain + _slashfwd + path + _end
		return
	}
	if err := isTlsChainValid(domain, path); err != nil {
		id := err.Error()
		debugLogChan <- _tlsterminated + id + _point + line
		outChan <- s[0] + _sep + _errTLS + id
		return
	}
	urlCache.Store(domain+path, true) // cacheable domain name
	outChan <- s[0] + _sep + _ok
}
