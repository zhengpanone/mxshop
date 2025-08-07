package utils

import (
	"fmt"
	"github.com/zhengpanone/mxshop/mxshop-api/common/global"
	"net"
	"os"
)

func GetIP() string {
	// 获取本机的所有网络接口。
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting interfaces : %v", err)
		os.Exit(1)
	}
	for _, iface := range interfaces {
		// 忽略未启用的接口或环回接口
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口地址
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Error getting addresses : %v", err)
			continue
		}

		for _, addr := range addrs {
			// 获取接口上的所有地址信息。
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 检查是否为IPv4，并排除环回地址
			if ip != nil && ip.To4() != nil && !ip.IsLoopback() {
				global.Logger.Info("IPV4 address: " + ip.String())
				return ip.String()
			}
		}

	}
	global.Logger.Warn("No valid IPv4 address found.")
	os.Exit(1)
	return ""
}
