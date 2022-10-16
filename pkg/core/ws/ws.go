package ws

import (
	// "errors"
	"context"
	"log"
	"net/http"
	"time"

	services "github.com/ByteGum/go-icms/pkg/service"
	utils "github.com/ByteGum/go-icms/utils"
	"github.com/gorilla/websocket"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type WsService struct {
	Ctx            *context.Context
	Cfg            *utils.Configuration
	MessageService *services.MessageService
	ChannelService *services.ChannelService
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewWsService(mainCtx *context.Context) *WsService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	return &WsService{
		Ctx:            mainCtx,
		Cfg:            cfg,
		MessageService: services.NewMessageService(mainCtx),
		ChannelService: services.NewChannelService(mainCtx),
	}
}

func newResponse(status string, data interface{}) *RpcResponse {
	d := RpcResponse{
		Status: status,
		Data:   data,
	}
	return &d
}

func (p *WsService) SendMessage(request utils.MessageJsonInput, reply *RpcResponse) error {
	utils.Logger.Info("SendMessage request:::", request)
	chatMsg := utils.CreateMessageFromJson(request)
	c, err := (*p.MessageService).Send(chatMsg, request.Signature)
	if err != nil {
		return err
	}
	reply = newResponse("success", c)
	return nil
}

func (p *WsService) Subscription(request utils.Subscription, reply *RpcResponse) error {
	utils.Logger.Info("Subscription request:::", request)
	err := (*p.ChannelService).ChannelSubscription(&request)
	if err != nil {
		return err
	}
	reply = newResponse("success", request)
	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ServeWebSocket(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	log.Print("New ServeWebSocket c : ", c.RemoteAddr())

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	hasVerifed := false
	time.AfterFunc(5000*time.Millisecond, func() {
		if !hasVerifed {
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
				// signature := string(message)
				verifiedRequest, _ := utils.VerificationRequestFromBytes(message)
				log.Println("verifiedRequest.Message: ", verifiedRequest.Message)

				if utils.VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
					// verifiedConn = append(verifiedConn, c)
					hasVerifed = true
					log.Println("Verification was successful: ", verifiedRequest)
				}
				log.Println("message:", string(message))
				log.Printf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
			}

		}
	}

}
