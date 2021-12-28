package models

import (
	"github.com/aiyijing/smart-gateway/config"
	"github.com/sdomino/scribble"
)

var DB *scribble.Driver

func init() {
	var err error
	DB, err = scribble.New(config.RunConf.Data, nil)
	if err != nil {
		panic(err)
	}
}
