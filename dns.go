package squidr

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"paepcke.de/dnsinfo"
	"paepcke.de/dnsresolver"
	"paepcke.de/reportstyle"
)

var (
	reportGoogle = &dnsinfo.Report{
		Type:     dnsresolver.TypeAll,
		Summary:  true,
		Style:    reportstyle.StyleText(),
		Resolver: dnsresolver.ResolverViaProvider("google", true),
	}
	reportCloudflare = &dnsinfo.Report{
		Type:     dnsresolver.TypeAll,
		Summary:  true,
		Style:    reportstyle.StyleText(),
		Resolver: dnsresolver.ResolverViaProvider("cloudflare", true),
	}
	reportQuad9 = &dnsinfo.Report{
		Type:     dnsresolver.TypeAll,
		Summary:  true,
		Style:    reportstyle.StyleText(),
		Resolver: dnsresolver.ResolverViaProvider("quad9", true),
	}
	reportLocal = &dnsinfo.Report{
		Type:    dnsresolver.TypeAll,
		Summary: true,
		Style:   reportstyle.StyleText(),
		Resolver: &dnsresolver.Resolver{
			Server:  "127.0.0.53:53",
			NoIP6:   true,
			Timeout: 6 * time.Second,
		},
	}
)

func reportDNS(domain string) error {
	s := strings.Split(domain, _slashfwd)
	domain = s[0]
	var buf bytes.Buffer
	reportDNSHead(domain, &buf)
	reportDNSBody(domain, &buf)
	reportDNSFooter(&buf)
	return writeReport(buf.Bytes(), _servDNS)
}

func reportDNSHead(domain string, buf *bytes.Buffer) {
	ts := time.Now()
	buf.WriteString(_head + "[dns]</title><pre>\n")
	buf.WriteString("[ -= SQUIDR DNS REPORT =- ]\n")
	buf.WriteString("Report Timestamp: " + strconv.FormatInt(ts.UnixNano(), 10) + _linefeed)
	buf.WriteString("Report Created  : " + ts.Format(time.RFC3339) + _linefeed)
	buf.WriteString("Target DNS      : " + domain + _linefeed)
}

func reportDNSFooter(buf *bytes.Buffer) {
	buf.WriteString("</pre></html>\n")
}

func reportDNSBody(domain string, buf *bytes.Buffer) {
	buf.WriteString("\nDNS Environment State Summary Report\n\n")
	reportLocal.Query, reportGoogle.Query, reportCloudflare.Query, reportQuad9.Query = domain, domain, domain, domain
	buf.WriteString(reportLocal.Generate() + _linefeed)
	buf.WriteString(reportGoogle.Generate() + _linefeed)
	buf.WriteString(reportCloudflare.Generate() + _linefeed)
	buf.WriteString(reportQuad9.Generate() + _linefeed + _linefeed)
}
