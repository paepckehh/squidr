package squidr

import (
	"bytes"
	"crypto/tls"
	"strconv"
	"time"

	"paepcke.de/squidr/openssltrace"
	"paepcke.de/tlsinfo"
)

const (
	_checkHeadLen = 27
	_head         = "<html><title>[squidr][report]"
)

func reportTLS(host, path string, conn *tls.Conn, connected bool, err error) error {
	var buf bytes.Buffer
	reportTLSHead(host, path, &buf)
	reportTLSErr(err, &buf)
	reportTLSTls(conn, connected, &buf)
	reportTLSDns(host, &buf)
	reportTLSOssl(host, &buf)
	reportTLSFooter(&buf)
	return writeReport(buf.Bytes(), _servTLS)
}

func reportTLSHead(host, path string, buf *bytes.Buffer) {
	ts := time.Now()
	if path == "" {
		path = _none
	}
	buf.WriteString(_head + "[tls]</title><pre>\n")
	buf.WriteString("[ -= SQUIDR TLS REPORT =- ]\n")
	buf.WriteString("Report Timestamp: " + strconv.FormatInt(ts.UnixNano(), 10) + _linefeed)
	buf.WriteString("Report Created  : " + ts.Format(time.RFC3339) + _linefeed)
	buf.WriteString("Target Domain   : " + host + _linefeed)
	buf.WriteString("Target Path     : " + path + _linefeed)
	for _, u := range tlsinfo.ExtCheckURLs {
		buf.WriteString("External Report : <a href=\"" + u + host + "\" target=\"_blank\">" + u + host + "</a>" + _linefeed)
	}
}

func reportTLSFooter(buf *bytes.Buffer) {
	buf.WriteString("</pre></html>\n")
}

func reportTLSErr(err error, buf *bytes.Buffer) {
	if err != nil {
		buf.WriteString("\nGolang Runtime Environment Error Message\n")
		buf.WriteString(err.Error())
		buf.WriteString(_linefeed)
		buf.WriteString(_linefeed)
	}
}

func reportTLSTls(conn *tls.Conn, connected bool, buf *bytes.Buffer) {
	if connected {
		buf.WriteString("\nGolang TLS Runtime Environment State Summary Report\n\n")
		buf.WriteString(tlsinfo.ReportConnText(conn))
	}
}

func reportTLSDns(host string, buf *bytes.Buffer) {
	buf.WriteString("\nDNS Environment State Summary Report (local resolver)\n\n")
	reportLocal.Query, reportGoogle.Query, reportCloudflare.Query = host, host, host
	buf.WriteString(reportLocal.Generate() + _linefeed)
}

func reportTLSOssl(host string, buf *bytes.Buffer) {
	buf.WriteString("\nOpenSSL TLS Trace Summary Report\n\n")
	buf.WriteString(openssltrace.OpenSSLReportHostText(host, _sslCaTrust, true) + _linefeed)
}
