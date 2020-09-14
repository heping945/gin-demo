package v1

import (
	"gin-demo/global/response"
	"gin-demo/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"path/filepath"
)

func UploadByFilename(c *gin.Context) {
	filename := c.DefaultQuery("filename", "")
	var err error
	ext := filepath.Ext(filename)
	if ext != ".img" {
		response.Fail("格式不对", c)
		return
	}
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	key := "upload/avatar/" + uuid.Must(uuid.NewV4(), err).String() + ext
	// 签名直传。
	bucket := utils.AliyunOssObject()
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		response.FailWithMessage("上传失败", err.Error(), c)
		return
	}
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		response.FailWithMessage("上传失败", err.Error(), c)
		return
	}
	Data := map[string]string{
		"key": key,
		"put": signedPutURL,
		"get": signedGetURL,
	}
	response.OkWithData(Data, c)
}
