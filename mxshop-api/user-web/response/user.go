package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stmp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"nickName"`
	//Birthday time.Time `json:"birthday"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
}

type RegisterResponse struct {
	Id        string `form:"id" json:"id"`
	Account   string `form:"account" json:"account" ` // 手机号码
	Nickname  string `form:"nickname" json:"nickname" `
	Token     string `form:"token" json:"token" `
	ExpiredAt int64  `form:"expiredAt" json:"expiredAt" `
}
