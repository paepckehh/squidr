package squidr

import (
	"crypto/tls"
	"crypto/x509"
	"strings"
	"sync"
)

const (
	// token
	_tcp         = "tcp"  // ip[6|4] happy eyeball
	_tcp4        = "tcp4" // ip4 only
	_tcp6        = "tcp6" // ip6 only
	_sslport     = ":443"
	_html        = ".html"
	_none        = "[none]"
	_checkMode   = "[squidr tls domain validation mode]"
	_sslCaTrust  = "/etc/ssl/external_trust.pem"
	_doublepoint = ":"
)

var cacheTls sync.Map

func isTlsChainValid(host, path string) error {
	// set ip4 only
	tcp := _tcp4

	// setup & switch to check mode if requested
	check, connected := false, false
	if host == _tls {
		s := strings.Split(path, _slashfwd)
		host, path, check = s[0], _checkMode, true
	}

	// check cache
	if _, ok := cacheTls.Load(host); ok && !check {
		return nil
	}

	// connection without any validation
	tlsconf := getTlsConfLegacy(x509.NewCertPool())
	tlsconf.InsecureSkipVerify = true
	connInsecure, err := tls.Dial(tcp, host+_sslport, tlsconf)
	if err != nil {
		debugLogChan <- _ioConnErr + host
		return reportTLS(host, path, connInsecure, connected, err)
	}
	connected = true
	connInsecure.Close()

	// connection without any validation
	tlsconf = getTlsConfLegacy(externalTrust)
	conn, err := tls.Dial(tcp, host+_sslport, tlsconf)
	if err != nil {
		debugLogChan <- _ioConnErr + host
		return reportTLS(host, path, connInsecure, connected, err) // report data from first touch
	}
	_ = conn.Handshake()
	defer conn.Close()

	// verify hostname
	err = conn.VerifyHostname(host)
	if err != nil {
		return reportTLS(host, path, conn, connected, err)
	}

	// cache valid host
	cacheTls.Store(host, true)

	// switch to check mode, full report, even without error
	if check {
		return reportTLS(host, path, conn, connected, err)
	}
	return nil
}
