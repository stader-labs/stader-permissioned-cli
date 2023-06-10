package node

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	types2 "github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

func EstimateAddValidatorKeys(pnr *stader.PermissionedNodeRegistryContractManager, pubKeys [][]byte, preDepositSignatures [][]byte, depositSignatures [][]byte, opts *bind.TransactOpts) (stader.GasInfo, error) {
	return pnr.PermissionedNodeRegistryContract.GetTransactionGasInfo(opts, "addValidatorKeys", pubKeys, preDepositSignatures, depositSignatures)
}

func AddValidatorKeys(pnr *stader.PermissionedNodeRegistryContractManager, pubKeys [][]byte, preDepositSignatures [][]byte, depositSignatures [][]byte, opts *bind.TransactOpts) (*types.Transaction, error) {
	tx, err := pnr.PermissionedNodeRegistry.AddValidatorKeys(opts, pubKeys, preDepositSignatures, depositSignatures)
	if err != nil {
		return nil, fmt.Errorf("could not add validator keys: %w", err)
	}

	return tx, nil
}

func EstimateDistributeRewards(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.TransactOpts) (stader.GasInfo, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return stader.GasInfo{}, err
	}

	return vwv.ValidatorWithdrawVaultContract.GetTransactionGasInfo(opts, "distributeRewards")
}

func DistributeRewards(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.TransactOpts) (*types.Transaction, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return nil, err
	}

	return vwv.ValidatorWithdrawVault.DistributeRewards(opts)
}

func GetTotalValidatorKeys(pnr *stader.PermissionedNodeRegistryContractManager, operatorId *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionedNodeRegistry.GetOperatorTotalKeys(opts, operatorId)
}

func GetTotalNonTerminalValidatorKeys(pnr *stader.PermissionedNodeRegistryContractManager, operatorAddress common.Address, maxPaginationIndex *big.Int, opts *bind.CallOpts) (uint64, error) {
	return pnr.PermissionedNodeRegistry.GetOperatorTotalNonTerminalKeys(opts, operatorAddress, big.NewInt(0), maxPaginationIndex)
}

func GetMaxValidatorKeysPerOperator(pnr *stader.PermissionedNodeRegistryContractManager, opts *bind.CallOpts) (uint64, error) {
	return pnr.PermissionedNodeRegistry.MaxNonTerminalKeyPerOperator(opts)
}

func GetValidatorIdByOperatorId(pnr *stader.PermissionedNodeRegistryContractManager, operatorId *big.Int, validatorIndex *big.Int, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionedNodeRegistry.ValidatorIdsByOperatorId(opts, operatorId, validatorIndex)
}

func GetValidatorInfo(pnr *stader.PermissionedNodeRegistryContractManager, validatorId *big.Int, opts *bind.CallOpts) (contracts.Validator, error) {
	return pnr.PermissionedNodeRegistry.ValidatorRegistry(opts, validatorId)
}

func ComputeWithdrawVaultAddress(vfcm *stader.VaultFactoryContractManager, poolType uint8, operatorId *big.Int, validatorCount *big.Int, opts *bind.CallOpts) (common.Address, error) {
	return vfcm.VaultFactory.ComputeWithdrawVaultAddress(opts, poolType, operatorId, validatorCount)
}

func GetValidatorWithdrawalCredential(vfcm *stader.VaultFactoryContractManager, withdrawVaultAddress common.Address, opts *bind.CallOpts) (common.Hash, error) {
	withdrawalCredentials := new(common.Hash)
	if err := vfcm.VaultFactoryContract.Call(opts, withdrawalCredentials, "getValidatorWithdrawCredential", withdrawVaultAddress); err != nil {
		return common.Hash{}, fmt.Errorf("Could not get vault withdrawal credentials: %w", err)
	}

	return *withdrawalCredentials, nil
}

func CalculateValidatorWithdrawVaultWithdrawShare(executionClient stader.ExecutionClient, validatorWithdrawVaultAddress common.Address, opts *bind.CallOpts) (types2.RewardShare, error) {
	vwv, err := stader.NewValidatorWithdrawVaultFactory(executionClient, validatorWithdrawVaultAddress)
	if err != nil {
		return types2.RewardShare{}, err
	}

	return vwv.ValidatorWithdrawVault.CalculateValidatorWithdrawalShare(opts)
}

func GetValidatorIdByPubKey(pnr *stader.PermissionedNodeRegistryContractManager, validatorPubKey []byte, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionedNodeRegistry.ValidatorIdByPubkey(opts, validatorPubKey)
}

func GetNextValidatorId(pnr *stader.PermissionedNodeRegistryContractManager, opts *bind.CallOpts) (*big.Int, error) {
	return pnr.PermissionedNodeRegistry.NextValidatorId(opts)
}

func GetInputKeyCountLimit(pnr *stader.PermissionedNodeRegistryContractManager, opts *bind.CallOpts) (uint16, error) {
	return pnr.PermissionedNodeRegistry.InputKeyCountLimit(opts)
}

func GetValidatorInfosByOperator(pnr *stader.PermissionedNodeRegistryContractManager, operatorAddress common.Address, pageNumber *big.Int, pageSize *big.Int, opts *bind.CallOpts) ([]contracts.Validator, error) {
	return pnr.PermissionedNodeRegistry.GetValidatorsByOperator(opts, operatorAddress, pageNumber, pageSize)
}

func GetAllValidatorsInfoByOperator(pnr *stader.PermissionedNodeRegistryContractManager, operatorAddress common.Address, opts *bind.CallOpts) ([]contracts.Validator, error) {
	finalValidators := []contracts.Validator{}
	pageNumber := big.NewInt(1)
	pageSize := big.NewInt(100)

	for {
		validators, err := pnr.PermissionedNodeRegistry.GetValidatorsByOperator(opts, operatorAddress, pageNumber, pageSize)
		if err != nil {
			return nil, err
		}
		if len(validators) == 0 {
			break
		}

		finalValidators = append(finalValidators, validators...)
		pageNumber.Add(pageNumber, big.NewInt(1))
	}

	return finalValidators, nil
}
