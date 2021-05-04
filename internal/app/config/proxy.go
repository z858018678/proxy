package config

import (
	"fmt"
	"proxy/internal/app/lib/utils"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/douyu/jupiter/pkg/conf"
	file_datasource "github.com/douyu/jupiter/pkg/datasource/file"
	"github.com/douyu/jupiter/pkg/server/governor"
	"github.com/douyu/jupiter/pkg/trace/jaeger"
	"github.com/douyu/jupiter/pkg/xlog"
)

type proxyConfig struct {
	Logger  *xlog.Config `toml:"logger"`
	Service service      `toml:"service"`
}

type service struct {
	// TradeEngineRegistry   *TcpClientConfig `toml:"trade_engine_registry"`
	MessageEngineRegistry *TcpClientConfig `toml:"message_engine_registry"`
	Governor              *governor.Config `toml:"governor"`
	Jaeger                *jaeger.Config   `toml:"jaeger"`
	WsServer              *wsServerConfig  `toml:"ws_server"`
	// WsClient              *wsClientConfig  `toml:"ws_client"`
}

type wsServerConfig struct {
	// 服务地址
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type wsClientConfig struct {
	changed     chan struct{}
	ServerAddrs []string `toml:"server_addrs"`
}

func (c *wsClientConfig) Changed() <-chan struct{} {
	return c.changed
}

func (s *wsServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func defaultWsServerConfig() *wsServerConfig {
	var c wsServerConfig
	c.Host = utils.GetInternalIP()
	return &c
}

func defaultWsClientConfig() *wsClientConfig {
	var c wsClientConfig
	c.changed = make(chan struct{}, 1)
	return &c
}

// 读配置
func (c *proxyConfig) Load(path string) {
	c.Service.Governor = governor.DefaultConfig()
	// c.Service.WsClient = defaultWsClientConfig()
	c.Service.WsServer = defaultWsServerConfig()
	c.Service.Jaeger = jaeger.DefaultConfig()
	c.Service.MessageEngineRegistry = DefaultDiscoveryConfig()
	c.Logger = xlog.DefaultConfig()

	provider := file_datasource.NewDataSource(path, true)
	var cf = conf.New()
	if err := cf.LoadFromDataSource(provider, toml.Unmarshal); err != nil {
		xlog.JupiterLogger.Panic("load config",
			xlog.String("path", path),
			xlog.Any("load config", path),
			xlog.String("error", err.Error()))
	}

	var updateConfig = func(cf *conf.Configuration) {
		// c.Service.WsClient.ServerAddrs = nil

		// 解析 client 配置
		cf.UnmarshalKey("", c, conf.TagName("toml"))

		// c.Service.WsClient.changed <- struct{}{}

		// 热更新日志等级
		lvText := strings.ToLower(c.Logger.Level)
		if lvText != "" {
			xlog.JupiterLogger.Info("update level", xlog.String("level", lvText))
			xlog.JupiterLogger.UnmarshalText([]byte(lvText))
		}

		xlog.JupiterLogger.Debugf("logger config: %+v", c.Logger)
		xlog.JupiterLogger.Debugf("ws server config: %+v", c.Service.WsServer)
		xlog.JupiterLogger.Debugf("message engine registry config: %+v", c.Service.MessageEngineRegistry)
		xlog.JupiterLogger.Debugf("governor config: %+v", c.Service.Governor)
		xlog.JupiterLogger.Debugf("jaeger config: %+v", c.Service.Jaeger)
	}

	// 热更新
	cf.OnChange(updateConfig)

	// 更新配置
	updateConfig(cf)
}
