package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/math"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/stader-labs/stader-node/stader-lib/utils/eth"
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
	fmt.Printf("Validator PubKey: %s\n", types.BytesToValidatorPubkey(response.ValidatorInfo.Pubkey).String())
	fmt.Printf("Validator WithdrawVault: %s\n", response.ValidatorInfo.WithdrawVaultAddress.String())
	fmt.Printf("Validator Status: %s\n", response.DisplayStatus)
	if response.ClRewards.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Validator Skimmed Rewards: %.6f\n", math.RoundDown(eth.WeiToEth(response.ClRewards), 18))
		fmt.Printf("To claim skimmed rewards use the %sstader-permissioned-cli node claim-cl-rewards %s command\n\n", log.ColorGreen, log.ColorReset)
	}
	if response.ValidatorInfo.DepositBlock.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Validator DepositBlock: %d\n", response.ValidatorInfo.DepositBlock)
	}
	if response.ValidatorInfo.WithdrawnBlock.Cmp(big.NewInt(0)) > 0 {
		fmt.Printf("Validator WithdrawnBlock: %d\n", response.ValidatorInfo.WithdrawnBlock)
	}

	return nil
}
