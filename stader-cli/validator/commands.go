package validator

import (
	"fmt"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/urfave/cli"
)

// Register commands
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Validator related commands",
		Subcommands: []cli.Command{
			{
				Name:      "register",
				Aliases:   []string{"d"},
				Usage:     "Register validators with stader",
				UsageText: "stader-permissioned-cli validator register [options]",
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

					validatorPubKeyList := c.String("validator-pub-key-list")
					if validatorPubKeyList == "" {
						return fmt.Errorf("validator-pub-key-list is required")
					}

					// Run
					return registerValidators(c, validatorPubKeyList)

				},
			},
			{
				Name:      "send-cl-rewards",
				Aliases:   []string{"ccr"},
				Usage:     "Send all Consensus Layer rewards to the claim vault.",
				UsageText: "stader-permissioned-cli validator send-cl-rewards --validator-pub-key",
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
					return SendClRewards(c, validatorPubKey)
				},
			},
			{
				Name:      "exit-validator",
				Aliases:   []string{"e"},
				Usage:     "Exit validator",
				UsageText: "stader-permissioned-cli validator exit-validator --validator-pub-key",
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
				Name:      "get-validator-info",
				Aliases:   []string{"gvi"},
				Usage:     "Get info for a specific validator",
				UsageText: "stader-permissioned-cli validator get-validator-info --validator-pub-key",
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
				Name:      "export",
				Aliases:   []string{"e"},
				Usage:     "Export validator info to a csv file",
				UsageText: "stader-permissioned-cli validator export",
				Flags:     []cli.Flag{},
				Action: func(c *cli.Context) error {

					// Run
					return exportValidatorStatus(c)
				},
			},
			{
				Name:      "status",
				Aliases:   []string{"s"},
				Usage:     "Validator Status",
				UsageText: "stader-permissioned-cli validator status",
				Flags:     []cli.Flag{},
				Action: func(c *cli.Context) error {

					// Run
					return getValidatorStatus(c)
				},
			},
		},
	})
}
