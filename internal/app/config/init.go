package config

import "flag"

var ProxyConfig proxyConfig

var config = flag.String("config", "", "config")

func InitProxyConfig() {
	ProxyConfig.Load(*config)
}
