package service

import (
	"flag"
	"fmt"
	"log"
	"proxy/internal/app/config"
	"proxy/internal/app/engine"

	"github.com/douyu/jupiter"
)

func RunProxy() {
	flag.Parse()
	// 加载配置
	config.InitProxyConfig()

	// 创建服务实例
	var proxy = engine.NewProxyEngine()

	// 退出前打印信息
	proxy.RegisterHooks(jupiter.StageAfterStop, func() error {
		fmt.Println("EXIT Proxy ...")
		return nil
	})

	// 启动服务
	if err := proxy.Run(); err != nil {
		log.Fatal(err)
	}
}
