package config

import "github.com/spf13/pflag"

type RunConfig struct {
	Address string `json:"address"`
	Data    string `json:"data"`
	Static  string `json:"static"`
	Debug   string `json:"debug"`
}

var RunConf = RunConfig{}

func addRunConfFlag() {
	pflag.StringVar(&RunConf.Address, "address", ":80", "listen network")
	pflag.StringVar(&RunConf.Data, "data", "/etc/gateway/", "user data dir")
	pflag.StringVar(&RunConf.Static, "static", "/static", "web static dir")
	pflag.StringVar(&RunConf.Debug, "debug", "release", "gin debug mode")
}

func init() {
	addRunConfFlag()
	pflag.Parse()
}
