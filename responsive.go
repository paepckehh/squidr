package squidr

import (
	"bytes"
	"io"
	"net/http"
	"sync"
	"time"
)

var responders, failed sync.Map // cache connectivity state

func isReachableViaProxy(targetMap rMap, idx int) bool {
	// check cache state
	target := targetMap.targets[idx]
	if _, ok := failed.Load(target); ok {
		return fail(target)
	}

	if _, ok := responders.Load(target); ok {
		return responsive(target)
	}

	// setup new https request
	client := getClient(getTransportProxy(_proxyURL, getTlsConf(proxyTrust)))
	request, err := getRequest(_https+target+targetMap.validateURL, _userAgent)
	if err != nil {
		return false
	}

	// get head first, warm caches
	request.Method = "HEAD"
	client.Timeout = time.Duration(5 * time.Second)
	resp, err := client.Do(request)
	if targetMap.headerOnly {
		if err != nil || resp.StatusCode != 200 || invalidBodySize(resp, targetMap) {
			return fail(target)
		}
		return responsive(target)
	}

	// get validateURL
	request.Method = "GET"
	client.Timeout = time.Duration(6 * time.Second)
	resp, err = client.Do(request)
	if err != nil || resp.StatusCode != 200 || invalidBodySize(resp, targetMap) || isInternalErrPage(resp) {
		return fail(target)
	}

	// reachable
	return responsive(target)
}

// isInternalErrPage response true if answer is an silent redirect to interal error
func isInternalErrPage(resp *http.Response) bool {
	head := make([]byte, _checkHeadLen)
	if _, err := io.ReadAtLeast(resp.Body, head, _checkHeadLen); err != nil {
		return false
	}
	defer resp.Body.Close()
	return bytes.Equal(head, []byte(_head))
}

// invalidBodySize response true if size requirement switch on an not meet
func invalidBodySize(resp *http.Response, targetMap rMap) bool {
	if targetMap.minSize < resp.ContentLength {
		return false
	}
	body, err := decodeResponse(resp)
	if err != nil {
		debugLogChan <- "### INTERNAL ERROR: unable to read or decompress response" + err.Error()
		return false
	}
	if targetMap.minSize < int64(len(body)) {
		return false
	}
	return true
}

// fail returns always false, syncs fail/responders stores for target
func fail(target string) bool {
	failed.Store(target, true)
	if _, ok := responders.Load(target); ok {
		responders.Delete(target)
	}
	return false
}

// responsive returns always true, syncs fail/responders stores for target
func responsive(target string) bool {
	responders.Store(target, true)
	if _, ok := failed.Load(target); ok {
		failed.Delete(target)
	}
	return true
}
