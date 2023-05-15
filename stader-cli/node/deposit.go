package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
)

func nodeDeposit(c *cli.Context, validatorPubKeyList string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	canNodeDepositResponse, err := staderClient.CanNodeDeposit(validatorPubKeyList, true)
	if err != nil {
		return err
	}
	if canNodeDepositResponse.InsufficientBalance {
		fmt.Printf("Account does not have enough balance!")
		return nil
	}
	if canNodeDepositResponse.DepositPaused {
		fmt.Printf("Deposits are currently paused!")
		return nil
	}
	if canNodeDepositResponse.MaxValidatorLimitReached {
		fmt.Printf("Max validator limit reached")
		return nil
	}
	if canNodeDepositResponse.OperatorNotRegistered {
		fmt.Printf("Operator not registered")
		return nil
	}
	if canNodeDepositResponse.OperatorNotActive {
		fmt.Printf("Operator not active")
		return nil
	}

	//Assign max fees
	err = gas.AssignMaxFeeAndLimit(canNodeDepositResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to deposit?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Make deposit
	response, err := staderClient.NodeDeposit(validatorPubKeyList, true)
	if err != nil {
		return err
	}

	fmt.Printf("Creating %s validators...\n", validatorPubKeyList)
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	fmt.Println("Your validators are now in Initialized status.")
	fmt.Println("Once the ETH deposits have been matched by the remaining 28ETH, it will move to Deposited status.")
	fmt.Println("You can check the status of your validator with `stader-permissioned-cli node status`.")

	return nil

}
