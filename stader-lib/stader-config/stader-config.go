package stader_config

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"math/big"
)

func GetRewardsThreshold(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetRewardsThreshold(opts)
}

func GetOperatorNameMaxLength(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return sdConfig.StaderConfig.GetOperatorMaxNameLength(opts)
}

func GetSocializingPoolContractAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetPermissionedSocializingPool(opts)
}

func GetPermissionedNodeRegistryAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetPermissionedNodeRegistry(opts)
}

func GetVaultFactoryAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetVaultFactory(opts)
}

func GetPermissionedPoolAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetPermissionedPool(opts)
}

func GetSdTokenAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetStaderToken(opts)
}

func GetETHxTokenAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetETHxToken(opts)
}

func GetStaderOracleAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetStaderOracle(opts)
}

func GetPoolUtilsAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetPoolUtils(opts)
}

func GetPenaltyTrackerAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetPenaltyContract(opts)
}

func GetStakePoolManagerAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetStakePoolManager(opts)
}

func GetOperatorRewardsCollectorAddress(sdConfig *stader.StaderConfigContractManager, opts *bind.CallOpts) (common.Address, error) {
	return sdConfig.StaderConfig.GetOperatorRewardsCollector(opts)
}
