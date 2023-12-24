package ws

import (
	// "errors"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	utils "github.com/mlayerprotocol/go-mlayer/utils"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type WsService struct {
	Ctx                    *context.Context
	Cfg                    *utils.Configuration
	ClientHandshakeChannel *chan *utils.ClientHandshake
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewWsService(mainCtx *context.Context) *WsService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	clientVerificationc, _ := (*mainCtx).Value(utils.ClientHandShackChId).(*chan *utils.ClientHandshake)
	return &WsService{
		Ctx:                    mainCtx,
		Cfg:                    cfg,
		ClientHandshakeChannel: clientVerificationc,
	}
}

func newResponse(status string, data interface{}) *RpcResponse {
	d := RpcResponse{
		Status: status,
		Data:   data,
	}
	return &d
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (p *WsService) ServeWebSocket(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	utils.Logger.Info("New ServeWebSocket c : ", c.RemoteAddr())

	if err != nil {
		utils.Logger.Debug("WS connection error:", err)
		return
	}
	defer c.Close()
	isVerifed := false
	time.AfterFunc(5000*time.Millisecond, func() {
		if !isVerifed {
			c.Close()
		}
	})
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break

		} else {
			err = c.WriteMessage(mt, (append(message, []byte("recieved Signature")...)))
			if err != nil {
				log.Println("Error:", err)
			} else {
				if(!isVerifed) {
					isVerifed = utils.ConnectClient(message, utils.WS, c, p.ClientHandshakeChannel)
					if(!isVerifed) {
						c.Close();
					}
					log.Println("message:", string(message))
					log.Printf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
					continue
				}
				// process message
			}

		}
	}

}
