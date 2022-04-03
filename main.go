/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/fero-tech/splanch/cmd"
	config "github.com/fero-tech/splanch/utils"
)

func main() {
	c := config.LoadConfig()

	fmt.Printf("ChainId %s \n", c.StakeContract)
	cmd.Execute()
}
