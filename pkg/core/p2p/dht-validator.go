package p2p

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"sort"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	"github.com/mlayerprotocol/go-mlayer/internal/chain"
)



type DhtValidator struct{
    config *configs.MainConfiguration
}

type NodeMultiAddressDataIndexed struct {
    Index int
    Data NodeMultiAddressData
}

var keyPrefixes = []string{"val","cost"}

func (v *DhtValidator) Validate(key string, value []byte) error {
    if strings.Index(key, "/") != 0 {
        return errors.New("DhtValidator: key must begin with /")
    }
    parts := strings.Split(string(key), "/")

	
	if parts[1] != "ml" {
		return errors.New("DhtValidator: Invalid key prefix")
	}
    
    if len(parts) < 4 {
        return errors.New("DhtValidator: Invalid key path section length")
    }
    if !slices.Contains(keyPrefixes, parts[2]) {
		return errors.New("DhtValidator: Invalid key parts[1] value")
	}
    switch parts[2] {
    case "val":
        return v.validateValidatorListKey(parts, value)
    case "cost":
        return v.validatePriceKey(parts, value)
    }

    return nil
}

func (v *DhtValidator) Select(key string, values [][]byte) (int, error) {
    if strings.Index(key, "/") != 0 {
        return 0, nil
    }
    parts := strings.Split(string(key), "/")
    if parts[2] == "val" {
		return v.selectFromValidatorList(parts, values)
	}
    // Handle selecting the valid value among multiple
    logger.Debugf("FOUND records %d", len(values))
    return 0, nil
}

func (v *DhtValidator) validateValidatorListKey(parts []string, value []byte ) error {
    if len(parts) != 4 {
		return errors.New("DhtValidator: key parts too short or long")
	}

    if len(parts[3]) != 66 && len(parts[3]) != 64 {
		return errors.New("DhtValidator: key value must be a valid public key")
	}
    addresses, err := UnpackNodeMultiAddressData(value)
    if err != nil {
        return fmt.Errorf("DhtValidator: Invalid validator multiaddress data - %v", err)
    }
    
    if !addresses.IsValid(v.config.ChainId) {
        return errors.New("DhtValidator: Invalid validator address signature")
    }
   
   
    if parts[3] != hex.EncodeToString(addresses.Signer) && parts[3] != hex.EncodeToString(addresses.PubKeyEDD) {
        return errors.New("DhtValidator: Signer and PubKeySecp does not match key public key")
    }
    isValidator,  _ := chain.NetworkInfo.IsValidator(hex.EncodeToString(addresses.Signer))
    if !isValidator {
        return errors.New("DhtValidator: Signer is not a validator")
    }
    // if chain.HasValidStake(addresses.Signer, &v.config) {
    //     return errors.New("DhtValidator: Signer is not a validator")
    // }

	return nil
}

func (v *DhtValidator) selectFromValidatorList(parts []string, value [][]byte ) (int, error) {
    result := []NodeMultiAddressDataIndexed{}
   
    for idx, b := range value {
      
        d, err := UnpackNodeMultiAddressData(b)
        if err != nil {
            continue
        }
        if len(d.Signer) != 32 {
            continue
        }
        if !d.IsValid(cfg.ChainId) {
           continue
        }
        if parts[3] != hex.EncodeToString(d.Signer) && parts[3] != hex.EncodeToString(d.PubKeyEDD) {
            continue
        }
        result = append(result, NodeMultiAddressDataIndexed{Data: d, Index: idx})
    }
    if len(result) == 0 {
        return 0, nil
    }
    logger.Debug("DHTLEN", len(result))
    sort.Slice(result, func(i, j int) bool {
        return result[i].Data.Timestamp > result[j].Data.Timestamp
    })
	return result[0].Index, nil
}

func (v *DhtValidator) validatePriceKey(parts []string, value []byte ) error {
    if len(parts) != 4 {
		return errors.New("DhtValidator: price key parts too short or long")
	}
	if !utils.IsNumericInt(parts[3]) {
		return errors.New("DhtValidator: price key value must be a numeric")
	}
    priceData, err := UnpackMessagePrice(value)
    if err != nil {
        return fmt.Errorf("DhtValidator: Invalid price data - %v", err)
    }
    logger.Debugf("PRICE_KEY %s, %d", parts[3], new(big.Int).SetBytes(priceData.Cycle))
    if parts[3] != fmt.Sprintf("%d", new(big.Int).SetBytes(priceData.Cycle)) {
        return errors.New("DhtValidator: price data cycle does not match key cycle")
    }
    if !priceData.IsValid(cfg.ChainId) {
        return errors.New("DhtValidator: Invalid price signature")
    }
   
   
    // check if signer is validator
    // if chain.HasValidStake(addresses.Signer, &v.config) {
    //     return errors.New("DhtValidator: Signer is not a validator")
    // }

	return nil
}