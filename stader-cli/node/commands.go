/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.4.0-beta-permissioned]

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
	"fmt"

	"github.com/urfave/cli"

	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Manage the node",
		Subcommands: []cli.Command{
			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Get the node's status",
				UsageText: "stader-permissioned-cli node status",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getStatus(c)

				},
			},
			{
				Name:      "exit-validator",
				Aliases:   []string{"e"},
				Usage:     "Exit validator",
				UsageText: "stader-permissioned-cli node exit-validator --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of validator we want to exit",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm validator exit",
					},
				},
				Action: func(c *cli.Context) error {

					//// Validate args
					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}

					// Run
					return ExitValidator(c, validatorPubKey)
				},
			},
			{
				Name:      "register",
				Aliases:   []string{"r"},
				Usage:     "Register the node with Stader",
				UsageText: "stader-permissioned-cli node register [options]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "operator-name, on",
						Usage: "The name of the operator",
					},
					cli.StringFlag{
						Name:  "operator-reward-address, ora",
						Usage: "The address at which operator will get rewards (will default to the current node address)",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm node registration",
					},
				},
				Action: func(c *cli.Context) error {
					// Validate flags
					if c.String("operator-name") == "" {
						return fmt.Errorf("operator-name is required")
					}

					if c.String("operator-reward-address") != "" {
						if _, err := cliutils.ValidateAddress("operator-reward-address", c.String("operator-reward-address")); err != nil {
							return err
						}
					}

					// Run
					return registerNode(c)

				},
			},
			{
				Name:      "deposit",
				Aliases:   []string{"d"},
				Usage:     "Make a deposit and create a validator",
				UsageText: "stader-permissioned-cli node deposit [options]",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm deposit",
					},
					cli.StringFlag{
						Name:  "validator-pub-key-list, vpk",
						Usage: "A list of validator public keys to deposit for.",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate flags
					if c.String("amount") != "" {
						if _, err := cliutils.ValidateDepositEthAmount("deposit amount", c.String("amount")); err != nil {
							return err
						}
					}
					if c.String("salt") != "" {
						if _, err := cliutils.ValidateBigInt("salt", c.String("salt")); err != nil {
							return err
						}
					}
					validatorPubKeyList := c.String("validator-pub-key-list")

					// Run
					return nodeDeposit(c, validatorPubKeyList)

				},
			},
			{
				Name:      "send",
				Aliases:   []string{"n"},
				Usage:     "Send ETH or SD, EthX tokens from the node account to an address.",
				UsageText: "stader-permissioned-cli node send [options] amount token to",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm token send",
					},
				},
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 3); err != nil {
						return err
					}
					amount, err := cliutils.ValidatePositiveEthAmount("send amount", c.Args().Get(0))
					if err != nil {
						return err
					}
					token, err := cliutils.ValidateTokenType("token type", c.Args().Get(1))
					if err != nil {
						return err
					}

					// Run
					return nodeSend(c, amount, token, c.Args().Get(2))

				},
			},
			{
				Name:      "get-contracts-info",
				Aliases:   []string{"c"},
				Usage:     "Get the current network contracts info",
				UsageText: "stader-permissioned-cli node get-contracts-info",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return getContractsInfo(c)
				},
			},
			{
				Name:      "claim-cl-rewards",
				Aliases:   []string{"ccr"},
				Usage:     "claim all Consensus Layer rewards to the node reward address.",
				UsageText: "stader-permissioned-cli node claim-cl-rewards --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of validator we want to exit",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm CL rewards withdrawal",
					},
				},
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}
					// Run
					return ClaimClRewards(c, validatorPubKey)
				},
			},
			{
				Name:      "get-validator-info",
				Aliases:   []string{"gvi"},
				Usage:     "Get info for a specific validator",
				UsageText: "stader-permissioned-cli node get-validator-info --validator-pub-key",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "validator-pub-key, vpk",
						Usage: "Public key of validator we want to exit",
					},
				},
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.String("validator-pub-key"))
					if err != nil {
						return err
					}
					// Run
					return GetValidatorInfo(c, validatorPubKey)
				},
			},
			{
				Name:      "claim-sp-rewards",
				Aliases:   []string{"cspr"},
				Usage:     "Claim Socializing Pool Rewards for given cycles",
				UsageText: "stader-permissioned-cli node claim-sp-rewards --download-merkles-proofs --yes",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm claim of rewards",
					},
				},
				Action: func(c *cli.Context) error {

					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Run
					return ClaimSpRewards(c)
				},
			},
			{
				Name:      "update-operator-name",
				Aliases:   []string{"uon"},
				Usage:     "Update Operator name",
				UsageText: "stader-permissioned-cli node update-operator-name --operator-name",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "operator-name, on",
						Usage: "The new operator name",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm claim of rewards",
					},
				},
				Action: func(c *cli.Context) error {

					operatorName := c.String("operator-name")
					if operatorName == "" {
						return fmt.Errorf("operator name can't be empty string")
					}

					// Run
					return updateOperatorName(c, operatorName)
				},
			},
			{
				Name:      "update-operator-reward-address",
				Aliases:   []string{"uod"},
				Usage:     "Update Operator reward address",
				UsageText: "stader-permissioned-cli node update-operator-reward-address --operator-reward-address",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "operator-reward-address, ora",
						Usage: "New operator reward address",
					},
					cli.BoolFlag{
						Name:  "yes, y",
						Usage: "Automatically confirm claim of rewards",
					},
				},
				Action: func(c *cli.Context) error {

					operatorRewardAddress, err := cliutils.ValidateAddress("operator-reward-address", c.String("operator-reward-address"))
					if err != nil {
						return err
					}

					// Run
					return updateOperatorRewardAddress(c, operatorRewardAddress)
				},
			},
		},
	})
}
