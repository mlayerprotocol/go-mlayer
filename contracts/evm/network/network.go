// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package network

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChainInfo is an auto generated low-level Go binding around an user-defined struct.
type ChainInfo struct {
	StartTime    *big.Int
	StartBlock   *big.Int
	CurrentBlock *big.Int
	CurrentEpoch *big.Int
	CurrentCycle *big.Int
	ChainId      *big.Int
}

// NetworkContractMetaData contains all meta data concerning the NetworkContract contract.
var NetworkContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_PRICE_MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cycleMessagePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentCycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"internalType\":\"structChainInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentMessagePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"getCurrentYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getMessagePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"priceHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"searchPriceHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setMessagePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"setStartBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NetworkContractABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkContractMetaData.ABI instead.
var NetworkContractABI = NetworkContractMetaData.ABI

// NetworkContract is an auto generated Go binding around an Ethereum contract.
type NetworkContract struct {
	NetworkContractCaller     // Read-only binding to the contract
	NetworkContractTransactor // Write-only binding to the contract
	NetworkContractFilterer   // Log filterer for contract events
}

// NetworkContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkContractSession struct {
	Contract     *NetworkContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkContractCallerSession struct {
	Contract *NetworkContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NetworkContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkContractTransactorSession struct {
	Contract     *NetworkContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NetworkContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkContractRaw struct {
	Contract *NetworkContract // Generic contract binding to access the raw methods on
}

// NetworkContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkContractCallerRaw struct {
	Contract *NetworkContractCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkContractTransactorRaw struct {
	Contract *NetworkContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkContract creates a new instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContract(address common.Address, backend bind.ContractBackend) (*NetworkContract, error) {
	contract, err := bindNetworkContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkContract{NetworkContractCaller: NetworkContractCaller{contract: contract}, NetworkContractTransactor: NetworkContractTransactor{contract: contract}, NetworkContractFilterer: NetworkContractFilterer{contract: contract}}, nil
}

// NewNetworkContractCaller creates a new read-only instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractCaller(address common.Address, caller bind.ContractCaller) (*NetworkContractCaller, error) {
	contract, err := bindNetworkContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkContractCaller{contract: contract}, nil
}

// NewNetworkContractTransactor creates a new write-only instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkContractTransactor, error) {
	contract, err := bindNetworkContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkContractTransactor{contract: contract}, nil
}

// NewNetworkContractFilterer creates a new log filterer instance of NetworkContract, bound to a specific deployed contract.
func NewNetworkContractFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkContractFilterer, error) {
	contract, err := bindNetworkContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkContractFilterer{contract: contract}, nil
}

// bindNetworkContract binds a generic wrapper to an already deployed contract.
func bindNetworkContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NetworkContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContract *NetworkContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkContract.Contract.NetworkContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContract *NetworkContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.Contract.NetworkContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContract *NetworkContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContract.Contract.NetworkContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkContract *NetworkContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkContract *NetworkContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkContract *NetworkContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NetworkContract *NetworkContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NetworkContract *NetworkContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NetworkContract.Contract.DEFAULTADMINROLE(&_NetworkContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NetworkContract *NetworkContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NetworkContract.Contract.DEFAULTADMINROLE(&_NetworkContract.CallOpts)
}

// MESSAGEPRICEMANAGER is a free data retrieval call binding the contract method 0x522872b3.
//
// Solidity: function MESSAGE_PRICE_MANAGER() view returns(bytes32)
func (_NetworkContract *NetworkContractCaller) MESSAGEPRICEMANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "MESSAGE_PRICE_MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGEPRICEMANAGER is a free data retrieval call binding the contract method 0x522872b3.
//
// Solidity: function MESSAGE_PRICE_MANAGER() view returns(bytes32)
func (_NetworkContract *NetworkContractSession) MESSAGEPRICEMANAGER() ([32]byte, error) {
	return _NetworkContract.Contract.MESSAGEPRICEMANAGER(&_NetworkContract.CallOpts)
}

// MESSAGEPRICEMANAGER is a free data retrieval call binding the contract method 0x522872b3.
//
// Solidity: function MESSAGE_PRICE_MANAGER() view returns(bytes32)
func (_NetworkContract *NetworkContractCallerSession) MESSAGEPRICEMANAGER() ([32]byte, error) {
	return _NetworkContract.Contract.MESSAGEPRICEMANAGER(&_NetworkContract.CallOpts)
}

// BlockTime is a free data retrieval call binding the contract method 0x48b15166.
//
// Solidity: function blockTime() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "blockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTime is a free data retrieval call binding the contract method 0x48b15166.
//
// Solidity: function blockTime() view returns(uint256)
func (_NetworkContract *NetworkContractSession) BlockTime() (*big.Int, error) {
	return _NetworkContract.Contract.BlockTime(&_NetworkContract.CallOpts)
}

// BlockTime is a free data retrieval call binding the contract method 0x48b15166.
//
// Solidity: function blockTime() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) BlockTime() (*big.Int, error) {
	return _NetworkContract.Contract.BlockTime(&_NetworkContract.CallOpts)
}

// CycleMessagePrice is a free data retrieval call binding the contract method 0x28240181.
//
// Solidity: function cycleMessagePrice(uint256 ) view returns(uint256)
func (_NetworkContract *NetworkContractCaller) CycleMessagePrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "cycleMessagePrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CycleMessagePrice is a free data retrieval call binding the contract method 0x28240181.
//
// Solidity: function cycleMessagePrice(uint256 ) view returns(uint256)
func (_NetworkContract *NetworkContractSession) CycleMessagePrice(arg0 *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.CycleMessagePrice(&_NetworkContract.CallOpts, arg0)
}

// CycleMessagePrice is a free data retrieval call binding the contract method 0x28240181.
//
// Solidity: function cycleMessagePrice(uint256 ) view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) CycleMessagePrice(arg0 *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.CycleMessagePrice(&_NetworkContract.CallOpts, arg0)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x21cae483.
//
// Solidity: function getChainInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_NetworkContract *NetworkContractCaller) GetChainInfo(opts *bind.CallOpts) (ChainInfo, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getChainInfo")

	if err != nil {
		return *new(ChainInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainInfo)).(*ChainInfo)

	return out0, err

}

// GetChainInfo is a free data retrieval call binding the contract method 0x21cae483.
//
// Solidity: function getChainInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_NetworkContract *NetworkContractSession) GetChainInfo() (ChainInfo, error) {
	return _NetworkContract.Contract.GetChainInfo(&_NetworkContract.CallOpts)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x21cae483.
//
// Solidity: function getChainInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_NetworkContract *NetworkContractCallerSession) GetChainInfo() (ChainInfo, error) {
	return _NetworkContract.Contract.GetChainInfo(&_NetworkContract.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetCurrentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getCurrentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentBlockNumber(&_NetworkContract.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentBlockNumber(&_NetworkContract.CallOpts)
}

// GetCurrentCycle is a free data retrieval call binding the contract method 0xbe26ed7f.
//
// Solidity: function getCurrentCycle() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetCurrentCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getCurrentCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentCycle is a free data retrieval call binding the contract method 0xbe26ed7f.
//
// Solidity: function getCurrentCycle() view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetCurrentCycle() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentCycle(&_NetworkContract.CallOpts)
}

// GetCurrentCycle is a free data retrieval call binding the contract method 0xbe26ed7f.
//
// Solidity: function getCurrentCycle() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetCurrentCycle() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentCycle(&_NetworkContract.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetCurrentEpoch() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentEpoch(&_NetworkContract.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentEpoch(&_NetworkContract.CallOpts)
}

// GetCurrentMessagePrice is a free data retrieval call binding the contract method 0x6d653312.
//
// Solidity: function getCurrentMessagePrice() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetCurrentMessagePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getCurrentMessagePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentMessagePrice is a free data retrieval call binding the contract method 0x6d653312.
//
// Solidity: function getCurrentMessagePrice() view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetCurrentMessagePrice() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentMessagePrice(&_NetworkContract.CallOpts)
}

// GetCurrentMessagePrice is a free data retrieval call binding the contract method 0x6d653312.
//
// Solidity: function getCurrentMessagePrice() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetCurrentMessagePrice() (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentMessagePrice(&_NetworkContract.CallOpts)
}

// GetCurrentYear is a free data retrieval call binding the contract method 0x318b2184.
//
// Solidity: function getCurrentYear(uint256 _block) view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetCurrentYear(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getCurrentYear", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentYear is a free data retrieval call binding the contract method 0x318b2184.
//
// Solidity: function getCurrentYear(uint256 _block) view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetCurrentYear(_block *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentYear(&_NetworkContract.CallOpts, _block)
}

// GetCurrentYear is a free data retrieval call binding the contract method 0x318b2184.
//
// Solidity: function getCurrentYear(uint256 _block) view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetCurrentYear(_block *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetCurrentYear(&_NetworkContract.CallOpts, _block)
}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 blockNumber) view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetCycle(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getCycle", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 blockNumber) view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetCycle(blockNumber *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetCycle(&_NetworkContract.CallOpts, blockNumber)
}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 blockNumber) view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetCycle(blockNumber *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetCycle(&_NetworkContract.CallOpts, blockNumber)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 blockNumber) view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetEpoch(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getEpoch", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 blockNumber) view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetEpoch(blockNumber *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetEpoch(&_NetworkContract.CallOpts, blockNumber)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 blockNumber) view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetEpoch(blockNumber *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetEpoch(&_NetworkContract.CallOpts, blockNumber)
}

// GetMessagePrice is a free data retrieval call binding the contract method 0x853ceffa.
//
// Solidity: function getMessagePrice(uint256 cycle) view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetMessagePrice(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getMessagePrice", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMessagePrice is a free data retrieval call binding the contract method 0x853ceffa.
//
// Solidity: function getMessagePrice(uint256 cycle) view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetMessagePrice(cycle *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetMessagePrice(&_NetworkContract.CallOpts, cycle)
}

// GetMessagePrice is a free data retrieval call binding the contract method 0x853ceffa.
//
// Solidity: function getMessagePrice(uint256 cycle) view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetMessagePrice(cycle *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.GetMessagePrice(&_NetworkContract.CallOpts, cycle)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NetworkContract *NetworkContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NetworkContract *NetworkContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NetworkContract.Contract.GetRoleAdmin(&_NetworkContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NetworkContract *NetworkContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NetworkContract.Contract.GetRoleAdmin(&_NetworkContract.CallOpts, role)
}

// GetStartBlock is a free data retrieval call binding the contract method 0xa5f18c01.
//
// Solidity: function getStartBlock() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartBlock is a free data retrieval call binding the contract method 0xa5f18c01.
//
// Solidity: function getStartBlock() view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetStartBlock() (*big.Int, error) {
	return _NetworkContract.Contract.GetStartBlock(&_NetworkContract.CallOpts)
}

// GetStartBlock is a free data retrieval call binding the contract method 0xa5f18c01.
//
// Solidity: function getStartBlock() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetStartBlock() (*big.Int, error) {
	return _NetworkContract.Contract.GetStartBlock(&_NetworkContract.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) GetStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "getStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() view returns(uint256)
func (_NetworkContract *NetworkContractSession) GetStartTime() (*big.Int, error) {
	return _NetworkContract.Contract.GetStartTime(&_NetworkContract.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) GetStartTime() (*big.Int, error) {
	return _NetworkContract.Contract.GetStartTime(&_NetworkContract.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NetworkContract *NetworkContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NetworkContract *NetworkContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NetworkContract.Contract.HasRole(&_NetworkContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NetworkContract *NetworkContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NetworkContract.Contract.HasRole(&_NetworkContract.CallOpts, role, account)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_NetworkContract *NetworkContractCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_NetworkContract *NetworkContractSession) Locked() (bool, error) {
	return _NetworkContract.Contract.Locked(&_NetworkContract.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_NetworkContract *NetworkContractCallerSession) Locked() (bool, error) {
	return _NetworkContract.Contract.Locked(&_NetworkContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NetworkContract *NetworkContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NetworkContract *NetworkContractSession) Owner() (common.Address, error) {
	return _NetworkContract.Contract.Owner(&_NetworkContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NetworkContract *NetworkContractCallerSession) Owner() (common.Address, error) {
	return _NetworkContract.Contract.Owner(&_NetworkContract.CallOpts)
}

// PriceHistory is a free data retrieval call binding the contract method 0xb27a0484.
//
// Solidity: function priceHistory(uint256 ) view returns(uint256 price, uint256 from, uint256 to)
func (_NetworkContract *NetworkContractCaller) PriceHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Price *big.Int
	From  *big.Int
	To    *big.Int
}, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "priceHistory", arg0)

	outstruct := new(struct {
		Price *big.Int
		From  *big.Int
		To    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.From = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.To = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PriceHistory is a free data retrieval call binding the contract method 0xb27a0484.
//
// Solidity: function priceHistory(uint256 ) view returns(uint256 price, uint256 from, uint256 to)
func (_NetworkContract *NetworkContractSession) PriceHistory(arg0 *big.Int) (struct {
	Price *big.Int
	From  *big.Int
	To    *big.Int
}, error) {
	return _NetworkContract.Contract.PriceHistory(&_NetworkContract.CallOpts, arg0)
}

// PriceHistory is a free data retrieval call binding the contract method 0xb27a0484.
//
// Solidity: function priceHistory(uint256 ) view returns(uint256 price, uint256 from, uint256 to)
func (_NetworkContract *NetworkContractCallerSession) PriceHistory(arg0 *big.Int) (struct {
	Price *big.Int
	From  *big.Int
	To    *big.Int
}, error) {
	return _NetworkContract.Contract.PriceHistory(&_NetworkContract.CallOpts, arg0)
}

// SearchPriceHistory is a free data retrieval call binding the contract method 0x80e0d2b4.
//
// Solidity: function searchPriceHistory(uint256 cycle) view returns(uint256)
func (_NetworkContract *NetworkContractCaller) SearchPriceHistory(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "searchPriceHistory", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SearchPriceHistory is a free data retrieval call binding the contract method 0x80e0d2b4.
//
// Solidity: function searchPriceHistory(uint256 cycle) view returns(uint256)
func (_NetworkContract *NetworkContractSession) SearchPriceHistory(cycle *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.SearchPriceHistory(&_NetworkContract.CallOpts, cycle)
}

// SearchPriceHistory is a free data retrieval call binding the contract method 0x80e0d2b4.
//
// Solidity: function searchPriceHistory(uint256 cycle) view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) SearchPriceHistory(cycle *big.Int) (*big.Int, error) {
	return _NetworkContract.Contract.SearchPriceHistory(&_NetworkContract.CallOpts, cycle)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_NetworkContract *NetworkContractCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_NetworkContract *NetworkContractSession) StartBlock() (*big.Int, error) {
	return _NetworkContract.Contract.StartBlock(&_NetworkContract.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_NetworkContract *NetworkContractCallerSession) StartBlock() (*big.Int, error) {
	return _NetworkContract.Contract.StartBlock(&_NetworkContract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NetworkContract *NetworkContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NetworkContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NetworkContract *NetworkContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NetworkContract.Contract.SupportsInterface(&_NetworkContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NetworkContract *NetworkContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NetworkContract.Contract.SupportsInterface(&_NetworkContract.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NetworkContract *NetworkContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NetworkContract *NetworkContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.GrantRole(&_NetworkContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NetworkContract *NetworkContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.GrantRole(&_NetworkContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _blockTime, uint256 _startBlock) returns()
func (_NetworkContract *NetworkContractTransactor) Initialize(opts *bind.TransactOpts, _blockTime *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "initialize", _blockTime, _startBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _blockTime, uint256 _startBlock) returns()
func (_NetworkContract *NetworkContractSession) Initialize(_blockTime *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.Initialize(&_NetworkContract.TransactOpts, _blockTime, _startBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xe4a30116.
//
// Solidity: function initialize(uint256 _blockTime, uint256 _startBlock) returns()
func (_NetworkContract *NetworkContractTransactorSession) Initialize(_blockTime *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.Initialize(&_NetworkContract.TransactOpts, _blockTime, _startBlock)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NetworkContract *NetworkContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NetworkContract *NetworkContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _NetworkContract.Contract.RenounceOwnership(&_NetworkContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NetworkContract *NetworkContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NetworkContract.Contract.RenounceOwnership(&_NetworkContract.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NetworkContract *NetworkContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NetworkContract *NetworkContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.RenounceRole(&_NetworkContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NetworkContract *NetworkContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.RenounceRole(&_NetworkContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NetworkContract *NetworkContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NetworkContract *NetworkContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.RevokeRole(&_NetworkContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NetworkContract *NetworkContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.RevokeRole(&_NetworkContract.TransactOpts, role, account)
}

// SetMessagePrice is a paid mutator transaction binding the contract method 0x69a4b708.
//
// Solidity: function setMessagePrice(uint256 price) returns()
func (_NetworkContract *NetworkContractTransactor) SetMessagePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setMessagePrice", price)
}

// SetMessagePrice is a paid mutator transaction binding the contract method 0x69a4b708.
//
// Solidity: function setMessagePrice(uint256 price) returns()
func (_NetworkContract *NetworkContractSession) SetMessagePrice(price *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMessagePrice(&_NetworkContract.TransactOpts, price)
}

// SetMessagePrice is a paid mutator transaction binding the contract method 0x69a4b708.
//
// Solidity: function setMessagePrice(uint256 price) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetMessagePrice(price *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetMessagePrice(&_NetworkContract.TransactOpts, price)
}

// SetStartBlock is a paid mutator transaction binding the contract method 0xf35e4a6e.
//
// Solidity: function setStartBlock(uint256 _blockNum) returns()
func (_NetworkContract *NetworkContractTransactor) SetStartBlock(opts *bind.TransactOpts, _blockNum *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "setStartBlock", _blockNum)
}

// SetStartBlock is a paid mutator transaction binding the contract method 0xf35e4a6e.
//
// Solidity: function setStartBlock(uint256 _blockNum) returns()
func (_NetworkContract *NetworkContractSession) SetStartBlock(_blockNum *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetStartBlock(&_NetworkContract.TransactOpts, _blockNum)
}

// SetStartBlock is a paid mutator transaction binding the contract method 0xf35e4a6e.
//
// Solidity: function setStartBlock(uint256 _blockNum) returns()
func (_NetworkContract *NetworkContractTransactorSession) SetStartBlock(_blockNum *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.SetStartBlock(&_NetworkContract.TransactOpts, _blockNum)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NetworkContract *NetworkContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NetworkContract *NetworkContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.TransferOwnership(&_NetworkContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NetworkContract *NetworkContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NetworkContract.Contract.TransferOwnership(&_NetworkContract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_NetworkContract *NetworkContractTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkContract.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_NetworkContract *NetworkContractSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.Withdraw(&_NetworkContract.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_NetworkContract *NetworkContractTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NetworkContract.Contract.Withdraw(&_NetworkContract.TransactOpts, token, to, amount)
}

// NetworkContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NetworkContract contract.
type NetworkContractInitializedIterator struct {
	Event *NetworkContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkContractInitialized represents a Initialized event raised by the NetworkContract contract.
type NetworkContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NetworkContract *NetworkContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*NetworkContractInitializedIterator, error) {

	logs, sub, err := _NetworkContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NetworkContractInitializedIterator{contract: _NetworkContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NetworkContract *NetworkContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NetworkContractInitialized) (event.Subscription, error) {

	logs, sub, err := _NetworkContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkContractInitialized)
				if err := _NetworkContract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NetworkContract *NetworkContractFilterer) ParseInitialized(log types.Log) (*NetworkContractInitialized, error) {
	event := new(NetworkContractInitialized)
	if err := _NetworkContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NetworkContract contract.
type NetworkContractOwnershipTransferredIterator struct {
	Event *NetworkContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkContractOwnershipTransferred represents a OwnershipTransferred event raised by the NetworkContract contract.
type NetworkContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NetworkContract *NetworkContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NetworkContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NetworkContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NetworkContractOwnershipTransferredIterator{contract: _NetworkContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NetworkContract *NetworkContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NetworkContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NetworkContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkContractOwnershipTransferred)
				if err := _NetworkContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NetworkContract *NetworkContractFilterer) ParseOwnershipTransferred(log types.Log) (*NetworkContractOwnershipTransferred, error) {
	event := new(NetworkContractOwnershipTransferred)
	if err := _NetworkContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkContractPriceUpdatedIterator is returned from FilterPriceUpdated and is used to iterate over the raw logs and unpacked data for PriceUpdated events raised by the NetworkContract contract.
type NetworkContractPriceUpdatedIterator struct {
	Event *NetworkContractPriceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkContractPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkContractPriceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkContractPriceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkContractPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkContractPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkContractPriceUpdated represents a PriceUpdated event raised by the NetworkContract contract.
type NetworkContractPriceUpdated struct {
	Cycle *big.Int
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceUpdated is a free log retrieval operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(uint256 cycle, uint256 price)
func (_NetworkContract *NetworkContractFilterer) FilterPriceUpdated(opts *bind.FilterOpts) (*NetworkContractPriceUpdatedIterator, error) {

	logs, sub, err := _NetworkContract.contract.FilterLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return &NetworkContractPriceUpdatedIterator{contract: _NetworkContract.contract, event: "PriceUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceUpdated is a free log subscription operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(uint256 cycle, uint256 price)
func (_NetworkContract *NetworkContractFilterer) WatchPriceUpdated(opts *bind.WatchOpts, sink chan<- *NetworkContractPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _NetworkContract.contract.WatchLogs(opts, "PriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkContractPriceUpdated)
				if err := _NetworkContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceUpdated is a log parse operation binding the contract event 0x945c1c4e99aa89f648fbfe3df471b916f719e16d960fcec0737d4d56bd696838.
//
// Solidity: event PriceUpdated(uint256 cycle, uint256 price)
func (_NetworkContract *NetworkContractFilterer) ParsePriceUpdated(log types.Log) (*NetworkContractPriceUpdated, error) {
	event := new(NetworkContractPriceUpdated)
	if err := _NetworkContract.contract.UnpackLog(event, "PriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NetworkContract contract.
type NetworkContractRoleAdminChangedIterator struct {
	Event *NetworkContractRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkContractRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkContractRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkContractRoleAdminChanged represents a RoleAdminChanged event raised by the NetworkContract contract.
type NetworkContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NetworkContract *NetworkContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NetworkContractRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NetworkContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NetworkContractRoleAdminChangedIterator{contract: _NetworkContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NetworkContract *NetworkContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NetworkContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NetworkContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkContractRoleAdminChanged)
				if err := _NetworkContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NetworkContract *NetworkContractFilterer) ParseRoleAdminChanged(log types.Log) (*NetworkContractRoleAdminChanged, error) {
	event := new(NetworkContractRoleAdminChanged)
	if err := _NetworkContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NetworkContract contract.
type NetworkContractRoleGrantedIterator struct {
	Event *NetworkContractRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkContractRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkContractRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkContractRoleGranted represents a RoleGranted event raised by the NetworkContract contract.
type NetworkContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkContract *NetworkContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NetworkContractRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NetworkContractRoleGrantedIterator{contract: _NetworkContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkContract *NetworkContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NetworkContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkContractRoleGranted)
				if err := _NetworkContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkContract *NetworkContractFilterer) ParseRoleGranted(log types.Log) (*NetworkContractRoleGranted, error) {
	event := new(NetworkContractRoleGranted)
	if err := _NetworkContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NetworkContract contract.
type NetworkContractRoleRevokedIterator struct {
	Event *NetworkContractRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NetworkContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkContractRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NetworkContractRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NetworkContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkContractRoleRevoked represents a RoleRevoked event raised by the NetworkContract contract.
type NetworkContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkContract *NetworkContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NetworkContractRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NetworkContractRoleRevokedIterator{contract: _NetworkContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkContract *NetworkContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NetworkContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkContractRoleRevoked)
				if err := _NetworkContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkContract *NetworkContractFilterer) ParseRoleRevoked(log types.Log) (*NetworkContractRoleRevoked, error) {
	event := new(NetworkContractRoleRevoked)
	if err := _NetworkContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
