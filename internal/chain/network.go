package chain

import (
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/chain/api"
)


var apis map[configs.ChainId]*api.IChainAPI = map[configs.ChainId]*api.IChainAPI{}
// var Network *MLChainAPI
// var DefaultProvider api.IChainAPI

// type MLChainAPI struct {
// 	config *configs.MainConfiguration
// 	apis map[configs.ChainId]*api.IChainAPI
// }

// func Init(config *configs.MainConfiguration) *MLChainAPI {
// 	return &MLChainAPI{config: config}
// }


// func (n MLChainAPI) Network(chainID configs.ChainId) api.IChainAPI {
// 	return *n.apis[chainID]
// }

// func (n MLChainAPI) DefaultProvider() api.IChainAPI {
// 	return *apis[n.config.ChainId]
// }

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

