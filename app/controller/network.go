package controller

import (
	"fmt"

	"github.com/aiyijing/smart-gateway/app/httputil"
	"github.com/aiyijing/smart-gateway/internal/sys"

	"github.com/gin-gonic/gin"
)

const (
	reload = "reload"
	flush  = "flush"
)

type NetworkController struct {
	mgr *sys.NetworkManger
}

func NewNetworkController(mgr *sys.NetworkManger) *NetworkController {
	return &NetworkController{mgr: mgr}
}

func (n *NetworkController) Action(g *gin.Context) {
	action := g.Query("action")
	switch action {
	case reload:
		httputil.Msg(g, n.reload())
		break
	case flush:
		httputil.Msg(g, n.flush())
		break
	default:
		httputil.MsgErr(g, fmt.Errorf("don't have any operate"))
	}
}

func (n *NetworkController) reload() error {

	return n.mgr.ReloadIptables()
}

func (n *NetworkController) flush() error {
	return n.mgr.FlushRule()
}
