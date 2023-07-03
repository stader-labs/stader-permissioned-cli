package stdr

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/stader-lib/contracts"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
	"time"
)

var ValidatorState = map[uint8]string{
	0: "Initialized",
	1: "Invalid Signature Submitted",
	2: "Validator Deposit Front Run",
	3: "Validator Queued for 31Eth deposit",
	4: "Validator has been matched with 31Eth deposit",
	5: "Funds Settled",
}

type ValidatorInfo struct {
	Status                           uint8
	StatusToDisplay                  string
	Pubkey                           []byte
	PreDepositSignature              []byte
	DepositSignature                 []byte
	WithdrawVaultAddress             common.Address
	WithdrawVaultRewardBalance       *big.Int
	WithdrawVaultWithdrawableBalance *big.Int
	CrossedRewardsThreshold          bool
	OperatorId                       *big.Int
	DepositBlock                     *big.Int
	DepositTime                      time.Time
	WithdrawnBlock                   *big.Int
	WithdrawnTime                    time.Time
}

func GetAllValidatorsRegisteredWithOperator(pnr *stader.PermissionedNodeRegistryContractManager, operatorId *big.Int, operatorAddress common.Address, opts *bind.CallOpts) (map[types.ValidatorPubkey]contracts.Validator, []types.ValidatorPubkey, error) {
	validators, err := node.GetAllValidatorsInfoByOperator(pnr, operatorAddress, opts)
	if err != nil {
		return nil, []types.ValidatorPubkey{}, err
	}

	validatorInfoMap := make(map[types.ValidatorPubkey]contracts.Validator)
	validatorPubKeys := []types.ValidatorPubkey{}
	for _, validator := range validators {
		pubKey := types.BytesToValidatorPubkey(validator.Pubkey)
		validatorInfoMap[pubKey] = validator
		validatorPubKeys = append(validatorPubKeys, pubKey)
	}

	return validatorInfoMap, validatorPubKeys, err

}

func IsValidatorTerminal(validatorInfo contracts.Validator) bool {
	return validatorInfo.Status == 1 || validatorInfo.Status == 2
}

func GetValidatorRunningStatus(beaconValidatorStatus beacon.ValidatorStatus, validatorContractInfo contracts.Validator) (string, error) {
	if validatorContractInfo.Status != 4 || !beaconValidatorStatus.Exists {
		return ValidatorState[validatorContractInfo.Status], nil
	}

	switch beaconValidatorStatus.Status {
	case beacon.ValidatorState_ActiveOngoing:
		return "Active and ongoing", nil
	case beacon.ValidatorState_ActiveSlashed:
		return "Active and slashed", nil
	case beacon.ValidatorState_ActiveExiting:
		return "Active and exiting", nil
	case beacon.ValidatorState_ExitedUnslashed:
		return "Exited and unslashed", nil
	case beacon.ValidatorState_ExitedSlashed:
		return "Exited and slashed", nil
	case beacon.ValidatorState_WithdrawalDone:
		return "Withdrawal done", nil
	case beacon.ValidatorState_WithdrawalPossible:
		return "Withdrawal possible", nil
	case beacon.ValidatorState_PendingInitialized:
		return "Pending initialized", nil
	case beacon.ValidatorState_PendingQueued:
		return "Pending queued", nil

	}

	return "", nil
}
