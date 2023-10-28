/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/cmd"
	config "github.com/mlayerprotocol/go-mlayer/utils"
)

func main() {
	c := config.LoadConfig()

	fmt.Printf("ChainId %s \n", c.StakeContract)
	cmd.Execute()
}
