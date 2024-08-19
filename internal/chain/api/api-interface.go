package api

import (
	"math/big"

	"github.com/mlayerprotocol/go-mlayer/entities"
)

type ChainInfo struct {
	StartTime *big.Int
	StartBlock *big.Int
	CurrentBlock *big.Int
	CurrentEpoch *big.Int
	CurrentCycle *big.Int
}
type IChainAPI interface {
	// general
	GetChainInfo() (ChainInfo, error)
	GetStartTime() (*big.Int, error) 
	GetStartBlock() (*big.Int, error) 
	GetEpoch(blockNumber *big.Int) (*big.Int, error) 
	GetCycle(blockNumber *big.Int) (*big.Int, error) 
	GetCurrentCycle() (*big.Int, error) 
	GetCurrentBlockNumber() (*big.Int, error) 
	GetCurrentEpoch() (*big.Int, error) 
	GetCurrentYear() (*big.Int, error) 

	// licence
	GetSentryActiveLicenseCount(cycle *big.Int)  (*big.Int, error)
	GetValidatorActiveLicenceCount(cycle *big.Int) (*big.Int, error)
	GetSentryOperatorLicenseCount(cycle *big.Int, operator []byte)  (*big.Int, error)
	GetValidatorOperatorLicenceCount(cycle *big.Int, operator []byte) (*big.Int, error)
	GetSentryLicenceOperator(license *big.Int) ([]byte, error)
	GetValidatorLicenceOperator(license *big.Int) ([]byte, error)


	IsValidatorNodeOperator(publicKey []byte) (bool, error)
	IsSentryNodeOperator(publicKey []byte) (bool, error)
	
	
	// GetStakeBalance(address entities.DIDString) big.Int

	// subnet
	GetSubnetBalance(id [16]byte) (*big.Int, error)

	// sentry
	GetMinStakeAmountForValidators() (*big.Int, error)
	GetMinStakeAmountForSentry() (*big.Int, error)
	GetCurrentMessagePrice() (*big.Int, error)
	GetMessagePrice(cycle *big.Int) (*big.Int, error)
	// GetChannelBalance(address entities.DIDString) *big.Int
	ClaimReward(claim entities.ClaimData) (string, error) 
}


