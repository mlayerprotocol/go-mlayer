package entities

import (

	// "math"

	"github.com/btcsuite/btcd/btcec/v2"
)

type ClaimData struct {
	Cycle uint64
	Signature [32]byte
	Commitment []byte
	PubKeys []*btcec.PublicKey
	SubnetRewardCount []SubnetCount
}
