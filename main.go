/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/cmd"
	"github.com/mlayerprotocol/go-mlayer/configs"
)

func main() {
	c := configs.LoadMainConfig()

	fmt.Printf("ChainId %s \n", c.StakeContract)
	cmd.Execute()
}
