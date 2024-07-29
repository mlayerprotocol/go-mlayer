package chain

import (
	"encoding/hex"
	"math/big"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
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

func (n *MLChainAPI) GetLicenseCount(cycle uint64) *big.Int {
	return big.NewInt(100)
}

func (n *MLChainAPI) GetCurrentEpoch() uint64 {
	return GetEpoch(n.GetCurrentBlockNumber())
}


func (n *MLChainAPI) GetCurrentCycle() uint64 {
	return GetCycle(n.GetCurrentBlockNumber())
}

func (n *MLChainAPI) LicenceOperator(license *big.Int) ([]byte, error) {
	
	return hex.DecodeString("02c4435e768b4bae8236eeba29dd113ed607813b4dc5419d33b9294f712ca79ff4")
}


type ClaimData struct {
	Cycle uint64
	Signature [32]byte
	Commitment []byte
	PubKeys []*btcec.PublicKey
	SubnetRewardCount []entities.SubnetCount
}
func (n *MLChainAPI) ClaimReward(claim ClaimData) (bool, error) {
	
	return true, nil
}

func (n *MLChainAPI) GetCurrentYear() uint64 {
	
	return n.GetCurrentCycle() / 12
}

func (n *MLChainAPI) Get() big.Int {
	bal := new(big.Int)
	bal.SetString("1000000000000000", 10)
	return *bal
}

func (n *MLChainAPI) GetStakeBalance(address entities.DIDString) big.Int {
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

func (n *MLChainAPI) GetCurrentMessageCost() (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("10000000", 10)
	return bal, nil
}

func (n *MLChainAPI) GetMessageCost(cycle uint64) (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("10000000", 10)
	return bal, nil
}

func (n *MLChainAPI) GetChannelBalance(address entities.DIDString) *big.Int {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
	return bal
}


