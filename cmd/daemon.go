/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"sync"

	"github.com/ByteGum/go-icms/pkg/core/db"
	p2p "github.com/ByteGum/go-icms/pkg/core/p2p"
	utils "github.com/ByteGum/go-icms/utils"
	"github.com/spf13/cobra"

	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcServer "github.com/ByteGum/go-icms/pkg/core/rpc"
)


var logger = utils.Logger

const (
	TESTNET string = "/icm/testing"
	MAINNET        = "/icm/mainnet"
)

type Flag string

const (
	PRIVATE_KEY Flag = "private-key"
	NETWORK          = "network"
)

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	daemonCmd.Flags().StringP(string(PRIVATE_KEY), "k", "", "Help message for toggle")
	daemonCmd.Flags().StringP(string(NETWORK), "m", MAINNET, "Network mode")
}

func daemonFunc(cmd *cobra.Command, args []string) {
	cfg := utils.Config
	ctx := context.Background()
	incomingMessagesc := make(chan utils.ClientMessage)
	outgoingMessagesc := make(chan utils.ClientMessage)

	privateKey, err := cmd.Flags().GetString(string(PRIVATE_KEY))
	if err != nil || len(privateKey) == 0 {
		if len(cfg.PrivateKey) == 0 {
			panic("Private key is required. Use --private-key flag or environment var ICM_PRIVATE_KEY")
		}
	}
	if len(privateKey) > 0 {
		cfg.PrivateKey = privateKey
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
	ctx = context.WithValue(ctx, "Config", cfg)
	ctx = context.WithValue(ctx, "IncomingMessageC", incomingMessagesc)
	ctx = context.WithValue(ctx, "OutgoingMessageC", outgoingMessagesc)
	var wg sync.WaitGroup
	errc := make(chan error)
	// dbPath, err := ioutil.TempDir("", "badger-test")
	// if err != nil {
	// 	errc <- fmt.Errorf("Could not read temp dir: %g", err)
	// }
	// ds, err := badgerds.NewDatastore(dbPath, nil)
	// if err != nil {
	// 	errc <- fmt.Errorf("Could not initialize ds: %g", err)
	// }
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
				logger.Info("Received new message %s\n", inMessage.Message.Body.Text)

				// attempt to push into outgoing message channel
			// case outMessage, ok := <-outgoingMessagesc:
			// 	if !ok {
			// 		logger.Errorf("Outgoing Message channel closed. Please restart server to try or adjust buffer size in config")
			// 		wg.Done()
			// 		return
			// 	}
			// 	// VALIDATE AND DISTRIBUTE
			// 	logger.Info("Received new message %s\n", outMessage.Message.Body.Text)

			}
			
		}
	}()

	wg.Add(1)
	go func() {

		if err := recover(); err != nil {
			wg.Done()
			errc <- fmt.Errorf("P2P error: %g", err)
		}
		// logger.WithFields(logrus.Fields{
		// 	"publicKey": "walrus",
		// }).Infof("publicKey %s", priv)
		p2p.Run(&ctx)
	}()
	wg.Add(1)
	go func() {
		if err := recover(); err != nil {
			wg.Done()
			errc <- fmt.Errorf("db error: %g", err)
		}
		defer wg.Done()
		db.Db()
	}()

	

	go func() {

		rpc.RegisterName("MessageService", rpcServer.NewMessageService(&ctx))
		listener, err := net.Listen("tcp", ":9000")
		if err != nil {
				log.Fatal("ListenTCP error: ", err)
		}
		// wg.Add(1)
		for {
				conn, err := listener.Accept()
				if err != nil {
					// wg.Done()
						log.Fatal("Accept error: ", err)
				}
				log.Printf("New connection: %+v\n", conn.RemoteAddr())
				
				go jsonrpc.ServeConn(conn)
		}
	}()
	

	// // sample test endpoint
	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(res, "RPC SERVER LIVE!")
	// })

	// // listen and serve default HTTP server
	// http.ListenAndServe(":9000", nil)
	// _chatInput := utils.MessageJsonInput{
	// 	Timestamp: 1663909754116,
	// 	From:      "111",
	// 	Receiver:  "111",
	// 	Platform:  "channel",
	// 	Type:      "html",
	// 	Message:   "hello world",
	// 	ChainId:   "",
	// 	Subject:   "Test Subject",
	// 	Signature: "909090",
	// 	Actions: []utils.ChatMessageAction{
	// 		{
	// 			Contract: "Contract",
	// 			Abi:      "Abi",
	// 			Action:   "Action",
	// 			Parameters: []string{
	// 				"good",
	// 				"Jon",
	// 				"Doe",
	// 			},
	// 		},
	// 	},
	// }
	// _chatMsg := utils.CreateMessageFromJson(_chatInput)
	// fmt.Printf("Testing my function%s, %t", "_chatMsg.ToString()", utils.IsValidMessage(_chatMsg, _chatInput.Signature))

	// r := gin.Default()
	// r = originatorRoutes.Init(r)
	// r.Run("localhost:8083")
}


// func checkError(err error) {
//     if err != nil {
//         fmt.Println("Fatal error ", err.Error())
//         os.Exit(1)
//     }
// }
