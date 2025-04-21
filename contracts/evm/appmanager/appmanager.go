// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package appmanager

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

// AppManagerClaim is an auto generated low-level Go binding around an user-defined struct.
type AppManagerClaim struct {
	Validator  []byte
	ClaimData  []AppManagerRewardClaimData
	Cycle      *big.Int
	Index      *big.Int
	Signers    []LibSecp256k1Point
	Commitment common.Address
	Signature  []byte
	TotalCost  *big.Int
}

// AppManagerLicenseInfo is an auto generated low-level Go binding around an user-defined struct.
type AppManagerLicenseInfo struct {
	Id          *big.Int
	Earned      *big.Int
	IsValidator bool
	DelegatedTo []byte
}

// AppManagerRewardClaimData is an auto generated low-level Go binding around an user-defined struct.
type AppManagerRewardClaimData struct {
	AppId  [16]byte
	Amount *big.Int
}

// AppManagerStakeStruct is an auto generated low-level Go binding around an user-defined struct.
type AppManagerStakeStruct struct {
	Amount    *big.Int
	Timestamp *big.Int
}

// LibSecp256k1Point is an auto generated low-level Go binding around an user-defined struct.
type LibSecp256k1Point struct {
	X *big.Int
	Y *big.Int
}

// AppmanagerContractMetaData contains all meta data concerning the AppmanagerContract contract.
var AppmanagerContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAppManager.StakeStruct\",\"name\":\"stake\",\"type\":\"tuple\"}],\"name\":\"StakeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAppManager.StakeStruct\",\"name\":\"stake\",\"type\":\"tuple\"}],\"name\":\"UnStakeEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accountCreditCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"activateApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"activateAppWithValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"prevCycleReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allTimeReward\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidator\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"delegatedTo\",\"type\":\"bytes\"}],\"internalType\":\"structAppManager.LicenseInfo[]\",\"name\":\"info\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"claimable\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"appActivationCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"}],\"name\":\"appBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"appBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"appCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"name\":\"appDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"appStakerBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structLibSecp256k1.Point[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"enableWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAppAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structLibSecp256k1.Point[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"name\":\"getClaimHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cycleNumLicences\",\"type\":\"uint256\"}],\"name\":\"getMinSignerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"}],\"name\":\"hashRewardData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sentryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorNodeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_appActivationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorAppActivationRewardPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformAppActivationRewardPercent\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"network\",\"outputs\":[{\"internalType\":\"contractINetwork\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"processedClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proofProviderRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sentryBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sentryContract\",\"outputs\":[{\"internalType\":\"contractINodeContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentryCycleRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentryLicenseRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setAppActivationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStakable\",\"type\":\"uint256\"}],\"name\":\"setMinStakable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setNetworkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentX10\",\"type\":\"uint256\"}],\"name\":\"setPlatformAppRewardPercentX10\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseReward\",\"type\":\"uint256\"}],\"name\":\"setSentryBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setSentryNodeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentX10\",\"type\":\"uint256\"}],\"name\":\"setValidatorAppRewardPercentX10\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseReward\",\"type\":\"uint256\"}],\"name\":\"setValidatorBaseReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"setValidatorNodeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"}],\"name\":\"unStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unstakeOrders\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorBaseReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorCycleRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorNodeContract\",\"outputs\":[{\"internalType\":\"contractINodeContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes16\",\"name\":\"appId\",\"type\":\"bytes16\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.RewardClaimData[]\",\"name\":\"claimData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cycle\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structLibSecp256k1.Point[]\",\"name\":\"signers\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"commitment\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalCost\",\"type\":\"uint256\"}],\"internalType\":\"structAppManager.Claim\",\"name\":\"claim\",\"type\":\"tuple\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AppmanagerContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AppmanagerContractMetaData.ABI instead.
var AppmanagerContractABI = AppmanagerContractMetaData.ABI

// AppmanagerContract is an auto generated Go binding around an Ethereum contract.
type AppmanagerContract struct {
	AppmanagerContractCaller     // Read-only binding to the contract
	AppmanagerContractTransactor // Write-only binding to the contract
	AppmanagerContractFilterer   // Log filterer for contract events
}

// AppmanagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppmanagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppmanagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppmanagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppmanagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppmanagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppmanagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppmanagerContractSession struct {
	Contract     *AppmanagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AppmanagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppmanagerContractCallerSession struct {
	Contract *AppmanagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AppmanagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppmanagerContractTransactorSession struct {
	Contract     *AppmanagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AppmanagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppmanagerContractRaw struct {
	Contract *AppmanagerContract // Generic contract binding to access the raw methods on
}

// AppmanagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppmanagerContractCallerRaw struct {
	Contract *AppmanagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// AppmanagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppmanagerContractTransactorRaw struct {
	Contract *AppmanagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppmanagerContract creates a new instance of AppmanagerContract, bound to a specific deployed contract.
func NewAppmanagerContract(address common.Address, backend bind.ContractBackend) (*AppmanagerContract, error) {
	contract, err := bindAppmanagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContract{AppmanagerContractCaller: AppmanagerContractCaller{contract: contract}, AppmanagerContractTransactor: AppmanagerContractTransactor{contract: contract}, AppmanagerContractFilterer: AppmanagerContractFilterer{contract: contract}}, nil
}

// NewAppmanagerContractCaller creates a new read-only instance of AppmanagerContract, bound to a specific deployed contract.
func NewAppmanagerContractCaller(address common.Address, caller bind.ContractCaller) (*AppmanagerContractCaller, error) {
	contract, err := bindAppmanagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractCaller{contract: contract}, nil
}

// NewAppmanagerContractTransactor creates a new write-only instance of AppmanagerContract, bound to a specific deployed contract.
func NewAppmanagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AppmanagerContractTransactor, error) {
	contract, err := bindAppmanagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractTransactor{contract: contract}, nil
}

// NewAppmanagerContractFilterer creates a new log filterer instance of AppmanagerContract, bound to a specific deployed contract.
func NewAppmanagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AppmanagerContractFilterer, error) {
	contract, err := bindAppmanagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractFilterer{contract: contract}, nil
}

// bindAppmanagerContract binds a generic wrapper to an already deployed contract.
func bindAppmanagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppmanagerContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppmanagerContract *AppmanagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppmanagerContract.Contract.AppmanagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppmanagerContract *AppmanagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.AppmanagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppmanagerContract *AppmanagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.AppmanagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppmanagerContract *AppmanagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppmanagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppmanagerContract *AppmanagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppmanagerContract *AppmanagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AppmanagerContract.Contract.DEFAULTADMINROLE(&_AppmanagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AppmanagerContract.Contract.DEFAULTADMINROLE(&_AppmanagerContract.CallOpts)
}

// AccountCreditCommitment is a free data retrieval call binding the contract method 0x547df96c.
//
// Solidity: function accountCreditCommitment(bytes32 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) AccountCreditCommitment(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "accountCreditCommitment", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AccountCreditCommitment is a free data retrieval call binding the contract method 0x547df96c.
//
// Solidity: function accountCreditCommitment(bytes32 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) AccountCreditCommitment(arg0 [32]byte) (bool, error) {
	return _AppmanagerContract.Contract.AccountCreditCommitment(&_AppmanagerContract.CallOpts, arg0)
}

// AccountCreditCommitment is a free data retrieval call binding the contract method 0x547df96c.
//
// Solidity: function accountCreditCommitment(bytes32 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) AccountCreditCommitment(arg0 [32]byte) (bool, error) {
	return _AppmanagerContract.Contract.AccountCreditCommitment(&_AppmanagerContract.CallOpts, arg0)
}

// AddressInfo is a free data retrieval call binding the contract method 0x2126fcb2.
//
// Solidity: function addressInfo(address addr) view returns(uint256 prevCycleReward, uint256 allTimeReward, (uint256,uint256,bool,bytes)[] info, uint256 claimable)
func (_AppmanagerContract *AppmanagerContractCaller) AddressInfo(opts *bind.CallOpts, addr common.Address) (struct {
	PrevCycleReward *big.Int
	AllTimeReward   *big.Int
	Info            []AppManagerLicenseInfo
	Claimable       *big.Int
}, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "addressInfo", addr)

	outstruct := new(struct {
		PrevCycleReward *big.Int
		AllTimeReward   *big.Int
		Info            []AppManagerLicenseInfo
		Claimable       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrevCycleReward = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AllTimeReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Info = *abi.ConvertType(out[2], new([]AppManagerLicenseInfo)).(*[]AppManagerLicenseInfo)
	outstruct.Claimable = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AddressInfo is a free data retrieval call binding the contract method 0x2126fcb2.
//
// Solidity: function addressInfo(address addr) view returns(uint256 prevCycleReward, uint256 allTimeReward, (uint256,uint256,bool,bytes)[] info, uint256 claimable)
func (_AppmanagerContract *AppmanagerContractSession) AddressInfo(addr common.Address) (struct {
	PrevCycleReward *big.Int
	AllTimeReward   *big.Int
	Info            []AppManagerLicenseInfo
	Claimable       *big.Int
}, error) {
	return _AppmanagerContract.Contract.AddressInfo(&_AppmanagerContract.CallOpts, addr)
}

// AddressInfo is a free data retrieval call binding the contract method 0x2126fcb2.
//
// Solidity: function addressInfo(address addr) view returns(uint256 prevCycleReward, uint256 allTimeReward, (uint256,uint256,bool,bytes)[] info, uint256 claimable)
func (_AppmanagerContract *AppmanagerContractCallerSession) AddressInfo(addr common.Address) (struct {
	PrevCycleReward *big.Int
	AllTimeReward   *big.Int
	Info            []AppManagerLicenseInfo
	Claimable       *big.Int
}, error) {
	return _AppmanagerContract.Contract.AddressInfo(&_AppmanagerContract.CallOpts, addr)
}

// AppActivationCommitment is a free data retrieval call binding the contract method 0xb308e3e6.
//
// Solidity: function appActivationCommitment(bytes32 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) AppActivationCommitment(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "appActivationCommitment", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AppActivationCommitment is a free data retrieval call binding the contract method 0xb308e3e6.
//
// Solidity: function appActivationCommitment(bytes32 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) AppActivationCommitment(arg0 [32]byte) (bool, error) {
	return _AppmanagerContract.Contract.AppActivationCommitment(&_AppmanagerContract.CallOpts, arg0)
}

// AppActivationCommitment is a free data retrieval call binding the contract method 0xb308e3e6.
//
// Solidity: function appActivationCommitment(bytes32 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) AppActivationCommitment(arg0 [32]byte) (bool, error) {
	return _AppmanagerContract.Contract.AppActivationCommitment(&_AppmanagerContract.CallOpts, arg0)
}

// AppBalance is a free data retrieval call binding the contract method 0x4d6c0ee2.
//
// Solidity: function appBalance(bytes16 appId) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) AppBalance(opts *bind.CallOpts, appId [16]byte) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "appBalance", appId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppBalance is a free data retrieval call binding the contract method 0x4d6c0ee2.
//
// Solidity: function appBalance(bytes16 appId) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) AppBalance(appId [16]byte) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppBalance(&_AppmanagerContract.CallOpts, appId)
}

// AppBalance is a free data retrieval call binding the contract method 0x4d6c0ee2.
//
// Solidity: function appBalance(bytes16 appId) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) AppBalance(appId [16]byte) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppBalance(&_AppmanagerContract.CallOpts, appId)
}

// AppBalances is a free data retrieval call binding the contract method 0xe97da002.
//
// Solidity: function appBalances(bytes16 , address , uint256 ) view returns(uint256 amount, uint256 timestamp)
func (_AppmanagerContract *AppmanagerContractCaller) AppBalances(opts *bind.CallOpts, arg0 [16]byte, arg1 common.Address, arg2 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "appBalances", arg0, arg1, arg2)

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

// AppBalances is a free data retrieval call binding the contract method 0xe97da002.
//
// Solidity: function appBalances(bytes16 , address , uint256 ) view returns(uint256 amount, uint256 timestamp)
func (_AppmanagerContract *AppmanagerContractSession) AppBalances(arg0 [16]byte, arg1 common.Address, arg2 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _AppmanagerContract.Contract.AppBalances(&_AppmanagerContract.CallOpts, arg0, arg1, arg2)
}

// AppBalances is a free data retrieval call binding the contract method 0xe97da002.
//
// Solidity: function appBalances(bytes16 , address , uint256 ) view returns(uint256 amount, uint256 timestamp)
func (_AppmanagerContract *AppmanagerContractCallerSession) AppBalances(arg0 [16]byte, arg1 common.Address, arg2 *big.Int) (struct {
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _AppmanagerContract.Contract.AppBalances(&_AppmanagerContract.CallOpts, arg0, arg1, arg2)
}

// AppCredit is a free data retrieval call binding the contract method 0x3eef4d28.
//
// Solidity: function appCredit(bytes16 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) AppCredit(opts *bind.CallOpts, arg0 [16]byte) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "appCredit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppCredit is a free data retrieval call binding the contract method 0x3eef4d28.
//
// Solidity: function appCredit(bytes16 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) AppCredit(arg0 [16]byte) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppCredit(&_AppmanagerContract.CallOpts, arg0)
}

// AppCredit is a free data retrieval call binding the contract method 0x3eef4d28.
//
// Solidity: function appCredit(bytes16 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) AppCredit(arg0 [16]byte) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppCredit(&_AppmanagerContract.CallOpts, arg0)
}

// AppDebt is a free data retrieval call binding the contract method 0xadb97faf.
//
// Solidity: function appDebt(bytes16 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) AppDebt(opts *bind.CallOpts, arg0 [16]byte) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "appDebt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppDebt is a free data retrieval call binding the contract method 0xadb97faf.
//
// Solidity: function appDebt(bytes16 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) AppDebt(arg0 [16]byte) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppDebt(&_AppmanagerContract.CallOpts, arg0)
}

// AppDebt is a free data retrieval call binding the contract method 0xadb97faf.
//
// Solidity: function appDebt(bytes16 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) AppDebt(arg0 [16]byte) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppDebt(&_AppmanagerContract.CallOpts, arg0)
}

// AppStakerBalances is a free data retrieval call binding the contract method 0xdf333c28.
//
// Solidity: function appStakerBalances(bytes16 , address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) AppStakerBalances(opts *bind.CallOpts, arg0 [16]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "appStakerBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AppStakerBalances is a free data retrieval call binding the contract method 0xdf333c28.
//
// Solidity: function appStakerBalances(bytes16 , address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) AppStakerBalances(arg0 [16]byte, arg1 common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppStakerBalances(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// AppStakerBalances is a free data retrieval call binding the contract method 0xdf333c28.
//
// Solidity: function appStakerBalances(bytes16 , address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) AppStakerBalances(arg0 [16]byte, arg1 common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.AppStakerBalances(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// GetAppAccountBalance is a free data retrieval call binding the contract method 0x6432c4d0.
//
// Solidity: function getAppAccountBalance(bytes16 appId, address addr) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) GetAppAccountBalance(opts *bind.CallOpts, appId [16]byte, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "getAppAccountBalance", appId, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAppAccountBalance is a free data retrieval call binding the contract method 0x6432c4d0.
//
// Solidity: function getAppAccountBalance(bytes16 appId, address addr) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) GetAppAccountBalance(appId [16]byte, addr common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.GetAppAccountBalance(&_AppmanagerContract.CallOpts, appId, addr)
}

// GetAppAccountBalance is a free data retrieval call binding the contract method 0x6432c4d0.
//
// Solidity: function getAppAccountBalance(bytes16 appId, address addr) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) GetAppAccountBalance(appId [16]byte, addr common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.GetAppAccountBalance(&_AppmanagerContract.CallOpts, appId, addr)
}

// GetClaimHash is a free data retrieval call binding the contract method 0xf6b85a9a.
//
// Solidity: function getClaimHash((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractCaller) GetClaimHash(opts *bind.CallOpts, claim AppManagerClaim) ([32]byte, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "getClaimHash", claim)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetClaimHash is a free data retrieval call binding the contract method 0xf6b85a9a.
//
// Solidity: function getClaimHash((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractSession) GetClaimHash(claim AppManagerClaim) ([32]byte, error) {
	return _AppmanagerContract.Contract.GetClaimHash(&_AppmanagerContract.CallOpts, claim)
}

// GetClaimHash is a free data retrieval call binding the contract method 0xf6b85a9a.
//
// Solidity: function getClaimHash((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractCallerSession) GetClaimHash(claim AppManagerClaim) ([32]byte, error) {
	return _AppmanagerContract.Contract.GetClaimHash(&_AppmanagerContract.CallOpts, claim)
}

// GetMinSignerCount is a free data retrieval call binding the contract method 0x9cee2fe3.
//
// Solidity: function getMinSignerCount(uint256 cycleNumLicences) pure returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) GetMinSignerCount(opts *bind.CallOpts, cycleNumLicences *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "getMinSignerCount", cycleNumLicences)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinSignerCount is a free data retrieval call binding the contract method 0x9cee2fe3.
//
// Solidity: function getMinSignerCount(uint256 cycleNumLicences) pure returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) GetMinSignerCount(cycleNumLicences *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.GetMinSignerCount(&_AppmanagerContract.CallOpts, cycleNumLicences)
}

// GetMinSignerCount is a free data retrieval call binding the contract method 0x9cee2fe3.
//
// Solidity: function getMinSignerCount(uint256 cycleNumLicences) pure returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) GetMinSignerCount(cycleNumLicences *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.GetMinSignerCount(&_AppmanagerContract.CallOpts, cycleNumLicences)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AppmanagerContract.Contract.GetRoleAdmin(&_AppmanagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AppmanagerContract *AppmanagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AppmanagerContract.Contract.GetRoleAdmin(&_AppmanagerContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AppmanagerContract.Contract.HasRole(&_AppmanagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AppmanagerContract.Contract.HasRole(&_AppmanagerContract.CallOpts, role, account)
}

// HashRewardData is a free data retrieval call binding the contract method 0x5a9a7b87.
//
// Solidity: function hashRewardData((bytes16,uint256)[] claimData) pure returns(bytes32 hash)
func (_AppmanagerContract *AppmanagerContractCaller) HashRewardData(opts *bind.CallOpts, claimData []AppManagerRewardClaimData) ([32]byte, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "hashRewardData", claimData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashRewardData is a free data retrieval call binding the contract method 0x5a9a7b87.
//
// Solidity: function hashRewardData((bytes16,uint256)[] claimData) pure returns(bytes32 hash)
func (_AppmanagerContract *AppmanagerContractSession) HashRewardData(claimData []AppManagerRewardClaimData) ([32]byte, error) {
	return _AppmanagerContract.Contract.HashRewardData(&_AppmanagerContract.CallOpts, claimData)
}

// HashRewardData is a free data retrieval call binding the contract method 0x5a9a7b87.
//
// Solidity: function hashRewardData((bytes16,uint256)[] claimData) pure returns(bytes32 hash)
func (_AppmanagerContract *AppmanagerContractCallerSession) HashRewardData(claimData []AppManagerRewardClaimData) ([32]byte, error) {
	return _AppmanagerContract.Contract.HashRewardData(&_AppmanagerContract.CallOpts, claimData)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) Locked() (bool, error) {
	return _AppmanagerContract.Contract.Locked(&_AppmanagerContract.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) Locked() (bool, error) {
	return _AppmanagerContract.Contract.Locked(&_AppmanagerContract.CallOpts)
}

// MinStakable is a free data retrieval call binding the contract method 0xc0d41476.
//
// Solidity: function minStakable() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) MinStakable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "minStakable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakable is a free data retrieval call binding the contract method 0xc0d41476.
//
// Solidity: function minStakable() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) MinStakable() (*big.Int, error) {
	return _AppmanagerContract.Contract.MinStakable(&_AppmanagerContract.CallOpts)
}

// MinStakable is a free data retrieval call binding the contract method 0xc0d41476.
//
// Solidity: function minStakable() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) MinStakable() (*big.Int, error) {
	return _AppmanagerContract.Contract.MinStakable(&_AppmanagerContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_AppmanagerContract *AppmanagerContractCaller) Network(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "network")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_AppmanagerContract *AppmanagerContractSession) Network() (common.Address, error) {
	return _AppmanagerContract.Contract.Network(&_AppmanagerContract.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns(address)
func (_AppmanagerContract *AppmanagerContractCallerSession) Network() (common.Address, error) {
	return _AppmanagerContract.Contract.Network(&_AppmanagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppmanagerContract *AppmanagerContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppmanagerContract *AppmanagerContractSession) Owner() (common.Address, error) {
	return _AppmanagerContract.Contract.Owner(&_AppmanagerContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AppmanagerContract *AppmanagerContractCallerSession) Owner() (common.Address, error) {
	return _AppmanagerContract.Contract.Owner(&_AppmanagerContract.CallOpts)
}

// ProcessedClaim is a free data retrieval call binding the contract method 0x2ad6f7f6.
//
// Solidity: function processedClaim(uint256 , bytes , uint256 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) ProcessedClaim(opts *bind.CallOpts, arg0 *big.Int, arg1 []byte, arg2 *big.Int) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "processedClaim", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedClaim is a free data retrieval call binding the contract method 0x2ad6f7f6.
//
// Solidity: function processedClaim(uint256 , bytes , uint256 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) ProcessedClaim(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (bool, error) {
	return _AppmanagerContract.Contract.ProcessedClaim(&_AppmanagerContract.CallOpts, arg0, arg1, arg2)
}

// ProcessedClaim is a free data retrieval call binding the contract method 0x2ad6f7f6.
//
// Solidity: function processedClaim(uint256 , bytes , uint256 ) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) ProcessedClaim(arg0 *big.Int, arg1 []byte, arg2 *big.Int) (bool, error) {
	return _AppmanagerContract.Contract.ProcessedClaim(&_AppmanagerContract.CallOpts, arg0, arg1, arg2)
}

// ProofProviderRewards is a free data retrieval call binding the contract method 0x58f25d13.
//
// Solidity: function proofProviderRewards(address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) ProofProviderRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "proofProviderRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProofProviderRewards is a free data retrieval call binding the contract method 0x58f25d13.
//
// Solidity: function proofProviderRewards(address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) ProofProviderRewards(arg0 common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.ProofProviderRewards(&_AppmanagerContract.CallOpts, arg0)
}

// ProofProviderRewards is a free data retrieval call binding the contract method 0x58f25d13.
//
// Solidity: function proofProviderRewards(address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) ProofProviderRewards(arg0 common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.ProofProviderRewards(&_AppmanagerContract.CallOpts, arg0)
}

// SentryBaseReward is a free data retrieval call binding the contract method 0x25a772a4.
//
// Solidity: function sentryBaseReward() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) SentryBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "sentryBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SentryBaseReward is a free data retrieval call binding the contract method 0x25a772a4.
//
// Solidity: function sentryBaseReward() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) SentryBaseReward() (*big.Int, error) {
	return _AppmanagerContract.Contract.SentryBaseReward(&_AppmanagerContract.CallOpts)
}

// SentryBaseReward is a free data retrieval call binding the contract method 0x25a772a4.
//
// Solidity: function sentryBaseReward() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) SentryBaseReward() (*big.Int, error) {
	return _AppmanagerContract.Contract.SentryBaseReward(&_AppmanagerContract.CallOpts)
}

// SentryContract is a free data retrieval call binding the contract method 0xdbca5c52.
//
// Solidity: function sentryContract() view returns(address)
func (_AppmanagerContract *AppmanagerContractCaller) SentryContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "sentryContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SentryContract is a free data retrieval call binding the contract method 0xdbca5c52.
//
// Solidity: function sentryContract() view returns(address)
func (_AppmanagerContract *AppmanagerContractSession) SentryContract() (common.Address, error) {
	return _AppmanagerContract.Contract.SentryContract(&_AppmanagerContract.CallOpts)
}

// SentryContract is a free data retrieval call binding the contract method 0xdbca5c52.
//
// Solidity: function sentryContract() view returns(address)
func (_AppmanagerContract *AppmanagerContractCallerSession) SentryContract() (common.Address, error) {
	return _AppmanagerContract.Contract.SentryContract(&_AppmanagerContract.CallOpts)
}

// SentryCycleRewards is a free data retrieval call binding the contract method 0x7773b6e9.
//
// Solidity: function sentryCycleRewards(address , uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) SentryCycleRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "sentryCycleRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SentryCycleRewards is a free data retrieval call binding the contract method 0x7773b6e9.
//
// Solidity: function sentryCycleRewards(address , uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) SentryCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.SentryCycleRewards(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// SentryCycleRewards is a free data retrieval call binding the contract method 0x7773b6e9.
//
// Solidity: function sentryCycleRewards(address , uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) SentryCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.SentryCycleRewards(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// SentryLicenseRevenue is a free data retrieval call binding the contract method 0x37b560f3.
//
// Solidity: function sentryLicenseRevenue(uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) SentryLicenseRevenue(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "sentryLicenseRevenue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SentryLicenseRevenue is a free data retrieval call binding the contract method 0x37b560f3.
//
// Solidity: function sentryLicenseRevenue(uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) SentryLicenseRevenue(arg0 *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.SentryLicenseRevenue(&_AppmanagerContract.CallOpts, arg0)
}

// SentryLicenseRevenue is a free data retrieval call binding the contract method 0x37b560f3.
//
// Solidity: function sentryLicenseRevenue(uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) SentryLicenseRevenue(arg0 *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.SentryLicenseRevenue(&_AppmanagerContract.CallOpts, arg0)
}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_AppmanagerContract *AppmanagerContractCaller) StakeAddresses(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "stakeAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_AppmanagerContract *AppmanagerContractSession) StakeAddresses(arg0 common.Address) (common.Address, error) {
	return _AppmanagerContract.Contract.StakeAddresses(&_AppmanagerContract.CallOpts, arg0)
}

// StakeAddresses is a free data retrieval call binding the contract method 0x6a7cce86.
//
// Solidity: function stakeAddresses(address ) view returns(address)
func (_AppmanagerContract *AppmanagerContractCallerSession) StakeAddresses(arg0 common.Address) (common.Address, error) {
	return _AppmanagerContract.Contract.StakeAddresses(&_AppmanagerContract.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AppmanagerContract.Contract.SupportsInterface(&_AppmanagerContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AppmanagerContract.Contract.SupportsInterface(&_AppmanagerContract.CallOpts, interfaceId)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) TotalValueLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "totalValueLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) TotalValueLocked() (*big.Int, error) {
	return _AppmanagerContract.Contract.TotalValueLocked(&_AppmanagerContract.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) TotalValueLocked() (*big.Int, error) {
	return _AppmanagerContract.Contract.TotalValueLocked(&_AppmanagerContract.CallOpts)
}

// UnstakeOrders is a free data retrieval call binding the contract method 0x9d922aca.
//
// Solidity: function unstakeOrders(address , bytes ) view returns(int32)
func (_AppmanagerContract *AppmanagerContractCaller) UnstakeOrders(opts *bind.CallOpts, arg0 common.Address, arg1 []byte) (int32, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "unstakeOrders", arg0, arg1)

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// UnstakeOrders is a free data retrieval call binding the contract method 0x9d922aca.
//
// Solidity: function unstakeOrders(address , bytes ) view returns(int32)
func (_AppmanagerContract *AppmanagerContractSession) UnstakeOrders(arg0 common.Address, arg1 []byte) (int32, error) {
	return _AppmanagerContract.Contract.UnstakeOrders(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// UnstakeOrders is a free data retrieval call binding the contract method 0x9d922aca.
//
// Solidity: function unstakeOrders(address , bytes ) view returns(int32)
func (_AppmanagerContract *AppmanagerContractCallerSession) UnstakeOrders(arg0 common.Address, arg1 []byte) (int32, error) {
	return _AppmanagerContract.Contract.UnstakeOrders(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// ValidatorBaseReward is a free data retrieval call binding the contract method 0x0063b8dd.
//
// Solidity: function validatorBaseReward() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) ValidatorBaseReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "validatorBaseReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorBaseReward is a free data retrieval call binding the contract method 0x0063b8dd.
//
// Solidity: function validatorBaseReward() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) ValidatorBaseReward() (*big.Int, error) {
	return _AppmanagerContract.Contract.ValidatorBaseReward(&_AppmanagerContract.CallOpts)
}

// ValidatorBaseReward is a free data retrieval call binding the contract method 0x0063b8dd.
//
// Solidity: function validatorBaseReward() view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) ValidatorBaseReward() (*big.Int, error) {
	return _AppmanagerContract.Contract.ValidatorBaseReward(&_AppmanagerContract.CallOpts)
}

// ValidatorCycleRewards is a free data retrieval call binding the contract method 0xbf19eb0c.
//
// Solidity: function validatorCycleRewards(address , uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) ValidatorCycleRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "validatorCycleRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCycleRewards is a free data retrieval call binding the contract method 0xbf19eb0c.
//
// Solidity: function validatorCycleRewards(address , uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) ValidatorCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.ValidatorCycleRewards(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// ValidatorCycleRewards is a free data retrieval call binding the contract method 0xbf19eb0c.
//
// Solidity: function validatorCycleRewards(address , uint256 ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) ValidatorCycleRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AppmanagerContract.Contract.ValidatorCycleRewards(&_AppmanagerContract.CallOpts, arg0, arg1)
}

// ValidatorNodeContract is a free data retrieval call binding the contract method 0x71f2053b.
//
// Solidity: function validatorNodeContract() view returns(address)
func (_AppmanagerContract *AppmanagerContractCaller) ValidatorNodeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "validatorNodeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorNodeContract is a free data retrieval call binding the contract method 0x71f2053b.
//
// Solidity: function validatorNodeContract() view returns(address)
func (_AppmanagerContract *AppmanagerContractSession) ValidatorNodeContract() (common.Address, error) {
	return _AppmanagerContract.Contract.ValidatorNodeContract(&_AppmanagerContract.CallOpts)
}

// ValidatorNodeContract is a free data retrieval call binding the contract method 0x71f2053b.
//
// Solidity: function validatorNodeContract() view returns(address)
func (_AppmanagerContract *AppmanagerContractCallerSession) ValidatorNodeContract() (common.Address, error) {
	return _AppmanagerContract.Contract.ValidatorNodeContract(&_AppmanagerContract.CallOpts)
}

// ValidatorRewards is a free data retrieval call binding the contract method 0xb1845c56.
//
// Solidity: function validatorRewards(address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) ValidatorRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "validatorRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRewards is a free data retrieval call binding the contract method 0xb1845c56.
//
// Solidity: function validatorRewards(address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) ValidatorRewards(arg0 common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.ValidatorRewards(&_AppmanagerContract.CallOpts, arg0)
}

// ValidatorRewards is a free data retrieval call binding the contract method 0xb1845c56.
//
// Solidity: function validatorRewards(address ) view returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) ValidatorRewards(arg0 common.Address) (*big.Int, error) {
	return _AppmanagerContract.Contract.ValidatorRewards(&_AppmanagerContract.CallOpts, arg0)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x8b2ed4c8.
//
// Solidity: function verifyClaim((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bool, bytes32)
func (_AppmanagerContract *AppmanagerContractCaller) VerifyClaim(opts *bind.CallOpts, claim AppManagerClaim) (bool, [32]byte, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "verifyClaim", claim)

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
func (_AppmanagerContract *AppmanagerContractSession) VerifyClaim(claim AppManagerClaim) (bool, [32]byte, error) {
	return _AppmanagerContract.Contract.VerifyClaim(&_AppmanagerContract.CallOpts, claim)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x8b2ed4c8.
//
// Solidity: function verifyClaim((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) view returns(bool, bytes32)
func (_AppmanagerContract *AppmanagerContractCallerSession) VerifyClaim(claim AppManagerClaim) (bool, [32]byte, error) {
	return _AppmanagerContract.Contract.VerifyClaim(&_AppmanagerContract.CallOpts, claim)
}

// WithdrawableAmount is a free data retrieval call binding the contract method 0x951303f5.
//
// Solidity: function withdrawableAmount() pure returns(uint256)
func (_AppmanagerContract *AppmanagerContractCaller) WithdrawableAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "withdrawableAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableAmount is a free data retrieval call binding the contract method 0x951303f5.
//
// Solidity: function withdrawableAmount() pure returns(uint256)
func (_AppmanagerContract *AppmanagerContractSession) WithdrawableAmount() (*big.Int, error) {
	return _AppmanagerContract.Contract.WithdrawableAmount(&_AppmanagerContract.CallOpts)
}

// WithdrawableAmount is a free data retrieval call binding the contract method 0x951303f5.
//
// Solidity: function withdrawableAmount() pure returns(uint256)
func (_AppmanagerContract *AppmanagerContractCallerSession) WithdrawableAmount() (*big.Int, error) {
	return _AppmanagerContract.Contract.WithdrawableAmount(&_AppmanagerContract.CallOpts)
}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_AppmanagerContract *AppmanagerContractCaller) WithdrawalEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AppmanagerContract.contract.Call(opts, &out, "withdrawalEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_AppmanagerContract *AppmanagerContractSession) WithdrawalEnabled() (bool, error) {
	return _AppmanagerContract.Contract.WithdrawalEnabled(&_AppmanagerContract.CallOpts)
}

// WithdrawalEnabled is a free data retrieval call binding the contract method 0xf8ea5daf.
//
// Solidity: function withdrawalEnabled() view returns(bool)
func (_AppmanagerContract *AppmanagerContractCallerSession) WithdrawalEnabled() (bool, error) {
	return _AppmanagerContract.Contract.WithdrawalEnabled(&_AppmanagerContract.CallOpts)
}

// ActivateApp is a paid mutator transaction binding the contract method 0x98715334.
//
// Solidity: function activateApp(bytes32 commitment, uint256 salt) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) ActivateApp(opts *bind.TransactOpts, commitment [32]byte, salt *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "activateApp", commitment, salt)
}

// ActivateApp is a paid mutator transaction binding the contract method 0x98715334.
//
// Solidity: function activateApp(bytes32 commitment, uint256 salt) returns()
func (_AppmanagerContract *AppmanagerContractSession) ActivateApp(commitment [32]byte, salt *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.ActivateApp(&_AppmanagerContract.TransactOpts, commitment, salt)
}

// ActivateApp is a paid mutator transaction binding the contract method 0x98715334.
//
// Solidity: function activateApp(bytes32 commitment, uint256 salt) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) ActivateApp(commitment [32]byte, salt *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.ActivateApp(&_AppmanagerContract.TransactOpts, commitment, salt)
}

// ActivateAppWithValidator is a paid mutator transaction binding the contract method 0xc1151557.
//
// Solidity: function activateAppWithValidator(bytes32 commitment, address validator) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) ActivateAppWithValidator(opts *bind.TransactOpts, commitment [32]byte, validator common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "activateAppWithValidator", commitment, validator)
}

// ActivateAppWithValidator is a paid mutator transaction binding the contract method 0xc1151557.
//
// Solidity: function activateAppWithValidator(bytes32 commitment, address validator) returns()
func (_AppmanagerContract *AppmanagerContractSession) ActivateAppWithValidator(commitment [32]byte, validator common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.ActivateAppWithValidator(&_AppmanagerContract.TransactOpts, commitment, validator)
}

// ActivateAppWithValidator is a paid mutator transaction binding the contract method 0xc1151557.
//
// Solidity: function activateAppWithValidator(bytes32 commitment, address validator) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) ActivateAppWithValidator(commitment [32]byte, validator common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.ActivateAppWithValidator(&_AppmanagerContract.TransactOpts, commitment, validator)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x25f3a9c0.
//
// Solidity: function claimReward((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) ClaimReward(opts *bind.TransactOpts, claim AppManagerClaim) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "claimReward", claim)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x25f3a9c0.
//
// Solidity: function claimReward((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) returns()
func (_AppmanagerContract *AppmanagerContractSession) ClaimReward(claim AppManagerClaim) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.ClaimReward(&_AppmanagerContract.TransactOpts, claim)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x25f3a9c0.
//
// Solidity: function claimReward((bytes,(bytes16,uint256)[],uint256,uint256,(uint256,uint256)[],address,bytes,uint256) claim) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) ClaimReward(claim AppManagerClaim) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.ClaimReward(&_AppmanagerContract.TransactOpts, claim)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) EnableWithdrawal(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "enableWithdrawal", _enabled)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_AppmanagerContract *AppmanagerContractSession) EnableWithdrawal(_enabled bool) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.EnableWithdrawal(&_AppmanagerContract.TransactOpts, _enabled)
}

// EnableWithdrawal is a paid mutator transaction binding the contract method 0x5636548f.
//
// Solidity: function enableWithdrawal(bool _enabled) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) EnableWithdrawal(_enabled bool) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.EnableWithdrawal(&_AppmanagerContract.TransactOpts, _enabled)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppmanagerContract *AppmanagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.GrantRole(&_AppmanagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.GrantRole(&_AppmanagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf5377fe.
//
// Solidity: function initialize(address _network, address tokenAddress, address xTokenAddress, address _sentryContract, address _validatorNodeContract, address _feeAddress, uint256 _appActivationFee, uint256 _validatorAppActivationRewardPercent, uint256 _platformAppActivationRewardPercent) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) Initialize(opts *bind.TransactOpts, _network common.Address, tokenAddress common.Address, xTokenAddress common.Address, _sentryContract common.Address, _validatorNodeContract common.Address, _feeAddress common.Address, _appActivationFee *big.Int, _validatorAppActivationRewardPercent *big.Int, _platformAppActivationRewardPercent *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "initialize", _network, tokenAddress, xTokenAddress, _sentryContract, _validatorNodeContract, _feeAddress, _appActivationFee, _validatorAppActivationRewardPercent, _platformAppActivationRewardPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf5377fe.
//
// Solidity: function initialize(address _network, address tokenAddress, address xTokenAddress, address _sentryContract, address _validatorNodeContract, address _feeAddress, uint256 _appActivationFee, uint256 _validatorAppActivationRewardPercent, uint256 _platformAppActivationRewardPercent) returns()
func (_AppmanagerContract *AppmanagerContractSession) Initialize(_network common.Address, tokenAddress common.Address, xTokenAddress common.Address, _sentryContract common.Address, _validatorNodeContract common.Address, _feeAddress common.Address, _appActivationFee *big.Int, _validatorAppActivationRewardPercent *big.Int, _platformAppActivationRewardPercent *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.Initialize(&_AppmanagerContract.TransactOpts, _network, tokenAddress, xTokenAddress, _sentryContract, _validatorNodeContract, _feeAddress, _appActivationFee, _validatorAppActivationRewardPercent, _platformAppActivationRewardPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xdf5377fe.
//
// Solidity: function initialize(address _network, address tokenAddress, address xTokenAddress, address _sentryContract, address _validatorNodeContract, address _feeAddress, uint256 _appActivationFee, uint256 _validatorAppActivationRewardPercent, uint256 _platformAppActivationRewardPercent) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) Initialize(_network common.Address, tokenAddress common.Address, xTokenAddress common.Address, _sentryContract common.Address, _validatorNodeContract common.Address, _feeAddress common.Address, _appActivationFee *big.Int, _validatorAppActivationRewardPercent *big.Int, _platformAppActivationRewardPercent *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.Initialize(&_AppmanagerContract.TransactOpts, _network, tokenAddress, xTokenAddress, _sentryContract, _validatorNodeContract, _feeAddress, _appActivationFee, _validatorAppActivationRewardPercent, _platformAppActivationRewardPercent)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppmanagerContract *AppmanagerContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppmanagerContract *AppmanagerContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppmanagerContract.Contract.RenounceOwnership(&_AppmanagerContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AppmanagerContract.Contract.RenounceOwnership(&_AppmanagerContract.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AppmanagerContract *AppmanagerContractSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.RenounceRole(&_AppmanagerContract.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.RenounceRole(&_AppmanagerContract.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppmanagerContract *AppmanagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.RevokeRole(&_AppmanagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.RevokeRole(&_AppmanagerContract.TransactOpts, role, account)
}

// SetAppActivationFee is a paid mutator transaction binding the contract method 0x3d2693e6.
//
// Solidity: function setAppActivationFee(uint256 fee) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetAppActivationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setAppActivationFee", fee)
}

// SetAppActivationFee is a paid mutator transaction binding the contract method 0x3d2693e6.
//
// Solidity: function setAppActivationFee(uint256 fee) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetAppActivationFee(fee *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetAppActivationFee(&_AppmanagerContract.TransactOpts, fee)
}

// SetAppActivationFee is a paid mutator transaction binding the contract method 0x3d2693e6.
//
// Solidity: function setAppActivationFee(uint256 fee) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetAppActivationFee(fee *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetAppActivationFee(&_AppmanagerContract.TransactOpts, fee)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _addr) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetFeeAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setFeeAddress", _addr)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _addr) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetFeeAddress(_addr common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetFeeAddress(&_AppmanagerContract.TransactOpts, _addr)
}

// SetFeeAddress is a paid mutator transaction binding the contract method 0x8705fcd4.
//
// Solidity: function setFeeAddress(address _addr) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetFeeAddress(_addr common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetFeeAddress(&_AppmanagerContract.TransactOpts, _addr)
}

// SetMinStakable is a paid mutator transaction binding the contract method 0xf65ff7bd.
//
// Solidity: function setMinStakable(uint256 _minStakable) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetMinStakable(opts *bind.TransactOpts, _minStakable *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setMinStakable", _minStakable)
}

// SetMinStakable is a paid mutator transaction binding the contract method 0xf65ff7bd.
//
// Solidity: function setMinStakable(uint256 _minStakable) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetMinStakable(_minStakable *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetMinStakable(&_AppmanagerContract.TransactOpts, _minStakable)
}

// SetMinStakable is a paid mutator transaction binding the contract method 0xf65ff7bd.
//
// Solidity: function setMinStakable(uint256 _minStakable) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetMinStakable(_minStakable *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetMinStakable(&_AppmanagerContract.TransactOpts, _minStakable)
}

// SetNetworkAddress is a paid mutator transaction binding the contract method 0x05f5dc95.
//
// Solidity: function setNetworkAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetNetworkAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setNetworkAddress", add)
}

// SetNetworkAddress is a paid mutator transaction binding the contract method 0x05f5dc95.
//
// Solidity: function setNetworkAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetNetworkAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetNetworkAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetNetworkAddress is a paid mutator transaction binding the contract method 0x05f5dc95.
//
// Solidity: function setNetworkAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetNetworkAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetNetworkAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetPlatformAppRewardPercentX10 is a paid mutator transaction binding the contract method 0x690a3309.
//
// Solidity: function setPlatformAppRewardPercentX10(uint256 percentX10) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetPlatformAppRewardPercentX10(opts *bind.TransactOpts, percentX10 *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setPlatformAppRewardPercentX10", percentX10)
}

// SetPlatformAppRewardPercentX10 is a paid mutator transaction binding the contract method 0x690a3309.
//
// Solidity: function setPlatformAppRewardPercentX10(uint256 percentX10) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetPlatformAppRewardPercentX10(percentX10 *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetPlatformAppRewardPercentX10(&_AppmanagerContract.TransactOpts, percentX10)
}

// SetPlatformAppRewardPercentX10 is a paid mutator transaction binding the contract method 0x690a3309.
//
// Solidity: function setPlatformAppRewardPercentX10(uint256 percentX10) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetPlatformAppRewardPercentX10(percentX10 *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetPlatformAppRewardPercentX10(&_AppmanagerContract.TransactOpts, percentX10)
}

// SetSentryBaseReward is a paid mutator transaction binding the contract method 0x1227d71a.
//
// Solidity: function setSentryBaseReward(uint256 _baseReward) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetSentryBaseReward(opts *bind.TransactOpts, _baseReward *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setSentryBaseReward", _baseReward)
}

// SetSentryBaseReward is a paid mutator transaction binding the contract method 0x1227d71a.
//
// Solidity: function setSentryBaseReward(uint256 _baseReward) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetSentryBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetSentryBaseReward(&_AppmanagerContract.TransactOpts, _baseReward)
}

// SetSentryBaseReward is a paid mutator transaction binding the contract method 0x1227d71a.
//
// Solidity: function setSentryBaseReward(uint256 _baseReward) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetSentryBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetSentryBaseReward(&_AppmanagerContract.TransactOpts, _baseReward)
}

// SetSentryNodeAddress is a paid mutator transaction binding the contract method 0x1e7fc3ea.
//
// Solidity: function setSentryNodeAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetSentryNodeAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setSentryNodeAddress", add)
}

// SetSentryNodeAddress is a paid mutator transaction binding the contract method 0x1e7fc3ea.
//
// Solidity: function setSentryNodeAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetSentryNodeAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetSentryNodeAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetSentryNodeAddress is a paid mutator transaction binding the contract method 0x1e7fc3ea.
//
// Solidity: function setSentryNodeAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetSentryNodeAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetSentryNodeAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetTokenAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setTokenAddress", add)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetTokenAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetTokenAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetTokenAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetTokenAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetValidatorAppRewardPercentX10 is a paid mutator transaction binding the contract method 0x0ed446ff.
//
// Solidity: function setValidatorAppRewardPercentX10(uint256 percentX10) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetValidatorAppRewardPercentX10(opts *bind.TransactOpts, percentX10 *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setValidatorAppRewardPercentX10", percentX10)
}

// SetValidatorAppRewardPercentX10 is a paid mutator transaction binding the contract method 0x0ed446ff.
//
// Solidity: function setValidatorAppRewardPercentX10(uint256 percentX10) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetValidatorAppRewardPercentX10(percentX10 *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetValidatorAppRewardPercentX10(&_AppmanagerContract.TransactOpts, percentX10)
}

// SetValidatorAppRewardPercentX10 is a paid mutator transaction binding the contract method 0x0ed446ff.
//
// Solidity: function setValidatorAppRewardPercentX10(uint256 percentX10) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetValidatorAppRewardPercentX10(percentX10 *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetValidatorAppRewardPercentX10(&_AppmanagerContract.TransactOpts, percentX10)
}

// SetValidatorBaseReward is a paid mutator transaction binding the contract method 0x8b043d02.
//
// Solidity: function setValidatorBaseReward(uint256 _baseReward) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetValidatorBaseReward(opts *bind.TransactOpts, _baseReward *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setValidatorBaseReward", _baseReward)
}

// SetValidatorBaseReward is a paid mutator transaction binding the contract method 0x8b043d02.
//
// Solidity: function setValidatorBaseReward(uint256 _baseReward) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetValidatorBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetValidatorBaseReward(&_AppmanagerContract.TransactOpts, _baseReward)
}

// SetValidatorBaseReward is a paid mutator transaction binding the contract method 0x8b043d02.
//
// Solidity: function setValidatorBaseReward(uint256 _baseReward) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetValidatorBaseReward(_baseReward *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetValidatorBaseReward(&_AppmanagerContract.TransactOpts, _baseReward)
}

// SetValidatorNodeAddress is a paid mutator transaction binding the contract method 0xb791bf39.
//
// Solidity: function setValidatorNodeAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) SetValidatorNodeAddress(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "setValidatorNodeAddress", add)
}

// SetValidatorNodeAddress is a paid mutator transaction binding the contract method 0xb791bf39.
//
// Solidity: function setValidatorNodeAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractSession) SetValidatorNodeAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetValidatorNodeAddress(&_AppmanagerContract.TransactOpts, add)
}

// SetValidatorNodeAddress is a paid mutator transaction binding the contract method 0xb791bf39.
//
// Solidity: function setValidatorNodeAddress(address add) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) SetValidatorNodeAddress(add common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.SetValidatorNodeAddress(&_AppmanagerContract.TransactOpts, add)
}

// Stake is a paid mutator transaction binding the contract method 0xf47cfa50.
//
// Solidity: function stake(bytes16 appId, uint256 amount) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) Stake(opts *bind.TransactOpts, appId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "stake", appId, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xf47cfa50.
//
// Solidity: function stake(bytes16 appId, uint256 amount) returns()
func (_AppmanagerContract *AppmanagerContractSession) Stake(appId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.Stake(&_AppmanagerContract.TransactOpts, appId, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xf47cfa50.
//
// Solidity: function stake(bytes16 appId, uint256 amount) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) Stake(appId [16]byte, amount *big.Int) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.Stake(&_AppmanagerContract.TransactOpts, appId, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppmanagerContract *AppmanagerContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.TransferOwnership(&_AppmanagerContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.TransferOwnership(&_AppmanagerContract.TransactOpts, newOwner)
}

// UnStake is a paid mutator transaction binding the contract method 0x7fd72127.
//
// Solidity: function unStake(bytes16 appId) returns()
func (_AppmanagerContract *AppmanagerContractTransactor) UnStake(opts *bind.TransactOpts, appId [16]byte) (*types.Transaction, error) {
	return _AppmanagerContract.contract.Transact(opts, "unStake", appId)
}

// UnStake is a paid mutator transaction binding the contract method 0x7fd72127.
//
// Solidity: function unStake(bytes16 appId) returns()
func (_AppmanagerContract *AppmanagerContractSession) UnStake(appId [16]byte) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.UnStake(&_AppmanagerContract.TransactOpts, appId)
}

// UnStake is a paid mutator transaction binding the contract method 0x7fd72127.
//
// Solidity: function unStake(bytes16 appId) returns()
func (_AppmanagerContract *AppmanagerContractTransactorSession) UnStake(appId [16]byte) (*types.Transaction, error) {
	return _AppmanagerContract.Contract.UnStake(&_AppmanagerContract.TransactOpts, appId)
}

// AppmanagerContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AppmanagerContract contract.
type AppmanagerContractInitializedIterator struct {
	Event *AppmanagerContractInitialized // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractInitialized)
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
		it.Event = new(AppmanagerContractInitialized)
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
func (it *AppmanagerContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractInitialized represents a Initialized event raised by the AppmanagerContract contract.
type AppmanagerContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*AppmanagerContractInitializedIterator, error) {

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractInitializedIterator{contract: _AppmanagerContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AppmanagerContractInitialized) (event.Subscription, error) {

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractInitialized)
				if err := _AppmanagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseInitialized(log types.Log) (*AppmanagerContractInitialized, error) {
	event := new(AppmanagerContractInitialized)
	if err := _AppmanagerContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppmanagerContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AppmanagerContract contract.
type AppmanagerContractOwnershipTransferredIterator struct {
	Event *AppmanagerContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractOwnershipTransferred)
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
		it.Event = new(AppmanagerContractOwnershipTransferred)
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
func (it *AppmanagerContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractOwnershipTransferred represents a OwnershipTransferred event raised by the AppmanagerContract contract.
type AppmanagerContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AppmanagerContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractOwnershipTransferredIterator{contract: _AppmanagerContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AppmanagerContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractOwnershipTransferred)
				if err := _AppmanagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseOwnershipTransferred(log types.Log) (*AppmanagerContractOwnershipTransferred, error) {
	event := new(AppmanagerContractOwnershipTransferred)
	if err := _AppmanagerContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppmanagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AppmanagerContract contract.
type AppmanagerContractRoleAdminChangedIterator struct {
	Event *AppmanagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractRoleAdminChanged)
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
		it.Event = new(AppmanagerContractRoleAdminChanged)
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
func (it *AppmanagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the AppmanagerContract contract.
type AppmanagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AppmanagerContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractRoleAdminChangedIterator{contract: _AppmanagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AppmanagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractRoleAdminChanged)
				if err := _AppmanagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*AppmanagerContractRoleAdminChanged, error) {
	event := new(AppmanagerContractRoleAdminChanged)
	if err := _AppmanagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppmanagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AppmanagerContract contract.
type AppmanagerContractRoleGrantedIterator struct {
	Event *AppmanagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractRoleGranted)
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
		it.Event = new(AppmanagerContractRoleGranted)
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
func (it *AppmanagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractRoleGranted represents a RoleGranted event raised by the AppmanagerContract contract.
type AppmanagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppmanagerContractRoleGrantedIterator, error) {

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

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractRoleGrantedIterator{contract: _AppmanagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AppmanagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractRoleGranted)
				if err := _AppmanagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseRoleGranted(log types.Log) (*AppmanagerContractRoleGranted, error) {
	event := new(AppmanagerContractRoleGranted)
	if err := _AppmanagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppmanagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AppmanagerContract contract.
type AppmanagerContractRoleRevokedIterator struct {
	Event *AppmanagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractRoleRevoked)
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
		it.Event = new(AppmanagerContractRoleRevoked)
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
func (it *AppmanagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractRoleRevoked represents a RoleRevoked event raised by the AppmanagerContract contract.
type AppmanagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AppmanagerContractRoleRevokedIterator, error) {

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

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractRoleRevokedIterator{contract: _AppmanagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AppmanagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractRoleRevoked)
				if err := _AppmanagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseRoleRevoked(log types.Log) (*AppmanagerContractRoleRevoked, error) {
	event := new(AppmanagerContractRoleRevoked)
	if err := _AppmanagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppmanagerContractStakeEventIterator is returned from FilterStakeEvent and is used to iterate over the raw logs and unpacked data for StakeEvent events raised by the AppmanagerContract contract.
type AppmanagerContractStakeEventIterator struct {
	Event *AppmanagerContractStakeEvent // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractStakeEvent)
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
		it.Event = new(AppmanagerContractStakeEvent)
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
func (it *AppmanagerContractStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractStakeEvent represents a StakeEvent event raised by the AppmanagerContract contract.
type AppmanagerContractStakeEvent struct {
	Account common.Address
	Stake   AppManagerStakeStruct
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeEvent is a free log retrieval operation binding the contract event 0x278a86cadccd34d129a359367ad378aae3808cfbe93b8a7f9d4fce6638cd2b87.
//
// Solidity: event StakeEvent(address indexed account, (uint256,uint256) stake)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterStakeEvent(opts *bind.FilterOpts, account []common.Address) (*AppmanagerContractStakeEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "StakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractStakeEventIterator{contract: _AppmanagerContract.contract, event: "StakeEvent", logs: logs, sub: sub}, nil
}

// WatchStakeEvent is a free log subscription operation binding the contract event 0x278a86cadccd34d129a359367ad378aae3808cfbe93b8a7f9d4fce6638cd2b87.
//
// Solidity: event StakeEvent(address indexed account, (uint256,uint256) stake)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchStakeEvent(opts *bind.WatchOpts, sink chan<- *AppmanagerContractStakeEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "StakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractStakeEvent)
				if err := _AppmanagerContract.contract.UnpackLog(event, "StakeEvent", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseStakeEvent(log types.Log) (*AppmanagerContractStakeEvent, error) {
	event := new(AppmanagerContractStakeEvent)
	if err := _AppmanagerContract.contract.UnpackLog(event, "StakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AppmanagerContractUnStakeEventIterator is returned from FilterUnStakeEvent and is used to iterate over the raw logs and unpacked data for UnStakeEvent events raised by the AppmanagerContract contract.
type AppmanagerContractUnStakeEventIterator struct {
	Event *AppmanagerContractUnStakeEvent // Event containing the contract specifics and raw log

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
func (it *AppmanagerContractUnStakeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AppmanagerContractUnStakeEvent)
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
		it.Event = new(AppmanagerContractUnStakeEvent)
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
func (it *AppmanagerContractUnStakeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AppmanagerContractUnStakeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AppmanagerContractUnStakeEvent represents a UnStakeEvent event raised by the AppmanagerContract contract.
type AppmanagerContractUnStakeEvent struct {
	Account common.Address
	Stake   AppManagerStakeStruct
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnStakeEvent is a free log retrieval operation binding the contract event 0xb8075e85c8e8c717443eca24e396af055e00019ff8c88d1cf27aeb39cba8353f.
//
// Solidity: event UnStakeEvent(address indexed account, (uint256,uint256) stake)
func (_AppmanagerContract *AppmanagerContractFilterer) FilterUnStakeEvent(opts *bind.FilterOpts, account []common.Address) (*AppmanagerContractUnStakeEventIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AppmanagerContract.contract.FilterLogs(opts, "UnStakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return &AppmanagerContractUnStakeEventIterator{contract: _AppmanagerContract.contract, event: "UnStakeEvent", logs: logs, sub: sub}, nil
}

// WatchUnStakeEvent is a free log subscription operation binding the contract event 0xb8075e85c8e8c717443eca24e396af055e00019ff8c88d1cf27aeb39cba8353f.
//
// Solidity: event UnStakeEvent(address indexed account, (uint256,uint256) stake)
func (_AppmanagerContract *AppmanagerContractFilterer) WatchUnStakeEvent(opts *bind.WatchOpts, sink chan<- *AppmanagerContractUnStakeEvent, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AppmanagerContract.contract.WatchLogs(opts, "UnStakeEvent", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AppmanagerContractUnStakeEvent)
				if err := _AppmanagerContract.contract.UnpackLog(event, "UnStakeEvent", log); err != nil {
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
func (_AppmanagerContract *AppmanagerContractFilterer) ParseUnStakeEvent(log types.Log) (*AppmanagerContractUnStakeEvent, error) {
	event := new(AppmanagerContractUnStakeEvent)
	if err := _AppmanagerContract.contract.UnpackLog(event, "UnStakeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
