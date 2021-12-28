package service

import (
	"fmt"
	"github.com/docker/docker/client"
)

var Mgr *Manager
var cli *client.Client

func init() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}

	Mgr = NewManager()
	v2ray := NewV2ray("v2ray", cli)
	Mgr.AddService("v2ray", v2ray)
}
