/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"slices"

	"sync"

	// "net/rpc/jsonrpc"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
	"github.com/mlayerprotocol/go-mlayer/internal/chain/api"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"

	// "github.com/mlayerprotocol/go-mlayer/entities"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/sql"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/mlayerprotocol/go-mlayer/pkg/node"
	"github.com/spf13/cobra"
)

var logger = &log.Logger



type Flag string

const (
	NETWORK_ADDRESS_PRFIX Flag = "network-address-prefix"
	CHAIN_ID Flag = "chain-id"
	PRIVATE_KEY Flag = "private-key"
	PROTOCOL_VERSION    Flag  = "protocol-version"
	RPC_PORT            Flag = "rpc-port"
	WS_ADDRESS          Flag = "ws-address"
	REST_ADDRESS        Flag = "rest-address"
	DATA_DIR            Flag = "data-dir"
	LISTENERS            Flag = "listen"
	KEYSTORE_DIR         Flag = "key_store_dir"
	KEYSTORE_PASSWORD         Flag = "key_store_password"
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
	daemonCmd.Flags().StringP(string(NETWORK_ADDRESS_PRFIX), "p", "", "The network address prefix. This determines the operational network e.g. ml=>mainnet, mldev=>devnet,mltest=>testnet")
	daemonCmd.Flags().StringP(string(PRIVATE_KEY), "k", "", "The deligated operators private key. This is the key used to sign handshakes and messages. The coresponding public key must be assigned to the validator")
	daemonCmd.Flags().StringP(string(PROTOCOL_VERSION), "v", constants.DefaultProtocolVersion, "Protocol version")
	daemonCmd.Flags().StringP(string(RPC_PORT), "r", constants.DefaultRPCPort, "RPC server port")
	daemonCmd.Flags().StringP(string(WS_ADDRESS), "w", constants.DefaultWebSocketAddress, "ws service address")
	daemonCmd.Flags().StringP(string(REST_ADDRESS), "R", constants.DefaultRestAddress, "rest api service address")
	daemonCmd.Flags().StringP(string(DATA_DIR), "d", constants.DefaultDataDir, "data storage directory")
	daemonCmd.Flags().StringSliceP(string(LISTENERS), "l", []string{}, "libp2p multiaddress array eg. [\"/ip4/127.0.0.1/tcp/5000/ws\", \"/ip4/127.0.0.1/tcp/5001\"]")
	daemonCmd.Flags().StringP(string(KEYSTORE_DIR), "K", "", "path to keystore directory")
	daemonCmd.Flags().StringP(string(KEYSTORE_PASSWORD), "P", "", "password for decripting key store")
}

func daemonFunc(cmd *cobra.Command, _ []string) {
	cfg := configs.Config
	ctx := context.Background()

	
	defer node.Start(&ctx)
	defer func () {
		// chain.Network = chain.Init(&cfg)
		chain.RegisterProvider(
			"31337", api.NewGenericAPI(),
		)
		ethAPI, err := api.NewEthAPI(cfg.ChainId, cfg.EvmRpcConfig[string(cfg.ChainId)], &cfg.PrivateKeySECP)
		if err != nil {
			panic(err)
		}
		chain.RegisterProvider(
			"84532", ethAPI,
		)
		// chain.DefaultProvider = chain.Network.Default()
		ownerAddress, _ := hex.DecodeString(constants.ADDRESS_ZERO)
		if cfg.Validator {
			ownerAddress, err = chain.Provider(cfg.ChainId).GetValidatorLicenseOwnerAddress(cfg.PublicKeySECP)
		} else {
			ownerAddress, err = chain.Provider(cfg.ChainId).GetSentryLicenseOwnerAddress(cfg.PublicKeySECP)
		}
		if err != nil  {
			logger.Fatalf("unable to get license owner: %v", err)
		}
		if hex.EncodeToString(ownerAddress) == constants.ADDRESS_ZERO {
			if cfg.Validator {
				logger.Fatalf("Failed to run in validator mode because no license is assigned to this operators public key (SECP).")
			}
			logger.Debugf("Operator not yet deligated. Running is archive mode.")
		}
		cfg.OwnerAddress = common.BytesToAddress(ownerAddress)
	}()
	defer sql.Init(&cfg)
	

	

	rpcPort, _ := cmd.Flags().GetString(string(RPC_PORT))
	wsAddress, _ := cmd.Flags().GetString(string(WS_ADDRESS))
	restAddress, _ := cmd.Flags().GetString(string(REST_ADDRESS))
	listeners, _ := cmd.Flags().GetStringSlice(string(LISTENERS))

	
	cfg.Context = &ctx

	if len(cfg.ChainId) == 0 {
		cfg.ChainId = "ml"
	}
	chainId, _ := cmd.Flags().GetString(string(CHAIN_ID))
	if len(chainId) > 0 {
		cfg.ChainId = configs.ChainId(chainId)
	}

	if len(cfg.ChainId) == 0 {
		cfg.ChainId = "ml"
	}
	prefix, _ := cmd.Flags().GetString(string(NETWORK_ADDRESS_PRFIX))
	if len(prefix) > 0 {
		cfg.AddressPrefix = prefix
	}
	cfg = injectPrivateKey(&cfg, cmd)

	if len(wsAddress) > 0 {
		cfg.WSAddress = wsAddress
	}

	if len(restAddress) > 0 {
		cfg.RestAddress = restAddress
	}

	dataDir, _ := cmd.Flags().GetString(string(DATA_DIR))
	if len(dataDir) > 0 {
		cfg.DataDir = dataDir
	}

	if len(cfg.SQLDB.DbStoragePath) == 0 {
		cfg.SQLDB.DbStoragePath = fmt.Sprintf("%s/store/sql", cfg.DataDir)
	}


	protocolVersion, _ := cmd.Flags().GetString(string(PROTOCOL_VERSION))
	
	if len(protocolVersion) > 0 && protocolVersion != constants.DefaultProtocolVersion  {
		cfg.ProtocolVersion = protocolVersion
	}
	if len(cfg.ProtocolVersion) == 0 {
		cfg.ProtocolVersion = constants.DefaultProtocolVersion
	}

	if !slices.Contains(constants.VALID_PROTOCOLS, cfg.ProtocolVersion) {
		panic("Invalid protocol version provided")
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
	
	if len(listeners) > 0 {
		cfg.ListenerAdresses = listeners
	}

	

	// ****** INITIALIZE CONTEXT ****** //

	ctx = context.WithValue(ctx, constants.ConfigKey, &cfg)

	// ADD EVENT  SUBSCRIPTION CHANNELS TO THE CONTEXT
	// ctx = context.WithValue(ctx, constants.IncomingAuthorizationEventChId, &channelpool.AuthorizationEvent_SubscriptionC)
	// ctx = context.WithValue(ctx, constants.IncomingTopicEventChId, &channelpool.IncomingTopicEventSubscriptionC)

	// ADD EVENT BROADCAST CHANNELS TO THE CONTEXT
	// ctx = context.WithValue(ctx, constants.BroadcastAuthorizationEventChId, &channelpool.AuthorizationEventPublishC)
	// ctx = context.WithValue(ctx, constants.BroadcastTopicEventChId, &channelpool.TopicEventPublishC)
	// ctx = context.WithValue(ctx, constants.BroadcastSubnetEventChId, &channelpool.SubnetEventPublishC)

	// // CLEANUP
	// ctx = context.WithValue(ctx, constants.IncomingMessageChId, &channelpool.IncomingMessageEvent_P2P_D_C)
	// ctx = context.WithValue(ctx, constants.OutgoingMessageChId, &channelpool.NewPayload_Cli_D_C)
	// ctx = context.WithValue(ctx, constants.OutgoingMessageDP2PChId, &channelpool.OutgoingMessageEvents_D_P2P_C)
	// incoming from client apps to daemon channel
	// ctx = context.WithValue(ctx, constants.SubscribeChId, &channelpool.Subscribers_RPC_D_C)
	// daemon to p2p channel
	// ctx = context.WithValue(ctx, constants.SubscriptionDP2PChId, &channelpool.Subscription_D_P2P_C)
	// ctx = context.WithValue(ctx, constants.ClientHandShackChId, &channelpool.ClientHandshakeC)
	// ctx = context.WithValue(ctx, constants.OutgoingDeliveryProof_BlockChId, &channelpool.OutgoingDeliveryProof_BlockC)
	// ctx = context.WithValue(ctx, constants.OutgoingDeliveryProofChId, &channelpool.OutgoingDeliveryProofC)
	// ctx = context.WithValue(ctx, constants.PubsubDeliverProofChId, &channelpool.PubSubInputBlockC)
	// ctx = context.WithValue(ctx, constants.PubSubBlockChId, &channelpool.PubSubInputProofC)
	// // receiving subscription from other nodes channel
	// ctx = context.WithValue(ctx, constants.PublishedSubChId, &channelpool.PublishedSubC)

	ctx = context.WithValue(ctx, constants.SQLDB, &sql.SqlDb)

	
	
	

}

func injectPrivateKey(cfg *configs.MainConfiguration, cmd *cobra.Command) configs.MainConfiguration {
	operatorPrivateKey, _ := cmd.Flags().GetString(string(PRIVATE_KEY))
	// if err != nil || len(operatorPrivateKey) == 0 {
	// 	logger.Fatal("operators private_key is required. Use --private-key flag or environment var ML_PRIVATE_KEY")

	// }
	pkFlagLen := len(operatorPrivateKey)
	if pkFlagLen > 0 {
		cfg.PrivateKey = operatorPrivateKey
	}
	if len(cfg.PrivateKey) == 0 {
		//check the keystore
		password, _ := cmd.Flags().GetString(string(KEYSTORE_PASSWORD))
		if password == "" {
			fmt.Println("Enter your keystore password: ")
			inputPass, err := readInputSecurely()
			if err!= nil {
				panic("provide a keystore password")
			}
			password = string(inputPass)
		}
		ksDir, _ := cmd.Flags().GetString(string(KEYSTORE_DIR))
		if len(ksDir) == 0 {
			ksDir = cfg.KeyStoreDir
		}
		if len(ksDir) == 0 {
			ksDir = "./data/keystores/"
		}
		privKey, err := loadPrivateKeyFromKeyStore(string(password), "account", ksDir)
		if err != nil {
			panic(err)
		}
		cfg.PrivateKey = hex.EncodeToString(privKey)
	}
	
	if  len(cfg.PrivateKey) != 64 {
		panic("--private-key must be 32 bytes long")
	}
	

	// conver private key to edd
	pk, err :=  hex.DecodeString(cfg.PrivateKey)
	if err != nil {
		panic( err)
	}
	// SECP KEYS
	cfg.PrivateKeySECP = pk
	_, pubKey := btcec.PrivKeyFromBytes(pk)
	cfg.PublicKeySECP = pubKey.SerializeCompressed()
	
	// EDD KEYS
	cfg.PrivateKeyEDD = ed25519.NewKeyFromSeed(pk)
	cfg.PrivateKey = hex.EncodeToString(cfg.PrivateKeyEDD)
	cfg.PublicKeyEDD = cfg.PrivateKeyEDD[32:]
	cfg.PublicKey = hex.EncodeToString(cfg.PublicKeyEDD)
	cfg.OperatorAddress = crypto.ToBech32Address(cfg.PublicKeySECP, string(cfg.AddressPrefix))

	return *cfg
}