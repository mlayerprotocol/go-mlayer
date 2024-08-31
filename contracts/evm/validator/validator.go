// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator

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

// ValidatorContractMetaData contains all meta data concerning the ValidatorContract contract.
var ValidatorContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PurchaseEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeLicensesIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calibrator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"deRegisterNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillLicenseCountGap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getCycleActiveLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLicencePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getOperatorCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"}],\"name\":\"getOperatorLicenses\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perPage\",\"type\":\"uint256\"}],\"name\":\"getOperators\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"opr\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getRegistrationData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"}],\"internalType\":\"structRegistrationData\",\"name\":\"regData\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"licensePrice\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseOperator\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"contractINetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesOwned\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"operatorLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"operatorsOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"purchaseLicense\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"}],\"internalType\":\"structRegistrationData\",\"name\":\"regData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"registerNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"regDataBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_calibrator\",\"type\":\"uint256\"}],\"name\":\"setCalibrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setInitialLicencePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ValidatorContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorContractMetaData.ABI instead.
var ValidatorContractABI = ValidatorContractMetaData.ABI

// ValidatorContract is an auto generated Go binding around an Ethereum contract.
type ValidatorContract struct {
	ValidatorContractCaller     // Read-only binding to the contract
	ValidatorContractTransactor // Write-only binding to the contract
	ValidatorContractFilterer   // Log filterer for contract events
}

// ValidatorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorContractSession struct {
	Contract     *ValidatorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorContractCallerSession struct {
	Contract *ValidatorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorContractTransactorSession struct {
	Contract     *ValidatorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorContractRaw struct {
	Contract *ValidatorContract // Generic contract binding to access the raw methods on
}

// ValidatorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorContractCallerRaw struct {
	Contract *ValidatorContractCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorContractTransactorRaw struct {
	Contract *ValidatorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorContract creates a new instance of ValidatorContract, bound to a specific deployed contract.
func NewValidatorContract(address common.Address, backend bind.ContractBackend) (*ValidatorContract, error) {
	contract, err := bindValidatorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorContract{ValidatorContractCaller: ValidatorContractCaller{contract: contract}, ValidatorContractTransactor: ValidatorContractTransactor{contract: contract}, ValidatorContractFilterer: ValidatorContractFilterer{contract: contract}}, nil
}

// NewValidatorContractCaller creates a new read-only instance of ValidatorContract, bound to a specific deployed contract.
func NewValidatorContractCaller(address common.Address, caller bind.ContractCaller) (*ValidatorContractCaller, error) {
	contract, err := bindValidatorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractCaller{contract: contract}, nil
}

// NewValidatorContractTransactor creates a new write-only instance of ValidatorContract, bound to a specific deployed contract.
func NewValidatorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorContractTransactor, error) {
	contract, err := bindValidatorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractTransactor{contract: contract}, nil
}

// NewValidatorContractFilterer creates a new log filterer instance of ValidatorContract, bound to a specific deployed contract.
func NewValidatorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorContractFilterer, error) {
	contract, err := bindValidatorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractFilterer{contract: contract}, nil
}

// bindValidatorContract binds a generic wrapper to an already deployed contract.
func bindValidatorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorContract *ValidatorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorContract.Contract.ValidatorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorContract *ValidatorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContract.Contract.ValidatorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorContract *ValidatorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorContract.Contract.ValidatorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorContract *ValidatorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorContract *ValidatorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorContract *ValidatorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorContract.Contract.contract.Transact(opts, method, params...)
}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) AccountLicenseCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "accountLicenseCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) AccountLicenseCount(arg0 common.Address) (*big.Int, error) {
	return _ValidatorContract.Contract.AccountLicenseCount(&_ValidatorContract.CallOpts, arg0)
}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) AccountLicenseCount(arg0 common.Address) (*big.Int, error) {
	return _ValidatorContract.Contract.AccountLicenseCount(&_ValidatorContract.CallOpts, arg0)
}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) AccountLicenses(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "accountLicenses", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) AccountLicenses(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.AccountLicenses(&_ValidatorContract.CallOpts, arg0, arg1)
}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) AccountLicenses(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.AccountLicenses(&_ValidatorContract.CallOpts, arg0, arg1)
}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) ActiveLicenseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "activeLicenseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) ActiveLicenseCount() (*big.Int, error) {
	return _ValidatorContract.Contract.ActiveLicenseCount(&_ValidatorContract.CallOpts)
}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) ActiveLicenseCount() (*big.Int, error) {
	return _ValidatorContract.Contract.ActiveLicenseCount(&_ValidatorContract.CallOpts)
}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) ActiveLicensesIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "activeLicensesIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) ActiveLicensesIndex(arg0 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.ActiveLicensesIndex(&_ValidatorContract.CallOpts, arg0)
}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) ActiveLicensesIndex(arg0 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.ActiveLicensesIndex(&_ValidatorContract.CallOpts, arg0)
}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) Calibrator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "calibrator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) Calibrator() (*big.Int, error) {
	return _ValidatorContract.Contract.Calibrator(&_ValidatorContract.CallOpts)
}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) Calibrator() (*big.Int, error) {
	return _ValidatorContract.Contract.Calibrator(&_ValidatorContract.CallOpts)
}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) GetCycleActiveLicenseCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getCycleActiveLicenseCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) GetCycleActiveLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.GetCycleActiveLicenseCount(&_ValidatorContract.CallOpts, cycle)
}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) GetCycleActiveLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.GetCycleActiveLicenseCount(&_ValidatorContract.CallOpts, cycle)
}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) GetCycleLicenseCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getCycleLicenseCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) GetCycleLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.GetCycleLicenseCount(&_ValidatorContract.CallOpts, cycle)
}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) GetCycleLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.GetCycleLicenseCount(&_ValidatorContract.CallOpts, cycle)
}

// GetLicencePrice is a free data retrieval call binding the contract method 0x1c9129a9.
//
// Solidity: function getLicencePrice(address token) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) GetLicencePrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getLicencePrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLicencePrice is a free data retrieval call binding the contract method 0x1c9129a9.
//
// Solidity: function getLicencePrice(address token) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) GetLicencePrice(token common.Address) (*big.Int, error) {
	return _ValidatorContract.Contract.GetLicencePrice(&_ValidatorContract.CallOpts, token)
}

// GetLicencePrice is a free data retrieval call binding the contract method 0x1c9129a9.
//
// Solidity: function getLicencePrice(address token) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) GetLicencePrice(token common.Address) (*big.Int, error) {
	return _ValidatorContract.Contract.GetLicencePrice(&_ValidatorContract.CallOpts, token)
}

// GetOperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x333825bd.
//
// Solidity: function getOperatorCycleLicenseCount(bytes operator, uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) GetOperatorCycleLicenseCount(opts *bind.CallOpts, operator []byte, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getOperatorCycleLicenseCount", operator, cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x333825bd.
//
// Solidity: function getOperatorCycleLicenseCount(bytes operator, uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) GetOperatorCycleLicenseCount(operator []byte, cycle *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.GetOperatorCycleLicenseCount(&_ValidatorContract.CallOpts, operator, cycle)
}

// GetOperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x333825bd.
//
// Solidity: function getOperatorCycleLicenseCount(bytes operator, uint256 cycle) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) GetOperatorCycleLicenseCount(operator []byte, cycle *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.GetOperatorCycleLicenseCount(&_ValidatorContract.CallOpts, operator, cycle)
}

// GetOperatorLicenses is a free data retrieval call binding the contract method 0xb6e1121f.
//
// Solidity: function getOperatorLicenses(bytes operator) view returns(uint256[])
func (_ValidatorContract *ValidatorContractCaller) GetOperatorLicenses(opts *bind.CallOpts, operator []byte) ([]*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getOperatorLicenses", operator)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorLicenses is a free data retrieval call binding the contract method 0xb6e1121f.
//
// Solidity: function getOperatorLicenses(bytes operator) view returns(uint256[])
func (_ValidatorContract *ValidatorContractSession) GetOperatorLicenses(operator []byte) ([]*big.Int, error) {
	return _ValidatorContract.Contract.GetOperatorLicenses(&_ValidatorContract.CallOpts, operator)
}

// GetOperatorLicenses is a free data retrieval call binding the contract method 0xb6e1121f.
//
// Solidity: function getOperatorLicenses(bytes operator) view returns(uint256[])
func (_ValidatorContract *ValidatorContractCallerSession) GetOperatorLicenses(operator []byte) ([]*big.Int, error) {
	return _ValidatorContract.Contract.GetOperatorLicenses(&_ValidatorContract.CallOpts, operator)
}

// GetOperators is a free data retrieval call binding the contract method 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 page, uint256 perPage) view returns(bytes[] opr)
func (_ValidatorContract *ValidatorContractCaller) GetOperators(opts *bind.CallOpts, page *big.Int, perPage *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getOperators", page, perPage)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 page, uint256 perPage) view returns(bytes[] opr)
func (_ValidatorContract *ValidatorContractSession) GetOperators(page *big.Int, perPage *big.Int) ([][]byte, error) {
	return _ValidatorContract.Contract.GetOperators(&_ValidatorContract.CallOpts, page, perPage)
}

// GetOperators is a free data retrieval call binding the contract method 0xea4dd2b9.
//
// Solidity: function getOperators(uint256 page, uint256 perPage) view returns(bytes[] opr)
func (_ValidatorContract *ValidatorContractCallerSession) GetOperators(page *big.Int, perPage *big.Int) ([][]byte, error) {
	return _ValidatorContract.Contract.GetOperators(&_ValidatorContract.CallOpts, page, perPage)
}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_ValidatorContract *ValidatorContractCaller) GetRegistrationData(opts *bind.CallOpts, data []byte) (RegistrationData, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "getRegistrationData", data)

	if err != nil {
		return *new(RegistrationData), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistrationData)).(*RegistrationData)

	return out0, err

}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_ValidatorContract *ValidatorContractSession) GetRegistrationData(data []byte) (RegistrationData, error) {
	return _ValidatorContract.Contract.GetRegistrationData(&_ValidatorContract.CallOpts, data)
}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_ValidatorContract *ValidatorContractCallerSession) GetRegistrationData(data []byte) (RegistrationData, error) {
	return _ValidatorContract.Contract.GetRegistrationData(&_ValidatorContract.CallOpts, data)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_ValidatorContract *ValidatorContractCaller) IsActive(opts *bind.CallOpts, license *big.Int) (bool, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "isActive", license)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_ValidatorContract *ValidatorContractSession) IsActive(license *big.Int) (bool, error) {
	return _ValidatorContract.Contract.IsActive(&_ValidatorContract.CallOpts, license)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_ValidatorContract *ValidatorContractCallerSession) IsActive(license *big.Int) (bool, error) {
	return _ValidatorContract.Contract.IsActive(&_ValidatorContract.CallOpts, license)
}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractCaller) LicenseOperator(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "licenseOperator", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractSession) LicenseOperator(arg0 *big.Int) ([]byte, error) {
	return _ValidatorContract.Contract.LicenseOperator(&_ValidatorContract.CallOpts, arg0)
}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractCallerSession) LicenseOperator(arg0 *big.Int) ([]byte, error) {
	return _ValidatorContract.Contract.LicenseOperator(&_ValidatorContract.CallOpts, arg0)
}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_ValidatorContract *ValidatorContractCaller) LicenseOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "licenseOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_ValidatorContract *ValidatorContractSession) LicenseOwner(arg0 *big.Int) (common.Address, error) {
	return _ValidatorContract.Contract.LicenseOwner(&_ValidatorContract.CallOpts, arg0)
}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_ValidatorContract *ValidatorContractCallerSession) LicenseOwner(arg0 *big.Int) (common.Address, error) {
	return _ValidatorContract.Contract.LicenseOwner(&_ValidatorContract.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_ValidatorContract *ValidatorContractCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_ValidatorContract *ValidatorContractSession) Locked() (bool, error) {
	return _ValidatorContract.Contract.Locked(&_ValidatorContract.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_ValidatorContract *ValidatorContractCallerSession) Locked() (bool, error) {
	return _ValidatorContract.Contract.Locked(&_ValidatorContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_ValidatorContract *ValidatorContractCaller) Network(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_ValidatorContract *ValidatorContractSession) Network() (common.Address, error) {
	return _ValidatorContract.Contract.Network(&_ValidatorContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_ValidatorContract *ValidatorContractCallerSession) Network() (common.Address, error) {
	return _ValidatorContract.Contract.Network(&_ValidatorContract.CallOpts)
}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractCaller) NodesOwned(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "nodesOwned", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractSession) NodesOwned(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _ValidatorContract.Contract.NodesOwned(&_ValidatorContract.CallOpts, arg0, arg1)
}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractCallerSession) NodesOwned(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _ValidatorContract.Contract.NodesOwned(&_ValidatorContract.CallOpts, arg0, arg1)
}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) OperatorCycleLicenseCount(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "operatorCycleLicenseCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) OperatorCycleLicenseCount(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.OperatorCycleLicenseCount(&_ValidatorContract.CallOpts, arg0, arg1)
}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) OperatorCycleLicenseCount(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.OperatorCycleLicenseCount(&_ValidatorContract.CallOpts, arg0, arg1)
}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) OperatorLicenseCount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "operatorLicenseCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) OperatorLicenseCount(arg0 []byte) (*big.Int, error) {
	return _ValidatorContract.Contract.OperatorLicenseCount(&_ValidatorContract.CallOpts, arg0)
}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) OperatorLicenseCount(arg0 []byte) (*big.Int, error) {
	return _ValidatorContract.Contract.OperatorLicenseCount(&_ValidatorContract.CallOpts, arg0)
}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) OperatorLicenses(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "operatorLicenses", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) OperatorLicenses(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.OperatorLicenses(&_ValidatorContract.CallOpts, arg0, arg1)
}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) OperatorLicenses(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _ValidatorContract.Contract.OperatorLicenses(&_ValidatorContract.CallOpts, arg0, arg1)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractCaller) Operators(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractSession) Operators(arg0 *big.Int) ([]byte, error) {
	return _ValidatorContract.Contract.Operators(&_ValidatorContract.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0xe28d4906.
//
// Solidity: function operators(uint256 ) view returns(bytes)
func (_ValidatorContract *ValidatorContractCallerSession) Operators(arg0 *big.Int) ([]byte, error) {
	return _ValidatorContract.Contract.Operators(&_ValidatorContract.CallOpts, arg0)
}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_ValidatorContract *ValidatorContractCaller) OperatorsOwner(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "operatorsOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_ValidatorContract *ValidatorContractSession) OperatorsOwner(arg0 []byte) (common.Address, error) {
	return _ValidatorContract.Contract.OperatorsOwner(&_ValidatorContract.CallOpts, arg0)
}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_ValidatorContract *ValidatorContractCallerSession) OperatorsOwner(arg0 []byte) (common.Address, error) {
	return _ValidatorContract.Contract.OperatorsOwner(&_ValidatorContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorContract *ValidatorContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorContract *ValidatorContractSession) Owner() (common.Address, error) {
	return _ValidatorContract.Contract.Owner(&_ValidatorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ValidatorContract *ValidatorContractCallerSession) Owner() (common.Address, error) {
	return _ValidatorContract.Contract.Owner(&_ValidatorContract.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_ValidatorContract *ValidatorContractCaller) TotalAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorContract.contract.Call(opts, &out, "totalAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_ValidatorContract *ValidatorContractSession) TotalAccounts() (*big.Int, error) {
	return _ValidatorContract.Contract.TotalAccounts(&_ValidatorContract.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_ValidatorContract *ValidatorContractCallerSession) TotalAccounts() (*big.Int, error) {
	return _ValidatorContract.Contract.TotalAccounts(&_ValidatorContract.CallOpts)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractTransactor) DeRegisterNodeOperator(opts *bind.TransactOpts, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "deRegisterNodeOperator", licenses)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractSession) DeRegisterNodeOperator(licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.DeRegisterNodeOperator(&_ValidatorContract.TransactOpts, licenses)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) DeRegisterNodeOperator(licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.DeRegisterNodeOperator(&_ValidatorContract.TransactOpts, licenses)
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_ValidatorContract *ValidatorContractTransactor) FillLicenseCountGap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "fillLicenseCountGap")
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_ValidatorContract *ValidatorContractSession) FillLicenseCountGap() (*types.Transaction, error) {
	return _ValidatorContract.Contract.FillLicenseCountGap(&_ValidatorContract.TransactOpts)
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_ValidatorContract *ValidatorContractTransactorSession) FillLicenseCountGap() (*types.Transaction, error) {
	return _ValidatorContract.Contract.FillLicenseCountGap(&_ValidatorContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _network, address _token, uint256 licensePrice) returns()
func (_ValidatorContract *ValidatorContractTransactor) Initialize(opts *bind.TransactOpts, _network common.Address, _token common.Address, licensePrice *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "initialize", _network, _token, licensePrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _network, address _token, uint256 licensePrice) returns()
func (_ValidatorContract *ValidatorContractSession) Initialize(_network common.Address, _token common.Address, licensePrice *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.Initialize(&_ValidatorContract.TransactOpts, _network, _token, licensePrice)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _network, address _token, uint256 licensePrice) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) Initialize(_network common.Address, _token common.Address, licensePrice *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.Initialize(&_ValidatorContract.TransactOpts, _network, _token, licensePrice)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0x003da2df.
//
// Solidity: function purchaseLicense(uint256 quantity, address token) payable returns(uint256[])
func (_ValidatorContract *ValidatorContractTransactor) PurchaseLicense(opts *bind.TransactOpts, quantity *big.Int, token common.Address) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "purchaseLicense", quantity, token)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0x003da2df.
//
// Solidity: function purchaseLicense(uint256 quantity, address token) payable returns(uint256[])
func (_ValidatorContract *ValidatorContractSession) PurchaseLicense(quantity *big.Int, token common.Address) (*types.Transaction, error) {
	return _ValidatorContract.Contract.PurchaseLicense(&_ValidatorContract.TransactOpts, quantity, token)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0x003da2df.
//
// Solidity: function purchaseLicense(uint256 quantity, address token) payable returns(uint256[])
func (_ValidatorContract *ValidatorContractTransactorSession) PurchaseLicense(quantity *big.Int, token common.Address) (*types.Transaction, error) {
	return _ValidatorContract.Contract.PurchaseLicense(&_ValidatorContract.TransactOpts, quantity, token)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractTransactor) RegisterNodeOperator(opts *bind.TransactOpts, regData RegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "registerNodeOperator", regData, licenses)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractSession) RegisterNodeOperator(regData RegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.RegisterNodeOperator(&_ValidatorContract.TransactOpts, regData, licenses)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) RegisterNodeOperator(regData RegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.RegisterNodeOperator(&_ValidatorContract.TransactOpts, regData, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractTransactor) RegisterOperator(opts *bind.TransactOpts, regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "registerOperator", regDataBytes, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractSession) RegisterOperator(regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.RegisterOperator(&_ValidatorContract.TransactOpts, regDataBytes, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) RegisterOperator(regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.RegisterOperator(&_ValidatorContract.TransactOpts, regDataBytes, licenses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorContract *ValidatorContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorContract *ValidatorContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorContract.Contract.RenounceOwnership(&_ValidatorContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ValidatorContract *ValidatorContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ValidatorContract.Contract.RenounceOwnership(&_ValidatorContract.TransactOpts)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_ValidatorContract *ValidatorContractTransactor) SetCalibrator(opts *bind.TransactOpts, _calibrator *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "setCalibrator", _calibrator)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_ValidatorContract *ValidatorContractSession) SetCalibrator(_calibrator *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.SetCalibrator(&_ValidatorContract.TransactOpts, _calibrator)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) SetCalibrator(_calibrator *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.SetCalibrator(&_ValidatorContract.TransactOpts, _calibrator)
}

// SetInitialLicencePrice is a paid mutator transaction binding the contract method 0x98bcd0b0.
//
// Solidity: function setInitialLicencePrice(address token, uint256 _price) returns()
func (_ValidatorContract *ValidatorContractTransactor) SetInitialLicencePrice(opts *bind.TransactOpts, token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "setInitialLicencePrice", token, _price)
}

// SetInitialLicencePrice is a paid mutator transaction binding the contract method 0x98bcd0b0.
//
// Solidity: function setInitialLicencePrice(address token, uint256 _price) returns()
func (_ValidatorContract *ValidatorContractSession) SetInitialLicencePrice(token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.SetInitialLicencePrice(&_ValidatorContract.TransactOpts, token, _price)
}

// SetInitialLicencePrice is a paid mutator transaction binding the contract method 0x98bcd0b0.
//
// Solidity: function setInitialLicencePrice(address token, uint256 _price) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) SetInitialLicencePrice(token common.Address, _price *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.SetInitialLicencePrice(&_ValidatorContract.TransactOpts, token, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorContract *ValidatorContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorContract *ValidatorContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorContract.Contract.TransferOwnership(&_ValidatorContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ValidatorContract.Contract.TransferOwnership(&_ValidatorContract.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_ValidatorContract *ValidatorContractTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_ValidatorContract *ValidatorContractSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.Withdraw(&_ValidatorContract.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ValidatorContract.Contract.Withdraw(&_ValidatorContract.TransactOpts, token, to, amount)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_ValidatorContract *ValidatorContractTransactor) WithdrawEthers(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ValidatorContract.contract.Transact(opts, "withdrawEthers", to)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_ValidatorContract *ValidatorContractSession) WithdrawEthers(to common.Address) (*types.Transaction, error) {
	return _ValidatorContract.Contract.WithdrawEthers(&_ValidatorContract.TransactOpts, to)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_ValidatorContract *ValidatorContractTransactorSession) WithdrawEthers(to common.Address) (*types.Transaction, error) {
	return _ValidatorContract.Contract.WithdrawEthers(&_ValidatorContract.TransactOpts, to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ValidatorContract *ValidatorContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ValidatorContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ValidatorContract *ValidatorContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ValidatorContract.Contract.Fallback(&_ValidatorContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ValidatorContract *ValidatorContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ValidatorContract.Contract.Fallback(&_ValidatorContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorContract *ValidatorContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorContract *ValidatorContractSession) Receive() (*types.Transaction, error) {
	return _ValidatorContract.Contract.Receive(&_ValidatorContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ValidatorContract *ValidatorContractTransactorSession) Receive() (*types.Transaction, error) {
	return _ValidatorContract.Contract.Receive(&_ValidatorContract.TransactOpts)
}

// ValidatorContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorContract contract.
type ValidatorContractInitializedIterator struct {
	Event *ValidatorContractInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorContractInitialized)
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
		it.Event = new(ValidatorContractInitialized)
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
func (it *ValidatorContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorContractInitialized represents a Initialized event raised by the ValidatorContract contract.
type ValidatorContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorContract *ValidatorContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorContractInitializedIterator, error) {

	logs, sub, err := _ValidatorContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorContractInitializedIterator{contract: _ValidatorContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ValidatorContract *ValidatorContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorContractInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorContractInitialized)
				if err := _ValidatorContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ValidatorContract *ValidatorContractFilterer) ParseInitialized(log types.Log) (*ValidatorContractInitialized, error) {
	event := new(ValidatorContractInitialized)
	if err := _ValidatorContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ValidatorContract contract.
type ValidatorContractOwnershipTransferredIterator struct {
	Event *ValidatorContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorContractOwnershipTransferred)
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
		it.Event = new(ValidatorContractOwnershipTransferred)
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
func (it *ValidatorContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorContractOwnershipTransferred represents a OwnershipTransferred event raised by the ValidatorContract contract.
type ValidatorContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorContract *ValidatorContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractOwnershipTransferredIterator{contract: _ValidatorContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ValidatorContract *ValidatorContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ValidatorContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorContractOwnershipTransferred)
				if err := _ValidatorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ValidatorContract *ValidatorContractFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorContractOwnershipTransferred, error) {
	event := new(ValidatorContractOwnershipTransferred)
	if err := _ValidatorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorContractPurchaseEventIterator is returned from FilterPurchaseEvent and is used to iterate over the raw logs and unpacked data for PurchaseEvent events raised by the ValidatorContract contract.
type ValidatorContractPurchaseEventIterator struct {
	Event *ValidatorContractPurchaseEvent // Event containing the contract specifics and raw log

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
func (it *ValidatorContractPurchaseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorContractPurchaseEvent)
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
		it.Event = new(ValidatorContractPurchaseEvent)
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
func (it *ValidatorContractPurchaseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorContractPurchaseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorContractPurchaseEvent represents a PurchaseEvent event raised by the ValidatorContract contract.
type ValidatorContractPurchaseEvent struct {
	Account   common.Address
	Price     *big.Int
	Quantity  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseEvent is a free log retrieval operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_ValidatorContract *ValidatorContractFilterer) FilterPurchaseEvent(opts *bind.FilterOpts, account []common.Address) (*ValidatorContractPurchaseEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ValidatorContract.contract.FilterLogs(opts, "PurchaseEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorContractPurchaseEventIterator{contract: _ValidatorContract.contract, event: "PurchaseEvent", logs: logs, sub: sub}, nil
}

// WatchPurchaseEvent is a free log subscription operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_ValidatorContract *ValidatorContractFilterer) WatchPurchaseEvent(opts *bind.WatchOpts, sink chan<- *ValidatorContractPurchaseEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ValidatorContract.contract.WatchLogs(opts, "PurchaseEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorContractPurchaseEvent)
				if err := _ValidatorContract.contract.UnpackLog(event, "PurchaseEvent", log); err != nil {
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
func (_ValidatorContract *ValidatorContractFilterer) ParsePurchaseEvent(log types.Log) (*ValidatorContractPurchaseEvent, error) {
	event := new(ValidatorContractPurchaseEvent)
	if err := _ValidatorContract.contract.UnpackLog(event, "PurchaseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorContractReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the ValidatorContract contract.
type ValidatorContractReceivedIterator struct {
	Event *ValidatorContractReceived // Event containing the contract specifics and raw log

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
func (it *ValidatorContractReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorContractReceived)
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
		it.Event = new(ValidatorContractReceived)
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
func (it *ValidatorContractReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorContractReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorContractReceived represents a Received event raised by the ValidatorContract contract.
type ValidatorContractReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ValidatorContract *ValidatorContractFilterer) FilterReceived(opts *bind.FilterOpts) (*ValidatorContractReceivedIterator, error) {

	logs, sub, err := _ValidatorContract.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &ValidatorContractReceivedIterator{contract: _ValidatorContract.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_ValidatorContract *ValidatorContractFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ValidatorContractReceived) (event.Subscription, error) {

	logs, sub, err := _ValidatorContract.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorContractReceived)
				if err := _ValidatorContract.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_ValidatorContract *ValidatorContractFilterer) ParseReceived(log types.Log) (*ValidatorContractReceived, error) {
	event := new(ValidatorContractReceived)
	if err := _ValidatorContract.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
