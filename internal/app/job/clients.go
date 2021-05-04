package job

import (
	"context"
	"sync"
	"time"

	xclient "github.com/douyu/jupiter/pkg/client/websocket"
)

// ws 客户端广播室
type WsClientsJob struct {
	ctx           context.Context
	cancel        context.CancelFunc
	clients       sync.Map
	broadcastChan chan []byte
}

// 检查 客户端 ID
func (j *WsClientsJob) CheckID(id string) bool {
	_, ok := j.clients.Load(id)
	return ok
}

func (j *WsClientsJob) RangeClients(f func(k, v interface{}) bool) {
	j.clients.Range(f)
}

func (j *WsClientsJob) LoadClient(id string) (*xclient.WsClient, bool) {
	var client, ok = j.clients.Load(id)
	if !ok {
		return nil, false
	}

	return client.(*xclient.WsClient), true
}

// 添加一个客户端到广播室
func (j *WsClientsJob) AddClient(id string, client *xclient.WsClient) {
	j.clients.Store(id, client)
	// 客户端断开连接或者广播室关闭时，从广播室中删除此客户端
	go func(id string) {
		select {
		case <-client.Stopped():
		case <-j.ctx.Done():
		}
		j.clients.Delete(id)
	}(id)
}

// 移除一个客户端
func (j *WsClientsJob) RemoveClient(id string) {
	var c, _ = j.clients.LoadAndDelete(id)
	var client, ok = c.(*xclient.WsClient)
	if ok && client != nil {
		client.Stop()
	}
}

// 广播消息方法
func (j *WsClientsJob) Broadcast(msg []byte) {
	j.broadcastChan <- msg
}

// 广播消息处理方式
func (j *WsClientsJob) broadcastMonitor() {
	for msg := range j.broadcastChan {
		j.RangeClients(func(_, v interface{}) bool {
			go func(v interface{}) {
				var ctx, _ = context.WithTimeout(j.ctx, time.Second)
				v.(*xclient.WsClient).Write(ctx, msg)
			}(v)
			return true
		})
	}
}

func (j *WsClientsJob) Run() error {
	go j.broadcastMonitor()
	return nil
}

func (j *WsClientsJob) Stop() error {
	j.cancel()
	return nil
}

func (j *WsClientsJob) GetContext() context.Context {
	return j.ctx
}

// 创建一个 client 广播室
func NewWsClientsJob(ctx context.Context) *WsClientsJob {
	var j WsClientsJob
	j.ctx, j.cancel = context.WithCancel(ctx)
	j.broadcastChan = make(chan []byte, 100)
	return &j
}
