package main

import "text/template"

var proxiesTemplate *template.Template

func init() {
	proxiesTemplate = template.Must(template.New("proxies").Parse(ProxiesTemplate))
}

// ProxiesTemplate is the configuration template
// for the `proxies` endpoint.
var ProxiesTemplate = `
[proxies]
type=endpoint
transport=transport-udp
context=default
disallow=all
allow=ulaw
aors=proxies
ice_support=no
rtp_symmetric=yes

[proxies]
type=aor
{{range .}}
contact=sip:{{.}}:5060
{{end}}

[proxies]
type=identify
endpoint=proxies
{{range .}}
match={{.}}
{{end}}

[acl]
type=acl
deny=0.0.0.0/0.0.0.0
permit=10.0.0.0/255.0.0.0
permit=10.2.0.0/255.255.0.0
permit=10.3.0.0/255.255.0.0
{{range .}}
permit={{.}}
{{end}}
`
