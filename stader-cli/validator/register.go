package validator

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/gas"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services/stader"
)

func registerValidators(c *cli.Context, validatorPubKeyList string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	canRegisterValidatorsResponse, err := staderClient.CanRegisterValidators(validatorPubKeyList)
	if err != nil {
		return err
	}
	if canRegisterValidatorsResponse.InsufficientBalance {
		fmt.Printf("Account does not have enough balance!")
		return nil
	}
	if canRegisterValidatorsResponse.DepositPaused {
		fmt.Printf("Deposits are currently paused!")
		return nil
	}
	if canRegisterValidatorsResponse.MaxValidatorLimitReached {
		fmt.Printf("Max validator limit reached")
		return nil
	}
	if canRegisterValidatorsResponse.OperatorNotRegistered {
		fmt.Printf("Operator not registered")
		return nil
	}
	if canRegisterValidatorsResponse.OperatorNotActive {
		fmt.Printf("Operator not active")
		return nil
	}
	if canRegisterValidatorsResponse.InputKeyLimitReached {
		fmt.Printf("You cannot add more than %d keys at a time", canRegisterValidatorsResponse.InputKeyLimit)
		return nil
	}

	//Assign max fees
	err = gas.AssignMaxFeeAndLimit(canRegisterValidatorsResponse.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to register the validators with stader?"))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Make deposit
	response, err := staderClient.RegisterValidators(validatorPubKeyList)
	if err != nil {
		return err
	}

	fmt.Printf("Registering validators with stader...\n")
	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	fmt.Println("Your validators are now in Initialized status.")
	fmt.Println("Your validators will be matched with a 1Eth deposit by stader which will be followed by the remaining 31Eth deposit")
	fmt.Println("You can check the status of your validator with `stader-permissioned-cli validator status`.")

	return nil

}
