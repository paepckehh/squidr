package squidr

import (
	"bytes"
	"strconv"
	"time"
)

func reportSEARX(query string) error {
	var buf bytes.Buffer
	reportSEARXHead(query, &buf)
	reportSEARXBody(query, &buf)
	reportSEARXFooter(&buf)
	return writeReport(buf.Bytes(), _servSSS)
}

func reportSEARXHead(query string, buf *bytes.Buffer) {
	ts := time.Now()
	buf.WriteString(_head + "[url]</title><pre>\n")
	buf.WriteString("[ -= SQUIDR SEARCH =- ]\n")
	buf.WriteString("Report Timestamp  : " + strconv.FormatInt(ts.UnixNano(), 10) + _linefeed)
	buf.WriteString("Report Created    : " + ts.Format(time.RFC3339) + _linefeed)
	buf.WriteString("Available Engines : " + itoa(len(_searx)) + _linefeed)
	buf.WriteString("Target Search     : " + query + _linefeed + _linefeed)
}

func reportSEARXFooter(buf *bytes.Buffer) {
	buf.WriteString("</pre></html>\n")
}

func reportSEARXBody(query string, buf *bytes.Buffer) {
	targetURL := _https + "searx.net/?q=" + query
	buf.WriteString(_linefeed)
	buf.WriteString("verified random engine link\n")
	buf.WriteString("<a href=\"" + targetURL + "\" target=\"_blank\" rel=\"noreferrer\">" + targetURL + "</a>" + _linefeed)
	buf.WriteString(_linefeed)
	for _, domain := range _searx {
		targetURL := _https + domain + "/?q=" + query
		buf.WriteString("<a href=\"")
		buf.WriteString(targetURL)
		buf.WriteString("\" target=\"_blank\" rel=\"noreferrer\">")
		buf.WriteString(targetURL)
		buf.WriteString("</a>")
		buf.WriteString(_linefeed)
	}
}
