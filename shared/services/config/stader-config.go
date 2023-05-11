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
package config

import (
	"fmt"
	"github.com/alessio/shellescape"
	"github.com/pbnjay/memory"
	"github.com/stader-labs/stader-node/shared"
	"github.com/stader-labs/stader-node/shared/types/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

// Constants
const (
	rootConfigName string = "root"

	ApiContainerName          string = "api"
	Eth1ContainerName         string = "eth1"
	Eth1FallbackContainerName string = "eth1-fallback"
	Eth2ContainerName         string = "eth2"
	ExporterContainerName     string = "exporter"
	GrafanaContainerName      string = "grafana"
	MevBoostContainerName     string = "mev-boost"
	NodeContainerName         string = "node"
	PrometheusContainerName   string = "prometheus"
	ValidatorContainerName    string = "validator"
	GuardianContainerName     string = "guardian"
	Web3SignerContainerName   string = "web3signer"

	FeeRecipientFileEnvVar string = "FEE_RECIPIENT_FILE"
	FeeRecipientEnvVar     string = "FEE_RECIPIENT"
)

// Defaults
const defaultBnMetricsPort uint16 = 9100
const defaultVcMetricsPort uint16 = 9101
const defaultNodeMetricsPort uint16 = 9104
const defaultExporterMetricsPort uint16 = 9103
const defaultEcMetricsPort uint16 = 9105

// The master configuration struct
type StaderConfig struct {
	Title string `yaml:"-"`

	Version string `yaml:"-"`

	StaderDirectory string `yaml:"-"`

	// The StaderNode configuration
	StaderNode *StaderNodeConfig `yaml:"stadernode,omitempty"`

	// Execution client configurations
	ExternalExecution *ExternalExecutionConfig `yaml:"externalExecution,omitempty"`
	ExternalBeacon    *ExternalBeaconConfig    `yaml:"externalBeacon,omitempty"`

	// Web3 Signer
	Web3SignerMode     config.Parameter          `yaml:"web3SignerMode,omitempty"`
	Web3Signer         *Web3SignerConfig         `yaml:"web3Signer,omitempty"`
	ExternalWeb3Signer *ExternalWeb3SignerConfig `yaml:"externalWeb3Signer,omitempty"`
}

// Load configuration settings from a file
func LoadFromFile(path string) (*StaderConfig, error) {

	// Return nil if the file doesn't exist
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, nil
	}

	// Read the file
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read Stader settings file at %s: %w", shellescape.Quote(path), err)
	}

	// Attempt to parse it out into a settings map
	var settings map[string]map[string]string
	if err := yaml.Unmarshal(configBytes, &settings); err != nil {
		return nil, fmt.Errorf("could not parse settings file: %w", err)
	}

	// Deserialize it into a config object
	cfg := NewStaderConfig(filepath.Dir(path))
	err = cfg.Deserialize(settings)
	if err != nil {
		return nil, fmt.Errorf("could not deserialize settings file: %w", err)
	}

	return cfg, nil

}

// Creates a new Stader configuration instance
func NewStaderConfig(staderDir string) *StaderConfig {

	clientModes := []config.ParameterOption{{
		Name:        "Locally Managed",
		Description: "Allow the Stadernode to manage the Execution and Consensus clients for you (Docker Mode)",
		Value:       config.Mode_Local,
	}, {
		Name:        "Externally Managed",
		Description: "Use existing Execution and Consensus clients that you manage on your own (Hybrid Mode)",
		Value:       config.Mode_External,
	}}

	cfg := &StaderConfig{
		Title:           "Top-level Settings",
		StaderDirectory: staderDir,

		Web3SignerMode: config.Parameter{
			ID:                   "web3SignerMode",
			Name:                 "Web3 Signer Mode",
			Description:          "Choose which mode to use for your Web3 Signer - locally managed (Docker Mode), or externally managed",
			Type:                 config.ParameterType_Choice,
			Default:              map[config.Network]interface{}{config.Network_All: config.Mode_Local},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api},
			EnvironmentVariables: []string{},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
			Options:              clientModes,
		},
	}

	cfg.StaderNode = NewStadernodeConfig(cfg)
	cfg.ExternalExecution = NewExternalExecutionConfig(cfg)
	cfg.ExternalBeacon = NewExternalBeaconConfig(cfg)
	cfg.Web3Signer = NewWeb3SignerConfig(cfg)
	cfg.ExternalWeb3Signer = NewExternalWeb3Signer(cfg)

	// Apply the default values for mainnet
	cfg.StaderNode.Network.Value = cfg.StaderNode.Network.Options[0].Value
	cfg.applyAllDefaults()

	return cfg
}

// Get a more verbose client description, including warnings
func getAugmentedEcDescription(client config.ExecutionClient, originalDescription string) string {

	switch client {
	case config.ExecutionClient_Nethermind:
		totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024
		if totalMemoryGB < 9 {
			return fmt.Sprintf("%s\n\n[red]WARNING: Nethermind currently requires over 8 GB of RAM to run smoothly. We do not recommend it for your system. This may be improved in a future release.", originalDescription)
		}
	}

	return originalDescription

}

// Create a copy of this configuration.
func (cfg *StaderConfig) CreateCopy() *StaderConfig {
	newConfig := NewStaderConfig(cfg.StaderDirectory)

	// Set the network
	network := cfg.StaderNode.Network.Value.(config.Network)
	newConfig.StaderNode.Network.Value = network

	newParams := newConfig.GetParameters()
	for i, param := range cfg.GetParameters() {
		newParams[i].Value = param.Value
		newParams[i].UpdateDescription(network)
	}

	newSubconfigs := newConfig.GetSubconfigs()
	for name, subConfig := range cfg.GetSubconfigs() {
		newParams := newSubconfigs[name].GetParameters()
		for i, param := range subConfig.GetParameters() {
			newParams[i].Value = param.Value
			newParams[i].UpdateDescription(network)
		}
	}

	return newConfig
}

// Get the parameters for this config
func (cfg *StaderConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Web3SignerMode,
	}
}

// Get the subconfigurations for this config
func (cfg *StaderConfig) GetSubconfigs() map[string]config.Config {
	return map[string]config.Config{
		"stadernode":         cfg.StaderNode,
		"externalExecution":  cfg.ExternalExecution,
		"externalBeacon":     cfg.ExternalBeacon,
		"web3signer":         cfg.Web3Signer,
		"externalWeb3Signer": cfg.ExternalWeb3Signer,
	}
}

// Handle a network change on all of the parameters
func (cfg *StaderConfig) ChangeNetwork(newNetwork config.Network) {

	// Get the current network
	oldNetwork, ok := cfg.StaderNode.Network.Value.(config.Network)
	if !ok {
		oldNetwork = config.Network_Unknown
	}
	if oldNetwork == newNetwork {
		return
	}
	cfg.StaderNode.Network.Value = newNetwork

	// Update the master parameters
	rootParams := cfg.GetParameters()
	for _, param := range rootParams {
		param.ChangeNetwork(oldNetwork, newNetwork)
	}

	// Update all of the child config objects
	subconfigs := cfg.GetSubconfigs()
	for _, subconfig := range subconfigs {
		for _, param := range subconfig.GetParameters() {
			param.ChangeNetwork(oldNetwork, newNetwork)
		}
	}

}

// Serializes the configuration into a map of maps, compatible with a settings file
func (cfg *StaderConfig) Serialize() map[string]map[string]string {

	masterMap := map[string]map[string]string{}

	// Serialize root params
	rootParams := map[string]string{}
	for _, param := range cfg.GetParameters() {
		param.Serialize(rootParams)
	}
	masterMap[rootConfigName] = rootParams
	masterMap[rootConfigName]["sdDir"] = cfg.StaderDirectory
	masterMap[rootConfigName]["version"] = fmt.Sprintf("v%s", shared.StaderVersion) // Update the version with the current Stadernode version

	// Serialize the subconfigs
	for name, subconfig := range cfg.GetSubconfigs() {
		subconfigParams := map[string]string{}
		for _, param := range subconfig.GetParameters() {
			param.Serialize(subconfigParams)
		}
		masterMap[name] = subconfigParams
	}

	return masterMap
}

// Deserializes a settings file into this config
func (cfg *StaderConfig) Deserialize(masterMap map[string]map[string]string) error {

	var err error
	// Get the network
	network := config.Network_Mainnet
	stadernodeConfig, exists := masterMap["stadernode"]
	if exists {
		networkString, exists := stadernodeConfig[cfg.StaderNode.Network.ID]
		if exists {
			valueType := reflect.TypeOf(networkString)
			paramType := reflect.TypeOf(network)
			if !valueType.ConvertibleTo(paramType) {
				return fmt.Errorf("can't get default network: value type %s cannot be converted to parameter type %s", valueType.Name(), paramType.Name())
			}
			network = reflect.ValueOf(networkString).Convert(paramType).Interface().(config.Network)
		}
	}

	// Deserialize root params
	rootParams := masterMap[rootConfigName]
	for _, param := range cfg.GetParameters() {
		// Note: if the root config doesn't exist, this will end up using the default values for all of its settings
		err := param.Deserialize(rootParams, network)
		if err != nil {
			return fmt.Errorf("error deserializing root config: %w", err)
		}
	}

	cfg.StaderDirectory = masterMap[rootConfigName]["sdDir"]
	if err != nil {
		return fmt.Errorf("error parsing isNative: %w", err)
	}
	cfg.Version = masterMap[rootConfigName]["version"]

	// Deserialize the subconfigs
	for name, subconfig := range cfg.GetSubconfigs() {
		subconfigParams := masterMap[name]
		for _, param := range subconfig.GetParameters() {
			// Note: if the subconfig doesn't exist, this will end up using the default values for all of its settings
			err := param.Deserialize(subconfigParams, network)
			if err != nil {
				return fmt.Errorf("error deserializing [%s]: %w", name, err)
			}
		}
	}

	return nil
}

// Generates a collection of environment variables based on this config's settings
func (cfg *StaderConfig) GenerateEnvironmentVariables() map[string]string {

	envVars := map[string]string{}

	// Basic variables and root parameters
	envVars["STADER_NODE_IMAGE"] = cfg.StaderNode.GetStadernodeContainerTag()
	envVars["STADER_FOLDER"] = cfg.StaderDirectory
	envVars["ETHX_ADDRESS"] = cfg.StaderNode.GetEthxTokenAddress().Hex()
	envVars[FeeRecipientFileEnvVar] = FeeRecipientFilename // If this is running, we're in Docker mode by definition so use the Docker fee recipient filename
	config.AddParametersToEnvVars(cfg.StaderNode.GetParameters(), envVars)
	config.AddParametersToEnvVars(cfg.GetParameters(), envVars)

	envVars["EC_CLIENT"] = "X" // X is for external / unknown
	config.AddParametersToEnvVars(cfg.ExternalExecution.GetParameters(), envVars)

	// Get the hostname of the Execution client, necessary for Prometheus to work in hybrid mode
	ecUrl, err := url.Parse(envVars["EC_HTTP_ENDPOINT"])
	if err == nil && ecUrl != nil {
		envVars["EC_HOSTNAME"] = ecUrl.Hostname()
	}

	versionString := fmt.Sprintf("v%s", shared.StaderVersion)
	envVars["STADER_VERSION"] = versionString

	// Get the hostname of the Consensus client, necessary for Prometheus to work in hybrid mode
	ccUrl, err := url.Parse(envVars["CC_API_ENDPOINT"])
	if err == nil && ccUrl != nil {
		envVars["CC_HOSTNAME"] = ccUrl.Hostname()
	}

	// Web3 Signer parameters
	// TODO - bchain - read these params from the config
	envVars["WEB3SIGNER_ETH2_KEYSTORES_PASSWORDS_PATH"] = cfg.StaderNode.GetWeb3SignerKeyStoresPasswordPath(true)
	envVars["WEB3SIGNER_ETH2_KEYSTORES_PATH"] = cfg.StaderNode.GetWeb3SignerKeyStorePath(true)
	envVars["WEB3SIGNER_CONTAINER_TAG"] = cfg.Web3Signer.GetContainerTag()
	envVars["WEB3SIGNER_HTTP_LISTEN_PORT"] = strconv.Itoa(int(cfg.Web3Signer.GetHttpListenPort()))
	//config.AddParametersToEnvVars(cfg.Web3Signer.GetParameters(), envVars)

	return envVars

}

// The the title for the config
func (cfg *StaderConfig) GetConfigTitle() string {
	return cfg.Title
}

// Update the default settings for all overwrite-on-upgrade parameters
func (cfg *StaderConfig) UpdateDefaults() error {
	// Update the root params
	currentNetwork := cfg.StaderNode.Network.Value.(config.Network)
	for _, param := range cfg.GetParameters() {
		defaultValue, err := param.GetDefault(currentNetwork)
		if err != nil {
			return fmt.Errorf("error getting defaults for root param [%s] on network [%v]: %w", param.ID, currentNetwork, err)
		}
		//if param.OverwriteOnUpgrade {
		param.Value = defaultValue
		//}
	}

	// Update the subconfigs
	for subconfigName, subconfig := range cfg.GetSubconfigs() {
		for _, param := range subconfig.GetParameters() {
			defaultValue, err := param.GetDefault(currentNetwork)
			if err != nil {
				return fmt.Errorf("error getting defaults for %s param [%s] on network [%v]: %w", subconfigName, param.ID, currentNetwork, err)
			}
			//if param.OverwriteOnUpgrade {
			param.Value = defaultValue
			//}
		}
	}

	return nil
}

// Get all of the settings that have changed between an old config and this config, and get all of the containers that are affected by those changes - also returns whether or not the selected network was changed
func (cfg *StaderConfig) GetChanges(oldConfig *StaderConfig) (map[string][]config.ChangedSetting, map[config.ContainerID]bool, bool) {
	// Get the map of changed settings by category
	changedSettings := getChangedSettingsMap(oldConfig, cfg)

	// Create a list of all of the container IDs that need to be restarted
	totalAffectedContainers := map[config.ContainerID]bool{}
	for _, settingList := range changedSettings {
		for _, setting := range settingList {
			for container := range setting.AffectedContainers {
				totalAffectedContainers[container] = true
			}
		}
	}

	// Check if the network has changed
	changeNetworks := false
	if oldConfig.StaderNode.Network.Value != cfg.StaderNode.Network.Value {
		changeNetworks = true
	}

	// Return everything
	return changedSettings, totalAffectedContainers, changeNetworks
}

// Checks to see if the current configuration is valid; if not, returns a list of errors
func (cfg *StaderConfig) Validate() []string {
	errors := []string{}

	return errors
}

// Applies all of the defaults to all of the settings that have them defined
func (cfg *StaderConfig) applyAllDefaults() error {
	for _, param := range cfg.GetParameters() {
		err := param.SetToDefault(cfg.StaderNode.Network.Value.(config.Network))
		if err != nil {
			return fmt.Errorf("error setting root parameter default: %w", err)
		}
	}

	for name, subconfig := range cfg.GetSubconfigs() {
		for _, param := range subconfig.GetParameters() {
			err := param.SetToDefault(cfg.StaderNode.Network.Value.(config.Network))
			if err != nil {
				return fmt.Errorf("error setting parameter default for %s: %w", name, err)
			}
		}
	}

	return nil
}

// Get all of the changed settings between an old and new config
func getChangedSettingsMap(oldConfig *StaderConfig, newConfig *StaderConfig) map[string][]config.ChangedSetting {
	changedSettings := map[string][]config.ChangedSetting{}

	// Root settings
	oldRootParams := oldConfig.GetParameters()
	newRootParams := newConfig.GetParameters()
	changedSettings[oldConfig.Title] = getChangedSettings(oldRootParams, newRootParams, newConfig)

	// Subconfig settings
	oldSubconfigs := oldConfig.GetSubconfigs()
	for name, subConfig := range newConfig.GetSubconfigs() {
		oldParams := oldSubconfigs[name].GetParameters()
		newParams := subConfig.GetParameters()
		changedSettings[subConfig.GetConfigTitle()] = getChangedSettings(oldParams, newParams, newConfig)
	}

	return changedSettings
}

// Get all of the settings that have changed between the given parameter lists.
// Assumes the parameter lists represent identical parameters (e.g. they have the same number of elements and
// each element has the same ID).
func getChangedSettings(oldParams []*config.Parameter, newParams []*config.Parameter, newConfig *StaderConfig) []config.ChangedSetting {
	changedSettings := []config.ChangedSetting{}

	for i, param := range newParams {
		oldValString := fmt.Sprint(oldParams[i].Value)
		newValString := fmt.Sprint(param.Value)
		if oldValString != newValString {
			changedSettings = append(changedSettings, config.ChangedSetting{
				Name:               param.Name,
				OldValue:           oldValString,
				NewValue:           newValString,
				AffectedContainers: getAffectedContainers(param, newConfig),
			})
		}
	}

	return changedSettings
}

// Handles custom container overrides
func getAffectedContainers(param *config.Parameter, cfg *StaderConfig) map[config.ContainerID]bool {

	affectedContainers := map[config.ContainerID]bool{}

	for _, container := range param.AffectsContainers {
		affectedContainers[container] = true
	}

	return affectedContainers

}
