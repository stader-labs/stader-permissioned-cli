package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	"math/big"
)

func GetValidatorInfo(c *cli.Context, validatorPubKey types.ValidatorPubkey) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	response, err := staderClient.GetValidatorInfo(validatorPubKey)
	if err != nil {
		return err
	}

	fmt.Printf("Validator info for %s\n", validatorPubKey.String())
	fmt.Printf("Validator PubKey: %d\n", response.ValidatorInfo.Pubkey)
	fmt.Printf("Validator WithdrawVault: %s\n", response.ValidatorInfo.WithdrawVaultAddress.String())
	fmt.Printf("Validator Status: %s\n", response.DisplayStatus)
	if response.ValidatorInfo.DepositBlock.Cmp(big.NewInt(0)) != 0 {
		fmt.Printf("Validator DepositBlock: %d\n", response.ValidatorInfo.DepositBlock)
	}
	if response.ValidatorInfo.WithdrawnBlock.Cmp(big.NewInt(0)) != 0 {
		fmt.Printf("Validator WithdrawnBlock: %d\n", response.ValidatorInfo.WithdrawnBlock)
	}

	return nil
}
