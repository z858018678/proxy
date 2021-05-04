package engine

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"proxy/internal/app/config"
	"proxy/internal/app/job"
	"proxy/internal/app/lib/utils"
	"strings"
	"time"

	pb "proxy/pb/aex_protocol/aex_registry_center/golang"

	"github.com/douyu/jupiter"
	xtcp "github.com/douyu/jupiter/pkg/client/tcp"
	xwsclient "github.com/douyu/jupiter/pkg/client/websocket"
	xroom "github.com/douyu/jupiter/pkg/server/websocket/room"
	xserver "github.com/douyu/jupiter/pkg/server/websocket/server"
	xsocket "github.com/douyu/jupiter/pkg/server/websocket/socket"
	"github.com/douyu/jupiter/pkg/util/xcompressor"
	"github.com/douyu/jupiter/pkg/util/xgo"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/gogo/protobuf/proto"
)

type ProxyEngine struct {
	*jupiter.Application
}

func NewProxyEngine() *ProxyEngine {
	var app = jupiter.DefaultApp()

	eng := &ProxyEngine{Application: app}

	eng.WithOptions(
		// 禁用默认 flag
		jupiter.WithDisable(jupiter.DisableParserFlag),
		// 禁用加载配置
		jupiter.WithDisable(jupiter.DisableLoadConfig),
		// 禁用默认的服务治理
		jupiter.WithDisable(jupiter.DisableDefaultGovernor),
	)

	// 服务启动项
	if err := eng.Startup(
		xgo.ParallelWithError(
			eng.initGovernor,
			// eng.runTradeEngineDiscovery,
			// eng.runTradeEngineWsClient,
			eng.runMessageEngineDiscovery,
			eng.runMessageEngineWsClient,
			eng.runWsServer,
		),
	); err != nil {
		xlog.Panic("startup engine", xlog.Any("err", err))
	}
	return eng
}

// 服务治理
func (eng *ProxyEngine) initGovernor() error {
	eng.Serve(config.ProxyConfig.Service.Governor.Build())
	return nil
}

var compressor = xcompressor.NewZlibCompressor(9)

/*
// 连接交易引擎服务
// 服务发现消息 channel
// 发现的服务会通过此 channel 传过来
var tradeEngineRegistryChan = make(chan []string, 1)

func handleTradeEngineRegistryMsg(msg []byte) []byte {
	return proxyhandleRecvMsg(msg, tradeEngineRegistryChan)
}

// 发现交易引擎服务
func (e *ProxyEngine) runTradeEngineDiscovery() error {
	var addr, err = net.ResolveTCPAddr("tcp", config.ProxyConfig.Service.TradeEngineRegistry.Address)
	if err != nil {
		xlog.JupiterLogger.Panic("resolve trade engine tcp address", xlog.FieldErr(err))
		return err
	}

	// ping
	var b pb.RegisterFoundREQMessage
	b.Cmd = "heartbeat"
	var hb pb.HeartBeat
	hb.Data = "ping"
	b.HBDData = &hb
	var pingData, _ = proto.Marshal(&b)

	// 启动服务发现 tcp 客户端
	var client = xtcp.NewClient(
		addr.String(),
		xtcp.WithMsgDelim([]byte("\r\n\r\n")),
		xtcp.WithMsgHandler(handleTradeEngineRegistryMsg),
		xtcp.WithPing(pingData, time.Second*10),
		xtcp.WithReconnect(true),
		xtcp.WithDialTimeout(time.Second*5),
	)

	// 发现服务消息
	var r pb.RegisterFoundREQMessage
	r.Cmd = "worker_register"
	var wr pb.BusinessWorkerRegisterData
	wr.Name = config.ProxyConfig.Service.TradeEngineRegistry.Name
	wr.Business = config.ProxyConfig.Service.TradeEngineRegistry.Business
	r.BWRData = &wr

	client.WithOnConnect(func() {
		var data, _ = proto.Marshal(&r)
		client.Send(data)
	})

	e.Schedule(client)
	return nil
}

// 连接交易引擎 websocket 服务
func (e *ProxyEngine) runTradeEngineWsClient() error {
	var clientsRoom = job.NewWsClientsJob(context.Background())
	e.Schedule(clientsRoom)

	go func() {
		for range config.ProxyConfig.Service.WsClient.Changed() {
			// 当前所有服务地址
			var totalServers []string
			// 初始化为服务发现的所有服务地址
			totalServers = config.ProxyConfig.Service.WsClient.ServerAddrs

			// 从已连接的服务中去校对
			// 将没有出现在发现列表中的服务关闭
			clientsRoom.RangeClients(func(k, v interface{}) bool {
				var addr, ok = k.(string)
				if !ok {
					clientsRoom.RemoveClient(addr)
					return true
				}

				if utils.InStringSlice(addr, totalServers) {
					return true
				}

				v.(*xwsclient.WsClient).Stop()
				return true
			})

			// 从已发现的服务中去校对
			// 连接还没有连接上的服务
			for _, server := range totalServers {
				var _, connected = clientsRoom.LoadClient(server)
				if !connected {
					// 连接服务，创建客户端
					var client = xwsclient.NewClient(
						server,
						xwsclient.WithContext(clientsRoom.GetContext()),
						xwsclient.WithDialTimeout(time.Second),
						xwsclient.WithPing([]byte("ping"), time.Second*15),
						xwsclient.WithReconnect(true),
						// xwsclient.WithCompressor(compressor, compressor),
						xwsclient.WithSendChanBuffer(5),
						xwsclient.WithRecvMsgMaxBytes(2000),
						xwsclient.WithOnRecvMsg(func(data []byte) []byte {
							msgForwardChan <- data
							return nil
						}),
					)

					clientsRoom.AddClient(server, client)
					client.Run()
				}
			}
		}
	}()

	return nil
}
*/

// 连接消息引擎服务
// 服务发现消息 channel
// 发现的服务会通过此 channel 传过来
var messageEngineRegistryChan = make(chan []string, 1)

func handleMessageEngineRegistryMsg(msg []byte) []byte {
	return proxyhandleRecvMsg(msg, messageEngineRegistryChan)
}

// 运行消息引擎服务发现
func (e *ProxyEngine) runMessageEngineDiscovery() error {
	var addr, err = net.ResolveTCPAddr("tcp", config.ProxyConfig.Service.MessageEngineRegistry.Address)
	if err != nil {
		xlog.JupiterLogger.Panic("resolve tcp address", xlog.FieldErr(err))
		return err
	}

	// ping
	var b pb.RegisterFoundREQMessage
	b.Cmd = "heartbeat"
	var hb pb.HeartBeat
	hb.Data = "ping"
	b.HBDData = &hb
	var pingData, _ = proto.Marshal(&b)

	// 启动服务发现 tcp 客户端
	var client = xtcp.NewClient(
		addr.String(),
		xtcp.WithMsgDelim([]byte("\r\n\r\n")),
		xtcp.WithMsgHandler(handleMessageEngineRegistryMsg),
		xtcp.WithPing(pingData, time.Second*10),
		xtcp.WithReconnect(true),
		xtcp.WithDialTimeout(time.Second*5),
	)

	// 发现服务消息
	var r pb.RegisterFoundREQMessage
	r.Cmd = "worker_register"
	var wr pb.BusinessWorkerRegisterData
	wr.Name = config.ProxyConfig.Service.MessageEngineRegistry.Name
	wr.Business = config.ProxyConfig.Service.MessageEngineRegistry.Business
	r.BWRData = &wr

	client.WithOnConnect(func() {
		var data, _ = proto.Marshal(&r)
		client.Write(data)
	})

	e.Schedule(client)
	return nil
}

// 转发消息 channel
// 任何接受到的消息都转发到此 channel 中
// 然后从此 channel 中取出消息广播出去
var msgForwardChan = make(chan []byte, 100)

// 启动 ws 客户端
func (e *ProxyEngine) runMessageEngineWsClient() error {
	// 启动一个客户端广播室
	// 用来给客户端广播消息
	var clientsRoom = job.NewWsClientsJob(context.Background())
	e.Schedule(clientsRoom)

	// 取出消息，广播给客户端
	go func() {
		for msg := range msgForwardChan {
			clientsRoom.Broadcast(msg)
		}
	}()

	// 与服务发现数据同步 ws 客户端连接
	go func() {
		for servers := range messageEngineRegistryChan {
			// 当前所有服务地址
			var totalServers []string
			// 初始化为服务发现的所有服务地址
			totalServers = parseAddress(servers...)

			// 从已连接的服务中去校对
			// 将没有出现在发现列表中的服务关闭
			clientsRoom.RangeClients(func(k, v interface{}) bool {
				var addr, ok = k.(string)
				if !ok {
					clientsRoom.RemoveClient(addr)
					return true
				}

				if utils.InStringSlice(addr, totalServers) {
					return true
				}

				v.(*xwsclient.WsClient).Stop()
				return true
			})

			// 从已发现的服务中去校对
			// 连接还没有连接上的服务
			for _, server := range totalServers {
				var _, connected = clientsRoom.LoadClient(server)
				if !connected {
					// 连接服务，创建客户端
					var client = xwsclient.NewClient(
						server,
						xwsclient.WithContext(clientsRoom.GetContext()),
						xwsclient.WithDialTimeout(time.Second),
						xwsclient.WithPing([]byte("ping"), time.Second*15),
						xwsclient.WithReconnect(true),
						// xwsclient.WithCompressor(compressor, compressor),
						xwsclient.WithSendChanBuffer(5),
						xwsclient.WithRecvMsgMaxBytes(2000),
					)

					// 客户端维护在广播室中
					clientsRoom.AddClient(server, client)
					client.Run()
				}
			}
		}
	}()

	return nil
}

// 启动 ws 服务
func (e *ProxyEngine) runWsServer() error {
	var server *xserver.WsServer

	server = xserver.NewServer(
		config.ProxyConfig.Service.WsServer.Addr(),
		xserver.WithLogger(xlog.JupiterLogger),
		xserver.WithContext(context.Background()),
	)

	// 注册服务
	server.RegisterPath(
		"/wsv3",
		xroom.WithSocketOptions(
			// xsocket.WithCompressor(compressor, compressor),
			xsocket.WithOnRecvMsg(func(data []byte) []byte {
				msgForwardChan <- data
				return nil
			}),
			// 对 socket 的心跳检测
			xsocket.WithHeartbeat([]byte("ping"), time.Second*25),
		),
	)

	e.Schedule(server)
	return nil
}

// 解析地址
func parseAddress(addrs ...string) []string {
	var res []string
	for _, addr := range addrs {
		var sli = strings.Split(addr, "|")
		if len(sli) != 4 {
			continue
		}

		var scheme = sli[0]
		var host = sli[1]
		var port = sli[2]
		var path = sli[3]

		var u = url.URL{
			Scheme: scheme,
			Host:   fmt.Sprintf("%s:%s", host, port),
			Path:   path,
		}
		res = append(res, u.String())
	}

	return res
}

// 处理服务发现消息
func proxyhandleRecvMsg(msg []byte, updateRegistryChan chan []string) []byte {
	if len(msg) == 0 {
		return nil
	}

	var logger = xlog.JupiterLogger.With(
		xlog.FieldMod("registry"),
	)

	var response pb.RegisterFoundRESMessage
	var err = proto.Unmarshal(msg, &response)
	if err != nil {
		logger.Error("unmarshal msg failed", xlog.FieldErr(err), xlog.FieldExtMessage(msg))
		return nil
	}

	switch response.GetCmd() {
	case "broadcast_addresses":
		logger.Info("broadcast addresses", xlog.FieldValue(response.GetMsg()))
		updateRegistryChan <- response.GetBARData().GetAddressList()

	case "register":
		// 注册失败
		switch response.GetCode() {
		case 1:
			logger.Info("register", xlog.FieldValue(response.GetMsg()))

		default:
			logger.Panic("register", xlog.Any("response", response))
		}

	case "connect":
		logger.Info("connection", xlog.FieldValue(response.GetMsg()))

	case "heartbeat":
		logger.Debug("pong", xlog.FieldValue(response.GetMsg()))

	default:
		logger.Warn("unknown msg", xlog.FieldExtMessage(msg))
	}

	return nil
}
