package openssltrace

import (
	"os/exec"
)

const (
	_version   = "version"
	_sclient   = "s_client"
	_cafile    = "-CAfile"
	_msg       = "-msg"
	_state     = "-state"
	_debug     = "-debug"
	_showcerts = "-showcerts"
	_openssl   = "openssl"
	_empty     = ""
	_linefeed  = "\n"
)

func isExec(file string) bool {
	if executeable, err := exec.LookPath(file); err != nil || executeable == _empty {
		return false
	}
	return true
}
