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
	incomingMessagesc := make(chan utils.NodeMessage)
	outgoingMessagesc := make(chan utils.NodeMessage)

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
	ctx = context.WithValue(ctx, "OutgoinMessageC", outgoingMessagesc)
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
					logger.Errorf("Incomming Message channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				// VALIDATE AND DISTRIBUTE
				logger.Info("Received new message %s\n", inMessage.Message.Body.Text)

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
			errc <- fmt.Errorf("Db error: %g", err)
		}
		defer wg.Done()
		db.Db()
	}()

	// r := gin.Default()
	// r = originatorRoutes.Init(r)
	// r.Run("localhost:8083")
}
