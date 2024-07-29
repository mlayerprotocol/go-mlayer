package chain

import (
	"fmt"
	"math/big"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/pkg/core/chain/evm"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/sirupsen/logrus"
)

var logger = &log.Logger

func  HasValidStake(address string, cfg *configs.MainConfiguration) bool {
	
		stakeContract, _, _, err := evm.StakeContract(cfg.EVMRPCHttp, cfg.StakeContract)
		if err != nil {
			logger.Errorf("EVM RPC error. Could not connect to validator contract: %s", err)
			return false
		}

		level, err := stakeContract.GetNodeLevel(nil, evm.ToHexAddress(address))
		i := new(big.Int).SetUint64(uint64(constants.ValidatorNodeType))
		fmt.Printf("level i ---  %s: %s -- %s\n", level, i, err)
		if level == nil || level.Cmp(i) >= 0 {
			logger.WithFields(logrus.Fields{"address": address, "accountType": level}).Infof("Inadequate stake balance for validator peer %v", err)
			return false
		}
	return true
}