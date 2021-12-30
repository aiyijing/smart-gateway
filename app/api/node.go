package api

import (
	"github.com/aiyijing/smart-gateway/app/httputil"
	"github.com/aiyijing/smart-gateway/database"
	"github.com/aiyijing/smart-gateway/models"
	"github.com/gin-gonic/gin"
)

func ListNode(g *gin.Context) {
	var nodes []models.Node
	database.GetInstance().Find(&nodes)
	httputil.MsgWithData(g, nodes)
}

func AddNode(g *gin.Context) {
	var node models.Node
	err := g.BindJSON(&node)
	if err != nil {
		httputil.MsgErr(g, err)
		return
	}
	rs := database.GetInstance().Create(&node)
	if rs.Error != nil {
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
	rs := database.GetInstance().Updates(&node)
	if rs.Error != nil {
		httputil.MsgErr(g, err)
		return
	}
	httputil.MsgWithData(g, node)
}

func DeleteNode(g *gin.Context) {
	var node models.Node
	name := g.Param("name")
	database.GetInstance().Delete(&node).Where("name = ?", name)
	httputil.MsgWithData(g, node)
}
