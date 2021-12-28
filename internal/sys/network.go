package sys

import (
	"fmt"
	"github.com/aiyijing/smart-gateway/internal/util"
	"strings"
)

type NetworkManger struct {
	Rule *NetConfig
}

func NewNetworkManger(config *NetConfig) *NetworkManger {
	return &NetworkManger{
		Rule: config,
	}
}

func (net *NetworkManger) FlushRule() error {
	err := util.Execute(
		"iptables -t mangle -F",
		"iptables -t mangle -X",
		"iptables -t mangle -Z",
	)
	if err != nil {
		return err
	}
	return nil
}

func (net *NetworkManger) ReloadIptables() error {
	if err := net.FlushRule(); err != nil {
		return err
	}
	rule, err := Render(net.Rule)
	if err != nil {
		return err
	}
	rules := strings.Split(rule, "\n")
	if err := util.Execute(rules...); err != nil {
		return err
	}
	return nil
}

func (net *NetworkManger) EnableNat() error {
	return net.nat(1)
}

func (net *NetworkManger) DisableNat() error {
	return net.nat(1)
}

func (net *NetworkManger) AddViaRoute() error {
	var cmds = []string{
		fmt.Sprintf("ip rule add fwmark %d table 100", net.Rule.ProxyMark),
		"ip route add local 0.0.0.0/0 dev lo table 100",
	}
	if err := util.Execute(cmds...); err != nil {
		return err
	}
	return nil
}

func (net *NetworkManger) DelViaRoute() error {
	var cmds = []string{
		fmt.Sprintf("ip rule del fwmark %d table 100", net.Rule.ProxyMark),
		"ip route del local 0.0.0.0/0 dev lo table 100",
	}
	if err := util.Execute(cmds...); err != nil {
		return err
	}
	return nil
}

func (net *NetworkManger) nat(enable int) error {
	var cmd = fmt.Sprintf("echo %d > /proc/sys/net/ipv4/ip_forward", enable)
	if err := util.Execute(cmd); err != nil {
		return err
	}
	return nil
}
