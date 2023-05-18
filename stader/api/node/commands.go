/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package node

import (
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/utils/api"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the node",
		Subcommands: []cli.Command{

			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node's status",
				UsageText: "stader-permissioned-cli api node status",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getStatus(c))
					return nil

				},
			},

			{
				Name:      "can-register",
				Usage:     "Check whether the node can be registered with Stader",
				UsageText: "stader-permissioned-cli api node can-register",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canRegisterNode(c, operatorName, operatorRewardAddress))
					return nil

				},
			},
			{
				Name:      "register",
				Aliases:   []string{"r"},
				Usage:     "Register the node with Stader",
				UsageText: "stader-permissioned-cli api node register operator-name operator-reward-address",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(registerNode(c, operatorName, operatorRewardAddress))
					return nil

				},
			},

			{
				Name:      "can-deposit",
				Usage:     "Check whether the node can make a deposit",
				UsageText: "stader-permissioned-cli api node can-deposit amount validator-list",
				Action: func(c *cli.Context) error {

					//// Validate args
					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}

					// string of public keys seperated by commas
					validatorList := c.Args().Get(0)

					api.PrintResponse(canNodeDeposit(c, validatorList))

					return nil

				},
			},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposit and create a validator, or just make and sign the transaction (when submit = false)",
				UsageText: "stader api node deposit amount validator-list submit",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}

					validatorList := c.Args().Get(0)

					submit, err := cliutils.ValidateBool("submit", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					response, err := nodeDeposit(c, validatorList, submit)
					if submit {
						api.PrintResponse(response, err)
					}
					return nil

				},
			},
			{
				Name:      "can-exit-validator",
				Usage:     "Can validator exit",
				UsageText: "stader-permissioned-cli api node can-exit-validator validator-pub-key",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(canExitValidator(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "exit-validator",
				Usage:     "Exit validator",
				UsageText: "stader-permissioned-cli api node exit-validator validator-pub-key",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(exitValidator(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "can-send",
				Usage:     "Check whether the node can send ETH or tokens to an address",
				UsageText: "stader-permissioned-cli api node can-send amount token",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 2); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(canNodeSend(c, amountWei, token))
					return nil

				},
			},
			{
				Name:      "send",
				Aliases:   []string{"n"},
				Usage:     "Send ETH or tokens from the node account to an address",
				UsageText: "stader-permissioned-cli api node send amount token to",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					amountWei, err := cliutils.ValidatePositiveWeiAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}
					toAddress, err := cliutils.ValidateAddress("to address", c.Args().Get(2))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(nodeSend(c, amountWei, token, toAddress))
					return nil

				},
			},

			{
				Name:      "get-contracts-info",
				Usage:     "Get information about the deposit contract and stader contract on the current network",
				UsageText: "stader-permissioned-cli api node get-contracts-info",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(getContractsInfo(c))
					return nil

				},
			},
			{
				Name:      "can-claim-cl-rewards",
				Usage:     "Can claim cl rewards of a validator",
				UsageText: "stader-permissioned-cli api node can-claim-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(CanClaimClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "claim-cl-rewards",
				Usage:     "Claim cl rewards of a validator",
				UsageText: "stader-permissioned-cli api node claim-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(ClaimClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "can-download-sp-merkle-proofs",
				Usage:     "Can we Download missing socializing merkle proofs",
				UsageText: "stader-permissioned-cli api node can-download-sp-merkle-proofs",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(canDownloadSpMerkleProofs(c))
					return nil

				},
			},
			{
				Name:      "download-sp-merkle-proofs",
				Usage:     "Download missing socializing merkle proofs",
				UsageText: "stader-permissioned-cli api node download-sp-merkle-proofs",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(downloadSpMerkleProofs(c))
					return nil

				},
			},
			{
				Name:      "detailed-cycles-info",
				Usage:     "Get detailed reward cycles info",
				UsageText: "stader-permissioned-cli api node detailed-cycles-info cycles",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					cycles := c.Args().Get(0)

					// Run
					api.PrintResponse(GetCyclesDetailedInfo(c, cycles))
					return nil

				},
			},
			{
				Name:      "can-claim-sp-rewards",
				Usage:     "Can we claim the SP rewards",
				UsageText: "stader-permissioned-cli api node can-claim-sp-rewards",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					api.PrintResponse(canClaimSpRewards(c))
					return nil

				},
			},
			{
				Name:      "claim-sp-rewards",
				Usage:     "Claim the SP rewards",
				UsageText: "stader-permissioned-cli api node claim-sp-rewards cycles",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					cycles := c.Args().Get(0)
					//fmt.Printf("cycles is %s\n", cycles)
					// Run
					api.PrintResponse(claimSpRewards(c, cycles))
					return nil

				},
			},
			{
				Name:      "estimate-claim-sp-rewards-gas",
				Usage:     "Estimate the gas required to claim the SP rewards",
				UsageText: "stader-permissioned-cli api node estimate-claim-sp-rewards-gas cycles",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					cycles := c.Args().Get(0)
					// Run
					api.PrintResponse(estimateSpRewardsGas(c, cycles))
					return nil

				},
			},
			{
				Name:      "can-update-operator-name",
				Usage:     "Can we update the operator name",
				UsageText: "stader-permissioned-cli api node can-update-operator-name operator-name",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					// Run
					api.PrintResponse(CanUpdateOperatorName(c, operatorName))
					return nil

				},
			},
			{
				Name:      "update-operator-name",
				Usage:     "Update the operator name",
				UsageText: "stader-permissioned-cli api node update-operator-name operator-name",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					operatorName := c.Args().Get(0)

					// Run
					api.PrintResponse(UpdateOperatorName(c, operatorName))
					return nil

				},
			},
			{
				Name:      "can-update-operator-reward-address",
				Usage:     "Can we update the operator reward address",
				UsageText: "stader-permissioned-cli api node can-update-operator-reward-address operator-name",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(CanUpdateOperatorRewardAddress(c, operatorRewardAddress))
					return nil

				},
			},
			{
				Name:      "update-operator-reward-address",
				Usage:     "Update the operator reward address",
				UsageText: "stader-permissioned-cli api node update-operator-reward-address operator-name",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.Args().Get(0))
					if err != nil {
						return err
					}

					// Run
					api.PrintResponse(UpdateOperatorRewardAddress(c, operatorRewardAddress))
					return nil

				},
			},
		},
	})
}
