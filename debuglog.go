package squidr

import (
	"os"
)

const (
	// config
	_debugLogFile = "/var/tmp/squidrdebug.log"
	// messages
	_loghdr        = "[squidr] "
	_ioDebugIN     = "[ in]  "
	_ioDebugOUT    = "[out] "
	_ioDebugDomain = "[domain:identified] "
	_ioDebugDNS    = "[dns] [err] "
	_ioDebugCAA    = "[dns] [caa] "
	_ioMethod      = "[http:methode] "
	_ioNoChain     = "[tls] [fail] [no trusted chain found] "
	_ioConnErr     = "[tls] [fail] [connect err] "
	_redirectHTTPS = "[redirect:https] "
	_broken        = "[request] [broken] "
	_cacheHit      = "[cache] [hit] "
	_connectExit   = "[method:connect] [exit] "
	_nopost        = "[method:post] [acl:no] "
	_terminated    = "[terminated] "
	_dnsterminated = "[dns] [fail] "
	_tlsterminated = "[tls] [fail] "
	_errcode       = " error code : "
	_space         = " : "
	_point         = " -> "
)

func debugLog() {
	go func() {
		f, err := os.OpenFile(_debugLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			panic("[squidr] [debuglog] [log:file] [unable to create log file]")
		}
		defer f.Close()
		for msg := range debugLogChan {
			if _, err := f.Write([]byte(_loghdr + msg + _linefeed)); err != nil {
				panic("[squidr] [debuglog] [log:file] [unable to write to log file])")
			}
		}
	}()
}
