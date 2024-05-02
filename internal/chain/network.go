package chain

import (
	"math/big"
	"time"

	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/entities"
)

func GetEpoch(blockNumber uint64) uint64 {
	return blockNumber / 14400
}
func GetCycle(blockNumber uint64) uint64 {
	cycle := 1 + (GetEpoch(blockNumber) / 30)
	return cycle
}

const NetworkStartDate = uint64(1706468787000)

var API MLChainAPI

type MLChainAPI struct {
	URL string
}


func Init(cfg *configs.MainConfiguration) {
	API = *NewMLChainAPI(cfg.MLBlockchainAPIUrl)
}
func NewMLChainAPI(url string) *MLChainAPI {
	return &MLChainAPI{URL: url}
}

func (n *MLChainAPI) GetCurrentBlockNumber() uint64 {
	return (uint64(time.Now().UnixMilli()) - NetworkStartDate) / 6000
}

func (n *MLChainAPI) GetCurrentEpoch() uint64 {
	return GetEpoch(n.GetCurrentBlockNumber())
}


func (n *MLChainAPI) GetCurrentCycle() uint64 {
	return GetCycle(n.GetCurrentBlockNumber())
}


func (n *MLChainAPI) GetCurrentYear() uint64 {
	
	return n.GetCurrentCycle() / 12
}

func (n *MLChainAPI) Get() big.Int {
	bal := new(big.Int)
	bal.SetString("1000000000000000", 10)
	return *bal
}

func (n *MLChainAPI) GetStakeBalance(address entities.AddressString) big.Int {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
	return *bal
}

func (n *MLChainAPI) GetSubnetBalance(hashOrId string) big.Int {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
	return *bal
}

func (n *MLChainAPI) GetMinStakeAmountForValidators() big.Int {
	bal := new(big.Int)
	bal.SetString("1000000000000000", 10)
	return *bal
}

func (n *MLChainAPI) GetCurrentMessageCost() *big.Int {
	bal := new(big.Int)
	bal.SetString("10000000", 10)
	return bal
}

func (n *MLChainAPI) GetChannelBalance(address entities.AddressString) *big.Int {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
	return bal
}


