package v1

import (
	"gin-demo/dao/req"
	"gin-demo/global/response"
	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
)

func GetPostList(c *gin.Context) {

}

func GetSinglePost(c *gin.Context) {
	id := c.Param("id")
	p, ok := utils.ValidateParam(id, -1)
	if !ok {
		response.FailWithMessage("failure", "参数错误", c)
		return
	}
	service.GetOnePost(p, c)

}

func CreatePost(c *gin.Context) {
	var cb req.CPBlog
	if err := c.ShouldBindJSON(&cb); err != nil {
		response.FailWithMessage("数据验证错误", utils.HandleValidateError(err), c)
		return
	}
	service.CreatePost(cb, c)

}
func DeletePost(c *gin.Context) {

}
func UpdatePost(c *gin.Context) {

}
