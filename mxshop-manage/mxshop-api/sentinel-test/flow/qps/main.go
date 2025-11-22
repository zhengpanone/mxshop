package main

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
	"time"
)

func main() {
	// 初始化sentinel
	err := sentinel.InitDefault()
	if err != nil {
		log.Printf("初始化sentinel异常%s", err)
	}
	// 配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test1", // 资源名
			TokenCalculateStrategy: flow.Direct,  // 当前流量控制器的Token计算策略。Direct表示直接使用字段 Threshold 作为阈值；WarmUp表示使用预热方式计算Token的阈值。
			ControlBehavior:        flow.Reject,  // 直接拒绝
			Threshold:              10,
			StatIntervalInMs:       1000,
		},
		{
			Resource:               "some-test2",    // 资源名
			TokenCalculateStrategy: flow.Direct,     // 当前流量控制器的Token计算策略。Direct表示直接使用字段 Threshold 作为阈值；WarmUp表示使用预热方式计算Token的阈值。
			ControlBehavior:        flow.Throttling, // 匀速通过
			Threshold:              10,
			StatIntervalInMs:       1000,
		},
	})
	if err != nil {
		log.Printf("加载规则失败,%s", err)
	}
	for i := 0; i < 12; i++ {
		entry, blockError := sentinel.Entry("some-test2", sentinel.WithTrafficType(base.Inbound)) // base.Inbound 入口流量、 Outbound 出口流量
		if blockError != nil {
			log.Printf("%d限流了%s", i, blockError)
		} else {
			log.Printf("检查通过%d", i)
			entry.Exit()
		}
		time.Sleep(101 * time.Millisecond)
	}

}
