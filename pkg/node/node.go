/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package node

import (
	"context"
	"encoding/hex"
	"math/big"
	"time"

	"strings"
	"sync"

	"net"
	"net/http"
	"net/rpc"

	// "net/rpc/jsonrpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"

	// "github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/common/apperror"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/service"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/chain/evm/abis/stake"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	p2p "github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/rest"
	rpcServer "github.com/mlayerprotocol/go-mlayer/pkg/core/rpc"
	ws "github.com/mlayerprotocol/go-mlayer/pkg/core/ws"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var Nodes = []*websocket.Conn{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
} // use default options

var logger = &log.Logger

func Start(mainCtx *context.Context) {
	ctx, cancel := context.WithCancel(*mainCtx)
	defer cancel()

	cfg, ok := ctx.Value(constants.ConfigKey).(*configs.MainConfiguration)
	if !ok {
		panic(apperror.Internal("Unable to load main config"))
	}
	connectedSubscribers := make(map[string]map[string][]interface{})

	// incomingEventsC := make(chan types.Log)

	var wg sync.WaitGroup
	// errc := make(chan error)

	// deliveryProofBlockStateStore := db.New(&ctx, constants.DeliveryProofBlockStateStore)
	// subscriptionBlockStateStore := db.New(&ctx, constants.SubscriptionBlockStateStore)

	// stores messages that have been validated
	// validMessagesStore := db.New(&ctx, constants.ValidMessageStore)
	// unProcessedClientPayloadStore := db.New(&ctx, fmt.Sprintf("%s", constants.UnprocessedClientPayloadStore))
	// unsentMessageP2pStore := db.New(&ctx, constants.UnsentMessageStore)
	// topicSubscriptionStore := db.New(&ctx, constants.TopicSubscriptionStore)
	// newTopicSubscriptionStore := db.New(&ctx, constants.NewTopicSubscriptionStore)
	// topicSubscriptionCountStore := db.New(&ctx, constants.TopicSubscriptionCountStore)
	// sentMessageStore := db.SqlDb(&ctx, constants.SentMessageStore)
	// deliveryProofStore := db.New(&ctx, constants.DeliveryProofStore)
	// localDPBlockStore := db.New(&ctx, constants.DeliveryProofBlockStore)
	// unconfirmedBlockStore := db.New(&ctx, constants.UnconfirmedDeliveryProofStore)
	eventCountStore := db.New(&ctx,   string(constants.EventCountStore))
	defer eventCountStore.Close()
	ctx = context.WithValue(ctx, constants.EventCountStore, eventCountStore)

	claimedRewardStore := db.New(&ctx,   string(constants.ClaimedRewardStore))
	defer claimedRewardStore.Close()
	ctx = context.WithValue(ctx, constants.ClaimedRewardStore, claimedRewardStore)
	// ctx = context.WithValue(ctx, constants.NewTopicSubscriptionStore, newTopicSubscriptionStore)
	// ctx = context.WithValue(ctx, constants.UnprocessedClientPayloadStore, unProcessedClientPayloadStore)
	ctx = context.WithValue(ctx, constants.ConnectedSubscribersMap, connectedSubscribers)

	// p2pDataStore := db.New(&ctx,   string(constants.P2PDataStore))
	// defer p2pDataStore.Close()
	// ctx = context.WithValue(ctx, constants.P2PDataStore, p2pDataStore)

	defer wg.Wait()

	//  wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	client.ListenForNewAuthEventFromPubSub(&ctx)
	// }()
	//  wg.Add(1)
	//  go func() {
	// 	defer wg.Done()
	// 	// go client.ListenForNewTopicEventFromPubSub(&ctx)
	//  }()

	wg.Add(1)
	go func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		defer wg.Done()
		for {
			select {
			case inEvent, ok := <-channelpool.IncomingMessageEvent_P2P_D_C:
				if !ok {
					logger.Errorf("Incoming Message channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				// VALIDATE, STORE AND DISTRIBUTE
				go func() {
					inMessage := inEvent.Payload.Data.(entities.Message)
					logger.Infof("Received new message %s\n", inMessage.DataHash)
					// validMessagesStore.Set(ctx, db.Key(inMessage.Key()), inMessage.MsgPack(), false)
					_reciever := inMessage.Receiver
					_recievers := strings.Split(string(_reciever), ":")
					_currentTopic := connectedSubscribers[_recievers[1]]
					logger.Info("connectedSubscribers : ", connectedSubscribers, "---", _reciever)
					logger.Info("_currentTopic : ", _currentTopic, "/n")
					for _, signerConn := range _currentTopic {
						for i := 0; i < len(signerConn); i++ {
							signerConn[i].(*websocket.Conn).WriteMessage(1, inMessage.MsgPack())
						}
					}
				}()

			}

		}
	}()

	// wg.Add(1)
	// go func() {
	// 	_, cancel := context.WithTimeout(context.Background(), time.Second)
	// 	defer cancel()
	// 	defer wg.Done()
	// 	subscription.ProcessNewSubscription(
	// 		ctx,
	// 		subscriptionBlockStateStore,
	// 		topicSubscriptionCountStore,
	// 		newTopicSubscriptionStore,
	// 		topicSubscriptionStore,
	// 		&wg)
	// }()

	// wg.Add(1)
	// go func() {
	// 	_, cancel := context.WithTimeout(context.Background(), time.Second)
	// 	defer cancel()
	// 	defer wg.Done()
	// 	message.ProcessNewMessageEvent(ctx, unsentMessageP2pStore, &wg)
	// }()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// defer func() {
		// if err := recover(); err != nil {
		// 	wg.Done()
		// 	errc <- fmt.Errorf("P2P error: %g", err)
		// }
		// }()
		
		go p2p.ProcessEventsReceivedFromOtherNodes(entities.AuthModel, &entities.AuthorizationPubSub, &ctx, service.HandleNewPubSubAuthEvent)
		go p2p.ProcessEventsReceivedFromOtherNodes(entities.TopicModel, &entities.TopicPubSub, &ctx, service.HandleNewPubSubTopicEvent)
		go p2p.ProcessEventsReceivedFromOtherNodes(entities.SubnetModel, &entities.SubnetPubSub, &ctx, service.HandleNewPubSubSubnetEvent)
		go p2p.ProcessEventsReceivedFromOtherNodes(entities.WalletModel, &entities.WalletPubSub, &ctx, service.HandleNewPubSubWalletEvent)
		go p2p.ProcessEventsReceivedFromOtherNodes(entities.SubnetModel, &entities.SubscriptionPubSub, &ctx, service.HandleNewPubSubSubscriptionEvent)
		go p2p.ProcessEventsReceivedFromOtherNodes(entities.MessageModel, &entities.MessagePubSub, &ctx, service.HandleNewPubSubMessageEvent)
		

		p2p.Run(&ctx)
		// if err != nil {
		// 	wg.Done()
		// 	panic(err)
		// }
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof("RPCCONFIG", cfg.EvmRpcConfig["31337"])
		GenerateRegsitrationData(cfg)
		for {
			time.Sleep(5 * time.Second)
			logger.Info("Generating batch...")
			batch, err := generateBatch(0, 0, &ctx)
			if batch == nil {
				
				break
			}
			if err != nil {
				
			} else {
				
				processSentryRewardBatch(ctx, cfg, batch)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	

	// wg.Add(1)
	// go func() {
	// 	_, cancel := context.WithTimeout(context.Background(), time.Second)
	// 	defer cancel()
	// 	defer wg.Done()
	// 	_, client, contractAddress, err := evm.StakeContract(cfg.EVMRPCWss, cfg.StakeContract)
	// 	if err != nil {
	// 		logger.Fatal(err, cfg.EVMRPCWss, cfg.StakeContract)
	// 	}
	// 	query := ethereum.FilterQuery{
	// 		// FromBlock: big.NewInt(23506010),
	// 		// ToBlock:   big.NewInt(23506110),

	// 		Addresses: []common.Address{contractAddress},
	// 	}

	// 	// logs, err := client.FilterLogs(context.Background(), query)
	// 	// if err != nil {
	// 	// 	logger.Fatal(err)
	// 	// }
	// 	// parserEvent(logs[0], "StakeEvent")

	// 	// logger.Infof("Past Events", logs)
	// 	// incomingEventsC

	// 	sub, err := client.SubscribeFilterLogs(context.Background(), query, incomingEventsC)
	// 	if err != nil {
	// 		logger.Fatal(err, "SubscribeFilterLogs")
	// 	}

	// 	for {
	// 		select {
	// 		case err := <-sub.Err():
	// 			logger.Fatal(err)
	// 		case vLog := <-incomingEventsC:
	// 			fmt.Println(vLog) // pointer to event log
	// 			parserEvent(vLog, "StakeEvent")
	// 		}
	// 	}

	// }()

	wg.Add(1)
	go func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		defer wg.Done()
		rpc.Register(rpcServer.NewRpcService(&ctx))
		rpc.HandleHTTP()
		host :=  cfg.RPCHost
		if host == "" {
			host = "127.0.0.1"
		}
		listener, err := net.Listen("tcp", host+":"+cfg.RPCPort)
		if err != nil {
			logger.Fatal("RPC failed to listen on TCP port: ", err)
		}
		defer listener.Close()
		logger.Infof("RPC server runing on: %+s", host+":"+cfg.RPCPort)
		go http.Serve(listener, nil)
		time.Sleep(time.Second) 
		sendHttp := rpcServer.NewHttpService(&ctx)
		err = sendHttp.Start(cfg.RPCPort)
		if err != nil {
			logger.Fatal("Http error: ", err)
		}
	}()


	wg.Add(1)
	go func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		defer wg.Done()
		wss := ws.NewWsService(&ctx)
		logger.Infof("WsAddress: %s\n", cfg.WSAddress)
		http.HandleFunc("/echo", wss.ServeWebSocket)

		logger.Fatal(http.ListenAndServe(cfg.WSAddress, nil))
	}()

	wg.Add(1)
	go func() {
		_, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		defer wg.Done()
		rest := rest.NewRestService(&ctx)

		router := rest.Initialize()
		logger.Infof("Starting REST api on: %s", cfg.RestAddress)
		logger.Fatal(router.Run(cfg.RestAddress))
	
	}()

	

}

func parserEvent(vLog types.Log, eventName string) {
	event := stake.StakeStakeEvent{}
	contractAbi, err := abi.JSON(strings.NewReader(string(stake.StakeMetaData.ABI)))

	if err != nil {
		logger.Fatal("contractAbi, err", err)
	}
	_err := contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
	if _err != nil {
		logger.Fatal("_err :  ", _err)
	}

	logger.Infof("Event Account: %s", event.Account) // foo
	logger.Infof("Event Amount: %d", event.Amount)
	logger.Infof("Event Timestamp: %d",event.Timestamp)
}

// var lobbyConn = []*websocket.Conn{}
// var verifiedConn = []*websocket.Conn{}
func GenerateRegsitrationData(cfg *configs.MainConfiguration) {
	regData := entities.RegisterationData{ChainId: "31337"}
	regData.Timestamp = 1723776438802; // uint64(time.Now().UnixMilli())
	// pkBig := new(big.Int).SetBytes(cfg.PublicKeySECP)
	// if !ok {
	// 	panic("Unable to generate big number")
	// }
	dHash := regData.GetHash()
	logger.Infof("DATAHHASH %s", new(big.Int).SetBytes(dHash))
	//pk, _ := hex.DecodeString(cfg.PrivateKey)
	signature, commitment, _ := regData.Sign(cfg.PrivateKeySECP)
	logger.Infof("RegData %s, %s, %s, %d, %s", hex.EncodeToString(signature),cfg.PrivateKey, hex.EncodeToString(commitment), regData.Timestamp, hex.EncodeToString(cfg.PublicKeySECP))
}
