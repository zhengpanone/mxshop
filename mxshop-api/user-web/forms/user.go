package forms

type PasswordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`           // 手机号码 自定义validator
	Password  string `form:"password" json:"password" binding:"required,min=3,max=20"` // 密码
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`           // 手机号码 自定义validator
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"` // 密码
	Code     string `form:"code" json:"code" binding:"required,min=5,max=5"`          // 短信验证码
}
