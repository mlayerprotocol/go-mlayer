// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package app

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

// LibSecp256k1Point is an auto generated low-level Go binding around an user-defined struct.
type LibSecp256k1Point struct {
	X *big.Int
	Y *big.Int
}

// ApplicationClaim is an auto generated low-level Go binding around an user-defined struct.
type ApplicationClaim struct {
	Validator  []byte
	ClaimData  []ApplicationRewardClaimData
	Cycle      *big.Int
	Index      *big.Int
	Signers    []LibSecp256k1Point
	Commitment common.Address
	Signature  []byte
	TotalCost  *big.Int
}

// ApplicationLicenseInfo is an auto generated low-level Go binding around an user-defined struct.
type ApplicationLicenseInfo struct {
	Id          *big.Int
	Earned      *big.Int
	IsValidator bool
	DelegatedTo []byte
}

// ApplicationRewardClaimData is an auto generated low-level Go binding around an user-defined struct.
type ApplicationRewardClaimData struct {
	ApplicationId [16]byte
	Amount   *big.Int
}

// ApplicationStakeStruct is an auto generated low-level Go binding around an user-defined struct.
type ApplicationStakeStruct struct {
	Amount    *big.Int
	Timestamp *big.Int
}

// ApplicationContractMetaData contains all meta data concerning the ApplicationContract contract.
var ApplicationContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structApplication.StakeStruct\",\"name\":\"stake\",\"type\":\"tuple\"}],\"name\":\"StakeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structApplication.StakeStruct\",\"name\":\"stake\",\"type\":\"tuple\"}],\"name\":\"UnStakeEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"prevCycleReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allTimeReward\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidator\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"delegatedTo\",\"type\":\"bytes\"}],\"internalType\":\"structApplication.LicenseInfo[]\",\"name\":\"info\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structLibSecp256k1.Point[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"enableWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structLibSecp256k1.Point[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"name\":\"getClaimHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycleNumLicences\",\"type\":\"uint256\"}],\"name\":\"getMinSignerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getApplicationAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"}],\"name\":\"hashRewardData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sentryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorNodeContract\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"contractINetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"processedClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proofProviderRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sentryBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sentryContract\",\"outputs\":[{\"internalType\":\"contractINodeContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentryCycleRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentryLicenseRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStakable\",\"type\":\"uint256\"}],\"name\":\"setMinStakable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setNetworkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseReward\",\"type\":\"uint256\"}],\"name\":\"setSentryBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setSentryNodeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseReward\",\"type\":\"uint256\"}],\"name\":\"setValidatorBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setValidatorNodeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"}],\"name\":\"appBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"appBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"appCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"appDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"appStakerBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"}],\"name\":\"unStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unstakeOrders\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorCycleRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorNodeContract\",\"outputs\":[{\"internalType\":\"contractINodeContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structLibSecp256k1.Point[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"internalType\":\"structApplication.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ApplicationContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ApplicationContractMetaData.ABI instead.
var ApplicationContractABI = ApplicationContractMetaData.ABI

// ApplicationContract is an auto generated Go binding around an Ethereum contract.
type ApplicationContract struct {
	ApplicationContractCaller     // Read-only binding to the contract
	ApplicationContractTransactor // Write-only binding to the contract
	ApplicationContractFilterer   // Log filterer for contract events
}

// ApplicationContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApplicationContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplicationContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApplicationContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplicationContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApplicationContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApplicationContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApplicationContractSession struct {
	Contract     *ApplicationContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApplicationContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApplicationContractCallerSession struct {
	Contract *ApplicationContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ApplicationContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApplicationContractTransactorSession struct {
	Contract     *ApplicationContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ApplicationContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApplicationContractRaw struct {
	Contract *ApplicationContract // Generic contract binding to access the raw methods on
}

// ApplicationContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApplicationContractCallerRaw struct {
	Contract *ApplicationContractCaller // Generic read-only contract binding to access the raw methods on
}

// ApplicationContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApplicationContractTransactorRaw struct {
	Contract *ApplicationContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApplicationContract creates a new instance of ApplicationContract, bound to a specific deployed contract.
func NewApplicationContract(address common.Address, backend bind.ContractBackend) (*ApplicationContract, error) {
	contract, err := bindApplicationContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApplicationContract{ApplicationContractCaller: ApplicationContractCaller{contract: contract}, ApplicationContractTransactor: ApplicationContractTransactor{contract: contract}, ApplicationContractFilterer: ApplicationContractFilterer{contract: contract}}, nil
}

// NewApplicationContractCaller creates a new read-only instance of ApplicationContract, bound to a specific deployed contract.
func NewApplicationContractCaller(address common.Address, caller bind.ContractCaller) (*ApplicationContractCaller, error) {
	contract, err := bindApplicationContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractCaller{contract: contract}, nil
}

// NewApplicationContractTransactor creates a new write-only instance of ApplicationContract, bound to a specific deployed contract.
func NewApplicationContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ApplicationContractTransactor, error) {
	contract, err := bindApplicationContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractTransactor{contract: contract}, nil
}

// NewApplicationContractFilterer creates a new log filterer instance of ApplicationContract, bound to a specific deployed contract.
func NewApplicationContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ApplicationContractFilterer, error) {
	contract, err := bindApplicationContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractFilterer{contract: contract}, nil
}

// bindApplicationContract binds a generic wrapper to an already deployed contract.
func bindApplicationContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApplicationContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApplicationContract *ApplicationContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApplicationContract.Contract.ApplicationContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApplicationContract *ApplicationContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApplicationContract.Contract.ApplicationContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApplicationContract *ApplicationContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApplicationContract.Contract.ApplicationContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApplicationContract *ApplicationContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApplicationContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApplicationContract *ApplicationContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApplicationContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApplicationContract *ApplicationContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApplicationContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ApplicationContract *ApplicationContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ApplicationContract *ApplicationContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ApplicationContract.Contract.DEFAULTADMINROLE(&_ApplicationContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ApplicationContract *ApplicationContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ApplicationContract.Contract.DEFAULTADMINROLE(&_ApplicationContract.CallOpts)
}

// AddressInfo is a free data retrieval call binding the contract method 0x2126fcb2.
//
// Solidity: function addressInfo(address addr) view returns(uint256 prevCycleReward, uint256 allTimeReward, (uint256,uint256,bool,bytes)[] info, uint256 claimable)
func (_ApplicationContract *ApplicationContractCaller) AddressInfo(opts *bind.CallOpts, addr common.Address) (struct {
	PrevCycleReward *big.Int
	AllTimeReward   *big.Int
	Info            []ApplicationLicenseInfo
	Claimable       *big.Int
}, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "addressInfo", addr)

	outstruct := new(struct {
		PrevCycleReward *big.Int
		AllTimeReward   *big.Int
		Info            []ApplicationLicenseInfo
		Claimable       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrevCycleReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AllTimeReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Info = *abi.ConvertType(out[2], new([]ApplicationLicenseInfo)).(*[]ApplicationLicenseInfo)
	outstruct.Claimable = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AddressInfo is a free data retrieval call binding the contract method 0x2126fcb2.
//
// Solidity: function addressInfo(address addr) view returns(uint256 prevCycleReward, uint256 allTimeReward, (uint256,uint256,bool,bytes)[] info, uint256 claimable)
func (_ApplicationContract *ApplicationContractSession) AddressInfo(addr common.Address) (struct {
	PrevCycleReward *big.Int
	AllTimeReward   *big.Int
	Info            []ApplicationLicenseInfo
	Claimable       *big.Int
}, error) {
	return _ApplicationContract.Contract.AddressInfo(&_ApplicationContract.CallOpts, addr)
}

// AddressInfo is a free data retrieval call binding the contract method 0x2126fcb2.
//
// Solidity: function addressInfo(address addr) view returns(uint256 prevCycleReward, uint256 allTimeReward, (uint256,uint256,bool,bytes)[] info, uint256 claimable)
func (_ApplicationContract *ApplicationContractCallerSession) AddressInfo(addr common.Address) (struct {
	PrevCycleReward *big.Int
	AllTimeReward   *big.Int
	Info            []ApplicationLicenseInfo
	Claimable       *big.Int
}, error) {
	return _ApplicationContract.Contract.AddressInfo(&_ApplicationContract.CallOpts, addr)
}

// GetClaimHash is a free data retrieval call binding the contract method 0xf6b85a9a.
//
// Solidity: function getClaimHash((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bytes32)
func (_ApplicationContract *ApplicationContractCaller) GetClaimHash(opts *bind.CallOpts, claim ApplicationClaim) ([32]byte, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "getClaimHash", claim)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetClaimHash is a free data retrieval call binding the contract method 0xf6b85a9a.
//
// Solidity: function getClaimHash((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bytes32)
func (_ApplicationContract *ApplicationContractSession) GetClaimHash(claim ApplicationClaim) ([32]byte, error) {
	return _ApplicationContract.Contract.GetClaimHash(&_ApplicationContract.CallOpts, claim)
}

// GetClaimHash is a free data retrieval call binding the contract method 0xf6b85a9a.
//
// Solidity: function getClaimHash((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bytes32)
func (_ApplicationContract *ApplicationContractCallerSession) GetClaimHash(claim ApplicationClaim) ([32]byte, error) {
	return _ApplicationContract.Contract.GetClaimHash(&_ApplicationContract.CallOpts, claim)
}

// GetMinSignerCount is a free data retrieval call binding the contract method 0x9cee2fe3.
//
// Solidity: function getMinSignerCount(uint256 cycleNumLicences) pure returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) GetMinSignerCount(opts *bind.CallOpts, cycleNumLicences *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "getMinSignerCount", cycleNumLicences)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinSignerCount is a free data retrieval call binding the contract method 0x9cee2fe3.
//
// Solidity: function getMinSignerCount(uint256 cycleNumLicences) pure returns(uint256)
func (_ApplicationContract *ApplicationContractSession) GetMinSignerCount(cycleNumLicences *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.GetMinSignerCount(&_ApplicationContract.CallOpts, cycleNumLicences)
}

// GetMinSignerCount is a free data retrieval call binding the contract method 0x9cee2fe3.
//
// Solidity: function getMinSignerCount(uint256 cycleNumLicences) pure returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) GetMinSignerCount(cycleNumLicences *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.GetMinSignerCount(&_ApplicationContract.CallOpts, cycleNumLicences)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ApplicationContract *ApplicationContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ApplicationContract *ApplicationContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ApplicationContract.Contract.GetRoleAdmin(&_ApplicationContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ApplicationContract *ApplicationContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ApplicationContract.Contract.GetRoleAdmin(&_ApplicationContract.CallOpts, role)
}

// GetApplicationAccountBalance is a free data retrieval call binding the contract method 0x09774aaf.
//
// Solidity: function getApplicationAccountBalance(bytes16 appId, address addr) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) GetApplicationAccountBalance(opts *bind.CallOpts, appId [16]byte, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "getApplicationAccountBalance", appId, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetApplicationAccountBalance is a free data retrieval call binding the contract method 0x09774aaf.
//
// Solidity: function getApplicationAccountBalance(bytes16 appId, address addr) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) GetApplicationAccountBalance(appId [16]byte, addr common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.GetApplicationAccountBalance(&_ApplicationContract.CallOpts, appId, addr)
}

// GetApplicationAccountBalance is a free data retrieval call binding the contract method 0x09774aaf.
//
// Solidity: function getApplicationAccountBalance(bytes16 appId, address addr) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) GetApplicationAccountBalance(appId [16]byte, addr common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.GetApplicationAccountBalance(&_ApplicationContract.CallOpts, appId, addr)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ApplicationContract *ApplicationContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ApplicationContract *ApplicationContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ApplicationContract.Contract.HasRole(&_ApplicationContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ApplicationContract *ApplicationContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ApplicationContract.Contract.HasRole(&_ApplicationContract.CallOpts, role, account)
}

// HashRewardData is a free data retrieval call binding the contract method 0x5a9a7b87.
//
// Solidity: function hashRewardData((bytes16,uint256)[] claimData) pure returns(bytes32 hash)
func (_ApplicationContract *ApplicationContractCaller) HashRewardData(opts *bind.CallOpts, claimData []ApplicationRewardClaimData) ([32]byte, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "hashRewardData", claimData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashRewardData is a free data retrieval call binding the contract method 0x5a9a7b87.
//
// Solidity: function hashRewardData((bytes16,uint256)[] claimData) pure returns(bytes32 hash)
func (_ApplicationContract *ApplicationContractSession) HashRewardData(claimData []ApplicationRewardClaimData) ([32]byte, error) {
	return _ApplicationContract.Contract.HashRewardData(&_ApplicationContract.CallOpts, claimData)
}

// HashRewardData is a free data retrieval call binding the contract method 0x5a9a7b87.
//
// Solidity: function hashRewardData((bytes16,uint256)[] claimData) pure returns(bytes32 hash)
func (_ApplicationContract *ApplicationContractCallerSession) HashRewardData(claimData []ApplicationRewardClaimData) ([32]byte, error) {
	return _ApplicationContract.Contract.HashRewardData(&_ApplicationContract.CallOpts, claimData)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_ApplicationContract *ApplicationContractCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_ApplicationContract *ApplicationContractSession) Locked() (bool, error) {
	return _ApplicationContract.Contract.Locked(&_ApplicationContract.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_ApplicationContract *ApplicationContractCallerSession) Locked() (bool, error) {
	return _ApplicationContract.Contract.Locked(&_ApplicationContract.CallOpts)
}

// MinStakable is a free data retrieval call binding the contract method 0xc0d41476.
//
// Solidity: function minStakable() view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) MinStakable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "minStakable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakable is a free data retrieval call binding the contract method 0xc0d41476.
//
// Solidity: function minStakable() view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) MinStakable() (*big.Int, error) {
	return _ApplicationContract.Contract.MinStakable(&_ApplicationContract.CallOpts)
}

// MinStakable is a free data retrieval call binding the contract method 0xc0d41476.
//
// Solidity: function minStakable() view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) MinStakable() (*big.Int, error) {
	return _ApplicationContract.Contract.MinStakable(&_ApplicationContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_ApplicationContract *ApplicationContractCaller) Network(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_ApplicationContract *ApplicationContractSession) Network() (common.Address, error) {
	return _ApplicationContract.Contract.Network(&_ApplicationContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_ApplicationContract *ApplicationContractCallerSession) Network() (common.Address, error) {
	return _ApplicationContract.Contract.Network(&_ApplicationContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApplicationContract *ApplicationContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApplicationContract *ApplicationContractSession) Owner() (common.Address, error) {
	return _ApplicationContract.Contract.Owner(&_ApplicationContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ApplicationContract *ApplicationContractCallerSession) Owner() (common.Address, error) {
	return _ApplicationContract.Contract.Owner(&_ApplicationContract.CallOpts)
}

// ProcessedClaim is a free data retrieval call binding the contract method 0x2ad6f7f6.
//
// Solidity: function processedClaim(uint256 , bytes , uint256 ) view returns(bool)
func (_ApplicationContract *ApplicationContractCaller) ProcessedClaim(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "processedClaim", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedClaim is a free data retrieval call binding the contract method 0x2ad6f7f6.
//
// Solidity: function processedClaim(uint256 , bytes , uint256 ) view returns(bool)
func (_ApplicationContract *ApplicationContractSession) ProcessedClaim(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (bool, error) {
	return _ApplicationContract.Contract.ProcessedClaim(&_ApplicationContract.CallOpts, arg0, arg1, arg2)
}

// ProcessedClaim is a free data retrieval call binding the contract method 0x2ad6f7f6.
//
// Solidity: function processedClaim(uint256 , bytes , uint256 ) view returns(bool)
func (_ApplicationContract *ApplicationContractCallerSession) ProcessedClaim(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (bool, error) {
	return _ApplicationContract.Contract.ProcessedClaim(&_ApplicationContract.CallOpts, arg0, arg1, arg2)
}

// ProofProviderRewards is a free data retrieval call binding the contract method 0x58f25d13.
//
// Solidity: function proofProviderRewards(address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ProofProviderRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "proofProviderRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProofProviderRewards is a free data retrieval call binding the contract method 0x58f25d13.
//
// Solidity: function proofProviderRewards(address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ProofProviderRewards(arg0 common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.ProofProviderRewards(&_ApplicationContract.CallOpts, arg0)
}

// ProofProviderRewards is a free data retrieval call binding the contract method 0x58f25d13.
//
// Solidity: function proofProviderRewards(address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ProofProviderRewards(arg0 common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.ProofProviderRewards(&_ApplicationContract.CallOpts, arg0)
}

// SentryBaseReward is a free data retrieval call binding the contract method 0x25a772a4.
//
// Solidity: function sentryBaseReward() view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) SentryBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "sentryBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SentryBaseReward is a free data retrieval call binding the contract method 0x25a772a4.
//
// Solidity: function sentryBaseReward() view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) SentryBaseReward() (*big.Int, error) {
	return _ApplicationContract.Contract.SentryBaseReward(&_ApplicationContract.CallOpts)
}

// SentryBaseReward is a free data retrieval call binding the contract method 0x25a772a4.
//
// Solidity: function sentryBaseReward() view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) SentryBaseReward() (*big.Int, error) {
	return _ApplicationContract.Contract.SentryBaseReward(&_ApplicationContract.CallOpts)
}

// SentryContract is a free data retrieval call binding the contract method 0xdbca5c52.
//
// Solidity: function sentryContract() view returns(address)
func (_ApplicationContract *ApplicationContractCaller) SentryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "sentryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SentryContract is a free data retrieval call binding the contract method 0xdbca5c52.
//
// Solidity: function sentryContract() view returns(address)
func (_ApplicationContract *ApplicationContractSession) SentryContract() (common.Address, error) {
	return _ApplicationContract.Contract.SentryContract(&_ApplicationContract.CallOpts)
}

// SentryContract is a free data retrieval call binding the contract method 0xdbca5c52.
//
// Solidity: function sentryContract() view returns(address)
func (_ApplicationContract *ApplicationContractCallerSession) SentryContract() (common.Address, error) {
	return _ApplicationContract.Contract.SentryContract(&_ApplicationContract.CallOpts)
}

// SentryCycleRewards is a free data retrieval call binding the contract method 0x7773b6e9.
//
// Solidity: function sentryCycleRewards(address , uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) SentryCycleRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "sentryCycleRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SentryCycleRewards is a free data retrieval call binding the contract method 0x7773b6e9.
//
// Solidity: function sentryCycleRewards(address , uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) SentryCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.SentryCycleRewards(&_ApplicationContract.CallOpts, arg0, arg1)
}

// SentryCycleRewards is a free data retrieval call binding the contract method 0x7773b6e9.
//
// Solidity: function sentryCycleRewards(address , uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) SentryCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.SentryCycleRewards(&_ApplicationContract.CallOpts, arg0, arg1)
}

// SentryLicenseRevenue is a free data retrieval call binding the contract method 0x37b560f3.
//
// Solidity: function sentryLicenseRevenue(uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) SentryLicenseRevenue(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "sentryLicenseRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SentryLicenseRevenue is a free data retrieval call binding the contract method 0x37b560f3.
//
// Solidity: function sentryLicenseRevenue(uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) SentryLicenseRevenue(arg0 *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.SentryLicenseRevenue(&_ApplicationContract.CallOpts, arg0)
}

// SentryLicenseRevenue is a free data retrieval call binding the contract method 0x37b560f3.
//
// Solidity: function sentryLicenseRevenue(uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) SentryLicenseRevenue(arg0 *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.SentryLicenseRevenue(&_ApplicationContract.CallOpts, arg0)
}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_ApplicationContract *ApplicationContractCaller) StakeAddresses(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "stakeAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_ApplicationContract *ApplicationContractSession) StakeAddresses(arg0 common.Address) (common.Address, error) {
	return _ApplicationContract.Contract.StakeAddresses(&_ApplicationContract.CallOpts, arg0)
}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_ApplicationContract *ApplicationContractCallerSession) StakeAddresses(arg0 common.Address) (common.Address, error) {
	return _ApplicationContract.Contract.StakeAddresses(&_ApplicationContract.CallOpts, arg0)
}

// ApplicationBalance is a free data retrieval call binding the contract method 0x4d1e4062.
//
// Solidity: function appBalance(bytes16 appId) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ApplicationBalance(opts *bind.CallOpts, appId [16]byte) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "appBalance", appId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationBalance is a free data retrieval call binding the contract method 0x4d1e4062.
//
// Solidity: function appBalance(bytes16 appId) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ApplicationBalance(appId [16]byte) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationBalance(&_ApplicationContract.CallOpts, appId)
}

// ApplicationBalance is a free data retrieval call binding the contract method 0x4d1e4062.
//
// Solidity: function appBalance(bytes16 appId) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ApplicationBalance(appId [16]byte) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationBalance(&_ApplicationContract.CallOpts, appId)
}

// ApplicationBalances is a free data retrieval call binding the contract method 0x5a50b09b.
//
// Solidity: function appBalances(bytes16 , address , uint256 ) view returns(uint256 amount, uint256 timestamp)
func (_ApplicationContract *ApplicationContractCaller) ApplicationBalances(opts *bind.CallOpts, arg0 [16]byte, arg1 common.Address, arg2 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "appBalances", arg0, arg1, arg2)

	outstruct := new(struct {
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ApplicationBalances is a free data retrieval call binding the contract method 0x5a50b09b.
//
// Solidity: function appBalances(bytes16 , address , uint256 ) view returns(uint256 amount, uint256 timestamp)
func (_ApplicationContract *ApplicationContractSession) ApplicationBalances(arg0 [16]byte, arg1 common.Address, arg2 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _ApplicationContract.Contract.ApplicationBalances(&_ApplicationContract.CallOpts, arg0, arg1, arg2)
}

// ApplicationBalances is a free data retrieval call binding the contract method 0x5a50b09b.
//
// Solidity: function appBalances(bytes16 , address , uint256 ) view returns(uint256 amount, uint256 timestamp)
func (_ApplicationContract *ApplicationContractCallerSession) ApplicationBalances(arg0 [16]byte, arg1 common.Address, arg2 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _ApplicationContract.Contract.ApplicationBalances(&_ApplicationContract.CallOpts, arg0, arg1, arg2)
}

// ApplicationCredit is a free data retrieval call binding the contract method 0x8d73c2fa.
//
// Solidity: function appCredit(bytes16 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ApplicationCredit(opts *bind.CallOpts, arg0 [16]byte) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "appCredit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationCredit is a free data retrieval call binding the contract method 0x8d73c2fa.
//
// Solidity: function appCredit(bytes16 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ApplicationCredit(arg0 [16]byte) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationCredit(&_ApplicationContract.CallOpts, arg0)
}

// ApplicationCredit is a free data retrieval call binding the contract method 0x8d73c2fa.
//
// Solidity: function appCredit(bytes16 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ApplicationCredit(arg0 [16]byte) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationCredit(&_ApplicationContract.CallOpts, arg0)
}

// ApplicationDebt is a free data retrieval call binding the contract method 0xe9af524e.
//
// Solidity: function appDebt(bytes16 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ApplicationDebt(opts *bind.CallOpts, arg0 [16]byte) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "appDebt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationDebt is a free data retrieval call binding the contract method 0xe9af524e.
//
// Solidity: function appDebt(bytes16 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ApplicationDebt(arg0 [16]byte) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationDebt(&_ApplicationContract.CallOpts, arg0)
}

// ApplicationDebt is a free data retrieval call binding the contract method 0xe9af524e.
//
// Solidity: function appDebt(bytes16 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ApplicationDebt(arg0 [16]byte) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationDebt(&_ApplicationContract.CallOpts, arg0)
}

// ApplicationStakerBalances is a free data retrieval call binding the contract method 0x78d25664.
//
// Solidity: function appStakerBalances(bytes16 , address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ApplicationStakerBalances(opts *bind.CallOpts, arg0 [16]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "appStakerBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApplicationStakerBalances is a free data retrieval call binding the contract method 0x78d25664.
//
// Solidity: function appStakerBalances(bytes16 , address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ApplicationStakerBalances(arg0 [16]byte, arg1 common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationStakerBalances(&_ApplicationContract.CallOpts, arg0, arg1)
}

// ApplicationStakerBalances is a free data retrieval call binding the contract method 0x78d25664.
//
// Solidity: function appStakerBalances(bytes16 , address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ApplicationStakerBalances(arg0 [16]byte, arg1 common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.ApplicationStakerBalances(&_ApplicationContract.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApplicationContract *ApplicationContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApplicationContract *ApplicationContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ApplicationContract.Contract.SupportsInterface(&_ApplicationContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ApplicationContract *ApplicationContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ApplicationContract.Contract.SupportsInterface(&_ApplicationContract.CallOpts, interfaceId)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) TotalValueLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "totalValueLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) TotalValueLocked() (*big.Int, error) {
	return _ApplicationContract.Contract.TotalValueLocked(&_ApplicationContract.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) TotalValueLocked() (*big.Int, error) {
	return _ApplicationContract.Contract.TotalValueLocked(&_ApplicationContract.CallOpts)
}

// UnstakeOrders is a free data retrieval call binding the contract method 0x9d922aca.
//
// Solidity: function unstakeOrders(address , bytes ) view returns(int32)
func (_ApplicationContract *ApplicationContractCaller) UnstakeOrders(opts *bind.CallOpts, arg0 common.Address, arg1 []byte) (int32, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "unstakeOrders", arg0, arg1)

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// UnstakeOrders is a free data retrieval call binding the contract method 0x9d922aca.
//
// Solidity: function unstakeOrders(address , bytes ) view returns(int32)
func (_ApplicationContract *ApplicationContractSession) UnstakeOrders(arg0 common.Address, arg1 []byte) (int32, error) {
	return _ApplicationContract.Contract.UnstakeOrders(&_ApplicationContract.CallOpts, arg0, arg1)
}

// UnstakeOrders is a free data retrieval call binding the contract method 0x9d922aca.
//
// Solidity: function unstakeOrders(address , bytes ) view returns(int32)
func (_ApplicationContract *ApplicationContractCallerSession) UnstakeOrders(arg0 common.Address, arg1 []byte) (int32, error) {
	return _ApplicationContract.Contract.UnstakeOrders(&_ApplicationContract.CallOpts, arg0, arg1)
}

// ValidatorBaseReward is a free data retrieval call binding the contract method 0x0063b8dd.
//
// Solidity: function validatorBaseReward() view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ValidatorBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "validatorBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorBaseReward is a free data retrieval call binding the contract method 0x0063b8dd.
//
// Solidity: function validatorBaseReward() view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ValidatorBaseReward() (*big.Int, error) {
	return _ApplicationContract.Contract.ValidatorBaseReward(&_ApplicationContract.CallOpts)
}

// ValidatorBaseReward is a free data retrieval call binding the contract method 0x0063b8dd.
//
// Solidity: function validatorBaseReward() view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ValidatorBaseReward() (*big.Int, error) {
	return _ApplicationContract.Contract.ValidatorBaseReward(&_ApplicationContract.CallOpts)
}

// ValidatorCycleRewards is a free data retrieval call binding the contract method 0xbf19eb0c.
//
// Solidity: function validatorCycleRewards(address , uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ValidatorCycleRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "validatorCycleRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCycleRewards is a free data retrieval call binding the contract method 0xbf19eb0c.
//
// Solidity: function validatorCycleRewards(address , uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ValidatorCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.ValidatorCycleRewards(&_ApplicationContract.CallOpts, arg0, arg1)
}

// ValidatorCycleRewards is a free data retrieval call binding the contract method 0xbf19eb0c.
//
// Solidity: function validatorCycleRewards(address , uint256 ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ValidatorCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ApplicationContract.Contract.ValidatorCycleRewards(&_ApplicationContract.CallOpts, arg0, arg1)
}

// ValidatorNodeContract is a free data retrieval call binding the contract method 0x71f2053b.
//
// Solidity: function validatorNodeContract() view returns(address)
func (_ApplicationContract *ApplicationContractCaller) ValidatorNodeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "validatorNodeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorNodeContract is a free data retrieval call binding the contract method 0x71f2053b.
//
// Solidity: function validatorNodeContract() view returns(address)
func (_ApplicationContract *ApplicationContractSession) ValidatorNodeContract() (common.Address, error) {
	return _ApplicationContract.Contract.ValidatorNodeContract(&_ApplicationContract.CallOpts)
}

// ValidatorNodeContract is a free data retrieval call binding the contract method 0x71f2053b.
//
// Solidity: function validatorNodeContract() view returns(address)
func (_ApplicationContract *ApplicationContractCallerSession) ValidatorNodeContract() (common.Address, error) {
	return _ApplicationContract.Contract.ValidatorNodeContract(&_ApplicationContract.CallOpts)
}

// ValidatorRewards is a free data retrieval call binding the contract method 0xb1845c56.
//
// Solidity: function validatorRewards(address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) ValidatorRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "validatorRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRewards is a free data retrieval call binding the contract method 0xb1845c56.
//
// Solidity: function validatorRewards(address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractSession) ValidatorRewards(arg0 common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.ValidatorRewards(&_ApplicationContract.CallOpts, arg0)
}

// ValidatorRewards is a free data retrieval call binding the contract method 0xb1845c56.
//
// Solidity: function validatorRewards(address ) view returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) ValidatorRewards(arg0 common.Address) (*big.Int, error) {
	return _ApplicationContract.Contract.ValidatorRewards(&_ApplicationContract.CallOpts, arg0)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x8b2ed4c8.
//
// Solidity: function verifyClaim((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bool, bytes32)
func (_ApplicationContract *ApplicationContractCaller) VerifyClaim(opts *bind.CallOpts, claim ApplicationClaim) (bool, [32]byte, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "verifyClaim", claim)

	if err != nil {
		return *new(bool), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0x8b2ed4c8.
//
// Solidity: function verifyClaim((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bool, bytes32)
func (_ApplicationContract *ApplicationContractSession) VerifyClaim(claim ApplicationClaim) (bool, [32]byte, error) {
	return _ApplicationContract.Contract.VerifyClaim(&_ApplicationContract.CallOpts, claim)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x8b2ed4c8.
//
// Solidity: function verifyClaim((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bool, bytes32)
func (_ApplicationContract *ApplicationContractCallerSession) VerifyClaim(claim ApplicationClaim) (bool, [32]byte, error) {
	return _ApplicationContract.Contract.VerifyClaim(&_ApplicationContract.CallOpts, claim)
}

// WithdrawableAmount is a free data retrieval call binding the contract method 0x951303f5.
//
// Solidity: function withdrawableAmount() pure returns(uint256)
func (_ApplicationContract *ApplicationContractCaller) WithdrawableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "withdrawableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableAmount is a free data retrieval call binding the contract method 0x951303f5.
//
// Solidity: function withdrawableAmount() pure returns(uint256)
func (_ApplicationContract *ApplicationContractSession) WithdrawableAmount() (*big.Int, error) {
	return _ApplicationContract.Contract.WithdrawableAmount(&_ApplicationContract.CallOpts)
}

// WithdrawableAmount is a free data retrieval call binding the contract method 0x951303f5.
//
// Solidity: function withdrawableAmount() pure returns(uint256)
func (_ApplicationContract *ApplicationContractCallerSession) WithdrawableAmount() (*big.Int, error) {
	return _ApplicationContract.Contract.WithdrawableAmount(&_ApplicationContract.CallOpts)
}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_ApplicationContract *ApplicationContractCaller) WithdrawalEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ApplicationContract.contract.Call(opts, &out, "withdrawalEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_ApplicationContract *ApplicationContractSession) WithdrawalEnabled() (bool, error) {
	return _ApplicationContract.Contract.WithdrawalEnabled(&_ApplicationContract.CallOpts)
}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_ApplicationContract *ApplicationContractCallerSession) WithdrawalEnabled() (bool, error) {
	return _ApplicationContract.Contract.WithdrawalEnabled(&_ApplicationContract.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x25f3a9c0.
//
// Solidity: function claimReward((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) returns()
func (_ApplicationContract *ApplicationContractTransactor) ClaimReward(opts *bind.TransactOpts, claim ApplicationClaim) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "claimReward", claim)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x25f3a9c0.
//
// Solidity: function claimReward((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) returns()
func (_ApplicationContract *ApplicationContractSession) ClaimReward(claim ApplicationClaim) (*types.Transaction, error) {
	return _ApplicationContract.Contract.ClaimReward(&_ApplicationContract.TransactOpts, claim)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x25f3a9c0.
//
// Solidity: function claimReward((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) ClaimReward(claim ApplicationClaim) (*types.Transaction, error) {
	return _ApplicationContract.Contract.ClaimReward(&_ApplicationContract.TransactOpts, claim)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_ApplicationContract *ApplicationContractTransactor) EnableWithdrawal(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "enableWithdrawal", _enabled)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_ApplicationContract *ApplicationContractSession) EnableWithdrawal(_enabled bool) (*types.Transaction, error) {
	return _ApplicationContract.Contract.EnableWithdrawal(&_ApplicationContract.TransactOpts, _enabled)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) EnableWithdrawal(_enabled bool) (*types.Transaction, error) {
	return _ApplicationContract.Contract.EnableWithdrawal(&_ApplicationContract.TransactOpts, _enabled)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ApplicationContract *ApplicationContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ApplicationContract *ApplicationContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.GrantRole(&_ApplicationContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.GrantRole(&_ApplicationContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _network, address tokenAddress, address xTokenAddress, address _sentryContract, address _validatorNodeContract) returns()
func (_ApplicationContract *ApplicationContractTransactor) Initialize(opts *bind.TransactOpts, _network common.Address, tokenAddress common.Address, xTokenAddress common.Address, _sentryContract common.Address, _validatorNodeContract common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "initialize", _network, tokenAddress, xTokenAddress, _sentryContract, _validatorNodeContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _network, address tokenAddress, address xTokenAddress, address _sentryContract, address _validatorNodeContract) returns()
func (_ApplicationContract *ApplicationContractSession) Initialize(_network common.Address, tokenAddress common.Address, xTokenAddress common.Address, _sentryContract common.Address, _validatorNodeContract common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.Initialize(&_ApplicationContract.TransactOpts, _network, tokenAddress, xTokenAddress, _sentryContract, _validatorNodeContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _network, address tokenAddress, address xTokenAddress, address _sentryContract, address _validatorNodeContract) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) Initialize(_network common.Address, tokenAddress common.Address, xTokenAddress common.Address, _sentryContract common.Address, _validatorNodeContract common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.Initialize(&_ApplicationContract.TransactOpts, _network, tokenAddress, xTokenAddress, _sentryContract, _validatorNodeContract)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ApplicationContract *ApplicationContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ApplicationContract *ApplicationContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ApplicationContract.Contract.RenounceOwnership(&_ApplicationContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ApplicationContract *ApplicationContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ApplicationContract.Contract.RenounceOwnership(&_ApplicationContract.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ApplicationContract *ApplicationContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ApplicationContract *ApplicationContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.RenounceRole(&_ApplicationContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.RenounceRole(&_ApplicationContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ApplicationContract *ApplicationContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ApplicationContract *ApplicationContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.RevokeRole(&_ApplicationContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.RevokeRole(&_ApplicationContract.TransactOpts, role, account)
}

// SetMinStakable is a paid mutator transaction binding the contract method 0xf65ff7bd.
//
// Solidity: function setMinStakable(uint256 _minStakable) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetMinStakable(opts *bind.TransactOpts, _minStakable *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setMinStakable", _minStakable)
}

// SetMinStakable is a paid mutator transaction binding the contract method 0xf65ff7bd.
//
// Solidity: function setMinStakable(uint256 _minStakable) returns()
func (_ApplicationContract *ApplicationContractSession) SetMinStakable(_minStakable *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetMinStakable(&_ApplicationContract.TransactOpts, _minStakable)
}

// SetMinStakable is a paid mutator transaction binding the contract method 0xf65ff7bd.
//
// Solidity: function setMinStakable(uint256 _minStakable) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetMinStakable(_minStakable *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetMinStakable(&_ApplicationContract.TransactOpts, _minStakable)
}

// SetNetworkAddress is a paid mutator transaction binding the contract method 0x05f5dc95.
//
// Solidity: function setNetworkAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetNetworkAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setNetworkAddress", add)
}

// SetNetworkAddress is a paid mutator transaction binding the contract method 0x05f5dc95.
//
// Solidity: function setNetworkAddress(address add) returns()
func (_ApplicationContract *ApplicationContractSession) SetNetworkAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetNetworkAddress(&_ApplicationContract.TransactOpts, add)
}

// SetNetworkAddress is a paid mutator transaction binding the contract method 0x05f5dc95.
//
// Solidity: function setNetworkAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetNetworkAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetNetworkAddress(&_ApplicationContract.TransactOpts, add)
}

// SetSentryBaseReward is a paid mutator transaction binding the contract method 0x1227d71a.
//
// Solidity: function setSentryBaseReward(uint256 _baseReward) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetSentryBaseReward(opts *bind.TransactOpts, _baseReward *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setSentryBaseReward", _baseReward)
}

// SetSentryBaseReward is a paid mutator transaction binding the contract method 0x1227d71a.
//
// Solidity: function setSentryBaseReward(uint256 _baseReward) returns()
func (_ApplicationContract *ApplicationContractSession) SetSentryBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetSentryBaseReward(&_ApplicationContract.TransactOpts, _baseReward)
}

// SetSentryBaseReward is a paid mutator transaction binding the contract method 0x1227d71a.
//
// Solidity: function setSentryBaseReward(uint256 _baseReward) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetSentryBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetSentryBaseReward(&_ApplicationContract.TransactOpts, _baseReward)
}

// SetSentryNodeAddress is a paid mutator transaction binding the contract method 0x1e7fc3ea.
//
// Solidity: function setSentryNodeAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetSentryNodeAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setSentryNodeAddress", add)
}

// SetSentryNodeAddress is a paid mutator transaction binding the contract method 0x1e7fc3ea.
//
// Solidity: function setSentryNodeAddress(address add) returns()
func (_ApplicationContract *ApplicationContractSession) SetSentryNodeAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetSentryNodeAddress(&_ApplicationContract.TransactOpts, add)
}

// SetSentryNodeAddress is a paid mutator transaction binding the contract method 0x1e7fc3ea.
//
// Solidity: function setSentryNodeAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetSentryNodeAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetSentryNodeAddress(&_ApplicationContract.TransactOpts, add)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetTokenAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setTokenAddress", add)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address add) returns()
func (_ApplicationContract *ApplicationContractSession) SetTokenAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetTokenAddress(&_ApplicationContract.TransactOpts, add)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetTokenAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetTokenAddress(&_ApplicationContract.TransactOpts, add)
}

// SetValidatorBaseReward is a paid mutator transaction binding the contract method 0x8b043d02.
//
// Solidity: function setValidatorBaseReward(uint256 _baseReward) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetValidatorBaseReward(opts *bind.TransactOpts, _baseReward *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setValidatorBaseReward", _baseReward)
}

// SetValidatorBaseReward is a paid mutator transaction binding the contract method 0x8b043d02.
//
// Solidity: function setValidatorBaseReward(uint256 _baseReward) returns()
func (_ApplicationContract *ApplicationContractSession) SetValidatorBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetValidatorBaseReward(&_ApplicationContract.TransactOpts, _baseReward)
}

// SetValidatorBaseReward is a paid mutator transaction binding the contract method 0x8b043d02.
//
// Solidity: function setValidatorBaseReward(uint256 _baseReward) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetValidatorBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetValidatorBaseReward(&_ApplicationContract.TransactOpts, _baseReward)
}

// SetValidatorNodeAddress is a paid mutator transaction binding the contract method 0xb791bf39.
//
// Solidity: function setValidatorNodeAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactor) SetValidatorNodeAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "setValidatorNodeAddress", add)
}

// SetValidatorNodeAddress is a paid mutator transaction binding the contract method 0xb791bf39.
//
// Solidity: function setValidatorNodeAddress(address add) returns()
func (_ApplicationContract *ApplicationContractSession) SetValidatorNodeAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetValidatorNodeAddress(&_ApplicationContract.TransactOpts, add)
}

// SetValidatorNodeAddress is a paid mutator transaction binding the contract method 0xb791bf39.
//
// Solidity: function setValidatorNodeAddress(address add) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) SetValidatorNodeAddress(add common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.SetValidatorNodeAddress(&_ApplicationContract.TransactOpts, add)
}

// Stake is a paid mutator transaction binding the contract method 0xf47cfa50.
//
// Solidity: function stake(bytes16 appId, uint256 amount) returns()
func (_ApplicationContract *ApplicationContractTransactor) Stake(opts *bind.TransactOpts, appId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "stake", appId, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xf47cfa50.
//
// Solidity: function stake(bytes16 appId, uint256 amount) returns()
func (_ApplicationContract *ApplicationContractSession) Stake(appId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.Stake(&_ApplicationContract.TransactOpts, appId, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xf47cfa50.
//
// Solidity: function stake(bytes16 appId, uint256 amount) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) Stake(appId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _ApplicationContract.Contract.Stake(&_ApplicationContract.TransactOpts, appId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApplicationContract *ApplicationContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApplicationContract *ApplicationContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.TransferOwnership(&_ApplicationContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ApplicationContract.Contract.TransferOwnership(&_ApplicationContract.TransactOpts, newOwner)
}

// UnStake is a paid mutator transaction binding the contract method 0x7fd72127.
//
// Solidity: function unStake(bytes16 appId) returns()
func (_ApplicationContract *ApplicationContractTransactor) UnStake(opts *bind.TransactOpts, appId [16]byte) (*types.Transaction, error) {
	return _ApplicationContract.contract.Transact(opts, "unStake", appId)
}

// UnStake is a paid mutator transaction binding the contract method 0x7fd72127.
//
// Solidity: function unStake(bytes16 appId) returns()
func (_ApplicationContract *ApplicationContractSession) UnStake(appId [16]byte) (*types.Transaction, error) {
	return _ApplicationContract.Contract.UnStake(&_ApplicationContract.TransactOpts, appId)
}

// UnStake is a paid mutator transaction binding the contract method 0x7fd72127.
//
// Solidity: function unStake(bytes16 appId) returns()
func (_ApplicationContract *ApplicationContractTransactorSession) UnStake(appId [16]byte) (*types.Transaction, error) {
	return _ApplicationContract.Contract.UnStake(&_ApplicationContract.TransactOpts, appId)
}

// ApplicationContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ApplicationContract contract.
type ApplicationContractInitializedIterator struct {
	Event *ApplicationContractInitialized // Event containing the contract specifics and raw log

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
func (it *ApplicationContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractInitialized)
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
		it.Event = new(ApplicationContractInitialized)
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
func (it *ApplicationContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractInitialized represents a Initialized event raised by the ApplicationContract contract.
type ApplicationContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ApplicationContract *ApplicationContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ApplicationContractInitializedIterator, error) {

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ApplicationContractInitializedIterator{contract: _ApplicationContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ApplicationContract *ApplicationContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ApplicationContractInitialized) (event.Subscription, error) {

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractInitialized)
				if err := _ApplicationContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ApplicationContract *ApplicationContractFilterer) ParseInitialized(log types.Log) (*ApplicationContractInitialized, error) {
	event := new(ApplicationContractInitialized)
	if err := _ApplicationContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApplicationContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ApplicationContract contract.
type ApplicationContractOwnershipTransferredIterator struct {
	Event *ApplicationContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ApplicationContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractOwnershipTransferred)
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
		it.Event = new(ApplicationContractOwnershipTransferred)
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
func (it *ApplicationContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractOwnershipTransferred represents a OwnershipTransferred event raised by the ApplicationContract contract.
type ApplicationContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ApplicationContract *ApplicationContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ApplicationContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractOwnershipTransferredIterator{contract: _ApplicationContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ApplicationContract *ApplicationContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ApplicationContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractOwnershipTransferred)
				if err := _ApplicationContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ApplicationContract *ApplicationContractFilterer) ParseOwnershipTransferred(log types.Log) (*ApplicationContractOwnershipTransferred, error) {
	event := new(ApplicationContractOwnershipTransferred)
	if err := _ApplicationContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApplicationContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ApplicationContract contract.
type ApplicationContractRoleAdminChangedIterator struct {
	Event *ApplicationContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ApplicationContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractRoleAdminChanged)
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
		it.Event = new(ApplicationContractRoleAdminChanged)
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
func (it *ApplicationContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractRoleAdminChanged represents a RoleAdminChanged event raised by the ApplicationContract contract.
type ApplicationContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ApplicationContract *ApplicationContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ApplicationContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractRoleAdminChangedIterator{contract: _ApplicationContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ApplicationContract *ApplicationContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ApplicationContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractRoleAdminChanged)
				if err := _ApplicationContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ApplicationContract *ApplicationContractFilterer) ParseRoleAdminChanged(log types.Log) (*ApplicationContractRoleAdminChanged, error) {
	event := new(ApplicationContractRoleAdminChanged)
	if err := _ApplicationContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApplicationContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ApplicationContract contract.
type ApplicationContractRoleGrantedIterator struct {
	Event *ApplicationContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *ApplicationContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractRoleGranted)
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
		it.Event = new(ApplicationContractRoleGranted)
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
func (it *ApplicationContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractRoleGranted represents a RoleGranted event raised by the ApplicationContract contract.
type ApplicationContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApplicationContract *ApplicationContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApplicationContractRoleGrantedIterator, error) {

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

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractRoleGrantedIterator{contract: _ApplicationContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApplicationContract *ApplicationContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ApplicationContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractRoleGranted)
				if err := _ApplicationContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ApplicationContract *ApplicationContractFilterer) ParseRoleGranted(log types.Log) (*ApplicationContractRoleGranted, error) {
	event := new(ApplicationContractRoleGranted)
	if err := _ApplicationContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApplicationContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ApplicationContract contract.
type ApplicationContractRoleRevokedIterator struct {
	Event *ApplicationContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ApplicationContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractRoleRevoked)
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
		it.Event = new(ApplicationContractRoleRevoked)
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
func (it *ApplicationContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractRoleRevoked represents a RoleRevoked event raised by the ApplicationContract contract.
type ApplicationContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApplicationContract *ApplicationContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApplicationContractRoleRevokedIterator, error) {

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

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractRoleRevokedIterator{contract: _ApplicationContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ApplicationContract *ApplicationContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ApplicationContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractRoleRevoked)
				if err := _ApplicationContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ApplicationContract *ApplicationContractFilterer) ParseRoleRevoked(log types.Log) (*ApplicationContractRoleRevoked, error) {
	event := new(ApplicationContractRoleRevoked)
	if err := _ApplicationContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApplicationContractStakeEventIterator is returned from FilterStakeEvent and is used to iterate over the raw logs and unpacked data for StakeEvent events raised by the ApplicationContract contract.
type ApplicationContractStakeEventIterator struct {
	Event *ApplicationContractStakeEvent // Event containing the contract specifics and raw log

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
func (it *ApplicationContractStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractStakeEvent)
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
		it.Event = new(ApplicationContractStakeEvent)
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
func (it *ApplicationContractStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractStakeEvent represents a StakeEvent event raised by the ApplicationContract contract.
type ApplicationContractStakeEvent struct {
	Account common.Address
	Stake   ApplicationStakeStruct
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeEvent is a free log retrieval operation binding the contract event 0x278a86cadccd34d129a359367ad378aae3808cfbe93b8a7f9d4fce6638cd2b87.
//
// Solidity: event StakeEvent(address indexed account, (uint256,uint256) stake)
func (_ApplicationContract *ApplicationContractFilterer) FilterStakeEvent(opts *bind.FilterOpts, account []common.Address) (*ApplicationContractStakeEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "StakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractStakeEventIterator{contract: _ApplicationContract.contract, event: "StakeEvent", logs: logs, sub: sub}, nil
}

// WatchStakeEvent is a free log subscription operation binding the contract event 0x278a86cadccd34d129a359367ad378aae3808cfbe93b8a7f9d4fce6638cd2b87.
//
// Solidity: event StakeEvent(address indexed account, (uint256,uint256) stake)
func (_ApplicationContract *ApplicationContractFilterer) WatchStakeEvent(opts *bind.WatchOpts, sink chan<- *ApplicationContractStakeEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "StakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractStakeEvent)
				if err := _ApplicationContract.contract.UnpackLog(event, "StakeEvent", log); err != nil {
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

// ParseStakeEvent is a log parse operation binding the contract event 0x278a86cadccd34d129a359367ad378aae3808cfbe93b8a7f9d4fce6638cd2b87.
//
// Solidity: event StakeEvent(address indexed account, (uint256,uint256) stake)
func (_ApplicationContract *ApplicationContractFilterer) ParseStakeEvent(log types.Log) (*ApplicationContractStakeEvent, error) {
	event := new(ApplicationContractStakeEvent)
	if err := _ApplicationContract.contract.UnpackLog(event, "StakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApplicationContractUnStakeEventIterator is returned from FilterUnStakeEvent and is used to iterate over the raw logs and unpacked data for UnStakeEvent events raised by the ApplicationContract contract.
type ApplicationContractUnStakeEventIterator struct {
	Event *ApplicationContractUnStakeEvent // Event containing the contract specifics and raw log

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
func (it *ApplicationContractUnStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApplicationContractUnStakeEvent)
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
		it.Event = new(ApplicationContractUnStakeEvent)
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
func (it *ApplicationContractUnStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApplicationContractUnStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApplicationContractUnStakeEvent represents a UnStakeEvent event raised by the ApplicationContract contract.
type ApplicationContractUnStakeEvent struct {
	Account common.Address
	Stake   ApplicationStakeStruct
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnStakeEvent is a free log retrieval operation binding the contract event 0xb8075e85c8e8c717443eca24e396af055e00019ff8c88d1cf27aeb39cba8353f.
//
// Solidity: event UnStakeEvent(address indexed account, (uint256,uint256) stake)
func (_ApplicationContract *ApplicationContractFilterer) FilterUnStakeEvent(opts *bind.FilterOpts, account []common.Address) (*ApplicationContractUnStakeEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ApplicationContract.contract.FilterLogs(opts, "UnStakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &ApplicationContractUnStakeEventIterator{contract: _ApplicationContract.contract, event: "UnStakeEvent", logs: logs, sub: sub}, nil
}

// WatchUnStakeEvent is a free log subscription operation binding the contract event 0xb8075e85c8e8c717443eca24e396af055e00019ff8c88d1cf27aeb39cba8353f.
//
// Solidity: event UnStakeEvent(address indexed account, (uint256,uint256) stake)
func (_ApplicationContract *ApplicationContractFilterer) WatchUnStakeEvent(opts *bind.WatchOpts, sink chan<- *ApplicationContractUnStakeEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ApplicationContract.contract.WatchLogs(opts, "UnStakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApplicationContractUnStakeEvent)
				if err := _ApplicationContract.contract.UnpackLog(event, "UnStakeEvent", log); err != nil {
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

// ParseUnStakeEvent is a log parse operation binding the contract event 0xb8075e85c8e8c717443eca24e396af055e00019ff8c88d1cf27aeb39cba8353f.
//
// Solidity: event UnStakeEvent(address indexed account, (uint256,uint256) stake)
func (_ApplicationContract *ApplicationContractFilterer) ParseUnStakeEvent(log types.Log) (*ApplicationContractUnStakeEvent, error) {
	event := new(ApplicationContractUnStakeEvent)
	if err := _ApplicationContract.contract.UnpackLog(event, "UnStakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
