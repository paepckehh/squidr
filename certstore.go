package squidr

import (
	"crypto/x509"
	_ "embed"
)

var (
	proxyTrust    = x509.NewCertPool()
	externalTrust = x509.NewCertPool()
)

func init() {
	if ok := externalTrust.AppendCertsFromPEM(_trustStoreExternal); !ok {
		debugLogChan <- _errTrustExit
		panic(_errTrustExit)
	}
	if ok := proxyTrust.AppendCertsFromPEM(_proxyCert); !ok {
		debugLogChan <- _errTrustExit
		panic(_errTrustExit)
	}
}

// squid mitm instance rootCA trust
//
//go:embed certstore/rootCA.pem
var _proxyCert []byte

// squid mitm mirrored external trust chain
//
//go:embed certstore/external_trust.pem
var _trustStoreExternal []byte
