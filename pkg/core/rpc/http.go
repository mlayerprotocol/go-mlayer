package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/rpc"
	"time"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	// "net/rpc/jsonrpc"
)

type JsonRpc struct {
	JsonRpcVersion string            `json:"jsonrpc"`
	Id             int               `json:"id"`
	Method         string            `json:"method"`
	Params         []json.RawMessage `json:"params"`
}

type HttpService struct {
	Ctx       *context.Context
	Cfg       *configs.MainConfiguration
	rpcClient *rpc.Client
}

func NewHttpService(mainCtx *context.Context) *HttpService {
	cfg, _ := (*mainCtx).Value(constants.ConfigKey).(*configs.MainConfiguration)
	return &HttpService{
		Ctx: mainCtx,
		Cfg: cfg,
	}
}

func (p *HttpService) sendHttp(w http.ResponseWriter, r *http.Request) {

	var jsonrpc JsonRpc
	err := json.NewDecoder(r.Body).Decode(&jsonrpc)

	payload := jsonrpc.Params
	// var params interface{}
	var reply RpcResponse

	err = p.rpcClient.Call("RpcService."+jsonrpc.Method, payload[0], &reply)

	if err != nil {
		reply = RpcResponse{
			Data:   err,
			Status: "failure",
		}
	}

	jData, err := json.Marshal(reply)
	if err != nil {
		logger.Errorf("marshal json error::", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

// func (p *HttpService) serveJSONRPC(w http.ResponseWriter, req *http.Request) {
// 	// if req.Method != "CONNECT" {
// 	// 	http.Error(w, "method must be connect", 405)
// 	// 	return
// 	// }
// 	conn, _, err := w.(http.Hijacker).Hijack()
// 	if err != nil {
// 		http.Error(w, "internal server error", 500)
// 		return
// 	}
// 	defer conn.Close()
// 	io.WriteString(conn, "HTTP/1.0 Connected\r\n\r\n")
// 	jsonrpc.ServeConn(conn)
// }

func (p *HttpService) Start(rpcPort string) error {
	_, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	hostname := "localhost"
	port := fmt.Sprintf(":%s", rpcPort)
	client, err := rpc.DialHTTP("tcp", hostname+port)

	if err != nil {
		logger.Errorf("Rpc Error::", err)
		return err
	}
	
	defer client.Close()
	p.rpcClient = client
	http.HandleFunc("/", p.sendHttp)
	// http.HandleFunc("/rpcendpoint", p.serveJSONRPC)
	err = http.ListenAndServe(":"+p.Cfg.RPCHttpPort, nil)
	return err
}
