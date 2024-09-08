/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mlayerprotocol/go-mlayer/cmd"
)
var version string
var releaseDate string

func main() {
	if version == "" {
		version = "v1.0.0"
	}
	if releaseDate == "" {
		releaseDate = "Released: October 5th 2024 03:23 UTC"
	}
	cmd.Execute(version, releaseDate)
}
