package utils

import (
	"net"
	"strings"
)

// GetLocalIPv4 : get local ipv4
func GetLocalIPv4() (ip string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, face := range interfaces {
		if strings.Contains(face.Name, "lo") {
			continue
		}
		addrs, err := face.Addrs()
		if err != nil {

		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
					if !strings.Contains(ip, ":") && ip != "127.0.0.1" {
						return ip
					}
				}
			}
		}
	}
	return ""
}

// GetLocalRandomPort : get local random available port
func GetLocalRandomPort() (port int) {
	l, _ := net.Listen("tcp", ":0") //listen on localhost
	defer l.Close()
	port = l.Addr().(*net.TCPAddr).Port
	return port
}
