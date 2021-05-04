package utils

import (
	"net"

	"github.com/douyu/jupiter/pkg/xlog"
)

// 获取本机内网IP
func GetInternalIP() string {
	var addrs, err = net.InterfaceAddrs()
	if err != nil {
		xlog.JupiterLogger.Error("get internal ip", xlog.FieldErr(err))
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

// 判断字符串是否在数组内
func InStringSlice(str string, sli []string) bool {
	for _, s := range sli {
		if str == s {
			return true
		}
	}

	return false
}
