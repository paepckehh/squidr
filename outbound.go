package squidr

import (
	"compress/gzip"
	"compress/zlib"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func getTlsConf(trust *x509.CertPool) *tls.Config {
	return &tls.Config{
		RootCAs:                trust,
		InsecureSkipVerify:     false,
		SessionTicketsDisabled: true,
		Renegotiation:          0,
		NextProtos:             []string{"http/1.1"},
		MinVersion:             tls.VersionTLS13,
		MaxVersion:             tls.VersionTLS13,
		CipherSuites:           []uint16{tls.TLS_CHACHA20_POLY1305_SHA256},
		CurvePreferences:       []tls.CurveID{tls.X25519},
	}
}

func getTlsConfLegacy(trust *x509.CertPool) *tls.Config {
	return &tls.Config{
		RootCAs:                trust,
		InsecureSkipVerify:     false,
		SessionTicketsDisabled: true,
		Renegotiation:          0,
		NextProtos:             []string{"http/1.1"},
		MinVersion:             tls.VersionTLS12,
		MaxVersion:             tls.VersionTLS13,
		CipherSuites: []uint16{
			tls.TLS_CHACHA20_POLY1305_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		CurvePreferences: []tls.CurveID{
			tls.X25519,
			tls.CurveP521,
			tls.CurveP384,
			tls.CurveP256,
		},
	}
}

func getTransportProxy(proxyname string, tlsconf *tls.Config) *http.Transport {
	proxy, err := url.Parse(proxyname)
	if err != nil {
		panic("[squidr] invalid proxy name URL")
	}
	transport := getTransport(tlsconf)
	transport.Proxy = http.ProxyURL(proxy)
	return transport
}

func getTransport(tlsconf *tls.Config) *http.Transport {
	return &http.Transport{
		TLSClientConfig:    tlsconf,
		DisableCompression: true,
		ForceAttemptHTTP2:  false,
	}
}

func getRequest(plainURL, ua string) (*http.Request, error) {
	targetURL, err := url.Parse(plainURL)
	if err != nil {
		return &http.Request{}, errors.New("invalid url")
	}
	return &http.Request{
		URL:        targetURL,
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header: http.Header{
			"User-Agent":      []string{ua},
			"Accept-Encoding": []string{"gzip", "deflate"},
		},
	}, nil
}

func getClient(transport *http.Transport) *http.Client {
	return &http.Client{
		CheckRedirect: nil,
		Jar:           nil,
		Transport:     transport,
	}
}

func decodeResponse(resp *http.Response) ([]byte, error) {
	var err error
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
	case "deflate":
		reader, err = zlib.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	if err != nil {
		return nil, errors.New("decode: " + err.Error())
	}
	return io.ReadAll(reader)
}
