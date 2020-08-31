package service

import (
	"gin-demo/dao"
	"gin-demo/dao/req"
	"gin-demo/dao/res"
	"gin-demo/global"
	"gin-demo/global/response"
	"gin-demo/middleware"
	"gin-demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"time"
)

func Register(user *dao.User, c *gin.Context) {
	username := user.Username
	notRegister := global.GVA_DB.Where("username = ?", username).First(&user).RecordNotFound()
	// notRegister为false表明读取到了 不能注册
	if !notRegister {
		response.Fail("用户已注册", c)
	} else {
		// 否则 附加uuid 密码md5简单加密 注册
		user.Password = utils.MD5V([]byte(user.Password))
		user.UUID = uuid.NewV4()
		if err := global.GVA_DB.Create(&user).Error; err != nil {
			response.Fail("注册用户数据失败", c)
		} else {
			//	注册成功
			response.OkDetailed(user, "创建用户成功", c)
		}
	}
}

func Login(user *dao.User, c *gin.Context) {
	user.Password = utils.MD5V([]byte(user.Password))
	notFound := global.GVA_DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).RecordNotFound()
	if notFound {
		response.Fail("账户或密码错误", c)
	} else {
		tokenNext(*user, c)
	}
}

// 登录以后签发jwt
func tokenNext(user dao.User, c *gin.Context) {
	j := &middleware.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := req.CustomClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "zhp",                          // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.Fail("获取token失败", c)
		return
	}

	response.OkDetailed(res.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)

}
