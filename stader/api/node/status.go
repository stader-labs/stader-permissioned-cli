package node

import (
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
	pool_utils "github.com/stader-labs/stader-node/stader-lib/pool-utils"
	socializing_pool "github.com/stader-labs/stader-node/stader-lib/socializing-pool"
	stader_config "github.com/stader-labs/stader-node/stader-lib/stader-config"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
	"time"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/tokens"
	"github.com/urfave/cli"
)

func GetClaimedAndUnclaimedSocializingPoolMerkles(c *cli.Context) ([]stader_backend.CycleMerkleProofs, []stader_backend.CycleMerkleProofs, error) {
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, nil, err
	}
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, nil, err
	}

	rewardDetails, err := socializing_pool.GetRewardDetails(sp, nil)
	if err != nil {
		return nil, nil, err
	}

	unclaimedMerkles := []stader_backend.CycleMerkleProofs{}
	claimedMerkles := []stader_backend.CycleMerkleProofs{}
	for i := int64(1); i < rewardDetails.CurrentIndex.Int64(); i++ {
		cycleMerkleProof, exists, err := cfg.StaderNode.ReadCycleCache(i)
		if err != nil {
			return nil, nil, err
		}
		if !exists {
			continue
		}
		claimed, err := socializing_pool.HasClaimedRewards(sp, nodeAccount.Address, big.NewInt(i), nil)
		if err != nil {
			return nil, nil, err
		}

		if claimed {
			claimedMerkles = append(claimedMerkles, cycleMerkleProof)
		} else {
			unclaimedMerkles = append(unclaimedMerkles, cycleMerkleProof)
		}
	}

	return claimedMerkles, unclaimedMerkles, nil
}

func getStatus(c *cli.Context) (*api.NodeStatusResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	sdcfg, err := services.GetStaderConfigContract(c)
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
	sp, err := services.GetSocializingPoolContract(c)
	if err != nil {
		return nil, err
	}
	w3signer, err := services.GetWeb3SignerClient(c)
	if err != nil {
		return nil, err
	}
	cfg, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.NodeStatusResponse{}

	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	response.AccountAddress = nodeAccount.Address

	isWhitelisted, err := node.IsOperatorWhitelisted(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	response.OperatorWhitelisted = isWhitelisted
	accountEthBalance, err := tokens.GetEthBalance(pnr.Client, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	response.AccountBalances.ETH = accountEthBalance

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("Getting operator info...\n")
	operatorRegistry, err := node.GetOperatorInfo(pnr, operatorId, nil)
	if err != nil {
		return nil, err
	}

	web3SignerUrl := cfg.ExternalWeb3Signer.HttpUrl.Value.(string)
	response.Web3SignerUrl = web3SignerUrl

	// get connection status
	err = w3signer.HealthCheck()
	if err != nil {
		response.Web3SignerConnectionSuccess = false
		response.Web3SignerConnectionError = err.Error()
	} else {
		response.Web3SignerConnectionSuccess = true
	}

	beaconUrl := cfg.ExternalBeacon.HttpUrl.Value.(string)
	executionUrl := cfg.ExternalExecution.HttpUrl.Value.(string)

	response.BeaconChainUrl = beaconUrl
	response.ExecutionChainUrl = executionUrl

	if operatorRegistry.OperatorName != "" {
		response.Registered = true
		response.OperatorId = operatorId
		response.OperatorName = operatorRegistry.OperatorName
		response.OperatorActive = operatorRegistry.Active
		response.OperatorAddress = operatorRegistry.OperatorAddress
		response.OperatorRewardAddress = operatorRegistry.OperatorRewardAddress

		operatorReward, err := tokens.GetEthBalance(pnr.Client, operatorRegistry.OperatorRewardAddress, nil)
		if err != nil {
			return nil, err
		}
		response.OperatorRewardInETH = operatorReward

		rewardCycleDetails, err := socializing_pool.GetRewardDetails(sp, nil)
		if err != nil {
			return nil, err
		}
		response.SocializingPoolRewardCycleDetails = rewardCycleDetails
		socializingPoolStartTimestamp := time.Now()
		response.SocializingPoolStartTime = socializingPoolStartTimestamp

		totalValidatorKeys, err := node.GetTotalValidatorKeys(pnr, operatorId, nil)
		if err != nil {
			return nil, err
		}

		totalNonTerminalValidatorKeys, err := node.GetTotalNonTerminalValidatorKeys(pnr, nodeAccount.Address, totalValidatorKeys, nil)
		if err != nil {
			return nil, err
		}

		response.TotalNonTerminalValidators = big.NewInt(int64(totalNonTerminalValidatorKeys))

		validatorInfoArray := make([]stdr.ValidatorInfo, totalValidatorKeys.Int64())

		totalClRewards := big.NewInt(0)
		for i := int64(0); i < totalValidatorKeys.Int64(); i++ {
			validatorIndex, err := node.GetValidatorIdByOperatorId(pnr, operatorId, big.NewInt(i), nil)
			if err != nil {
				return nil, err
			}
			validatorContractInfo, err := node.GetValidatorInfo(pnr, validatorIndex, nil)
			if err != nil {
				return nil, err
			}
			withdrawVaultBalance, err := tokens.GetEthBalance(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
			if err != nil {
				return nil, err
			}
			withdrawVaultRewardShares, err := pool_utils.CalculateRewardShare(putils, 2, withdrawVaultBalance, nil)
			if err != nil {
				return nil, err
			}
			rewardsThreshold, err := stader_config.GetRewardsThreshold(sdcfg, nil)
			if err != nil {
				return nil, err
			}
			crossedRewardThreshold := false
			if withdrawVaultBalance.Cmp(rewardsThreshold) > 0 {
				crossedRewardThreshold = true
			}

			withdrawVaultWithdrawShares, err := node.CalculateValidatorWithdrawVaultWithdrawShare(pnr.Client, validatorContractInfo.WithdrawVaultAddress, nil)
			if err != nil {
				return nil, err
			}
			validatorWithdrawVaultWithdrawShares := withdrawVaultWithdrawShares.OperatorShare

			validatorBeaconStatus, err := bc.GetValidatorStatus(types.BytesToValidatorPubkey(validatorContractInfo.Pubkey), nil)
			if err != nil {
				return nil, err
			}

			validatorDisplayStatus, err := stdr.GetValidatorRunningStatus(validatorBeaconStatus, validatorContractInfo)
			if err != nil {
				return nil, err
			}

			depositTime, err := eth1.ConvertBlockToTimestamp(c, validatorContractInfo.DepositBlock.Int64())
			if err != nil {
				return nil, err
			}
			withdrawTime, err := eth1.ConvertBlockToTimestamp(c, validatorContractInfo.WithdrawnBlock.Int64())
			if err != nil {
				return nil, err
			}

			validatorInfo := stdr.ValidatorInfo{
				Status:                           validatorContractInfo.Status,
				StatusToDisplay:                  validatorDisplayStatus,
				Pubkey:                           validatorContractInfo.Pubkey,
				PreDepositSignature:              validatorContractInfo.PreDepositSignature,
				DepositSignature:                 validatorContractInfo.DepositSignature,
				WithdrawVaultAddress:             validatorContractInfo.WithdrawVaultAddress,
				WithdrawVaultRewardBalance:       withdrawVaultRewardShares.OperatorShare,
				CrossedRewardsThreshold:          crossedRewardThreshold,
				WithdrawVaultWithdrawableBalance: validatorWithdrawVaultWithdrawShares,
				OperatorId:                       validatorContractInfo.OperatorId,
				DepositBlock:                     validatorContractInfo.DepositBlock,
				DepositTime:                      depositTime,
				WithdrawnBlock:                   validatorContractInfo.WithdrawnBlock,
				WithdrawnTime:                    withdrawTime,
			}

			totalClRewards.Add(totalClRewards, withdrawVaultRewardShares.OperatorShare)
			validatorInfoArray[i] = validatorInfo
		}

		response.ValidatorInfos = validatorInfoArray
		response.TotalClRewards = totalClRewards

		claimedMerkles, unclaimedMerkles, err := GetClaimedAndUnclaimedSocializingPoolMerkles(c)
		if err != nil {
			return nil, err
		}

		response.ClaimedSocializingPoolMerkles = claimedMerkles
		response.UnclaimedSocializingPoolMerkles = unclaimedMerkles

	} else {
		response.ValidatorInfos = []stdr.ValidatorInfo{}
		response.Registered = false
	}

	// Return response
	return &response, nil
}
