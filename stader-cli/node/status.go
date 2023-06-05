package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
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
		if !status.OperatorWhitelisted {
			fmt.Printf("The operator is not whitelisted with Stader. Please connect with the Stader Team to whitelist your address to be a part of the permissioned operator pool")
			return nil
		} else {
			fmt.Printf("The operator is not registered with Stader. Please use the %sstader-permissioned-cli node register%s to register with Stader", log.ColorGreen, log.ColorReset)
			return nil
		}
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
		fmt.Printf("The Operator has %.6f ETH as unclaimed socializing pool rewards\n", math.RoundDown(eth.WeiToEth(totalUnclaimedSocializingPoolEth), 18))
		fmt.Printf("To claim socializing pool rewards use the %sstader-permissioned-cli node claim-sp-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
	}

	if status.OperatorClaimVaultBalance.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf(
			"The Operator %s%s%s has aggregated total rewards of %.6f ETH in the claim vault\n\n",
			log.ColorBlue,
			status.AccountAddress,
			log.ColorReset,
			math.RoundDown(eth.WeiToEth(status.OperatorClaimVaultBalance), 6))
		fmt.Printf("To transfer the claims to your operator reward address use the %sstader-permissioned-cli node claim-rewards%s command\n\n", log.ColorGreen, log.ColorReset)
	}

	fmt.Printf("%s=== Registered Validator Details ===%s\n", log.ColorGreen, log.ColorReset)

	fmt.Printf("To view details of each validator, please use the %sstader-permissioned-cli validator status%s command\n\n", log.ColorGreen, log.ColorReset)

	return nil
}
