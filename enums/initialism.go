package enums

import (
	"strings"
	"unicode"
)

//go:generate ctenum create enum initialism.go
//ctenum:package=enums make-file=true
const (
	acl   = "ACL"
	api   = "API"
	ascii = "ASCII"
	cpu   = "CPU"
	// Cascading style sheets
	css = "CSS"
	// eof is End of File.
	eof  = "EOF"
	gpu  = "GPU"
	guid = "GUID"
	// HyperTextMarkupLanguage.
	html  = "HTML"
	http  = "HTTP"
	https = "HTTPS"
	id    = "ID"
	ip    = "IP"
	json  = "JSON"
	jwt   = "JWT"
	ram   = "RAM"
	rest  = "REST"
	sql   = "SQL"
	tcp   = "TCP"
	tls   = "TLS"
	udp   = "UDP"
	uri   = "URI"
	url   = "URL"
	utf   = "UTF"
	uuid  = "UUID"
	xml   = "XML"
	yaml  = "YAML"
	zip   = "ZIP"
)

func FormatInitialism(s string) string {
	flds := func(s string) []string {
		var flds []string
		start := 0
		for i, r := range s {
			if !unicode.IsUpper(r) || i <= start {
				continue
			}
			flds = append(flds, s[start:i])
			start = i
		}
		return append(flds, s[start:])
	}(s)

	for i, f := range flds {
		maybe := strings.ToUpper(f)
		if _, err := ToInitialismEnum(maybe); err != nil {
			continue
		}
		flds[i] = maybe
	}
	return strings.Join(flds, "")
}
