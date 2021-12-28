package controller

import (
	"encoding/json"
	"fmt"

	"github.com/aiyijing/smart-gateway/app/httputil"
	"github.com/aiyijing/smart-gateway/internal/service"
	"github.com/aiyijing/smart-gateway/internal/template"
	"github.com/aiyijing/smart-gateway/models"

	"github.com/gin-gonic/gin"
)

const (
	start   = "start"
	stop    = "stop"
	restart = "restart"
)

type V2rayController struct {
	mgr *service.Manager
}

func NewV2rayController(mgr *service.Manager) *V2rayController {
	return &V2rayController{mgr: mgr}
}

func (v *V2rayController) Action(g *gin.Context) {
	action := g.Query("action")
	switch action {
	case start:
		httputil.Msg(g, v.start())
		break
	case restart:
		httputil.Msg(g, v.restart())
		break
	case stop:
		httputil.Msg(g, v.stop())
		break
	default:
		httputil.MsgErr(g, fmt.Errorf("don't have any operate"))
	}

}

func (v *V2rayController) UpdateV2rayRawConfig(g *gin.Context) {
	data, err := g.GetRawData()
	if err != nil {
		httputil.Msg(g, err)
	}
	err = v.mgr.UpdateConfig("v2ray", string(data))
	httputil.Msg(g, err)
}

func (v *V2rayController) GetV2rayRawConfig(g *gin.Context) {
	data, err := v.mgr.GetCurrentConfig("v2ray")
	if err != nil {
		httputil.Msg(g, err)
	}
	httputil.MsgWithJsonString(g, data)
}

func (v *V2rayController) restart() error {
	nodes, err := models.GetAllNode()
	if err != nil {
		return err
	}
	var availableNode models.Node
	for _, n := range nodes {
		if n.Enable ==true{
			availableNode = n
		}
	}
	outbound, err := json.MarshalIndent(availableNode.OutBound, "", "  ")
	if err != nil {
		return fmt.Errorf("unable marshal outbound: %s", availableNode.Name)
	}

	tplArgs := &template.V2rayTplArgs{
		OutBounds: []string{
			string(outbound),
		},
	}

	newConfig, err := template.Render(tplArgs)
	if err != nil {
		return fmt.Errorf("unable render v2ray tpl: %v", err)
	}

	err = v.mgr.UpdateConfig("v2ray", newConfig)
	if err != nil {
		return fmt.Errorf("update v2ray config: %v", err)
	}

	return v.mgr.RestartService("v2ray")
}

func (v *V2rayController) stop() error {
	return v.mgr.StopService("v2ray")
}

func (v *V2rayController) start() error {
	return v.mgr.StartService("v2ray")
}
