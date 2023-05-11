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

// PermissionedPoolMetaData contains all meta data concerning the PermissionedPool contract.
var PermissionedPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaderContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CouldNotDetermineExcessETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorFeeUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeeUnchanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedCollateralETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedInsuranceFund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferredETHToSSPMForDefectiveKeys\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorFee\",\"type\":\"uint256\"}],\"name\":\"UpdatedCommissionFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staderConfig\",\"type\":\"address\"}],\"name\":\"UpdatedStaderConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorDepositedOnBeaconChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorPreDepositedOnBeaconChain\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_COMMISSION_LIMIT_BIPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_withdrawCredential\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"computeDepositDataRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkey\",\"type\":\"bytes[]\"}],\"name\":\"fullDepositOnBeaconChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndex\",\"type\":\"uint256\"}],\"name\":\"getOperatorTotalNonTerminalKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocializingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQueuedValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operAddr\",\"type\":\"address\"}],\"name\":\"isExistingOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isExistingPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preDepositValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInsuranceFund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_operatorFee\",\"type\":\"uint256\"}],\"name\":\"setCommissionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staderConfig\",\"outputs\":[{\"internalType\":\"contractIStaderConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeUserETHToBeaconChain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defectiveKeyCount\",\"type\":\"uint256\"}],\"name\":\"transferETHOfDefectiveKeysToSSPM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferExcessETHToSSPM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staderConfig\",\"type\":\"address\"}],\"name\":\"updateStaderConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PermissionedPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PermissionedPoolMetaData.ABI instead.
var PermissionedPoolABI = PermissionedPoolMetaData.ABI

// PermissionedPool is an auto generated Go binding around an Ethereum contract.
type PermissionedPool struct {
	PermissionedPoolCaller     // Read-only binding to the contract
	PermissionedPoolTransactor // Write-only binding to the contract
	PermissionedPoolFilterer   // Log filterer for contract events
}

// PermissionedPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermissionedPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermissionedPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermissionedPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermissionedPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermissionedPoolSession struct {
	Contract     *PermissionedPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermissionedPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermissionedPoolCallerSession struct {
	Contract *PermissionedPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PermissionedPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermissionedPoolTransactorSession struct {
	Contract     *PermissionedPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PermissionedPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermissionedPoolRaw struct {
	Contract *PermissionedPool // Generic contract binding to access the raw methods on
}

// PermissionedPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermissionedPoolCallerRaw struct {
	Contract *PermissionedPoolCaller // Generic read-only contract binding to access the raw methods on
}

// PermissionedPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermissionedPoolTransactorRaw struct {
	Contract *PermissionedPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermissionedPool creates a new instance of PermissionedPool, bound to a specific deployed contract.
func NewPermissionedPool(address common.Address, backend bind.ContractBackend) (*PermissionedPool, error) {
	contract, err := bindPermissionedPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermissionedPool{PermissionedPoolCaller: PermissionedPoolCaller{contract: contract}, PermissionedPoolTransactor: PermissionedPoolTransactor{contract: contract}, PermissionedPoolFilterer: PermissionedPoolFilterer{contract: contract}}, nil
}

// NewPermissionedPoolCaller creates a new read-only instance of PermissionedPool, bound to a specific deployed contract.
func NewPermissionedPoolCaller(address common.Address, caller bind.ContractCaller) (*PermissionedPoolCaller, error) {
	contract, err := bindPermissionedPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolCaller{contract: contract}, nil
}

// NewPermissionedPoolTransactor creates a new write-only instance of PermissionedPool, bound to a specific deployed contract.
func NewPermissionedPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PermissionedPoolTransactor, error) {
	contract, err := bindPermissionedPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolTransactor{contract: contract}, nil
}

// NewPermissionedPoolFilterer creates a new log filterer instance of PermissionedPool, bound to a specific deployed contract.
func NewPermissionedPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PermissionedPoolFilterer, error) {
	contract, err := bindPermissionedPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolFilterer{contract: contract}, nil
}

// bindPermissionedPool binds a generic wrapper to an already deployed contract.
func bindPermissionedPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermissionedPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedPool *PermissionedPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedPool.Contract.PermissionedPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedPool *PermissionedPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedPool.Contract.PermissionedPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedPool *PermissionedPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedPool.Contract.PermissionedPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermissionedPool *PermissionedPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermissionedPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermissionedPool *PermissionedPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermissionedPool *PermissionedPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermissionedPool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionedPool *PermissionedPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionedPool *PermissionedPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionedPool.Contract.DEFAULTADMINROLE(&_PermissionedPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_PermissionedPool *PermissionedPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _PermissionedPool.Contract.DEFAULTADMINROLE(&_PermissionedPool.CallOpts)
}

// MAXCOMMISSIONLIMITBIPS is a free data retrieval call binding the contract method 0x9cd6dd56.
//
// Solidity: function MAX_COMMISSION_LIMIT_BIPS() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) MAXCOMMISSIONLIMITBIPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "MAX_COMMISSION_LIMIT_BIPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOMMISSIONLIMITBIPS is a free data retrieval call binding the contract method 0x9cd6dd56.
//
// Solidity: function MAX_COMMISSION_LIMIT_BIPS() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) MAXCOMMISSIONLIMITBIPS() (*big.Int, error) {
	return _PermissionedPool.Contract.MAXCOMMISSIONLIMITBIPS(&_PermissionedPool.CallOpts)
}

// MAXCOMMISSIONLIMITBIPS is a free data retrieval call binding the contract method 0x9cd6dd56.
//
// Solidity: function MAX_COMMISSION_LIMIT_BIPS() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) MAXCOMMISSIONLIMITBIPS() (*big.Int, error) {
	return _PermissionedPool.Contract.MAXCOMMISSIONLIMITBIPS(&_PermissionedPool.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionedPool *PermissionedPoolCaller) POOLID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "POOL_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionedPool *PermissionedPoolSession) POOLID() (uint8, error) {
	return _PermissionedPool.Contract.POOLID(&_PermissionedPool.CallOpts)
}

// POOLID is a free data retrieval call binding the contract method 0xe0d7d0e9.
//
// Solidity: function POOL_ID() view returns(uint8)
func (_PermissionedPool *PermissionedPoolCallerSession) POOLID() (uint8, error) {
	return _PermissionedPool.Contract.POOLID(&_PermissionedPool.CallOpts)
}

// ComputeDepositDataRoot is a free data retrieval call binding the contract method 0x5c164e53.
//
// Solidity: function computeDepositDataRoot(bytes _pubkey, bytes _signature, bytes _withdrawCredential, uint256 _depositAmount) pure returns(bytes32)
func (_PermissionedPool *PermissionedPoolCaller) ComputeDepositDataRoot(opts *bind.CallOpts, _pubkey []byte, _signature []byte, _withdrawCredential []byte, _depositAmount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "computeDepositDataRoot", _pubkey, _signature, _withdrawCredential, _depositAmount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeDepositDataRoot is a free data retrieval call binding the contract method 0x5c164e53.
//
// Solidity: function computeDepositDataRoot(bytes _pubkey, bytes _signature, bytes _withdrawCredential, uint256 _depositAmount) pure returns(bytes32)
func (_PermissionedPool *PermissionedPoolSession) ComputeDepositDataRoot(_pubkey []byte, _signature []byte, _withdrawCredential []byte, _depositAmount *big.Int) ([32]byte, error) {
	return _PermissionedPool.Contract.ComputeDepositDataRoot(&_PermissionedPool.CallOpts, _pubkey, _signature, _withdrawCredential, _depositAmount)
}

// ComputeDepositDataRoot is a free data retrieval call binding the contract method 0x5c164e53.
//
// Solidity: function computeDepositDataRoot(bytes _pubkey, bytes _signature, bytes _withdrawCredential, uint256 _depositAmount) pure returns(bytes32)
func (_PermissionedPool *PermissionedPoolCallerSession) ComputeDepositDataRoot(_pubkey []byte, _signature []byte, _withdrawCredential []byte, _depositAmount *big.Int) ([32]byte, error) {
	return _PermissionedPool.Contract.ComputeDepositDataRoot(&_PermissionedPool.CallOpts, _pubkey, _signature, _withdrawCredential, _depositAmount)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) GetCollateralETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getCollateralETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) GetCollateralETH() (*big.Int, error) {
	return _PermissionedPool.Contract.GetCollateralETH(&_PermissionedPool.CallOpts)
}

// GetCollateralETH is a free data retrieval call binding the contract method 0xb01db078.
//
// Solidity: function getCollateralETH() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) GetCollateralETH() (*big.Int, error) {
	return _PermissionedPool.Contract.GetCollateralETH(&_PermissionedPool.CallOpts)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0xb6fb3fac.
//
// Solidity: function getNodeRegistry() view returns(address)
func (_PermissionedPool *PermissionedPoolCaller) GetNodeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getNodeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeRegistry is a free data retrieval call binding the contract method 0xb6fb3fac.
//
// Solidity: function getNodeRegistry() view returns(address)
func (_PermissionedPool *PermissionedPoolSession) GetNodeRegistry() (common.Address, error) {
	return _PermissionedPool.Contract.GetNodeRegistry(&_PermissionedPool.CallOpts)
}

// GetNodeRegistry is a free data retrieval call binding the contract method 0xb6fb3fac.
//
// Solidity: function getNodeRegistry() view returns(address)
func (_PermissionedPool *PermissionedPoolCallerSession) GetNodeRegistry() (common.Address, error) {
	return _PermissionedPool.Contract.GetNodeRegistry(&_PermissionedPool.CallOpts)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) GetOperatorTotalNonTerminalKeys(opts *bind.CallOpts, _nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getOperatorTotalNonTerminalKeys", _nodeOperator, _startIndex, _endIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PermissionedPool.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionedPool.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetOperatorTotalNonTerminalKeys is a free data retrieval call binding the contract method 0x8a25bcec.
//
// Solidity: function getOperatorTotalNonTerminalKeys(address _nodeOperator, uint256 _startIndex, uint256 _endIndex) view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) GetOperatorTotalNonTerminalKeys(_nodeOperator common.Address, _startIndex *big.Int, _endIndex *big.Int) (*big.Int, error) {
	return _PermissionedPool.Contract.GetOperatorTotalNonTerminalKeys(&_PermissionedPool.CallOpts, _nodeOperator, _startIndex, _endIndex)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionedPool *PermissionedPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionedPool *PermissionedPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionedPool.Contract.GetRoleAdmin(&_PermissionedPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_PermissionedPool *PermissionedPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _PermissionedPool.Contract.GetRoleAdmin(&_PermissionedPool.CallOpts, role)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0xf74b4cd1.
//
// Solidity: function getSocializingPoolAddress() view returns(address)
func (_PermissionedPool *PermissionedPoolCaller) GetSocializingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getSocializingPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0xf74b4cd1.
//
// Solidity: function getSocializingPoolAddress() view returns(address)
func (_PermissionedPool *PermissionedPoolSession) GetSocializingPoolAddress() (common.Address, error) {
	return _PermissionedPool.Contract.GetSocializingPoolAddress(&_PermissionedPool.CallOpts)
}

// GetSocializingPoolAddress is a free data retrieval call binding the contract method 0xf74b4cd1.
//
// Solidity: function getSocializingPoolAddress() view returns(address)
func (_PermissionedPool *PermissionedPoolCallerSession) GetSocializingPoolAddress() (common.Address, error) {
	return _PermissionedPool.Contract.GetSocializingPoolAddress(&_PermissionedPool.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) GetTotalActiveValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getTotalActiveValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionedPool.Contract.GetTotalActiveValidatorCount(&_PermissionedPool.CallOpts)
}

// GetTotalActiveValidatorCount is a free data retrieval call binding the contract method 0x77c359e1.
//
// Solidity: function getTotalActiveValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) GetTotalActiveValidatorCount() (*big.Int, error) {
	return _PermissionedPool.Contract.GetTotalActiveValidatorCount(&_PermissionedPool.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) GetTotalQueuedValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "getTotalQueuedValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionedPool.Contract.GetTotalQueuedValidatorCount(&_PermissionedPool.CallOpts)
}

// GetTotalQueuedValidatorCount is a free data retrieval call binding the contract method 0x7bd977d9.
//
// Solidity: function getTotalQueuedValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) GetTotalQueuedValidatorCount() (*big.Int, error) {
	return _PermissionedPool.Contract.GetTotalQueuedValidatorCount(&_PermissionedPool.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionedPool *PermissionedPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionedPool *PermissionedPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionedPool.Contract.HasRole(&_PermissionedPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_PermissionedPool *PermissionedPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _PermissionedPool.Contract.HasRole(&_PermissionedPool.CallOpts, role, account)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionedPool *PermissionedPoolCaller) IsExistingOperator(opts *bind.CallOpts, _operAddr common.Address) (bool, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "isExistingOperator", _operAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionedPool *PermissionedPoolSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionedPool.Contract.IsExistingOperator(&_PermissionedPool.CallOpts, _operAddr)
}

// IsExistingOperator is a free data retrieval call binding the contract method 0xf9c4dda4.
//
// Solidity: function isExistingOperator(address _operAddr) view returns(bool)
func (_PermissionedPool *PermissionedPoolCallerSession) IsExistingOperator(_operAddr common.Address) (bool, error) {
	return _PermissionedPool.Contract.IsExistingOperator(&_PermissionedPool.CallOpts, _operAddr)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionedPool *PermissionedPoolCaller) IsExistingPubkey(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "isExistingPubkey", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionedPool *PermissionedPoolSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionedPool.Contract.IsExistingPubkey(&_PermissionedPool.CallOpts, _pubkey)
}

// IsExistingPubkey is a free data retrieval call binding the contract method 0x36514d9f.
//
// Solidity: function isExistingPubkey(bytes _pubkey) view returns(bool)
func (_PermissionedPool *PermissionedPoolCallerSession) IsExistingPubkey(_pubkey []byte) (bool, error) {
	return _PermissionedPool.Contract.IsExistingPubkey(&_PermissionedPool.CallOpts, _pubkey)
}

// OperatorFee is a free data retrieval call binding the contract method 0x89afc0f1.
//
// Solidity: function operatorFee() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) OperatorFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "operatorFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorFee is a free data retrieval call binding the contract method 0x89afc0f1.
//
// Solidity: function operatorFee() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) OperatorFee() (*big.Int, error) {
	return _PermissionedPool.Contract.OperatorFee(&_PermissionedPool.CallOpts)
}

// OperatorFee is a free data retrieval call binding the contract method 0x89afc0f1.
//
// Solidity: function operatorFee() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) OperatorFee() (*big.Int, error) {
	return _PermissionedPool.Contract.OperatorFee(&_PermissionedPool.CallOpts)
}

// PreDepositValidatorCount is a free data retrieval call binding the contract method 0x29976506.
//
// Solidity: function preDepositValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) PreDepositValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "preDepositValidatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreDepositValidatorCount is a free data retrieval call binding the contract method 0x29976506.
//
// Solidity: function preDepositValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) PreDepositValidatorCount() (*big.Int, error) {
	return _PermissionedPool.Contract.PreDepositValidatorCount(&_PermissionedPool.CallOpts)
}

// PreDepositValidatorCount is a free data retrieval call binding the contract method 0x29976506.
//
// Solidity: function preDepositValidatorCount() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) PreDepositValidatorCount() (*big.Int, error) {
	return _PermissionedPool.Contract.PreDepositValidatorCount(&_PermissionedPool.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PermissionedPool *PermissionedPoolSession) ProtocolFee() (*big.Int, error) {
	return _PermissionedPool.Contract.ProtocolFee(&_PermissionedPool.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_PermissionedPool *PermissionedPoolCallerSession) ProtocolFee() (*big.Int, error) {
	return _PermissionedPool.Contract.ProtocolFee(&_PermissionedPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionedPool *PermissionedPoolCaller) StaderConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "staderConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionedPool *PermissionedPoolSession) StaderConfig() (common.Address, error) {
	return _PermissionedPool.Contract.StaderConfig(&_PermissionedPool.CallOpts)
}

// StaderConfig is a free data retrieval call binding the contract method 0x490ffa35.
//
// Solidity: function staderConfig() view returns(address)
func (_PermissionedPool *PermissionedPoolCallerSession) StaderConfig() (common.Address, error) {
	return _PermissionedPool.Contract.StaderConfig(&_PermissionedPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionedPool *PermissionedPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PermissionedPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionedPool *PermissionedPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionedPool.Contract.SupportsInterface(&_PermissionedPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PermissionedPool *PermissionedPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PermissionedPool.Contract.SupportsInterface(&_PermissionedPool.CallOpts, interfaceId)
}

// FullDepositOnBeaconChain is a paid mutator transaction binding the contract method 0x4bd1e4c5.
//
// Solidity: function fullDepositOnBeaconChain(bytes[] _pubkey) returns()
func (_PermissionedPool *PermissionedPoolTransactor) FullDepositOnBeaconChain(opts *bind.TransactOpts, _pubkey [][]byte) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "fullDepositOnBeaconChain", _pubkey)
}

// FullDepositOnBeaconChain is a paid mutator transaction binding the contract method 0x4bd1e4c5.
//
// Solidity: function fullDepositOnBeaconChain(bytes[] _pubkey) returns()
func (_PermissionedPool *PermissionedPoolSession) FullDepositOnBeaconChain(_pubkey [][]byte) (*types.Transaction, error) {
	return _PermissionedPool.Contract.FullDepositOnBeaconChain(&_PermissionedPool.TransactOpts, _pubkey)
}

// FullDepositOnBeaconChain is a paid mutator transaction binding the contract method 0x4bd1e4c5.
//
// Solidity: function fullDepositOnBeaconChain(bytes[] _pubkey) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) FullDepositOnBeaconChain(_pubkey [][]byte) (*types.Transaction, error) {
	return _PermissionedPool.Contract.FullDepositOnBeaconChain(&_PermissionedPool.TransactOpts, _pubkey)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.GrantRole(&_PermissionedPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.GrantRole(&_PermissionedPool.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionedPool *PermissionedPoolTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "initialize", _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionedPool *PermissionedPoolSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.Initialize(&_PermissionedPool.TransactOpts, _admin, _staderConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _admin, address _staderConfig) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) Initialize(_admin common.Address, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.Initialize(&_PermissionedPool.TransactOpts, _admin, _staderConfig)
}

// ReceiveInsuranceFund is a paid mutator transaction binding the contract method 0x016557e7.
//
// Solidity: function receiveInsuranceFund() payable returns()
func (_PermissionedPool *PermissionedPoolTransactor) ReceiveInsuranceFund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "receiveInsuranceFund")
}

// ReceiveInsuranceFund is a paid mutator transaction binding the contract method 0x016557e7.
//
// Solidity: function receiveInsuranceFund() payable returns()
func (_PermissionedPool *PermissionedPoolSession) ReceiveInsuranceFund() (*types.Transaction, error) {
	return _PermissionedPool.Contract.ReceiveInsuranceFund(&_PermissionedPool.TransactOpts)
}

// ReceiveInsuranceFund is a paid mutator transaction binding the contract method 0x016557e7.
//
// Solidity: function receiveInsuranceFund() payable returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) ReceiveInsuranceFund() (*types.Transaction, error) {
	return _PermissionedPool.Contract.ReceiveInsuranceFund(&_PermissionedPool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.RenounceRole(&_PermissionedPool.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.RenounceRole(&_PermissionedPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.RevokeRole(&_PermissionedPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.RevokeRole(&_PermissionedPool.TransactOpts, role, account)
}

// SetCommissionFees is a paid mutator transaction binding the contract method 0x21066d18.
//
// Solidity: function setCommissionFees(uint256 _protocolFee, uint256 _operatorFee) returns()
func (_PermissionedPool *PermissionedPoolTransactor) SetCommissionFees(opts *bind.TransactOpts, _protocolFee *big.Int, _operatorFee *big.Int) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "setCommissionFees", _protocolFee, _operatorFee)
}

// SetCommissionFees is a paid mutator transaction binding the contract method 0x21066d18.
//
// Solidity: function setCommissionFees(uint256 _protocolFee, uint256 _operatorFee) returns()
func (_PermissionedPool *PermissionedPoolSession) SetCommissionFees(_protocolFee *big.Int, _operatorFee *big.Int) (*types.Transaction, error) {
	return _PermissionedPool.Contract.SetCommissionFees(&_PermissionedPool.TransactOpts, _protocolFee, _operatorFee)
}

// SetCommissionFees is a paid mutator transaction binding the contract method 0x21066d18.
//
// Solidity: function setCommissionFees(uint256 _protocolFee, uint256 _operatorFee) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) SetCommissionFees(_protocolFee *big.Int, _operatorFee *big.Int) (*types.Transaction, error) {
	return _PermissionedPool.Contract.SetCommissionFees(&_PermissionedPool.TransactOpts, _protocolFee, _operatorFee)
}

// StakeUserETHToBeaconChain is a paid mutator transaction binding the contract method 0x9b26728e.
//
// Solidity: function stakeUserETHToBeaconChain() payable returns()
func (_PermissionedPool *PermissionedPoolTransactor) StakeUserETHToBeaconChain(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "stakeUserETHToBeaconChain")
}

// StakeUserETHToBeaconChain is a paid mutator transaction binding the contract method 0x9b26728e.
//
// Solidity: function stakeUserETHToBeaconChain() payable returns()
func (_PermissionedPool *PermissionedPoolSession) StakeUserETHToBeaconChain() (*types.Transaction, error) {
	return _PermissionedPool.Contract.StakeUserETHToBeaconChain(&_PermissionedPool.TransactOpts)
}

// StakeUserETHToBeaconChain is a paid mutator transaction binding the contract method 0x9b26728e.
//
// Solidity: function stakeUserETHToBeaconChain() payable returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) StakeUserETHToBeaconChain() (*types.Transaction, error) {
	return _PermissionedPool.Contract.StakeUserETHToBeaconChain(&_PermissionedPool.TransactOpts)
}

// TransferETHOfDefectiveKeysToSSPM is a paid mutator transaction binding the contract method 0xd5c851a4.
//
// Solidity: function transferETHOfDefectiveKeysToSSPM(uint256 _defectiveKeyCount) returns()
func (_PermissionedPool *PermissionedPoolTransactor) TransferETHOfDefectiveKeysToSSPM(opts *bind.TransactOpts, _defectiveKeyCount *big.Int) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "transferETHOfDefectiveKeysToSSPM", _defectiveKeyCount)
}

// TransferETHOfDefectiveKeysToSSPM is a paid mutator transaction binding the contract method 0xd5c851a4.
//
// Solidity: function transferETHOfDefectiveKeysToSSPM(uint256 _defectiveKeyCount) returns()
func (_PermissionedPool *PermissionedPoolSession) TransferETHOfDefectiveKeysToSSPM(_defectiveKeyCount *big.Int) (*types.Transaction, error) {
	return _PermissionedPool.Contract.TransferETHOfDefectiveKeysToSSPM(&_PermissionedPool.TransactOpts, _defectiveKeyCount)
}

// TransferETHOfDefectiveKeysToSSPM is a paid mutator transaction binding the contract method 0xd5c851a4.
//
// Solidity: function transferETHOfDefectiveKeysToSSPM(uint256 _defectiveKeyCount) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) TransferETHOfDefectiveKeysToSSPM(_defectiveKeyCount *big.Int) (*types.Transaction, error) {
	return _PermissionedPool.Contract.TransferETHOfDefectiveKeysToSSPM(&_PermissionedPool.TransactOpts, _defectiveKeyCount)
}

// TransferExcessETHToSSPM is a paid mutator transaction binding the contract method 0x5e638738.
//
// Solidity: function transferExcessETHToSSPM() returns()
func (_PermissionedPool *PermissionedPoolTransactor) TransferExcessETHToSSPM(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "transferExcessETHToSSPM")
}

// TransferExcessETHToSSPM is a paid mutator transaction binding the contract method 0x5e638738.
//
// Solidity: function transferExcessETHToSSPM() returns()
func (_PermissionedPool *PermissionedPoolSession) TransferExcessETHToSSPM() (*types.Transaction, error) {
	return _PermissionedPool.Contract.TransferExcessETHToSSPM(&_PermissionedPool.TransactOpts)
}

// TransferExcessETHToSSPM is a paid mutator transaction binding the contract method 0x5e638738.
//
// Solidity: function transferExcessETHToSSPM() returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) TransferExcessETHToSSPM() (*types.Transaction, error) {
	return _PermissionedPool.Contract.TransferExcessETHToSSPM(&_PermissionedPool.TransactOpts)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionedPool *PermissionedPoolTransactor) UpdateStaderConfig(opts *bind.TransactOpts, _staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedPool.contract.Transact(opts, "updateStaderConfig", _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionedPool *PermissionedPoolSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.UpdateStaderConfig(&_PermissionedPool.TransactOpts, _staderConfig)
}

// UpdateStaderConfig is a paid mutator transaction binding the contract method 0x9ee804cb.
//
// Solidity: function updateStaderConfig(address _staderConfig) returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) UpdateStaderConfig(_staderConfig common.Address) (*types.Transaction, error) {
	return _PermissionedPool.Contract.UpdateStaderConfig(&_PermissionedPool.TransactOpts, _staderConfig)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PermissionedPool *PermissionedPoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PermissionedPool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PermissionedPool *PermissionedPoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PermissionedPool.Contract.Fallback(&_PermissionedPool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PermissionedPool.Contract.Fallback(&_PermissionedPool.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PermissionedPool *PermissionedPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermissionedPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PermissionedPool *PermissionedPoolSession) Receive() (*types.Transaction, error) {
	return _PermissionedPool.Contract.Receive(&_PermissionedPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PermissionedPool *PermissionedPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _PermissionedPool.Contract.Receive(&_PermissionedPool.TransactOpts)
}

// PermissionedPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PermissionedPool contract.
type PermissionedPoolInitializedIterator struct {
	Event *PermissionedPoolInitialized // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolInitialized)
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
		it.Event = new(PermissionedPoolInitialized)
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
func (it *PermissionedPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolInitialized represents a Initialized event raised by the PermissionedPool contract.
type PermissionedPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionedPool *PermissionedPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*PermissionedPoolInitializedIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolInitializedIterator{contract: _PermissionedPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PermissionedPool *PermissionedPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PermissionedPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolInitialized)
				if err := _PermissionedPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PermissionedPool *PermissionedPoolFilterer) ParseInitialized(log types.Log) (*PermissionedPoolInitialized, error) {
	event := new(PermissionedPoolInitialized)
	if err := _PermissionedPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolReceivedCollateralETHIterator is returned from FilterReceivedCollateralETH and is used to iterate over the raw logs and unpacked data for ReceivedCollateralETH events raised by the PermissionedPool contract.
type PermissionedPoolReceivedCollateralETHIterator struct {
	Event *PermissionedPoolReceivedCollateralETH // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolReceivedCollateralETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolReceivedCollateralETH)
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
		it.Event = new(PermissionedPoolReceivedCollateralETH)
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
func (it *PermissionedPoolReceivedCollateralETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolReceivedCollateralETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolReceivedCollateralETH represents a ReceivedCollateralETH event raised by the PermissionedPool contract.
type PermissionedPoolReceivedCollateralETH struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedCollateralETH is a free log retrieval operation binding the contract event 0x3726a70cbf9252aacb06774b2f9f7108d837fc60afd7143f86eec77d7e3da94b.
//
// Solidity: event ReceivedCollateralETH(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) FilterReceivedCollateralETH(opts *bind.FilterOpts) (*PermissionedPoolReceivedCollateralETHIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "ReceivedCollateralETH")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolReceivedCollateralETHIterator{contract: _PermissionedPool.contract, event: "ReceivedCollateralETH", logs: logs, sub: sub}, nil
}

// WatchReceivedCollateralETH is a free log subscription operation binding the contract event 0x3726a70cbf9252aacb06774b2f9f7108d837fc60afd7143f86eec77d7e3da94b.
//
// Solidity: event ReceivedCollateralETH(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) WatchReceivedCollateralETH(opts *bind.WatchOpts, sink chan<- *PermissionedPoolReceivedCollateralETH) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "ReceivedCollateralETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolReceivedCollateralETH)
				if err := _PermissionedPool.contract.UnpackLog(event, "ReceivedCollateralETH", log); err != nil {
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

// ParseReceivedCollateralETH is a log parse operation binding the contract event 0x3726a70cbf9252aacb06774b2f9f7108d837fc60afd7143f86eec77d7e3da94b.
//
// Solidity: event ReceivedCollateralETH(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) ParseReceivedCollateralETH(log types.Log) (*PermissionedPoolReceivedCollateralETH, error) {
	event := new(PermissionedPoolReceivedCollateralETH)
	if err := _PermissionedPool.contract.UnpackLog(event, "ReceivedCollateralETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolReceivedInsuranceFundIterator is returned from FilterReceivedInsuranceFund and is used to iterate over the raw logs and unpacked data for ReceivedInsuranceFund events raised by the PermissionedPool contract.
type PermissionedPoolReceivedInsuranceFundIterator struct {
	Event *PermissionedPoolReceivedInsuranceFund // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolReceivedInsuranceFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolReceivedInsuranceFund)
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
		it.Event = new(PermissionedPoolReceivedInsuranceFund)
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
func (it *PermissionedPoolReceivedInsuranceFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolReceivedInsuranceFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolReceivedInsuranceFund represents a ReceivedInsuranceFund event raised by the PermissionedPool contract.
type PermissionedPoolReceivedInsuranceFund struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedInsuranceFund is a free log retrieval operation binding the contract event 0xf43316e35306192a57665c134c0731f8e965fd6b625c620983cc68c2dd183726.
//
// Solidity: event ReceivedInsuranceFund(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) FilterReceivedInsuranceFund(opts *bind.FilterOpts) (*PermissionedPoolReceivedInsuranceFundIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "ReceivedInsuranceFund")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolReceivedInsuranceFundIterator{contract: _PermissionedPool.contract, event: "ReceivedInsuranceFund", logs: logs, sub: sub}, nil
}

// WatchReceivedInsuranceFund is a free log subscription operation binding the contract event 0xf43316e35306192a57665c134c0731f8e965fd6b625c620983cc68c2dd183726.
//
// Solidity: event ReceivedInsuranceFund(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) WatchReceivedInsuranceFund(opts *bind.WatchOpts, sink chan<- *PermissionedPoolReceivedInsuranceFund) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "ReceivedInsuranceFund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolReceivedInsuranceFund)
				if err := _PermissionedPool.contract.UnpackLog(event, "ReceivedInsuranceFund", log); err != nil {
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

// ParseReceivedInsuranceFund is a log parse operation binding the contract event 0xf43316e35306192a57665c134c0731f8e965fd6b625c620983cc68c2dd183726.
//
// Solidity: event ReceivedInsuranceFund(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) ParseReceivedInsuranceFund(log types.Log) (*PermissionedPoolReceivedInsuranceFund, error) {
	event := new(PermissionedPoolReceivedInsuranceFund)
	if err := _PermissionedPool.contract.UnpackLog(event, "ReceivedInsuranceFund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the PermissionedPool contract.
type PermissionedPoolRoleAdminChangedIterator struct {
	Event *PermissionedPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolRoleAdminChanged)
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
		it.Event = new(PermissionedPoolRoleAdminChanged)
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
func (it *PermissionedPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolRoleAdminChanged represents a RoleAdminChanged event raised by the PermissionedPool contract.
type PermissionedPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionedPool *PermissionedPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PermissionedPoolRoleAdminChangedIterator, error) {

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

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolRoleAdminChangedIterator{contract: _PermissionedPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_PermissionedPool *PermissionedPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PermissionedPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolRoleAdminChanged)
				if err := _PermissionedPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_PermissionedPool *PermissionedPoolFilterer) ParseRoleAdminChanged(log types.Log) (*PermissionedPoolRoleAdminChanged, error) {
	event := new(PermissionedPoolRoleAdminChanged)
	if err := _PermissionedPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the PermissionedPool contract.
type PermissionedPoolRoleGrantedIterator struct {
	Event *PermissionedPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolRoleGranted)
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
		it.Event = new(PermissionedPoolRoleGranted)
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
func (it *PermissionedPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolRoleGranted represents a RoleGranted event raised by the PermissionedPool contract.
type PermissionedPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedPool *PermissionedPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionedPoolRoleGrantedIterator, error) {

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

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolRoleGrantedIterator{contract: _PermissionedPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedPool *PermissionedPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PermissionedPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolRoleGranted)
				if err := _PermissionedPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_PermissionedPool *PermissionedPoolFilterer) ParseRoleGranted(log types.Log) (*PermissionedPoolRoleGranted, error) {
	event := new(PermissionedPoolRoleGranted)
	if err := _PermissionedPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the PermissionedPool contract.
type PermissionedPoolRoleRevokedIterator struct {
	Event *PermissionedPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolRoleRevoked)
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
		it.Event = new(PermissionedPoolRoleRevoked)
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
func (it *PermissionedPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolRoleRevoked represents a RoleRevoked event raised by the PermissionedPool contract.
type PermissionedPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedPool *PermissionedPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PermissionedPoolRoleRevokedIterator, error) {

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

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolRoleRevokedIterator{contract: _PermissionedPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_PermissionedPool *PermissionedPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PermissionedPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolRoleRevoked)
				if err := _PermissionedPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_PermissionedPool *PermissionedPoolFilterer) ParseRoleRevoked(log types.Log) (*PermissionedPoolRoleRevoked, error) {
	event := new(PermissionedPoolRoleRevoked)
	if err := _PermissionedPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator is returned from FilterTransferredETHToSSPMForDefectiveKeys and is used to iterate over the raw logs and unpacked data for TransferredETHToSSPMForDefectiveKeys events raised by the PermissionedPool contract.
type PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator struct {
	Event *PermissionedPoolTransferredETHToSSPMForDefectiveKeys // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolTransferredETHToSSPMForDefectiveKeys)
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
		it.Event = new(PermissionedPoolTransferredETHToSSPMForDefectiveKeys)
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
func (it *PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolTransferredETHToSSPMForDefectiveKeys represents a TransferredETHToSSPMForDefectiveKeys event raised by the PermissionedPool contract.
type PermissionedPoolTransferredETHToSSPMForDefectiveKeys struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferredETHToSSPMForDefectiveKeys is a free log retrieval operation binding the contract event 0x1149ac7366ffa1a2dee383a534b6742d013df82e7c15fd5720fd1ce82592a36f.
//
// Solidity: event TransferredETHToSSPMForDefectiveKeys(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) FilterTransferredETHToSSPMForDefectiveKeys(opts *bind.FilterOpts) (*PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "TransferredETHToSSPMForDefectiveKeys")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolTransferredETHToSSPMForDefectiveKeysIterator{contract: _PermissionedPool.contract, event: "TransferredETHToSSPMForDefectiveKeys", logs: logs, sub: sub}, nil
}

// WatchTransferredETHToSSPMForDefectiveKeys is a free log subscription operation binding the contract event 0x1149ac7366ffa1a2dee383a534b6742d013df82e7c15fd5720fd1ce82592a36f.
//
// Solidity: event TransferredETHToSSPMForDefectiveKeys(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) WatchTransferredETHToSSPMForDefectiveKeys(opts *bind.WatchOpts, sink chan<- *PermissionedPoolTransferredETHToSSPMForDefectiveKeys) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "TransferredETHToSSPMForDefectiveKeys")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolTransferredETHToSSPMForDefectiveKeys)
				if err := _PermissionedPool.contract.UnpackLog(event, "TransferredETHToSSPMForDefectiveKeys", log); err != nil {
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

// ParseTransferredETHToSSPMForDefectiveKeys is a log parse operation binding the contract event 0x1149ac7366ffa1a2dee383a534b6742d013df82e7c15fd5720fd1ce82592a36f.
//
// Solidity: event TransferredETHToSSPMForDefectiveKeys(uint256 amount)
func (_PermissionedPool *PermissionedPoolFilterer) ParseTransferredETHToSSPMForDefectiveKeys(log types.Log) (*PermissionedPoolTransferredETHToSSPMForDefectiveKeys, error) {
	event := new(PermissionedPoolTransferredETHToSSPMForDefectiveKeys)
	if err := _PermissionedPool.contract.UnpackLog(event, "TransferredETHToSSPMForDefectiveKeys", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolUpdatedCommissionFeesIterator is returned from FilterUpdatedCommissionFees and is used to iterate over the raw logs and unpacked data for UpdatedCommissionFees events raised by the PermissionedPool contract.
type PermissionedPoolUpdatedCommissionFeesIterator struct {
	Event *PermissionedPoolUpdatedCommissionFees // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolUpdatedCommissionFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolUpdatedCommissionFees)
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
		it.Event = new(PermissionedPoolUpdatedCommissionFees)
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
func (it *PermissionedPoolUpdatedCommissionFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolUpdatedCommissionFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolUpdatedCommissionFees represents a UpdatedCommissionFees event raised by the PermissionedPool contract.
type PermissionedPoolUpdatedCommissionFees struct {
	ProtocolFee *big.Int
	OperatorFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedCommissionFees is a free log retrieval operation binding the contract event 0xe8525a7862504bbe091b0440fe96979769664cae1625591ff0a9816512a5287b.
//
// Solidity: event UpdatedCommissionFees(uint256 protocolFee, uint256 operatorFee)
func (_PermissionedPool *PermissionedPoolFilterer) FilterUpdatedCommissionFees(opts *bind.FilterOpts) (*PermissionedPoolUpdatedCommissionFeesIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "UpdatedCommissionFees")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolUpdatedCommissionFeesIterator{contract: _PermissionedPool.contract, event: "UpdatedCommissionFees", logs: logs, sub: sub}, nil
}

// WatchUpdatedCommissionFees is a free log subscription operation binding the contract event 0xe8525a7862504bbe091b0440fe96979769664cae1625591ff0a9816512a5287b.
//
// Solidity: event UpdatedCommissionFees(uint256 protocolFee, uint256 operatorFee)
func (_PermissionedPool *PermissionedPoolFilterer) WatchUpdatedCommissionFees(opts *bind.WatchOpts, sink chan<- *PermissionedPoolUpdatedCommissionFees) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "UpdatedCommissionFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolUpdatedCommissionFees)
				if err := _PermissionedPool.contract.UnpackLog(event, "UpdatedCommissionFees", log); err != nil {
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

// ParseUpdatedCommissionFees is a log parse operation binding the contract event 0xe8525a7862504bbe091b0440fe96979769664cae1625591ff0a9816512a5287b.
//
// Solidity: event UpdatedCommissionFees(uint256 protocolFee, uint256 operatorFee)
func (_PermissionedPool *PermissionedPoolFilterer) ParseUpdatedCommissionFees(log types.Log) (*PermissionedPoolUpdatedCommissionFees, error) {
	event := new(PermissionedPoolUpdatedCommissionFees)
	if err := _PermissionedPool.contract.UnpackLog(event, "UpdatedCommissionFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolUpdatedStaderConfigIterator is returned from FilterUpdatedStaderConfig and is used to iterate over the raw logs and unpacked data for UpdatedStaderConfig events raised by the PermissionedPool contract.
type PermissionedPoolUpdatedStaderConfigIterator struct {
	Event *PermissionedPoolUpdatedStaderConfig // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolUpdatedStaderConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolUpdatedStaderConfig)
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
		it.Event = new(PermissionedPoolUpdatedStaderConfig)
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
func (it *PermissionedPoolUpdatedStaderConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolUpdatedStaderConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolUpdatedStaderConfig represents a UpdatedStaderConfig event raised by the PermissionedPool contract.
type PermissionedPoolUpdatedStaderConfig struct {
	StaderConfig common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedStaderConfig is a free log retrieval operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionedPool *PermissionedPoolFilterer) FilterUpdatedStaderConfig(opts *bind.FilterOpts) (*PermissionedPoolUpdatedStaderConfigIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolUpdatedStaderConfigIterator{contract: _PermissionedPool.contract, event: "UpdatedStaderConfig", logs: logs, sub: sub}, nil
}

// WatchUpdatedStaderConfig is a free log subscription operation binding the contract event 0xdb2219043d7b197cb235f1af0cf6d782d77dee3de19e3f4fb6d39aae633b4485.
//
// Solidity: event UpdatedStaderConfig(address staderConfig)
func (_PermissionedPool *PermissionedPoolFilterer) WatchUpdatedStaderConfig(opts *bind.WatchOpts, sink chan<- *PermissionedPoolUpdatedStaderConfig) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "UpdatedStaderConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolUpdatedStaderConfig)
				if err := _PermissionedPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
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
func (_PermissionedPool *PermissionedPoolFilterer) ParseUpdatedStaderConfig(log types.Log) (*PermissionedPoolUpdatedStaderConfig, error) {
	event := new(PermissionedPoolUpdatedStaderConfig)
	if err := _PermissionedPool.contract.UnpackLog(event, "UpdatedStaderConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolValidatorDepositedOnBeaconChainIterator is returned from FilterValidatorDepositedOnBeaconChain and is used to iterate over the raw logs and unpacked data for ValidatorDepositedOnBeaconChain events raised by the PermissionedPool contract.
type PermissionedPoolValidatorDepositedOnBeaconChainIterator struct {
	Event *PermissionedPoolValidatorDepositedOnBeaconChain // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolValidatorDepositedOnBeaconChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolValidatorDepositedOnBeaconChain)
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
		it.Event = new(PermissionedPoolValidatorDepositedOnBeaconChain)
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
func (it *PermissionedPoolValidatorDepositedOnBeaconChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolValidatorDepositedOnBeaconChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolValidatorDepositedOnBeaconChain represents a ValidatorDepositedOnBeaconChain event raised by the PermissionedPool contract.
type PermissionedPoolValidatorDepositedOnBeaconChain struct {
	ValidatorId *big.Int
	PubKey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorDepositedOnBeaconChain is a free log retrieval operation binding the contract event 0xbef89de94658b7ef396ba7f9316542858c893c9011602906b1a2ad18d0a99c35.
//
// Solidity: event ValidatorDepositedOnBeaconChain(uint256 indexed validatorId, bytes pubKey)
func (_PermissionedPool *PermissionedPoolFilterer) FilterValidatorDepositedOnBeaconChain(opts *bind.FilterOpts, validatorId []*big.Int) (*PermissionedPoolValidatorDepositedOnBeaconChainIterator, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "ValidatorDepositedOnBeaconChain", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolValidatorDepositedOnBeaconChainIterator{contract: _PermissionedPool.contract, event: "ValidatorDepositedOnBeaconChain", logs: logs, sub: sub}, nil
}

// WatchValidatorDepositedOnBeaconChain is a free log subscription operation binding the contract event 0xbef89de94658b7ef396ba7f9316542858c893c9011602906b1a2ad18d0a99c35.
//
// Solidity: event ValidatorDepositedOnBeaconChain(uint256 indexed validatorId, bytes pubKey)
func (_PermissionedPool *PermissionedPoolFilterer) WatchValidatorDepositedOnBeaconChain(opts *bind.WatchOpts, sink chan<- *PermissionedPoolValidatorDepositedOnBeaconChain, validatorId []*big.Int) (event.Subscription, error) {

	var validatorIdRule []interface{}
	for _, validatorIdItem := range validatorId {
		validatorIdRule = append(validatorIdRule, validatorIdItem)
	}

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "ValidatorDepositedOnBeaconChain", validatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolValidatorDepositedOnBeaconChain)
				if err := _PermissionedPool.contract.UnpackLog(event, "ValidatorDepositedOnBeaconChain", log); err != nil {
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

// ParseValidatorDepositedOnBeaconChain is a log parse operation binding the contract event 0xbef89de94658b7ef396ba7f9316542858c893c9011602906b1a2ad18d0a99c35.
//
// Solidity: event ValidatorDepositedOnBeaconChain(uint256 indexed validatorId, bytes pubKey)
func (_PermissionedPool *PermissionedPoolFilterer) ParseValidatorDepositedOnBeaconChain(log types.Log) (*PermissionedPoolValidatorDepositedOnBeaconChain, error) {
	event := new(PermissionedPoolValidatorDepositedOnBeaconChain)
	if err := _PermissionedPool.contract.UnpackLog(event, "ValidatorDepositedOnBeaconChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermissionedPoolValidatorPreDepositedOnBeaconChainIterator is returned from FilterValidatorPreDepositedOnBeaconChain and is used to iterate over the raw logs and unpacked data for ValidatorPreDepositedOnBeaconChain events raised by the PermissionedPool contract.
type PermissionedPoolValidatorPreDepositedOnBeaconChainIterator struct {
	Event *PermissionedPoolValidatorPreDepositedOnBeaconChain // Event containing the contract specifics and raw log

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
func (it *PermissionedPoolValidatorPreDepositedOnBeaconChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermissionedPoolValidatorPreDepositedOnBeaconChain)
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
		it.Event = new(PermissionedPoolValidatorPreDepositedOnBeaconChain)
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
func (it *PermissionedPoolValidatorPreDepositedOnBeaconChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermissionedPoolValidatorPreDepositedOnBeaconChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermissionedPoolValidatorPreDepositedOnBeaconChain represents a ValidatorPreDepositedOnBeaconChain event raised by the PermissionedPool contract.
type PermissionedPoolValidatorPreDepositedOnBeaconChain struct {
	PubKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorPreDepositedOnBeaconChain is a free log retrieval operation binding the contract event 0xa35366ad083efbaee7949ff15c68508b95e2c0441248e80d73da97ac82bc1f10.
//
// Solidity: event ValidatorPreDepositedOnBeaconChain(bytes pubKey)
func (_PermissionedPool *PermissionedPoolFilterer) FilterValidatorPreDepositedOnBeaconChain(opts *bind.FilterOpts) (*PermissionedPoolValidatorPreDepositedOnBeaconChainIterator, error) {

	logs, sub, err := _PermissionedPool.contract.FilterLogs(opts, "ValidatorPreDepositedOnBeaconChain")
	if err != nil {
		return nil, err
	}
	return &PermissionedPoolValidatorPreDepositedOnBeaconChainIterator{contract: _PermissionedPool.contract, event: "ValidatorPreDepositedOnBeaconChain", logs: logs, sub: sub}, nil
}

// WatchValidatorPreDepositedOnBeaconChain is a free log subscription operation binding the contract event 0xa35366ad083efbaee7949ff15c68508b95e2c0441248e80d73da97ac82bc1f10.
//
// Solidity: event ValidatorPreDepositedOnBeaconChain(bytes pubKey)
func (_PermissionedPool *PermissionedPoolFilterer) WatchValidatorPreDepositedOnBeaconChain(opts *bind.WatchOpts, sink chan<- *PermissionedPoolValidatorPreDepositedOnBeaconChain) (event.Subscription, error) {

	logs, sub, err := _PermissionedPool.contract.WatchLogs(opts, "ValidatorPreDepositedOnBeaconChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermissionedPoolValidatorPreDepositedOnBeaconChain)
				if err := _PermissionedPool.contract.UnpackLog(event, "ValidatorPreDepositedOnBeaconChain", log); err != nil {
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

// ParseValidatorPreDepositedOnBeaconChain is a log parse operation binding the contract event 0xa35366ad083efbaee7949ff15c68508b95e2c0441248e80d73da97ac82bc1f10.
//
// Solidity: event ValidatorPreDepositedOnBeaconChain(bytes pubKey)
func (_PermissionedPool *PermissionedPoolFilterer) ParseValidatorPreDepositedOnBeaconChain(log types.Log) (*PermissionedPoolValidatorPreDepositedOnBeaconChain, error) {
	event := new(PermissionedPoolValidatorPreDepositedOnBeaconChain)
	if err := _PermissionedPool.contract.UnpackLog(event, "ValidatorPreDepositedOnBeaconChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
