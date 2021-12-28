package sys

import (
	"bytes"
	"github.com/aiyijing/smart-gateway/internal/util"
	"text/template"

)

var iptables = `
iptables -t mangle -N V2RAY
{{range .ignoreCidr -}}
iptables -t mangle -A V2RAY -d {{.}} -j RETURN
{{end -}}
iptables -t mangle -A V2RAY -d {{.lanCidr}} -p tcp -j RETURN
iptables -t mangle -A V2RAY -d {{.lanCidr}} -p udp ! --dport 53 -j RETURN
iptables -t mangle -A V2RAY -p udp -j TPROXY --on-port {{.poxyPort}} --tproxy-mark {{.proxyMark}}
iptables -t mangle -A V2RAY -p tcp -j TPROXY --on-port {{.poxyPort}} --tproxy-mark {{.proxyMark}}
iptables -t mangle -A PREROUTING -j V2RAY
`

type NetConfig struct {
	IgnoreCidr []string `json:"ignoreCidr"`
	LanCidr    string   `json:"lanCidr"`
	PoxyPort   int      `json:"poxyPort"`
	ProxyMark  int      `json:"proxyMark"`
}

func Render(config *NetConfig) (string, error) {
	t, err := template.New("rules").Parse(iptables)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	args := util.StructToMapByJsonTag(config)
	if err := t.Execute(&buffer, args); err != nil {
		return "", err
	}
	return string(buffer.Bytes()), nil
}
