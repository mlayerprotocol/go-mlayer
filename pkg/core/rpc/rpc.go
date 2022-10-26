package rpc

import (
	// "errors"
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	services "github.com/ByteGum/go-icms/pkg/service"
	utils "github.com/ByteGum/go-icms/utils"
	shell "github.com/ipfs/go-ipfs-api"
)

type Flag string

type Ipfs struct {
	Message string
	Subject string
}

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type RpcService struct {
	Ctx            *context.Context
	Cfg            *utils.Configuration
	MessageService *services.MessageService
	ChannelService *services.ChannelService
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRpcService(mainCtx *context.Context) *RpcService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	return &RpcService{
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

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}

func (p *RpcService) SendMessage(request utils.MessageJsonInput, reply *RpcResponse) error {
	var messageCID string
	utils.Logger.Infof("ipfs address %s", p.Cfg.Ipfs.Host)
	client := NewClient(p.Cfg.Ipfs.ProjectId, p.Cfg.Ipfs.ProjectSecret)
	sh := shell.NewShellWithClient(p.Cfg.Ipfs.Host, client)
	if len(request.Subject) > 0 && len(request.Message) > 0 {
		ipfs := &Ipfs{
			Message: request.Message,
			Subject: request.Subject,
		}
		tsdBin, _ := json.Marshal(ipfs)
		reader := bytes.NewReader(tsdBin)

		cid, err := sh.Add(reader)
		if err != nil {
			utils.Logger.Errorf("ipfs error:: %w", err)
		}
		messageCID = cid
		utils.Logger.Infof("IPFS messageCID::: %s", messageCID)
	}
	// if message and subject specified
	//create a json object with subject and message fields
	// push to ipfs
	// add cid to client message
	chatMsg, err := utils.CreateMessageFromJson(request)
	if err != nil {
		return err
	}
	c, err := (*p.MessageService).Send(chatMsg, request.Signature, messageCID)
	if err != nil {
		return err
	}
	reply = newResponse("success", c)
	return nil
}

func (p *RpcService) Subscription(request utils.Subscription, reply *RpcResponse) error {
	utils.Logger.Debug("Subscription request:::", request)
	err := (*p.ChannelService).ChannelSubscription(&request)
	if err != nil {
		return err
	}
	reply = newResponse("success", request)
	return nil
}
