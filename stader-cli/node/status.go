package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
	"math/big"
)

func getStatus(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	clientStatus, err := staderClient.GetClientStatus()
	if err != nil {
		return err
	}

	// Get node status
	status, err := staderClient.NodeStatus()
	if err != nil {
		return err
	}

	totalRegisteredValidators := status.TotalNonTerminalValidators

	totalUnclaimedSocializingPoolEth := big.NewInt(0)
	totalUnclaimedSocializingPoolSd := big.NewInt(0)
	for _, merkle := range status.UnclaimedSocializingPoolMerkles {
		ethRewards, ok := big.NewInt(0).SetString(merkle.Eth, 10)
		if !ok {
			return fmt.Errorf("error while converting eth rewards: %s to big int", merkle.Eth)
		}

		sdRewards, ok := big.NewInt(0).SetString(merkle.Sd, 10)
		if !ok {
			return fmt.Errorf("error while converting sd rewards: %s to big int", merkle.Sd)
		}

		totalUnclaimedSocializingPoolEth.Add(totalUnclaimedSocializingPoolEth, ethRewards)
		totalUnclaimedSocializingPoolSd.Add(totalUnclaimedSocializingPoolSd, sdRewards)
	}

	fmt.Printf("%s=== Services Connection Status ===%s\n", log.ColorGreen, log.ColorReset)
	if status.Web3SignerConnectionSuccess {
		fmt.Printf("The node is connected to the Web3Signer at %s.\n\n", status.Web3SignerUrl)
	} else {
		fmt.Printf("Web3Signer Connection Failed with error: %s. \n\n", status.Web3SignerConnectionError)
	}

	if clientStatus.BcManagerStatus.PrimaryClientStatus.Error != "" {
		fmt.Printf("Connection to beacon chain failed with: %s\n\n", clientStatus.BcManagerStatus.PrimaryClientStatus.Error)
	} else {
		fmt.Printf("The node is using %s to connect to the beacon chain.\n\n", status.BeaconChainUrl)
	}

	if clientStatus.EcManagerStatus.PrimaryClientStatus.Error != "" {
		fmt.Printf("Connection to execution chain failed with: %s\n\n", clientStatus.EcManagerStatus.PrimaryClientStatus.Error)
	} else {
		fmt.Printf("The node is using %s to connect to the execution chain.\n\n", status.ExecutionChainUrl)
	}

	// Account address & balances
	fmt.Printf("%s=== Account and Balances ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf(
		"The node %s%s%s has a balance of %.6f ETH.\n\n",
		log.ColorBlue,
		status.AccountAddress,
		log.ColorReset,
		math.RoundDown(eth.WeiToEth(status.AccountBalances.ETH), 6))

	fmt.Printf("%s=== Operator Registration Details ===%s\n", log.ColorGreen, log.ColorReset)

	if !status.Registered {
		fmt.Printf("The node is not registered with Stader. Please use the %sstader-permissioned-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
		return nil
	}

	fmt.Printf("The Operator is registered with Stader. Below are operator details:\n")
	fmt.Printf("Operator Id: %d\n\n", status.OperatorId)
	fmt.Printf("Operator Name: %s\n\n", status.OperatorName)
	fmt.Printf("Operator Address: %s\n\n", status.OperatorAddress.String())
	fmt.Printf("Operator Reward Address: %s\n\n", status.OperatorRewardAddress.String())
	if status.OperatorActive {
		fmt.Printf("Operator Status: Active\n\n")
	} else {
		fmt.Printf("Operator Status: Not Active\n\n")
	}
	fmt.Printf("The Operator has registered a total of %d validators\n\n", len(status.ValidatorInfos))

	if totalUnclaimedSocializingPoolEth.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("The Operator has %.6f ETH as unclaimed ETH rewards\n\n", math.RoundDown(eth.WeiToEth(totalUnclaimedSocializingPoolEth), 18))
		fmt.Printf("To claim ETH rewards using the %sstader-permissioned-cli node claim-sp-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
	}

	fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)

	if totalRegisteredValidators.Int64() <= 0 {
		fmt.Printf("The node has no registered validators. Please use the %sstader-permissioned-cli node deposit%s command to register a validator with Stader\n\n", log.ColorGreen, log.ColorReset)
		return nil
	}

	for i := 0; i < len(status.ValidatorInfos); i++ {
		fmt.Printf("%d)\n", i+1)
		validatorInfo := status.ValidatorInfos[i]
		validatorPubKey := types.BytesToValidatorPubkey(validatorInfo.Pubkey)
		fmt.Printf("-Validator Pub Key: %s\n\n", validatorPubKey)
		fmt.Printf("-Validator Status: %s\n\n", validatorInfo.StatusToDisplay)
		fmt.Printf("-Validator Withdraw Vault: %s\n\n", validatorInfo.WithdrawVaultAddress)
		if validatorInfo.WithdrawVaultRewardBalance.Int64() > 0 && !validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("-Validator Skimmed Rewards: %.6f\n", math.RoundDown(eth.WeiToEth(validatorInfo.WithdrawVaultRewardBalance), 18))
			fmt.Printf("To claim skimmed rewards use the %sstader-permissioned-cli node claim-cl-rewards %s command\n\n", log.ColorGreen, log.ColorReset)
		} else if validatorInfo.CrossedRewardsThreshold {
			fmt.Printf("The validator has crossed the reward threshold.\n")
			fmt.Printf("If you have exited the validator, Please wait for Stader Oracles to settle your funds!\n")
			fmt.Printf("If you have not exited the validator, Please reach out to the Stader Team on discord!\n\n")
		}

		if validatorInfo.Status > 3 {
			fmt.Printf("-Deposit time: %s\n\n", validatorInfo.DepositTime.Format("2006-01-02 15:04:05"))
		}

		if validatorInfo.WithdrawnBlock.Int64() > 0 {
			fmt.Printf("-Withdraw Time: %s\n\n", validatorInfo.WithdrawnTime.Format("2006-01-02 15:04:05"))
		}

		fmt.Printf("\n\n")
	}

	return nil
}
