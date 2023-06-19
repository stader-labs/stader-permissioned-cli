/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
package service

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/stader"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	cliutils "github.com/stader-labs/stader-node/shared/utils/cli"
	"github.com/stader-labs/stader-node/shared/utils/sys"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Settings
const (
	NodeContainerSuffix     string = "_node"
	GuardianContainerSuffix string = "guardian"

	colorReset     string = "\033[0m"
	colorRed       string = "\033[31m"
	colorYellow    string = "\033[33m"
	colorGreen     string = "\033[32m"
	colorLightBlue string = "\033[36m"
)

// Install the Stader service
func installService(c *cli.Context) error {
	dataPath := ""

	if c.String("network") != "" {
		fmt.Printf("%sNOTE: The --network flag is deprecated. You no longer need to specify it.%s\n\n", colorLightBlue, colorReset)
	}

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf(
		"The Stader service will be installed --Version: %s\n\n%sIf you're upgrading, your existing configuration will be backed up and preserved.\nAll of your previous settings will be migrated automatically.%s\nAre you sure you want to continue?",
		c.String("version"), colorGreen, colorReset,
	))) {
		fmt.Println("Cancelled.")
		return nil
	}

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Attempt to load the config to see if any settings need to be passed along to the install script
	cfg, isNew, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading old configuration: %w", err)
	}
	if !isNew {
		dataPath = cfg.StaderNode.DataPath.Value.(string)
		dataPath, err = homedir.Expand(dataPath)
		if err != nil {
			return fmt.Errorf("error getting data path from old configuration: %w", err)
		}
	}
	fmt.Println("Installing Stader service...")

	// Install service
	err = staderClient.InstallService(c.Bool("verbose"), c.Bool("no-deps"), c.String("network"), c.String("version"), c.String("path"), dataPath)
	if err != nil {
		return err
	}

	// Print success message & return
	fmt.Println("")
	fmt.Println("The Stader service was successfully installed!")

	// printPatchNotes(c)

	// Reload the config after installation
	cfg, isNew, err = staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading new configuration: %w", err)
	}

	fmt.Printf("%s\n=== Saved default Config ===\n", colorLightBlue)
	if isNew {
		err = staderClient.SaveConfig(cfg)
		if err != nil {
			return fmt.Errorf("error saveConfig: %w", err)
		}
	}

	// Print the docker permissions notice
	if isNew {
		// Report next steps
		fmt.Printf("%s\n=== Next Steps ===\n", colorLightBlue)
		fmt.Printf("Open %s/user-settings.yml and add the URLs for the beacon chain node, execution chain node and your web3signer setup. %s\n", cfg.StaderDirectory, colorReset)

		fmt.Printf("\n%sNOTE:\nSince this is your first time installing Stader, please start a new shell session by logging out and back in or restarting the machine.\n", colorYellow)
		fmt.Printf("This is necessary for your user account to have permissions to use Docker.%s", colorReset)
	}

	// Remove the upgrade flag if it's there
	return staderClient.RemoveUpgradeFlagFile()
}

// View the Stader service status
func serviceStatus(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(staderClient)
	if err != nil {
		return err
	}

	// Print service status
	return staderClient.PrintServiceStatus(getComposeFiles(c))

}

func ConvertBoolToString(boolValue bool) string {
	if boolValue {
		return "Yes"
	}

	return "No"
}

func ConvertStringToBool(stringValue string) bool {
	if stringValue == "Yes" {
		return true
	} else {
		return false
	}
}

// Updates a configuration from the provided CLI arguments headlessly
func configureHeadless(c *cli.Context, cfg *config.StaderConfig) error {

	// Root params
	for _, param := range cfg.GetParameters() {
		err := updateConfigParamFromCliArg(c, "", param, cfg)
		if err != nil {
			return err
		}
	}

	// Subconfigs
	for sectionName, subconfig := range cfg.GetSubconfigs() {
		for _, param := range subconfig.GetParameters() {
			err := updateConfigParamFromCliArg(c, sectionName, param, cfg)
			if err != nil {
				return err
			}
		}
	}

	return nil

}

// Updates a config parameter from a CLI flag
func updateConfigParamFromCliArg(c *cli.Context, sectionName string, param *cfgtypes.Parameter, cfg *config.StaderConfig) error {

	var paramName string
	if sectionName == "" {
		paramName = param.ID
	} else {
		paramName = fmt.Sprintf("%s-%s", sectionName, param.ID)
	}

	if c.IsSet(paramName) {
		switch param.Type {
		case cfgtypes.ParameterType_Bool:
			param.Value = c.Bool(paramName)
		case cfgtypes.ParameterType_Int:
			param.Value = c.Int(paramName)
		case cfgtypes.ParameterType_Float:
			param.Value = c.Float64(paramName)
		case cfgtypes.ParameterType_String:
			setting := c.String(paramName)
			if param.MaxLength > 0 && len(setting) > param.MaxLength {
				return fmt.Errorf("error setting value for %s: [%s] is too long (max length %d)", paramName, setting, param.MaxLength)
			}
			param.Value = c.String(paramName)
		case cfgtypes.ParameterType_Uint:
			param.Value = c.Uint(paramName)
		case cfgtypes.ParameterType_Uint16:
			param.Value = uint16(c.Uint(paramName))
		case cfgtypes.ParameterType_Choice:
			selection := c.String(paramName)
			found := false
			for _, option := range param.Options {
				if fmt.Sprint(option.Value) == selection {
					param.Value = option.Value
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("error setting value for %s: [%s] is not one of the valid options", paramName, selection)
			}
		}
	}

	return nil

}

// Start the Stader service
func startService(c *cli.Context, ignoreConfigSuggestion bool) error {

	println("Starting Stader Service")

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Update the Prometheus template with the assigned ports
	cfg, _, err := staderClient.LoadConfig()
	if err != nil {
		return fmt.Errorf("Error loading user settings: %w", err)
	}

	if !ignoreConfigSuggestion {
		if c.Bool("yes") || cliutils.Confirm("Stadernode upgrade detected - starting will overwrite certain settings with the latest defaults (such as container versions).\nYou may want to run `service config` first to see what's changed.\n\nWould you like to continue starting the service?") {
			err = cfg.UpdateDefaults()
			if err != nil {
				return fmt.Errorf("error upgrading configuration with the latest parameters: %w", err)
			}
			staderClient.SaveConfig(cfg)
			fmt.Printf("%sUpdated settings successfully.%s\n", colorGreen, colorReset)
		} else {
			fmt.Println("Cancelled.")
			return nil
		}
	}

	// Validate the config
	errors := cfg.Validate()
	if len(errors) > 0 {
		fmt.Printf("%sYour configuration encountered errors. You must correct the following in order to start Stader:\n\n", colorRed)
		for _, err := range errors {
			fmt.Printf("%s\n\n", err)
		}
		fmt.Println(colorReset)
		return nil
	}

	// Start service
	err = staderClient.StartService(getComposeFiles(c))
	if err != nil {
		return err
	}

	return nil

}

// Pause the Stader service
func pauseService(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm("Are you sure you want to pause the Stader service? Any staking validators will be penalized!")) {
		fmt.Println("Cancelled.")
		return nil
	}

	// Pause service
	return staderClient.PauseService(getComposeFiles(c))

}

// Stop the Stader service
func stopService(c *cli.Context) error {

	// Prompt for confirmation
	if !(c.Bool("yes") || cliutils.Confirm(fmt.Sprintf("%sWARNING: Are you sure you want to terminate the Stader service? Any validators will be penalized, your ETH1 and ETH2 chain databases will be deleted, you will lose ALL of your sync progress, and you will lose your Prometheus metrics database!%s", colorRed, colorReset))) {
		fmt.Println("Cancelled.")
		return nil
	}

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Stop service
	return staderClient.StopService(getComposeFiles(c))

}

// View the Stader service logs
func serviceLogs(c *cli.Context, serviceNames ...string) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print service logs
	return staderClient.PrintServiceLogs(getComposeFiles(c), c.String("tail"), serviceNames...)

}

// View the Stader service stats
func serviceStats(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print service stats
	return staderClient.PrintServiceStats(getComposeFiles(c))

}

// View the Stader service compose config
func serviceCompose(c *cli.Context) error {

	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Print service compose config
	return staderClient.PrintServiceCompose(getComposeFiles(c))

}

// View the Stader service version information
func serviceVersion(c *cli.Context) error {

	stader, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer stader.Close()

	// Print what network we're on
	err = cliutils.PrintNetwork(stader)
	if err != nil {
		return err
	}

	serviceVersion, err := stader.GetServiceVersion()
	if err != nil {
		return err
	}

	// Get config
	_, isNew, err := stader.LoadConfig()
	if err != nil {
		return err
	}
	if isNew {
		return fmt.Errorf("settings file not found. Please run `stader-permissioned-cli service config` to set up your Stadernode")
	}

	// Print version info
	fmt.Printf("Stader client version: %s\n", c.App.Version)
	fmt.Printf("Stader service version: %s\n", serviceVersion)
	return nil

}

// Get the compose file paths for a CLI context
func getComposeFiles(c *cli.Context) []string {
	return c.Parent().StringSlice("compose-file")
}

// Generate a YAML file that shows the current configuration schema, including all of the parameters and their descriptions
func getConfigYaml(c *cli.Context) error {
	cfg := config.NewStaderConfig("")
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("error serializing configuration schema: %w", err)
	}

	fmt.Println(string(bytes))
	return nil
}

// Get the list of features required for modern client containers but not supported by the CPU
func checkCpuFeatures() error {
	unsupportedFeatures := sys.GetMissingModernCpuFeatures()
	if len(unsupportedFeatures) > 0 {
		fmt.Println("Your CPU is missing support for the following features:")
		for _, name := range unsupportedFeatures {
			fmt.Printf("  - %s\n", name)
		}

		fmt.Println("\nYou must use the 'portable' image.")
		return nil
	}

	fmt.Println("Your CPU supports all required features for 'modern' images.")
	return nil
}
