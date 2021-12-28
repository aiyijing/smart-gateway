package app

import (
	"github.com/aiyijing/smart-gateway/app/api"
	"github.com/aiyijing/smart-gateway/app/controller"
	"github.com/aiyijing/smart-gateway/app/middleware"
	"github.com/aiyijing/smart-gateway/config"
	"github.com/aiyijing/smart-gateway/internal/service"
	"github.com/aiyijing/smart-gateway/internal/sys"

	"github.com/gin-gonic/gin"
)

func registerApi(restful *gin.RouterGroup) {
	restful.GET("/nodes", api.ListNode)
	restful.POST("/nodes", api.AddNode)
	restful.PUT("/nodes/:name", api.UpdateNode)
	restful.DELETE("/nodes/:name", api.DeleteNode)

}

func registerController(serviceManager *service.Manager,
	networkManager *sys.NetworkManger, g *gin.RouterGroup,
) {
	v2ray := controller.NewV2rayController(serviceManager)
	v2rayRouter := g.Group("v2ray")
	{
		v2rayRouter.GET("/action", v2ray.Action)
		v2rayRouter.PUT("/config", v2ray.UpdateV2rayRawConfig)
		v2rayRouter.GET("/config", v2ray.GetV2rayRawConfig)
	}
	network := controller.NewNetworkController(networkManager)
	networkRouter := g.Group("network")
	{
		networkRouter.GET("/action", network.Action)
	}
}

func addMiddleware(router *gin.Engine) {
	router.Use(
		middleware.Serve("/", middleware.LocalFile(
			config.RunConf.Static, true,
		)),
	)
}

func noRoute(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.File(config.RunConf.Static)
	})
}

func StartApp(serviceManager *service.Manager, networkManager *sys.NetworkManger) error {
	gin.SetMode(config.RunConf.Debug)

	router := gin.Default()
	noRoute(router)
	addMiddleware(router)

	registerApi(router.Group("api"))
	registerController(serviceManager, networkManager, router.Group("controller"))

	err := router.Run(config.RunConf.Address)
	return err
}
