package service

import (
	"gin-demo/dao"
	"gin-demo/dao/req"
	"gin-demo/global"
	"gin-demo/global/response"
	"github.com/gin-gonic/gin"
)

func GetOnePost(id int, c *gin.Context) {
	var post dao.Post
	if global.GVA_DB.First(&post, id).RecordNotFound() {
		response.FailWithMessage("failure", "数据不存在", c)
	} else {
		response.OkWithData(&post, c)
	}
}

func CreatePost(blog req.CPBlog, c *gin.Context) {
	p := blog.CreateOrPut()
	if err := global.GVA_DB.Create(p).Error; err != nil {
		response.FailWithMessage("数据库创建失败", err.Error(), c)
	} else {
		response.OkWithData(p, c)
	}
}
