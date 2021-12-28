package models

import (
	"encoding/json"
	"fmt"
	"github.com/aiyijing/smart-gateway/config"
	"github.com/aiyijing/smart-gateway/internal/sys"
	"github.com/aiyijing/smart-gateway/internal/util"
	"path"
)

var network *sys.NetConfig

func init() {
	err := load()
	if err != nil {
		fmt.Printf("load network err: %v\n", err)
		network = &sys.NetConfig{
			LanCidr: "192.168.1.1/24",
			IgnoreCidr: []string{
				"127.0.0.1/32",
				"224.0.0.0/4",
				"255.255.255.255/32",
			},
			ProxyMark: 1,
			PoxyPort:  12345,
		}
	}
}

func UpdateNetwork(net *sys.NetConfig) error {
	network = net
	dir := config.RunConf.Data
	p := path.Join(dir, "network.json")

	data, err := json.Marshal(net)
	if err != nil {
		return nil
	}

	return util.WriteOnCreate(p, data)
}

func GetNetwork() *sys.NetConfig {
	return network
}

func load() error {
	dir := config.RunConf.Data
	p := path.Join(dir, "network.json")

	data, err := util.ReadAll(p)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, network)
	if err != nil {
		return err
	}
	return nil
}
