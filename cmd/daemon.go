/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/ByteGum/go-icms/pkg/core/chain/evm"
	"github.com/ByteGum/go-icms/pkg/core/chain/evm/abis/stake"
	"github.com/ByteGum/go-icms/pkg/core/db"
	p2p "github.com/ByteGum/go-icms/pkg/core/p2p"
	utils "github.com/ByteGum/go-icms/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"github.com/ipfs/go-datastore/query"
	"github.com/spf13/cobra"

	rpcServer "github.com/ByteGum/go-icms/pkg/core/rpc"
	ws "github.com/ByteGum/go-icms/pkg/core/ws"
)

var logger = utils.Logger

const (
	TESTNET string = "/icm/testing"
	MAINNET        = "/icm/mainnet"
)

type Flag string

const (
	PRIVATE_KEY      Flag = "private-key"
	EVM_PRIVATE_KEY  Flag = "evm-private-key"
	NODE_PRIVATE_KEY Flag = "node-private-key"
	NETWORK               = "network"
	RPC_PORT         Flag = "rpc-port"
	WS_ADDRESS       Flag = "ws-address"
)
const MaxBatchSize = 1000

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemonFunc(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	daemonCmd.Flags().StringP(string(PRIVATE_KEY), "r", "", "(Deprecated) The evm private key. Please use --evm-private-key")
	daemonCmd.Flags().StringP(string(EVM_PRIVATE_KEY), "e", "", "The evm private key. This is the key used to sign handshakes and messages")
	daemonCmd.Flags().StringP(string(NODE_PRIVATE_KEY), "k", "", "The node private key. This is the nodes identity")
	daemonCmd.Flags().StringP(string(NETWORK), "m", MAINNET, "Network mode")
	daemonCmd.Flags().StringP(string(RPC_PORT), "p", utils.DefaultRPCPort, "RPC server port")
	daemonCmd.Flags().StringP(string(WS_ADDRESS), "w", utils.DefaultWebSocketAddress, "http service address")
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

	privateKey, err := cmd.Flags().GetString(string(PRIVATE_KEY))
	evmPrivateKey, err := cmd.Flags().GetString(string(EVM_PRIVATE_KEY))
	rpcPort, err := cmd.Flags().GetString(string(RPC_PORT))
	wsAddress, err := cmd.Flags().GetString(string(WS_ADDRESS))

	if err != nil || len(privateKey) == 0 {
		if len(evmPrivateKey) == 0 {
			panic("Private key is required. Use --private-key flag or environment var ICM_EVM_PRIVATE_KEY")
		} else {
			privateKey = evmPrivateKey
		}
	}
	if len(privateKey) > 0 {
		cfg.EvmPrivateKey = privateKey
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

	incomingMessagesc := make(chan *utils.ClientMessage)
	outgoingMessagesc := make(chan *utils.ClientMessage)
	outgoingMessagesP2Pc := make(chan *utils.ClientMessage)
	subscribersc := make(chan *utils.Subscription)
	subscriptiondp2pc := make(chan *utils.Subscription)
	clientHandshakec := make(chan *utils.ClientHandshake)
	incomingDeliveryProofsc := make(chan *utils.DeliveryProof)
	outgoingBatchc := make(chan *utils.Batch)
	outgoingProofc := make(chan *utils.DeliveryProof)
	pubsubInputBatchc := make(chan *utils.Batch)
	pubsubInputProofc := make(chan *utils.DeliveryProof)
	publishedSubc := make(chan *utils.Subscription)

	ctx = context.WithValue(ctx, utils.ConfigKey, &cfg)
	ctx = context.WithValue(ctx, utils.IncomingMessageCh, &incomingMessagesc)
	ctx = context.WithValue(ctx, utils.OutgoingMessageCh, &outgoingMessagesc)
	ctx = context.WithValue(ctx, utils.OutgoingMessageDP2PCh, &outgoingMessagesP2Pc)
	// incoming from client apps to daemon channel
	ctx = context.WithValue(ctx, utils.SubscribeCh, &subscribersc)
	// daemon to p2p channel
	ctx = context.WithValue(ctx, utils.SubscriptionDP2PCh, &subscriptiondp2pc)
	ctx = context.WithValue(ctx, utils.ClientHandShackCh, &clientHandshakec)
	ctx = context.WithValue(ctx, utils.IncomingDeliveryProofsCh, &incomingDeliveryProofsc)
	ctx = context.WithValue(ctx, utils.OutgoingBatchCh, &outgoingBatchc)
	ctx = context.WithValue(ctx, utils.OutgoingDeliveryProofCh, &outgoingProofc)
	ctx = context.WithValue(ctx, utils.PubsubDeliverProofCh, &pubsubInputBatchc)
	ctx = context.WithValue(ctx, utils.PubsubBatchCh, &pubsubInputProofc)
	// receiving subscription from other nodes channel
	ctx = context.WithValue(ctx, utils.PublishedSubCh, &publishedSubc)

	var wg sync.WaitGroup
	errc := make(chan error)

	stateStore := db.New(&ctx, utils.StateStore)
	validMessagesStore := db.New(&ctx, utils.ValidMessageStore)
	// unsentMessageStore := db.New(&ctx, utils.UnsentMessageStore)
	unsentMessageP2pStore := db.New(&ctx, utils.UnsentMessageStore)
	channelSubscriberStore := db.New(&ctx, utils.ChannelSubscriberStore)
	channelSubscribersCountStore := db.New(&ctx, utils.ChannelSubscribersCountStore)
	// sentMessageStore := db.Db(&ctx, utils.SentMessageStore)
	deliveryProofStore := db.New(&ctx, utils.DeliveryProofStore)
	localBatchStore := db.New(&ctx, utils.BatchStore)
	unconfurmedBatchStore := db.New(&ctx, utils.UnconfirmedDeliveryProofStore)

	defer wg.Wait()

	wg.Add(1)
	go func() {
		for {
			select {
			case inMessage, ok := <-incomingMessagesc:
				if !ok {
					logger.Errorf("Incoming Message channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				// VALIDATE AND DISTRIBUTE
				go func() {
					logger.Info("Received new message %s\n", inMessage.Message.Body.Message)
					validMessagesStore.Set(ctx, db.Key(inMessage.Key()), inMessage.ToJSON(), false)
					_reciever := inMessage.Message.Header.Receiver
					_recievers := strings.Split(_reciever, ":")
					_currentChannel := connectedSubscribers[_recievers[1]]
					for _, signerConn := range _currentChannel {
						for i := 0; i < len(signerConn); i++ {
							__msg := utils.SocketMessage{
								Type: "new-message",
								Data: inMessage.ToJSON(),
							}
							signerConn[i].WriteMessage(1, __msg.ToJSON())
						}
					}
				}()

			}

		}
	}()

	wg.Add(1)
	go func() {
		for {
			select {

			// attempt to push into outgoing message channel
			case outMessage, ok := <-outgoingMessagesc:
				if !ok {
					logger.Errorf("Outgoing Message channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					// VALIDATE AND DISTRIBUTE
					logger.Infof("\nSending out message %s\n", outMessage.Message.Body.Message)
					unsentMessageP2pStore.Set(ctx, db.Key(outMessage.Key()), outMessage.ToJSON(), false)
					outgoingMessagesP2Pc <- outMessage
					incomingMessagesc <- outMessage
					logger.Infof("\nSending out complete\n")
				}()

			case sub, ok := <-subscribersc:
				if !ok {
					logger.Errorf("Subscription channel closed!")
					return
				}
				if !utils.IsValidSubscription(*sub) {
					utils.Logger.Info("ITS NOT VALID!")
					continue
				}
				go func() {
					if sub.Broadcast {
						subscriptiondp2pc <- sub
					}
					trx, err := channelSubscribersCountStore.NewTransaction(ctx, false)
					logger.Info("TRANSACTION INITIATED ******")
					if err != nil {
						logger.Infof("Transaction err::: %w", err)
					}
					cscstore, err := trx.Get(ctx, db.Key(sub.Key()))
					increment := -1
					if sub.Action == utils.Join {
						increment = 1
						channelSubscriberStore.Set(ctx, db.Key(sub.Key()), sub.ToJSON(), false)

					} else {
						channelSubscriberStore.Delete(ctx, db.Key(sub.Key()))
					}
					if len(cscstore) == 0 {
						cscstore = []byte("0")
					}
					cscstoreint, err := strconv.Atoi(string(cscstore))
					cscstoreint += increment
					channelSubscribersCountStore.Set(ctx, db.Key(sub.Channel), []byte(strconv.Itoa(cscstoreint)), true)
					logger.Info("TRANSACTION ENDED ******")
					trx.Commit(ctx)
				}()

			case clientHandshake, ok := <-clientHandshakec:
				if !ok {
					logger.Errorf("Verification channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					// VALIDATE AND DISTRIBUTE
					logger.Infof("Signer:  %s\n", clientHandshake.Signer)
					results, err := channelSubscriberStore.Query(ctx, query.Query{
						Prefix: "/" + clientHandshake.Signer,
					})
					if err != nil {
						logger.Errorf("Channel Subscriber Store Query Error %w", err)
						return
					}
					entries, _err := results.Rest()
					for i := 0; i < len(entries); i++ {
						_sub, _ := utils.SubscriptionFromBytes(entries[i].Value)
						if connectedSubscribers[_sub.Channel] == nil {
							connectedSubscribers[_sub.Channel] = map[string][]*websocket.Conn{}
						}
						connectedSubscribers[_sub.Channel][_sub.Subscriber] = append(connectedSubscribers[_sub.Channel][_sub.Subscriber], clientHandshake.Socket)

					}
					logger.Infof("results:  %s  -  %w\n", entries[0].Value, _err)
				}()

			case proof, ok := <-incomingDeliveryProofsc:
				if !ok {
					logger.Errorf("Incoming delivery proof channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					err := deliveryProofStore.Set(ctx, db.Key(proof.Key()), proof.ToJSON(), true)
					if err == nil {
						// msg, err := validMessagesStore.Get(ctx, db.Key(fmt.Sprintf("/%s/%s", proof.MessageSender, proof.MessageHash)))
						// if err != nil {
						// 	// invalid proof or proof has been tampered with
						// 	return
						// }
						// get signer of proof
						susbscriber, err := utils.GetSigner(proof.ToString(), proof.Signature)
						if err != nil {
							// invalid proof or proof has been tampered with
							return
						}
						// check if the signer of the proof is a member of the channel
						isSubscriber, err := channelSubscriberStore.Has(ctx, db.Key("/"+susbscriber+"/"+proof.MessageSignature))
						if isSubscriber {
							// proof is valid, so we should add to a new or existing batch
							var batch utils.Batch
							var err error
							txn, err := stateStore.NewTransaction(ctx, false)
							if err != nil {
								logger.Errorf("State query errror %w", err)
								// invalid proof or proof has been tampered with
								return
							}
							batchData, err := txn.Get(ctx, db.Key(utils.CurrentBatchState))
							if err != nil {
								logger.Errorf("State query errror %w", err)
								// invalid proof or proof has been tampered with
								txn.Discard(ctx)
								return
							}
							if len(batchData) > 0 && batch.Size < MaxBatchSize {
								batch, err = utils.BatchFromBytes(batchData)
								if err != nil {
									logger.Errorf("Invalid batch %w", err)
									// invalid proof or proof has been tampered with
									txn.Discard(ctx)
									return
								}
							} else {
								// generate a new batch
								batch = utils.NewBatch()

							}
							batch.Size += 1
							if batch.Size >= MaxBatchSize {
								batch.Closed = true
								batch.NodeHeight = utils.GetNodeHeight()
							}
							// save the proof and the batch
							batch.Hash = hexutil.Encode(utils.Hash(proof.Signature + batch.Hash))
							err = txn.Put(ctx, db.Key(utils.CurrentBatchState), batch.ToJSON())
							if err != nil {
								logger.Errorf("Unable to udate State store errror %w", err)
								txn.Discard(ctx)
								return
							}
							proof.Batch = batch.BatchId
							proof.Index = batch.Size
							err = deliveryProofStore.Put(ctx, db.Key(proof.Key()), proof.ToJSON())
							if err != nil {
								txn.Discard(ctx)
								logger.Errorf("Unable to save proof to store error %w", err)
								return
							}
							err = localBatchStore.Put(ctx, db.Key(utils.CurrentBatchState), batch.ToJSON())
							if err != nil {
								logger.Errorf("Unable to save batch error %w", err)
								txn.Discard(ctx)
								return
							}
							err = txn.Commit(ctx)
							if err != nil {
								logger.Errorf("Unable to commit state update transaction errror %w", err)
								txn.Discard(ctx)
								return
							}
							// dispatch the proof and the batch
							if batch.Closed {
								outgoingBatchc <- &batch
							}
							outgoingProofc <- proof

						}

					}

				}()
			case batch, ok := <-pubsubInputBatchc:
				if !ok {
					logger.Errorf("PubsubInputBatch channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					unconfurmedBatchStore.Put(ctx, db.Key(batch.Key()), batch.ToJSON())
				}()
			case proof, ok := <-pubsubInputProofc:
				if !ok {
					logger.Errorf("PubsubInputBatch channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					unconfurmedBatchStore.Put(ctx, db.Key(proof.BatchKey()), proof.ToJSON())
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
		rpc.RegisterName("RpcService", rpcServer.NewRpcService(&ctx))
		listener, err := net.Listen("tcp", cfg.RPCHost+":"+rpcPort)
		if err != nil {
			logger.Fatal("ListenTCP error: ", err)
		}
		logger.Infof("RPC server runing on: %+s", cfg.RPCHost+":"+rpcPort)
		for {
			conn, err := listener.Accept()
			if err != nil {
				// wg.Done()
				logger.Fatalf("Accept error: ", err)
			}
			logger.Infof("New connection: %+v\n", conn.RemoteAddr())

			go jsonrpc.ServeConn(conn)
		}
	}()

	wg.Add(1)
	go func() {
		wss := ws.NewWsService(&ctx)
		logger.Infof("wsAddress: %s\n", wsAddress)
		http.HandleFunc("/echo", wss.ServeWebSocket)

		log.Fatal(http.ListenAndServe(wsAddress, nil))
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
// 				verifiedRequest, _ := utils.VerificationRequestFromBytes(message)
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
