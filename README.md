# Proxy(交易消息转发服务)

当前架构：
```
Users(Web/App)
┏━┓┏━┓┏━┓┏━┓┏━┓
┃u┃┃u┃┃u┃┃u┃┃u┃
┗━┛┗━┛┗━┛┗━┛┗━┛ Ws Client (客户端，连接并订阅消息服务，接受 market 数据)


Websocket 消息服务
┏━━━┓┏━━━┓┏━━━┓ Ws Server (针对用户的订阅发送 market 数据)
┃ S ┃┃ S ┃┃ S ┃
┗━━━┛┗━━━┛┗━━━┛ Ws Server + Regsitry (接收撮合服务消息)


撮合数据发送服务 + Proxy 消息转发
┏━━━━━━┓┏━━━━━━┓
┃┏━━━━┓┃┃┏━━━━┓┃ Ws Client + Discovery (Proxy 向每一个 TradeWsServer 发送撮合消息)
┃┃ P  ┃┃┃┃ P  ┃┃  
┃┗━━━━┛┃┃┗━━━━┛┃ Ws Server (Proxy 从每一个撮合引擎接受撮合消息)
┃      ┃┃      ┃ 
┃┏┓┏┓┏┓┃┃┏┓┏┓┏┓┃ 
┃┗┛┗┛┗┛┃┃┗┛┗┛┗┛┃ Ws Client (撮合殷勤连接 Proxy，向其发送自己的撮合数据)
┗━━━━━━┛┗━━━━━━┛
```

编译方法:

```bash

// 第一次拉取项目时，应该下面的命令应该先跑一下
git submodule update --init

// 更新 submodule
git submodule foreach git pull origin master
git submodule update

cd AEX_SERVICE/trade_engine_service/proxy
make install

$GOBIN/tradeProxy -config AEX_SERVICE/trade_engine_service/proxy/config/config.toml
```
编译之后，二进制文件在 $GOBIN 中，文件名为: tradeProxy
