package evm

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/fero-tech/splanch/pkg/core/chain/evm/abis/registry"
	"github.com/fero-tech/splanch/pkg/core/chain/evm/abis/stake"
	"github.com/fero-tech/splanch/pkg/core/chain/evm/abis/token"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ToHexAddress(address string) common.Address {
	return common.HexToAddress(address)
}
func StakeContract(rpc_url string, contractAddress string) (*stake.Stake, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := stake.NewStake(address, client)
	if err != nil {
		return nil, err
	}
	return instance, err
}

func TokenContract(rpc_url string, contractAddress string) (*token.Abi, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := token.NewAbi(address, client)
	if err != nil {
		return nil, err
	}
	return instance, err
}

func RegistryContract(rpc_url string, contractAddress string) (*registry.Abi, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := registry.NewAbi(address, client)
	if err != nil {
		return nil, err
	}
	return instance, err
}

func AuthOption(privateKey string, client ethclient.Client) *bind.TransactOpts {
	pKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := pKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(pKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth
}
