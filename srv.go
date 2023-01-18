package squidr

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
)

const (
	_servBase       = "/var/squid/reports"
	_servPreview    = "/var/squid/reports/www"
	_servTLS        = _servBase + "/tls"
	_servURL        = _servBase + "/url"
	_servDNS        = _servBase + "/dns"
	_servSSS        = _servBase + "/sss"
	_defaultErrFile = "err.html"
	_defaultErrHtml = "<HTML><H1>SQUIDR DEFAULT ERROR</H1></HTML>"
)

type httpFS struct {
	http.FileSystem
}

func servReports() {
	_ = os.MkdirAll(_servBase, 0o770)
	_ = os.MkdirAll(_servURL, 0o770)
	_ = os.MkdirAll(_servDNS, 0o770)
	_ = os.MkdirAll(_servTLS, 0o770)
	_ = os.MkdirAll(_servSSS, 0o770)
	_ = os.WriteFile(_servBase+_slashfwd+_defaultErrFile, []byte(_defaultErrHtml), 0o660)
	reports := http.FileServer(httpFS{http.Dir(_servBase)})
	http.Handle("/", reports)
	panic(http.ListenAndServeTLS("127.0.0.80:9292", "/etc/app/squid/proxylocal.pem", "/etc/app/squid/proxylocal.key", nil))
}

func servPreview() {
	preview := http.FileServer(httpFS{http.Dir(_servPreview)})
	http.Handle("/www", preview)
	panic(http.ListenAndServeTLS("preview.paepcke.pnoc:4443", "/etc/app/squid/preview.pem", "/etc/app/squid/preview.key", nil))
}

func writeReport(data []byte, category string) error {
	hash := sha512.Sum512(data)
	id := base64.RawURLEncoding.EncodeToString(hash[:])
	filename := category + _slashfwd + id
	if err := os.WriteFile(filename, data, 0o660); err != nil {
		debugLogChan <- _errFile + filename + _space + err.Error()
	}
	return errors.New(id)
}
