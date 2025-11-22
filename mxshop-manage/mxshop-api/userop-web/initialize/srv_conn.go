package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	"github.com/zhengpanone/mxshop/mxshop-api/userop-web/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitSrvConn 初始化Srv连接
// https://blog.csdn.net/zhoupenghui168/article/details/131196225
func InitSrvConn() {

	// 商品服务 如果已有连接先关闭
	if global.GoodsConn != nil {
		global.GoodsConn.Close()
	}
	consul := global.ServerConfig.Consul
	url := fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.GoodsSrvConfig.Name)
	goodsConn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//轮询调度策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitGoodsSrvConn]连接【商品Srv服务失败】")
	}

	// 用户操作服务 如果已有连接先关闭
	if global.UserOpConn != nil {
		global.UserOpConn.Close()
	}

	userOpConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.UserOpSrvConfig.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//轮询调度策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitUserOpSrvConn]连接【用户操作Srv服务失败】")
	}
	global.GoodsConn = goodsConn
	global.UserOpConn = userOpConn

	// 注册客户端
	global.GoodsSrvClient = commonpb.NewGoodsClient(goodsConn)
	global.MessageSrvClient = commonpb.NewMessageClient(userOpConn)
	global.UserFavSrvClient = commonpb.NewUserFavClient(userOpConn)
	global.AddressSrvClient = commonpb.NewAddressClient(userOpConn)

}
