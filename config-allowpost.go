package squidr

// example config
// var _allowPost contains a generated list of all currently allowed git repo targets
var _allowPost = map[string]bool{
	// captcha
	"api.funcaptcha.com":        true,
	"api.hcaptcha.com":          true,
	"www.recaptcha.net":         true,
	"challenges.cloudflare.com": true,
	// webmail
	"account-api.proton.me": true,
	"account.proton.me":     true,
	"mail.proton.me":        true,
	"accounts.google.com":   true,
	"store.google.com":      true,
	"mail.google.com":       true,
	"mail.tutanota.com":     true,
	// social
	"news.ycombinator.com": true,
	"infosec.exchange":     true,
	// dev freebsd
	"bugs.freebsd.org":    true,
	"reviews.freebsd.org": true,
	"lists.freebsd.org":   true,
	// dev golang
	"pkg.go.dev": true,
	// dev repos
	"meta.sr.ht":                true,
	"gitlab-api.arkoselabs.com": true,
	"support.github.com":        true,
	"uploads.github.com":        true,
	// search & tools
	"duckduckgo.com":        true,
	"lite.duckduckgo.com":   true,
	"html.duckduckgo.com":   true,
	"tineye.com":            true,
	"compresspng.com":       true,
	"tinypng.com":           true,
	"mxtoolbox.com":         true,
	"www.openstreetmap.org": true,
	// openai api
	"api.openai.com":      true,
	"ogs.google.com":      true,
	"beta.openai.com":     true,
	"auth0.openai.com":    true,
	"labs.openai.com":     true,
	"platform.openai.com": true,
	"chat.openai.com":     true,
	// gov (german)
	"www.lbv-termine.de":           true,
	"api-iam.intercom.io":          true,
	"sso.arbeitsagentur.de":        true,
	"www.arbeitsagentur.de":        true,
	"con.arbeitsagentur.de":        true,
	"web.arbeitsagentur.de":        true,
	"rest.arbeitsagentur.de":       true,
	"web.intern.arbeitsagentur.de": true,
	// e-commerce & IT
	"support.hp.com":   true,
	"www.dell.com":     true,
	"www.reichelt.de":  true,
	"www.berrybase.de": true,
	"www.paypal.com":   true,
	"link.tink.com":    true,
	// public transport (german)
	"www.hvv.de":            true,
	"shop.hvv.de":           true,
	"accounts.bahn.de":      true,
	"fahrkarten.bahn.de":    true,
	"reiseauskunft.bahn.de": true,
	// IoT api
	"api.kaiterra.com": true,
	// telco provider (eu)
	"login.o2online.de": true,
	"www.o2online.de":   true,
}
