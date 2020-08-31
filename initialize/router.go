package initialize

import (
	"gin-demo/global"
	"gin-demo/middleware"
	"gin-demo/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	Router := gin.Default()
	Router.Use(middleware.Cors())
	global.GVA_LOG.Debug("启动了cors中间件")
	ApiGroup := Router.Group("api/v1")
	{
		router.InitBaseRouter(ApiGroup)
	}
	return Router
}
