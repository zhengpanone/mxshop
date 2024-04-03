package utils

import (
	"fmt"
	"net"
	"os"
)

func GetIP() {
	// 获取本机的所有网络接口。
	interfaces, err := net.InterfaceByName("en0")
	if err != nil {
		fmt.Printf("Error getting interfaces : %v", err)
		os.Exit(1)
	}

	addrs, err := interfaces.Addrs()
	if err != nil {
		fmt.Printf("Error getting addresses : %v", err)
		os.Exit(1)
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

		// 仅选择IPv4地址，忽略环回地址，并且确保IP不是nil。
		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}

		fmt.Printf("Interface Name: %v, IP Address: %v\n", interfaces.Name, ip)
	}
}
