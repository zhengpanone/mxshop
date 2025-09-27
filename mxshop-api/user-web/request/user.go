package request

type PasswordLoginForm struct {
	Account     string `form:"account" json:"account" binding:"required,mobile"`         // 手机号码 自定义validator
	Password    string `form:"password" json:"password" binding:"required,min=3,max=20"` // 密码
	CaptchaText string `form:"captchaText" json:"captchaText"`
	CaptchaId   string `form:"captchaId" json:"captchaId"`
}

type RegisterForm struct {
	Account  string `form:"account" json:"account" binding:"required,mobile"`         // 手机号码 自定义validator
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"` // 密码
	Code     string `form:"code" json:"code" binding:"required,min=5,max=5"`          // 短信验证码
}
