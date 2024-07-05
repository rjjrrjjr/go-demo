package ips

import (
	"fmt"
	"net"
	"testing"
)

func TestGetIp(t *testing.T) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("获取IP地址失败:", err)
		return
	}

	for _, address := range addrs {
		// 检查地址类型，排除IPv6地址和回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("当前服务器的IP地址:", ipnet.IP.String())
			}
		}
	}
}

func TestGetIp2(t *testing.T) {
	fmt.Println(getLocalIP())
}

func getLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addresses {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""
}
