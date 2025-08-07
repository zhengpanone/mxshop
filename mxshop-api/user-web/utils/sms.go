package utils

import (
	"context"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	commonGlobal "github.com/zhengpanone/mxshop/mxshop-api/common/global"
	"github.com/zhengpanone/mxshop/mxshop-api/user-web/global"
	"time"
)

func SendSms(mobile, smsCode string) (err error) {
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
	request.QueryParams["SignName"] = smsConfig.SignName                // 阿里云验证过的项目名
	request.QueryParams["TemplateCode"] = smsConfig.TemplateCode        //阿里云短信模版号
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" // 短信模版中验证码内容
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		return err
	}
	fmt.Println(client.DoAction(request, response))
	err = commonGlobal.RedisClient.Set(context.Background(), mobile, smsCode, 3*time.Minute).Err()
	return err
}
