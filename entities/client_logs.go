package entities

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/spaolacci/murmur3"
)


func init() {
	for _, t := range EntityModels {
		eventModelsAsByte = append(eventModelsAsByte, []byte(t))
	} 
}
type SocketSubscriptionId struct {
	Conn *websocket.Conn
	Id string
}
type WsClientLog struct {
	counter map[*websocket.Conn]map[uint64]int
	Clients map[uint64][]*SocketSubscriptionId
	mutex  *sync.Mutex
}

func NewWsClientLog() WsClientLog {
	return WsClientLog{Clients:map[uint64][]*SocketSubscriptionId{}, mutex: &sync.Mutex{}, counter: make(map[*websocket.Conn]map[uint64]int) }
}
var eventModelsAsByte = [][]byte{}

func (c *WsClientLog) RegisterClient(subscription *ClientWsSubscription) {
	
	keys := []uint64{}

	for snet, val := range subscription.Filter {
		// keys = append(keys, key)
		for _, _type := range val {
			if _type == "*" {
				for _, t := range eventModelsAsByte {
					logger.Debugf("REGISTERING: %s", snet, string(t) )
					keys = append(keys, murmur3.Sum64(append(utils.UuidToBytes(snet), t...)))
				}
				
			} else {
				logger.Debugf("REGISTERING: %s", snet, string(_type) )
				keys = append(keys, murmur3.Sum64(append(utils.UuidToBytes(snet), []byte(_type)...)))
			}
		}

	}
	// process channel
	c.mutex.Lock()
	for _, key := range keys {
		// wsClients[key] = make(map[*websocket.Conn]*ClientWsSubscription)
		// wsClients[key][subscription.Conn] = subscription
		// for _, val2 := range subscription.Filter[key] {
		// 	wsClients[val2][subscription.Conn] = subscription
		// }
		if c.counter[subscription.Conn] == nil {
			c.counter[subscription.Conn] = make(map[uint64]int)
		}
		if c.Clients[key] == nil {
			c.Clients[key] = []*SocketSubscriptionId{}
		}
		c.counter[subscription.Conn][key] = len(c.Clients[key])
		logger.Debugf("Received PAYLOAD FILTEr 2: %v",subscription )
		c.Clients[key] = append(c.Clients[key], &SocketSubscriptionId{Conn: subscription.Conn, Id: subscription.Id})
		// wsClients[key][subscription.Conn] = subscription.Account
	}
	c.mutex.Unlock()
	logger.Debugf("Received PAYLOAD FILTEr: %v, %v", c.Clients)
}

func (c *WsClientLog) RemoveClient(conn *websocket.Conn) {
	c.mutex.Lock()
	for key, v := range c.counter[conn] {
		if c.Clients[key] == nil || len(c.Clients[key]) == 0 {
			continue
		}
		n := len(c.Clients[key])
		c.Clients[key][v] = c.Clients[key][n-1]
		c.Clients[key] = c.Clients[key][:n-1]
	}
	c.mutex.Unlock()
}

func (c *WsClientLog) GetClients(subnetId string, entittyModel string ) []*SocketSubscriptionId {
	
	hash := murmur3.Sum64(append(utils.UuidToBytes(subnetId), []byte(entittyModel)...))
	if c.Clients[hash] == nil {
		return []*SocketSubscriptionId{}
	}
	return c.Clients[hash]
}
