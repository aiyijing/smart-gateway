package main

import (
	"github.com/aiyijing/smart-gateway/app"
	"github.com/aiyijing/smart-gateway/internal/service"
	"github.com/aiyijing/smart-gateway/internal/sys"
	"github.com/aiyijing/smart-gateway/models"
)

func main() {
	netMgr := sys.NewNetworkManger(models.GetNetwork())
	err := netMgr.ReloadIptables()
	if err != nil {
		panic(err)
	}
	err = netMgr.EnableNat()
	if err != nil {
		panic(err)
	}
	_ = netMgr.DelViaRoute()
	err = netMgr.AddViaRoute()
	if err != nil {
		panic(err)
	}

	err = app.StartApp(service.Mgr, netMgr)
	if err != nil {
		panic(err)
	}
}
