package req

type Register struct {
	Username   string `json:"username"  binding:"required,min=2,max=12"`
	Password   string `json:"password"  binding:"required,min=6,max=16"`
	RePassword string `json:"re_password" binding:"required,min=6,max=16,eqfield=Password"`
}

type Login struct {
	Username  string `json:"username"  binding:"required,min=2,max=12"`
	Password  string `json:"password"  binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}
