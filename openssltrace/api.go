// package openssltrace is an openssl wrapper for golang
package openssltrace

import (
	"os/exec"
	"strings"

	"paepcke.de/reportstyle"
)

//
// TEXT REPORT
//

// OpenSSLReportHostText report an client connect trace as text
func OpenSSLReportHostText(host, sslCaTrustFile string, debug bool) string {
	return OpenSSLReportHost(host, sslCaTrustFile, GetBinary(), reportstyle.StyleText(), debug)
}

//
// GENERIC BACKEND
//

// OpenSSLReportHost generic
func OpenSSLReportHost(host, sslCaTrustFile, openSSLBinary string, e *reportstyle.Style, debug bool) string {
	if !isExec(openSSLBinary) {
		return "[openssltrace] openssl binary not found: " + openSSLBinary
	}
	var s strings.Builder
	s.WriteString(e.L1 + "[openssl:setup]" + e.LE)
	s.WriteString(e.L1 + " OpenSSL Binary         " + openSSLBinary + e.LE)
	s.WriteString(e.L1 + " OpenSSL Version        " + OpenSSLVersion(openSSLBinary) + e.LE)
	s.WriteString(e.L1 + " CA Trust bundle file   " + sslCaTrustFile + e.LE)
	s.WriteString(e.L1 + "[openssl:trace]" + e.LE)
	s.WriteString(OpenSSLTrace(host, sslCaTrustFile, openSSLBinary, e, debug))
	return s.String()
}

// OpenSSLVersion report
func OpenSSLVersion(openSSLBinary string) string {
	if !isExec(openSSLBinary) {
		return "[openssltrace] openssl binary not found " + openSSLBinary
	}
	var s strings.Builder
	cmd := exec.Command(openSSLBinary, _version)
	cmd.Stdout = &s
	cmd.Stderr = &s
	err := cmd.Start()
	if err != nil {
		return _empty
	}
	err = cmd.Wait()
	if err != nil {
		return _empty
	}
	out := s.String()
	return out[:len(out)-1]
}

// OpenSSLTrace report
func OpenSSLTrace(host, sslCaTrustFile, openSSLBinary string, e *reportstyle.Style, debug bool) string {
	if !isExec(openSSLBinary) {
		return "[openssltrace] openssl binary not found: " + openSSLBinary
	}
	var s strings.Builder
	s.WriteString(e.PS)
	if debug {
		cmd := exec.Command(openSSLBinary, _sclient, _cafile, sslCaTrustFile, _showcerts, _msg, _state, _debug, host)
		cmd.Stdout = &s
		cmd.Stderr = &s
		err := cmd.Run()
		if err != nil {
			return _empty
		}
	} else {
		cmd := exec.Command(openSSLBinary, _sclient, _cafile, sslCaTrustFile, _showcerts, _state, host)
		cmd.Stdout = &s
		cmd.Stderr = &s
		err := cmd.Run()
		if err != nil {
			return _empty
		}
	}
	s.WriteString(e.PE)
	return s.String()
}

//
// LITTLE HELPER
//

// GetBinary ...
func GetBinary() string {
	executeable, err := exec.LookPath(_openssl)
	if err != nil || executeable == _empty {
		return err.Error()
	}
	return executeable
}
