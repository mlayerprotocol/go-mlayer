package api

import (
	"encoding/hex"
	"math/big"
	"time"

	"github.com/mlayerprotocol/go-mlayer/entities"
)




type Protocol string
const (
	Ws Protocol = "ws"
	Https Protocol = "https"
)
type EthereumAPI struct {
	Protocol Protocol
	Url string
	OperatorContract string
	ValidatorContract string
	SubnetContract string
}




func (n EthereumAPI) GetEpoch(blockNumber *big.Int) (*big.Int, error) {
	return new(big.Int).Div(blockNumber , big.NewInt(14400)), nil
}
func (n EthereumAPI) GetCycle(blockNumber *big.Int) (*big.Int, error) {
	b, err := n.GetEpoch(blockNumber)
	var cycle *big.Int
	if err == nil {
		cycle = new(big.Int).Div(b, big.NewInt(30)).Add(big.NewInt(1), new(big.Int))
	}
	return cycle, err
}

func (n EthereumAPI) GetCurrentBlockNumber() (*big.Int, error) {
	return big.NewInt(int64((time.Now().UnixMilli()) - NetworkStartDate) / 6000), nil
}
func (n EthereumAPI) GetTotalSentryLicenseCount(cycle *big.Int)  (*big.Int, error)  {
	return big.NewInt(0), nil
}
func (n EthereumAPI) GetTotalValidatorLicenceCount(cycle *big.Int) (*big.Int, error) {
	return big.NewInt(0), nil
}
func (n EthereumAPI) GetSentryLicenseCount(cycle  *big.Int, operator []byte) (*big.Int, error) {
	return big.NewInt(100), nil
}
func (n EthereumAPI) GetValidatorLicenceCount(cycle  *big.Int ,operator []byte) (*big.Int, error) {
	return big.NewInt(100), nil
}
func (n EthereumAPI) GetSentryLicenceOperator(license *big.Int) ([]byte, error) {
	return hex.DecodeString("02c4435e768b4bae8236eeba29dd113ed607813b4dc5419d33b9294f712ca79ff4")
}
func (n EthereumAPI) GetValidatorLicenceOperator(license *big.Int) ([]byte, error) {
	return hex.DecodeString("02c4435e768b4bae8236eeba29dd113ed607813b4dc5419d33b9294f712ca79ff4")
}

func (n EthereumAPI) GetCurrentEpoch() (*big.Int, error) {
	b, err := n.GetCurrentBlockNumber()
	if err == nil {
		return n.GetEpoch(b)
	}
	return nil, err
}


func (n EthereumAPI) GetCurrentCycle() (*big.Int, error) {
	b, err := n.GetCurrentBlockNumber()
	if err == nil {
		return n.GetCycle(b)
	}
	return nil, err
}

func (n *EthereumAPI) LicenceOperator(license *big.Int) ([]byte, error) {
	
	return hex.DecodeString("02c4435e768b4bae8236eeba29dd113ed607813b4dc5419d33b9294f712ca79ff4")
}




func (n EthereumAPI) GetCurrentYear() (*big.Int, error) {
	 r, err := n.GetCurrentCycle()
	return new(big.Int).Div(r, big.NewInt(12)), err
}


func (n EthereumAPI) GetStakeBalance(address entities.DIDString) big.Int {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
	return *bal
}

func (n EthereumAPI) GetSubnetBalance(id [16]byte) (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
	return bal, nil
}

func (n EthereumAPI) GetMinStakeAmountForValidators() (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("1000000000000000", 10)
	return bal, nil
}

func (n EthereumAPI) GetCurrentMessagePrice() (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("10000000", 10)
	return bal, nil
}

func (n EthereumAPI) GetMessagePrice(cycle *big.Int) (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("10000000", 10)
	return bal, nil
}


func (n EthereumAPI) GetChainInfo() (ChainInfo, error) {
	curBlock, _ := n.GetCurrentBlockNumber()
	startBlock, _ := n.GetStartBlock()
 return ChainInfo{CurrentBlock: curBlock, StartBlock: startBlock }, nil
}
func (n EthereumAPI)  GetStartTime() (*big.Int, error) {
	return big.NewInt(0), nil
}
func (n EthereumAPI)  GetStartBlock() (*big.Int, error) {
	return big.NewInt(0), nil
}

func (n EthereumAPI) ClaimReward(claim entities.ClaimData) (hash string, err error) {
	
	return "", nil
}


func  (n EthereumAPI) GetMinStakeAmountForSentry() (*big.Int, error) {
	return new(big.Int), nil
}

