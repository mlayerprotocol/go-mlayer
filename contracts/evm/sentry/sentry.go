// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sentry

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

// RegistrationData is an auto generated low-level Go binding around an user-defined struct.
type RegistrationData struct {
	PublicKey  []byte
	Nonce      *big.Int
	Signature  []byte
	Commitment common.Address
}

// SentryContractMetaData contains all meta data concerning the SentryContract contract.
var SentryContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PurchaseEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeLicensesIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calibrator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"deRegisterNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillLicenseCountGap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getCycleActiveLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLicencePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getOperatorCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"}],\"name\":\"getOperatorLicenses\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perPage\",\"type\":\"uint256\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"opr\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getRegistrationData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"}],\"internalType\":\"structRegistrationData\",\"name\":\"regData\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"licensePrice\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseOperator\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"contractINetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesOwned\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"operatorLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"operatorsOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"purchaseLicense\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"}],\"internalType\":\"structRegistrationData\",\"name\":\"regData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"registerNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"regDataBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_calibrator\",\"type\":\"uint256\"}],\"name\":\"setCalibrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setInitialLicencePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SentryContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SentryContractMetaData.ABI instead.
var SentryContractABI = SentryContractMetaData.ABI

// SentryContract is an auto generated Go binding around an Ethereum contract.
type SentryContract struct {
	SentryContractCaller     // Read-only binding to the contract
	SentryContractTransactor // Write-only binding to the contract
	SentryContractFilterer   // Log filterer for contract events
}

// SentryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SentryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SentryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SentryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SentryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SentryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SentryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SentryContractSession struct {
	Contract     *SentryContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SentryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SentryContractCallerSession struct {
	Contract *SentryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SentryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SentryContractTransactorSession struct {
	Contract     *SentryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SentryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SentryContractRaw struct {
	Contract *SentryContract // Generic contract binding to access the raw methods on
}

// SentryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SentryContractCallerRaw struct {
	Contract *SentryContractCaller // Generic read-only contract binding to access the raw methods on
}

// SentryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SentryContractTransactorRaw struct {
	Contract *SentryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSentryContract creates a new instance of SentryContract, bound to a specific deployed contract.
func NewSentryContract(address common.Address, backend bind.ContractBackend) (*SentryContract, error) {
	contract, err := bindSentryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SentryContract{SentryContractCaller: SentryContractCaller{contract: contract}, SentryContractTransactor: SentryContractTransactor{contract: contract}, SentryContractFilterer: SentryContractFilterer{contract: contract}}, nil
}

// NewSentryContractCaller creates a new read-only instance of SentryContract, bound to a specific deployed contract.
func NewSentryContractCaller(address common.Address, caller bind.ContractCaller) (*SentryContractCaller, error) {
	contract, err := bindSentryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SentryContractCaller{contract: contract}, nil
}

// NewSentryContractTransactor creates a new write-only instance of SentryContract, bound to a specific deployed contract.
func NewSentryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SentryContractTransactor, error) {
	contract, err := bindSentryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SentryContractTransactor{contract: contract}, nil
}

// NewSentryContractFilterer creates a new log filterer instance of SentryContract, bound to a specific deployed contract.
func NewSentryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SentryContractFilterer, error) {
	contract, err := bindSentryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SentryContractFilterer{contract: contract}, nil
}

// bindSentryContract binds a generic wrapper to an already deployed contract.
func bindSentryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SentryContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SentryContract *SentryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SentryContract.Contract.SentryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SentryContract *SentryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentryContract.Contract.SentryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SentryContract *SentryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SentryContract.Contract.SentryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SentryContract *SentryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SentryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SentryContract *SentryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SentryContract *SentryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SentryContract.Contract.contract.Transact(opts, method, params...)
}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_SentryContract *SentryContractCaller) AccountLicenseCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "accountLicenseCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_SentryContract *SentryContractSession) AccountLicenseCount(arg0 common.Address) (*big.Int, error) {
	return _SentryContract.Contract.AccountLicenseCount(&_SentryContract.CallOpts, arg0)
}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) AccountLicenseCount(arg0 common.Address) (*big.Int, error) {
	return _SentryContract.Contract.AccountLicenseCount(&_SentryContract.CallOpts, arg0)
}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCaller) AccountLicenses(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "accountLicenses", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractSession) AccountLicenses(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.AccountLicenses(&_SentryContract.CallOpts, arg0, arg1)
}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) AccountLicenses(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.AccountLicenses(&_SentryContract.CallOpts, arg0, arg1)
}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_SentryContract *SentryContractCaller) ActiveLicenseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "activeLicenseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_SentryContract *SentryContractSession) ActiveLicenseCount() (*big.Int, error) {
	return _SentryContract.Contract.ActiveLicenseCount(&_SentryContract.CallOpts)
}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_SentryContract *SentryContractCallerSession) ActiveLicenseCount() (*big.Int, error) {
	return _SentryContract.Contract.ActiveLicenseCount(&_SentryContract.CallOpts)
}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCaller) ActiveLicensesIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "activeLicensesIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_SentryContract *SentryContractSession) ActiveLicensesIndex(arg0 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.ActiveLicensesIndex(&_SentryContract.CallOpts, arg0)
}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) ActiveLicensesIndex(arg0 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.ActiveLicensesIndex(&_SentryContract.CallOpts, arg0)
}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_SentryContract *SentryContractCaller) Calibrator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "calibrator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_SentryContract *SentryContractSession) Calibrator() (*big.Int, error) {
	return _SentryContract.Contract.Calibrator(&_SentryContract.CallOpts)
}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_SentryContract *SentryContractCallerSession) Calibrator() (*big.Int, error) {
	return _SentryContract.Contract.Calibrator(&_SentryContract.CallOpts)
}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractCaller) GetCycleActiveLicenseCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getCycleActiveLicenseCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractSession) GetCycleActiveLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.GetCycleActiveLicenseCount(&_SentryContract.CallOpts, cycle)
}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) GetCycleActiveLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.GetCycleActiveLicenseCount(&_SentryContract.CallOpts, cycle)
}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractCaller) GetCycleLicenseCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getCycleLicenseCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractSession) GetCycleLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.GetCycleLicenseCount(&_SentryContract.CallOpts, cycle)
}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) GetCycleLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.GetCycleLicenseCount(&_SentryContract.CallOpts, cycle)
}

// GetLicencePrice is a free data retrieval call binding the contract method 0x1c9129a9.
//
// Solidity: function getLicencePrice(address token) view returns(uint256)
func (_SentryContract *SentryContractCaller) GetLicencePrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getLicencePrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLicencePrice is a free data retrieval call binding the contract method 0x1c9129a9.
//
// Solidity: function getLicencePrice(address token) view returns(uint256)
func (_SentryContract *SentryContractSession) GetLicencePrice(token common.Address) (*big.Int, error) {
	return _SentryContract.Contract.GetLicencePrice(&_SentryContract.CallOpts, token)
}

// GetLicencePrice is a free data retrieval call binding the contract method 0x1c9129a9.
//
// Solidity: function getLicencePrice(address token) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) GetLicencePrice(token common.Address) (*big.Int, error) {
	return _SentryContract.Contract.GetLicencePrice(&_SentryContract.CallOpts, token)
}

// GetOperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x333825bd.
//
// Solidity: function getOperatorCycleLicenseCount(bytes operator, uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractCaller) GetOperatorCycleLicenseCount(opts *bind.CallOpts, operator []byte, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getOperatorCycleLicenseCount", operator, cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x333825bd.
//
// Solidity: function getOperatorCycleLicenseCount(bytes operator, uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractSession) GetOperatorCycleLicenseCount(operator []byte, cycle *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.GetOperatorCycleLicenseCount(&_SentryContract.CallOpts, operator, cycle)
}

// GetOperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x333825bd.
//
// Solidity: function getOperatorCycleLicenseCount(bytes operator, uint256 cycle) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) GetOperatorCycleLicenseCount(operator []byte, cycle *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.GetOperatorCycleLicenseCount(&_SentryContract.CallOpts, operator, cycle)
}

// GetOperatorLicenses is a free data retrieval call binding the contract method 0xb6e1121f.
//
// Solidity: function getOperatorLicenses(bytes operator) view returns(uint256[])
func (_SentryContract *SentryContractCaller) GetOperatorLicenses(opts *bind.CallOpts, operator []byte) ([]*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getOperatorLicenses", operator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorLicenses is a free data retrieval call binding the contract method 0xb6e1121f.
//
// Solidity: function getOperatorLicenses(bytes operator) view returns(uint256[])
func (_SentryContract *SentryContractSession) GetOperatorLicenses(operator []byte) ([]*big.Int, error) {
	return _SentryContract.Contract.GetOperatorLicenses(&_SentryContract.CallOpts, operator)
}

// GetOperatorLicenses is a free data retrieval call binding the contract method 0xb6e1121f.
//
// Solidity: function getOperatorLicenses(bytes operator) view returns(uint256[])
func (_SentryContract *SentryContractCallerSession) GetOperatorLicenses(operator []byte) ([]*big.Int, error) {
	return _SentryContract.Contract.GetOperatorLicenses(&_SentryContract.CallOpts, operator)
}

// GetOperators is a free data retrieval call binding the contract method 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 page, uint256 perPage) view returns(bytes[] opr)
func (_SentryContract *SentryContractCaller) GetOperators(opts *bind.CallOpts, page *big.Int, perPage *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getOperators", page, perPage)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 page, uint256 perPage) view returns(bytes[] opr)
func (_SentryContract *SentryContractSession) GetOperators(page *big.Int, perPage *big.Int) ([][]byte, error) {
	return _SentryContract.Contract.GetOperators(&_SentryContract.CallOpts, page, perPage)
}

// GetOperators is a free data retrieval call binding the contract method 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 page, uint256 perPage) view returns(bytes[] opr)
func (_SentryContract *SentryContractCallerSession) GetOperators(page *big.Int, perPage *big.Int) ([][]byte, error) {
	return _SentryContract.Contract.GetOperators(&_SentryContract.CallOpts, page, perPage)
}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_SentryContract *SentryContractCaller) GetRegistrationData(opts *bind.CallOpts, data []byte) (RegistrationData, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "getRegistrationData", data)

	if err != nil {
		return *new(RegistrationData), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistrationData)).(*RegistrationData)

	return out0, err

}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_SentryContract *SentryContractSession) GetRegistrationData(data []byte) (RegistrationData, error) {
	return _SentryContract.Contract.GetRegistrationData(&_SentryContract.CallOpts, data)
}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_SentryContract *SentryContractCallerSession) GetRegistrationData(data []byte) (RegistrationData, error) {
	return _SentryContract.Contract.GetRegistrationData(&_SentryContract.CallOpts, data)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_SentryContract *SentryContractCaller) IsActive(opts *bind.CallOpts, license *big.Int) (bool, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "isActive", license)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_SentryContract *SentryContractSession) IsActive(license *big.Int) (bool, error) {
	return _SentryContract.Contract.IsActive(&_SentryContract.CallOpts, license)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_SentryContract *SentryContractCallerSession) IsActive(license *big.Int) (bool, error) {
	return _SentryContract.Contract.IsActive(&_SentryContract.CallOpts, license)
}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_SentryContract *SentryContractCaller) LicenseOperator(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "licenseOperator", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_SentryContract *SentryContractSession) LicenseOperator(arg0 *big.Int) ([]byte, error) {
	return _SentryContract.Contract.LicenseOperator(&_SentryContract.CallOpts, arg0)
}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_SentryContract *SentryContractCallerSession) LicenseOperator(arg0 *big.Int) ([]byte, error) {
	return _SentryContract.Contract.LicenseOperator(&_SentryContract.CallOpts, arg0)
}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_SentryContract *SentryContractCaller) LicenseOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "licenseOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_SentryContract *SentryContractSession) LicenseOwner(arg0 *big.Int) (common.Address, error) {
	return _SentryContract.Contract.LicenseOwner(&_SentryContract.CallOpts, arg0)
}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_SentryContract *SentryContractCallerSession) LicenseOwner(arg0 *big.Int) (common.Address, error) {
	return _SentryContract.Contract.LicenseOwner(&_SentryContract.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_SentryContract *SentryContractCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_SentryContract *SentryContractSession) Locked() (bool, error) {
	return _SentryContract.Contract.Locked(&_SentryContract.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_SentryContract *SentryContractCallerSession) Locked() (bool, error) {
	return _SentryContract.Contract.Locked(&_SentryContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_SentryContract *SentryContractCaller) Network(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_SentryContract *SentryContractSession) Network() (common.Address, error) {
	return _SentryContract.Contract.Network(&_SentryContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_SentryContract *SentryContractCallerSession) Network() (common.Address, error) {
	return _SentryContract.Contract.Network(&_SentryContract.CallOpts)
}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_SentryContract *SentryContractCaller) NodesOwned(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "nodesOwned", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_SentryContract *SentryContractSession) NodesOwned(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _SentryContract.Contract.NodesOwned(&_SentryContract.CallOpts, arg0, arg1)
}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_SentryContract *SentryContractCallerSession) NodesOwned(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _SentryContract.Contract.NodesOwned(&_SentryContract.CallOpts, arg0, arg1)
}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCaller) OperatorCycleLicenseCount(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "operatorCycleLicenseCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractSession) OperatorCycleLicenseCount(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.OperatorCycleLicenseCount(&_SentryContract.CallOpts, arg0, arg1)
}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) OperatorCycleLicenseCount(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.OperatorCycleLicenseCount(&_SentryContract.CallOpts, arg0, arg1)
}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_SentryContract *SentryContractCaller) OperatorLicenseCount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "operatorLicenseCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_SentryContract *SentryContractSession) OperatorLicenseCount(arg0 []byte) (*big.Int, error) {
	return _SentryContract.Contract.OperatorLicenseCount(&_SentryContract.CallOpts, arg0)
}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) OperatorLicenseCount(arg0 []byte) (*big.Int, error) {
	return _SentryContract.Contract.OperatorLicenseCount(&_SentryContract.CallOpts, arg0)
}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCaller) OperatorLicenses(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "operatorLicenses", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractSession) OperatorLicenses(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.OperatorLicenses(&_SentryContract.CallOpts, arg0, arg1)
}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_SentryContract *SentryContractCallerSession) OperatorLicenses(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _SentryContract.Contract.OperatorLicenses(&_SentryContract.CallOpts, arg0, arg1)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(bytes)
func (_SentryContract *SentryContractCaller) Operators(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(bytes)
func (_SentryContract *SentryContractSession) Operators(arg0 *big.Int) ([]byte, error) {
	return _SentryContract.Contract.Operators(&_SentryContract.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(bytes)
func (_SentryContract *SentryContractCallerSession) Operators(arg0 *big.Int) ([]byte, error) {
	return _SentryContract.Contract.Operators(&_SentryContract.CallOpts, arg0)
}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_SentryContract *SentryContractCaller) OperatorsOwner(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "operatorsOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_SentryContract *SentryContractSession) OperatorsOwner(arg0 []byte) (common.Address, error) {
	return _SentryContract.Contract.OperatorsOwner(&_SentryContract.CallOpts, arg0)
}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_SentryContract *SentryContractCallerSession) OperatorsOwner(arg0 []byte) (common.Address, error) {
	return _SentryContract.Contract.OperatorsOwner(&_SentryContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SentryContract *SentryContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SentryContract *SentryContractSession) Owner() (common.Address, error) {
	return _SentryContract.Contract.Owner(&_SentryContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SentryContract *SentryContractCallerSession) Owner() (common.Address, error) {
	return _SentryContract.Contract.Owner(&_SentryContract.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_SentryContract *SentryContractCaller) TotalAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SentryContract.contract.Call(opts, &out, "totalAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_SentryContract *SentryContractSession) TotalAccounts() (*big.Int, error) {
	return _SentryContract.Contract.TotalAccounts(&_SentryContract.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_SentryContract *SentryContractCallerSession) TotalAccounts() (*big.Int, error) {
	return _SentryContract.Contract.TotalAccounts(&_SentryContract.CallOpts)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_SentryContract *SentryContractTransactor) DeRegisterNodeOperator(opts *bind.TransactOpts, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "deRegisterNodeOperator", licenses)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_SentryContract *SentryContractSession) DeRegisterNodeOperator(licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.DeRegisterNodeOperator(&_SentryContract.TransactOpts, licenses)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_SentryContract *SentryContractTransactorSession) DeRegisterNodeOperator(licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.DeRegisterNodeOperator(&_SentryContract.TransactOpts, licenses)
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_SentryContract *SentryContractTransactor) FillLicenseCountGap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "fillLicenseCountGap")
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_SentryContract *SentryContractSession) FillLicenseCountGap() (*types.Transaction, error) {
	return _SentryContract.Contract.FillLicenseCountGap(&_SentryContract.TransactOpts)
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_SentryContract *SentryContractTransactorSession) FillLicenseCountGap() (*types.Transaction, error) {
	return _SentryContract.Contract.FillLicenseCountGap(&_SentryContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _network, address _token, uint256 licensePrice) returns()
func (_SentryContract *SentryContractTransactor) Initialize(opts *bind.TransactOpts, _network common.Address, _token common.Address, licensePrice *big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "initialize", _network, _token, licensePrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _network, address _token, uint256 licensePrice) returns()
func (_SentryContract *SentryContractSession) Initialize(_network common.Address, _token common.Address, licensePrice *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.Initialize(&_SentryContract.TransactOpts, _network, _token, licensePrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _network, address _token, uint256 licensePrice) returns()
func (_SentryContract *SentryContractTransactorSession) Initialize(_network common.Address, _token common.Address, licensePrice *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.Initialize(&_SentryContract.TransactOpts, _network, _token, licensePrice)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0x003da2df.
//
// Solidity: function purchaseLicense(uint256 quantity, address token) payable returns(uint256[])
func (_SentryContract *SentryContractTransactor) PurchaseLicense(opts *bind.TransactOpts, quantity *big.Int, token common.Address) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "purchaseLicense", quantity, token)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0x003da2df.
//
// Solidity: function purchaseLicense(uint256 quantity, address token) payable returns(uint256[])
func (_SentryContract *SentryContractSession) PurchaseLicense(quantity *big.Int, token common.Address) (*types.Transaction, error) {
	return _SentryContract.Contract.PurchaseLicense(&_SentryContract.TransactOpts, quantity, token)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0x003da2df.
//
// Solidity: function purchaseLicense(uint256 quantity, address token) payable returns(uint256[])
func (_SentryContract *SentryContractTransactorSession) PurchaseLicense(quantity *big.Int, token common.Address) (*types.Transaction, error) {
	return _SentryContract.Contract.PurchaseLicense(&_SentryContract.TransactOpts, quantity, token)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_SentryContract *SentryContractTransactor) RegisterNodeOperator(opts *bind.TransactOpts, regData RegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "registerNodeOperator", regData, licenses)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_SentryContract *SentryContractSession) RegisterNodeOperator(regData RegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.RegisterNodeOperator(&_SentryContract.TransactOpts, regData, licenses)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_SentryContract *SentryContractTransactorSession) RegisterNodeOperator(regData RegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.RegisterNodeOperator(&_SentryContract.TransactOpts, regData, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_SentryContract *SentryContractTransactor) RegisterOperator(opts *bind.TransactOpts, regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "registerOperator", regDataBytes, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_SentryContract *SentryContractSession) RegisterOperator(regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.RegisterOperator(&_SentryContract.TransactOpts, regDataBytes, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_SentryContract *SentryContractTransactorSession) RegisterOperator(regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.RegisterOperator(&_SentryContract.TransactOpts, regDataBytes, licenses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SentryContract *SentryContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SentryContract *SentryContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SentryContract.Contract.RenounceOwnership(&_SentryContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SentryContract *SentryContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SentryContract.Contract.RenounceOwnership(&_SentryContract.TransactOpts)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_SentryContract *SentryContractTransactor) SetCalibrator(opts *bind.TransactOpts, _calibrator *big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "setCalibrator", _calibrator)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_SentryContract *SentryContractSession) SetCalibrator(_calibrator *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.SetCalibrator(&_SentryContract.TransactOpts, _calibrator)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_SentryContract *SentryContractTransactorSession) SetCalibrator(_calibrator *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.SetCalibrator(&_SentryContract.TransactOpts, _calibrator)
}

// SetInitialLicencePrice is a paid mutator transaction binding the contract method 0x98bcd0b0.
//
// Solidity: function setInitialLicencePrice(address token, uint256 _price) returns()
func (_SentryContract *SentryContractTransactor) SetInitialLicencePrice(opts *bind.TransactOpts, token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "setInitialLicencePrice", token, _price)
}

// SetInitialLicencePrice is a paid mutator transaction binding the contract method 0x98bcd0b0.
//
// Solidity: function setInitialLicencePrice(address token, uint256 _price) returns()
func (_SentryContract *SentryContractSession) SetInitialLicencePrice(token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.SetInitialLicencePrice(&_SentryContract.TransactOpts, token, _price)
}

// SetInitialLicencePrice is a paid mutator transaction binding the contract method 0x98bcd0b0.
//
// Solidity: function setInitialLicencePrice(address token, uint256 _price) returns()
func (_SentryContract *SentryContractTransactorSession) SetInitialLicencePrice(token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.SetInitialLicencePrice(&_SentryContract.TransactOpts, token, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SentryContract *SentryContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SentryContract *SentryContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SentryContract.Contract.TransferOwnership(&_SentryContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SentryContract *SentryContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SentryContract.Contract.TransferOwnership(&_SentryContract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_SentryContract *SentryContractTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_SentryContract *SentryContractSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.Withdraw(&_SentryContract.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_SentryContract *SentryContractTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SentryContract.Contract.Withdraw(&_SentryContract.TransactOpts, token, to, amount)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_SentryContract *SentryContractTransactor) WithdrawEthers(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SentryContract.contract.Transact(opts, "withdrawEthers", to)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_SentryContract *SentryContractSession) WithdrawEthers(to common.Address) (*types.Transaction, error) {
	return _SentryContract.Contract.WithdrawEthers(&_SentryContract.TransactOpts, to)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_SentryContract *SentryContractTransactorSession) WithdrawEthers(to common.Address) (*types.Transaction, error) {
	return _SentryContract.Contract.WithdrawEthers(&_SentryContract.TransactOpts, to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SentryContract *SentryContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SentryContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SentryContract *SentryContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SentryContract.Contract.Fallback(&_SentryContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SentryContract *SentryContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SentryContract.Contract.Fallback(&_SentryContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SentryContract *SentryContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SentryContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SentryContract *SentryContractSession) Receive() (*types.Transaction, error) {
	return _SentryContract.Contract.Receive(&_SentryContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SentryContract *SentryContractTransactorSession) Receive() (*types.Transaction, error) {
	return _SentryContract.Contract.Receive(&_SentryContract.TransactOpts)
}

// SentryContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SentryContract contract.
type SentryContractInitializedIterator struct {
	Event *SentryContractInitialized // Event containing the contract specifics and raw log

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
func (it *SentryContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SentryContractInitialized)
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
		it.Event = new(SentryContractInitialized)
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
func (it *SentryContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SentryContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SentryContractInitialized represents a Initialized event raised by the SentryContract contract.
type SentryContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SentryContract *SentryContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*SentryContractInitializedIterator, error) {

	logs, sub, err := _SentryContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SentryContractInitializedIterator{contract: _SentryContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SentryContract *SentryContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SentryContractInitialized) (event.Subscription, error) {

	logs, sub, err := _SentryContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SentryContractInitialized)
				if err := _SentryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SentryContract *SentryContractFilterer) ParseInitialized(log types.Log) (*SentryContractInitialized, error) {
	event := new(SentryContractInitialized)
	if err := _SentryContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SentryContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SentryContract contract.
type SentryContractOwnershipTransferredIterator struct {
	Event *SentryContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SentryContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SentryContractOwnershipTransferred)
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
		it.Event = new(SentryContractOwnershipTransferred)
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
func (it *SentryContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SentryContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SentryContractOwnershipTransferred represents a OwnershipTransferred event raised by the SentryContract contract.
type SentryContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SentryContract *SentryContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SentryContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SentryContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SentryContractOwnershipTransferredIterator{contract: _SentryContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SentryContract *SentryContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SentryContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SentryContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SentryContractOwnershipTransferred)
				if err := _SentryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SentryContract *SentryContractFilterer) ParseOwnershipTransferred(log types.Log) (*SentryContractOwnershipTransferred, error) {
	event := new(SentryContractOwnershipTransferred)
	if err := _SentryContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SentryContractPurchaseEventIterator is returned from FilterPurchaseEvent and is used to iterate over the raw logs and unpacked data for PurchaseEvent events raised by the SentryContract contract.
type SentryContractPurchaseEventIterator struct {
	Event *SentryContractPurchaseEvent // Event containing the contract specifics and raw log

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
func (it *SentryContractPurchaseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SentryContractPurchaseEvent)
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
		it.Event = new(SentryContractPurchaseEvent)
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
func (it *SentryContractPurchaseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SentryContractPurchaseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SentryContractPurchaseEvent represents a PurchaseEvent event raised by the SentryContract contract.
type SentryContractPurchaseEvent struct {
	Account   common.Address
	Price     *big.Int
	Quantity  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseEvent is a free log retrieval operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_SentryContract *SentryContractFilterer) FilterPurchaseEvent(opts *bind.FilterOpts, account []common.Address) (*SentryContractPurchaseEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SentryContract.contract.FilterLogs(opts, "PurchaseEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &SentryContractPurchaseEventIterator{contract: _SentryContract.contract, event: "PurchaseEvent", logs: logs, sub: sub}, nil
}

// WatchPurchaseEvent is a free log subscription operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_SentryContract *SentryContractFilterer) WatchPurchaseEvent(opts *bind.WatchOpts, sink chan<- *SentryContractPurchaseEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SentryContract.contract.WatchLogs(opts, "PurchaseEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SentryContractPurchaseEvent)
				if err := _SentryContract.contract.UnpackLog(event, "PurchaseEvent", log); err != nil {
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

// ParsePurchaseEvent is a log parse operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_SentryContract *SentryContractFilterer) ParsePurchaseEvent(log types.Log) (*SentryContractPurchaseEvent, error) {
	event := new(SentryContractPurchaseEvent)
	if err := _SentryContract.contract.UnpackLog(event, "PurchaseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SentryContractReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the SentryContract contract.
type SentryContractReceivedIterator struct {
	Event *SentryContractReceived // Event containing the contract specifics and raw log

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
func (it *SentryContractReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SentryContractReceived)
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
		it.Event = new(SentryContractReceived)
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
func (it *SentryContractReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SentryContractReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SentryContractReceived represents a Received event raised by the SentryContract contract.
type SentryContractReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_SentryContract *SentryContractFilterer) FilterReceived(opts *bind.FilterOpts) (*SentryContractReceivedIterator, error) {

	logs, sub, err := _SentryContract.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &SentryContractReceivedIterator{contract: _SentryContract.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_SentryContract *SentryContractFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *SentryContractReceived) (event.Subscription, error) {

	logs, sub, err := _SentryContract.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SentryContractReceived)
				if err := _SentryContract.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_SentryContract *SentryContractFilterer) ParseReceived(log types.Log) (*SentryContractReceived, error) {
	event := new(SentryContractReceived)
	if err := _SentryContract.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
