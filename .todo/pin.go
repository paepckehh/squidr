package squidr

import "time"

type pin struct {
	dns []string
	pin [32]byte
	exp time.Time
}

type rawPIN struct {
	dns string
	pin string
	exp string
}

var rawPINs = []rawPIN{
	dns: "DNS:*.pr.tn, DNS:*.proton.me, DNS:*.storage.proton.me, DNS:pr.tn, DNS:proton.me",
	pin: "CT56BhOTmj5ZIPgb/xD5mH8rY3BLo/MlhP7oPyJUEDo=",
	exp: "2023-01-31 13:11:28 +0000 UTC",
}
