package ws

import (
	// "errors"
	"context"

	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger;
type Flag string

// Main websocket service that receives messages and directs them to the neccessry processing channel

// type msgError struct {
// 	code int
// 	message string
// }

type WsService struct {
	Ctx                    *context.Context
	Cfg                    *configs.MainConfiguration
	ClientHandshakeChannel *chan *entities.ClientHandshake
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewWsService(mainCtx *context.Context) *WsService {
	cfg, _ := (*mainCtx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	clientVerificationc, _ := (*mainCtx).Value(constants.ClientHandShackChId).(*chan *entities.ClientHandshake)
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
	logger.Debug("New ServeWebSocket c : ", c.RemoteAddr())

	if err != nil {
		logger.Debug("WS connection error:", err)
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
			logger.Println("read:", err)
			break

		} else {
			err = c.WriteMessage(mt, (append(message, []byte("recieved Signature")...)))
			if err != nil {
				logger.Println("Error:", err)
			} else {
				if(!isVerifed) {
					verifiedRequest, err := service.ConnectClient(message, constants.WS, c,)
					if (err != nil) {
						c.Close()
						continue
					}
					*p.ClientHandshakeChannel <- verifiedRequest
					
					logger.Debugf("message:", string(message))
					logger.Debugf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
					continue
				}
				// process message
			}

		}
	}

}
