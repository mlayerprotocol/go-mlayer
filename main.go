/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/ByteGum/go-icms/cmd"
	config "github.com/ByteGum/go-icms/utils"
)

func main() {
	c := config.LoadConfig()

	fmt.Printf("ChainId %s \n", c.StakeContract)
	cmd.Execute()
}
