package solidity

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"./abis/registry"
	"./abis/stake"
	"./abis/token"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InitStake(rpc_url string, contractAddress string) (*stake.Abi, *ethclient.Client, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := stake.NewAbi(address, client)
	if err != nil {
		return nil, nil, err
	}
	return instance, client, err
}

func InitToken(rpc_url string, contractAddress string) (*token.Abi, *ethclient.Client, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := token.NewAbi(address, client)
	if err != nil {
		return nil, nil, err
	}
	return instance, client, err
}

func InitRegistry(rpc_url string, contractAddress string) (*registry.Abi, *ethclient.Client, error) {
	client, err := ethclient.Dial(rpc_url)
	if err != nil {
		return nil, nil, err
	}

	address := common.HexToAddress(contractAddress)

	instance, err := registry.NewAbi(address, client)
	if err != nil {
		return nil, nil, err
	}
	return instance, client, err
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
