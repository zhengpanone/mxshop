package request

type CaptchaForm struct {
	Height int `form:"height" json:"height" ` // 高度
	Width  int `form:"width" json:"width" `   // 宽度
	Length int `form:"length" json:"length"`  // 长度
}
