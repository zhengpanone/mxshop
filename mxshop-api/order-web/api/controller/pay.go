package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	commonResponse "github.com/zhengpanone/mxshop/mxshop-api/common/response"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/global"
	"net/http"
)

type PayApi struct {
}

// AlipayDemo 支付宝支付测试Demo
func (*PayApi) AlipayDemo() {
	appID := "9021000128666235"
	// 应用私钥
	privateKey := "MIIEowIBAAKCAQEAkjCC0bKG4lwSie5CHuEESUsb7r29K12OIVFCXewqMmZ8x3V5woSZxupng7Qoq3om0jZ479a1ZVIJIh8i9wB2usT/CUK67BPNnoypgCZMLDh0An6vBO85HnWQoNSK9tcJIhS3/vq1clJ/i7BOvLy75P411+7zg6UyxVWoh66VnctKQFqeYrfNJHhNS+lyC7JgolvSpbgzqV1824ixEDl4+5/bTZ+6ApZSlTKCjm5RD2bT2UJjgfswkhGo0CdNXb2uL2nahsDdNg+BYi6466bM5oUYao5QWuEFO5AadWbfKTEQpjiXPvKJpLBEqLj15HV+Gjx6gnPO1NAd/ebjnuTt3wIDAQABAoIBAHjnuiooVro7n/GHphPX0i2z+uQW9K867uPLSvJW8gdBEA3+sLcZ5/zFvNsGU2SO4DCXcKobj2a+1GLuEYLrVUbeynckQ2ggcLyiZUVhZzpjbj7p+2I/X6Q7Y2RApLXF3v3a2Nn/C7YDWQ10wYoDJfsb6/gs3iWQqU7fq0ScNY2MYqqRDfbbfqdxk4bBLWdF6iNpLX5ZV77jow8nmP0EYfjVhXpGMwg/L8SM90c0O27yqpz7FdVU9r4POak0x7H+ji3CqZ7KshHpzJYY/r8hmE9B8oeA5SXX3j7OjYVc68oSebtyHR3qr5Fc2ozzukIqCGTRntUSID81dJN8paX9wbECgYEA4Sjkxh/VLg+tUWNmcCisce6JNvdDNIX6DIeDH/7k8yJFGGpkpdeZgsEvfEHiTqiUlcz+LDHCn2bbmD618h2ucDyGjPIaMQRtFCMOq+FU+r2MJ7+6/13ovk/y2xnUJuJ93DNgGhbgd5Fnq3nP0qpaWlilRZU0daegEmhfzsNf79sCgYEApjaRW9/1U7v9FPBghzePWSnnGhxlkhTVt++KPLOGAaw9/KIEH9dcGr8t+h9Dr8oeYWDPlxElMVLH/+JbPR9w2SNVTVQE3/5hW/0sl2gn8HpoC/NKIHz4QfoSkSvSHe5PqY2IILw7/E+pTvTTNRcUDWvnk4H7YInHu4Ocq4VAK00CgYBBSzeUkNSkR92N9ZJWQiVH4NGfw/KUP6n0ijOnSqagLzb0Tp4jTbNxrI4VrZFwGAkGq+ylakSzLwPNUZo3vQ3B9HtcUjTwNAVhyozNoUmgcOk8+afDuDrvPYYCranNIO84tRlQV4P+iIcUvf1bbRVIj95VoGAXImYUEHqAHj4q8wKBgHEr8HddG36DHoQ2U5Nd6jXsyRVHRoVbuFVAPaCtH85dx+sPKb3Adk8j4xtOVFkKRvtI3q/elbNqyRDawALzOHQwkbFQRu15GDN5Q/ZnjeI4hkW0xsEuNh8+NYwxCSUmEYnB+3FGmZVnbEF9g52/dADbetc+BropthxgNbm3xgR1AoGBAKIJtSY51NQjftwzcYeZgihgkXwfu4/hjqU0pYZS7SnnEYffsbbouicqNG+2WucZ/4YLIzzatIXydYCbOFhGD9AaWmVp/d00LRvr+nawfvMads9/eOrJPd876uSUEmWErixiAJmtTKCprwbFSVs3EL0id1C35svQRCSk5jTbmVXr"
	// 支付宝公钥
	aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsbPUJnYoLTI8huhSt2UIT3bjkp/aiNFDIXVcAEQJ9vBANaOmg5C8fuuYK7rjQtL+wb1Y2oij1YM4+j4ABFzkKdAPgrIQ1VYQ7DC/4rQ4rQ1bdI6CeHgO3rabesjejckIzTJz87Scg/orhbUZUmPhb+NcEjhstSGXzKpDTdwKEdhUFiR2U12rHqvZeNXWajKJnf2C1TZ5an4wZkdEijgxtYOqkTdc7/N33oTagJ4XsfiveNiz8UCaOU+4zC6xg+QPjDpdp1NwQlkeRjW4HRlw+7SkKD3cjkgMNUhz5N0SkxYAl645M1yUxRrMmTfxIZ0saVicyvgV41hMZhPRUVHB0QIDAQAB"
	var client, err = alipay.New(appID, privateKey, false)
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayPublicKey(aliPublicKey)
	if err != nil {
		panic(err)
	}

	var p = alipay.TradePagePay{}
	p.NotifyURL = "https://platform.frp.yn.asqy.net/#/login?redirect=%2Findex"
	p.ReturnURL = "https://portal.frp.yn.asqy.net/#/login?redirect=%2Findex"
	p.Subject = "慕学生鲜"
	p.OutTradeNo = "20240714151901"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(url.String())
}

// AliPayNotify 支付宝回调通知
func (*PayApi) AliPayNotify(ctx *gin.Context) {

	client, _ := alipay.New(global.ServerConfig.AliPayConfig.AppId, global.ServerConfig.AliPayConfig.PrivateKey, false)

	err := client.LoadAliPayPublicKey(global.ServerConfig.AliPayConfig.AliPublicKey)

	if err != nil {
		global.Logger.Error(err.Error())
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	notification, err := client.DecodeNotification(ctx.Request.Form)
	if err != nil {
		global.Logger.Error(err.Error())
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	global.Logger.Info(fmt.Sprintf("交易状态%s", notification.TradeStatus))
	_, err = global.OrderSrvClient.UpdateOrderStatus(context.Background(), &commonpb.OrderStatus{
		OrderSn: notification.OutTradeNo,
		Status:  string(notification.TradeStatus),
		// TODO payTime
	})
	if err != nil {
		global.Logger.Error("新建订单详情失败")
		commonResponse.ErrorWithCodeAndMsg(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	commonResponse.OkWithMsg(ctx, "success")
}
