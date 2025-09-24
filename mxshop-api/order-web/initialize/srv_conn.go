package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	commonpb "github.com/zhengpanone/mxshop/mxshop-api/common/proto/pb"
	"github.com/zhengpanone/mxshop/mxshop-api/order-web/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// https://blog.csdn.net/zhoupenghui168/article/details/131196225
func InitSrvConn() {

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
	)
	if err != nil {
		zap.S().Fatal("[InitGoodsSrvConn]连接【商品Srv服务失败】")
	}
	global.GoodsConn = goodsConn
	// 注册客户端
	goodsSrvClient := commonpb.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient

	// 如果已有连接先关闭
	if global.OrderConn != nil {
		global.OrderConn.Close()
	}

	orderConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.OrderSrvConfig.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//轮询调度策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitOrderSrvConn]连接【订单Srv服务失败】")
	}
	global.OrderConn = orderConn
	// 注册客户端
	orderSrvConn := commonpb.NewOrderClient(orderConn)
	global.OrderSrvClient = orderSrvConn

	// 如果已有连接先关闭
	if global.InventoryConn != nil {
		global.InventoryConn.Close()
	}

	inventoryConn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.InventorySrvConfig.Name),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//轮询调度策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitInventorySrvConn]连接【库存Srv服务失败】")
	}
	global.InventoryConn = inventoryConn

	// 注册客户端
	inventorySrvClient := commonpb.NewInventoryClient(inventoryConn)
	global.InventorySrvClient = inventorySrvClient
}
