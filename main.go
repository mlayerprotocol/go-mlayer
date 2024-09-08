/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/cmd"
	"github.com/spf13/cobra"
)
var version string
var releaseDate string
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "returns the version",
	Long: `mLayer (message layer) is an open, decentralized communication network that enables the creation, 
	transmission and termination of data of all sizes, leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with the ever-advancing realm of cryptography. 
	Given its protocol-centric nature, it is an adaptable and universally integrable tool conceived for the 
	decentralized era. Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	daemonFunc(cmd, args)
	// },
}
func main() {
	if version == "" {
		version = "v1.0.0"
	}
	if releaseDate == "" {
		releaseDate = "Released: October 5th 2024 03:23 UTC"
	}
	cmd.Execute(fmt.Sprintf("%s - %s", version, releaseDate))
}
