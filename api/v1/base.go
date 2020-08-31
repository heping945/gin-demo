package v1

import (
	"gin-demo/dao"
	"gin-demo/dao/req"
	"gin-demo/global/response"
	"gin-demo/service"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var R req.Register
	if err := c.ShouldBindJSON(&R); err != nil {
		response.FailWithMessage("数据验证错误", utils.HandleValidateError(err), c)
		return
	}
	user := &dao.User{Username: R.Username, Password: R.Password}
	service.Register(user, c)
}

func Login(c *gin.Context) {
	var L req.Login
	if err := c.ShouldBindJSON(&L); err != nil {
		//response.FailWithMessage(fmt.Sprintf("错误：%s",err), c)
		response.FailWithMessage("数据验证错误", utils.HandleValidateError(err), c)
		return
	}
	if store.Verify(L.CaptchaId, L.Captcha, true) {
		U := &dao.User{Username: L.Username, Password: L.Password}
		service.Login(U, c)
	} else {
		response.Fail("验证码验证失败", c)
	}

}
