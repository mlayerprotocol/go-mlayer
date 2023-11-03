/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
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
	"github.com/mlayerprotocol/go-mlayer/pkg/core/chain/evm"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/chain/evm/abis/stake"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/db"
	p2p "github.com/mlayerprotocol/go-mlayer/pkg/core/p2p"
	processor "github.com/mlayerprotocol/go-mlayer/pkg/core/processor"
	rpcServer "github.com/mlayerprotocol/go-mlayer/pkg/core/rpc"
	ws "github.com/mlayerprotocol/go-mlayer/pkg/core/ws"
	utils "github.com/mlayerprotocol/go-mlayer/utils"
	"github.com/spf13/cobra"
)

var logger = &utils.Logger

const (
	TESTNET string = "/icm/testing"
	MAINNET        = "/icm/mainnet"
)

type Flag string

const (
	NETWORK_PRIVATE_KEY Flag = "network-private-key"
	NODE_PRIVATE_KEY    Flag = "node-private-key"
	NETWORK                  = "network"
	RPC_PORT            Flag = "rpc-port"
	WS_ADDRESS          Flag = "ws-address"
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
	daemonCmd.Flags().StringP(string(RPC_PORT), "p", utils.DefaultRPCPort, "RPC server port")
	daemonCmd.Flags().StringP(string(WS_ADDRESS), "w", utils.DefaultWebSocketAddress, "http service address")
	daemonCmd.Flags().StringP(string(DATA_DIR), "d", utils.DefaultDataDir, "data directory")

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
} // use default options

func daemonFunc(cmd *cobra.Command, args []string) {
	cfg := utils.Config
	ctx := context.Background()

	connectedSubscribers := map[string]map[string][]*websocket.Conn{}

	incomingEventsC := make(chan types.Log)

	networkPrivateKey, err := cmd.Flags().GetString(string(NETWORK_PRIVATE_KEY))
	rpcPort, err := cmd.Flags().GetString(string(RPC_PORT))
	wsAddress, err := cmd.Flags().GetString(string(WS_ADDRESS))

	if err != nil || len(networkPrivateKey) == 0 {
		panic("network_private_key is required. Use --network-private-key flag or environment var ML_NETWORK_PRIVATE_KEY")
	}
	if len(networkPrivateKey) > 0 {
		cfg.NetworkPrivateKey = networkPrivateKey
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

	if rpcPort == utils.DefaultRPCPort && len(cfg.RPCPort) > 0 {
		rpcPort = cfg.RPCPort
	}

	ctx = context.WithValue(ctx, utils.ConfigKey, &cfg)
	ctx = context.WithValue(ctx, utils.IncomingMessageChId, &utils.IncomingMessagesP2P2_D_c)
	ctx = context.WithValue(ctx, utils.OutgoingMessageChId, &utils.SentMessagesRPC_D_c)
	ctx = context.WithValue(ctx, utils.OutgoingMessageDP2PChId, &utils.OutgoingMessagesD_P2P_c)
	// incoming from client apps to daemon channel
	ctx = context.WithValue(ctx, utils.SubscribeChId, &utils.SubscribersRPC_D_c)
	// daemon to p2p channel
	ctx = context.WithValue(ctx, utils.SubscriptionDP2PChId, &utils.SubscriptionD_P2P_C)
	ctx = context.WithValue(ctx, utils.ClientHandShackChId, &utils.ClientHandshakeC)
	ctx = context.WithValue(ctx, utils.OutgoingDeliveryProof_BlockChId, &utils.OutgoingDeliveryProof_BlockC)
	ctx = context.WithValue(ctx, utils.OutgoingDeliveryProofChId, &utils.OutgoingDeliveryProofC)
	ctx = context.WithValue(ctx, utils.PubsubDeliverProofChId, &utils.PubSubInputBlockC)
	ctx = context.WithValue(ctx, utils.PubSubBlockChId, &utils.PubSubInputProofC)
	// receiving subscription from other nodes channel
	ctx = context.WithValue(ctx, utils.PublishedSubChId, &utils.PublishedSubC)

	var wg sync.WaitGroup
	errc := make(chan error)

	deliveryProofBlockStateStore := db.New(&ctx, utils.DeliveryProofBlockStateStore)
	subscriptionBlockStateStore := db.New(&ctx, utils.SubscriptionBlockStateStore)

	// stores messages that have been validated
	validMessagesStore := db.New(&ctx, utils.ValidMessageStore)
	// unsentMessageStore := db.New(&ctx, utils.UnsentMessageStore)
	unsentMessageP2pStore := db.New(&ctx, utils.UnsentMessageStore)
	channelSubscriptionStore := db.New(&ctx, utils.ChannelSubscriptionStore)
	newChannelSubscriptionStore := db.New(&ctx, utils.NewChannelSubscriptionStore)
	channelsubscriptionCountStore := db.New(&ctx, utils.ChannelSubscriptionCountStore)
	// sentMessageStore := db.Db(&ctx, utils.SentMessageStore)
	deliveryProofStore := db.New(&ctx, utils.DeliveryProofStore)
	localDPBlockStore := db.New(&ctx, utils.DeliveryProofBlockStore)
	unconfurmedBlockStore := db.New(&ctx, utils.UnconfirmedDeliveryProofStore)

	ctx = context.WithValue(ctx, utils.NewChannelSubscriptionStore, newChannelSubscriptionStore)

	defer wg.Wait()

	wg.Add(1)
	go func() {
		for {
			select {
			case inMessage, ok := <-utils.IncomingMessagesP2P2_D_c:
				if !ok {
					logger.Errorf("Incoming Message channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				// VALIDATE AND DISTRIBUTE
				go func() {
					logger.Infof("Received new message %s\n", inMessage.Message.Body.MessageHash)
					validMessagesStore.Set(ctx, db.Key(inMessage.Key()), inMessage.Pack(), false)
					_reciever := inMessage.Message.Header.Receiver
					_recievers := strings.Split(_reciever, ":")
					_currentChannel := connectedSubscribers[_recievers[1]]
					logger.Info("connectedSubscribers : ", connectedSubscribers, "---", _reciever)
					logger.Info("_currentChannel : ", _currentChannel, "/n")
					for _, signerConn := range _currentChannel {
						for i := 0; i < len(signerConn); i++ {
							signerConn[i].WriteMessage(1, inMessage.Pack())
						}
					}
				}()

			}

		}
	}()
	wg.Add(1)
	go processor.ProcessNewSubscription(
		ctx,
		subscriptionBlockStateStore,
		channelsubscriptionCountStore,
		newChannelSubscriptionStore,
		channelSubscriptionStore,
		&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {

			// attempt to push into outgoing message channel
			case outMessage, ok := <-utils.SentMessagesRPC_D_c:
				if !ok {
					logger.Errorf("Outgoing Message channel closed. Please restart server to try or adjust buffer size in config")

					return
				}
				go processor.ProcessSentMessage(ctx, unsentMessageP2pStore, outMessage)

			case sub, ok := <-utils.SubscribersRPC_D_c:
				if !ok {
					logger.Errorf("Subscription channel closed!")
					return
				}
				if !utils.IsValidSubscription(*sub, true) {
					utils.Logger.Info("ITS NOT VALID!")
					continue
				}
				// go processor.ProcessNewSubscription(ctx, sub, channelsubscribersRPC_D_countStore, channelSubscriptionStore)

			case clientHandshake, ok := <-utils.ClientHandshakeC:
				if !ok {
					logger.Errorf("Verification channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go processor.ValidateMessageClient(ctx, &connectedSubscribers, clientHandshake, channelSubscriptionStore)

			case proof, ok := <-utils.IncomingDeliveryProofsC:
				if !ok {
					logger.Errorf("Incoming delivery proof channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go processor.ValidateAndAddToDeliveryProofToBlock(ctx,
					proof,
					deliveryProofStore,
					channelSubscriptionStore,
					deliveryProofBlockStateStore,
					localDPBlockStore,
					MaxDeliveryProofBlockSize,
					&deliveryProofBlockMutex,
				)

			case batch, ok := <-utils.PubSubInputBlockC:
				if !ok {
					logger.Errorf("PubsubInputBlock channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					unconfurmedBlockStore.Put(ctx, db.Key(batch.Key()), batch.Pack())
				}()
			case proof, ok := <-utils.PubSubInputProofC:
				if !ok {
					logger.Errorf("PubsubInputBlock channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					unconfurmedBlockStore.Put(ctx, db.Key(proof.BlockKey()), proof.Pack())
				}()

			}

		}
	}()

	wg.Add(1)
	go func() {

		if err := recover(); err != nil {
			wg.Done()
			errc <- fmt.Errorf("P2P error: %g", err)
		}
		p2p.Run(&ctx)
	}()
	wg.Add(1)

	wg.Add(1)
	go func() {
		_, client, contractAddress, err := evm.StakeContract(cfg.EVMRPCWss, cfg.StakeContract)
		if err != nil {
			log.Fatal(err, cfg.EVMRPCWss, cfg.StakeContract)
		}
		query := ethereum.FilterQuery{
			// FromBlock: big.NewInt(23506010),
			// ToBlock:   big.NewInt(23506110),

			Addresses: []common.Address{contractAddress},
		}

		// logs, err := client.FilterLogs(context.Background(), query)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// parserEvent(logs[0], "StakeEvent")

		// logger.Infof("Past Events", logs)
		// incomingEventsC

		sub, err := client.SubscribeFilterLogs(context.Background(), query, incomingEventsC)
		if err != nil {
			log.Fatal(err, "SubscribeFilterLogs")
		}

		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-incomingEventsC:
				fmt.Println(vLog) // pointer to event log
				parserEvent(vLog, "StakeEvent")
			}
		}

	}()

	wg.Add(1)
	go func() {
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
		wss := ws.NewWsService(&ctx)
		logger.Infof("wsAddress: %s\n", wsAddress)
		http.HandleFunc("/echo", wss.ServeWebSocket)

		log.Fatal(http.ListenAndServe(wsAddress, nil))
	}()

	wg.Add(1)
	go func() {
		sendHttp := rpcServer.NewHttpService(&ctx)
		err := sendHttp.Start()
		if err != nil {
			logger.Fatalf("Http error: ", err)
		}
		logger.Infof("New http connection")
	}()

}

func parserEvent(vLog types.Log, eventName string) {
	event := stake.StakeStakeEvent{}
	contractAbi, err := abi.JSON(strings.NewReader(string(stake.StakeMetaData.ABI)))

	if err != nil {
		log.Fatal("contractAbi, err", err)
	}
	_err := contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
	if _err != nil {
		log.Fatal("_err :  ", _err)
	}

	fmt.Println(event.Account) // foo
	fmt.Println(event.Amount)
	fmt.Println(event.Timestamp)
}

var lobbyConn = []*websocket.Conn{}
var verifiedConn = []*websocket.Conn{}

// func ServeWebSocket(w http.ResponseWriter, r *http.Request) {

// 	c, err := upgrader.Upgrade(w, r, nil)
// 	log.Print("New ServeWebSocket c : ", c.RemoteAddr())

// 	if err != nil {
// 		log.Print("upgrade:", err)
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
// 			log.Println("read:", err)
// 			break

// 		} else {
// 			err = c.WriteMessage(mt, (append(message, []byte("recieved Signature")...)))
// 			if err != nil {
// 				log.Println("Error:", err)
// 			} else {
// 				// signature := string(message)
// 				verifiedRequest, _ := utils.UnpackVerificationRequest(message)
// 				log.Println("verifiedRequest.Message: ", verifiedRequest.Message)

// 				if utils.VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
// 					verifiedConn = append(verifiedConn, c)
// 					hasVerifed = true
// 					log.Println("Verification was successful: ", verifiedRequest)
// 				}
// 				log.Println("message:", string(message))
// 				log.Printf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
// 			}

// 		}
// 	}

// }
