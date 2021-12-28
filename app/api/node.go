package api

import (
	"github.com/aiyijing/smart-gateway/app/httputil"
	"github.com/aiyijing/smart-gateway/models"
	"github.com/gin-gonic/gin"
)

func ListNode(g *gin.Context) {
	nodes, err := models.GetAllNode()
	if err != nil {
		httputil.MsgErr(g, err)
		return
	}
	httputil.MsgWithData(g, nodes)
}

func AddNode(g *gin.Context) {
	var node models.Node
	err := g.BindJSON(&node)
	if err != nil {
		httputil.MsgErr(g, err)
		return
	}
	err = models.AddNode(node)
	if err != nil {
		httputil.MsgErr(g, err)
		return
	}
	httputil.MsgWithData(g, node)
}

func UpdateNode(g *gin.Context) {
	var node models.Node
	err := g.BindJSON(&node)
	if err != nil {
		httputil.MsgErr(g, err)
		return
	}

	err = models.UpdateNode(node)
	if err != nil {
		httputil.MsgErr(g, err)
		return
	}

	httputil.MsgWithData(g, node)
}

func DeleteNode(g *gin.Context) {
	name := g.Param("name")
	err := models.DeleteNode(name)
	httputil.Msg(g, err)
}
