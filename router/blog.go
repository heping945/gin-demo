package router

import (
	"gin-demo/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBlogRouter(Router *gin.RouterGroup) {
	PostRouter := Router.Group("posts")

	{
		PostRouter.GET("/", v1.GetPostList)
		PostRouter.GET("/:id", v1.GetSinglePost)
		PostRouter.POST("/", v1.CreatePost)
		PostRouter.DELETE("/:id", v1.DeletePost)
		PostRouter.PUT("/:id", v1.UpdatePost)
	}

	CategoryRouter := Router.Group("category")
	{
		CategoryRouter.GET("/", v1.Register)
	}
}
