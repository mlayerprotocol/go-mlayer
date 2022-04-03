/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sync"

	"github.com/fero-tech/splanch/pkg/core/db"
	p2p "github.com/fero-tech/splanch/pkg/core/p2p"
	originatorRoutes "github.com/fero-tech/splanch/pkg/core/rest/originator"
	utils "github.com/fero-tech/splanch/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

const (
	TESTNET string = "/icm/testing"
	MAINNET        = "/icm/mainnet"
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
		daemonFunc()
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
}

func daemonFunc() {
	c := utils.LoadConfig()
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

	go func(priv string) {
		wg.Add(1)
		if err := recover(); err != nil {
			errc <- fmt.Errorf("P2P error: %g", err)
		}
		fmt.Printf("publicKey %s", priv)
		p2p.Run(MAINNET, priv)
	}(c.PrivateKey)
	go func() {
		wg.Add(1)
		if err := recover(); err != nil {
			errc <- fmt.Errorf("Db error: %g", err)
		}
		db.Db()
	}()

	r := gin.Default()
	r = originatorRoutes.Init(r)
	r.Run("localhost:8084")
}
