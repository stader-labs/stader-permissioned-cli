// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ExchangeRate is an auto generated low-level Go binding around an user-defined struct.
type ExchangeRate struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}

// MissedAttestationPenaltyData is an auto generated low-level Go binding around an user-defined struct.
type MissedAttestationPenaltyData struct {
	ReportingBlockNumber *big.Int
	Index                *big.Int
	SortedPubkeys        [][]byte
}

//// RewardsData is an auto generated low-level Go binding around an user-defined struct.
//type RewardsData struct {
//	ReportingBlockNumber *big.Int
//	Index                *big.Int
//	MerkleRoot           [32]byte
//	PoolId               uint8
//	OperatorETHRewards   *big.Int
//	UserETHRewards       *big.Int
//	ProtocolETHRewards   *big.Int
//	OperatorSDRewards    *big.Int
//}

// SDPriceData is an auto generated low-level Go binding around an user-defined struct.
type SDPriceData struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}

// ValidatorStats is an auto generated low-level Go binding around an user-defined struct.
type ValidatorStats struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}

// WithdrawnValidators is an auto generated low-level Go binding around an user-defined struct.
type WithdrawnValidators struct {
	ReportingBlockNumber *big.Int
	NodeRegistry         common.Address
	SortedPubkeys        [][]byte
}

// StaderOracleMetaData contains all meta data concerning the StaderOracle contract.
var StaderOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CooldownNotComplete\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSubmissionFromNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERChangeLimitCrossed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERChangeLimitNotCrossed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERPermissibleChangeOutofBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FrequencyUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InspectionModeActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientTrustedNodes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidERDataSource\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMAPDIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleRootIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPubkeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportingBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUpdate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotTrusted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotATrustedNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PageNumberAlreadyReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportingFutureBlockData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportingPreviousCycleData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateFrequencyNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFrequency\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPORBasedERData\",\"type\":\"bool\"}],\"name\":\"ERDataSourceToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"erInspectionMode\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ERInspectionModeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ExchangeRateSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalEth\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ExchangeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"MissedAttestationPenaltySubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"MissedAttestationPenaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SDPriceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reportedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SDPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SafeModeDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SafeModeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SocializingRewardsMerkleRootSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"}],\"name\":\"SocializingRewardsMerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"TrustedNodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"TrustedNodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updateFrequency\",\"type\":\"uint256\"}],\"name\":\"UpdateFrequencyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erChangeLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedERChangeLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawnValidatorsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"WithdrawnValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ER_CHANGE_MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETHX_ER_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ER_UPDATE_FREQUENCY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TRUSTED_NODES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MISSED_ATTESTATION_PENALTY_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SD_PRICE_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_STATS_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWN_VALIDATORS_UF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"addTrustedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeERInspectionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableERInspectionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableSafeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableSafeMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erChangeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erInspectionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erInspectionModeStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getCurrentRewardsIndexByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getERReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeRate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getMerkleRootReportableBlockByPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMissedAttestationPenaltyReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDPriceInETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSDPriceReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorStats\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"exitingValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"exitingValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"internalType\":\"structValidatorStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorStatsReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawnValidatorReportableBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inspectionModeExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPORFeedBasedERData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isTrustedNode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReportedMAPDIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastReportedSDPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"missedAttestationPenalty\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"removeTrustedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reportingBlockNumberForWithdrawnValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setERUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setMissedAttestationPenaltyUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setSDPriceUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setValidatorStatsUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"}],\"name\":\"setWithdrawnValidatorsUpdateFrequency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalETHXSupply\",\"type\":\"uint256\"}],\"internalType\":\"structExchangeRate\",\"name\":\"_exchangeRate\",\"type\":\"tuple\"}],\"name\":\"submitExchangeRateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedPubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structMissedAttestationPenaltyData\",\"name\":\"_mapd\",\"type\":\"tuple\"}],\"name\":\"submitMissedAttestationPenalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sdPriceInETH\",\"type\":\"uint256\"}],\"internalType\":\"structSDPriceData\",\"name\":\"_sdPriceData\",\"type\":\"tuple\"}],\"name\":\"submitSDPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"poolId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"operatorETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolETHRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"operatorSDRewards\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsData\",\"name\":\"_rewardsData\",\"type\":\"tuple\"}],\"name\":\"submitSocializingRewardsMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"exitingValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"exitingValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"internalType\":\"structValidatorStats\",\"name\":\"_validatorStats\",\"type\":\"tuple\"}],\"name\":\"submitValidatorStats\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"sortedPubkeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structWithdrawnValidators\",\"name\":\"_withdrawnValidators\",\"type\":\"tuple\"}],\"name\":\"submitWithdrawnValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePORFeedBasedERData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedNodesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_erChangeLimit\",\"type\":\"uint256\"}],\"name\":\"updateERChangeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateERFromPORFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"updateFrequencyMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reportingBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"exitingValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"exitedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"slashedValidatorsBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"exitingValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"slashedValidatorsCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StaderOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use StaderOracleMetaData.ABI instead.
var StaderOracleABI = StaderOracleMetaData.ABI

// StaderOracle is an auto generated Go binding around an Ethereum contract.
type StaderOracle struct {
	StaderOracleCaller     // Read-only binding to the contract
	StaderOracleTransactor // Write-only binding to the contract
	StaderOracleFilterer   // Log filterer for contract events
}

// StaderOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaderOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaderOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaderOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaderOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaderOracleSession struct {
	Contract     *StaderOracle     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaderOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaderOracleCallerSession struct {
	Contract *StaderOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StaderOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaderOracleTransactorSession struct {
	Contract     *StaderOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StaderOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaderOracleRaw struct {
	Contract *StaderOracle // Generic contract binding to access the raw methods on
}

// StaderOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaderOracleCallerRaw struct {
	Contract *StaderOracleCaller // Generic read-only contract binding to access the raw methods on
}

// StaderOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaderOracleTransactorRaw struct {
	Contract *StaderOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaderOracle creates a new instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracle(address common.Address, backend bind.ContractBackend) (*StaderOracle, error) {
	contract, err := bindStaderOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StaderOracle{StaderOracleCaller: StaderOracleCaller{contract: contract}, StaderOracleTransactor: StaderOracleTransactor{contract: contract}, StaderOracleFilterer: StaderOracleFilterer{contract: contract}}, nil
}

// NewStaderOracleCaller creates a new read-only instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracleCaller(address common.Address, caller bind.ContractCaller) (*StaderOracleCaller, error) {
	contract, err := bindStaderOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaderOracleCaller{contract: contract}, nil
}

// NewStaderOracleTransactor creates a new write-only instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*StaderOracleTransactor, error) {
	contract, err := bindStaderOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaderOracleTransactor{contract: contract}, nil
}

// NewStaderOracleFilterer creates a new log filterer instance of StaderOracle, bound to a specific deployed contract.
func NewStaderOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*StaderOracleFilterer, error) {
	contract, err := bindStaderOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaderOracleFilterer{contract: contract}, nil
}

// bindStaderOracle binds a generic wrapper to an already deployed contract.
func bindStaderOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaderOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderOracle *StaderOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderOracle.Contract.StaderOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderOracle *StaderOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.Contract.StaderOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderOracle *StaderOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderOracle.Contract.StaderOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StaderOracle *StaderOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StaderOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StaderOracle *StaderOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StaderOracle *StaderOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StaderOracle.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderOracle.Contract.DEFAULTADMINROLE(&_StaderOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StaderOracle.Contract.DEFAULTADMINROLE(&_StaderOracle.CallOpts)
}

// ERCHANGEMAXBPS is a free data retrieval call binding the contract method 0xe0bcb378.
//
// Solidity: function ER_CHANGE_MAX_BPS() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ERCHANGEMAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "ER_CHANGE_MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERCHANGEMAXBPS is a free data retrieval call binding the contract method 0xe0bcb378.
//
// Solidity: function ER_CHANGE_MAX_BPS() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ERCHANGEMAXBPS() (*big.Int, error) {
	return _StaderOracle.Contract.ERCHANGEMAXBPS(&_StaderOracle.CallOpts)
}

// ERCHANGEMAXBPS is a free data retrieval call binding the contract method 0xe0bcb378.
//
// Solidity: function ER_CHANGE_MAX_BPS() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ERCHANGEMAXBPS() (*big.Int, error) {
	return _StaderOracle.Contract.ERCHANGEMAXBPS(&_StaderOracle.CallOpts)
}

// ETHXERUF is a free data retrieval call binding the contract method 0x12710361.
//
// Solidity: function ETHX_ER_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) ETHXERUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "ETHX_ER_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ETHXERUF is a free data retrieval call binding the contract method 0x12710361.
//
// Solidity: function ETHX_ER_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) ETHXERUF() ([32]byte, error) {
	return _StaderOracle.Contract.ETHXERUF(&_StaderOracle.CallOpts)
}

// ETHXERUF is a free data retrieval call binding the contract method 0x12710361.
//
// Solidity: function ETHX_ER_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) ETHXERUF() ([32]byte, error) {
	return _StaderOracle.Contract.ETHXERUF(&_StaderOracle.CallOpts)
}

// MAXERUPDATEFREQUENCY is a free data retrieval call binding the contract method 0xf10b2569.
//
// Solidity: function MAX_ER_UPDATE_FREQUENCY() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) MAXERUPDATEFREQUENCY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MAX_ER_UPDATE_FREQUENCY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXERUPDATEFREQUENCY is a free data retrieval call binding the contract method 0xf10b2569.
//
// Solidity: function MAX_ER_UPDATE_FREQUENCY() view returns(uint256)
func (_StaderOracle *StaderOracleSession) MAXERUPDATEFREQUENCY() (*big.Int, error) {
	return _StaderOracle.Contract.MAXERUPDATEFREQUENCY(&_StaderOracle.CallOpts)
}

// MAXERUPDATEFREQUENCY is a free data retrieval call binding the contract method 0xf10b2569.
//
// Solidity: function MAX_ER_UPDATE_FREQUENCY() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) MAXERUPDATEFREQUENCY() (*big.Int, error) {
	return _StaderOracle.Contract.MAXERUPDATEFREQUENCY(&_StaderOracle.CallOpts)
}

// MINTRUSTEDNODES is a free data retrieval call binding the contract method 0xe36c3140.
//
// Solidity: function MIN_TRUSTED_NODES() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) MINTRUSTEDNODES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MIN_TRUSTED_NODES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTRUSTEDNODES is a free data retrieval call binding the contract method 0xe36c3140.
//
// Solidity: function MIN_TRUSTED_NODES() view returns(uint256)
func (_StaderOracle *StaderOracleSession) MINTRUSTEDNODES() (*big.Int, error) {
	return _StaderOracle.Contract.MINTRUSTEDNODES(&_StaderOracle.CallOpts)
}

// MINTRUSTEDNODES is a free data retrieval call binding the contract method 0xe36c3140.
//
// Solidity: function MIN_TRUSTED_NODES() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) MINTRUSTEDNODES() (*big.Int, error) {
	return _StaderOracle.Contract.MINTRUSTEDNODES(&_StaderOracle.CallOpts)
}

// MISSEDATTESTATIONPENALTYUF is a free data retrieval call binding the contract method 0x052a6840.
//
// Solidity: function MISSED_ATTESTATION_PENALTY_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) MISSEDATTESTATIONPENALTYUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "MISSED_ATTESTATION_PENALTY_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MISSEDATTESTATIONPENALTYUF is a free data retrieval call binding the contract method 0x052a6840.
//
// Solidity: function MISSED_ATTESTATION_PENALTY_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) MISSEDATTESTATIONPENALTYUF() ([32]byte, error) {
	return _StaderOracle.Contract.MISSEDATTESTATIONPENALTYUF(&_StaderOracle.CallOpts)
}

// MISSEDATTESTATIONPENALTYUF is a free data retrieval call binding the contract method 0x052a6840.
//
// Solidity: function MISSED_ATTESTATION_PENALTY_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) MISSEDATTESTATIONPENALTYUF() ([32]byte, error) {
	return _StaderOracle.Contract.MISSEDATTESTATIONPENALTYUF(&_StaderOracle.CallOpts)
}

// SDPRICEUF is a free data retrieval call binding the contract method 0xe2f63392.
//
// Solidity: function SD_PRICE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) SDPRICEUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "SD_PRICE_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SDPRICEUF is a free data retrieval call binding the contract method 0xe2f63392.
//
// Solidity: function SD_PRICE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) SDPRICEUF() ([32]byte, error) {
	return _StaderOracle.Contract.SDPRICEUF(&_StaderOracle.CallOpts)
}

// SDPRICEUF is a free data retrieval call binding the contract method 0xe2f63392.
//
// Solidity: function SD_PRICE_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) SDPRICEUF() ([32]byte, error) {
	return _StaderOracle.Contract.SDPRICEUF(&_StaderOracle.CallOpts)
}

// VALIDATORSTATSUF is a free data retrieval call binding the contract method 0x29f96856.
//
// Solidity: function VALIDATOR_STATS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) VALIDATORSTATSUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "VALIDATOR_STATS_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORSTATSUF is a free data retrieval call binding the contract method 0x29f96856.
//
// Solidity: function VALIDATOR_STATS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) VALIDATORSTATSUF() ([32]byte, error) {
	return _StaderOracle.Contract.VALIDATORSTATSUF(&_StaderOracle.CallOpts)
}

// VALIDATORSTATSUF is a free data retrieval call binding the contract method 0x29f96856.
//
// Solidity: function VALIDATOR_STATS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) VALIDATORSTATSUF() ([32]byte, error) {
	return _StaderOracle.Contract.VALIDATORSTATSUF(&_StaderOracle.CallOpts)
}

// WITHDRAWNVALIDATORSUF is a free data retrieval call binding the contract method 0xe61befa7.
//
// Solidity: function WITHDRAWN_VALIDATORS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) WITHDRAWNVALIDATORSUF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "WITHDRAWN_VALIDATORS_UF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWNVALIDATORSUF is a free data retrieval call binding the contract method 0xe61befa7.
//
// Solidity: function WITHDRAWN_VALIDATORS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleSession) WITHDRAWNVALIDATORSUF() ([32]byte, error) {
	return _StaderOracle.Contract.WITHDRAWNVALIDATORSUF(&_StaderOracle.CallOpts)
}

// WITHDRAWNVALIDATORSUF is a free data retrieval call binding the contract method 0xe61befa7.
//
// Solidity: function WITHDRAWN_VALIDATORS_UF() view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) WITHDRAWNVALIDATORSUF() ([32]byte, error) {
	return _StaderOracle.Contract.WITHDRAWNVALIDATORSUF(&_StaderOracle.CallOpts)
}

// ErChangeLimit is a free data retrieval call binding the contract method 0x97a3a10a.
//
// Solidity: function erChangeLimit() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ErChangeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "erChangeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErChangeLimit is a free data retrieval call binding the contract method 0x97a3a10a.
//
// Solidity: function erChangeLimit() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ErChangeLimit() (*big.Int, error) {
	return _StaderOracle.Contract.ErChangeLimit(&_StaderOracle.CallOpts)
}

// ErChangeLimit is a free data retrieval call binding the contract method 0x97a3a10a.
//
// Solidity: function erChangeLimit() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ErChangeLimit() (*big.Int, error) {
	return _StaderOracle.Contract.ErChangeLimit(&_StaderOracle.CallOpts)
}

// ErInspectionMode is a free data retrieval call binding the contract method 0x342280b3.
//
// Solidity: function erInspectionMode() view returns(bool)
func (_StaderOracle *StaderOracleCaller) ErInspectionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "erInspectionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ErInspectionMode is a free data retrieval call binding the contract method 0x342280b3.
//
// Solidity: function erInspectionMode() view returns(bool)
func (_StaderOracle *StaderOracleSession) ErInspectionMode() (bool, error) {
	return _StaderOracle.Contract.ErInspectionMode(&_StaderOracle.CallOpts)
}

// ErInspectionMode is a free data retrieval call binding the contract method 0x342280b3.
//
// Solidity: function erInspectionMode() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) ErInspectionMode() (bool, error) {
	return _StaderOracle.Contract.ErInspectionMode(&_StaderOracle.CallOpts)
}

// ErInspectionModeStartBlock is a free data retrieval call binding the contract method 0xde271c6d.
//
// Solidity: function erInspectionModeStartBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ErInspectionModeStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "erInspectionModeStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErInspectionModeStartBlock is a free data retrieval call binding the contract method 0xde271c6d.
//
// Solidity: function erInspectionModeStartBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ErInspectionModeStartBlock() (*big.Int, error) {
	return _StaderOracle.Contract.ErInspectionModeStartBlock(&_StaderOracle.CallOpts)
}

// ErInspectionModeStartBlock is a free data retrieval call binding the contract method 0xde271c6d.
//
// Solidity: function erInspectionModeStartBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ErInspectionModeStartBlock() (*big.Int, error) {
	return _StaderOracle.Contract.ErInspectionModeStartBlock(&_StaderOracle.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCaller) ExchangeRate(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "exchangeRate")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		TotalETHBalance      *big.Int
		TotalETHXSupply      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalETHBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalETHXSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleSession) ExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.ExchangeRate(&_StaderOracle.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCallerSession) ExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.ExchangeRate(&_StaderOracle.CallOpts)
}

// GetCurrentRewardsIndexByPoolId is a free data retrieval call binding the contract method 0xd06628ed.
//
// Solidity: function getCurrentRewardsIndexByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetCurrentRewardsIndexByPoolId(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getCurrentRewardsIndexByPoolId", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardsIndexByPoolId is a free data retrieval call binding the contract method 0xd06628ed.
//
// Solidity: function getCurrentRewardsIndexByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetCurrentRewardsIndexByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetCurrentRewardsIndexByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetCurrentRewardsIndexByPoolId is a free data retrieval call binding the contract method 0xd06628ed.
//
// Solidity: function getCurrentRewardsIndexByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetCurrentRewardsIndexByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetCurrentRewardsIndexByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetERReportableBlock is a free data retrieval call binding the contract method 0xb5c25ba6.
//
// Solidity: function getERReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetERReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getERReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetERReportableBlock is a free data retrieval call binding the contract method 0xb5c25ba6.
//
// Solidity: function getERReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetERReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetERReportableBlock(&_StaderOracle.CallOpts)
}

// GetERReportableBlock is a free data retrieval call binding the contract method 0xb5c25ba6.
//
// Solidity: function getERReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetERReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetERReportableBlock(&_StaderOracle.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256))
func (_StaderOracle *StaderOracleCaller) GetExchangeRate(opts *bind.CallOpts) (ExchangeRate, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(ExchangeRate), err
	}

	out0 := *abi.ConvertType(out[0], new(ExchangeRate)).(*ExchangeRate)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256))
func (_StaderOracle *StaderOracleSession) GetExchangeRate() (ExchangeRate, error) {
	return _StaderOracle.Contract.GetExchangeRate(&_StaderOracle.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns((uint256,uint256,uint256))
func (_StaderOracle *StaderOracleCallerSession) GetExchangeRate() (ExchangeRate, error) {
	return _StaderOracle.Contract.GetExchangeRate(&_StaderOracle.CallOpts)
}

// GetMerkleRootReportableBlockByPoolId is a free data retrieval call binding the contract method 0xa0c54387.
//
// Solidity: function getMerkleRootReportableBlockByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetMerkleRootReportableBlockByPoolId(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getMerkleRootReportableBlockByPoolId", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleRootReportableBlockByPoolId is a free data retrieval call binding the contract method 0xa0c54387.
//
// Solidity: function getMerkleRootReportableBlockByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetMerkleRootReportableBlockByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetMerkleRootReportableBlockByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetMerkleRootReportableBlockByPoolId is a free data retrieval call binding the contract method 0xa0c54387.
//
// Solidity: function getMerkleRootReportableBlockByPoolId(uint8 _poolId) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetMerkleRootReportableBlockByPoolId(_poolId uint8) (*big.Int, error) {
	return _StaderOracle.Contract.GetMerkleRootReportableBlockByPoolId(&_StaderOracle.CallOpts, _poolId)
}

// GetMissedAttestationPenaltyReportableBlock is a free data retrieval call binding the contract method 0xa71b3907.
//
// Solidity: function getMissedAttestationPenaltyReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetMissedAttestationPenaltyReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getMissedAttestationPenaltyReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMissedAttestationPenaltyReportableBlock is a free data retrieval call binding the contract method 0xa71b3907.
//
// Solidity: function getMissedAttestationPenaltyReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetMissedAttestationPenaltyReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetMissedAttestationPenaltyReportableBlock(&_StaderOracle.CallOpts)
}

// GetMissedAttestationPenaltyReportableBlock is a free data retrieval call binding the contract method 0xa71b3907.
//
// Solidity: function getMissedAttestationPenaltyReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetMissedAttestationPenaltyReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetMissedAttestationPenaltyReportableBlock(&_StaderOracle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderOracle *StaderOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderOracle *StaderOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderOracle.Contract.GetRoleAdmin(&_StaderOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StaderOracle *StaderOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StaderOracle.Contract.GetRoleAdmin(&_StaderOracle.CallOpts, role)
}

// GetSDPriceInETH is a free data retrieval call binding the contract method 0xa6870e5b.
//
// Solidity: function getSDPriceInETH() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetSDPriceInETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getSDPriceInETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSDPriceInETH is a free data retrieval call binding the contract method 0xa6870e5b.
//
// Solidity: function getSDPriceInETH() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetSDPriceInETH() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceInETH(&_StaderOracle.CallOpts)
}

// GetSDPriceInETH is a free data retrieval call binding the contract method 0xa6870e5b.
//
// Solidity: function getSDPriceInETH() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetSDPriceInETH() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceInETH(&_StaderOracle.CallOpts)
}

// GetSDPriceReportableBlock is a free data retrieval call binding the contract method 0xfc8b821c.
//
// Solidity: function getSDPriceReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetSDPriceReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getSDPriceReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSDPriceReportableBlock is a free data retrieval call binding the contract method 0xfc8b821c.
//
// Solidity: function getSDPriceReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetSDPriceReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceReportableBlock(&_StaderOracle.CallOpts)
}

// GetSDPriceReportableBlock is a free data retrieval call binding the contract method 0xfc8b821c.
//
// Solidity: function getSDPriceReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetSDPriceReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetSDPriceReportableBlock(&_StaderOracle.CallOpts)
}

// GetValidatorStats is a free data retrieval call binding the contract method 0x3e23a827.
//
// Solidity: function getValidatorStats() view returns((uint256,uint128,uint128,uint128,uint32,uint32,uint32))
func (_StaderOracle *StaderOracleCaller) GetValidatorStats(opts *bind.CallOpts) (ValidatorStats, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getValidatorStats")

	if err != nil {
		return *new(ValidatorStats), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorStats)).(*ValidatorStats)

	return out0, err

}

// GetValidatorStats is a free data retrieval call binding the contract method 0x3e23a827.
//
// Solidity: function getValidatorStats() view returns((uint256,uint128,uint128,uint128,uint32,uint32,uint32))
func (_StaderOracle *StaderOracleSession) GetValidatorStats() (ValidatorStats, error) {
	return _StaderOracle.Contract.GetValidatorStats(&_StaderOracle.CallOpts)
}

// GetValidatorStats is a free data retrieval call binding the contract method 0x3e23a827.
//
// Solidity: function getValidatorStats() view returns((uint256,uint128,uint128,uint128,uint32,uint32,uint32))
func (_StaderOracle *StaderOracleCallerSession) GetValidatorStats() (ValidatorStats, error) {
	return _StaderOracle.Contract.GetValidatorStats(&_StaderOracle.CallOpts)
}

// GetValidatorStatsReportableBlock is a free data retrieval call binding the contract method 0x49115a2e.
//
// Solidity: function getValidatorStatsReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetValidatorStatsReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getValidatorStatsReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorStatsReportableBlock is a free data retrieval call binding the contract method 0x49115a2e.
//
// Solidity: function getValidatorStatsReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetValidatorStatsReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetValidatorStatsReportableBlock(&_StaderOracle.CallOpts)
}

// GetValidatorStatsReportableBlock is a free data retrieval call binding the contract method 0x49115a2e.
//
// Solidity: function getValidatorStatsReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetValidatorStatsReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetValidatorStatsReportableBlock(&_StaderOracle.CallOpts)
}

// GetWithdrawnValidatorReportableBlock is a free data retrieval call binding the contract method 0x5063b5bd.
//
// Solidity: function getWithdrawnValidatorReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) GetWithdrawnValidatorReportableBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "getWithdrawnValidatorReportableBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawnValidatorReportableBlock is a free data retrieval call binding the contract method 0x5063b5bd.
//
// Solidity: function getWithdrawnValidatorReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleSession) GetWithdrawnValidatorReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetWithdrawnValidatorReportableBlock(&_StaderOracle.CallOpts)
}

// GetWithdrawnValidatorReportableBlock is a free data retrieval call binding the contract method 0x5063b5bd.
//
// Solidity: function getWithdrawnValidatorReportableBlock() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) GetWithdrawnValidatorReportableBlock() (*big.Int, error) {
	return _StaderOracle.Contract.GetWithdrawnValidatorReportableBlock(&_StaderOracle.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderOracle *StaderOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderOracle *StaderOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderOracle.Contract.HasRole(&_StaderOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StaderOracle.Contract.HasRole(&_StaderOracle.CallOpts, role, account)
}

// InspectionModeExchangeRate is a free data retrieval call binding the contract method 0xb940a003.
//
// Solidity: function inspectionModeExchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCaller) InspectionModeExchangeRate(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "inspectionModeExchangeRate")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		TotalETHBalance      *big.Int
		TotalETHXSupply      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalETHBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalETHXSupply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// InspectionModeExchangeRate is a free data retrieval call binding the contract method 0xb940a003.
//
// Solidity: function inspectionModeExchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleSession) InspectionModeExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.InspectionModeExchangeRate(&_StaderOracle.CallOpts)
}

// InspectionModeExchangeRate is a free data retrieval call binding the contract method 0xb940a003.
//
// Solidity: function inspectionModeExchangeRate() view returns(uint256 reportingBlockNumber, uint256 totalETHBalance, uint256 totalETHXSupply)
func (_StaderOracle *StaderOracleCallerSession) InspectionModeExchangeRate() (struct {
	ReportingBlockNumber *big.Int
	TotalETHBalance      *big.Int
	TotalETHXSupply      *big.Int
}, error) {
	return _StaderOracle.Contract.InspectionModeExchangeRate(&_StaderOracle.CallOpts)
}

// IsPORFeedBasedERData is a free data retrieval call binding the contract method 0x16515428.
//
// Solidity: function isPORFeedBasedERData() view returns(bool)
func (_StaderOracle *StaderOracleCaller) IsPORFeedBasedERData(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "isPORFeedBasedERData")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPORFeedBasedERData is a free data retrieval call binding the contract method 0x16515428.
//
// Solidity: function isPORFeedBasedERData() view returns(bool)
func (_StaderOracle *StaderOracleSession) IsPORFeedBasedERData() (bool, error) {
	return _StaderOracle.Contract.IsPORFeedBasedERData(&_StaderOracle.CallOpts)
}

// IsPORFeedBasedERData is a free data retrieval call binding the contract method 0x16515428.
//
// Solidity: function isPORFeedBasedERData() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) IsPORFeedBasedERData() (bool, error) {
	return _StaderOracle.Contract.IsPORFeedBasedERData(&_StaderOracle.CallOpts)
}

// IsTrustedNode is a free data retrieval call binding the contract method 0x2f739b1d.
//
// Solidity: function isTrustedNode(address ) view returns(bool)
func (_StaderOracle *StaderOracleCaller) IsTrustedNode(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "isTrustedNode", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedNode is a free data retrieval call binding the contract method 0x2f739b1d.
//
// Solidity: function isTrustedNode(address ) view returns(bool)
func (_StaderOracle *StaderOracleSession) IsTrustedNode(arg0 common.Address) (bool, error) {
	return _StaderOracle.Contract.IsTrustedNode(&_StaderOracle.CallOpts, arg0)
}

// IsTrustedNode is a free data retrieval call binding the contract method 0x2f739b1d.
//
// Solidity: function isTrustedNode(address ) view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) IsTrustedNode(arg0 common.Address) (bool, error) {
	return _StaderOracle.Contract.IsTrustedNode(&_StaderOracle.CallOpts, arg0)
}

// LastReportedMAPDIndex is a free data retrieval call binding the contract method 0xa3737869.
//
// Solidity: function lastReportedMAPDIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) LastReportedMAPDIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastReportedMAPDIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastReportedMAPDIndex is a free data retrieval call binding the contract method 0xa3737869.
//
// Solidity: function lastReportedMAPDIndex() view returns(uint256)
func (_StaderOracle *StaderOracleSession) LastReportedMAPDIndex() (*big.Int, error) {
	return _StaderOracle.Contract.LastReportedMAPDIndex(&_StaderOracle.CallOpts)
}

// LastReportedMAPDIndex is a free data retrieval call binding the contract method 0xa3737869.
//
// Solidity: function lastReportedMAPDIndex() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) LastReportedMAPDIndex() (*big.Int, error) {
	return _StaderOracle.Contract.LastReportedMAPDIndex(&_StaderOracle.CallOpts)
}

// LastReportedSDPriceData is a free data retrieval call binding the contract method 0xa8c3a3a8.
//
// Solidity: function lastReportedSDPriceData() view returns(uint256 reportingBlockNumber, uint256 sdPriceInETH)
func (_StaderOracle *StaderOracleCaller) LastReportedSDPriceData(opts *bind.CallOpts) (struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "lastReportedSDPriceData")

	outstruct := new(struct {
		ReportingBlockNumber *big.Int
		SdPriceInETH         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SdPriceInETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LastReportedSDPriceData is a free data retrieval call binding the contract method 0xa8c3a3a8.
//
// Solidity: function lastReportedSDPriceData() view returns(uint256 reportingBlockNumber, uint256 sdPriceInETH)
func (_StaderOracle *StaderOracleSession) LastReportedSDPriceData() (struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}, error) {
	return _StaderOracle.Contract.LastReportedSDPriceData(&_StaderOracle.CallOpts)
}

// LastReportedSDPriceData is a free data retrieval call binding the contract method 0xa8c3a3a8.
//
// Solidity: function lastReportedSDPriceData() view returns(uint256 reportingBlockNumber, uint256 sdPriceInETH)
func (_StaderOracle *StaderOracleCallerSession) LastReportedSDPriceData() (struct {
	ReportingBlockNumber *big.Int
	SdPriceInETH         *big.Int
}, error) {
	return _StaderOracle.Contract.LastReportedSDPriceData(&_StaderOracle.CallOpts)
}

// MissedAttestationPenalty is a free data retrieval call binding the contract method 0x9773ee60.
//
// Solidity: function missedAttestationPenalty(bytes32 ) view returns(uint16)
func (_StaderOracle *StaderOracleCaller) MissedAttestationPenalty(opts *bind.CallOpts, arg0 [32]byte) (uint16, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "missedAttestationPenalty", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MissedAttestationPenalty is a free data retrieval call binding the contract method 0x9773ee60.
//
// Solidity: function missedAttestationPenalty(bytes32 ) view returns(uint16)
func (_StaderOracle *StaderOracleSession) MissedAttestationPenalty(arg0 [32]byte) (uint16, error) {
	return _StaderOracle.Contract.MissedAttestationPenalty(&_StaderOracle.CallOpts, arg0)
}

// MissedAttestationPenalty is a free data retrieval call binding the contract method 0x9773ee60.
//
// Solidity: function missedAttestationPenalty(bytes32 ) view returns(uint16)
func (_StaderOracle *StaderOracleCallerSession) MissedAttestationPenalty(arg0 [32]byte) (uint16, error) {
	return _StaderOracle.Contract.MissedAttestationPenalty(&_StaderOracle.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderOracle *StaderOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderOracle *StaderOracleSession) Paused() (bool, error) {
	return _StaderOracle.Contract.Paused(&_StaderOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) Paused() (bool, error) {
	return _StaderOracle.Contract.Paused(&_StaderOracle.CallOpts)
}

// ReportingBlockNumberForWithdrawnValidators is a free data retrieval call binding the contract method 0x1445f7a0.
//
// Solidity: function reportingBlockNumberForWithdrawnValidators() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) ReportingBlockNumberForWithdrawnValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "reportingBlockNumberForWithdrawnValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReportingBlockNumberForWithdrawnValidators is a free data retrieval call binding the contract method 0x1445f7a0.
//
// Solidity: function reportingBlockNumberForWithdrawnValidators() view returns(uint256)
func (_StaderOracle *StaderOracleSession) ReportingBlockNumberForWithdrawnValidators() (*big.Int, error) {
	return _StaderOracle.Contract.ReportingBlockNumberForWithdrawnValidators(&_StaderOracle.CallOpts)
}

// ReportingBlockNumberForWithdrawnValidators is a free data retrieval call binding the contract method 0x1445f7a0.
//
// Solidity: function reportingBlockNumberForWithdrawnValidators() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) ReportingBlockNumberForWithdrawnValidators() (*big.Int, error) {
	return _StaderOracle.Contract.ReportingBlockNumberForWithdrawnValidators(&_StaderOracle.CallOpts)
}

// SafeMode is a free data retrieval call binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() view returns(bool)
func (_StaderOracle *StaderOracleCaller) SafeMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "safeMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SafeMode is a free data retrieval call binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() view returns(bool)
func (_StaderOracle *StaderOracleSession) SafeMode() (bool, error) {
	return _StaderOracle.Contract.SafeMode(&_StaderOracle.CallOpts)
}

// SafeMode is a free data retrieval call binding the contract method 0xabe3219c.
//
// Solidity: function safeMode() view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) SafeMode() (bool, error) {
	return _StaderOracle.Contract.SafeMode(&_StaderOracle.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StaderOracle *StaderOracleCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StaderOracle *StaderOracleSession) StaderConfig() (common.Address, error) {
	return _StaderOracle.Contract.StaderConfig(&_StaderOracle.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_StaderOracle *StaderOracleCallerSession) StaderConfig() (common.Address, error) {
	return _StaderOracle.Contract.StaderConfig(&_StaderOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderOracle *StaderOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderOracle *StaderOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderOracle.Contract.SupportsInterface(&_StaderOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StaderOracle *StaderOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StaderOracle.Contract.SupportsInterface(&_StaderOracle.CallOpts, interfaceId)
}

// TrustedNodesCount is a free data retrieval call binding the contract method 0xae815a04.
//
// Solidity: function trustedNodesCount() view returns(uint256)
func (_StaderOracle *StaderOracleCaller) TrustedNodesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "trustedNodesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrustedNodesCount is a free data retrieval call binding the contract method 0xae815a04.
//
// Solidity: function trustedNodesCount() view returns(uint256)
func (_StaderOracle *StaderOracleSession) TrustedNodesCount() (*big.Int, error) {
	return _StaderOracle.Contract.TrustedNodesCount(&_StaderOracle.CallOpts)
}

// TrustedNodesCount is a free data retrieval call binding the contract method 0xae815a04.
//
// Solidity: function trustedNodesCount() view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) TrustedNodesCount() (*big.Int, error) {
	return _StaderOracle.Contract.TrustedNodesCount(&_StaderOracle.CallOpts)
}

// UpdateFrequencyMap is a free data retrieval call binding the contract method 0x7150bc5b.
//
// Solidity: function updateFrequencyMap(bytes32 ) view returns(uint256)
func (_StaderOracle *StaderOracleCaller) UpdateFrequencyMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "updateFrequencyMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateFrequencyMap is a free data retrieval call binding the contract method 0x7150bc5b.
//
// Solidity: function updateFrequencyMap(bytes32 ) view returns(uint256)
func (_StaderOracle *StaderOracleSession) UpdateFrequencyMap(arg0 [32]byte) (*big.Int, error) {
	return _StaderOracle.Contract.UpdateFrequencyMap(&_StaderOracle.CallOpts, arg0)
}

// UpdateFrequencyMap is a free data retrieval call binding the contract method 0x7150bc5b.
//
// Solidity: function updateFrequencyMap(bytes32 ) view returns(uint256)
func (_StaderOracle *StaderOracleCallerSession) UpdateFrequencyMap(arg0 [32]byte) (*big.Int, error) {
	return _StaderOracle.Contract.UpdateFrequencyMap(&_StaderOracle.CallOpts, arg0)
}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 exitingValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 exitingValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleCaller) ValidatorStats(opts *bind.CallOpts) (struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	var out []interface{}
	err := _StaderOracle.contract.Call(opts, &out, "validatorStats")

	outstruct := new(struct {
		ReportingBlockNumber     *big.Int
		ExitingValidatorsBalance *big.Int
		ExitedValidatorsBalance  *big.Int
		SlashedValidatorsBalance *big.Int
		ExitingValidatorsCount   uint32
		ExitedValidatorsCount    uint32
		SlashedValidatorsCount   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReportingBlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExitingValidatorsBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExitedValidatorsBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SlashedValidatorsBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExitingValidatorsCount = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ExitedValidatorsCount = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.SlashedValidatorsCount = *abi.ConvertType(out[6], new(uint32)).(*uint32)

	return *outstruct, err

}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 exitingValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 exitingValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleSession) ValidatorStats() (struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	return _StaderOracle.Contract.ValidatorStats(&_StaderOracle.CallOpts)
}

// ValidatorStats is a free data retrieval call binding the contract method 0x3b5eb03a.
//
// Solidity: function validatorStats() view returns(uint256 reportingBlockNumber, uint128 exitingValidatorsBalance, uint128 exitedValidatorsBalance, uint128 slashedValidatorsBalance, uint32 exitingValidatorsCount, uint32 exitedValidatorsCount, uint32 slashedValidatorsCount)
func (_StaderOracle *StaderOracleCallerSession) ValidatorStats() (struct {
	ReportingBlockNumber     *big.Int
	ExitingValidatorsBalance *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ExitingValidatorsCount   uint32
	ExitedValidatorsCount    uint32
	SlashedValidatorsCount   uint32
}, error) {
	return _StaderOracle.Contract.ValidatorStats(&_StaderOracle.CallOpts)
}

// AddTrustedNode is a paid mutator transaction binding the contract method 0xd6275dd7.
//
// Solidity: function addTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactor) AddTrustedNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "addTrustedNode", _nodeAddress)
}

// AddTrustedNode is a paid mutator transaction binding the contract method 0xd6275dd7.
//
// Solidity: function addTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleSession) AddTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.AddTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// AddTrustedNode is a paid mutator transaction binding the contract method 0xd6275dd7.
//
// Solidity: function addTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactorSession) AddTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.AddTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// CloseERInspectionMode is a paid mutator transaction binding the contract method 0x101b6e34.
//
// Solidity: function closeERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactor) CloseERInspectionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "closeERInspectionMode")
}

// CloseERInspectionMode is a paid mutator transaction binding the contract method 0x101b6e34.
//
// Solidity: function closeERInspectionMode() returns()
func (_StaderOracle *StaderOracleSession) CloseERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.CloseERInspectionMode(&_StaderOracle.TransactOpts)
}

// CloseERInspectionMode is a paid mutator transaction binding the contract method 0x101b6e34.
//
// Solidity: function closeERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) CloseERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.CloseERInspectionMode(&_StaderOracle.TransactOpts)
}

// DisableERInspectionMode is a paid mutator transaction binding the contract method 0xe10025e6.
//
// Solidity: function disableERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactor) DisableERInspectionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "disableERInspectionMode")
}

// DisableERInspectionMode is a paid mutator transaction binding the contract method 0xe10025e6.
//
// Solidity: function disableERInspectionMode() returns()
func (_StaderOracle *StaderOracleSession) DisableERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableERInspectionMode(&_StaderOracle.TransactOpts)
}

// DisableERInspectionMode is a paid mutator transaction binding the contract method 0xe10025e6.
//
// Solidity: function disableERInspectionMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) DisableERInspectionMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableERInspectionMode(&_StaderOracle.TransactOpts)
}

// DisableSafeMode is a paid mutator transaction binding the contract method 0xe514fe55.
//
// Solidity: function disableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactor) DisableSafeMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "disableSafeMode")
}

// DisableSafeMode is a paid mutator transaction binding the contract method 0xe514fe55.
//
// Solidity: function disableSafeMode() returns()
func (_StaderOracle *StaderOracleSession) DisableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableSafeMode(&_StaderOracle.TransactOpts)
}

// DisableSafeMode is a paid mutator transaction binding the contract method 0xe514fe55.
//
// Solidity: function disableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) DisableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.DisableSafeMode(&_StaderOracle.TransactOpts)
}

// EnableSafeMode is a paid mutator transaction binding the contract method 0x4f560abd.
//
// Solidity: function enableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactor) EnableSafeMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "enableSafeMode")
}

// EnableSafeMode is a paid mutator transaction binding the contract method 0x4f560abd.
//
// Solidity: function enableSafeMode() returns()
func (_StaderOracle *StaderOracleSession) EnableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.EnableSafeMode(&_StaderOracle.TransactOpts)
}

// EnableSafeMode is a paid mutator transaction binding the contract method 0x4f560abd.
//
// Solidity: function enableSafeMode() returns()
func (_StaderOracle *StaderOracleTransactorSession) EnableSafeMode() (*types.Transaction, error) {
	return _StaderOracle.Contract.EnableSafeMode(&_StaderOracle.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.GrantRole(&_StaderOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.GrantRole(&_StaderOracle.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_StaderOracle *StaderOracleSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.Initialize(&_StaderOracle.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.Initialize(&_StaderOracle.TransactOpts, _admin, _staderConfig)
}

// RemoveTrustedNode is a paid mutator transaction binding the contract method 0x52e0fc80.
//
// Solidity: function removeTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactor) RemoveTrustedNode(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "removeTrustedNode", _nodeAddress)
}

// RemoveTrustedNode is a paid mutator transaction binding the contract method 0x52e0fc80.
//
// Solidity: function removeTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleSession) RemoveTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RemoveTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// RemoveTrustedNode is a paid mutator transaction binding the contract method 0x52e0fc80.
//
// Solidity: function removeTrustedNode(address _nodeAddress) returns()
func (_StaderOracle *StaderOracleTransactorSession) RemoveTrustedNode(_nodeAddress common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RemoveTrustedNode(&_StaderOracle.TransactOpts, _nodeAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RenounceRole(&_StaderOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RenounceRole(&_StaderOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RevokeRole(&_StaderOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StaderOracle *StaderOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.RevokeRole(&_StaderOracle.TransactOpts, role, account)
}

// SetERUpdateFrequency is a paid mutator transaction binding the contract method 0xd0a8f679.
//
// Solidity: function setERUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetERUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setERUpdateFrequency", _updateFrequency)
}

// SetERUpdateFrequency is a paid mutator transaction binding the contract method 0xd0a8f679.
//
// Solidity: function setERUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetERUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetERUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetERUpdateFrequency is a paid mutator transaction binding the contract method 0xd0a8f679.
//
// Solidity: function setERUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetERUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetERUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetMissedAttestationPenaltyUpdateFrequency is a paid mutator transaction binding the contract method 0xf51c0fe7.
//
// Solidity: function setMissedAttestationPenaltyUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetMissedAttestationPenaltyUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setMissedAttestationPenaltyUpdateFrequency", _updateFrequency)
}

// SetMissedAttestationPenaltyUpdateFrequency is a paid mutator transaction binding the contract method 0xf51c0fe7.
//
// Solidity: function setMissedAttestationPenaltyUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetMissedAttestationPenaltyUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetMissedAttestationPenaltyUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetMissedAttestationPenaltyUpdateFrequency is a paid mutator transaction binding the contract method 0xf51c0fe7.
//
// Solidity: function setMissedAttestationPenaltyUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetMissedAttestationPenaltyUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetMissedAttestationPenaltyUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetSDPriceUpdateFrequency is a paid mutator transaction binding the contract method 0x749f7d8a.
//
// Solidity: function setSDPriceUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetSDPriceUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setSDPriceUpdateFrequency", _updateFrequency)
}

// SetSDPriceUpdateFrequency is a paid mutator transaction binding the contract method 0x749f7d8a.
//
// Solidity: function setSDPriceUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetSDPriceUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetSDPriceUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetSDPriceUpdateFrequency is a paid mutator transaction binding the contract method 0x749f7d8a.
//
// Solidity: function setSDPriceUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetSDPriceUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetSDPriceUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetValidatorStatsUpdateFrequency is a paid mutator transaction binding the contract method 0x844007fe.
//
// Solidity: function setValidatorStatsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetValidatorStatsUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setValidatorStatsUpdateFrequency", _updateFrequency)
}

// SetValidatorStatsUpdateFrequency is a paid mutator transaction binding the contract method 0x844007fe.
//
// Solidity: function setValidatorStatsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetValidatorStatsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetValidatorStatsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetValidatorStatsUpdateFrequency is a paid mutator transaction binding the contract method 0x844007fe.
//
// Solidity: function setValidatorStatsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetValidatorStatsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetValidatorStatsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetWithdrawnValidatorsUpdateFrequency is a paid mutator transaction binding the contract method 0xc06a6201.
//
// Solidity: function setWithdrawnValidatorsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactor) SetWithdrawnValidatorsUpdateFrequency(opts *bind.TransactOpts, _updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "setWithdrawnValidatorsUpdateFrequency", _updateFrequency)
}

// SetWithdrawnValidatorsUpdateFrequency is a paid mutator transaction binding the contract method 0xc06a6201.
//
// Solidity: function setWithdrawnValidatorsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleSession) SetWithdrawnValidatorsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetWithdrawnValidatorsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SetWithdrawnValidatorsUpdateFrequency is a paid mutator transaction binding the contract method 0xc06a6201.
//
// Solidity: function setWithdrawnValidatorsUpdateFrequency(uint256 _updateFrequency) returns()
func (_StaderOracle *StaderOracleTransactorSession) SetWithdrawnValidatorsUpdateFrequency(_updateFrequency *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.SetWithdrawnValidatorsUpdateFrequency(&_StaderOracle.TransactOpts, _updateFrequency)
}

// SubmitExchangeRateData is a paid mutator transaction binding the contract method 0x818c8b26.
//
// Solidity: function submitExchangeRateData((uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitExchangeRateData(opts *bind.TransactOpts, _exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitExchangeRateData", _exchangeRate)
}

// SubmitExchangeRateData is a paid mutator transaction binding the contract method 0x818c8b26.
//
// Solidity: function submitExchangeRateData((uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleSession) SubmitExchangeRateData(_exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitExchangeRateData(&_StaderOracle.TransactOpts, _exchangeRate)
}

// SubmitExchangeRateData is a paid mutator transaction binding the contract method 0x818c8b26.
//
// Solidity: function submitExchangeRateData((uint256,uint256,uint256) _exchangeRate) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitExchangeRateData(_exchangeRate ExchangeRate) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitExchangeRateData(&_StaderOracle.TransactOpts, _exchangeRate)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x67fbf731.
//
// Solidity: function submitMissedAttestationPenalties((uint256,uint256,bytes[]) _mapd) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitMissedAttestationPenalties(opts *bind.TransactOpts, _mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitMissedAttestationPenalties", _mapd)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x67fbf731.
//
// Solidity: function submitMissedAttestationPenalties((uint256,uint256,bytes[]) _mapd) returns()
func (_StaderOracle *StaderOracleSession) SubmitMissedAttestationPenalties(_mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitMissedAttestationPenalties(&_StaderOracle.TransactOpts, _mapd)
}

// SubmitMissedAttestationPenalties is a paid mutator transaction binding the contract method 0x67fbf731.
//
// Solidity: function submitMissedAttestationPenalties((uint256,uint256,bytes[]) _mapd) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitMissedAttestationPenalties(_mapd MissedAttestationPenaltyData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitMissedAttestationPenalties(&_StaderOracle.TransactOpts, _mapd)
}

// SubmitSDPrice is a paid mutator transaction binding the contract method 0xf6a3c090.
//
// Solidity: function submitSDPrice((uint256,uint256) _sdPriceData) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitSDPrice(opts *bind.TransactOpts, _sdPriceData SDPriceData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitSDPrice", _sdPriceData)
}

// SubmitSDPrice is a paid mutator transaction binding the contract method 0xf6a3c090.
//
// Solidity: function submitSDPrice((uint256,uint256) _sdPriceData) returns()
func (_StaderOracle *StaderOracleSession) SubmitSDPrice(_sdPriceData SDPriceData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSDPrice(&_StaderOracle.TransactOpts, _sdPriceData)
}

// SubmitSDPrice is a paid mutator transaction binding the contract method 0xf6a3c090.
//
// Solidity: function submitSDPrice((uint256,uint256) _sdPriceData) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitSDPrice(_sdPriceData SDPriceData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSDPrice(&_StaderOracle.TransactOpts, _sdPriceData)
}

// SubmitSocializingRewardsMerkleRoot is a paid mutator transaction binding the contract method 0xae541d65.
//
// Solidity: function submitSocializingRewardsMerkleRoot((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitSocializingRewardsMerkleRoot(opts *bind.TransactOpts, _rewardsData RewardsData) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitSocializingRewardsMerkleRoot", _rewardsData)
}

// SubmitSocializingRewardsMerkleRoot is a paid mutator transaction binding the contract method 0xae541d65.
//
// Solidity: function submitSocializingRewardsMerkleRoot((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_StaderOracle *StaderOracleSession) SubmitSocializingRewardsMerkleRoot(_rewardsData RewardsData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSocializingRewardsMerkleRoot(&_StaderOracle.TransactOpts, _rewardsData)
}

// SubmitSocializingRewardsMerkleRoot is a paid mutator transaction binding the contract method 0xae541d65.
//
// Solidity: function submitSocializingRewardsMerkleRoot((uint256,uint256,bytes32,uint8,uint256,uint256,uint256,uint256) _rewardsData) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitSocializingRewardsMerkleRoot(_rewardsData RewardsData) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitSocializingRewardsMerkleRoot(&_StaderOracle.TransactOpts, _rewardsData)
}

// SubmitValidatorStats is a paid mutator transaction binding the contract method 0x5c7ccd3b.
//
// Solidity: function submitValidatorStats((uint256,uint128,uint128,uint128,uint32,uint32,uint32) _validatorStats) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitValidatorStats(opts *bind.TransactOpts, _validatorStats ValidatorStats) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitValidatorStats", _validatorStats)
}

// SubmitValidatorStats is a paid mutator transaction binding the contract method 0x5c7ccd3b.
//
// Solidity: function submitValidatorStats((uint256,uint128,uint128,uint128,uint32,uint32,uint32) _validatorStats) returns()
func (_StaderOracle *StaderOracleSession) SubmitValidatorStats(_validatorStats ValidatorStats) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitValidatorStats(&_StaderOracle.TransactOpts, _validatorStats)
}

// SubmitValidatorStats is a paid mutator transaction binding the contract method 0x5c7ccd3b.
//
// Solidity: function submitValidatorStats((uint256,uint128,uint128,uint128,uint32,uint32,uint32) _validatorStats) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitValidatorStats(_validatorStats ValidatorStats) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitValidatorStats(&_StaderOracle.TransactOpts, _validatorStats)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xf0626d72.
//
// Solidity: function submitWithdrawnValidators((uint256,address,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleTransactor) SubmitWithdrawnValidators(opts *bind.TransactOpts, _withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "submitWithdrawnValidators", _withdrawnValidators)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xf0626d72.
//
// Solidity: function submitWithdrawnValidators((uint256,address,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleSession) SubmitWithdrawnValidators(_withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitWithdrawnValidators(&_StaderOracle.TransactOpts, _withdrawnValidators)
}

// SubmitWithdrawnValidators is a paid mutator transaction binding the contract method 0xf0626d72.
//
// Solidity: function submitWithdrawnValidators((uint256,address,bytes[]) _withdrawnValidators) returns()
func (_StaderOracle *StaderOracleTransactorSession) SubmitWithdrawnValidators(_withdrawnValidators WithdrawnValidators) (*types.Transaction, error) {
	return _StaderOracle.Contract.SubmitWithdrawnValidators(&_StaderOracle.TransactOpts, _withdrawnValidators)
}

// TogglePORFeedBasedERData is a paid mutator transaction binding the contract method 0x712033eb.
//
// Solidity: function togglePORFeedBasedERData() returns()
func (_StaderOracle *StaderOracleTransactor) TogglePORFeedBasedERData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "togglePORFeedBasedERData")
}

// TogglePORFeedBasedERData is a paid mutator transaction binding the contract method 0x712033eb.
//
// Solidity: function togglePORFeedBasedERData() returns()
func (_StaderOracle *StaderOracleSession) TogglePORFeedBasedERData() (*types.Transaction, error) {
	return _StaderOracle.Contract.TogglePORFeedBasedERData(&_StaderOracle.TransactOpts)
}

// TogglePORFeedBasedERData is a paid mutator transaction binding the contract method 0x712033eb.
//
// Solidity: function togglePORFeedBasedERData() returns()
func (_StaderOracle *StaderOracleTransactorSession) TogglePORFeedBasedERData() (*types.Transaction, error) {
	return _StaderOracle.Contract.TogglePORFeedBasedERData(&_StaderOracle.TransactOpts)
}

// UpdateERChangeLimit is a paid mutator transaction binding the contract method 0x8ca8703c.
//
// Solidity: function updateERChangeLimit(uint256 _erChangeLimit) returns()
func (_StaderOracle *StaderOracleTransactor) UpdateERChangeLimit(opts *bind.TransactOpts, _erChangeLimit *big.Int) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateERChangeLimit", _erChangeLimit)
}

// UpdateERChangeLimit is a paid mutator transaction binding the contract method 0x8ca8703c.
//
// Solidity: function updateERChangeLimit(uint256 _erChangeLimit) returns()
func (_StaderOracle *StaderOracleSession) UpdateERChangeLimit(_erChangeLimit *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERChangeLimit(&_StaderOracle.TransactOpts, _erChangeLimit)
}

// UpdateERChangeLimit is a paid mutator transaction binding the contract method 0x8ca8703c.
//
// Solidity: function updateERChangeLimit(uint256 _erChangeLimit) returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateERChangeLimit(_erChangeLimit *big.Int) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERChangeLimit(&_StaderOracle.TransactOpts, _erChangeLimit)
}

// UpdateERFromPORFeed is a paid mutator transaction binding the contract method 0x9bfdf9a4.
//
// Solidity: function updateERFromPORFeed() returns()
func (_StaderOracle *StaderOracleTransactor) UpdateERFromPORFeed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateERFromPORFeed")
}

// UpdateERFromPORFeed is a paid mutator transaction binding the contract method 0x9bfdf9a4.
//
// Solidity: function updateERFromPORFeed() returns()
func (_StaderOracle *StaderOracleSession) UpdateERFromPORFeed() (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERFromPORFeed(&_StaderOracle.TransactOpts)
}

// UpdateERFromPORFeed is a paid mutator transaction binding the contract method 0x9bfdf9a4.
//
// Solidity: function updateERFromPORFeed() returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateERFromPORFeed() (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateERFromPORFeed(&_StaderOracle.TransactOpts)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StaderOracle *StaderOracleSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateStaderConfig(&_StaderOracle.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_StaderOracle *StaderOracleTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _StaderOracle.Contract.UpdateStaderConfig(&_StaderOracle.TransactOpts, _staderConfig)
}

// StaderOracleERDataSourceToggledIterator is returned from FilterERDataSourceToggled and is used to iterate over the raw logs and unpacked data for ERDataSourceToggled events raised by the StaderOracle contract.
type StaderOracleERDataSourceToggledIterator struct {
	Event *StaderOracleERDataSourceToggled // Event containing the contract specifics and raw log

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
func (it *StaderOracleERDataSourceToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleERDataSourceToggled)
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
		it.Event = new(StaderOracleERDataSourceToggled)
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
func (it *StaderOracleERDataSourceToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleERDataSourceToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleERDataSourceToggled represents a ERDataSourceToggled event raised by the StaderOracle contract.
type StaderOracleERDataSourceToggled struct {
	IsPORBasedERData bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERDataSourceToggled is a free log retrieval operation binding the contract event 0xc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5.
//
// Solidity: event ERDataSourceToggled(bool isPORBasedERData)
func (_StaderOracle *StaderOracleFilterer) FilterERDataSourceToggled(opts *bind.FilterOpts) (*StaderOracleERDataSourceToggledIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ERDataSourceToggled")
	if err != nil {
		return nil, err
	}
	return &StaderOracleERDataSourceToggledIterator{contract: _StaderOracle.contract, event: "ERDataSourceToggled", logs: logs, sub: sub}, nil
}

// WatchERDataSourceToggled is a free log subscription operation binding the contract event 0xc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5.
//
// Solidity: event ERDataSourceToggled(bool isPORBasedERData)
func (_StaderOracle *StaderOracleFilterer) WatchERDataSourceToggled(opts *bind.WatchOpts, sink chan<- *StaderOracleERDataSourceToggled) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ERDataSourceToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleERDataSourceToggled)
				if err := _StaderOracle.contract.UnpackLog(event, "ERDataSourceToggled", log); err != nil {
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

// ParseERDataSourceToggled is a log parse operation binding the contract event 0xc59a5de02f9d69be770ff0d61bbc894968433bb599f9fd9c2016e149c509c5e5.
//
// Solidity: event ERDataSourceToggled(bool isPORBasedERData)
func (_StaderOracle *StaderOracleFilterer) ParseERDataSourceToggled(log types.Log) (*StaderOracleERDataSourceToggled, error) {
	event := new(StaderOracleERDataSourceToggled)
	if err := _StaderOracle.contract.UnpackLog(event, "ERDataSourceToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleERInspectionModeActivatedIterator is returned from FilterERInspectionModeActivated and is used to iterate over the raw logs and unpacked data for ERInspectionModeActivated events raised by the StaderOracle contract.
type StaderOracleERInspectionModeActivatedIterator struct {
	Event *StaderOracleERInspectionModeActivated // Event containing the contract specifics and raw log

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
func (it *StaderOracleERInspectionModeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleERInspectionModeActivated)
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
		it.Event = new(StaderOracleERInspectionModeActivated)
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
func (it *StaderOracleERInspectionModeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleERInspectionModeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleERInspectionModeActivated represents a ERInspectionModeActivated event raised by the StaderOracle contract.
type StaderOracleERInspectionModeActivated struct {
	ErInspectionMode bool
	Time             *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERInspectionModeActivated is a free log retrieval operation binding the contract event 0x9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415.
//
// Solidity: event ERInspectionModeActivated(bool erInspectionMode, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterERInspectionModeActivated(opts *bind.FilterOpts) (*StaderOracleERInspectionModeActivatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ERInspectionModeActivated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleERInspectionModeActivatedIterator{contract: _StaderOracle.contract, event: "ERInspectionModeActivated", logs: logs, sub: sub}, nil
}

// WatchERInspectionModeActivated is a free log subscription operation binding the contract event 0x9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415.
//
// Solidity: event ERInspectionModeActivated(bool erInspectionMode, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchERInspectionModeActivated(opts *bind.WatchOpts, sink chan<- *StaderOracleERInspectionModeActivated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ERInspectionModeActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleERInspectionModeActivated)
				if err := _StaderOracle.contract.UnpackLog(event, "ERInspectionModeActivated", log); err != nil {
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

// ParseERInspectionModeActivated is a log parse operation binding the contract event 0x9dea6ddefbcfcf9c4f6c1c086e462c2ab380c0be199d0289bf23b09f20814415.
//
// Solidity: event ERInspectionModeActivated(bool erInspectionMode, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseERInspectionModeActivated(log types.Log) (*StaderOracleERInspectionModeActivated, error) {
	event := new(StaderOracleERInspectionModeActivated)
	if err := _StaderOracle.contract.UnpackLog(event, "ERInspectionModeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleExchangeRateSubmittedIterator is returned from FilterExchangeRateSubmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateSubmitted events raised by the StaderOracle contract.
type StaderOracleExchangeRateSubmittedIterator struct {
	Event *StaderOracleExchangeRateSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleExchangeRateSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleExchangeRateSubmitted)
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
		it.Event = new(StaderOracleExchangeRateSubmitted)
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
func (it *StaderOracleExchangeRateSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleExchangeRateSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleExchangeRateSubmitted represents a ExchangeRateSubmitted event raised by the StaderOracle contract.
type StaderOracleExchangeRateSubmitted struct {
	From       common.Address
	Block      *big.Int
	TotalEth   *big.Int
	EthxSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateSubmitted is a free log retrieval operation binding the contract event 0x73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb12.
//
// Solidity: event ExchangeRateSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterExchangeRateSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleExchangeRateSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ExchangeRateSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleExchangeRateSubmittedIterator{contract: _StaderOracle.contract, event: "ExchangeRateSubmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateSubmitted is a free log subscription operation binding the contract event 0x73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb12.
//
// Solidity: event ExchangeRateSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchExchangeRateSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleExchangeRateSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ExchangeRateSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleExchangeRateSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateSubmitted", log); err != nil {
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

// ParseExchangeRateSubmitted is a log parse operation binding the contract event 0x73327a5c0fdb3104b4a0f993bc20ce59885ac5bfe5f23e4bfdd19c21fb79cb12.
//
// Solidity: event ExchangeRateSubmitted(address indexed from, uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseExchangeRateSubmitted(log types.Log) (*StaderOracleExchangeRateSubmitted, error) {
	event := new(StaderOracleExchangeRateSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleExchangeRateUpdatedIterator is returned from FilterExchangeRateUpdated and is used to iterate over the raw logs and unpacked data for ExchangeRateUpdated events raised by the StaderOracle contract.
type StaderOracleExchangeRateUpdatedIterator struct {
	Event *StaderOracleExchangeRateUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleExchangeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleExchangeRateUpdated)
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
		it.Event = new(StaderOracleExchangeRateUpdated)
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
func (it *StaderOracleExchangeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleExchangeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleExchangeRateUpdated represents a ExchangeRateUpdated event raised by the StaderOracle contract.
type StaderOracleExchangeRateUpdated struct {
	Block      *big.Int
	TotalEth   *big.Int
	EthxSupply *big.Int
	Time       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateUpdated is a free log retrieval operation binding the contract event 0xf525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e3.
//
// Solidity: event ExchangeRateUpdated(uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterExchangeRateUpdated(opts *bind.FilterOpts) (*StaderOracleExchangeRateUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleExchangeRateUpdatedIterator{contract: _StaderOracle.contract, event: "ExchangeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchExchangeRateUpdated is a free log subscription operation binding the contract event 0xf525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e3.
//
// Solidity: event ExchangeRateUpdated(uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchExchangeRateUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleExchangeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleExchangeRateUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
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

// ParseExchangeRateUpdated is a log parse operation binding the contract event 0xf525f19964f35afcb9a475680bb27abecbc5e62b4d6d4f66a81a5bd7e8a107e3.
//
// Solidity: event ExchangeRateUpdated(uint256 block, uint256 totalEth, uint256 ethxSupply, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseExchangeRateUpdated(log types.Log) (*StaderOracleExchangeRateUpdated, error) {
	event := new(StaderOracleExchangeRateUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StaderOracle contract.
type StaderOracleInitializedIterator struct {
	Event *StaderOracleInitialized // Event containing the contract specifics and raw log

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
func (it *StaderOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleInitialized)
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
		it.Event = new(StaderOracleInitialized)
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
func (it *StaderOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleInitialized represents a Initialized event raised by the StaderOracle contract.
type StaderOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderOracle *StaderOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*StaderOracleInitializedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StaderOracleInitializedIterator{contract: _StaderOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderOracle *StaderOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StaderOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleInitialized)
				if err := _StaderOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StaderOracle *StaderOracleFilterer) ParseInitialized(log types.Log) (*StaderOracleInitialized, error) {
	event := new(StaderOracleInitialized)
	if err := _StaderOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleMissedAttestationPenaltySubmittedIterator is returned from FilterMissedAttestationPenaltySubmitted and is used to iterate over the raw logs and unpacked data for MissedAttestationPenaltySubmitted events raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltySubmittedIterator struct {
	Event *StaderOracleMissedAttestationPenaltySubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleMissedAttestationPenaltySubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleMissedAttestationPenaltySubmitted)
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
		it.Event = new(StaderOracleMissedAttestationPenaltySubmitted)
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
func (it *StaderOracleMissedAttestationPenaltySubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleMissedAttestationPenaltySubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleMissedAttestationPenaltySubmitted represents a MissedAttestationPenaltySubmitted event raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltySubmitted struct {
	Node                 common.Address
	Index                *big.Int
	Block                *big.Int
	ReportingBlockNumber *big.Int
	Pubkeys              [][]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMissedAttestationPenaltySubmitted is a free log retrieval operation binding the contract event 0x51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 block, uint256 reportingBlockNumber, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) FilterMissedAttestationPenaltySubmitted(opts *bind.FilterOpts, node []common.Address) (*StaderOracleMissedAttestationPenaltySubmittedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "MissedAttestationPenaltySubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleMissedAttestationPenaltySubmittedIterator{contract: _StaderOracle.contract, event: "MissedAttestationPenaltySubmitted", logs: logs, sub: sub}, nil
}

// WatchMissedAttestationPenaltySubmitted is a free log subscription operation binding the contract event 0x51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 block, uint256 reportingBlockNumber, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) WatchMissedAttestationPenaltySubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleMissedAttestationPenaltySubmitted, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "MissedAttestationPenaltySubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleMissedAttestationPenaltySubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltySubmitted", log); err != nil {
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

// ParseMissedAttestationPenaltySubmitted is a log parse operation binding the contract event 0x51308cad1da8fe95d4be43112c17d5651d3e3713a675ec61c2214fa16d9f6beb.
//
// Solidity: event MissedAttestationPenaltySubmitted(address indexed node, uint256 index, uint256 block, uint256 reportingBlockNumber, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) ParseMissedAttestationPenaltySubmitted(log types.Log) (*StaderOracleMissedAttestationPenaltySubmitted, error) {
	event := new(StaderOracleMissedAttestationPenaltySubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltySubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleMissedAttestationPenaltyUpdatedIterator is returned from FilterMissedAttestationPenaltyUpdated and is used to iterate over the raw logs and unpacked data for MissedAttestationPenaltyUpdated events raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltyUpdatedIterator struct {
	Event *StaderOracleMissedAttestationPenaltyUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleMissedAttestationPenaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleMissedAttestationPenaltyUpdated)
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
		it.Event = new(StaderOracleMissedAttestationPenaltyUpdated)
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
func (it *StaderOracleMissedAttestationPenaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleMissedAttestationPenaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleMissedAttestationPenaltyUpdated represents a MissedAttestationPenaltyUpdated event raised by the StaderOracle contract.
type StaderOracleMissedAttestationPenaltyUpdated struct {
	Index   *big.Int
	Block   *big.Int
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMissedAttestationPenaltyUpdated is a free log retrieval operation binding the contract event 0x5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a191.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) FilterMissedAttestationPenaltyUpdated(opts *bind.FilterOpts) (*StaderOracleMissedAttestationPenaltyUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "MissedAttestationPenaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleMissedAttestationPenaltyUpdatedIterator{contract: _StaderOracle.contract, event: "MissedAttestationPenaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchMissedAttestationPenaltyUpdated is a free log subscription operation binding the contract event 0x5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a191.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) WatchMissedAttestationPenaltyUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleMissedAttestationPenaltyUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "MissedAttestationPenaltyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleMissedAttestationPenaltyUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltyUpdated", log); err != nil {
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

// ParseMissedAttestationPenaltyUpdated is a log parse operation binding the contract event 0x5454855cf2eeb89296b9e10ba0881425fa305f06ce9ccb1b0ce47bc2f107a191.
//
// Solidity: event MissedAttestationPenaltyUpdated(uint256 index, uint256 block, bytes[] pubkeys)
func (_StaderOracle *StaderOracleFilterer) ParseMissedAttestationPenaltyUpdated(log types.Log) (*StaderOracleMissedAttestationPenaltyUpdated, error) {
	event := new(StaderOracleMissedAttestationPenaltyUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "MissedAttestationPenaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StaderOracle contract.
type StaderOraclePausedIterator struct {
	Event *StaderOraclePaused // Event containing the contract specifics and raw log

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
func (it *StaderOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOraclePaused)
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
		it.Event = new(StaderOraclePaused)
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
func (it *StaderOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOraclePaused represents a Paused event raised by the StaderOracle contract.
type StaderOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StaderOracle *StaderOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*StaderOraclePausedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StaderOraclePausedIterator{contract: _StaderOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StaderOracle *StaderOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StaderOraclePaused) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOraclePaused)
				if err := _StaderOracle.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StaderOracle *StaderOracleFilterer) ParsePaused(log types.Log) (*StaderOraclePaused, error) {
	event := new(StaderOraclePaused)
	if err := _StaderOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StaderOracle contract.
type StaderOracleRoleAdminChangedIterator struct {
	Event *StaderOracleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StaderOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleRoleAdminChanged)
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
		it.Event = new(StaderOracleRoleAdminChanged)
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
func (it *StaderOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleRoleAdminChanged represents a RoleAdminChanged event raised by the StaderOracle contract.
type StaderOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderOracle *StaderOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StaderOracleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleRoleAdminChangedIterator{contract: _StaderOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StaderOracle *StaderOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StaderOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleRoleAdminChanged)
				if err := _StaderOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseRoleAdminChanged(log types.Log) (*StaderOracleRoleAdminChanged, error) {
	event := new(StaderOracleRoleAdminChanged)
	if err := _StaderOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StaderOracle contract.
type StaderOracleRoleGrantedIterator struct {
	Event *StaderOracleRoleGranted // Event containing the contract specifics and raw log

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
func (it *StaderOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleRoleGranted)
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
		it.Event = new(StaderOracleRoleGranted)
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
func (it *StaderOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleRoleGranted represents a RoleGranted event raised by the StaderOracle contract.
type StaderOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StaderOracleRoleGrantedIterator, error) {

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

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleRoleGrantedIterator{contract: _StaderOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StaderOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleRoleGranted)
				if err := _StaderOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseRoleGranted(log types.Log) (*StaderOracleRoleGranted, error) {
	event := new(StaderOracleRoleGranted)
	if err := _StaderOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StaderOracle contract.
type StaderOracleRoleRevokedIterator struct {
	Event *StaderOracleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StaderOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleRoleRevoked)
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
		it.Event = new(StaderOracleRoleRevoked)
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
func (it *StaderOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleRoleRevoked represents a RoleRevoked event raised by the StaderOracle contract.
type StaderOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StaderOracleRoleRevokedIterator, error) {

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

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleRoleRevokedIterator{contract: _StaderOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StaderOracle *StaderOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StaderOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleRoleRevoked)
				if err := _StaderOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StaderOracle *StaderOracleFilterer) ParseRoleRevoked(log types.Log) (*StaderOracleRoleRevoked, error) {
	event := new(StaderOracleRoleRevoked)
	if err := _StaderOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSDPriceSubmittedIterator is returned from FilterSDPriceSubmitted and is used to iterate over the raw logs and unpacked data for SDPriceSubmitted events raised by the StaderOracle contract.
type StaderOracleSDPriceSubmittedIterator struct {
	Event *StaderOracleSDPriceSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleSDPriceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSDPriceSubmitted)
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
		it.Event = new(StaderOracleSDPriceSubmitted)
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
func (it *StaderOracleSDPriceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSDPriceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSDPriceSubmitted represents a SDPriceSubmitted event raised by the StaderOracle contract.
type StaderOracleSDPriceSubmitted struct {
	Node          common.Address
	SdPriceInETH  *big.Int
	ReportedBlock *big.Int
	Block         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSDPriceSubmitted is a free log retrieval operation binding the contract event 0x6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8.
//
// Solidity: event SDPriceSubmitted(address indexed node, uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSDPriceSubmitted(opts *bind.FilterOpts, node []common.Address) (*StaderOracleSDPriceSubmittedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SDPriceSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleSDPriceSubmittedIterator{contract: _StaderOracle.contract, event: "SDPriceSubmitted", logs: logs, sub: sub}, nil
}

// WatchSDPriceSubmitted is a free log subscription operation binding the contract event 0x6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8.
//
// Solidity: event SDPriceSubmitted(address indexed node, uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSDPriceSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleSDPriceSubmitted, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SDPriceSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSDPriceSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "SDPriceSubmitted", log); err != nil {
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

// ParseSDPriceSubmitted is a log parse operation binding the contract event 0x6c291a7ce9906b2982643002c104cb0ba9f2b9fecc8b38e9cc3cf5cfca65b4e8.
//
// Solidity: event SDPriceSubmitted(address indexed node, uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSDPriceSubmitted(log types.Log) (*StaderOracleSDPriceSubmitted, error) {
	event := new(StaderOracleSDPriceSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "SDPriceSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSDPriceUpdatedIterator is returned from FilterSDPriceUpdated and is used to iterate over the raw logs and unpacked data for SDPriceUpdated events raised by the StaderOracle contract.
type StaderOracleSDPriceUpdatedIterator struct {
	Event *StaderOracleSDPriceUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleSDPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSDPriceUpdated)
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
		it.Event = new(StaderOracleSDPriceUpdated)
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
func (it *StaderOracleSDPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSDPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSDPriceUpdated represents a SDPriceUpdated event raised by the StaderOracle contract.
type StaderOracleSDPriceUpdated struct {
	SdPriceInETH  *big.Int
	ReportedBlock *big.Int
	Block         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSDPriceUpdated is a free log retrieval operation binding the contract event 0xbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc.
//
// Solidity: event SDPriceUpdated(uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSDPriceUpdated(opts *bind.FilterOpts) (*StaderOracleSDPriceUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SDPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSDPriceUpdatedIterator{contract: _StaderOracle.contract, event: "SDPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchSDPriceUpdated is a free log subscription operation binding the contract event 0xbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc.
//
// Solidity: event SDPriceUpdated(uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSDPriceUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleSDPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SDPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSDPriceUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "SDPriceUpdated", log); err != nil {
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

// ParseSDPriceUpdated is a log parse operation binding the contract event 0xbc1a0795e699bbeaa982f6049ae9689f4d0e3e22d554adb7c46626be62f3b8bc.
//
// Solidity: event SDPriceUpdated(uint256 sdPriceInETH, uint256 reportedBlock, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSDPriceUpdated(log types.Log) (*StaderOracleSDPriceUpdated, error) {
	event := new(StaderOracleSDPriceUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "SDPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSafeModeDisabledIterator is returned from FilterSafeModeDisabled and is used to iterate over the raw logs and unpacked data for SafeModeDisabled events raised by the StaderOracle contract.
type StaderOracleSafeModeDisabledIterator struct {
	Event *StaderOracleSafeModeDisabled // Event containing the contract specifics and raw log

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
func (it *StaderOracleSafeModeDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSafeModeDisabled)
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
		it.Event = new(StaderOracleSafeModeDisabled)
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
func (it *StaderOracleSafeModeDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSafeModeDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSafeModeDisabled represents a SafeModeDisabled event raised by the StaderOracle contract.
type StaderOracleSafeModeDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSafeModeDisabled is a free log retrieval operation binding the contract event 0xf29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292.
//
// Solidity: event SafeModeDisabled()
func (_StaderOracle *StaderOracleFilterer) FilterSafeModeDisabled(opts *bind.FilterOpts) (*StaderOracleSafeModeDisabledIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SafeModeDisabled")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSafeModeDisabledIterator{contract: _StaderOracle.contract, event: "SafeModeDisabled", logs: logs, sub: sub}, nil
}

// WatchSafeModeDisabled is a free log subscription operation binding the contract event 0xf29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292.
//
// Solidity: event SafeModeDisabled()
func (_StaderOracle *StaderOracleFilterer) WatchSafeModeDisabled(opts *bind.WatchOpts, sink chan<- *StaderOracleSafeModeDisabled) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SafeModeDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSafeModeDisabled)
				if err := _StaderOracle.contract.UnpackLog(event, "SafeModeDisabled", log); err != nil {
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

// ParseSafeModeDisabled is a log parse operation binding the contract event 0xf29e751b3c28b619a215d25fee98a7828471a8e554626186d3f8d122f1650292.
//
// Solidity: event SafeModeDisabled()
func (_StaderOracle *StaderOracleFilterer) ParseSafeModeDisabled(log types.Log) (*StaderOracleSafeModeDisabled, error) {
	event := new(StaderOracleSafeModeDisabled)
	if err := _StaderOracle.contract.UnpackLog(event, "SafeModeDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSafeModeEnabledIterator is returned from FilterSafeModeEnabled and is used to iterate over the raw logs and unpacked data for SafeModeEnabled events raised by the StaderOracle contract.
type StaderOracleSafeModeEnabledIterator struct {
	Event *StaderOracleSafeModeEnabled // Event containing the contract specifics and raw log

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
func (it *StaderOracleSafeModeEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSafeModeEnabled)
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
		it.Event = new(StaderOracleSafeModeEnabled)
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
func (it *StaderOracleSafeModeEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSafeModeEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSafeModeEnabled represents a SafeModeEnabled event raised by the StaderOracle contract.
type StaderOracleSafeModeEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSafeModeEnabled is a free log retrieval operation binding the contract event 0x3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c.
//
// Solidity: event SafeModeEnabled()
func (_StaderOracle *StaderOracleFilterer) FilterSafeModeEnabled(opts *bind.FilterOpts) (*StaderOracleSafeModeEnabledIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SafeModeEnabled")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSafeModeEnabledIterator{contract: _StaderOracle.contract, event: "SafeModeEnabled", logs: logs, sub: sub}, nil
}

// WatchSafeModeEnabled is a free log subscription operation binding the contract event 0x3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c.
//
// Solidity: event SafeModeEnabled()
func (_StaderOracle *StaderOracleFilterer) WatchSafeModeEnabled(opts *bind.WatchOpts, sink chan<- *StaderOracleSafeModeEnabled) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SafeModeEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSafeModeEnabled)
				if err := _StaderOracle.contract.UnpackLog(event, "SafeModeEnabled", log); err != nil {
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

// ParseSafeModeEnabled is a log parse operation binding the contract event 0x3328bda355014adfb66d5d9086ab2c3204d1af7f83a69a3276daeed721f6ce3c.
//
// Solidity: event SafeModeEnabled()
func (_StaderOracle *StaderOracleFilterer) ParseSafeModeEnabled(log types.Log) (*StaderOracleSafeModeEnabled, error) {
	event := new(StaderOracleSafeModeEnabled)
	if err := _StaderOracle.contract.UnpackLog(event, "SafeModeEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSocializingRewardsMerkleRootSubmittedIterator is returned from FilterSocializingRewardsMerkleRootSubmitted and is used to iterate over the raw logs and unpacked data for SocializingRewardsMerkleRootSubmitted events raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootSubmittedIterator struct {
	Event *StaderOracleSocializingRewardsMerkleRootSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleSocializingRewardsMerkleRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSocializingRewardsMerkleRootSubmitted)
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
		it.Event = new(StaderOracleSocializingRewardsMerkleRootSubmitted)
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
func (it *StaderOracleSocializingRewardsMerkleRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSocializingRewardsMerkleRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSocializingRewardsMerkleRootSubmitted represents a SocializingRewardsMerkleRootSubmitted event raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootSubmitted struct {
	Node       common.Address
	Index      *big.Int
	MerkleRoot [32]byte
	PoolId     uint8
	Block      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSocializingRewardsMerkleRootSubmitted is a free log retrieval operation binding the contract event 0x97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd23377.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSocializingRewardsMerkleRootSubmitted(opts *bind.FilterOpts, node []common.Address) (*StaderOracleSocializingRewardsMerkleRootSubmittedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SocializingRewardsMerkleRootSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleSocializingRewardsMerkleRootSubmittedIterator{contract: _StaderOracle.contract, event: "SocializingRewardsMerkleRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchSocializingRewardsMerkleRootSubmitted is a free log subscription operation binding the contract event 0x97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd23377.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSocializingRewardsMerkleRootSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleSocializingRewardsMerkleRootSubmitted, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SocializingRewardsMerkleRootSubmitted", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSocializingRewardsMerkleRootSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootSubmitted", log); err != nil {
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

// ParseSocializingRewardsMerkleRootSubmitted is a log parse operation binding the contract event 0x97f29f2f9a3ad2e8ebffd3fb4a6dbf5035b92b432c8344609b8368407dd23377.
//
// Solidity: event SocializingRewardsMerkleRootSubmitted(address indexed node, uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSocializingRewardsMerkleRootSubmitted(log types.Log) (*StaderOracleSocializingRewardsMerkleRootSubmitted, error) {
	event := new(StaderOracleSocializingRewardsMerkleRootSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleSocializingRewardsMerkleRootUpdatedIterator is returned from FilterSocializingRewardsMerkleRootUpdated and is used to iterate over the raw logs and unpacked data for SocializingRewardsMerkleRootUpdated events raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootUpdatedIterator struct {
	Event *StaderOracleSocializingRewardsMerkleRootUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleSocializingRewardsMerkleRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleSocializingRewardsMerkleRootUpdated)
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
		it.Event = new(StaderOracleSocializingRewardsMerkleRootUpdated)
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
func (it *StaderOracleSocializingRewardsMerkleRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleSocializingRewardsMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleSocializingRewardsMerkleRootUpdated represents a SocializingRewardsMerkleRootUpdated event raised by the StaderOracle contract.
type StaderOracleSocializingRewardsMerkleRootUpdated struct {
	Index      *big.Int
	MerkleRoot [32]byte
	PoolId     uint8
	Block      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSocializingRewardsMerkleRootUpdated is a free log retrieval operation binding the contract event 0x4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) FilterSocializingRewardsMerkleRootUpdated(opts *bind.FilterOpts) (*StaderOracleSocializingRewardsMerkleRootUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "SocializingRewardsMerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleSocializingRewardsMerkleRootUpdatedIterator{contract: _StaderOracle.contract, event: "SocializingRewardsMerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSocializingRewardsMerkleRootUpdated is a free log subscription operation binding the contract event 0x4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) WatchSocializingRewardsMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleSocializingRewardsMerkleRootUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "SocializingRewardsMerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleSocializingRewardsMerkleRootUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootUpdated", log); err != nil {
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

// ParseSocializingRewardsMerkleRootUpdated is a log parse operation binding the contract event 0x4394ee7a4ca89453c6900058c69bfaf14014d9fc4d510ead54ae47ac06d1f05e.
//
// Solidity: event SocializingRewardsMerkleRootUpdated(uint256 index, bytes32 merkleRoot, uint8 poolId, uint256 block)
func (_StaderOracle *StaderOracleFilterer) ParseSocializingRewardsMerkleRootUpdated(log types.Log) (*StaderOracleSocializingRewardsMerkleRootUpdated, error) {
	event := new(StaderOracleSocializingRewardsMerkleRootUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "SocializingRewardsMerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleTrustedNodeAddedIterator is returned from FilterTrustedNodeAdded and is used to iterate over the raw logs and unpacked data for TrustedNodeAdded events raised by the StaderOracle contract.
type StaderOracleTrustedNodeAddedIterator struct {
	Event *StaderOracleTrustedNodeAdded // Event containing the contract specifics and raw log

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
func (it *StaderOracleTrustedNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleTrustedNodeAdded)
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
		it.Event = new(StaderOracleTrustedNodeAdded)
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
func (it *StaderOracleTrustedNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleTrustedNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleTrustedNodeAdded represents a TrustedNodeAdded event raised by the StaderOracle contract.
type StaderOracleTrustedNodeAdded struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedNodeAdded is a free log retrieval operation binding the contract event 0xff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8.
//
// Solidity: event TrustedNodeAdded(address indexed node)
func (_StaderOracle *StaderOracleFilterer) FilterTrustedNodeAdded(opts *bind.FilterOpts, node []common.Address) (*StaderOracleTrustedNodeAddedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "TrustedNodeAdded", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleTrustedNodeAddedIterator{contract: _StaderOracle.contract, event: "TrustedNodeAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedNodeAdded is a free log subscription operation binding the contract event 0xff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8.
//
// Solidity: event TrustedNodeAdded(address indexed node)
func (_StaderOracle *StaderOracleFilterer) WatchTrustedNodeAdded(opts *bind.WatchOpts, sink chan<- *StaderOracleTrustedNodeAdded, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "TrustedNodeAdded", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleTrustedNodeAdded)
				if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeAdded", log); err != nil {
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

// ParseTrustedNodeAdded is a log parse operation binding the contract event 0xff4a290f0500d113133708d378eb9a151c32d91cb8f5778cfda6328d89cfc4b8.
//
// Solidity: event TrustedNodeAdded(address indexed node)
func (_StaderOracle *StaderOracleFilterer) ParseTrustedNodeAdded(log types.Log) (*StaderOracleTrustedNodeAdded, error) {
	event := new(StaderOracleTrustedNodeAdded)
	if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleTrustedNodeRemovedIterator is returned from FilterTrustedNodeRemoved and is used to iterate over the raw logs and unpacked data for TrustedNodeRemoved events raised by the StaderOracle contract.
type StaderOracleTrustedNodeRemovedIterator struct {
	Event *StaderOracleTrustedNodeRemoved // Event containing the contract specifics and raw log

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
func (it *StaderOracleTrustedNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleTrustedNodeRemoved)
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
		it.Event = new(StaderOracleTrustedNodeRemoved)
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
func (it *StaderOracleTrustedNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleTrustedNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleTrustedNodeRemoved represents a TrustedNodeRemoved event raised by the StaderOracle contract.
type StaderOracleTrustedNodeRemoved struct {
	Node common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedNodeRemoved is a free log retrieval operation binding the contract event 0x6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421.
//
// Solidity: event TrustedNodeRemoved(address indexed node)
func (_StaderOracle *StaderOracleFilterer) FilterTrustedNodeRemoved(opts *bind.FilterOpts, node []common.Address) (*StaderOracleTrustedNodeRemovedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "TrustedNodeRemoved", nodeRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleTrustedNodeRemovedIterator{contract: _StaderOracle.contract, event: "TrustedNodeRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedNodeRemoved is a free log subscription operation binding the contract event 0x6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421.
//
// Solidity: event TrustedNodeRemoved(address indexed node)
func (_StaderOracle *StaderOracleFilterer) WatchTrustedNodeRemoved(opts *bind.WatchOpts, sink chan<- *StaderOracleTrustedNodeRemoved, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "TrustedNodeRemoved", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleTrustedNodeRemoved)
				if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeRemoved", log); err != nil {
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

// ParseTrustedNodeRemoved is a log parse operation binding the contract event 0x6d1de2fb0c63996bae7ba6277c720c0560ba42874ea34c1ebe8e1423b9b47421.
//
// Solidity: event TrustedNodeRemoved(address indexed node)
func (_StaderOracle *StaderOracleFilterer) ParseTrustedNodeRemoved(log types.Log) (*StaderOracleTrustedNodeRemoved, error) {
	event := new(StaderOracleTrustedNodeRemoved)
	if err := _StaderOracle.contract.UnpackLog(event, "TrustedNodeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StaderOracle contract.
type StaderOracleUnpausedIterator struct {
	Event *StaderOracleUnpaused // Event containing the contract specifics and raw log

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
func (it *StaderOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUnpaused)
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
		it.Event = new(StaderOracleUnpaused)
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
func (it *StaderOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUnpaused represents a Unpaused event raised by the StaderOracle contract.
type StaderOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StaderOracle *StaderOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StaderOracleUnpausedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUnpausedIterator{contract: _StaderOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StaderOracle *StaderOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StaderOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUnpaused)
				if err := _StaderOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StaderOracle *StaderOracleFilterer) ParseUnpaused(log types.Log) (*StaderOracleUnpaused, error) {
	event := new(StaderOracleUnpaused)
	if err := _StaderOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUpdateFrequencyUpdatedIterator is returned from FilterUpdateFrequencyUpdated and is used to iterate over the raw logs and unpacked data for UpdateFrequencyUpdated events raised by the StaderOracle contract.
type StaderOracleUpdateFrequencyUpdatedIterator struct {
	Event *StaderOracleUpdateFrequencyUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdateFrequencyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdateFrequencyUpdated)
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
		it.Event = new(StaderOracleUpdateFrequencyUpdated)
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
func (it *StaderOracleUpdateFrequencyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdateFrequencyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdateFrequencyUpdated represents a UpdateFrequencyUpdated event raised by the StaderOracle contract.
type StaderOracleUpdateFrequencyUpdated struct {
	UpdateFrequency *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateFrequencyUpdated is a free log retrieval operation binding the contract event 0x6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d391138.
//
// Solidity: event UpdateFrequencyUpdated(uint256 updateFrequency)
func (_StaderOracle *StaderOracleFilterer) FilterUpdateFrequencyUpdated(opts *bind.FilterOpts) (*StaderOracleUpdateFrequencyUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdateFrequencyUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdateFrequencyUpdatedIterator{contract: _StaderOracle.contract, event: "UpdateFrequencyUpdated", logs: logs, sub: sub}, nil
}

// WatchUpdateFrequencyUpdated is a free log subscription operation binding the contract event 0x6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d391138.
//
// Solidity: event UpdateFrequencyUpdated(uint256 updateFrequency)
func (_StaderOracle *StaderOracleFilterer) WatchUpdateFrequencyUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdateFrequencyUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdateFrequencyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdateFrequencyUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdateFrequencyUpdated", log); err != nil {
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

// ParseUpdateFrequencyUpdated is a log parse operation binding the contract event 0x6231a3e049e2072a042ae727208e7650b487871f4080458371e84d6e7d391138.
//
// Solidity: event UpdateFrequencyUpdated(uint256 updateFrequency)
func (_StaderOracle *StaderOracleFilterer) ParseUpdateFrequencyUpdated(log types.Log) (*StaderOracleUpdateFrequencyUpdated, error) {
	event := new(StaderOracleUpdateFrequencyUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdateFrequencyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUpdatedERChangeLimitIterator is returned from FilterUpdatedERChangeLimit and is used to iterate over the raw logs and unpacked data for UpdatedERChangeLimit events raised by the StaderOracle contract.
type StaderOracleUpdatedERChangeLimitIterator struct {
	Event *StaderOracleUpdatedERChangeLimit // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdatedERChangeLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdatedERChangeLimit)
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
		it.Event = new(StaderOracleUpdatedERChangeLimit)
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
func (it *StaderOracleUpdatedERChangeLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdatedERChangeLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdatedERChangeLimit represents a UpdatedERChangeLimit event raised by the StaderOracle contract.
type StaderOracleUpdatedERChangeLimit struct {
	ErChangeLimit *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedERChangeLimit is a free log retrieval operation binding the contract event 0x94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520.
//
// Solidity: event UpdatedERChangeLimit(uint256 erChangeLimit)
func (_StaderOracle *StaderOracleFilterer) FilterUpdatedERChangeLimit(opts *bind.FilterOpts) (*StaderOracleUpdatedERChangeLimitIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdatedERChangeLimit")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdatedERChangeLimitIterator{contract: _StaderOracle.contract, event: "UpdatedERChangeLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedERChangeLimit is a free log subscription operation binding the contract event 0x94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520.
//
// Solidity: event UpdatedERChangeLimit(uint256 erChangeLimit)
func (_StaderOracle *StaderOracleFilterer) WatchUpdatedERChangeLimit(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdatedERChangeLimit) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdatedERChangeLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdatedERChangeLimit)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdatedERChangeLimit", log); err != nil {
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

// ParseUpdatedERChangeLimit is a log parse operation binding the contract event 0x94a97bfc9c7a83fe5f75c66931ca7d2d16372fdc244afc5db36044f3ca52a520.
//
// Solidity: event UpdatedERChangeLimit(uint256 erChangeLimit)
func (_StaderOracle *StaderOracleFilterer) ParseUpdatedERChangeLimit(log types.Log) (*StaderOracleUpdatedERChangeLimit, error) {
	event := new(StaderOracleUpdatedERChangeLimit)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdatedERChangeLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the StaderOracle contract.
type StaderOracleUpdatedStaderConfigIterator struct {
	Event *StaderOracleUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *StaderOracleUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleUpdatedStaderConfig)
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
		it.Event = new(StaderOracleUpdatedStaderConfig)
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
func (it *StaderOracleUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the StaderOracle contract.
type StaderOracleUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StaderOracle *StaderOracleFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*StaderOracleUpdatedStaderConfigIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &StaderOracleUpdatedStaderConfigIterator{contract: _StaderOracle.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StaderOracle *StaderOracleFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *StaderOracleUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleUpdatedStaderConfig)
				if err := _StaderOracle.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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

// ParseUpdatedStaderConfig is a log parse operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_StaderOracle *StaderOracleFilterer) ParseUpdatedStaderConfig(log types.Log) (*StaderOracleUpdatedStaderConfig, error) {
	event := new(StaderOracleUpdatedStaderConfig)
	if err := _StaderOracle.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleValidatorStatsSubmittedIterator is returned from FilterValidatorStatsSubmitted and is used to iterate over the raw logs and unpacked data for ValidatorStatsSubmitted events raised by the StaderOracle contract.
type StaderOracleValidatorStatsSubmittedIterator struct {
	Event *StaderOracleValidatorStatsSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleValidatorStatsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleValidatorStatsSubmitted)
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
		it.Event = new(StaderOracleValidatorStatsSubmitted)
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
func (it *StaderOracleValidatorStatsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleValidatorStatsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleValidatorStatsSubmitted represents a ValidatorStatsSubmitted event raised by the StaderOracle contract.
type StaderOracleValidatorStatsSubmitted struct {
	From                     common.Address
	Block                    *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    *big.Int
	ExitedValidatorsCount    *big.Int
	SlashedValidatorsCount   *big.Int
	Time                     *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatsSubmitted is a free log retrieval operation binding the contract event 0x72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc.
//
// Solidity: event ValidatorStatsSubmitted(address indexed from, uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterValidatorStatsSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleValidatorStatsSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ValidatorStatsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleValidatorStatsSubmittedIterator{contract: _StaderOracle.contract, event: "ValidatorStatsSubmitted", logs: logs, sub: sub}, nil
}

// WatchValidatorStatsSubmitted is a free log subscription operation binding the contract event 0x72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc.
//
// Solidity: event ValidatorStatsSubmitted(address indexed from, uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchValidatorStatsSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleValidatorStatsSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ValidatorStatsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleValidatorStatsSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsSubmitted", log); err != nil {
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

// ParseValidatorStatsSubmitted is a log parse operation binding the contract event 0x72745b0271618e5d84d738ea416e3a3be6eb267e0f639198f63c0ef124c29ffc.
//
// Solidity: event ValidatorStatsSubmitted(address indexed from, uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseValidatorStatsSubmitted(log types.Log) (*StaderOracleValidatorStatsSubmitted, error) {
	event := new(StaderOracleValidatorStatsSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleValidatorStatsUpdatedIterator is returned from FilterValidatorStatsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorStatsUpdated events raised by the StaderOracle contract.
type StaderOracleValidatorStatsUpdatedIterator struct {
	Event *StaderOracleValidatorStatsUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleValidatorStatsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleValidatorStatsUpdated)
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
		it.Event = new(StaderOracleValidatorStatsUpdated)
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
func (it *StaderOracleValidatorStatsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleValidatorStatsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleValidatorStatsUpdated represents a ValidatorStatsUpdated event raised by the StaderOracle contract.
type StaderOracleValidatorStatsUpdated struct {
	Block                    *big.Int
	ActiveValidatorsBalance  *big.Int
	ExitedValidatorsBalance  *big.Int
	SlashedValidatorsBalance *big.Int
	ActiveValidatorsCount    *big.Int
	ExitedValidatorsCount    *big.Int
	SlashedValidatorsCount   *big.Int
	Time                     *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatsUpdated is a free log retrieval operation binding the contract event 0x6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c94.
//
// Solidity: event ValidatorStatsUpdated(uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterValidatorStatsUpdated(opts *bind.FilterOpts) (*StaderOracleValidatorStatsUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "ValidatorStatsUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleValidatorStatsUpdatedIterator{contract: _StaderOracle.contract, event: "ValidatorStatsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorStatsUpdated is a free log subscription operation binding the contract event 0x6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c94.
//
// Solidity: event ValidatorStatsUpdated(uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchValidatorStatsUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleValidatorStatsUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "ValidatorStatsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleValidatorStatsUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsUpdated", log); err != nil {
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

// ParseValidatorStatsUpdated is a log parse operation binding the contract event 0x6988248fd82a7ce842fbdce0c49889ad926bf1548bae4194de0006498d069c94.
//
// Solidity: event ValidatorStatsUpdated(uint256 block, uint256 activeValidatorsBalance, uint256 exitedValidatorsBalance, uint256 slashedValidatorsBalance, uint256 activeValidatorsCount, uint256 exitedValidatorsCount, uint256 slashedValidatorsCount, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseValidatorStatsUpdated(log types.Log) (*StaderOracleValidatorStatsUpdated, error) {
	event := new(StaderOracleValidatorStatsUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "ValidatorStatsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleWithdrawnValidatorsSubmittedIterator is returned from FilterWithdrawnValidatorsSubmitted and is used to iterate over the raw logs and unpacked data for WithdrawnValidatorsSubmitted events raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsSubmittedIterator struct {
	Event *StaderOracleWithdrawnValidatorsSubmitted // Event containing the contract specifics and raw log

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
func (it *StaderOracleWithdrawnValidatorsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleWithdrawnValidatorsSubmitted)
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
		it.Event = new(StaderOracleWithdrawnValidatorsSubmitted)
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
func (it *StaderOracleWithdrawnValidatorsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleWithdrawnValidatorsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleWithdrawnValidatorsSubmitted represents a WithdrawnValidatorsSubmitted event raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsSubmitted struct {
	From         common.Address
	Block        *big.Int
	NodeRegistry common.Address
	Pubkeys      [][]byte
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnValidatorsSubmitted is a free log retrieval operation binding the contract event 0x1e004d900c9787bf2cfed87d7704cf2a6b1956a7141dab327aea41eb6aa143f2.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterWithdrawnValidatorsSubmitted(opts *bind.FilterOpts, from []common.Address) (*StaderOracleWithdrawnValidatorsSubmittedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "WithdrawnValidatorsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return &StaderOracleWithdrawnValidatorsSubmittedIterator{contract: _StaderOracle.contract, event: "WithdrawnValidatorsSubmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnValidatorsSubmitted is a free log subscription operation binding the contract event 0x1e004d900c9787bf2cfed87d7704cf2a6b1956a7141dab327aea41eb6aa143f2.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchWithdrawnValidatorsSubmitted(opts *bind.WatchOpts, sink chan<- *StaderOracleWithdrawnValidatorsSubmitted, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "WithdrawnValidatorsSubmitted", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleWithdrawnValidatorsSubmitted)
				if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsSubmitted", log); err != nil {
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

// ParseWithdrawnValidatorsSubmitted is a log parse operation binding the contract event 0x1e004d900c9787bf2cfed87d7704cf2a6b1956a7141dab327aea41eb6aa143f2.
//
// Solidity: event WithdrawnValidatorsSubmitted(address indexed from, uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseWithdrawnValidatorsSubmitted(log types.Log) (*StaderOracleWithdrawnValidatorsSubmitted, error) {
	event := new(StaderOracleWithdrawnValidatorsSubmitted)
	if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StaderOracleWithdrawnValidatorsUpdatedIterator is returned from FilterWithdrawnValidatorsUpdated and is used to iterate over the raw logs and unpacked data for WithdrawnValidatorsUpdated events raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsUpdatedIterator struct {
	Event *StaderOracleWithdrawnValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *StaderOracleWithdrawnValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StaderOracleWithdrawnValidatorsUpdated)
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
		it.Event = new(StaderOracleWithdrawnValidatorsUpdated)
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
func (it *StaderOracleWithdrawnValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StaderOracleWithdrawnValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StaderOracleWithdrawnValidatorsUpdated represents a WithdrawnValidatorsUpdated event raised by the StaderOracle contract.
type StaderOracleWithdrawnValidatorsUpdated struct {
	Block        *big.Int
	NodeRegistry common.Address
	Pubkeys      [][]byte
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnValidatorsUpdated is a free log retrieval operation binding the contract event 0x70a26ae3b547e42cc7a3819a181db1a108440c2dedd23eb34c2357680cfd056b.
//
// Solidity: event WithdrawnValidatorsUpdated(uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) FilterWithdrawnValidatorsUpdated(opts *bind.FilterOpts) (*StaderOracleWithdrawnValidatorsUpdatedIterator, error) {

	logs, sub, err := _StaderOracle.contract.FilterLogs(opts, "WithdrawnValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &StaderOracleWithdrawnValidatorsUpdatedIterator{contract: _StaderOracle.contract, event: "WithdrawnValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawnValidatorsUpdated is a free log subscription operation binding the contract event 0x70a26ae3b547e42cc7a3819a181db1a108440c2dedd23eb34c2357680cfd056b.
//
// Solidity: event WithdrawnValidatorsUpdated(uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) WatchWithdrawnValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *StaderOracleWithdrawnValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _StaderOracle.contract.WatchLogs(opts, "WithdrawnValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StaderOracleWithdrawnValidatorsUpdated)
				if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsUpdated", log); err != nil {
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

// ParseWithdrawnValidatorsUpdated is a log parse operation binding the contract event 0x70a26ae3b547e42cc7a3819a181db1a108440c2dedd23eb34c2357680cfd056b.
//
// Solidity: event WithdrawnValidatorsUpdated(uint256 block, address nodeRegistry, bytes[] pubkeys, uint256 time)
func (_StaderOracle *StaderOracleFilterer) ParseWithdrawnValidatorsUpdated(log types.Log) (*StaderOracleWithdrawnValidatorsUpdated, error) {
	event := new(StaderOracleWithdrawnValidatorsUpdated)
	if err := _StaderOracle.contract.UnpackLog(event, "WithdrawnValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
