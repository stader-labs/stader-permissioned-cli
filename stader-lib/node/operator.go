package node

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	types2 "github.com/stader-labs/stader-node/stader-lib/types"
)

func EstimateOnboardNodeOperator(pnr *stader.PermissionedNodeRegistryContractManager, operatorName string, operatorRewarderAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionedNodeRegistryContract.GetTransactionGasInfo(opts, "onboardNodeOperator", operatorName, operatorRewarderAddress)
}

func OnboardNodeOperator(pnr *stader.PermissionedNodeRegistryContractManager, operatorName string, operatorRewarderAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionedNodeRegistry.OnboardNodeOperator(opts, operatorName, operatorRewarderAddress)
	if err != nil {
		return nil, fmt.Errorf("Could not onboard node operator: %w", err)
	}

	return tx, nil
}

func IsOperatorWhitelisted(pnr *stader.PermissionedNodeRegistryContractManager, operatorAddress common.Address, opts *bind.CallOpts) (bool, error) {
	return pnr.PermissionedNodeRegistry.PermissionList(opts, operatorAddress)
}

func EstimateUpdateOperatorName(pnr *stader.PermissionedNodeRegistryContractManager, operatorName string, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionedNodeRegistryContract.GetTransactionGasInfo(opts, "updateOperatorName", operatorName)
}

func UpdateOperatorName(pnr *stader.PermissionedNodeRegistryContractManager, operatorName string, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionedNodeRegistry.UpdateOperatorName(opts, operatorName)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func EstimateUpdateOperatorRewardAddress(pnr *stader.PermissionedNodeRegistryContractManager, operatorAddress common.Address, rewardAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionedNodeRegistryContract.GetTransactionGasInfo(opts, "proposeRewardAddress", operatorAddress, rewardAddress)
}

func UpdateOperatorRewardAddress(pnr *stader.PermissionedNodeRegistryContractManager, operatorAddress common.Address, rewardAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionedNodeRegistry.ProposeRewardAddress(opts, operatorAddress, rewardAddress)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func EstimateClaimOperatorRewards(orc *stader.OperatorRewardsCollectorContractManager, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return orc.OperatorRewardsCollectorContract.GetTransactionGasInfo(opts, "claim")
}

func ClaimOperatorRewards(orc *stader.OperatorRewardsCollectorContractManager, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := orc.OperatorRewardsCollector.Claim(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not claim operator rewards: %w", err)
	}

	return tx, nil
}

func GetOperatorId(pnr *stader.PermissionedNodeRegistryContractManager, nodeAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	operatorId, err := pnr.PermissionedNodeRegistry.OperatorIDByAddress(opts, nodeAddress)
	if err != nil {
		return nil, err
	}

	return operatorId, nil
}

func GetOperatorInfo(pnr *stader.PermissionedNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (types2.OperatorInfo, error) {
	operatorInfo, err := pnr.PermissionedNodeRegistry.OperatorStructById(opts, operatorId)
	if err != nil {
		return types2.OperatorInfo{}, err
	}

	return operatorInfo, nil
}

func GetNodeElRewardAddress(vf *stader.VaultFactoryContractManager, poolId uint8, operatorId *big.Int, opts *bind.CallOpts) (common.Address, error) {
	return vf.VaultFactory.ComputeNodeELRewardVaultAddress(opts, poolId, operatorId)
}

func GetOperatorRewardsCollectorBalance(orc *stader.OperatorRewardsCollectorContractManager, operatorAddress common.Address, opts *bind.CallOpts) (*big.Int, error) {
	return orc.OperatorRewardsCollector.Balances(opts, operatorAddress)
}

func GetSocializingPoolContract(pp *stader.PermissionlessPoolContractManager, opts *bind.CallOpts) (common.Address, error) {
	return pp.PermissionlessPool.GetSocializingPoolAddress(opts)
}
