/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.0]

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
				Name:      "send",
				Aliases:   []string{"n"},
				Usage:     "Send ETH,SD or ETHx tokens from the node account to an address.",
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
				Name:      "claim-rewards",
				Aliases:   []string{"wer"},
				Usage:     "Claim rewards from claim vault to the claim vault",
				UsageText: "stader-permissioned-cli node claim-rewards",
				Flags: []cli.Flag{cli.BoolFlag{
					Name:  "yes, y",
					Usage: "Automatically confirm rewards transfer to claim vault",
				}},
				Action: func(c *cli.Context) error {
					// Run
					return ClaimRewards(c)
				},
			},
			{
				Name:      "claim-sp-rewards",
				Aliases:   []string{"cspr"},
				Usage:     "Claim Socializing Pool Rewards for given cycles",
				UsageText: "stader-permissioned-cli node claim-sp-rewards",
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
