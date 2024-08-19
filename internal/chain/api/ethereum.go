package api

import (
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/contracts/evm/sentry"
	"github.com/mlayerprotocol/go-mlayer/entities"
)




type Protocol string
const (
	Ws Protocol = "ws"
	Https Protocol = "https"
)
type EthereumAPI struct {
	// Protocol Protocol
	EthConfig configs.EthConfig
	sentryContract *sentry.SentryContract
	chainId *configs.ChainId
	client *ethclient.Client
}


func NewEthAPI(chainId configs.ChainId, ethConfig configs.EthConfig) (*EthereumAPI, error) {
	api := EthereumAPI{EthConfig: ethConfig}
	client, err := ethclient.Dial(ethConfig.Http)
    if err != nil {
        return nil, err
    }
	api.client = client
	api.sentryContract, err = sentry.NewSentryContract(common.HexToAddress(ethConfig.SentryNodeContract), api.client)
	if err != nil {
		return nil, err
	}
	return &api, nil
}

func (n EthereumAPI) GetEpoch(blockNumber *big.Int) (*big.Int, error) {
	return n.sentryContract.GetEpoch( nil,blockNumber)
	// return new(big.Int).Div(blockNumber , big.NewInt(14400)), nil
}
func (n EthereumAPI) GetCycle(blockNumber *big.Int) (*big.Int, error) {
	return n.sentryContract.GetCycle( nil,blockNumber)
}

func (n EthereumAPI) GetCurrentBlockNumber() (*big.Int, error) {
	return n.sentryContract.GetCurrentBlockNumber( nil)
}
func (n EthereumAPI) GetSentryActiveLicenseCount(cycle *big.Int)  (*big.Int, error)  {
	return n.sentryContract.ActiveLicenseCount( nil)
}
func (n EthereumAPI) GetValidatorActiveLicenceCount(cycle *big.Int) (*big.Int, error) {
	return big.NewInt(2), nil
}
func (n EthereumAPI) GetSentryOperatorLicenseCount(operator []byte, cycle  *big.Int) (*big.Int, error) {
	return n.sentryContract.OperatorCycleLicenseCount( nil, operator, cycle)
}
func (n EthereumAPI) GetValidatorOperatorLicenceCount(cycle  *big.Int ,operator []byte) (*big.Int, error) {
	return big.NewInt(100), nil
}
func (n EthereumAPI) GetSentryLicenceOperator(license *big.Int) ([]byte, error) {
	return n.sentryContract.LicenseOperator(nil, license)
}
func (n EthereumAPI) GetValidatorLicenceOperator(license *big.Int) ([]byte, error) {
	return hex.DecodeString("02c4435e768b4bae8236eeba29dd113ed607813b4dc5419d33b9294f712ca79ff4")
}

func (n EthereumAPI) IsValidatorNodeOperator(publicKey []byte) (bool, error) {
	return true, nil
}

func (n EthereumAPI) IsSentryNodeOperator(publicKey []byte, cycle *big.Int) (bool, error) {
	licenseCount, err := n.sentryContract.OperatorCycleLicenseCount(nil, publicKey, cycle)
	if err != nil {
		return false, err
	}
	if licenseCount.Cmp(big.NewInt(0)) == 1 {
		return false, nil
	}
	return true, nil
}

func (n EthereumAPI) GetCurrentEpoch() (*big.Int, error) {
	return n.sentryContract.GetCurrentEpoch(nil)
}


func (n EthereumAPI) GetCurrentCycle() (*big.Int, error) {
	return n.sentryContract.GetCurrentCycle()
}

func (n *EthereumAPI) LicenceOperator(license *big.Int) ([]byte, error) {
	return n.sentryContract.LicenseOperator(nil, license)
}


func (n EthereumAPI) GetCurrentYear() (*big.Int, error) {
	 r, err := n.sentryContract.StartTime(nil)
	return new(big.Int).Div(r, big.NewInt(int64((((time.Hour)*24*365)+(time.Hour*6))/time.Millisecond))), err
}


func (n EthereumAPI) GetSubnetBalance(id [16]byte) (*big.Int, error) {
	bal := new(big.Int)
	bal.SetString("100000000000000000000000000", 10)
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

