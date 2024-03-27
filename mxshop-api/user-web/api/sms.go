package api

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"golang.org/x/exp/rand"
	"mxshop-api/user-web/global"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {
	// 生成width的短信验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(uint64(time.Now().UnixNano()))
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms() {
	smsConfig := global.ServerConfig.SMSConfig
	client, err := dysmsapi.NewClientWithAccessKey(smsConfig.RegionId, smsConfig.AccessKeyId, smsConfig.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = smsConfig.Protocol // https | http
	request.Domain = smsConfig.Domain
	request.Version = "2017-05-25"
	request.ApiName = smsConfig.ApiName
	request.QueryParams["RegionId"] = smsConfig.RegionId
	request.QueryParams["PhoneNumbers"] = "15527300572"
	request.QueryParams["SignName"] = smsConfig.SignName               // 阿里云验证过的项目名
	request.QueryParams["TemplateCode"] = smsConfig.TemplateCode       //阿里云短信模版号
	request.QueryParams["TemplateParam"] = "{\"code\":" + "7777" + "}" // 短信模版中验证码内容
	response, err := client.ProcessCommonRequest(request)
	fmt.Println(client.DoAction(request, response))
}
