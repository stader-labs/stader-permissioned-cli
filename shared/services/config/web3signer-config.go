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

// Constants
const web3SignerTag string = "consensys/web3signer:develop"

// Defaults
const defaultWeb3SignerPort uint16 = 9000

// Configuration for Web3Signer
type Web3SignerConfig struct {
	Title string `yaml:"-"`

	// The HTTP port to serve on
	Port config.Parameter `yaml:"port,omitempty"`

	// The Docker Hub tag for Grafana
	ContainerTag config.Parameter `yaml:"containerTag,omitempty"`
}

// Generates a new Grafana config
func NewWeb3SignerConfig(cfg *StaderConfig) *Web3SignerConfig {
	return &Web3SignerConfig{
		Title: "Web3Signer Settings",

		Port: config.Parameter{
			ID:                   "port",
			Name:                 "Web3Signer Port",
			Description:          "The port Grafana should run its HTTP server on - this is the port you will connect to in your browser.",
			Type:                 config.ParameterType_Uint16,
			Default:              map[config.Network]interface{}{config.Network_All: defaultWeb3SignerPort},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Web3Signer},
			EnvironmentVariables: []string{"WEB3SIGNER_HTTP_LISTEN_PORT"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},

		ContainerTag: config.Parameter{
			ID:                   "containerTag",
			Name:                 "Web3Signer Container Tag",
			Description:          "The tag name of the Web3Signer container you want to use on Docker Hub.",
			Type:                 config.ParameterType_String,
			Default:              map[config.Network]interface{}{config.Network_All: web3SignerTag},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Web3Signer},
			EnvironmentVariables: []string{"WEB3SIGNER_CONTAINER_TAG"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   true,
		},
	}
}

// Get the parameters for this config
func (cfg *Web3SignerConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.Port,
		&cfg.ContainerTag,
	}
}

// The the title for the config
func (cfg *Web3SignerConfig) GetConfigTitle() string {
	return cfg.Title
}

func (cfg *Web3SignerConfig) GetContainerTag() string {
	return web3SignerTag
}

func (cfg *Web3SignerConfig) GetHttpListenPort() uint16 {
	return defaultWeb3SignerPort
}
