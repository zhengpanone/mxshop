package main

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"golang.org/x/exp/rand"
	"log"
	"time"
)

// 预热、冷启动
func main() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Printf("初始化sentinel异常%s", err)
	}
	// 配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test1", // 资源名
			TokenCalculateStrategy: flow.WarmUp,  // 冷启动你
			ControlBehavior:        flow.Reject,  // 直接拒绝
			WarmUpPeriodSec:        30,           // 30秒
			Threshold:              1000,
		},
	})
	if err != nil {
		log.Printf("加载规则失败,%s", err)
	}
	// 每秒统计一次，通过了多少，block了多少
	var globalTotal int
	var passTotal int
	var blockTotal int
	ch := make(chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			for {
				globalTotal++
				entry, blockError := sentinel.Entry("some-test1", sentinel.WithTrafficType(base.Inbound))
				if blockError != nil {
					blockTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					passTotal++
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					entry.Exit()
				}
			}
		}()
	}
	go func() {
		var oldTotal int
		var oldPass int
		var oldBlock int
		for {
			oneSecondTotal := globalTotal - oldTotal
			oldTotal = globalTotal

			oneSecondPass := passTotal - oldPass
			oldPass = passTotal

			oneSecondBlock := blockTotal - oldBlock
			oldBlock = blockTotal

			time.Sleep(time.Second)
			log.Printf("total:%d pass:%d block:%d", oneSecondTotal, oneSecondPass, oneSecondBlock)
		}
	}()
	<-ch
}
