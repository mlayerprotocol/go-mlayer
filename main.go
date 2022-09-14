/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/fero-tech/icms/cmd"
	config "github.com/fero-tech/icms/utils"
)

func main() {
	c := config.LoadConfig()

	fmt.Printf("ChainId %s \n", c.StakeContract)
	cmd.Execute()
}
