// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake

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
)

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LAUNCH_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLevel\",\"outputs\":[{\"internalType\":\"enumStake.AccountLevels\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStakeForOriginators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStakeForRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"initiateUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeForOriginators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeForRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingStakeDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountNoDecimals\",\"type\":\"uint256\"}],\"name\":\"setMinStakeForOriginators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountNoDecimals\",\"type\":\"uint256\"}],\"name\":\"setMinStakeForRelayers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dayz\",\"type\":\"uint256\"}],\"name\":\"setPendingDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingIndex\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// LAUNCHDAY is a free data retrieval call binding the contract method 0x734c2cff.
//
// Solidity: function LAUNCH_DAY() view returns(uint256)
func (_Stake *StakeCaller) LAUNCHDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "LAUNCH_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LAUNCHDAY is a free data retrieval call binding the contract method 0x734c2cff.
//
// Solidity: function LAUNCH_DAY() view returns(uint256)
func (_Stake *StakeSession) LAUNCHDAY() (*big.Int, error) {
	return _Stake.Contract.LAUNCHDAY(&_Stake.CallOpts)
}

// LAUNCHDAY is a free data retrieval call binding the contract method 0x734c2cff.
//
// Solidity: function LAUNCH_DAY() view returns(uint256)
func (_Stake *StakeCallerSession) LAUNCHDAY() (*big.Int, error) {
	return _Stake.Contract.LAUNCHDAY(&_Stake.CallOpts)
}

// GetAccountLevel is a free data retrieval call binding the contract method 0x189048fc.
//
// Solidity: function getAccountLevel(address account) view returns(uint8)
func (_Stake *StakeCaller) GetAccountLevel(opts *bind.CallOpts, account common.Address) (uint8, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getAccountLevel", account)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAccountLevel is a free data retrieval call binding the contract method 0x189048fc.
//
// Solidity: function getAccountLevel(address account) view returns(uint8)
func (_Stake *StakeSession) GetAccountLevel(account common.Address) (uint8, error) {
	return _Stake.Contract.GetAccountLevel(&_Stake.CallOpts, account)
}

// GetAccountLevel is a free data retrieval call binding the contract method 0x189048fc.
//
// Solidity: function getAccountLevel(address account) view returns(uint8)
func (_Stake *StakeCallerSession) GetAccountLevel(account common.Address) (uint8, error) {
	return _Stake.Contract.GetAccountLevel(&_Stake.CallOpts, account)
}

// GetMinStakeForOriginators is a free data retrieval call binding the contract method 0xed1d5864.
//
// Solidity: function getMinStakeForOriginators() view returns(uint256)
func (_Stake *StakeCaller) GetMinStakeForOriginators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getMinStakeForOriginators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinStakeForOriginators is a free data retrieval call binding the contract method 0xed1d5864.
//
// Solidity: function getMinStakeForOriginators() view returns(uint256)
func (_Stake *StakeSession) GetMinStakeForOriginators() (*big.Int, error) {
	return _Stake.Contract.GetMinStakeForOriginators(&_Stake.CallOpts)
}

// GetMinStakeForOriginators is a free data retrieval call binding the contract method 0xed1d5864.
//
// Solidity: function getMinStakeForOriginators() view returns(uint256)
func (_Stake *StakeCallerSession) GetMinStakeForOriginators() (*big.Int, error) {
	return _Stake.Contract.GetMinStakeForOriginators(&_Stake.CallOpts)
}

// GetMinStakeForRelayers is a free data retrieval call binding the contract method 0x04685b4f.
//
// Solidity: function getMinStakeForRelayers() view returns(uint256)
func (_Stake *StakeCaller) GetMinStakeForRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getMinStakeForRelayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinStakeForRelayers is a free data retrieval call binding the contract method 0x04685b4f.
//
// Solidity: function getMinStakeForRelayers() view returns(uint256)
func (_Stake *StakeSession) GetMinStakeForRelayers() (*big.Int, error) {
	return _Stake.Contract.GetMinStakeForRelayers(&_Stake.CallOpts)
}

// GetMinStakeForRelayers is a free data retrieval call binding the contract method 0x04685b4f.
//
// Solidity: function getMinStakeForRelayers() view returns(uint256)
func (_Stake *StakeCallerSession) GetMinStakeForRelayers() (*big.Int, error) {
	return _Stake.Contract.GetMinStakeForRelayers(&_Stake.CallOpts)
}

// MinimumStakeForOriginators is a free data retrieval call binding the contract method 0x9d9f3368.
//
// Solidity: function minimumStakeForOriginators() view returns(uint256)
func (_Stake *StakeCaller) MinimumStakeForOriginators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "minimumStakeForOriginators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeForOriginators is a free data retrieval call binding the contract method 0x9d9f3368.
//
// Solidity: function minimumStakeForOriginators() view returns(uint256)
func (_Stake *StakeSession) MinimumStakeForOriginators() (*big.Int, error) {
	return _Stake.Contract.MinimumStakeForOriginators(&_Stake.CallOpts)
}

// MinimumStakeForOriginators is a free data retrieval call binding the contract method 0x9d9f3368.
//
// Solidity: function minimumStakeForOriginators() view returns(uint256)
func (_Stake *StakeCallerSession) MinimumStakeForOriginators() (*big.Int, error) {
	return _Stake.Contract.MinimumStakeForOriginators(&_Stake.CallOpts)
}

// MinimumStakeForRelayers is a free data retrieval call binding the contract method 0x1542b43c.
//
// Solidity: function minimumStakeForRelayers() view returns(uint256)
func (_Stake *StakeCaller) MinimumStakeForRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "minimumStakeForRelayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeForRelayers is a free data retrieval call binding the contract method 0x1542b43c.
//
// Solidity: function minimumStakeForRelayers() view returns(uint256)
func (_Stake *StakeSession) MinimumStakeForRelayers() (*big.Int, error) {
	return _Stake.Contract.MinimumStakeForRelayers(&_Stake.CallOpts)
}

// MinimumStakeForRelayers is a free data retrieval call binding the contract method 0x1542b43c.
//
// Solidity: function minimumStakeForRelayers() view returns(uint256)
func (_Stake *StakeCallerSession) MinimumStakeForRelayers() (*big.Int, error) {
	return _Stake.Contract.MinimumStakeForRelayers(&_Stake.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Stake *StakeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Stake *StakeSession) Operators(arg0 common.Address) (bool, error) {
	return _Stake.Contract.Operators(&_Stake.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(bool)
func (_Stake *StakeCallerSession) Operators(arg0 common.Address) (bool, error) {
	return _Stake.Contract.Operators(&_Stake.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// PendingStakeDuration is a free data retrieval call binding the contract method 0xd3a219ff.
//
// Solidity: function pendingStakeDuration() view returns(uint256)
func (_Stake *StakeCaller) PendingStakeDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "pendingStakeDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingStakeDuration is a free data retrieval call binding the contract method 0xd3a219ff.
//
// Solidity: function pendingStakeDuration() view returns(uint256)
func (_Stake *StakeSession) PendingStakeDuration() (*big.Int, error) {
	return _Stake.Contract.PendingStakeDuration(&_Stake.CallOpts)
}

// PendingStakeDuration is a free data retrieval call binding the contract method 0xd3a219ff.
//
// Solidity: function pendingStakeDuration() view returns(uint256)
func (_Stake *StakeCallerSession) PendingStakeDuration() (*big.Int, error) {
	return _Stake.Contract.PendingStakeDuration(&_Stake.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_Stake *StakeCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_Stake *StakeSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.Stakes(&_Stake.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_Stake *StakeCallerSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _Stake.Contract.Stakes(&_Stake.CallOpts, arg0)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 amount) returns()
func (_Stake *StakeTransactor) InitiateUnstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "initiateUnstake", amount)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 amount) returns()
func (_Stake *StakeSession) InitiateUnstake(amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.InitiateUnstake(&_Stake.TransactOpts, amount)
}

// InitiateUnstake is a paid mutator transaction binding the contract method 0xae5ac921.
//
// Solidity: function initiateUnstake(uint256 amount) returns()
func (_Stake *StakeTransactorSession) InitiateUnstake(amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.InitiateUnstake(&_Stake.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// SetMinStakeForOriginators is a paid mutator transaction binding the contract method 0xc3273aaf.
//
// Solidity: function setMinStakeForOriginators(uint256 amountNoDecimals) returns()
func (_Stake *StakeTransactor) SetMinStakeForOriginators(opts *bind.TransactOpts, amountNoDecimals *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setMinStakeForOriginators", amountNoDecimals)
}

// SetMinStakeForOriginators is a paid mutator transaction binding the contract method 0xc3273aaf.
//
// Solidity: function setMinStakeForOriginators(uint256 amountNoDecimals) returns()
func (_Stake *StakeSession) SetMinStakeForOriginators(amountNoDecimals *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetMinStakeForOriginators(&_Stake.TransactOpts, amountNoDecimals)
}

// SetMinStakeForOriginators is a paid mutator transaction binding the contract method 0xc3273aaf.
//
// Solidity: function setMinStakeForOriginators(uint256 amountNoDecimals) returns()
func (_Stake *StakeTransactorSession) SetMinStakeForOriginators(amountNoDecimals *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetMinStakeForOriginators(&_Stake.TransactOpts, amountNoDecimals)
}

// SetMinStakeForRelayers is a paid mutator transaction binding the contract method 0xe68beeac.
//
// Solidity: function setMinStakeForRelayers(uint256 amountNoDecimals) returns()
func (_Stake *StakeTransactor) SetMinStakeForRelayers(opts *bind.TransactOpts, amountNoDecimals *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setMinStakeForRelayers", amountNoDecimals)
}

// SetMinStakeForRelayers is a paid mutator transaction binding the contract method 0xe68beeac.
//
// Solidity: function setMinStakeForRelayers(uint256 amountNoDecimals) returns()
func (_Stake *StakeSession) SetMinStakeForRelayers(amountNoDecimals *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetMinStakeForRelayers(&_Stake.TransactOpts, amountNoDecimals)
}

// SetMinStakeForRelayers is a paid mutator transaction binding the contract method 0xe68beeac.
//
// Solidity: function setMinStakeForRelayers(uint256 amountNoDecimals) returns()
func (_Stake *StakeTransactorSession) SetMinStakeForRelayers(amountNoDecimals *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetMinStakeForRelayers(&_Stake.TransactOpts, amountNoDecimals)
}

// SetPendingDuration is a paid mutator transaction binding the contract method 0xccbebbd3.
//
// Solidity: function setPendingDuration(uint256 dayz) returns()
func (_Stake *StakeTransactor) SetPendingDuration(opts *bind.TransactOpts, dayz *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setPendingDuration", dayz)
}

// SetPendingDuration is a paid mutator transaction binding the contract method 0xccbebbd3.
//
// Solidity: function setPendingDuration(uint256 dayz) returns()
func (_Stake *StakeSession) SetPendingDuration(dayz *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetPendingDuration(&_Stake.TransactOpts, dayz)
}

// SetPendingDuration is a paid mutator transaction binding the contract method 0xccbebbd3.
//
// Solidity: function setPendingDuration(uint256 dayz) returns()
func (_Stake *StakeTransactorSession) SetPendingDuration(dayz *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SetPendingDuration(&_Stake.TransactOpts, dayz)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stake *StakeTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stake *StakeSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Stake *StakeTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 pendingIndex) returns()
func (_Stake *StakeTransactor) Unstake(opts *bind.TransactOpts, pendingIndex *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "unstake", pendingIndex)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 pendingIndex) returns()
func (_Stake *StakeSession) Unstake(pendingIndex *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Unstake(&_Stake.TransactOpts, pendingIndex)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 pendingIndex) returns()
func (_Stake *StakeTransactorSession) Unstake(pendingIndex *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Unstake(&_Stake.TransactOpts, pendingIndex)
}

// StakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stake contract.
type StakeOwnershipTransferredIterator struct {
	Event *StakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeOwnershipTransferred)
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
		it.Event = new(StakeOwnershipTransferred)
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
func (it *StakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeOwnershipTransferred represents a OwnershipTransferred event raised by the Stake contract.
type StakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeOwnershipTransferredIterator{contract: _Stake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeOwnershipTransferred)
				if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stake *StakeFilterer) ParseOwnershipTransferred(log types.Log) (*StakeOwnershipTransferred, error) {
	event := new(StakeOwnershipTransferred)
	if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stake contract.
type StakeStakedIterator struct {
	Event *StakeStaked // Event containing the contract specifics and raw log

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
func (it *StakeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStaked)
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
		it.Event = new(StakeStaked)
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
func (it *StakeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStaked represents a Staked event raised by the Stake contract.
type StakeStaked struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 time)
func (_Stake *StakeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StakeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakedIterator{contract: _Stake.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 time)
func (_Stake *StakeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakeStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStaked)
				if err := _Stake.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 time)
func (_Stake *StakeFilterer) ParseStaked(log types.Log) (*StakeStaked, error) {
	event := new(StakeStaked)
	if err := _Stake.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Stake contract.
type StakeUnstakedIterator struct {
	Event *StakeUnstaked // Event containing the contract specifics and raw log

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
func (it *StakeUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeUnstaked)
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
		it.Event = new(StakeUnstaked)
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
func (it *StakeUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeUnstaked represents a Unstaked event raised by the Stake contract.
type StakeUnstaked struct {
	User   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 time)
func (_Stake *StakeFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*StakeUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeUnstakedIterator{contract: _Stake.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 time)
func (_Stake *StakeFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakeUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeUnstaked)
				if err := _Stake.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 amount, uint256 time)
func (_Stake *StakeFilterer) ParseUnstaked(log types.Log) (*StakeUnstaked, error) {
	event := new(StakeUnstaked)
	if err := _Stake.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
