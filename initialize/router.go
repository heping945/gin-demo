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
		router.InitBaseRouter(ApiGroup)             // 基础路由，不做鉴权验证，包括登录，注册和验证码
		router.InitAliyunFileServerRouter(ApiGroup) // 阿里云的oss使用
		router.InitBlogRouter(ApiGroup)             // 一个blog简单的增删改查的实例
	}
	return Router
}
