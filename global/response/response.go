package response

import (
	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 1001
	ERROR   = 1004
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

func Result(code int, data interface{}, msg, err string, c *gin.Context) {
	// 开始时间
	c.JSON(200, Response{
		code,
		data,
		msg,
		err,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", "", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, "", c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", "", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, "", c)
}

func Fail(err string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", err, c)
}

func FailWithMessage(message, err string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, err, c)
}

func FailWithDetailed(code int, data interface{}, message, err string, c *gin.Context) {
	Result(code, data, message, err, c)
}
