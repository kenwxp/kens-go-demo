package util

import (
	"fmt"
	"net"
	"net/http"
)

//获取本机ip
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("get local ip failed")
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//获取本机Mac
func GetMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Get loacl Mac failed")
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		if mac.String() != "" {
			return mac.String()
		}

	}
	return ""
}

//获取请求中的ip
func GetRemoteIp(req *http.Request) string {
	var remoteAddr string
	// RemoteAddr
	remoteAddr = req.RemoteAddr
	if remoteAddr != "" {
		return remoteAddr
	}
	// ipv4
	remoteAddr = req.Header.Get("ipv4")
	if remoteAddr != "" {
		return remoteAddr
	}
	//
	remoteAddr = req.Header.Get("XForwardedFor")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Forwarded-For
	remoteAddr = req.Header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Real-Ip
	remoteAddr = req.Header.Get("X-Real-Ip")
	if remoteAddr != "" {
		return remoteAddr
	} else {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
