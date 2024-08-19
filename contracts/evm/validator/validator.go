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
)

// ChainInfo is an auto generated low-level Go binding around an user-defined struct.
type ChainInfo struct {
	StartTime    *big.Int
	StartBlock   *big.Int
	CurrentBlock *big.Int
	CurrentEpoch *big.Int
	CurrentCycle *big.Int
}

// SentryV2NodeRegistrationData is an auto generated low-level Go binding around an user-defined struct.
type SentryV2NodeRegistrationData struct {
	PublicKey  []byte
	Nonce      *big.Int
	Signature  []byte
	Commitment common.Address
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PurchaseEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"activeLicensesIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calibrator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"deRegisterNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochDivider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillLicenseCountGap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentCycle\",\"type\":\"uint256\"}],\"internalType\":\"structChainInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCycle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getCycleActiveLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLicencePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getRegistrationData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"}],\"internalType\":\"structSentryV2Node.RegistrationData\",\"name\":\"regData\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"}],\"name\":\"getSentryLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"}],\"name\":\"getTotalValidatorLicenceCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"operator\",\"type\":\"bytes\"}],\"name\":\"getValidatorLicenceCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"getYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"}],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseOperator\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodesOwned\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorCycleLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"operatorLicenseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorLicenses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"operatorsOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"purchaseLicense\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"}],\"internalType\":\"structSentryV2Node.RegistrationData\",\"name\":\"regData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"registerNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"regDataBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"licenses\",\"type\":\"uint256[]\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_calibrator\",\"type\":\"uint256\"}],\"name\":\"setCalibrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setStartNodePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startNodePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawEthers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_Storage *StorageCaller) AccountLicenseCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "accountLicenseCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_Storage *StorageSession) AccountLicenseCount(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.AccountLicenseCount(&_Storage.CallOpts, arg0)
}

// AccountLicenseCount is a free data retrieval call binding the contract method 0x2a7fca77.
//
// Solidity: function accountLicenseCount(address ) view returns(uint256)
func (_Storage *StorageCallerSession) AccountLicenseCount(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.AccountLicenseCount(&_Storage.CallOpts, arg0)
}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_Storage *StorageCaller) AccountLicenses(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "accountLicenses", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_Storage *StorageSession) AccountLicenses(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Storage.Contract.AccountLicenses(&_Storage.CallOpts, arg0, arg1)
}

// AccountLicenses is a free data retrieval call binding the contract method 0xf80c2adc.
//
// Solidity: function accountLicenses(address , uint256 ) view returns(uint256)
func (_Storage *StorageCallerSession) AccountLicenses(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Storage.Contract.AccountLicenses(&_Storage.CallOpts, arg0, arg1)
}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_Storage *StorageCaller) ActiveLicenseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "activeLicenseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_Storage *StorageSession) ActiveLicenseCount() (*big.Int, error) {
	return _Storage.Contract.ActiveLicenseCount(&_Storage.CallOpts)
}

// ActiveLicenseCount is a free data retrieval call binding the contract method 0x4104991d.
//
// Solidity: function activeLicenseCount() view returns(uint256)
func (_Storage *StorageCallerSession) ActiveLicenseCount() (*big.Int, error) {
	return _Storage.Contract.ActiveLicenseCount(&_Storage.CallOpts)
}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_Storage *StorageCaller) ActiveLicensesIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "activeLicensesIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_Storage *StorageSession) ActiveLicensesIndex(arg0 *big.Int) (*big.Int, error) {
	return _Storage.Contract.ActiveLicensesIndex(&_Storage.CallOpts, arg0)
}

// ActiveLicensesIndex is a free data retrieval call binding the contract method 0xaee5a56c.
//
// Solidity: function activeLicensesIndex(uint256 ) view returns(uint256)
func (_Storage *StorageCallerSession) ActiveLicensesIndex(arg0 *big.Int) (*big.Int, error) {
	return _Storage.Contract.ActiveLicensesIndex(&_Storage.CallOpts, arg0)
}

// BlockTime is a free data retrieval call binding the contract method 0x48b15166.
//
// Solidity: function blockTime() view returns(uint256)
func (_Storage *StorageCaller) BlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "blockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockTime is a free data retrieval call binding the contract method 0x48b15166.
//
// Solidity: function blockTime() view returns(uint256)
func (_Storage *StorageSession) BlockTime() (*big.Int, error) {
	return _Storage.Contract.BlockTime(&_Storage.CallOpts)
}

// BlockTime is a free data retrieval call binding the contract method 0x48b15166.
//
// Solidity: function blockTime() view returns(uint256)
func (_Storage *StorageCallerSession) BlockTime() (*big.Int, error) {
	return _Storage.Contract.BlockTime(&_Storage.CallOpts)
}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_Storage *StorageCaller) Calibrator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "calibrator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_Storage *StorageSession) Calibrator() (*big.Int, error) {
	return _Storage.Contract.Calibrator(&_Storage.CallOpts)
}

// Calibrator is a free data retrieval call binding the contract method 0x399f0f64.
//
// Solidity: function calibrator() view returns(uint256)
func (_Storage *StorageCallerSession) Calibrator() (*big.Int, error) {
	return _Storage.Contract.Calibrator(&_Storage.CallOpts)
}

// EpochDivider is a free data retrieval call binding the contract method 0x6ca677b1.
//
// Solidity: function epochDivider() view returns(uint256)
func (_Storage *StorageCaller) EpochDivider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "epochDivider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDivider is a free data retrieval call binding the contract method 0x6ca677b1.
//
// Solidity: function epochDivider() view returns(uint256)
func (_Storage *StorageSession) EpochDivider() (*big.Int, error) {
	return _Storage.Contract.EpochDivider(&_Storage.CallOpts)
}

// EpochDivider is a free data retrieval call binding the contract method 0x6ca677b1.
//
// Solidity: function epochDivider() view returns(uint256)
func (_Storage *StorageCallerSession) EpochDivider() (*big.Int, error) {
	return _Storage.Contract.EpochDivider(&_Storage.CallOpts)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x21cae483.
//
// Solidity: function getChainInfo() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCaller) GetChainInfo(opts *bind.CallOpts) (ChainInfo, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getChainInfo")

	if err != nil {
		return *new(ChainInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ChainInfo)).(*ChainInfo)

	return out0, err

}

// GetChainInfo is a free data retrieval call binding the contract method 0x21cae483.
//
// Solidity: function getChainInfo() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageSession) GetChainInfo() (ChainInfo, error) {
	return _Storage.Contract.GetChainInfo(&_Storage.CallOpts)
}

// GetChainInfo is a free data retrieval call binding the contract method 0x21cae483.
//
// Solidity: function getChainInfo() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Storage *StorageCallerSession) GetChainInfo() (ChainInfo, error) {
	return _Storage.Contract.GetChainInfo(&_Storage.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_Storage *StorageCaller) GetCurrentBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCurrentBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_Storage *StorageSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _Storage.Contract.GetCurrentBlockNumber(&_Storage.CallOpts)
}

// GetCurrentBlockNumber is a free data retrieval call binding the contract method 0x6fd902e1.
//
// Solidity: function getCurrentBlockNumber() view returns(uint256)
func (_Storage *StorageCallerSession) GetCurrentBlockNumber() (*big.Int, error) {
	return _Storage.Contract.GetCurrentBlockNumber(&_Storage.CallOpts)
}

// GetCurrentCycle is a free data retrieval call binding the contract method 0xbe26ed7f.
//
// Solidity: function getCurrentCycle() view returns(uint256)
func (_Storage *StorageCaller) GetCurrentCycle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCurrentCycle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentCycle is a free data retrieval call binding the contract method 0xbe26ed7f.
//
// Solidity: function getCurrentCycle() view returns(uint256)
func (_Storage *StorageSession) GetCurrentCycle() (*big.Int, error) {
	return _Storage.Contract.GetCurrentCycle(&_Storage.CallOpts)
}

// GetCurrentCycle is a free data retrieval call binding the contract method 0xbe26ed7f.
//
// Solidity: function getCurrentCycle() view returns(uint256)
func (_Storage *StorageCallerSession) GetCurrentCycle() (*big.Int, error) {
	return _Storage.Contract.GetCurrentCycle(&_Storage.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Storage *StorageCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Storage *StorageSession) GetCurrentEpoch() (*big.Int, error) {
	return _Storage.Contract.GetCurrentEpoch(&_Storage.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_Storage *StorageCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _Storage.Contract.GetCurrentEpoch(&_Storage.CallOpts)
}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 blockNumber) view returns(uint256)
func (_Storage *StorageCaller) GetCycle(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCycle", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 blockNumber) view returns(uint256)
func (_Storage *StorageSession) GetCycle(blockNumber *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetCycle(&_Storage.CallOpts, blockNumber)
}

// GetCycle is a free data retrieval call binding the contract method 0x2026f638.
//
// Solidity: function getCycle(uint256 blockNumber) view returns(uint256)
func (_Storage *StorageCallerSession) GetCycle(blockNumber *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetCycle(&_Storage.CallOpts, blockNumber)
}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageCaller) GetCycleActiveLicenseCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCycleActiveLicenseCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageSession) GetCycleActiveLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetCycleActiveLicenseCount(&_Storage.CallOpts, cycle)
}

// GetCycleActiveLicenseCount is a free data retrieval call binding the contract method 0xd01b1fd3.
//
// Solidity: function getCycleActiveLicenseCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageCallerSession) GetCycleActiveLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetCycleActiveLicenseCount(&_Storage.CallOpts, cycle)
}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageCaller) GetCycleLicenseCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getCycleLicenseCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageSession) GetCycleLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetCycleLicenseCount(&_Storage.CallOpts, cycle)
}

// GetCycleLicenseCount is a free data retrieval call binding the contract method 0x6830ab97.
//
// Solidity: function getCycleLicenseCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageCallerSession) GetCycleLicenseCount(cycle *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetCycleLicenseCount(&_Storage.CallOpts, cycle)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 blockNumber) view returns(uint256)
func (_Storage *StorageCaller) GetEpoch(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getEpoch", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 blockNumber) view returns(uint256)
func (_Storage *StorageSession) GetEpoch(blockNumber *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetEpoch(&_Storage.CallOpts, blockNumber)
}

// GetEpoch is a free data retrieval call binding the contract method 0xbc0bc6ba.
//
// Solidity: function getEpoch(uint256 blockNumber) view returns(uint256)
func (_Storage *StorageCallerSession) GetEpoch(blockNumber *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetEpoch(&_Storage.CallOpts, blockNumber)
}

// GetLicencePrice is a free data retrieval call binding the contract method 0x9909d33d.
//
// Solidity: function getLicencePrice() view returns(uint256)
func (_Storage *StorageCaller) GetLicencePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getLicencePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLicencePrice is a free data retrieval call binding the contract method 0x9909d33d.
//
// Solidity: function getLicencePrice() view returns(uint256)
func (_Storage *StorageSession) GetLicencePrice() (*big.Int, error) {
	return _Storage.Contract.GetLicencePrice(&_Storage.CallOpts)
}

// GetLicencePrice is a free data retrieval call binding the contract method 0x9909d33d.
//
// Solidity: function getLicencePrice() view returns(uint256)
func (_Storage *StorageCallerSession) GetLicencePrice() (*big.Int, error) {
	return _Storage.Contract.GetLicencePrice(&_Storage.CallOpts)
}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_Storage *StorageCaller) GetRegistrationData(opts *bind.CallOpts, data []byte) (SentryV2NodeRegistrationData, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getRegistrationData", data)

	if err != nil {
		return *new(SentryV2NodeRegistrationData), err
	}

	out0 := *abi.ConvertType(out[0], new(SentryV2NodeRegistrationData)).(*SentryV2NodeRegistrationData)

	return out0, err

}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_Storage *StorageSession) GetRegistrationData(data []byte) (SentryV2NodeRegistrationData, error) {
	return _Storage.Contract.GetRegistrationData(&_Storage.CallOpts, data)
}

// GetRegistrationData is a free data retrieval call binding the contract method 0x274a77b9.
//
// Solidity: function getRegistrationData(bytes data) pure returns((bytes,uint256,bytes,address) regData)
func (_Storage *StorageCallerSession) GetRegistrationData(data []byte) (SentryV2NodeRegistrationData, error) {
	return _Storage.Contract.GetRegistrationData(&_Storage.CallOpts, data)
}

// GetSentryLicenseCount is a free data retrieval call binding the contract method 0x76a2a9f3.
//
// Solidity: function getSentryLicenseCount(uint256 cycle, bytes operator) view returns(uint256)
func (_Storage *StorageCaller) GetSentryLicenseCount(opts *bind.CallOpts, cycle *big.Int, operator []byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getSentryLicenseCount", cycle, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSentryLicenseCount is a free data retrieval call binding the contract method 0x76a2a9f3.
//
// Solidity: function getSentryLicenseCount(uint256 cycle, bytes operator) view returns(uint256)
func (_Storage *StorageSession) GetSentryLicenseCount(cycle *big.Int, operator []byte) (*big.Int, error) {
	return _Storage.Contract.GetSentryLicenseCount(&_Storage.CallOpts, cycle, operator)
}

// GetSentryLicenseCount is a free data retrieval call binding the contract method 0x76a2a9f3.
//
// Solidity: function getSentryLicenseCount(uint256 cycle, bytes operator) view returns(uint256)
func (_Storage *StorageCallerSession) GetSentryLicenseCount(cycle *big.Int, operator []byte) (*big.Int, error) {
	return _Storage.Contract.GetSentryLicenseCount(&_Storage.CallOpts, cycle, operator)
}

// GetStartBlock is a free data retrieval call binding the contract method 0xa5f18c01.
//
// Solidity: function getStartBlock() view returns(uint256)
func (_Storage *StorageCaller) GetStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartBlock is a free data retrieval call binding the contract method 0xa5f18c01.
//
// Solidity: function getStartBlock() view returns(uint256)
func (_Storage *StorageSession) GetStartBlock() (*big.Int, error) {
	return _Storage.Contract.GetStartBlock(&_Storage.CallOpts)
}

// GetStartBlock is a free data retrieval call binding the contract method 0xa5f18c01.
//
// Solidity: function getStartBlock() view returns(uint256)
func (_Storage *StorageCallerSession) GetStartBlock() (*big.Int, error) {
	return _Storage.Contract.GetStartBlock(&_Storage.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() view returns(uint256)
func (_Storage *StorageCaller) GetStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() view returns(uint256)
func (_Storage *StorageSession) GetStartTime() (*big.Int, error) {
	return _Storage.Contract.GetStartTime(&_Storage.CallOpts)
}

// GetStartTime is a free data retrieval call binding the contract method 0xc828371e.
//
// Solidity: function getStartTime() view returns(uint256)
func (_Storage *StorageCallerSession) GetStartTime() (*big.Int, error) {
	return _Storage.Contract.GetStartTime(&_Storage.CallOpts)
}

// GetTotalValidatorLicenceCount is a free data retrieval call binding the contract method 0xa6a8733e.
//
// Solidity: function getTotalValidatorLicenceCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageCaller) GetTotalValidatorLicenceCount(opts *bind.CallOpts, cycle *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getTotalValidatorLicenceCount", cycle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValidatorLicenceCount is a free data retrieval call binding the contract method 0xa6a8733e.
//
// Solidity: function getTotalValidatorLicenceCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageSession) GetTotalValidatorLicenceCount(cycle *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetTotalValidatorLicenceCount(&_Storage.CallOpts, cycle)
}

// GetTotalValidatorLicenceCount is a free data retrieval call binding the contract method 0xa6a8733e.
//
// Solidity: function getTotalValidatorLicenceCount(uint256 cycle) view returns(uint256)
func (_Storage *StorageCallerSession) GetTotalValidatorLicenceCount(cycle *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetTotalValidatorLicenceCount(&_Storage.CallOpts, cycle)
}

// GetValidatorLicenceCount is a free data retrieval call binding the contract method 0x852dbeb9.
//
// Solidity: function getValidatorLicenceCount(uint256 cycle, bytes operator) view returns(uint256)
func (_Storage *StorageCaller) GetValidatorLicenceCount(opts *bind.CallOpts, cycle *big.Int, operator []byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getValidatorLicenceCount", cycle, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorLicenceCount is a free data retrieval call binding the contract method 0x852dbeb9.
//
// Solidity: function getValidatorLicenceCount(uint256 cycle, bytes operator) view returns(uint256)
func (_Storage *StorageSession) GetValidatorLicenceCount(cycle *big.Int, operator []byte) (*big.Int, error) {
	return _Storage.Contract.GetValidatorLicenceCount(&_Storage.CallOpts, cycle, operator)
}

// GetValidatorLicenceCount is a free data retrieval call binding the contract method 0x852dbeb9.
//
// Solidity: function getValidatorLicenceCount(uint256 cycle, bytes operator) view returns(uint256)
func (_Storage *StorageCallerSession) GetValidatorLicenceCount(cycle *big.Int, operator []byte) (*big.Int, error) {
	return _Storage.Contract.GetValidatorLicenceCount(&_Storage.CallOpts, cycle, operator)
}

// GetYear is a free data retrieval call binding the contract method 0x92d66313.
//
// Solidity: function getYear(uint256 _block) view returns(uint256)
func (_Storage *StorageCaller) GetYear(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getYear", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetYear is a free data retrieval call binding the contract method 0x92d66313.
//
// Solidity: function getYear(uint256 _block) view returns(uint256)
func (_Storage *StorageSession) GetYear(_block *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetYear(&_Storage.CallOpts, _block)
}

// GetYear is a free data retrieval call binding the contract method 0x92d66313.
//
// Solidity: function getYear(uint256 _block) view returns(uint256)
func (_Storage *StorageCallerSession) GetYear(_block *big.Int) (*big.Int, error) {
	return _Storage.Contract.GetYear(&_Storage.CallOpts, _block)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_Storage *StorageCaller) IsActive(opts *bind.CallOpts, license *big.Int) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isActive", license)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_Storage *StorageSession) IsActive(license *big.Int) (bool, error) {
	return _Storage.Contract.IsActive(&_Storage.CallOpts, license)
}

// IsActive is a free data retrieval call binding the contract method 0x82afd23b.
//
// Solidity: function isActive(uint256 license) view returns(bool)
func (_Storage *StorageCallerSession) IsActive(license *big.Int) (bool, error) {
	return _Storage.Contract.IsActive(&_Storage.CallOpts, license)
}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_Storage *StorageCaller) LicenseOperator(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "licenseOperator", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_Storage *StorageSession) LicenseOperator(arg0 *big.Int) ([]byte, error) {
	return _Storage.Contract.LicenseOperator(&_Storage.CallOpts, arg0)
}

// LicenseOperator is a free data retrieval call binding the contract method 0xacb905f1.
//
// Solidity: function licenseOperator(uint256 ) view returns(bytes)
func (_Storage *StorageCallerSession) LicenseOperator(arg0 *big.Int) ([]byte, error) {
	return _Storage.Contract.LicenseOperator(&_Storage.CallOpts, arg0)
}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_Storage *StorageCaller) LicenseOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "licenseOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_Storage *StorageSession) LicenseOwner(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.LicenseOwner(&_Storage.CallOpts, arg0)
}

// LicenseOwner is a free data retrieval call binding the contract method 0x452dd0f7.
//
// Solidity: function licenseOwner(uint256 ) view returns(address)
func (_Storage *StorageCallerSession) LicenseOwner(arg0 *big.Int) (common.Address, error) {
	return _Storage.Contract.LicenseOwner(&_Storage.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Storage *StorageCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Storage *StorageSession) Locked() (bool, error) {
	return _Storage.Contract.Locked(&_Storage.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Storage *StorageCallerSession) Locked() (bool, error) {
	return _Storage.Contract.Locked(&_Storage.CallOpts)
}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_Storage *StorageCaller) NodesOwned(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "nodesOwned", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_Storage *StorageSession) NodesOwned(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Storage.Contract.NodesOwned(&_Storage.CallOpts, arg0, arg1)
}

// NodesOwned is a free data retrieval call binding the contract method 0x0117320b.
//
// Solidity: function nodesOwned(address , uint256 ) view returns(bytes)
func (_Storage *StorageCallerSession) NodesOwned(arg0 common.Address, arg1 *big.Int) ([]byte, error) {
	return _Storage.Contract.NodesOwned(&_Storage.CallOpts, arg0, arg1)
}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_Storage *StorageCaller) OperatorCycleLicenseCount(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "operatorCycleLicenseCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_Storage *StorageSession) OperatorCycleLicenseCount(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _Storage.Contract.OperatorCycleLicenseCount(&_Storage.CallOpts, arg0, arg1)
}

// OperatorCycleLicenseCount is a free data retrieval call binding the contract method 0x33bf29f2.
//
// Solidity: function operatorCycleLicenseCount(bytes , uint256 ) view returns(uint256)
func (_Storage *StorageCallerSession) OperatorCycleLicenseCount(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _Storage.Contract.OperatorCycleLicenseCount(&_Storage.CallOpts, arg0, arg1)
}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_Storage *StorageCaller) OperatorLicenseCount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "operatorLicenseCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_Storage *StorageSession) OperatorLicenseCount(arg0 []byte) (*big.Int, error) {
	return _Storage.Contract.OperatorLicenseCount(&_Storage.CallOpts, arg0)
}

// OperatorLicenseCount is a free data retrieval call binding the contract method 0x7e297e2f.
//
// Solidity: function operatorLicenseCount(bytes ) view returns(uint256)
func (_Storage *StorageCallerSession) OperatorLicenseCount(arg0 []byte) (*big.Int, error) {
	return _Storage.Contract.OperatorLicenseCount(&_Storage.CallOpts, arg0)
}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_Storage *StorageCaller) OperatorLicenses(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "operatorLicenses", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_Storage *StorageSession) OperatorLicenses(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _Storage.Contract.OperatorLicenses(&_Storage.CallOpts, arg0, arg1)
}

// OperatorLicenses is a free data retrieval call binding the contract method 0x846be5bd.
//
// Solidity: function operatorLicenses(bytes , uint256 ) view returns(uint256)
func (_Storage *StorageCallerSession) OperatorLicenses(arg0 []byte, arg1 *big.Int) (*big.Int, error) {
	return _Storage.Contract.OperatorLicenses(&_Storage.CallOpts, arg0, arg1)
}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_Storage *StorageCaller) OperatorsOwner(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "operatorsOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_Storage *StorageSession) OperatorsOwner(arg0 []byte) (common.Address, error) {
	return _Storage.Contract.OperatorsOwner(&_Storage.CallOpts, arg0)
}

// OperatorsOwner is a free data retrieval call binding the contract method 0xa7f5afee.
//
// Solidity: function operatorsOwner(bytes ) view returns(address)
func (_Storage *StorageCallerSession) OperatorsOwner(arg0 []byte) (common.Address, error) {
	return _Storage.Contract.OperatorsOwner(&_Storage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Storage *StorageCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Storage *StorageSession) StartBlock() (*big.Int, error) {
	return _Storage.Contract.StartBlock(&_Storage.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Storage *StorageCallerSession) StartBlock() (*big.Int, error) {
	return _Storage.Contract.StartBlock(&_Storage.CallOpts)
}

// StartNodePrice is a free data retrieval call binding the contract method 0xb99a92c0.
//
// Solidity: function startNodePrice() view returns(uint256)
func (_Storage *StorageCaller) StartNodePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "startNodePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartNodePrice is a free data retrieval call binding the contract method 0xb99a92c0.
//
// Solidity: function startNodePrice() view returns(uint256)
func (_Storage *StorageSession) StartNodePrice() (*big.Int, error) {
	return _Storage.Contract.StartNodePrice(&_Storage.CallOpts)
}

// StartNodePrice is a free data retrieval call binding the contract method 0xb99a92c0.
//
// Solidity: function startNodePrice() view returns(uint256)
func (_Storage *StorageCallerSession) StartNodePrice() (*big.Int, error) {
	return _Storage.Contract.StartNodePrice(&_Storage.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Storage *StorageCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Storage *StorageSession) StartTime() (*big.Int, error) {
	return _Storage.Contract.StartTime(&_Storage.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Storage *StorageCallerSession) StartTime() (*big.Int, error) {
	return _Storage.Contract.StartTime(&_Storage.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_Storage *StorageCaller) TotalAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_Storage *StorageSession) TotalAccounts() (*big.Int, error) {
	return _Storage.Contract.TotalAccounts(&_Storage.CallOpts)
}

// TotalAccounts is a free data retrieval call binding the contract method 0x58451f97.
//
// Solidity: function totalAccounts() view returns(uint256)
func (_Storage *StorageCallerSession) TotalAccounts() (*big.Int, error) {
	return _Storage.Contract.TotalAccounts(&_Storage.CallOpts)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_Storage *StorageTransactor) DeRegisterNodeOperator(opts *bind.TransactOpts, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deRegisterNodeOperator", licenses)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_Storage *StorageSession) DeRegisterNodeOperator(licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeRegisterNodeOperator(&_Storage.TransactOpts, licenses)
}

// DeRegisterNodeOperator is a paid mutator transaction binding the contract method 0x632402cc.
//
// Solidity: function deRegisterNodeOperator(uint256[] licenses) returns()
func (_Storage *StorageTransactorSession) DeRegisterNodeOperator(licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeRegisterNodeOperator(&_Storage.TransactOpts, licenses)
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_Storage *StorageTransactor) FillLicenseCountGap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillLicenseCountGap")
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_Storage *StorageSession) FillLicenseCountGap() (*types.Transaction, error) {
	return _Storage.Contract.FillLicenseCountGap(&_Storage.TransactOpts)
}

// FillLicenseCountGap is a paid mutator transaction binding the contract method 0x7b7de55d.
//
// Solidity: function fillLicenseCountGap() returns()
func (_Storage *StorageTransactorSession) FillLicenseCountGap() (*types.Transaction, error) {
	return _Storage.Contract.FillLicenseCountGap(&_Storage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _token, uint256 _blockTime, uint256 _startBlock) returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _blockTime *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize", _token, _blockTime, _startBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _token, uint256 _blockTime, uint256 _startBlock) returns()
func (_Storage *StorageSession) Initialize(_token common.Address, _blockTime *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, _token, _blockTime, _startBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _token, uint256 _blockTime, uint256 _startBlock) returns()
func (_Storage *StorageTransactorSession) Initialize(_token common.Address, _blockTime *big.Int, _startBlock *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, _token, _blockTime, _startBlock)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0xc8a028a8.
//
// Solidity: function purchaseLicense(uint256 quantity) payable returns(uint256[])
func (_Storage *StorageTransactor) PurchaseLicense(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "purchaseLicense", quantity)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0xc8a028a8.
//
// Solidity: function purchaseLicense(uint256 quantity) payable returns(uint256[])
func (_Storage *StorageSession) PurchaseLicense(quantity *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.PurchaseLicense(&_Storage.TransactOpts, quantity)
}

// PurchaseLicense is a paid mutator transaction binding the contract method 0xc8a028a8.
//
// Solidity: function purchaseLicense(uint256 quantity) payable returns(uint256[])
func (_Storage *StorageTransactorSession) PurchaseLicense(quantity *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.PurchaseLicense(&_Storage.TransactOpts, quantity)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_Storage *StorageTransactor) RegisterNodeOperator(opts *bind.TransactOpts, regData SentryV2NodeRegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "registerNodeOperator", regData, licenses)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_Storage *StorageSession) RegisterNodeOperator(regData SentryV2NodeRegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RegisterNodeOperator(&_Storage.TransactOpts, regData, licenses)
}

// RegisterNodeOperator is a paid mutator transaction binding the contract method 0x4ca5628f.
//
// Solidity: function registerNodeOperator((bytes,uint256,bytes,address) regData, uint256[] licenses) returns()
func (_Storage *StorageTransactorSession) RegisterNodeOperator(regData SentryV2NodeRegistrationData, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RegisterNodeOperator(&_Storage.TransactOpts, regData, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_Storage *StorageTransactor) RegisterOperator(opts *bind.TransactOpts, regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "registerOperator", regDataBytes, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_Storage *StorageSession) RegisterOperator(regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RegisterOperator(&_Storage.TransactOpts, regDataBytes, licenses)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xcdadf777.
//
// Solidity: function registerOperator(bytes regDataBytes, uint256[] licenses) returns()
func (_Storage *StorageTransactorSession) RegisterOperator(regDataBytes []byte, licenses []*big.Int) (*types.Transaction, error) {
	return _Storage.Contract.RegisterOperator(&_Storage.TransactOpts, regDataBytes, licenses)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Storage *StorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Storage.Contract.RenounceOwnership(&_Storage.TransactOpts)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_Storage *StorageTransactor) SetCalibrator(opts *bind.TransactOpts, _calibrator *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setCalibrator", _calibrator)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_Storage *StorageSession) SetCalibrator(_calibrator *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetCalibrator(&_Storage.TransactOpts, _calibrator)
}

// SetCalibrator is a paid mutator transaction binding the contract method 0x48922da5.
//
// Solidity: function setCalibrator(uint256 _calibrator) returns()
func (_Storage *StorageTransactorSession) SetCalibrator(_calibrator *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetCalibrator(&_Storage.TransactOpts, _calibrator)
}

// SetStartNodePrice is a paid mutator transaction binding the contract method 0xfd1c7362.
//
// Solidity: function setStartNodePrice(uint256 _price) returns()
func (_Storage *StorageTransactor) SetStartNodePrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setStartNodePrice", _price)
}

// SetStartNodePrice is a paid mutator transaction binding the contract method 0xfd1c7362.
//
// Solidity: function setStartNodePrice(uint256 _price) returns()
func (_Storage *StorageSession) SetStartNodePrice(_price *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetStartNodePrice(&_Storage.TransactOpts, _price)
}

// SetStartNodePrice is a paid mutator transaction binding the contract method 0xfd1c7362.
//
// Solidity: function setStartNodePrice(uint256 _price) returns()
func (_Storage *StorageTransactorSession) SetStartNodePrice(_price *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetStartNodePrice(&_Storage.TransactOpts, _price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Storage *StorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwnership(&_Storage.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Storage *StorageTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Storage *StorageSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Withdraw(&_Storage.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Storage *StorageTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.Withdraw(&_Storage.TransactOpts, token, to, amount)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_Storage *StorageTransactor) WithdrawEthers(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "withdrawEthers", to)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_Storage *StorageSession) WithdrawEthers(to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawEthers(&_Storage.TransactOpts, to)
}

// WithdrawEthers is a paid mutator transaction binding the contract method 0x2988a9f0.
//
// Solidity: function withdrawEthers(address to) returns()
func (_Storage *StorageTransactorSession) WithdrawEthers(to common.Address) (*types.Transaction, error) {
	return _Storage.Contract.WithdrawEthers(&_Storage.TransactOpts, to)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Storage.Contract.Fallback(&_Storage.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Storage *StorageTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Storage.Contract.Fallback(&_Storage.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Storage *StorageTransactorSession) Receive() (*types.Transaction, error) {
	return _Storage.Contract.Receive(&_Storage.TransactOpts)
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

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
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
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
		it.Event = new(StorageInitialized)
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
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Storage contract.
type StorageOwnershipTransferredIterator struct {
	Event *StorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnershipTransferred)
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
		it.Event = new(StorageOwnershipTransferred)
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
func (it *StorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnershipTransferred represents a OwnershipTransferred event raised by the Storage contract.
type StorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnershipTransferredIterator{contract: _Storage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Storage *StorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnershipTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Storage *StorageFilterer) ParseOwnershipTransferred(log types.Log) (*StorageOwnershipTransferred, error) {
	event := new(StorageOwnershipTransferred)
	if err := _Storage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoragePurchaseEventIterator is returned from FilterPurchaseEvent and is used to iterate over the raw logs and unpacked data for PurchaseEvent events raised by the Storage contract.
type StoragePurchaseEventIterator struct {
	Event *StoragePurchaseEvent // Event containing the contract specifics and raw log

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
func (it *StoragePurchaseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoragePurchaseEvent)
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
		it.Event = new(StoragePurchaseEvent)
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
func (it *StoragePurchaseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoragePurchaseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoragePurchaseEvent represents a PurchaseEvent event raised by the Storage contract.
type StoragePurchaseEvent struct {
	Account   common.Address
	Price     *big.Int
	Quantity  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseEvent is a free log retrieval operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_Storage *StorageFilterer) FilterPurchaseEvent(opts *bind.FilterOpts, account []common.Address) (*StoragePurchaseEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "PurchaseEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &StoragePurchaseEventIterator{contract: _Storage.contract, event: "PurchaseEvent", logs: logs, sub: sub}, nil
}

// WatchPurchaseEvent is a free log subscription operation binding the contract event 0x4d28b0527b61511e95e214c4b5dc5ef6a46f03f9484a44eb6168f446530a239b.
//
// Solidity: event PurchaseEvent(address indexed account, uint256 price, uint256 quantity, uint256 timestamp)
func (_Storage *StorageFilterer) WatchPurchaseEvent(opts *bind.WatchOpts, sink chan<- *StoragePurchaseEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "PurchaseEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoragePurchaseEvent)
				if err := _Storage.contract.UnpackLog(event, "PurchaseEvent", log); err != nil {
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
func (_Storage *StorageFilterer) ParsePurchaseEvent(log types.Log) (*StoragePurchaseEvent, error) {
	event := new(StoragePurchaseEvent)
	if err := _Storage.contract.UnpackLog(event, "PurchaseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Storage contract.
type StorageReceivedIterator struct {
	Event *StorageReceived // Event containing the contract specifics and raw log

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
func (it *StorageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageReceived)
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
		it.Event = new(StorageReceived)
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
func (it *StorageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageReceived represents a Received event raised by the Storage contract.
type StorageReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Storage *StorageFilterer) FilterReceived(opts *bind.FilterOpts) (*StorageReceivedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &StorageReceivedIterator{contract: _Storage.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Storage *StorageFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *StorageReceived) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageReceived)
				if err := _Storage.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_Storage *StorageFilterer) ParseReceived(log types.Log) (*StorageReceived, error) {
	event := new(StorageReceived)
	if err := _Storage.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
