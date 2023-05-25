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

	AllowAnyCa config.Parameter `yaml:"allowAnyCa,omitempty"`
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

		AllowAnyCa: config.Parameter{
			ID:                   "AllowAnyCa",
			Name:                 "Allow Any CA",
			Description:          "Configure whether you want to trust the CA of the web3signer. If the web3signer you are connected to uses a self-signed certificate, you will need to set this to true.",
			Type:                 config.ParameterType_Bool,
			Default:              map[config.Network]interface{}{config.Network_All: true},
			AffectsContainers:    []config.ContainerID{config.ContainerID_Api},
			EnvironmentVariables: []string{"W3_SIGNER_TRUST_CA"},
			CanBeBlank:           false,
			OverwriteOnUpgrade:   false,
		},
	}
}

func (cfg *ExternalWeb3SignerConfig) GetParameters() []*config.Parameter {
	return []*config.Parameter{
		&cfg.HttpUrl,
		&cfg.AllowAnyCa,
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
