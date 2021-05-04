package config

const (
	ServiceName   = "trade_ws"
	Registration  = ""
	Discovery     = "TradeWsProxyDiscovery"
	ServiceScheme = "ws"
)

type TcpClientConfig struct {
	// 注册中心的地址
	Address string `toml:"address"`

	// 服务名字
	Name string `toml:"name"`
	// 服务发现时，会通过此 business 来发现服务
	Business string `toml:"business"`
	Scheme   string `toml:"scheme"`
}

func DefaultRegistryConfig() *TcpClientConfig {
	var r TcpClientConfig
	r.Name = Registration
	r.Business = ServiceName
	r.Scheme = ServiceScheme
	return &r
}

func DefaultDiscoveryConfig() *TcpClientConfig {
	var r TcpClientConfig
	r.Name = Discovery
	r.Business = ServiceName
	r.Scheme = ServiceScheme
	return &r
}
