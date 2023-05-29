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
package validator

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
				Name:      "can-register",
				Usage:     "Check whether the node can register the validators",
				UsageText: "stader-permissioned-cli api validator can-register validator-list",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					// string of public keys seperated by commas
					validatorList := c.Args().Get(0)

					api.PrintResponse(canRegisterValidators(c, validatorList))

					return nil

				},
			},
			{
				Name:      "register",
				Aliases:   []string{"d"},
				Usage:     "Register the validators with stader",
				UsageText: "stader api validator register validator-list",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 1); err != nil {
						return err
					}

					validatorList := c.Args().Get(0)

					// Run
					response, err := registerValidators(c, validatorList)
					api.PrintResponse(response, err)

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
				Name:      "can-send-cl-rewards",
				Usage:     "Can send cl rewards of a validator to operator reward collector",
				UsageText: "stader-permissioned-cli api node can-send-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(CanSendClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "send-cl-rewards",
				Usage:     "send cl rewards of a validator to operator reward collector",
				UsageText: "stader-permissioned-cli api node send-cl-rewards --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(SendClRewards(c, validatorPubKey))
					return nil

				},
			},
			{
				Name:      "get-validator-info",
				Usage:     "Get info of a particular validator",
				UsageText: "stader-permissioned-cli api node get-validator-info --validator-pub-key",
				Action: func(c *cli.Context) error {

					validatorPubKey, err := cliutils.ValidatePubkey("validator-pub-key", c.Args().Get(0))
					if err != nil {
						return err
					}

					api.PrintResponse(GetValidatorContractInfo(c, validatorPubKey))
					return nil

				},
			},
		},
	})
}
