/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/hex"

	"sync"

	// "net/rpc/jsonrpc"

	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	// "github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/internal/channelpool"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/mlayerprotocol/go-mlayer/pkg/node"
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
	REST_ADDRESS        Flag = "rest-address"
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

func daemonFunc(cmd *cobra.Command, args []string) {
	cfg := configs.Config
	ctx := context.Background()

	sql.Init()

	rpcPort, _ := cmd.Flags().GetString(string(RPC_PORT))
	wsAddress, _ := cmd.Flags().GetString(string(WS_ADDRESS))
	restAddress, _ := cmd.Flags().GetString(string(REST_ADDRESS))

	networkPrivateKey, err := cmd.Flags().GetString(string(NETWORK_PRIVATE_KEY))
	if err != nil || len(networkPrivateKey) == 0 {
		panic("network_private_key is required. Use --network-private-key flag or environment var ML_NETWORK_PRIVATE_KEY")
	}
	if len(cfg.AddressPrefix) == 0 {
		cfg.AddressPrefix = "ml"
	}
	if len(networkPrivateKey) > 0 {
		cfg.NetworkPrivateKey = networkPrivateKey
		cfg.NetworkPublicKey = crypto.GetPublicKeyEDD(networkPrivateKey)
		key, err := hex.DecodeString(cfg.NetworkPublicKey)
		if err != nil {
			panic(err)
		}
		cfg.NetworkKeyAddress = crypto.ToBech32Address(key, cfg.AddressPrefix)
	}

	if len(wsAddress) > 0 {
		cfg.WSAddress = wsAddress
	}

	if len(restAddress) > 0 {
		cfg.RestAddress = restAddress
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
	if len(rpcPort) > 0 {
		cfg.RPCPort = rpcPort
	}
	if len(cfg.RPCPort) == 0 {
		cfg.RPCPort = constants.DefaultRPCPort
	}

	chain.Init(&cfg)

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

	node.Start(&ctx)
}

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
