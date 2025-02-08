package initialize

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/zhengpanone/mxshop/common/interceptor"
	"github.com/zhengpanone/mxshop/common/utils"
	"github.com/zhengpanone/mxshop/goods-web/global"
	"github.com/zhengpanone/mxshop/goods-web/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitSrvConn 初始化Srv连接
// https://blog.csdn.net/zhoupenghui168/article/details/131196225
func InitSrvConn() {
	tracer, closer := utils.InitJaeger("goods-grpc", global.ServerConfig.Jaeger.Host, global.ServerConfig.Jaeger.Port)
	defer closer.Close()
	// 如果已有连接先关闭
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
		grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
			interceptor.ClientInterceptor(tracer))),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn]连接【商品Srv服务失败】")
	}
	global.GoodsConn = goodsConn
	// 注册客户端
	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient

}
