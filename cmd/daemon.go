/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"strings"
	"sync"

	"net"
	"net/http"
	"net/rpc"

	// "net/rpc/jsonrpc"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	// "github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/internal/message"
	"github.com/mlayerprotocol/go-mlayer/internal/subscription"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/chain/evm"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/chain/evm/abis/stake"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	p2p "github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/rest"
	rpcServer "github.com/mlayerprotocol/go-mlayer/pkg/core/rpc"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	ws "github.com/mlayerprotocol/go-mlayer/pkg/core/ws"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/spf13/cobra"
)

var logger = &log.Logger

const (
	TESTNET string = "/mlayer/testing"
	MAINNET        = "/mlayer/mainnet"
)

type Flag string

const (
	NETWORK_PRIVATE_KEY Flag = "network-private-key"
	NODE_PRIVATE_KEY    Flag = "node-private-key"
	NETWORK                  = "network"
	RPC_PORT            Flag = "rpc-port"
	WS_ADDRESS          Flag = "ws-address"
	REST_ADDRESS          Flag = "rest-address"
	DATA_DIR            Flag = "data-dir"
)
const MaxDeliveryProofBlockSize = 1000

var deliveryProofBlockMutex sync.RWMutex

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Runs goml as a daemon",
	Long: `Use this command to run goml as a daemon:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemonFunc(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.Flags().StringP(string(NETWORK_PRIVATE_KEY), "e", "", "The network private key. This is the key used to sign handshakes and messages")
	daemonCmd.Flags().StringP(string(NODE_PRIVATE_KEY), "k", "", "The node private key. This is the nodes identity")
	daemonCmd.Flags().StringP(string(NETWORK), "m", MAINNET, "Network mode")
	daemonCmd.Flags().StringP(string(RPC_PORT), "p", constants.DefaultRPCPort, "RPC server port")
	daemonCmd.Flags().StringP(string(WS_ADDRESS), "w", constants.DefaultWebSocketAddress, "ws service address")
	daemonCmd.Flags().StringP(string(REST_ADDRESS), "r", constants.DefaultRestAddress, "rest api service address")
	daemonCmd.Flags().StringP(string(DATA_DIR), "d", constants.DefaultDataDir, "data directory")

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
} // use default options

func daemonFunc(cmd *cobra.Command, args []string) {
	cfg := configs.Config
	ctx := context.Background()

	sql.Init()

	connectedSubscribers := make(map[string]map[string][]interface{})

	incomingEventsC := make(chan types.Log)

	networkPrivateKey, err := cmd.Flags().GetString(string(NETWORK_PRIVATE_KEY))
	rpcPort, err := cmd.Flags().GetString(string(RPC_PORT))
	wsAddress, err := cmd.Flags().GetString(string(WS_ADDRESS))
	restAddress, err := cmd.Flags().GetString(string(REST_ADDRESS))

	if err != nil || len(networkPrivateKey) == 0 {
		panic("network_private_key is required. Use --network-private-key flag or environment var ML_NETWORK_PRIVATE_KEY")
	}
	if len(networkPrivateKey) > 0 {
		cfg.NetworkPrivateKey = networkPrivateKey
		cfg.NetworkPublicKey = crypto.GetPublicKeyEDD(networkPrivateKey)
		cfg.NetworkKeyAddress = crypto.ToBech32Address(cfg.NetworkPublicKey)
	}

	dataDir, err := cmd.Flags().GetString(string(DATA_DIR))
	if len(dataDir) > 0 {
		cfg.DataDir = dataDir
	}
	network, err := cmd.Flags().GetString(string(NETWORK))
	if err != nil || len(network) == 0 {
		if len(cfg.Network) == 0 {
			panic("Network required")
		}
	}
	if len(network) > 0 {
		cfg.Network = network
	}

	if rpcPort == constants.DefaultRPCPort && len(cfg.RPCPort) > 0 {
		rpcPort = cfg.RPCPort
	}

	chain.InitializeMlChain(&cfg)

	// ****** INITIALIZE CONTEXT ****** //

	ctx = context.WithValue(ctx, constants.ConfigKey, &cfg)

	// ADD EVENT  SUBSCRIPTION CHANNELS TO THE CONTEXT
	// ctx = context.WithValue(ctx, constants.IncomingAuthorizationEventChId, &channelpool.AuthorizationEvent_SubscriptionC)
	// ctx = context.WithValue(ctx, constants.IncomingTopicEventChId, &channelpool.IncomingTopicEventSubscriptionC)

	// ADD EVENT BROADCAST CHANNELS TO THE CONTEXT
	ctx = context.WithValue(ctx, constants.BroadcastAuthorizationEventChId, &channelpool.AuthorizationEventPublishC)
	ctx = context.WithValue(ctx, constants.BroadcastTopicEventChId, &channelpool.TopicEventPublishC)
	
	
	// CLEANUP
	ctx = context.WithValue(ctx, constants.IncomingMessageChId, &channelpool.IncomingMessageEvent_P2P_D_C)
	ctx = context.WithValue(ctx, constants.OutgoingMessageChId, &channelpool.NewPayload_Cli_D_C)
	ctx = context.WithValue(ctx, constants.OutgoingMessageDP2PChId, &channelpool.OutgoingMessageEvents_D_P2P_C)
	// incoming from client apps to daemon channel
	ctx = context.WithValue(ctx, constants.SubscribeChId, &channelpool.Subscribers_RPC_D_C)
	// daemon to p2p channel
	ctx = context.WithValue(ctx, constants.SubscriptionDP2PChId, &channelpool.Subscription_D_P2P_C)
	ctx = context.WithValue(ctx, constants.ClientHandShackChId, &channelpool.ClientHandshakeC)
	ctx = context.WithValue(ctx, constants.OutgoingDeliveryProof_BlockChId, &channelpool.OutgoingDeliveryProof_BlockC)
	ctx = context.WithValue(ctx, constants.OutgoingDeliveryProofChId, &channelpool.OutgoingDeliveryProofC)
	ctx = context.WithValue(ctx, constants.PubsubDeliverProofChId, &channelpool.PubSubInputBlockC)
	ctx = context.WithValue(ctx, constants.PubSubBlockChId, &channelpool.PubSubInputProofC)
	// receiving subscription from other nodes channel
	ctx = context.WithValue(ctx, constants.PublishedSubChId, &channelpool.PublishedSubC)

	ctx = context.WithValue(ctx, constants.SQLDB, &sql.Db)

	var wg sync.WaitGroup
	// errc := make(chan error)

	// deliveryProofBlockStateStore := db.New(&ctx, constants.DeliveryProofBlockStateStore)
	subscriptionBlockStateStore := db.New(&ctx, constants.SubscriptionBlockStateStore)

	// stores messages that have been validated
	// validMessagesStore := db.New(&ctx, constants.ValidMessageStore)
	unProcessedClientPayloadStore := db.New(&ctx, constants.UnprocessedClientPayloadStore)
	unsentMessageP2pStore := db.New(&ctx, constants.UnsentMessageStore)
	topicSubscriptionStore := db.New(&ctx, constants.TopicSubscriptionStore)
	newTopicSubscriptionStore := db.New(&ctx, constants.NewTopicSubscriptionStore)
	topicSubscriptionCountStore := db.New(&ctx, constants.TopicSubscriptionCountStore)
	// sentMessageStore := db.Db(&ctx, constants.SentMessageStore)
	// deliveryProofStore := db.New(&ctx, constants.DeliveryProofStore)
	// localDPBlockStore := db.New(&ctx, constants.DeliveryProofBlockStore)
	// unconfirmedBlockStore := db.New(&ctx, constants.UnconfirmedDeliveryProofStore)

	ctx = context.WithValue(ctx, constants.NewTopicSubscriptionStore, newTopicSubscriptionStore)
	ctx = context.WithValue(ctx, constants.UnprocessedClientPayloadStore, unProcessedClientPayloadStore)

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
					inMessage := inEvent.Payload.Data.(entities.ChatMessage)
					logger.Infof("Received new message %s\n", inMessage.Body.MessageHash)
					// validMessagesStore.Set(ctx, db.Key(inMessage.Key()), inMessage.MsgPack(), false)
					_reciever := inMessage.Header.Receiver
					_recievers := strings.Split(_reciever, ":")
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


	 wg.Add(1)
	go func() {
		defer wg.Done()
		subscription.ProcessNewSubscription(
		ctx,
		subscriptionBlockStateStore,
		topicSubscriptionCountStore,
		newTopicSubscriptionStore,
		topicSubscriptionStore,
		&wg)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		message.ProcessNewMessageEvent(ctx, unsentMessageP2pStore, &wg)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// defer func() {
		// if err := recover(); err != nil {
		// 	wg.Done()
		// 	errc <- fmt.Errorf("P2P error: %g", err)
		// }
		// }()
		
		p2p.Run(&ctx)
		if err != nil {
			wg.Done()
			panic(err)
		}
	}()
	

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, client, contractAddress, err := evm.StakeContract(cfg.EVMRPCWss, cfg.StakeContract)
		if err != nil {
			logger.Fatal(err, cfg.EVMRPCWss, cfg.StakeContract)
		}
		query := ethereum.FilterQuery{
			// FromBlock: big.NewInt(23506010),
			// ToBlock:   big.NewInt(23506110),

			Addresses: []common.Address{contractAddress},
		}

		// logs, err := client.FilterLogs(context.Background(), query)
		// if err != nil {
		// 	logger.Fatal(err)
		// }
		// parserEvent(logs[0], "StakeEvent")

		// logger.Infof("Past Events", logs)
		// incomingEventsC

		sub, err := client.SubscribeFilterLogs(context.Background(), query, incomingEventsC)
		if err != nil {
			logger.Fatal(err, "SubscribeFilterLogs")
		}

		for {
			select {
			case err := <-sub.Err():
				logger.Fatal(err)
			case vLog := <-incomingEventsC:
				fmt.Println(vLog) // pointer to event log
				parserEvent(vLog, "StakeEvent")
			}
		}

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rpc.Register(rpcServer.NewRpcService(&ctx))
		rpc.HandleHTTP()
		listener, err := net.Listen("tcp", cfg.RPCHost+":"+rpcPort)
		if err != nil {
			logger.Fatal("RPC failed to listen on TCP port: ", err)
		}
		logger.Infof("RPC server runing on: %+s", cfg.RPCHost+":"+rpcPort)
		go http.Serve(listener, nil)
		// for {
		// 	conn, err := listener.Accept()
		// 	if err != nil {
		// 		// wg.Done()
		// 		logger.Fatalf("Accept error: ", err)
		// 	}
		// 	logger.Infof("New connection: %+v\n", conn.RemoteAddr())

		// }

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wss := ws.NewWsService(&ctx)
		logger.Infof("wsAddress: %s\n", wsAddress)
		http.HandleFunc("/echo", wss.ServeWebSocket)

		logger.Fatal(http.ListenAndServe(wsAddress, nil))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rest := rest.NewRestService(&ctx)
		
		router := rest.Initialize()
		logger.Infof("Starting REST api on: %s", restAddress)
		logger.Fatal(router.Run(restAddress))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sendHttp := rpcServer.NewHttpService(&ctx)
		err := sendHttp.Start()
		if err != nil {
			logger.Fatal("Http error: ", err)
		}
		logger.Infof("New http connection")
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

	fmt.Println(event.Account) // foo
	fmt.Println(event.Amount)
	fmt.Println(event.Timestamp)
}

var lobbyConn = []*websocket.Conn{}
var verifiedConn = []*websocket.Conn{}

// func ServeWebSocket(w http.ResponseWriter, r *http.Request) {

// 	c, err := upgrader.Upgrade(w, r, nil)
// 	logger.Print("New ServeWebSocket c : ", c.RemoteAddr())

// 	if err != nil {
// 		logger.Print("upgrade:", err)
// 		return
// 	}
// 	defer c.Close()
// 	hasVerifed := false
// 	time.AfterFunc(5000*time.Millisecond, func() {

// 		if !hasVerifed {
// 			c.Close()
// 		}
// 	})
// 	_close := func(code int, t string) error {
// 		logger.Infof("code: %d, t: %s \n", code, t)
// 		return errors.New("Closed ")
// 	}
// 	c.SetCloseHandler(_close)
// 	for {
// 		mt, message, err := c.ReadMessage()
// 		if err != nil {
// 			logger.Println("read:", err)
// 			break

// 		} else {
// 			err = c.WriteMessage(mt, (append(message, []byte("recieved Signature")...)))
// 			if err != nil {
// 				logger.Println("Error:", err)
// 			} else {
// 				// signature := string(message)
// 				verifiedRequest, _ := entities.UnpackVerificationRequest(message)
// 				logger.Println("verifiedRequest.Message: ", verifiedRequest.Message)

// 				if constants.VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
// 					verifiedConn = append(verifiedConn, c)
// 					hasVerifed = true
// 					logger.Println("Verification was successful: ", verifiedRequest)
// 				}
// 				logger.Println("message:", string(message))
// 				logger.Printf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
// 			}

// 		}
// 	}

// }
