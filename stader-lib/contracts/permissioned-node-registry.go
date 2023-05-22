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

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}

// PermissionedNodeRegistryMetaData contains all meta data concerning the PermissionedNodeRegistry contract.
var PermissionedNodeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartAndEndIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxOperatorLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MisMatchingInputKeysSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAPermissionedNodeOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughSDCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyDeactivate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyOnBoardedInProtocol\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorIsDeactivate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotOnBoarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PageNumberIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyAlreadyExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyVerifiedKeysReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyWithdrawnKeysReported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UNEXPECTED_STATUS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"maxKeyLimitReached\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"AddedValidatorKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalActiveValidatorCount\",\"type\":\"uint256\"}],\"name\":\"DecreasedTotalActiveValidatorCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalActiveValidatorCount\",\"type\":\"uint256\"}],\"name\":\"IncreasedTotalActiveValidatorCount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"MarkedValidatorStatusAsPreDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxOperatorId\",\"type\":\"uint256\"}],\"name\":\"MaxOperatorIdLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nodeRewardAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"}],\"name\":\"OnboardedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorID\",\"type\":\"uint256\"}],\"name\":\"OperatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorID\",\"type\":\"uint256\"}],\"name\":\"OperatorDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"permissionedNO\",\"type\":\"address\"}],\"name\":\"OperatorWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batchKeyDepositLimit\",\"type\":\"uint256\"}],\"name\":\"UpdatedInputKeyCountLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxNonTerminalKeyPerOperator\",\"type\":\"uint64\"}],\"name\":\"UpdatedMaxNonTerminalKeyPerOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"UpdatedOperatorDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextQueuedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"UpdatedQueuedValidatorIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"}],\"name\":\"UpdatedValidatorDepositBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verifiedKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"UpdatedVerifiedKeyBatchSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"UpdatedWithdrawnKeyBatchSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorMarkedAsFrontRunned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorStatusMarkedAsInvalidSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"}],\"name\":\"ValidatorWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"activateNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_preDepositSignature\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_depositSignature\",\"type\":\"bytes[]\"}],\"name\":\"addValidatorKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numValidators\",\"type\":\"uint256\"}],\"name\":\"allocateValidatorsAndUpdateOperatorId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"selectedOperatorCapacity\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"deactivateNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"}],\"name\":\"getAllActiveValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnBlock\",\"type\":\"uint256\"}],\"internalType\":\"structValidator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorRewardAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalKeys\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"}],\"name\":\"getSocializingPoolStateChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQueuedValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"increaseTotalActiveValidatorCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inputKeyCountLimit\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"isExistingOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_readyToDepositPubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_frontRunPubkey\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_invalidSignaturePubkey\",\"type\":\"bytes[]\"}],\"name\":\"markValidatorReadyToDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"markValidatorStatusAsPreDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNonTerminalKeyPerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nextQueuedValidatorIndexByOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_operatorRewardAddress\",\"type\":\"address\"}],\"name\":\"onboardNodeOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"onlyPreDepositValidator\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorIDByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorIdForExcessDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorStructById\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"optedForSocializingPool\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"operatorRewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissionList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"socializingPoolStateChangeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveOperatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"}],\"name\":\"updateDepositStatusAndBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_inputKeyCountLimit\",\"type\":\"uint16\"}],\"name\":\"updateInputKeyCountLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_maxNonTerminalKeyPerOperator\",\"type\":\"uint64\"}],\"name\":\"updateMaxNonTerminalKeyPerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxOperatorId\",\"type\":\"uint256\"}],\"name\":\"updateMaxOperatorId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_operatorName\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_rewardAddress\",\"type\":\"address\"}],\"name\":\"updateOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextQueuedValidatorIndex\",\"type\":\"uint256\"}],\"name\":\"updateQueuedValidatorIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_verifiedKeysBatchSize\",\"type\":\"uint256\"}],\"name\":\"updateVerifiedKeysBatchSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validatorIdByPubkey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorIdsByOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"enumValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"preDepositSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"withdrawVaultAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"operatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifiedKeyBatchSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_permissionedNOs\",\"type\":\"address[]\"}],\"name\":\"whitelistPermissionedNOs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"withdrawnValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PermissionedNodeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionedNodeRegistryMetaData.ABI instead.
var PermissionedNodeRegistryABI = PermissionedNodeRegistryMetaData.ABI

// PermissionedNodeRegistry is an auto generated Go binding around an Ethereum contract.
type PermissionedNodeRegistry struct {
	PermissionedNodeRegistryCaller     // Read-only binding to the contract
	PermissionedNodeRegistryTransactor // Write-only binding to the contract
	PermissionedNodeRegistryFilterer   // Log filterer for contract events
}

// PermissionedNodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionedNodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedNodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionedNodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedNodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionedNodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedNodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionedNodeRegistrySession struct {
	Contract     *PermissionedNodeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PermissionedNodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionedNodeRegistryCallerSession struct {
	Contract *PermissionedNodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// PermissionedNodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionedNodeRegistryTransactorSession struct {
	Contract     *PermissionedNodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// PermissionedNodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionedNodeRegistryRaw struct {
	Contract *PermissionedNodeRegistry // Generic contract binding to access the raw methods on
}

// PermissionedNodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionedNodeRegistryCallerRaw struct {
	Contract *PermissionedNodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionedNodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionedNodeRegistryTransactorRaw struct {
	Contract *PermissionedNodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionedNodeRegistry creates a new instance of PermissionedNodeRegistry, bound to a specific deployed contract.
func NewPermissionedNodeRegistry(address common.Address, backend bind.ContractBackend) (*PermissionedNodeRegistry, error) {
	contract, err := bindPermissionedNodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistry{PermissionedNodeRegistryCaller: PermissionedNodeRegistryCaller{contract: contract}, PermissionedNodeRegistryTransactor: PermissionedNodeRegistryTransactor{contract: contract}, PermissionedNodeRegistryFilterer: PermissionedNodeRegistryFilterer{contract: contract}}, nil
}

// NewPermissionedNodeRegistryCaller creates a new read-only instance of PermissionedNodeRegistry, bound to a specific deployed contract.
func NewPermissionedNodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*PermissionedNodeRegistryCaller, error) {
	contract, err := bindPermissionedNodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryCaller{contract: contract}, nil
}

// NewPermissionedNodeRegistryTransactor creates a new write-only instance of PermissionedNodeRegistry, bound to a specific deployed contract.
func NewPermissionedNodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionedNodeRegistryTransactor, error) {
	contract, err := bindPermissionedNodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryTransactor{contract: contract}, nil
}

// NewPermissionedNodeRegistryFilterer creates a new log filterer instance of PermissionedNodeRegistry, bound to a specific deployed contract.
func NewPermissionedNodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionedNodeRegistryFilterer, error) {
	contract, err := bindPermissionedNodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryFilterer{contract: contract}, nil
}

// bindPermissionedNodeRegistry binds a generic wrapper to an already deployed contract.
func bindPermissionedNodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionedNodeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedNodeRegistry *PermissionedNodeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedNodeRegistry.Contract.PermissionedNodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedNodeRegistry *PermissionedNodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.PermissionedNodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedNodeRegistry *PermissionedNodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.PermissionedNodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedNodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionedNodeRegistry.Contract.DEFAULTADMINROLE(&_PermissionedNodeRegistry.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionedNodeRegistry.Contract.DEFAULTADMINROLE(&_PermissionedNodeRegistry.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) POOLID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "POOL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) POOLID() (uint8, error) {
	return _PermissionedNodeRegistry.Contract.POOLID(&_PermissionedNodeRegistry.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) POOLID() (uint8, error) {
	return _PermissionedNodeRegistry.Contract.POOLID(&_PermissionedNodeRegistry.CallOpts)
}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0x99888898.
//
// Solidity: function getAllActiveValidators(uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetAllActiveValidators(opts *bind.CallOpts, _pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getAllActiveValidators", _pageNumber, _pageSize)

	if err != nil {
		return *new([]Validator), err
	}

	out0 := *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)

	return out0, err

}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0x99888898.
//
// Solidity: function getAllActiveValidators(uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetAllActiveValidators(_pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	return _PermissionedNodeRegistry.Contract.GetAllActiveValidators(&_PermissionedNodeRegistry.CallOpts, _pageNumber, _pageSize)
}

// GetAllActiveValidators is a free data retrieval call binding the contract method 0x99888898.
//
// Solidity: function getAllActiveValidators(uint256 _pageNumber, uint256 _pageSize) view returns((uint8,bytes,bytes,bytes,address,uint256,uint256,uint256)[])
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetAllActiveValidators(_pageNumber *big.Int, _pageSize *big.Int) ([]Validator, error) {
	return _PermissionedNodeRegistry.Contract.GetAllActiveValidators(&_PermissionedNodeRegistry.CallOpts, _pageNumber, _pageSize)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() pure returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetCollateralETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getCollateralETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() pure returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetCollateralETH() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetCollateralETH(&_PermissionedNodeRegistry.CallOpts)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() pure returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetCollateralETH() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetCollateralETH(&_PermissionedNodeRegistry.CallOpts)
}

// GetOperatorRewardAddress is a free data retrieval call binding the contract method 0x83ea2358.
//
// Solidity: function getOperatorRewardAddress(uint256 _operatorId) view returns(address)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetOperatorRewardAddress(opts *bind.CallOpts, _operatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getOperatorRewardAddress", _operatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorRewardAddress is a free data retrieval call binding the contract method 0x83ea2358.
//
// Solidity: function getOperatorRewardAddress(uint256 _operatorId) view returns(address)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetOperatorRewardAddress(_operatorId *big.Int) (common.Address, error) {
	return _PermissionedNodeRegistry.Contract.GetOperatorRewardAddress(&_PermissionedNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorRewardAddress is a free data retrieval call binding the contract method 0x83ea2358.
//
// Solidity: function getOperatorRewardAddress(uint256 _operatorId) view returns(address)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetOperatorRewardAddress(_operatorId *big.Int) (common.Address, error) {
	return _PermissionedNodeRegistry.Contract.GetOperatorRewardAddress(&_PermissionedNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorTotalKeys is a free data retrieval call binding the contract method 0xc34ade5c.
//
// Solidity: function getOperatorTotalKeys(uint256 _operatorId) view returns(uint256 _totalKeys)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetOperatorTotalKeys(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getOperatorTotalKeys", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalKeys is a free data retrieval call binding the contract method 0xc34ade5c.
//
// Solidity: function getOperatorTotalKeys(uint256 _operatorId) view returns(uint256 _totalKeys)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetOperatorTotalKeys(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetOperatorTotalKeys(&_PermissionedNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorTotalKeys is a free data retrieval call binding the contract method 0xc34ade5c.
//
// Solidity: function getOperatorTotalKeys(uint256 _operatorId) view returns(uint256 _totalKeys)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetOperatorTotalKeys(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetOperatorTotalKeys(&_PermissionedNodeRegistry.CallOpts, _operatorId)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint64)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (uint64, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _nodeOperator, _startIndex, _endIndex)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint64)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (uint64, error) {
	return _PermissionedNodeRegistry.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionedNodeRegistry.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint64)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (uint64, error) {
	return _PermissionedNodeRegistry.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionedNodeRegistry.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionedNodeRegistry.Contract.GetRoleAdmin(&_PermissionedNodeRegistry.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionedNodeRegistry.Contract.GetRoleAdmin(&_PermissionedNodeRegistry.CallOpts, role)
}

// GetSocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0xebb5c174.
//
// Solidity: function getSocializingPoolStateChangeBlock(uint256 _operatorId) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetSocializingPoolStateChangeBlock(opts *bind.CallOpts, _operatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getSocializingPoolStateChangeBlock", _operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0xebb5c174.
//
// Solidity: function getSocializingPoolStateChangeBlock(uint256 _operatorId) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetSocializingPoolStateChangeBlock(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetSocializingPoolStateChangeBlock(&_PermissionedNodeRegistry.CallOpts, _operatorId)
}

// GetSocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0xebb5c174.
//
// Solidity: function getSocializingPoolStateChangeBlock(uint256 _operatorId) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetSocializingPoolStateChangeBlock(_operatorId *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetSocializingPoolStateChangeBlock(&_PermissionedNodeRegistry.CallOpts, _operatorId)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetTotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getTotalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetTotalActiveValidatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetTotalActiveValidatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) GetTotalQueuedValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "getTotalQueuedValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetTotalQueuedValidatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.GetTotalQueuedValidatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionedNodeRegistry.Contract.HasRole(&_PermissionedNodeRegistry.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionedNodeRegistry.Contract.HasRole(&_PermissionedNodeRegistry.CallOpts, role, account)
}

// InputKeyCountLimit is a free data retrieval call binding the contract method 0xe0bf8b53.
//
// Solidity: function inputKeyCountLimit() view returns(uint16)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) InputKeyCountLimit(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "inputKeyCountLimit")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InputKeyCountLimit is a free data retrieval call binding the contract method 0xe0bf8b53.
//
// Solidity: function inputKeyCountLimit() view returns(uint16)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) InputKeyCountLimit() (uint16, error) {
	return _PermissionedNodeRegistry.Contract.InputKeyCountLimit(&_PermissionedNodeRegistry.CallOpts)
}

// InputKeyCountLimit is a free data retrieval call binding the contract method 0xe0bf8b53.
//
// Solidity: function inputKeyCountLimit() view returns(uint16)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) InputKeyCountLimit() (uint16, error) {
	return _PermissionedNodeRegistry.Contract.InputKeyCountLimit(&_PermissionedNodeRegistry.CallOpts)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) IsExistingOperator(opts *bind.CallOpts, _operAddr common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "isExistingOperator", _operAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionedNodeRegistry.Contract.IsExistingOperator(&_PermissionedNodeRegistry.CallOpts, _operAddr)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionedNodeRegistry.Contract.IsExistingOperator(&_PermissionedNodeRegistry.CallOpts, _operAddr)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) IsExistingPubkey(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "isExistingPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionedNodeRegistry.Contract.IsExistingPubkey(&_PermissionedNodeRegistry.CallOpts, _pubkey)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionedNodeRegistry.Contract.IsExistingPubkey(&_PermissionedNodeRegistry.CallOpts, _pubkey)
}

// MaxNonTerminalKeyPerOperator is a free data retrieval call binding the contract method 0x50d5d7ab.
//
// Solidity: function maxNonTerminalKeyPerOperator() view returns(uint64)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) MaxNonTerminalKeyPerOperator(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "maxNonTerminalKeyPerOperator")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxNonTerminalKeyPerOperator is a free data retrieval call binding the contract method 0x50d5d7ab.
//
// Solidity: function maxNonTerminalKeyPerOperator() view returns(uint64)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) MaxNonTerminalKeyPerOperator() (uint64, error) {
	return _PermissionedNodeRegistry.Contract.MaxNonTerminalKeyPerOperator(&_PermissionedNodeRegistry.CallOpts)
}

// MaxNonTerminalKeyPerOperator is a free data retrieval call binding the contract method 0x50d5d7ab.
//
// Solidity: function maxNonTerminalKeyPerOperator() view returns(uint64)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) MaxNonTerminalKeyPerOperator() (uint64, error) {
	return _PermissionedNodeRegistry.Contract.MaxNonTerminalKeyPerOperator(&_PermissionedNodeRegistry.CallOpts)
}

// MaxOperatorId is a free data retrieval call binding the contract method 0x7ba89581.
//
// Solidity: function maxOperatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) MaxOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "maxOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxOperatorId is a free data retrieval call binding the contract method 0x7ba89581.
//
// Solidity: function maxOperatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) MaxOperatorId() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.MaxOperatorId(&_PermissionedNodeRegistry.CallOpts)
}

// MaxOperatorId is a free data retrieval call binding the contract method 0x7ba89581.
//
// Solidity: function maxOperatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) MaxOperatorId() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.MaxOperatorId(&_PermissionedNodeRegistry.CallOpts)
}

// NextOperatorId is a free data retrieval call binding the contract method 0x2d1dbd74.
//
// Solidity: function nextOperatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) NextOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "nextOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOperatorId is a free data retrieval call binding the contract method 0x2d1dbd74.
//
// Solidity: function nextOperatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) NextOperatorId() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.NextOperatorId(&_PermissionedNodeRegistry.CallOpts)
}

// NextOperatorId is a free data retrieval call binding the contract method 0x2d1dbd74.
//
// Solidity: function nextOperatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) NextOperatorId() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.NextOperatorId(&_PermissionedNodeRegistry.CallOpts)
}

// NextQueuedValidatorIndexByOperatorId is a free data retrieval call binding the contract method 0xe23c1f45.
//
// Solidity: function nextQueuedValidatorIndexByOperatorId(uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) NextQueuedValidatorIndexByOperatorId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "nextQueuedValidatorIndexByOperatorId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextQueuedValidatorIndexByOperatorId is a free data retrieval call binding the contract method 0xe23c1f45.
//
// Solidity: function nextQueuedValidatorIndexByOperatorId(uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) NextQueuedValidatorIndexByOperatorId(arg0 *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.NextQueuedValidatorIndexByOperatorId(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// NextQueuedValidatorIndexByOperatorId is a free data retrieval call binding the contract method 0xe23c1f45.
//
// Solidity: function nextQueuedValidatorIndexByOperatorId(uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) NextQueuedValidatorIndexByOperatorId(arg0 *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.NextQueuedValidatorIndexByOperatorId(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) NextValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "nextValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) NextValidatorId() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.NextValidatorId(&_PermissionedNodeRegistry.CallOpts)
}

// NextValidatorId is a free data retrieval call binding the contract method 0xf7c09189.
//
// Solidity: function nextValidatorId() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) NextValidatorId() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.NextValidatorId(&_PermissionedNodeRegistry.CallOpts)
}

// OnlyPreDepositValidator is a free data retrieval call binding the contract method 0x1b02ff40.
//
// Solidity: function onlyPreDepositValidator(bytes _pubkey) view returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) OnlyPreDepositValidator(opts *bind.CallOpts, _pubkey []byte) error {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "onlyPreDepositValidator", _pubkey)

	if err != nil {
		return err
	}

	return err

}

// OnlyPreDepositValidator is a free data retrieval call binding the contract method 0x1b02ff40.
//
// Solidity: function onlyPreDepositValidator(bytes _pubkey) view returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) OnlyPreDepositValidator(_pubkey []byte) error {
	return _PermissionedNodeRegistry.Contract.OnlyPreDepositValidator(&_PermissionedNodeRegistry.CallOpts, _pubkey)
}

// OnlyPreDepositValidator is a free data retrieval call binding the contract method 0x1b02ff40.
//
// Solidity: function onlyPreDepositValidator(bytes _pubkey) view returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) OnlyPreDepositValidator(_pubkey []byte) error {
	return _PermissionedNodeRegistry.Contract.OnlyPreDepositValidator(&_PermissionedNodeRegistry.CallOpts, _pubkey)
}

// OperatorIDByAddress is a free data retrieval call binding the contract method 0xcac8b306.
//
// Solidity: function operatorIDByAddress(address ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) OperatorIDByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "operatorIDByAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorIDByAddress is a free data retrieval call binding the contract method 0xcac8b306.
//
// Solidity: function operatorIDByAddress(address ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) OperatorIDByAddress(arg0 common.Address) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.OperatorIDByAddress(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// OperatorIDByAddress is a free data retrieval call binding the contract method 0xcac8b306.
//
// Solidity: function operatorIDByAddress(address ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) OperatorIDByAddress(arg0 common.Address) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.OperatorIDByAddress(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// OperatorIdForExcessDeposit is a free data retrieval call binding the contract method 0x9e03fb7e.
//
// Solidity: function operatorIdForExcessDeposit() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) OperatorIdForExcessDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "operatorIdForExcessDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorIdForExcessDeposit is a free data retrieval call binding the contract method 0x9e03fb7e.
//
// Solidity: function operatorIdForExcessDeposit() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) OperatorIdForExcessDeposit() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.OperatorIdForExcessDeposit(&_PermissionedNodeRegistry.CallOpts)
}

// OperatorIdForExcessDeposit is a free data retrieval call binding the contract method 0x9e03fb7e.
//
// Solidity: function operatorIdForExcessDeposit() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) OperatorIdForExcessDeposit() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.OperatorIdForExcessDeposit(&_PermissionedNodeRegistry.CallOpts)
}

// OperatorStructById is a free data retrieval call binding the contract method 0xc8a00e7a.
//
// Solidity: function operatorStructById(uint256 ) view returns(bool active, bool optedForSocializingPool, string operatorName, address operatorRewardAddress, address operatorAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) OperatorStructById(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "operatorStructById", arg0)

	outstruct := new(struct {
		Active                  bool
		OptedForSocializingPool bool
		OperatorName            string
		OperatorRewardAddress   common.Address
		OperatorAddress         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.OptedForSocializingPool = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.OperatorName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.OperatorRewardAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.OperatorAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// OperatorStructById is a free data retrieval call binding the contract method 0xc8a00e7a.
//
// Solidity: function operatorStructById(uint256 ) view returns(bool active, bool optedForSocializingPool, string operatorName, address operatorRewardAddress, address operatorAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) OperatorStructById(arg0 *big.Int) (struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}, error) {
	return _PermissionedNodeRegistry.Contract.OperatorStructById(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// OperatorStructById is a free data retrieval call binding the contract method 0xc8a00e7a.
//
// Solidity: function operatorStructById(uint256 ) view returns(bool active, bool optedForSocializingPool, string operatorName, address operatorRewardAddress, address operatorAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) OperatorStructById(arg0 *big.Int) (struct {
	Active                  bool
	OptedForSocializingPool bool
	OperatorName            string
	OperatorRewardAddress   common.Address
	OperatorAddress         common.Address
}, error) {
	return _PermissionedNodeRegistry.Contract.OperatorStructById(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) Paused() (bool, error) {
	return _PermissionedNodeRegistry.Contract.Paused(&_PermissionedNodeRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) Paused() (bool, error) {
	return _PermissionedNodeRegistry.Contract.Paused(&_PermissionedNodeRegistry.CallOpts)
}

// PermissionList is a free data retrieval call binding the contract method 0xbb0a5c76.
//
// Solidity: function permissionList(address ) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) PermissionList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "permissionList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PermissionList is a free data retrieval call binding the contract method 0xbb0a5c76.
//
// Solidity: function permissionList(address ) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) PermissionList(arg0 common.Address) (bool, error) {
	return _PermissionedNodeRegistry.Contract.PermissionList(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// PermissionList is a free data retrieval call binding the contract method 0xbb0a5c76.
//
// Solidity: function permissionList(address ) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) PermissionList(arg0 common.Address) (bool, error) {
	return _PermissionedNodeRegistry.Contract.PermissionList(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// SocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0x84522a6d.
//
// Solidity: function socializingPoolStateChangeBlock(uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) SocializingPoolStateChangeBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "socializingPoolStateChangeBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0x84522a6d.
//
// Solidity: function socializingPoolStateChangeBlock(uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) SocializingPoolStateChangeBlock(arg0 *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.SocializingPoolStateChangeBlock(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// SocializingPoolStateChangeBlock is a free data retrieval call binding the contract method 0x84522a6d.
//
// Solidity: function socializingPoolStateChangeBlock(uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) SocializingPoolStateChangeBlock(arg0 *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.SocializingPoolStateChangeBlock(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) StaderConfig() (common.Address, error) {
	return _PermissionedNodeRegistry.Contract.StaderConfig(&_PermissionedNodeRegistry.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) StaderConfig() (common.Address, error) {
	return _PermissionedNodeRegistry.Contract.StaderConfig(&_PermissionedNodeRegistry.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionedNodeRegistry.Contract.SupportsInterface(&_PermissionedNodeRegistry.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionedNodeRegistry.Contract.SupportsInterface(&_PermissionedNodeRegistry.CallOpts, interfaceId)
}

// TotalActiveOperatorCount is a free data retrieval call binding the contract method 0x063c6963.
//
// Solidity: function totalActiveOperatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) TotalActiveOperatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "totalActiveOperatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveOperatorCount is a free data retrieval call binding the contract method 0x063c6963.
//
// Solidity: function totalActiveOperatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) TotalActiveOperatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.TotalActiveOperatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// TotalActiveOperatorCount is a free data retrieval call binding the contract method 0x063c6963.
//
// Solidity: function totalActiveOperatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) TotalActiveOperatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.TotalActiveOperatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// TotalActiveValidatorCount is a free data retrieval call binding the contract method 0x84b0fa4c.
//
// Solidity: function totalActiveValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) TotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "totalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveValidatorCount is a free data retrieval call binding the contract method 0x84b0fa4c.
//
// Solidity: function totalActiveValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) TotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.TotalActiveValidatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// TotalActiveValidatorCount is a free data retrieval call binding the contract method 0x84b0fa4c.
//
// Solidity: function totalActiveValidatorCount() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) TotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.TotalActiveValidatorCount(&_PermissionedNodeRegistry.CallOpts)
}

// ValidatorIdByPubkey is a free data retrieval call binding the contract method 0x5c2c30a5.
//
// Solidity: function validatorIdByPubkey(bytes ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) ValidatorIdByPubkey(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "validatorIdByPubkey", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorIdByPubkey is a free data retrieval call binding the contract method 0x5c2c30a5.
//
// Solidity: function validatorIdByPubkey(bytes ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) ValidatorIdByPubkey(arg0 []byte) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.ValidatorIdByPubkey(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// ValidatorIdByPubkey is a free data retrieval call binding the contract method 0x5c2c30a5.
//
// Solidity: function validatorIdByPubkey(bytes ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) ValidatorIdByPubkey(arg0 []byte) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.ValidatorIdByPubkey(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// ValidatorIdsByOperatorId is a free data retrieval call binding the contract method 0xd5e1e5ce.
//
// Solidity: function validatorIdsByOperatorId(uint256 , uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) ValidatorIdsByOperatorId(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "validatorIdsByOperatorId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorIdsByOperatorId is a free data retrieval call binding the contract method 0xd5e1e5ce.
//
// Solidity: function validatorIdsByOperatorId(uint256 , uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) ValidatorIdsByOperatorId(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.ValidatorIdsByOperatorId(&_PermissionedNodeRegistry.CallOpts, arg0, arg1)
}

// ValidatorIdsByOperatorId is a free data retrieval call binding the contract method 0xd5e1e5ce.
//
// Solidity: function validatorIdsByOperatorId(uint256 , uint256 ) view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) ValidatorIdsByOperatorId(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.ValidatorIdsByOperatorId(&_PermissionedNodeRegistry.CallOpts, arg0, arg1)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 depositBlock, uint256 withdrawnBlock)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) ValidatorRegistry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "validatorRegistry", arg0)

	outstruct := new(struct {
		Status               uint8
		Pubkey               []byte
		PreDepositSignature  []byte
		DepositSignature     []byte
		WithdrawVaultAddress common.Address
		OperatorId           *big.Int
		DepositBlock         *big.Int
		WithdrawnBlock       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Pubkey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.PreDepositSignature = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.DepositSignature = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.WithdrawVaultAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.OperatorId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DepositBlock = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.WithdrawnBlock = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 depositBlock, uint256 withdrawnBlock)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) ValidatorRegistry(arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}, error) {
	return _PermissionedNodeRegistry.Contract.ValidatorRegistry(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0x5a1239c1.
//
// Solidity: function validatorRegistry(uint256 ) view returns(uint8 status, bytes pubkey, bytes preDepositSignature, bytes depositSignature, address withdrawVaultAddress, uint256 operatorId, uint256 depositBlock, uint256 withdrawnBlock)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) ValidatorRegistry(arg0 *big.Int) (struct {
	Status               uint8
	Pubkey               []byte
	PreDepositSignature  []byte
	DepositSignature     []byte
	WithdrawVaultAddress common.Address
	OperatorId           *big.Int
	DepositBlock         *big.Int
	WithdrawnBlock       *big.Int
}, error) {
	return _PermissionedNodeRegistry.Contract.ValidatorRegistry(&_PermissionedNodeRegistry.CallOpts, arg0)
}

// VerifiedKeyBatchSize is a free data retrieval call binding the contract method 0xab3e71eb.
//
// Solidity: function verifiedKeyBatchSize() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCaller) VerifiedKeyBatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedNodeRegistry.contract.Call(opts, &out, "verifiedKeyBatchSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifiedKeyBatchSize is a free data retrieval call binding the contract method 0xab3e71eb.
//
// Solidity: function verifiedKeyBatchSize() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) VerifiedKeyBatchSize() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.VerifiedKeyBatchSize(&_PermissionedNodeRegistry.CallOpts)
}

// VerifiedKeyBatchSize is a free data retrieval call binding the contract method 0xab3e71eb.
//
// Solidity: function verifiedKeyBatchSize() view returns(uint256)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryCallerSession) VerifiedKeyBatchSize() (*big.Int, error) {
	return _PermissionedNodeRegistry.Contract.VerifiedKeyBatchSize(&_PermissionedNodeRegistry.CallOpts)
}

// ActivateNodeOperator is a paid mutator transaction binding the contract method 0x91dcd6b2.
//
// Solidity: function activateNodeOperator(uint256 _operatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) ActivateNodeOperator(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "activateNodeOperator", _operatorId)
}

// ActivateNodeOperator is a paid mutator transaction binding the contract method 0x91dcd6b2.
//
// Solidity: function activateNodeOperator(uint256 _operatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) ActivateNodeOperator(_operatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.ActivateNodeOperator(&_PermissionedNodeRegistry.TransactOpts, _operatorId)
}

// ActivateNodeOperator is a paid mutator transaction binding the contract method 0x91dcd6b2.
//
// Solidity: function activateNodeOperator(uint256 _operatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) ActivateNodeOperator(_operatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.ActivateNodeOperator(&_PermissionedNodeRegistry.TransactOpts, _operatorId)
}

// AddValidatorKeys is a paid mutator transaction binding the contract method 0xdeacde2b.
//
// Solidity: function addValidatorKeys(bytes[] _pubkey, bytes[] _preDepositSignature, bytes[] _depositSignature) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) AddValidatorKeys(opts *bind.TransactOpts, _pubkey [][]byte, _preDepositSignature [][]byte, _depositSignature [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "addValidatorKeys", _pubkey, _preDepositSignature, _depositSignature)
}

// AddValidatorKeys is a paid mutator transaction binding the contract method 0xdeacde2b.
//
// Solidity: function addValidatorKeys(bytes[] _pubkey, bytes[] _preDepositSignature, bytes[] _depositSignature) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) AddValidatorKeys(_pubkey [][]byte, _preDepositSignature [][]byte, _depositSignature [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.AddValidatorKeys(&_PermissionedNodeRegistry.TransactOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// AddValidatorKeys is a paid mutator transaction binding the contract method 0xdeacde2b.
//
// Solidity: function addValidatorKeys(bytes[] _pubkey, bytes[] _preDepositSignature, bytes[] _depositSignature) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) AddValidatorKeys(_pubkey [][]byte, _preDepositSignature [][]byte, _depositSignature [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.AddValidatorKeys(&_PermissionedNodeRegistry.TransactOpts, _pubkey, _preDepositSignature, _depositSignature)
}

// AllocateValidatorsAndUpdateOperatorId is a paid mutator transaction binding the contract method 0x73e6e569.
//
// Solidity: function allocateValidatorsAndUpdateOperatorId(uint256 _numValidators) returns(uint256[] selectedOperatorCapacity)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) AllocateValidatorsAndUpdateOperatorId(opts *bind.TransactOpts, _numValidators *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "allocateValidatorsAndUpdateOperatorId", _numValidators)
}

// AllocateValidatorsAndUpdateOperatorId is a paid mutator transaction binding the contract method 0x73e6e569.
//
// Solidity: function allocateValidatorsAndUpdateOperatorId(uint256 _numValidators) returns(uint256[] selectedOperatorCapacity)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) AllocateValidatorsAndUpdateOperatorId(_numValidators *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.AllocateValidatorsAndUpdateOperatorId(&_PermissionedNodeRegistry.TransactOpts, _numValidators)
}

// AllocateValidatorsAndUpdateOperatorId is a paid mutator transaction binding the contract method 0x73e6e569.
//
// Solidity: function allocateValidatorsAndUpdateOperatorId(uint256 _numValidators) returns(uint256[] selectedOperatorCapacity)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) AllocateValidatorsAndUpdateOperatorId(_numValidators *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.AllocateValidatorsAndUpdateOperatorId(&_PermissionedNodeRegistry.TransactOpts, _numValidators)
}

// DeactivateNodeOperator is a paid mutator transaction binding the contract method 0x75a080d5.
//
// Solidity: function deactivateNodeOperator(uint256 _operatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) DeactivateNodeOperator(opts *bind.TransactOpts, _operatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "deactivateNodeOperator", _operatorId)
}

// DeactivateNodeOperator is a paid mutator transaction binding the contract method 0x75a080d5.
//
// Solidity: function deactivateNodeOperator(uint256 _operatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) DeactivateNodeOperator(_operatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.DeactivateNodeOperator(&_PermissionedNodeRegistry.TransactOpts, _operatorId)
}

// DeactivateNodeOperator is a paid mutator transaction binding the contract method 0x75a080d5.
//
// Solidity: function deactivateNodeOperator(uint256 _operatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) DeactivateNodeOperator(_operatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.DeactivateNodeOperator(&_PermissionedNodeRegistry.TransactOpts, _operatorId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.GrantRole(&_PermissionedNodeRegistry.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.GrantRole(&_PermissionedNodeRegistry.TransactOpts, role, account)
}

// IncreaseTotalActiveValidatorCount is a paid mutator transaction binding the contract method 0x59c3c9b7.
//
// Solidity: function increaseTotalActiveValidatorCount(uint256 _count) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) IncreaseTotalActiveValidatorCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "increaseTotalActiveValidatorCount", _count)
}

// IncreaseTotalActiveValidatorCount is a paid mutator transaction binding the contract method 0x59c3c9b7.
//
// Solidity: function increaseTotalActiveValidatorCount(uint256 _count) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) IncreaseTotalActiveValidatorCount(_count *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.IncreaseTotalActiveValidatorCount(&_PermissionedNodeRegistry.TransactOpts, _count)
}

// IncreaseTotalActiveValidatorCount is a paid mutator transaction binding the contract method 0x59c3c9b7.
//
// Solidity: function increaseTotalActiveValidatorCount(uint256 _count) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) IncreaseTotalActiveValidatorCount(_count *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.IncreaseTotalActiveValidatorCount(&_PermissionedNodeRegistry.TransactOpts, _count)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.Initialize(&_PermissionedNodeRegistry.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.Initialize(&_PermissionedNodeRegistry.TransactOpts, _admin, _staderConfig)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) MarkValidatorReadyToDeposit(opts *bind.TransactOpts, _readyToDepositPubkey [][]byte, _frontRunPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "markValidatorReadyToDeposit", _readyToDepositPubkey, _frontRunPubkey, _invalidSignaturePubkey)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) MarkValidatorReadyToDeposit(_readyToDepositPubkey [][]byte, _frontRunPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.MarkValidatorReadyToDeposit(&_PermissionedNodeRegistry.TransactOpts, _readyToDepositPubkey, _frontRunPubkey, _invalidSignaturePubkey)
}

// MarkValidatorReadyToDeposit is a paid mutator transaction binding the contract method 0x13797bff.
//
// Solidity: function markValidatorReadyToDeposit(bytes[] _readyToDepositPubkey, bytes[] _frontRunPubkey, bytes[] _invalidSignaturePubkey) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) MarkValidatorReadyToDeposit(_readyToDepositPubkey [][]byte, _frontRunPubkey [][]byte, _invalidSignaturePubkey [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.MarkValidatorReadyToDeposit(&_PermissionedNodeRegistry.TransactOpts, _readyToDepositPubkey, _frontRunPubkey, _invalidSignaturePubkey)
}

// MarkValidatorStatusAsPreDeposit is a paid mutator transaction binding the contract method 0x3f06739f.
//
// Solidity: function markValidatorStatusAsPreDeposit(bytes _pubkey) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) MarkValidatorStatusAsPreDeposit(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "markValidatorStatusAsPreDeposit", _pubkey)
}

// MarkValidatorStatusAsPreDeposit is a paid mutator transaction binding the contract method 0x3f06739f.
//
// Solidity: function markValidatorStatusAsPreDeposit(bytes _pubkey) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) MarkValidatorStatusAsPreDeposit(_pubkey []byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.MarkValidatorStatusAsPreDeposit(&_PermissionedNodeRegistry.TransactOpts, _pubkey)
}

// MarkValidatorStatusAsPreDeposit is a paid mutator transaction binding the contract method 0x3f06739f.
//
// Solidity: function markValidatorStatusAsPreDeposit(bytes _pubkey) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) MarkValidatorStatusAsPreDeposit(_pubkey []byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.MarkValidatorStatusAsPreDeposit(&_PermissionedNodeRegistry.TransactOpts, _pubkey)
}

// OnboardNodeOperator is a paid mutator transaction binding the contract method 0x21eb5a37.
//
// Solidity: function onboardNodeOperator(string _operatorName, address _operatorRewardAddress) returns(address feeRecipientAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) OnboardNodeOperator(opts *bind.TransactOpts, _operatorName string, _operatorRewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "onboardNodeOperator", _operatorName, _operatorRewardAddress)
}

// OnboardNodeOperator is a paid mutator transaction binding the contract method 0x21eb5a37.
//
// Solidity: function onboardNodeOperator(string _operatorName, address _operatorRewardAddress) returns(address feeRecipientAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) OnboardNodeOperator(_operatorName string, _operatorRewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.OnboardNodeOperator(&_PermissionedNodeRegistry.TransactOpts, _operatorName, _operatorRewardAddress)
}

// OnboardNodeOperator is a paid mutator transaction binding the contract method 0x21eb5a37.
//
// Solidity: function onboardNodeOperator(string _operatorName, address _operatorRewardAddress) returns(address feeRecipientAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) OnboardNodeOperator(_operatorName string, _operatorRewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.OnboardNodeOperator(&_PermissionedNodeRegistry.TransactOpts, _operatorName, _operatorRewardAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) Pause() (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.Pause(&_PermissionedNodeRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.Pause(&_PermissionedNodeRegistry.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.RenounceRole(&_PermissionedNodeRegistry.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.RenounceRole(&_PermissionedNodeRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.RevokeRole(&_PermissionedNodeRegistry.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.RevokeRole(&_PermissionedNodeRegistry.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) Unpause() (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.Unpause(&_PermissionedNodeRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.Unpause(&_PermissionedNodeRegistry.TransactOpts)
}

// UpdateDepositStatusAndBlock is a paid mutator transaction binding the contract method 0x186d9541.
//
// Solidity: function updateDepositStatusAndBlock(uint256 _validatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateDepositStatusAndBlock(opts *bind.TransactOpts, _validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateDepositStatusAndBlock", _validatorId)
}

// UpdateDepositStatusAndBlock is a paid mutator transaction binding the contract method 0x186d9541.
//
// Solidity: function updateDepositStatusAndBlock(uint256 _validatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateDepositStatusAndBlock(_validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateDepositStatusAndBlock(&_PermissionedNodeRegistry.TransactOpts, _validatorId)
}

// UpdateDepositStatusAndBlock is a paid mutator transaction binding the contract method 0x186d9541.
//
// Solidity: function updateDepositStatusAndBlock(uint256 _validatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateDepositStatusAndBlock(_validatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateDepositStatusAndBlock(&_PermissionedNodeRegistry.TransactOpts, _validatorId)
}

// UpdateInputKeyCountLimit is a paid mutator transaction binding the contract method 0x2517cfbf.
//
// Solidity: function updateInputKeyCountLimit(uint16 _inputKeyCountLimit) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateInputKeyCountLimit(opts *bind.TransactOpts, _inputKeyCountLimit uint16) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateInputKeyCountLimit", _inputKeyCountLimit)
}

// UpdateInputKeyCountLimit is a paid mutator transaction binding the contract method 0x2517cfbf.
//
// Solidity: function updateInputKeyCountLimit(uint16 _inputKeyCountLimit) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateInputKeyCountLimit(_inputKeyCountLimit uint16) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateInputKeyCountLimit(&_PermissionedNodeRegistry.TransactOpts, _inputKeyCountLimit)
}

// UpdateInputKeyCountLimit is a paid mutator transaction binding the contract method 0x2517cfbf.
//
// Solidity: function updateInputKeyCountLimit(uint16 _inputKeyCountLimit) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateInputKeyCountLimit(_inputKeyCountLimit uint16) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateInputKeyCountLimit(&_PermissionedNodeRegistry.TransactOpts, _inputKeyCountLimit)
}

// UpdateMaxNonTerminalKeyPerOperator is a paid mutator transaction binding the contract method 0x60c3cf3f.
//
// Solidity: function updateMaxNonTerminalKeyPerOperator(uint64 _maxNonTerminalKeyPerOperator) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateMaxNonTerminalKeyPerOperator(opts *bind.TransactOpts, _maxNonTerminalKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateMaxNonTerminalKeyPerOperator", _maxNonTerminalKeyPerOperator)
}

// UpdateMaxNonTerminalKeyPerOperator is a paid mutator transaction binding the contract method 0x60c3cf3f.
//
// Solidity: function updateMaxNonTerminalKeyPerOperator(uint64 _maxNonTerminalKeyPerOperator) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateMaxNonTerminalKeyPerOperator(_maxNonTerminalKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateMaxNonTerminalKeyPerOperator(&_PermissionedNodeRegistry.TransactOpts, _maxNonTerminalKeyPerOperator)
}

// UpdateMaxNonTerminalKeyPerOperator is a paid mutator transaction binding the contract method 0x60c3cf3f.
//
// Solidity: function updateMaxNonTerminalKeyPerOperator(uint64 _maxNonTerminalKeyPerOperator) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateMaxNonTerminalKeyPerOperator(_maxNonTerminalKeyPerOperator uint64) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateMaxNonTerminalKeyPerOperator(&_PermissionedNodeRegistry.TransactOpts, _maxNonTerminalKeyPerOperator)
}

// UpdateMaxOperatorId is a paid mutator transaction binding the contract method 0x2360bf21.
//
// Solidity: function updateMaxOperatorId(uint256 _maxOperatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateMaxOperatorId(opts *bind.TransactOpts, _maxOperatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateMaxOperatorId", _maxOperatorId)
}

// UpdateMaxOperatorId is a paid mutator transaction binding the contract method 0x2360bf21.
//
// Solidity: function updateMaxOperatorId(uint256 _maxOperatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateMaxOperatorId(_maxOperatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateMaxOperatorId(&_PermissionedNodeRegistry.TransactOpts, _maxOperatorId)
}

// UpdateMaxOperatorId is a paid mutator transaction binding the contract method 0x2360bf21.
//
// Solidity: function updateMaxOperatorId(uint256 _maxOperatorId) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateMaxOperatorId(_maxOperatorId *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateMaxOperatorId(&_PermissionedNodeRegistry.TransactOpts, _maxOperatorId)
}

// UpdateOperatorDetails is a paid mutator transaction binding the contract method 0x58a994ea.
//
// Solidity: function updateOperatorDetails(string _operatorName, address _rewardAddress) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateOperatorDetails(opts *bind.TransactOpts, _operatorName string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateOperatorDetails", _operatorName, _rewardAddress)
}

// UpdateOperatorDetails is a paid mutator transaction binding the contract method 0x58a994ea.
//
// Solidity: function updateOperatorDetails(string _operatorName, address _rewardAddress) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateOperatorDetails(_operatorName string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateOperatorDetails(&_PermissionedNodeRegistry.TransactOpts, _operatorName, _rewardAddress)
}

// UpdateOperatorDetails is a paid mutator transaction binding the contract method 0x58a994ea.
//
// Solidity: function updateOperatorDetails(string _operatorName, address _rewardAddress) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateOperatorDetails(_operatorName string, _rewardAddress common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateOperatorDetails(&_PermissionedNodeRegistry.TransactOpts, _operatorName, _rewardAddress)
}

// UpdateQueuedValidatorIndex is a paid mutator transaction binding the contract method 0xd1ac892c.
//
// Solidity: function updateQueuedValidatorIndex(uint256 _operatorId, uint256 _nextQueuedValidatorIndex) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateQueuedValidatorIndex(opts *bind.TransactOpts, _operatorId *big.Int, _nextQueuedValidatorIndex *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateQueuedValidatorIndex", _operatorId, _nextQueuedValidatorIndex)
}

// UpdateQueuedValidatorIndex is a paid mutator transaction binding the contract method 0xd1ac892c.
//
// Solidity: function updateQueuedValidatorIndex(uint256 _operatorId, uint256 _nextQueuedValidatorIndex) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateQueuedValidatorIndex(_operatorId *big.Int, _nextQueuedValidatorIndex *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateQueuedValidatorIndex(&_PermissionedNodeRegistry.TransactOpts, _operatorId, _nextQueuedValidatorIndex)
}

// UpdateQueuedValidatorIndex is a paid mutator transaction binding the contract method 0xd1ac892c.
//
// Solidity: function updateQueuedValidatorIndex(uint256 _operatorId, uint256 _nextQueuedValidatorIndex) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateQueuedValidatorIndex(_operatorId *big.Int, _nextQueuedValidatorIndex *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateQueuedValidatorIndex(&_PermissionedNodeRegistry.TransactOpts, _operatorId, _nextQueuedValidatorIndex)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateStaderConfig(&_PermissionedNodeRegistry.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateStaderConfig(&_PermissionedNodeRegistry.TransactOpts, _staderConfig)
}

// UpdateVerifiedKeysBatchSize is a paid mutator transaction binding the contract method 0xaf533aa8.
//
// Solidity: function updateVerifiedKeysBatchSize(uint256 _verifiedKeysBatchSize) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) UpdateVerifiedKeysBatchSize(opts *bind.TransactOpts, _verifiedKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "updateVerifiedKeysBatchSize", _verifiedKeysBatchSize)
}

// UpdateVerifiedKeysBatchSize is a paid mutator transaction binding the contract method 0xaf533aa8.
//
// Solidity: function updateVerifiedKeysBatchSize(uint256 _verifiedKeysBatchSize) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) UpdateVerifiedKeysBatchSize(_verifiedKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateVerifiedKeysBatchSize(&_PermissionedNodeRegistry.TransactOpts, _verifiedKeysBatchSize)
}

// UpdateVerifiedKeysBatchSize is a paid mutator transaction binding the contract method 0xaf533aa8.
//
// Solidity: function updateVerifiedKeysBatchSize(uint256 _verifiedKeysBatchSize) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) UpdateVerifiedKeysBatchSize(_verifiedKeysBatchSize *big.Int) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.UpdateVerifiedKeysBatchSize(&_PermissionedNodeRegistry.TransactOpts, _verifiedKeysBatchSize)
}

// WhitelistPermissionedNOs is a paid mutator transaction binding the contract method 0xf6dd6098.
//
// Solidity: function whitelistPermissionedNOs(address[] _permissionedNOs) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) WhitelistPermissionedNOs(opts *bind.TransactOpts, _permissionedNOs []common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "whitelistPermissionedNOs", _permissionedNOs)
}

// WhitelistPermissionedNOs is a paid mutator transaction binding the contract method 0xf6dd6098.
//
// Solidity: function whitelistPermissionedNOs(address[] _permissionedNOs) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) WhitelistPermissionedNOs(_permissionedNOs []common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.WhitelistPermissionedNOs(&_PermissionedNodeRegistry.TransactOpts, _permissionedNOs)
}

// WhitelistPermissionedNOs is a paid mutator transaction binding the contract method 0xf6dd6098.
//
// Solidity: function whitelistPermissionedNOs(address[] _permissionedNOs) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) WhitelistPermissionedNOs(_permissionedNOs []common.Address) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.WhitelistPermissionedNOs(&_PermissionedNodeRegistry.TransactOpts, _permissionedNOs)
}

// WithdrawnValidators is a paid mutator transaction binding the contract method 0x264f27f3.
//
// Solidity: function withdrawnValidators(bytes[] _pubkeys) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactor) WithdrawnValidators(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.contract.Transact(opts, "withdrawnValidators", _pubkeys)
}

// WithdrawnValidators is a paid mutator transaction binding the contract method 0x264f27f3.
//
// Solidity: function withdrawnValidators(bytes[] _pubkeys) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistrySession) WithdrawnValidators(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.WithdrawnValidators(&_PermissionedNodeRegistry.TransactOpts, _pubkeys)
}

// WithdrawnValidators is a paid mutator transaction binding the contract method 0x264f27f3.
//
// Solidity: function withdrawnValidators(bytes[] _pubkeys) returns()
func (_PermissionedNodeRegistry *PermissionedNodeRegistryTransactorSession) WithdrawnValidators(_pubkeys [][]byte) (*types.Transaction, error) {
	return _PermissionedNodeRegistry.Contract.WithdrawnValidators(&_PermissionedNodeRegistry.TransactOpts, _pubkeys)
}

// PermissionedNodeRegistryAddedValidatorKeyIterator is returned from FilterAddedValidatorKey and is used to iterate over the raw logs and unpacked data for AddedValidatorKey events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryAddedValidatorKeyIterator struct {
	Event *PermissionedNodeRegistryAddedValidatorKey // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryAddedValidatorKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryAddedValidatorKey)
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
		it.Event = new(PermissionedNodeRegistryAddedValidatorKey)
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
func (it *PermissionedNodeRegistryAddedValidatorKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryAddedValidatorKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryAddedValidatorKey represents a AddedValidatorKey event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryAddedValidatorKey struct {
	NodeOperator common.Address
	Pubkey       []byte
	ValidatorId  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddedValidatorKey is a free log retrieval operation binding the contract event 0xab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf.
//
// Solidity: event AddedValidatorKey(address indexed nodeOperator, bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterAddedValidatorKey(opts *bind.FilterOpts, nodeOperator []common.Address) (*PermissionedNodeRegistryAddedValidatorKeyIterator, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "AddedValidatorKey", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryAddedValidatorKeyIterator{contract: _PermissionedNodeRegistry.contract, event: "AddedValidatorKey", logs: logs, sub: sub}, nil
}

// WatchAddedValidatorKey is a free log subscription operation binding the contract event 0xab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf.
//
// Solidity: event AddedValidatorKey(address indexed nodeOperator, bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchAddedValidatorKey(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryAddedValidatorKey, nodeOperator []common.Address) (event.Subscription, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "AddedValidatorKey", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryAddedValidatorKey)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "AddedValidatorKey", log); err != nil {
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

// ParseAddedValidatorKey is a log parse operation binding the contract event 0xab5128638b64e6216e80dfafa70d3cb6d54913a536dc41e76eb4a04cfbe979cf.
//
// Solidity: event AddedValidatorKey(address indexed nodeOperator, bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseAddedValidatorKey(log types.Log) (*PermissionedNodeRegistryAddedValidatorKey, error) {
	event := new(PermissionedNodeRegistryAddedValidatorKey)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "AddedValidatorKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator is returned from FilterDecreasedTotalActiveValidatorCount and is used to iterate over the raw logs and unpacked data for DecreasedTotalActiveValidatorCount events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator struct {
	Event *PermissionedNodeRegistryDecreasedTotalActiveValidatorCount // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryDecreasedTotalActiveValidatorCount)
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
		it.Event = new(PermissionedNodeRegistryDecreasedTotalActiveValidatorCount)
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
func (it *PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryDecreasedTotalActiveValidatorCount represents a DecreasedTotalActiveValidatorCount event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryDecreasedTotalActiveValidatorCount struct {
	TotalActiveValidatorCount *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterDecreasedTotalActiveValidatorCount is a free log retrieval operation binding the contract event 0x5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece3.
//
// Solidity: event DecreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterDecreasedTotalActiveValidatorCount(opts *bind.FilterOpts) (*PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "DecreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryDecreasedTotalActiveValidatorCountIterator{contract: _PermissionedNodeRegistry.contract, event: "DecreasedTotalActiveValidatorCount", logs: logs, sub: sub}, nil
}

// WatchDecreasedTotalActiveValidatorCount is a free log subscription operation binding the contract event 0x5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece3.
//
// Solidity: event DecreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchDecreasedTotalActiveValidatorCount(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryDecreasedTotalActiveValidatorCount) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "DecreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryDecreasedTotalActiveValidatorCount)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "DecreasedTotalActiveValidatorCount", log); err != nil {
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

// ParseDecreasedTotalActiveValidatorCount is a log parse operation binding the contract event 0x5040a06a11b7d9b75fc56fbbd207905dbaa4ac86c0dc9cc7fff40cd1d92aece3.
//
// Solidity: event DecreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseDecreasedTotalActiveValidatorCount(log types.Log) (*PermissionedNodeRegistryDecreasedTotalActiveValidatorCount, error) {
	event := new(PermissionedNodeRegistryDecreasedTotalActiveValidatorCount)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "DecreasedTotalActiveValidatorCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator is returned from FilterIncreasedTotalActiveValidatorCount and is used to iterate over the raw logs and unpacked data for IncreasedTotalActiveValidatorCount events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator struct {
	Event *PermissionedNodeRegistryIncreasedTotalActiveValidatorCount // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryIncreasedTotalActiveValidatorCount)
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
		it.Event = new(PermissionedNodeRegistryIncreasedTotalActiveValidatorCount)
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
func (it *PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryIncreasedTotalActiveValidatorCount represents a IncreasedTotalActiveValidatorCount event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryIncreasedTotalActiveValidatorCount struct {
	TotalActiveValidatorCount *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterIncreasedTotalActiveValidatorCount is a free log retrieval operation binding the contract event 0x5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f2.
//
// Solidity: event IncreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterIncreasedTotalActiveValidatorCount(opts *bind.FilterOpts) (*PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "IncreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryIncreasedTotalActiveValidatorCountIterator{contract: _PermissionedNodeRegistry.contract, event: "IncreasedTotalActiveValidatorCount", logs: logs, sub: sub}, nil
}

// WatchIncreasedTotalActiveValidatorCount is a free log subscription operation binding the contract event 0x5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f2.
//
// Solidity: event IncreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchIncreasedTotalActiveValidatorCount(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryIncreasedTotalActiveValidatorCount) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "IncreasedTotalActiveValidatorCount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryIncreasedTotalActiveValidatorCount)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "IncreasedTotalActiveValidatorCount", log); err != nil {
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

// ParseIncreasedTotalActiveValidatorCount is a log parse operation binding the contract event 0x5818a627697795ff3c3403f320c7549835866cfb64a0b06a6f7f077bc478e9f2.
//
// Solidity: event IncreasedTotalActiveValidatorCount(uint256 totalActiveValidatorCount)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseIncreasedTotalActiveValidatorCount(log types.Log) (*PermissionedNodeRegistryIncreasedTotalActiveValidatorCount, error) {
	event := new(PermissionedNodeRegistryIncreasedTotalActiveValidatorCount)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "IncreasedTotalActiveValidatorCount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryInitializedIterator struct {
	Event *PermissionedNodeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryInitialized)
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
		it.Event = new(PermissionedNodeRegistryInitialized)
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
func (it *PermissionedNodeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryInitialized represents a Initialized event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermissionedNodeRegistryInitializedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryInitializedIterator{contract: _PermissionedNodeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryInitialized)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseInitialized(log types.Log) (*PermissionedNodeRegistryInitialized, error) {
	event := new(PermissionedNodeRegistryInitialized)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator is returned from FilterMarkedValidatorStatusAsPreDeposit and is used to iterate over the raw logs and unpacked data for MarkedValidatorStatusAsPreDeposit events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator struct {
	Event *PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit)
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
		it.Event = new(PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit)
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
func (it *PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit represents a MarkedValidatorStatusAsPreDeposit event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit struct {
	Pubkey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarkedValidatorStatusAsPreDeposit is a free log retrieval operation binding the contract event 0x4e171cd135eaf1ffe293a6b83fe6d28dadc58b9ecf77eae6f85136ab76349001.
//
// Solidity: event MarkedValidatorStatusAsPreDeposit(bytes pubkey)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterMarkedValidatorStatusAsPreDeposit(opts *bind.FilterOpts) (*PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "MarkedValidatorStatusAsPreDeposit")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryMarkedValidatorStatusAsPreDepositIterator{contract: _PermissionedNodeRegistry.contract, event: "MarkedValidatorStatusAsPreDeposit", logs: logs, sub: sub}, nil
}

// WatchMarkedValidatorStatusAsPreDeposit is a free log subscription operation binding the contract event 0x4e171cd135eaf1ffe293a6b83fe6d28dadc58b9ecf77eae6f85136ab76349001.
//
// Solidity: event MarkedValidatorStatusAsPreDeposit(bytes pubkey)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchMarkedValidatorStatusAsPreDeposit(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "MarkedValidatorStatusAsPreDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "MarkedValidatorStatusAsPreDeposit", log); err != nil {
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

// ParseMarkedValidatorStatusAsPreDeposit is a log parse operation binding the contract event 0x4e171cd135eaf1ffe293a6b83fe6d28dadc58b9ecf77eae6f85136ab76349001.
//
// Solidity: event MarkedValidatorStatusAsPreDeposit(bytes pubkey)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseMarkedValidatorStatusAsPreDeposit(log types.Log) (*PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit, error) {
	event := new(PermissionedNodeRegistryMarkedValidatorStatusAsPreDeposit)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "MarkedValidatorStatusAsPreDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator is returned from FilterMaxOperatorIdLimitChanged and is used to iterate over the raw logs and unpacked data for MaxOperatorIdLimitChanged events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator struct {
	Event *PermissionedNodeRegistryMaxOperatorIdLimitChanged // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryMaxOperatorIdLimitChanged)
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
		it.Event = new(PermissionedNodeRegistryMaxOperatorIdLimitChanged)
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
func (it *PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryMaxOperatorIdLimitChanged represents a MaxOperatorIdLimitChanged event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryMaxOperatorIdLimitChanged struct {
	MaxOperatorId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxOperatorIdLimitChanged is a free log retrieval operation binding the contract event 0x5098143eb105a8137c163f0207dbaab72f98e85a7affd34bcbc75341fd47522c.
//
// Solidity: event MaxOperatorIdLimitChanged(uint256 maxOperatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterMaxOperatorIdLimitChanged(opts *bind.FilterOpts) (*PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "MaxOperatorIdLimitChanged")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryMaxOperatorIdLimitChangedIterator{contract: _PermissionedNodeRegistry.contract, event: "MaxOperatorIdLimitChanged", logs: logs, sub: sub}, nil
}

// WatchMaxOperatorIdLimitChanged is a free log subscription operation binding the contract event 0x5098143eb105a8137c163f0207dbaab72f98e85a7affd34bcbc75341fd47522c.
//
// Solidity: event MaxOperatorIdLimitChanged(uint256 maxOperatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchMaxOperatorIdLimitChanged(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryMaxOperatorIdLimitChanged) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "MaxOperatorIdLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryMaxOperatorIdLimitChanged)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "MaxOperatorIdLimitChanged", log); err != nil {
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

// ParseMaxOperatorIdLimitChanged is a log parse operation binding the contract event 0x5098143eb105a8137c163f0207dbaab72f98e85a7affd34bcbc75341fd47522c.
//
// Solidity: event MaxOperatorIdLimitChanged(uint256 maxOperatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseMaxOperatorIdLimitChanged(log types.Log) (*PermissionedNodeRegistryMaxOperatorIdLimitChanged, error) {
	event := new(PermissionedNodeRegistryMaxOperatorIdLimitChanged)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "MaxOperatorIdLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryOnboardedOperatorIterator is returned from FilterOnboardedOperator and is used to iterate over the raw logs and unpacked data for OnboardedOperator events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOnboardedOperatorIterator struct {
	Event *PermissionedNodeRegistryOnboardedOperator // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryOnboardedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryOnboardedOperator)
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
		it.Event = new(PermissionedNodeRegistryOnboardedOperator)
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
func (it *PermissionedNodeRegistryOnboardedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryOnboardedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryOnboardedOperator represents a OnboardedOperator event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOnboardedOperator struct {
	NodeOperator      common.Address
	NodeRewardAddress common.Address
	OperatorId        *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOnboardedOperator is a free log retrieval operation binding the contract event 0x4d1311c747b9dccb91ce1767d13459f59d76abb25b79a9b032e5c277927fc736.
//
// Solidity: event OnboardedOperator(address indexed nodeOperator, address nodeRewardAddress, uint256 operatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterOnboardedOperator(opts *bind.FilterOpts, nodeOperator []common.Address) (*PermissionedNodeRegistryOnboardedOperatorIterator, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "OnboardedOperator", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryOnboardedOperatorIterator{contract: _PermissionedNodeRegistry.contract, event: "OnboardedOperator", logs: logs, sub: sub}, nil
}

// WatchOnboardedOperator is a free log subscription operation binding the contract event 0x4d1311c747b9dccb91ce1767d13459f59d76abb25b79a9b032e5c277927fc736.
//
// Solidity: event OnboardedOperator(address indexed nodeOperator, address nodeRewardAddress, uint256 operatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchOnboardedOperator(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryOnboardedOperator, nodeOperator []common.Address) (event.Subscription, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "OnboardedOperator", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryOnboardedOperator)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OnboardedOperator", log); err != nil {
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

// ParseOnboardedOperator is a log parse operation binding the contract event 0x4d1311c747b9dccb91ce1767d13459f59d76abb25b79a9b032e5c277927fc736.
//
// Solidity: event OnboardedOperator(address indexed nodeOperator, address nodeRewardAddress, uint256 operatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseOnboardedOperator(log types.Log) (*PermissionedNodeRegistryOnboardedOperator, error) {
	event := new(PermissionedNodeRegistryOnboardedOperator)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OnboardedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryOperatorActivatedIterator is returned from FilterOperatorActivated and is used to iterate over the raw logs and unpacked data for OperatorActivated events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOperatorActivatedIterator struct {
	Event *PermissionedNodeRegistryOperatorActivated // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryOperatorActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryOperatorActivated)
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
		it.Event = new(PermissionedNodeRegistryOperatorActivated)
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
func (it *PermissionedNodeRegistryOperatorActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryOperatorActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryOperatorActivated represents a OperatorActivated event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOperatorActivated struct {
	OperatorID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorActivated is a free log retrieval operation binding the contract event 0xcb6a50a41db3ab5eb9e2592806fadcf034ef05eb2cc4138c92a418d8e08d5b03.
//
// Solidity: event OperatorActivated(uint256 operatorID)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterOperatorActivated(opts *bind.FilterOpts) (*PermissionedNodeRegistryOperatorActivatedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "OperatorActivated")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryOperatorActivatedIterator{contract: _PermissionedNodeRegistry.contract, event: "OperatorActivated", logs: logs, sub: sub}, nil
}

// WatchOperatorActivated is a free log subscription operation binding the contract event 0xcb6a50a41db3ab5eb9e2592806fadcf034ef05eb2cc4138c92a418d8e08d5b03.
//
// Solidity: event OperatorActivated(uint256 operatorID)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchOperatorActivated(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryOperatorActivated) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "OperatorActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryOperatorActivated)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OperatorActivated", log); err != nil {
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

// ParseOperatorActivated is a log parse operation binding the contract event 0xcb6a50a41db3ab5eb9e2592806fadcf034ef05eb2cc4138c92a418d8e08d5b03.
//
// Solidity: event OperatorActivated(uint256 operatorID)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseOperatorActivated(log types.Log) (*PermissionedNodeRegistryOperatorActivated, error) {
	event := new(PermissionedNodeRegistryOperatorActivated)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OperatorActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryOperatorDeactivatedIterator is returned from FilterOperatorDeactivated and is used to iterate over the raw logs and unpacked data for OperatorDeactivated events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOperatorDeactivatedIterator struct {
	Event *PermissionedNodeRegistryOperatorDeactivated // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryOperatorDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryOperatorDeactivated)
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
		it.Event = new(PermissionedNodeRegistryOperatorDeactivated)
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
func (it *PermissionedNodeRegistryOperatorDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryOperatorDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryOperatorDeactivated represents a OperatorDeactivated event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOperatorDeactivated struct {
	OperatorID *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeactivated is a free log retrieval operation binding the contract event 0x6ff0b7b14b65a91172fa7e0a7b49d909add202dadcef1b08c80f3c136a914fa8.
//
// Solidity: event OperatorDeactivated(uint256 operatorID)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterOperatorDeactivated(opts *bind.FilterOpts) (*PermissionedNodeRegistryOperatorDeactivatedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "OperatorDeactivated")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryOperatorDeactivatedIterator{contract: _PermissionedNodeRegistry.contract, event: "OperatorDeactivated", logs: logs, sub: sub}, nil
}

// WatchOperatorDeactivated is a free log subscription operation binding the contract event 0x6ff0b7b14b65a91172fa7e0a7b49d909add202dadcef1b08c80f3c136a914fa8.
//
// Solidity: event OperatorDeactivated(uint256 operatorID)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchOperatorDeactivated(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryOperatorDeactivated) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "OperatorDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryOperatorDeactivated)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OperatorDeactivated", log); err != nil {
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

// ParseOperatorDeactivated is a log parse operation binding the contract event 0x6ff0b7b14b65a91172fa7e0a7b49d909add202dadcef1b08c80f3c136a914fa8.
//
// Solidity: event OperatorDeactivated(uint256 operatorID)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseOperatorDeactivated(log types.Log) (*PermissionedNodeRegistryOperatorDeactivated, error) {
	event := new(PermissionedNodeRegistryOperatorDeactivated)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OperatorDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryOperatorWhitelistedIterator is returned from FilterOperatorWhitelisted and is used to iterate over the raw logs and unpacked data for OperatorWhitelisted events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOperatorWhitelistedIterator struct {
	Event *PermissionedNodeRegistryOperatorWhitelisted // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryOperatorWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryOperatorWhitelisted)
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
		it.Event = new(PermissionedNodeRegistryOperatorWhitelisted)
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
func (it *PermissionedNodeRegistryOperatorWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryOperatorWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryOperatorWhitelisted represents a OperatorWhitelisted event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryOperatorWhitelisted struct {
	PermissionedNO common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelisted is a free log retrieval operation binding the contract event 0xa1a8e424f5c573f369672146cef50befb2515d5c6929b30703731f51af04daf4.
//
// Solidity: event OperatorWhitelisted(address permissionedNO)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterOperatorWhitelisted(opts *bind.FilterOpts) (*PermissionedNodeRegistryOperatorWhitelistedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "OperatorWhitelisted")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryOperatorWhitelistedIterator{contract: _PermissionedNodeRegistry.contract, event: "OperatorWhitelisted", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelisted is a free log subscription operation binding the contract event 0xa1a8e424f5c573f369672146cef50befb2515d5c6929b30703731f51af04daf4.
//
// Solidity: event OperatorWhitelisted(address permissionedNO)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchOperatorWhitelisted(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryOperatorWhitelisted) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "OperatorWhitelisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryOperatorWhitelisted)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OperatorWhitelisted", log); err != nil {
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

// ParseOperatorWhitelisted is a log parse operation binding the contract event 0xa1a8e424f5c573f369672146cef50befb2515d5c6929b30703731f51af04daf4.
//
// Solidity: event OperatorWhitelisted(address permissionedNO)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseOperatorWhitelisted(log types.Log) (*PermissionedNodeRegistryOperatorWhitelisted, error) {
	event := new(PermissionedNodeRegistryOperatorWhitelisted)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "OperatorWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryPausedIterator struct {
	Event *PermissionedNodeRegistryPaused // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryPaused)
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
		it.Event = new(PermissionedNodeRegistryPaused)
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
func (it *PermissionedNodeRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryPaused represents a Paused event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*PermissionedNodeRegistryPausedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryPausedIterator{contract: _PermissionedNodeRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryPaused)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParsePaused(log types.Log) (*PermissionedNodeRegistryPaused, error) {
	event := new(PermissionedNodeRegistryPaused)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryRoleAdminChangedIterator struct {
	Event *PermissionedNodeRegistryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryRoleAdminChanged)
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
		it.Event = new(PermissionedNodeRegistryRoleAdminChanged)
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
func (it *PermissionedNodeRegistryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryRoleAdminChanged represents a RoleAdminChanged event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PermissionedNodeRegistryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryRoleAdminChangedIterator{contract: _PermissionedNodeRegistry.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryRoleAdminChanged)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseRoleAdminChanged(log types.Log) (*PermissionedNodeRegistryRoleAdminChanged, error) {
	event := new(PermissionedNodeRegistryRoleAdminChanged)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryRoleGrantedIterator struct {
	Event *PermissionedNodeRegistryRoleGranted // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryRoleGranted)
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
		it.Event = new(PermissionedNodeRegistryRoleGranted)
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
func (it *PermissionedNodeRegistryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryRoleGranted represents a RoleGranted event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionedNodeRegistryRoleGrantedIterator, error) {

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

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryRoleGrantedIterator{contract: _PermissionedNodeRegistry.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryRoleGranted)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseRoleGranted(log types.Log) (*PermissionedNodeRegistryRoleGranted, error) {
	event := new(PermissionedNodeRegistryRoleGranted)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryRoleRevokedIterator struct {
	Event *PermissionedNodeRegistryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryRoleRevoked)
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
		it.Event = new(PermissionedNodeRegistryRoleRevoked)
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
func (it *PermissionedNodeRegistryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryRoleRevoked represents a RoleRevoked event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionedNodeRegistryRoleRevokedIterator, error) {

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

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryRoleRevokedIterator{contract: _PermissionedNodeRegistry.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryRoleRevoked)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseRoleRevoked(log types.Log) (*PermissionedNodeRegistryRoleRevoked, error) {
	event := new(PermissionedNodeRegistryRoleRevoked)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUnpausedIterator struct {
	Event *PermissionedNodeRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUnpaused)
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
		it.Event = new(PermissionedNodeRegistryUnpaused)
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
func (it *PermissionedNodeRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUnpaused represents a Unpaused event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PermissionedNodeRegistryUnpausedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUnpausedIterator{contract: _PermissionedNodeRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUnpaused)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUnpaused(log types.Log) (*PermissionedNodeRegistryUnpaused, error) {
	event := new(PermissionedNodeRegistryUnpaused)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator is returned from FilterUpdatedInputKeyCountLimit and is used to iterate over the raw logs and unpacked data for UpdatedInputKeyCountLimit events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator struct {
	Event *PermissionedNodeRegistryUpdatedInputKeyCountLimit // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedInputKeyCountLimit)
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
		it.Event = new(PermissionedNodeRegistryUpdatedInputKeyCountLimit)
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
func (it *PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedInputKeyCountLimit represents a UpdatedInputKeyCountLimit event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedInputKeyCountLimit struct {
	BatchKeyDepositLimit *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatedInputKeyCountLimit is a free log retrieval operation binding the contract event 0x5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c36.
//
// Solidity: event UpdatedInputKeyCountLimit(uint256 batchKeyDepositLimit)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedInputKeyCountLimit(opts *bind.FilterOpts) (*PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedInputKeyCountLimit")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedInputKeyCountLimitIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedInputKeyCountLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedInputKeyCountLimit is a free log subscription operation binding the contract event 0x5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c36.
//
// Solidity: event UpdatedInputKeyCountLimit(uint256 batchKeyDepositLimit)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedInputKeyCountLimit(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedInputKeyCountLimit) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedInputKeyCountLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedInputKeyCountLimit)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedInputKeyCountLimit", log); err != nil {
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

// ParseUpdatedInputKeyCountLimit is a log parse operation binding the contract event 0x5fd0fcd821abb4c92d47c4740e5f4a25ef35e99ee092d170faa0e5cb47013c36.
//
// Solidity: event UpdatedInputKeyCountLimit(uint256 batchKeyDepositLimit)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedInputKeyCountLimit(log types.Log) (*PermissionedNodeRegistryUpdatedInputKeyCountLimit, error) {
	event := new(PermissionedNodeRegistryUpdatedInputKeyCountLimit)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedInputKeyCountLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator is returned from FilterUpdatedMaxNonTerminalKeyPerOperator and is used to iterate over the raw logs and unpacked data for UpdatedMaxNonTerminalKeyPerOperator events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator struct {
	Event *PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
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
		it.Event = new(PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
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
func (it *PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator represents a UpdatedMaxNonTerminalKeyPerOperator event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator struct {
	MaxNonTerminalKeyPerOperator uint64
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxNonTerminalKeyPerOperator is a free log retrieval operation binding the contract event 0xacda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a2.
//
// Solidity: event UpdatedMaxNonTerminalKeyPerOperator(uint64 maxNonTerminalKeyPerOperator)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedMaxNonTerminalKeyPerOperator(opts *bind.FilterOpts) (*PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedMaxNonTerminalKeyPerOperator")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperatorIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedMaxNonTerminalKeyPerOperator", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxNonTerminalKeyPerOperator is a free log subscription operation binding the contract event 0xacda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a2.
//
// Solidity: event UpdatedMaxNonTerminalKeyPerOperator(uint64 maxNonTerminalKeyPerOperator)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedMaxNonTerminalKeyPerOperator(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedMaxNonTerminalKeyPerOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedMaxNonTerminalKeyPerOperator", log); err != nil {
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

// ParseUpdatedMaxNonTerminalKeyPerOperator is a log parse operation binding the contract event 0xacda2fe79efeffc359206ddeeb45f26ba1596223e01e1585458603af76e880a2.
//
// Solidity: event UpdatedMaxNonTerminalKeyPerOperator(uint64 maxNonTerminalKeyPerOperator)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedMaxNonTerminalKeyPerOperator(log types.Log) (*PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator, error) {
	event := new(PermissionedNodeRegistryUpdatedMaxNonTerminalKeyPerOperator)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedMaxNonTerminalKeyPerOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedOperatorDetailsIterator is returned from FilterUpdatedOperatorDetails and is used to iterate over the raw logs and unpacked data for UpdatedOperatorDetails events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedOperatorDetailsIterator struct {
	Event *PermissionedNodeRegistryUpdatedOperatorDetails // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedOperatorDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedOperatorDetails)
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
		it.Event = new(PermissionedNodeRegistryUpdatedOperatorDetails)
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
func (it *PermissionedNodeRegistryUpdatedOperatorDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedOperatorDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedOperatorDetails represents a UpdatedOperatorDetails event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedOperatorDetails struct {
	NodeOperator  common.Address
	OperatorName  string
	RewardAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOperatorDetails is a free log retrieval operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed nodeOperator, string operatorName, address rewardAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedOperatorDetails(opts *bind.FilterOpts, nodeOperator []common.Address) (*PermissionedNodeRegistryUpdatedOperatorDetailsIterator, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedOperatorDetails", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedOperatorDetailsIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorDetails is a free log subscription operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed nodeOperator, string operatorName, address rewardAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedOperatorDetails(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedOperatorDetails, nodeOperator []common.Address) (event.Subscription, error) {

	var nodeOperatorRule []interface{}
	for _, nodeOperatorItem := range nodeOperator {
		nodeOperatorRule = append(nodeOperatorRule, nodeOperatorItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedOperatorDetails", nodeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedOperatorDetails)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedOperatorDetails", log); err != nil {
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

// ParseUpdatedOperatorDetails is a log parse operation binding the contract event 0xadc8722095edf061d7fdcb583105c05bf9eb15488503b621c39e254d87269777.
//
// Solidity: event UpdatedOperatorDetails(address indexed nodeOperator, string operatorName, address rewardAddress)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedOperatorDetails(log types.Log) (*PermissionedNodeRegistryUpdatedOperatorDetails, error) {
	event := new(PermissionedNodeRegistryUpdatedOperatorDetails)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator is returned from FilterUpdatedQueuedValidatorIndex and is used to iterate over the raw logs and unpacked data for UpdatedQueuedValidatorIndex events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator struct {
	Event *PermissionedNodeRegistryUpdatedQueuedValidatorIndex // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedQueuedValidatorIndex)
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
		it.Event = new(PermissionedNodeRegistryUpdatedQueuedValidatorIndex)
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
func (it *PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedQueuedValidatorIndex represents a UpdatedQueuedValidatorIndex event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedQueuedValidatorIndex struct {
	OperatorId               *big.Int
	NextQueuedValidatorIndex *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedQueuedValidatorIndex is a free log retrieval operation binding the contract event 0xe70c90ab9752cfe420f7c454a7182fdc437d1b4d580c0daf54a6edb6b0307a90.
//
// Solidity: event UpdatedQueuedValidatorIndex(uint256 indexed operatorId, uint256 nextQueuedValidatorIndex)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedQueuedValidatorIndex(opts *bind.FilterOpts, operatorId []*big.Int) (*PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedQueuedValidatorIndex", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedQueuedValidatorIndexIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedQueuedValidatorIndex", logs: logs, sub: sub}, nil
}

// WatchUpdatedQueuedValidatorIndex is a free log subscription operation binding the contract event 0xe70c90ab9752cfe420f7c454a7182fdc437d1b4d580c0daf54a6edb6b0307a90.
//
// Solidity: event UpdatedQueuedValidatorIndex(uint256 indexed operatorId, uint256 nextQueuedValidatorIndex)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedQueuedValidatorIndex(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedQueuedValidatorIndex, operatorId []*big.Int) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedQueuedValidatorIndex", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedQueuedValidatorIndex)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedQueuedValidatorIndex", log); err != nil {
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

// ParseUpdatedQueuedValidatorIndex is a log parse operation binding the contract event 0xe70c90ab9752cfe420f7c454a7182fdc437d1b4d580c0daf54a6edb6b0307a90.
//
// Solidity: event UpdatedQueuedValidatorIndex(uint256 indexed operatorId, uint256 nextQueuedValidatorIndex)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedQueuedValidatorIndex(log types.Log) (*PermissionedNodeRegistryUpdatedQueuedValidatorIndex, error) {
	event := new(PermissionedNodeRegistryUpdatedQueuedValidatorIndex)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedQueuedValidatorIndex", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedStaderConfigIterator struct {
	Event *PermissionedNodeRegistryUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedStaderConfig)
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
		it.Event = new(PermissionedNodeRegistryUpdatedStaderConfig)
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
func (it *PermissionedNodeRegistryUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PermissionedNodeRegistryUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedStaderConfigIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedStaderConfig)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedStaderConfig(log types.Log) (*PermissionedNodeRegistryUpdatedStaderConfig, error) {
	event := new(PermissionedNodeRegistryUpdatedStaderConfig)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator is returned from FilterUpdatedValidatorDepositBlock and is used to iterate over the raw logs and unpacked data for UpdatedValidatorDepositBlock events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator struct {
	Event *PermissionedNodeRegistryUpdatedValidatorDepositBlock // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedValidatorDepositBlock)
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
		it.Event = new(PermissionedNodeRegistryUpdatedValidatorDepositBlock)
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
func (it *PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedValidatorDepositBlock represents a UpdatedValidatorDepositBlock event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedValidatorDepositBlock struct {
	ValidatorId  *big.Int
	DepositBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedValidatorDepositBlock is a free log retrieval operation binding the contract event 0xce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c01.
//
// Solidity: event UpdatedValidatorDepositBlock(uint256 validatorId, uint256 depositBlock)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedValidatorDepositBlock(opts *bind.FilterOpts) (*PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedValidatorDepositBlock")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedValidatorDepositBlockIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedValidatorDepositBlock", logs: logs, sub: sub}, nil
}

// WatchUpdatedValidatorDepositBlock is a free log subscription operation binding the contract event 0xce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c01.
//
// Solidity: event UpdatedValidatorDepositBlock(uint256 validatorId, uint256 depositBlock)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedValidatorDepositBlock(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedValidatorDepositBlock) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedValidatorDepositBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedValidatorDepositBlock)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedValidatorDepositBlock", log); err != nil {
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

// ParseUpdatedValidatorDepositBlock is a log parse operation binding the contract event 0xce479ab1b7a806fa3704c907b8fae15a191ad8da9a1671659e4f411f516c4c01.
//
// Solidity: event UpdatedValidatorDepositBlock(uint256 validatorId, uint256 depositBlock)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedValidatorDepositBlock(log types.Log) (*PermissionedNodeRegistryUpdatedValidatorDepositBlock, error) {
	event := new(PermissionedNodeRegistryUpdatedValidatorDepositBlock)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedValidatorDepositBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator is returned from FilterUpdatedVerifiedKeyBatchSize and is used to iterate over the raw logs and unpacked data for UpdatedVerifiedKeyBatchSize events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator struct {
	Event *PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize)
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
		it.Event = new(PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize)
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
func (it *PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize represents a UpdatedVerifiedKeyBatchSize event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize struct {
	VerifiedKeysBatchSize *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdatedVerifiedKeyBatchSize is a free log retrieval operation binding the contract event 0x5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b4.
//
// Solidity: event UpdatedVerifiedKeyBatchSize(uint256 verifiedKeysBatchSize)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedVerifiedKeyBatchSize(opts *bind.FilterOpts) (*PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedVerifiedKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedVerifiedKeyBatchSizeIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedVerifiedKeyBatchSize", logs: logs, sub: sub}, nil
}

// WatchUpdatedVerifiedKeyBatchSize is a free log subscription operation binding the contract event 0x5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b4.
//
// Solidity: event UpdatedVerifiedKeyBatchSize(uint256 verifiedKeysBatchSize)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedVerifiedKeyBatchSize(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedVerifiedKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedVerifiedKeyBatchSize", log); err != nil {
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

// ParseUpdatedVerifiedKeyBatchSize is a log parse operation binding the contract event 0x5d19c92c6893231b764f3320c712a4d056ff157295c8b620d893dbbed1a869b4.
//
// Solidity: event UpdatedVerifiedKeyBatchSize(uint256 verifiedKeysBatchSize)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedVerifiedKeyBatchSize(log types.Log) (*PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize, error) {
	event := new(PermissionedNodeRegistryUpdatedVerifiedKeyBatchSize)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedVerifiedKeyBatchSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator is returned from FilterUpdatedWithdrawnKeyBatchSize and is used to iterate over the raw logs and unpacked data for UpdatedWithdrawnKeyBatchSize events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator struct {
	Event *PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize)
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
		it.Event = new(PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize)
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
func (it *PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize represents a UpdatedWithdrawnKeyBatchSize event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize struct {
	WithdrawnKeysBatchSize *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedWithdrawnKeyBatchSize is a free log retrieval operation binding the contract event 0x5aa519ec64f29fb81c513568f7c6839ee0265b5799bb434dfa467be612125950.
//
// Solidity: event UpdatedWithdrawnKeyBatchSize(uint256 withdrawnKeysBatchSize)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterUpdatedWithdrawnKeyBatchSize(opts *bind.FilterOpts) (*PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "UpdatedWithdrawnKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSizeIterator{contract: _PermissionedNodeRegistry.contract, event: "UpdatedWithdrawnKeyBatchSize", logs: logs, sub: sub}, nil
}

// WatchUpdatedWithdrawnKeyBatchSize is a free log subscription operation binding the contract event 0x5aa519ec64f29fb81c513568f7c6839ee0265b5799bb434dfa467be612125950.
//
// Solidity: event UpdatedWithdrawnKeyBatchSize(uint256 withdrawnKeysBatchSize)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchUpdatedWithdrawnKeyBatchSize(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "UpdatedWithdrawnKeyBatchSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedWithdrawnKeyBatchSize", log); err != nil {
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

// ParseUpdatedWithdrawnKeyBatchSize is a log parse operation binding the contract event 0x5aa519ec64f29fb81c513568f7c6839ee0265b5799bb434dfa467be612125950.
//
// Solidity: event UpdatedWithdrawnKeyBatchSize(uint256 withdrawnKeysBatchSize)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseUpdatedWithdrawnKeyBatchSize(log types.Log) (*PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize, error) {
	event := new(PermissionedNodeRegistryUpdatedWithdrawnKeyBatchSize)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "UpdatedWithdrawnKeyBatchSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator is returned from FilterValidatorMarkedAsFrontRunned and is used to iterate over the raw logs and unpacked data for ValidatorMarkedAsFrontRunned events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator struct {
	Event *PermissionedNodeRegistryValidatorMarkedAsFrontRunned // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryValidatorMarkedAsFrontRunned)
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
		it.Event = new(PermissionedNodeRegistryValidatorMarkedAsFrontRunned)
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
func (it *PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryValidatorMarkedAsFrontRunned represents a ValidatorMarkedAsFrontRunned event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryValidatorMarkedAsFrontRunned struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorMarkedAsFrontRunned is a free log retrieval operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterValidatorMarkedAsFrontRunned(opts *bind.FilterOpts) (*PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "ValidatorMarkedAsFrontRunned")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryValidatorMarkedAsFrontRunnedIterator{contract: _PermissionedNodeRegistry.contract, event: "ValidatorMarkedAsFrontRunned", logs: logs, sub: sub}, nil
}

// WatchValidatorMarkedAsFrontRunned is a free log subscription operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchValidatorMarkedAsFrontRunned(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryValidatorMarkedAsFrontRunned) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "ValidatorMarkedAsFrontRunned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryValidatorMarkedAsFrontRunned)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "ValidatorMarkedAsFrontRunned", log); err != nil {
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

// ParseValidatorMarkedAsFrontRunned is a log parse operation binding the contract event 0x4e93215f00bc729272f0ff71afd3d0f385208cbf6c999fe776ad07c623b83466.
//
// Solidity: event ValidatorMarkedAsFrontRunned(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseValidatorMarkedAsFrontRunned(log types.Log) (*PermissionedNodeRegistryValidatorMarkedAsFrontRunned, error) {
	event := new(PermissionedNodeRegistryValidatorMarkedAsFrontRunned)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "ValidatorMarkedAsFrontRunned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator is returned from FilterValidatorStatusMarkedAsInvalidSignature and is used to iterate over the raw logs and unpacked data for ValidatorStatusMarkedAsInvalidSignature events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator struct {
	Event *PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature)
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
		it.Event = new(PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature)
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
func (it *PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature represents a ValidatorStatusMarkedAsInvalidSignature event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusMarkedAsInvalidSignature is a free log retrieval operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterValidatorStatusMarkedAsInvalidSignature(opts *bind.FilterOpts) (*PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "ValidatorStatusMarkedAsInvalidSignature")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignatureIterator{contract: _PermissionedNodeRegistry.contract, event: "ValidatorStatusMarkedAsInvalidSignature", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusMarkedAsInvalidSignature is a free log subscription operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchValidatorStatusMarkedAsInvalidSignature(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "ValidatorStatusMarkedAsInvalidSignature")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "ValidatorStatusMarkedAsInvalidSignature", log); err != nil {
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

// ParseValidatorStatusMarkedAsInvalidSignature is a log parse operation binding the contract event 0x596ee835bed6cb827d21ba1785c468f0755ee40d33d87132df5d2ec90b645f9f.
//
// Solidity: event ValidatorStatusMarkedAsInvalidSignature(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseValidatorStatusMarkedAsInvalidSignature(log types.Log) (*PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature, error) {
	event := new(PermissionedNodeRegistryValidatorStatusMarkedAsInvalidSignature)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "ValidatorStatusMarkedAsInvalidSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedNodeRegistryValidatorWithdrawnIterator is returned from FilterValidatorWithdrawn and is used to iterate over the raw logs and unpacked data for ValidatorWithdrawn events raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryValidatorWithdrawnIterator struct {
	Event *PermissionedNodeRegistryValidatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *PermissionedNodeRegistryValidatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedNodeRegistryValidatorWithdrawn)
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
		it.Event = new(PermissionedNodeRegistryValidatorWithdrawn)
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
func (it *PermissionedNodeRegistryValidatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedNodeRegistryValidatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedNodeRegistryValidatorWithdrawn represents a ValidatorWithdrawn event raised by the PermissionedNodeRegistry contract.
type PermissionedNodeRegistryValidatorWithdrawn struct {
	Pubkey      []byte
	ValidatorId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorWithdrawn is a free log retrieval operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) FilterValidatorWithdrawn(opts *bind.FilterOpts) (*PermissionedNodeRegistryValidatorWithdrawnIterator, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.FilterLogs(opts, "ValidatorWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PermissionedNodeRegistryValidatorWithdrawnIterator{contract: _PermissionedNodeRegistry.contract, event: "ValidatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchValidatorWithdrawn is a free log subscription operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) WatchValidatorWithdrawn(opts *bind.WatchOpts, sink chan<- *PermissionedNodeRegistryValidatorWithdrawn) (event.Subscription, error) {

	logs, sub, err := _PermissionedNodeRegistry.contract.WatchLogs(opts, "ValidatorWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedNodeRegistryValidatorWithdrawn)
				if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
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

// ParseValidatorWithdrawn is a log parse operation binding the contract event 0x450186694fefe67df6156f60235e4073b623160f28a0b85908ebc864316abf79.
//
// Solidity: event ValidatorWithdrawn(bytes pubkey, uint256 validatorId)
func (_PermissionedNodeRegistry *PermissionedNodeRegistryFilterer) ParseValidatorWithdrawn(log types.Log) (*PermissionedNodeRegistryValidatorWithdrawn, error) {
	event := new(PermissionedNodeRegistryValidatorWithdrawn)
	if err := _PermissionedNodeRegistry.contract.UnpackLog(event, "ValidatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
