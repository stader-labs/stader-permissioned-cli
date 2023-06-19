package validator

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"

	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
	"github.com/urfave/cli"
)

func SendClRewards(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)

	canSendClRewardsResponse, err := staderClient.CanSendClRewards(validatorPubKey)
	if err != nil {
		return err
	}
	if canSendClRewardsResponse.OperatorNotRegistered {
		fmt.Printf("Operator not registered\n")
		return nil
	}
	if canSendClRewardsResponse.NoClRewards {
		fmt.Printf("No CL rewards to withdraw for validator %s\n", validatorPubKey.String())
		return nil
	}
	if canSendClRewardsResponse.TooManyClRewards {
		fmt.Printf("It appears that the validator has exited. Please wait for oracles to settle the validator funds after funds are withdrawn in the beacon chain.")
		return nil
	}
	if canSendClRewardsResponse.ValidatorNotFound {
		fmt.Printf("Validator %s not found\n", validatorPubKey.String())
		return nil
	}
	if canSendClRewardsResponse.VaultAlreadySettled {
		fmt.Printf("Vault for validator %s has already been settled\n", validatorPubKey.String())
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(canSendClRewardsResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to Send CL rewards for validator %s to the claim vault?", validatorPubKey))) {
		fmt.Println("Cancelled.")
		return nil
	}

	res, err := staderClient.SendClRewards(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Sending %.6f CL Rewards to Claim Vault\n\n", math.RoundDown(eth.WeiToEth(res.ClRewardsAmount), 6))
	cliutils.PrintTransactionHash(staderClient, res.TxHash)
	if _, err = staderClient.WaitForTransaction(res.TxHash); err != nil {
		return err
	}

	// Log & return
	fmt.Printf("Successfully Sent %.6f CL Rewards to Claim Vault\n\n", math.RoundDown(eth.WeiToEth(res.ClRewardsAmount), 6))

	return nil
}
