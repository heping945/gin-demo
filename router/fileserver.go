package router

import (
	"gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAliyunFileServerRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("file")

	{
		BaseRouter.POST("upload/", v1.UploadByFilename)
	}

}
