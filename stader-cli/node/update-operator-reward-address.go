package node

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/gas"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfTypes "github.com/stader-labs/stader-node/shared/types/config"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

const colorReset string = "\033[0m"
const colorRed string = "\033[31m"
const colorGreen string = "\033[32m"
const colorYellow string = "\033[33m"
const colorLightBlue string = "\033[36m"

func updateOperatorRewardAddress(c *cli.Context, operatorRewardAddress common.Address) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// check if we can update the el
	res, err := staderClient.CanUpdateOperatorRewardAddress(operatorRewardAddress)
	if err != nil {
		return err
	}
	if res.OperatorNotActive {
		fmt.Println("Operator not active")
		return nil
	}
	if res.OperatorRewardAddressZero {
		fmt.Println("Operator reward address cannot be zero")
		return nil
	}
	if res.NothingToUpdate {
		fmt.Println("Nothing to update")
		return nil
	}

	cfg, err := services.GetConfig(c)
	if err != nil {
		return err
	}

	if res.OperatorAddressAndRewardNotTheSame {
		promptHowToChangeReward(cfg)
		return nil
	}

	err = gas.AssignMaxFeeAndLimit(res.GasInfo, staderClient, c.Bool("yes"))
	if err != nil {
		return err
	}

	confirmMessage := `
This action will change your Reward Address. Once it's changed, all future SD and ETH rewards will be sent to the New Reward Address.

After you propose the change, your New Reward Address will initially be in a 'Confirmation pending' state until you confirm the change using your New Reward Address on the PermissionedNodeRegistry Smart Contract. Please make sure that your New Reward Address is linked to a web3-compatible wallet, such as MetaMask, to connect with the Smart Contract

Do you wish to proceed with the Reward Address change?`

	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"\n%s %s %s", colorLightBlue, confirmMessage, colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	// update the socializing pool el
	response, err := staderClient.UpdateOperatorRewardAddress(operatorRewardAddress)
	if err != nil {
		return err
	}

	fmt.Println("Updating operator reward address...")

	cliutils.PrintTransactionHash(staderClient, response.TxHash)
	_, err = staderClient.WaitForTransaction(response.TxHash)
	if err != nil {
		return err
	}

	promptSuccessChangedRewardAndNextStep(cfg)

	return nil
}

func promptHowToChangeReward(cfg *config.StaderConfig) {
	network := cfg.StaderNode.Network.Value.(cfTypes.Network)
	switch network {
	case cfTypes.Network_Mainnet:
		msg := `
For node security, only your existing Reward Address can propose a change. To propose and confirm a Reward Address change, please use the PermissionlessNodeRegistry Smart Contract: https://etherscan.io/address/0xaf42d795A6D279e9DCc19DC0eE1cE3ecd4ecf5dD#writeProxyContract
Follow these steps for your Reward Address change:
Step 1: Propose the Reward Address change by connecting your Existing Reward Address wallet with the Smart Contract and execute the "ProposeRewardAddress" function.
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function
Please refer to the Reward Address change guide here - https://staderlabs.notion.site/staderlabs/Stader-ETHx-Reward-address-change-flow-Mainnet-Permissioned-3f2f83405cee47ca8d158a3409b90b31
Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your New Reward Address.`
		fmt.Printf("%s %s %s\n", colorLightBlue, msg, colorReset)
	case cfTypes.Network_Prater:
		msg := `
For node security, only your existing Reward Address can propose a change. To propose and confirm a Reward Address change, please use the PermissionlessNodeRegistry Smart Contract: https://goerli.etherscan.io/address/0xA8BbaD2C6d3F2a28EdB85Fc1c87B300EAC77A00d#writeProxyContract
Follow these steps for your Reward Address change:
Step 1: Propose the Reward Address change by connecting your Existing Reward Address wallet with the Smart Contract and execute the "ProposeRewardAddress" function.
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function
Please refer to the Reward Address change guide here - https://staderlabs.notion.site/staderlabs/Stader-ETHx-Reward-address-change-flow-Testnet-Permissioned-e8418ad8ac7f41ccae107f2996e85a96
Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your New Reward Address.`
		fmt.Printf("%s %s %s\n", colorLightBlue, msg, colorReset)
	default:
		fmt.Println("Unsupported network")
	}
}

func promptSuccessChangedRewardAndNextStep(cfg *config.StaderConfig) {
	switch cfg.StaderNode.Network.Value.(cfTypes.Network) {
	case cfTypes.Network_Mainnet:
		msg := `
You have successfully raised a request to change your Reward Address.

To confirm the Reward Address change please follow these steps:
Step 1: Visit the PermissionedNodeRegistry Smart Contract: https://etherscan.io/address/0xaf42d795A6D279e9DCc19DC0eE1cE3ecd4ecf5dD#writeProxyContract
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function.

Please refer to the Reward Address change guide here - https://staderlabs.notion.site/staderlabs/Stader-ETHx-Reward-address-change-flow-Mainnet-Permissioned-3f2f83405cee47ca8d158a3409b90b31

Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your new Reward Address.`
		fmt.Printf("%s %s %s\n", colorLightBlue, msg, colorReset)
	case cfTypes.Network_Prater:
		msg := `
You have successfully raised a request to change your Reward Address.

To confirm the Reward Address change please follow these steps:
Step 1: Visit the PermissionedNodeRegistry Smart Contract: https://goerli.etherscan.io/address/0xA8BbaD2C6d3F2a28EdB85Fc1c87B300EAC77A00d#writeProxyContract
Step 2: Confirm the Reward Address change by connecting your New Reward Address wallet with the Smart Contract and execute the "ConfirmRewardAddressChange" function.

Please refer to the Reward Address change guide here - https://staderlabs.notion.site/staderlabs/Stader-ETHx-Reward-address-change-flow-Testnet-Permissioned-e8418ad8ac7f41ccae107f2996e85a96

Note: Stader will continue to send rewards to your existing Reward Address until you confirm the change using your new Reward Address.`
		fmt.Printf("%s %s %s\n", colorLightBlue, msg, colorReset)
	default:
		fmt.Println("Unsupported network")
	}
}
