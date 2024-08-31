package chain

import (
	"context"
	"encoding/hex"
	"math/big"
	"sync"
	"time"

	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/chain/api"
)
var syncMu sync.Mutex
type NetworkParams struct {
	StartTime *big.Int
	StartBlock *big.Int
	CurrentCycle *big.Int
	CurrentBlock *big.Int
	Validators map[string]bool
	Sentries map[string]uint64
	Config *configs.MainConfiguration
	Synced bool
}
func (n *NetworkParams) IsValidator(pubKey string) (bool, error) {
	return n.Validators[pubKey], nil
}
func (n *NetworkParams) IsSentry(pubKey string) (bool, error)  {
	if n.Sentries[pubKey] > 0 &&  uint64(time.Now().UnixMicro()) - n.Sentries[pubKey] < uint64(10 * time.Minute.Microseconds()) {
		return true, nil
	}
	b, err := hex.DecodeString(pubKey)
	if err != nil {
		return false, err
	}
	count, err := DefaultProvider(n.Config).GetSentryOperatorCycleLicenseCount(b, n.CurrentCycle)
	if err != nil {
		return false, err
	}
	hasLicense := count.Cmp(big.NewInt(0)) != 0
	if hasLicense {
		n.Sentries[pubKey] = uint64(time.Now().UnixMicro())
	}
	return count.Cmp(big.NewInt(0)) != 0, nil
}

func (n *NetworkParams) Sync(ctx *context.Context, syncFunc func () bool) {
		syncMu.Lock()
		defer syncMu.Unlock()
		if n.Synced {
			return 
		}
		n.Synced = syncFunc()
}
var networkInfo NetworkParams
var NetworkInfo = &networkInfo
var apis map[configs.ChainId]*api.IChainAPI = map[configs.ChainId]*api.IChainAPI{}

func DefaultProvider(cfg *configs.MainConfiguration) api.IChainAPI {
	return *apis[cfg.ChainId]
}
// func (n MLChainAPI)  RegisterNetwork(chainId configs.ChainId, api api.IChainAPI) {
// 	n.apis[chainId] = &api
// }

func RegisterProvider (chainId configs.ChainId, api api.IChainAPI) {
	apis[chainId] = &api
}
func Provider (chainId configs.ChainId) api.IChainAPI {
	return *apis[chainId]
}


// func (n MLChainAPI) GetEpoch(blockNumber uint64) uint64 {
// 	return blockNumber / 14400
// }
// func GetCycle(blockNumber uint64) uint64 {
// 	cycle := 1 + (GetEpoch(blockNumber) / 30)
// 	return cycle
// }

// func (n *MLChainAPI) GetCurrentBlockNumber() uint64 {
// 	return (uint64(time.Now().UnixMilli()) - NetworkStartDate) / 6000
// }

// func (n *MLChainAPI) GetLicenseCount(cycle uint64) *big.Int {
// 	return big.NewInt(100)
// }

// func (n *MLChainAPI) GetCurrentEpoch() uint64 {
// 	return GetEpoch(n.GetCurrentBlockNumber())
// }


// func (n *MLChainAPI) GetCurrentCycle() uint64 {
// 	return GetCycle(n.GetCurrentBlockNumber())
// }

// func (n *MLChainAPI) LicenceOperator(license *big.Int) ([]byte, error) {
	
// 	return hex.DecodeString("02c4435e768b4bae8236eeba29dd113ed607813b4dc5419d33b9294f712ca79ff4")
// }




// func (n *MLChainAPI) GetCurrentYear() uint64 {
	
// 	return n.GetCurrentCycle() / 12
// }

// func (n *MLChainAPI) Get() big.Int {
// 	bal := new(big.Int)
// 	bal.SetString("1000000000000000", 10)
// 	return *bal
// }

// func (n *MLChainAPI) GetStakeBalance(address entities.DIDString) big.Int {
// 	bal := new(big.Int)
// 	bal.SetString("100000000000000000000000000", 10)
// 	return *bal
// }

// func (n *MLChainAPI) GetSubnetBalance(hashOrId string) big.Int {
// 	bal := new(big.Int)
// 	bal.SetString("100000000000000000000000000", 10)
// 	return *bal
// }

// func (n *MLChainAPI) GetMinStakeAmountForValidators() big.Int {
// 	bal := new(big.Int)
// 	bal.SetString("1000000000000000", 10)
// 	return *bal
// }

// func (n *MLChainAPI) GetCurrentMessageCost() (*big.Int, error) {
// 	bal := new(big.Int)
// 	bal.SetString("10000000", 10)
// 	return bal, nil
// }

// func (n *MLChainAPI) GetMessageCost(cycle uint64) (*big.Int, error) {
// 	bal := new(big.Int)
// 	bal.SetString("10000000", 10)
// 	return bal, nil
// }

// func (n *MLChainAPI) GetChannelBalance(address entities.DIDString) *big.Int {
// 	bal := new(big.Int)
// 	bal.SetString("100000000000000000000000000", 10)
// 	return bal
// }

// func (n *MLChainAPI) ClaimReward(claim entities.ClaimData) (bool, error) {
	
// 	return true, nil
// }

