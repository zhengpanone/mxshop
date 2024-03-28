package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` // 手机号码 自定义validator
	Type   string `form:"type" json:"type" binding:"required,oneof=1 2"`  // 1、注册发送验证码，2、动态验证码发送验证码，校验手机号码是否在系统中存在
}
