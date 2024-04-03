package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"user-web/global"
	"user-web/proto"
)

// https://blog.csdn.net/zhoupenghui168/article/details/131196225
func InitSrvConn() {

	// 如果已有连接先关闭
	if global.UserConn != nil {
		global.UserConn.Close()
	}
	consul := global.ServerConfig.Consul
	url := fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.UserSrvConfig.Name)
	userConn, err := grpc.Dial(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//轮询调度策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn]连接【用户服务失败】")
	}
	global.UserConn = userConn
	// 注册客户端
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = &userSrvClient

}

func InitSrvConnBack() {

	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	consul := global.ServerConfig.Consul
	cfg.Address = fmt.Sprintf("%s:%d", consul.Host, consul.Port)
	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSrvConfig.Name))
	if err != nil {
		panic(err)
	}

	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn]连接【用户服务失败】")
		return
	}
	// 拨号连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接【用户服务失败】", "msg", err.Error())
	}
	// 生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = &userSrvClient
}
