package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func SendSignedPresignedMessage(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	canSendPresignedMsgRes, err := staderClient.CanSendPresignedMessage(validatorPubKey)
	if err != nil {
		return err
	}
	if canSendPresignedMsgRes.ValidatorNotRegistered {
		fmt.Println("Validator not registered on beacon chain!")
		return nil
	}
	if canSendPresignedMsgRes.ValidatorPreSignKeyAlreadyRegistered {
		fmt.Println("Validator pre sign key is already registered!")
		return nil
	}
	if canSendPresignedMsgRes.ValidatorIsNotActive {
		fmt.Println("Validator is not active")
		return nil
	}
	if canSendPresignedMsgRes.OperatorNotRegistered {
		fmt.Println("Operator not registered!")
		return nil
	}
	if canSendPresignedMsgRes.ValidatorNotRegisteredWithStader {
		fmt.Println("Validator not registered with stader!")
		return nil
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"Are you sure you want to send presigned message for validator %s?", validatorPubKey))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// send presigned message
	res, err := staderClient.SendPresignedMessage(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Sent the presigned message successfully for validator %s\n", res.ValidatorPubKey)

	return nil
}
