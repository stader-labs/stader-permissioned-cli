package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func GetValidatorContractInfo(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.ValidatorInfoResponse, error) {
	prn, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	putils, err := services.GetPoolUtilsContract(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}

	response := api.ValidatorInfoResponse{}

	validatorId, err := node.GetValidatorIdByPubKey(prn, validatorPubKey.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	validatorInfo, err := node.GetValidatorInfo(prn, validatorId, nil)
	if err != nil {
		return nil, err
	}
	withdrawVaultBalance, err := tokens.GetEthBalance(prn.Client, validatorInfo.WithdrawVaultAddress, nil)
	if err != nil {
		return nil, err
	}
	withdrawVaultRewardShares, err := pool_utils.CalculateRewardShare(putils, 2, withdrawVaultBalance, nil)
	if err != nil {
		return nil, err
	}

	validatorBeaconStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		return nil, err
	}

	validatorDisplayStatus, err := stdr.GetValidatorRunningStatus(validatorBeaconStatus, validatorInfo)
	if err != nil {
		return nil, err
	}

	response.ClRewards = withdrawVaultRewardShares.OperatorShare
	response.ValidatorInfo = validatorInfo
	response.DisplayStatus = validatorDisplayStatus

	return &response, nil
}
