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
	"github.com/stader-labs/stader-node/shared/types/config"
)

// Configuration for external Execution clients
type ExternalExecutionConfig struct {
	Title string `yaml:"-"`

	// The URL of the HTTP endpoint
	HttpUrl config.Parameter `yaml:"httpUrl,omitempty"`

	// The URL of the websocket endpoint
	WsUrl config.Parameter `yaml:"wsUrl,omitempty"`
}

// Configuration for external Execution clients
type ExternalBeaconConfig struct {
	Title string `yaml:"-"`

	// The URL of the HTTP endpoint
	HttpUrl config.Parameter `yaml:"httpUrl,omitempty"`
}

// Configuration for external web3signer clients
type ExternalWeb3SignerConfig struct {
	Title string `yaml:"-"`

	HttpUrl config.Parameter `yaml:"httpUrl,omitempty"`
}

func (cfg *ExternalWeb3SignerConfig) GetConfigTitle() string {
	return "externalWeb3Signer"
}

// Generates a new ExternalExecutionConfig configuration
func NewExternalExecutionConfig(cfg *StaderConfig) *ExternalExecutionConfig {
	return &ExternalExecutionConfig{
		Title: "External Execution Client Settings",

		HttpUrl: config.Parameter{
			ID:                   "httpUrl",
			Name:                 "HTTP URL",
			Description:          "The URL of the HTTP RPC endpoint for your external Execution client.\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api},
			EnvironmentVariables: []string{"EC_HTTP_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		WsUrl: config.Parameter{
			ID:                   "wsUrl",
			Name:                 "Websocket URL",
			Description:          "The URL of the Websocket RPC endpoint for your %s Execution client.\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api},
			EnvironmentVariables: []string{"EC_WS_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

// Generates a new ExternalExecutionConfig configuration
func NewExternalBeaconConfig(cfg *StaderConfig) *ExternalBeaconConfig {
	return &ExternalBeaconConfig{
		Title: "External Beacon Client Settings",

		HttpUrl: config.Parameter{
			ID:                   "httpUrl",
			Name:                 "HTTP URL",
			Description:          "The URL of the HTTP RPC endpoint for your external Beacon client.\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api},
			EnvironmentVariables: []string{"BC_HTTP_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

func NewExternalWeb3Signer(cfg *StaderConfig) *ExternalWeb3SignerConfig {
	return &ExternalWeb3SignerConfig{
		Title: "External Web3Signer Client Settings",

		HttpUrl: config.Parameter{
			ID:                   "httpUrl",
			Name:                 "HTTP URL",
			Description:          "The URL of the HTTP RPC endpoint for your external Web3Signer client.\nNOTE: If you are running it on the same machine as the Stadernode, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: ""},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api},
			EnvironmentVariables: []string{"W3_SIGNER_HTTP_ENDPOINT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

func (cfg *ExternalWeb3SignerConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.HttpUrl,
	}
}

// Get the parameters for this config
func (cfg *ExternalExecutionConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.HttpUrl,
		&cfg.WsUrl,
	}
}

func (cfg *ExternalBeaconConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.HttpUrl,
	}
}

func (cfg *ExternalWeb3SignerConfig) GetName() string {
	return "Web3Signer"
}

func (cfg *ExternalBeaconConfig) GetName() string {
	return "Beacon"
}

// The the title for the config
func (cfg *ExternalExecutionConfig) GetConfigTitle() string {
	return cfg.Title
}

func (cfg *ExternalBeaconConfig) GetConfigTitle() string {
	return cfg.Title
}
