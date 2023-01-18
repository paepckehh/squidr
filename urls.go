package squidr

import (
	"bytes"
	"errors"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"mvdan.cc/xurls/v2"
)

func reportURL(targetURL string) error {
	targetURL = _https + targetURL
	var buf bytes.Buffer
	reportURLHead(targetURL, &buf)
	reportURLBody(targetURL, &buf)
	reportURLFooter(&buf)
	return writeReport(buf.Bytes(), _servURL)
}

func reportURLHead(targetURL string, buf *bytes.Buffer) {
	ts := time.Now()
	buf.WriteString(_head + "[url]</title><pre>\n")
	buf.WriteString("[ -= SQUIDR URL REPORT =- ]\n")
	buf.WriteString("Report Timestamp: " + strconv.FormatInt(ts.UnixNano(), 10) + _linefeed)
	buf.WriteString("Report Created  : " + ts.Format(time.RFC3339) + _linefeed)
	buf.WriteString("Target URL      : " + targetURL + _linefeed)
}

func reportURLFooter(buf *bytes.Buffer) {
	buf.WriteString("</pre></html>\n")
}

func reportURLBody(targetURL string, buf *bytes.Buffer) {
	body, err := getUrlsFromPageViaProxy(targetURL)
	if err != nil {
		buf.WriteString("[error] [unable to fetch]: " + targetURL + _sep + err.Error())
		return
	}
	buf.WriteString("Target PageSize : " + strconv.Itoa(len(body)) + " bytes" + _linefeed + _linefeed)
	buf.WriteString(urls(string(body), false, false))
}

func getUrlsFromPageViaProxy(targetURL string) ([]byte, error) {
	// sanitize url
	_, err := url.Parse(targetURL)
	if err != nil {
		return nil, errors.New("invalid url, unable to parse url: " + err.Error())
	}

	// setup
	client := getClient(getTransportProxy(_proxyURL, getTlsConf(proxyTrust)))
	request, err := getRequest(targetURL, _userAgent)
	if err != nil {
		return nil, errors.New("unable to fetch via proxy: " + err.Error())
	}

	// get
	request.Method = "GET"
	client.Timeout = time.Duration(5 * time.Second)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("invalid response: ")
	}
	return decodeResponse(resp)
}

func urls(in string, head, css bool) string {
	var s strings.Builder
	parse := xurls.Strict()
	array := parse.FindAllString(in, -1)
	uMap := make(map[string]bool)
	for _, v := range array {
		uMap[v] = true
	}
	uniq := make([]string, len(uMap))
	for k := range uMap {
		uniq = append(uniq, k)
	}
	sort.Strings(uniq)
	if len(uniq) > 0 {
		switch head {
		case true:
			s.WriteString("\n<br>\n[url list]")
		default:
		}
		switch css {
		case true:
			s.WriteString("\n\t<ol>")
		default:
		}
		for _, e := range uniq {
			if strings.Contains(e, "w3.org") || strings.Contains(e, "W3.org") || strings.Contains(e, "schema.org") {
				continue
			}
			if len(e) < 4 {
				continue
			}
			if e[:4] != "http" {
				e = "https://" + e
			}
			if len(e) < 5 {
				continue
			}
			if e[:5] == "http:" {
				e = "https:" + e[5:]
			}
			switch css {
			case true:
				s.WriteString("\n\t\t<li><a style=\"color:#AAA;text-decoration:none;\"} href=\"")
				s.WriteString(e)
				s.WriteString("\" target=\"_blank\" rel=\"noreferrer\">")
				s.WriteString(e)
				s.WriteString("</a></li><br>")
			case false:
				s.WriteString("<a href=\"")
				s.WriteString(e)
				s.WriteString("\" target=\"_blank\" rel=\"noreferrer\">")
				s.WriteString(e)
				s.WriteString("</a><br>")
			}
		}
		switch css {
		case true:
			s.WriteString("\n\t</ol>")
		default:
		}
	}
	return s.String()
}
