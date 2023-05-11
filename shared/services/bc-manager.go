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
package services

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/beacon/client"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/types"
)

// This is a proxy for multiple Beacon clients, providing natural fallback support if one of them fails.
type BeaconClientManager struct {
	primaryBc       beacon.Client
	fallbackBc      beacon.Client
	logger          log.ColorLogger
	primaryReady    bool
	fallbackReady   bool
	ignoreSyncCheck bool
}

// This is a signature for a wrapped Beacon client function that only returns an error
type bcFunction0 func(beacon.Client) error

// This is a signature for a wrapped Beacon client function that returns 1 var and an error
type bcFunction1 func(beacon.Client) (interface{}, error)

// This is a signature for a wrapped Beacon client function that returns 2 vars and an error
type bcFunction2 func(beacon.Client) (interface{}, interface{}, error)

// Creates a new BeaconClientManager instance based on the Stader config
func NewBeaconClientManager(cfg *config.StaderConfig) (*BeaconClientManager, error) {

	// TODO - bchain - allow users to pass this is as an option
	primaryProvider := "https://beacon-nd-942-489-268.p2pify.com/c450ba1e6c5025d33dd14dc4c54f5cf6"
	fallbackProvider := "https://beacon-nd-942-489-268.p2pify.com/c450ba1e6c5025d33dd14dc4c54f5cf6"

	//primaryProvider := cfg.ExternalBeacon.HttpUrl.Value.(string)
	//fallbackProvider := cfg.ExternalBeacon.HttpUrl.Value.(string)

	fmt.Printf("primaryBc: %s\n", primaryProvider)
	fmt.Printf("fallbackBc: %s\n", fallbackProvider)

	primaryBc := client.NewStandardHttpClient(primaryProvider)
	fallbackBc := client.NewStandardHttpClient(fallbackProvider)

	return &BeaconClientManager{
		primaryBc:     primaryBc,
		fallbackBc:    fallbackBc,
		logger:        log.NewColorLogger(color.FgHiBlue),
		primaryReady:  true,
		fallbackReady: fallbackBc != nil,
	}, nil

}

/// ======================
/// BeaconClient Functions
/// ======================

// Get the client's process mode
func (m *BeaconClientManager) GetClientType() (beacon.BeaconClientType, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetClientType()
	})
	if err != nil {
		return beacon.Unknown, err
	}
	return result.(beacon.BeaconClientType), nil
}

// Get the client's sync status
func (m *BeaconClientManager) GetSyncStatus() (beacon.SyncStatus, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetSyncStatus()
	})
	if err != nil {
		return beacon.SyncStatus{}, err
	}
	return result.(beacon.SyncStatus), nil
}

// Get the Beacon configuration
func (m *BeaconClientManager) GetEth2Config() (beacon.Eth2Config, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetEth2Config()
	})
	if err != nil {
		return beacon.Eth2Config{}, err
	}
	return result.(beacon.Eth2Config), nil
}

// Get the Beacon configuration
func (m *BeaconClientManager) GetEth2DepositContract() (beacon.Eth2DepositContract, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetEth2DepositContract()
	})
	if err != nil {
		return beacon.Eth2DepositContract{}, err
	}
	return result.(beacon.Eth2DepositContract), nil
}

// Get the attestations in a Beacon chain block
func (m *BeaconClientManager) GetAttestations(blockId string) ([]beacon.AttestationInfo, bool, error) {
	result1, result2, err := m.runFunction2(func(client beacon.Client) (interface{}, interface{}, error) {
		return client.GetAttestations(blockId)
	})
	if err != nil {
		return nil, false, err
	}
	return result1.([]beacon.AttestationInfo), result2.(bool), nil
}

// Get a Beacon chain block
func (m *BeaconClientManager) GetBeaconBlock(blockId string) (beacon.BeaconBlock, bool, error) {
	result1, result2, err := m.runFunction2(func(client beacon.Client) (interface{}, interface{}, error) {
		return client.GetBeaconBlock(blockId)
	})
	if err != nil {
		return beacon.BeaconBlock{}, false, err
	}
	return result1.(beacon.BeaconBlock), result2.(bool), nil
}

// Get the Beacon chain's head information
func (m *BeaconClientManager) GetBeaconHead() (beacon.BeaconHead, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetBeaconHead()
	})
	if err != nil {
		return beacon.BeaconHead{}, err
	}
	return result.(beacon.BeaconHead), nil
}

// Get a validator's status by its index
func (m *BeaconClientManager) GetValidatorStatusByIndex(index string, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetValidatorStatusByIndex(index, opts)
	})
	if err != nil {
		return beacon.ValidatorStatus{}, err
	}
	return result.(beacon.ValidatorStatus), nil
}

// Get a validator's status by its pubkey
func (m *BeaconClientManager) GetValidatorStatus(pubkey types.ValidatorPubkey, opts *beacon.ValidatorStatusOptions) (beacon.ValidatorStatus, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetValidatorStatus(pubkey, opts)
	})
	if err != nil {
		return beacon.ValidatorStatus{}, err
	}
	return result.(beacon.ValidatorStatus), nil
}

// Get the statuses of multiple validators by their pubkeys
func (m *BeaconClientManager) GetValidatorStatuses(pubkeys []types.ValidatorPubkey, opts *beacon.ValidatorStatusOptions) (map[types.ValidatorPubkey]beacon.ValidatorStatus, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetValidatorStatuses(pubkeys, opts)
	})
	if err != nil {
		return nil, err
	}
	return result.(map[types.ValidatorPubkey]beacon.ValidatorStatus), nil
}

// Get a validator's index
func (m *BeaconClientManager) GetValidatorIndex(pubkey types.ValidatorPubkey) (uint64, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetValidatorIndex(pubkey)
	})
	if err != nil {
		return 0, err
	}
	return result.(uint64), nil
}

// Get a validator's sync duties
func (m *BeaconClientManager) GetValidatorSyncDuties(indices []uint64, epoch uint64) (map[uint64]bool, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetValidatorSyncDuties(indices, epoch)
	})
	if err != nil {
		return nil, err
	}
	return result.(map[uint64]bool), nil
}

// Get a validator's proposer duties
func (m *BeaconClientManager) GetValidatorProposerDuties(indices []uint64, epoch uint64) (map[uint64]uint64, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetValidatorProposerDuties(indices, epoch)
	})
	if err != nil {
		return nil, err
	}
	return result.(map[uint64]uint64), nil
}

// Get the Beacon chain's domain data
func (m *BeaconClientManager) GetDomainData(domainType []byte, epoch uint64, useGenesisFork bool) ([]byte, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetDomainData(domainType, epoch, useGenesisFork)
	})
	if err != nil {
		return nil, err
	}
	return result.([]byte), nil
}

// Voluntarily exit a validator
func (m *BeaconClientManager) ExitValidator(validatorIndex, epoch uint64, signature types.ValidatorSignature) error {
	err := m.runFunction0(func(client beacon.Client) error {
		return client.ExitValidator(validatorIndex, epoch, signature)
	})
	return err
}

// Close the connection to the Beacon client
func (m *BeaconClientManager) Close() error {
	err := m.runFunction0(func(client beacon.Client) error {
		return client.Close()
	})
	return err
}

// Get the EL data for a CL block
func (m *BeaconClientManager) GetEth1DataForEth2Block(blockId string) (beacon.Eth1Data, bool, error) {
	result1, result2, err := m.runFunction2(func(client beacon.Client) (interface{}, interface{}, error) {
		return client.GetEth1DataForEth2Block(blockId)
	})
	if err != nil {
		return beacon.Eth1Data{}, false, err
	}
	return result1.(beacon.Eth1Data), result2.(bool), nil
}

// Get the attestation committees for an epoch
func (m *BeaconClientManager) GetCommitteesForEpoch(epoch *uint64) ([]beacon.Committee, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetCommitteesForEpoch(epoch)
	})
	if err != nil {
		return nil, err
	}
	return result.([]beacon.Committee), nil
}

func (m *BeaconClientManager) GetForkInfo() (beacon.ForkInfo, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetForkInfo()
	})
	if err != nil {
		return beacon.ForkInfo{}, err
	}
	return result.(beacon.ForkInfo), nil
}

func (m *BeaconClientManager) GetGenesisInfo() (beacon.GenesisInfo, error) {
	result, err := m.runFunction1(func(client beacon.Client) (interface{}, error) {
		return client.GetGenesisInfo()
	})
	if err != nil {
		return beacon.GenesisInfo{}, err
	}
	return result.(beacon.GenesisInfo), nil
}

/// ==================
/// Internal Functions
/// ==================

func (m *BeaconClientManager) CheckStatus() *api.ClientManagerStatus {

	status := &api.ClientManagerStatus{
		FallbackEnabled: m.fallbackBc != nil,
	}

	// Ignore the sync check and just use the predefined settings if requested
	if m.ignoreSyncCheck {
		status.PrimaryClientStatus.IsWorking = m.primaryReady
		status.PrimaryClientStatus.IsSynced = m.primaryReady
		if status.FallbackEnabled {
			status.FallbackClientStatus.IsWorking = m.fallbackReady
			status.FallbackClientStatus.IsSynced = m.fallbackReady
		}
		return status
	}

	// Get the primary BC status
	status.PrimaryClientStatus = checkBcStatus(m.primaryBc)

	// Get the fallback BC status if applicable
	if status.FallbackEnabled {
		status.FallbackClientStatus = checkBcStatus(m.fallbackBc)
	}

	// Flag the ready clients
	m.primaryReady = (status.PrimaryClientStatus.IsWorking && status.PrimaryClientStatus.IsSynced)
	m.fallbackReady = (status.FallbackEnabled && status.FallbackClientStatus.IsWorking && status.FallbackClientStatus.IsSynced)

	return status

}

// Check the client status
func checkBcStatus(client beacon.Client) api.ClientStatus {

	status := api.ClientStatus{}

	// Get the fallback's sync progress
	syncStatus, err := client.GetSyncStatus()
	if err != nil {
		status.Error = fmt.Sprintf("Sync progress check failed with [%s]", err.Error())
		status.IsSynced = false
		status.IsWorking = false
		return status
	}

	// Return the sync status
	if !syncStatus.Syncing {
		status.IsWorking = true
		status.IsSynced = true
		status.SyncProgress = 1
	} else {
		status.IsWorking = true
		status.IsSynced = false
		status.SyncProgress = syncStatus.Progress
	}
	return status

}

// Attempts to run a function progressively through each client until one succeeds or they all fail.
func (m *BeaconClientManager) runFunction0(function bcFunction0) error {

	// Check if we can use the primary
	if m.primaryReady {
		// Try to run the function on the primary
		err := function(m.primaryBc)
		if err != nil {
			if m.isDisconnected(err) {
				// If it's disconnected, log it and try the fallback
				m.logger.Printlnf("WARNING: Primary Beacon client disconnected (%s), using fallback...", err.Error())
				m.primaryReady = false
				return m.runFunction0(function)
			}
			// If it's a different error, just return it
			return err
		}
		// If there's no error, return the result
		return nil
	}

	if m.fallbackReady {
		// Try to run the function on the fallback
		err := function(m.fallbackBc)
		if err != nil {
			if m.isDisconnected(err) {
				// If it's disconnected, log it and try the fallback
				m.logger.Printlnf("WARNING: Fallback Beacon client disconnected (%s)", err.Error())
				m.fallbackReady = false
				return fmt.Errorf("all Beacon clients failed")
			}

			// If it's a different error, just return it
			return err
		}
		// If there's no error, return the result
		return nil
	}

	return fmt.Errorf("no Beacon clients were ready")
}

// Attempts to run a function progressively through each client until one succeeds or they all fail.
func (m *BeaconClientManager) runFunction1(function bcFunction1) (interface{}, error) {

	// Check if we can use the primary
	if m.primaryReady {
		// Try to run the function on the primary
		result, err := function(m.primaryBc)
		if err != nil {
			if m.isDisconnected(err) {
				// If it's disconnected, log it and try the fallback
				m.logger.Printlnf("WARNING: Primary Beacon client disconnected (%s), using fallback...", err.Error())
				m.primaryReady = false
				return m.runFunction1(function)
			}
			// If it's a different error, just return it
			return nil, err
		}
		// If there's no error, return the result
		return result, nil
	}

	if m.fallbackReady {
		// Try to run the function on the fallback
		result, err := function(m.fallbackBc)
		if err != nil {
			if m.isDisconnected(err) {
				// If it's disconnected, log it and try the fallback
				m.logger.Printlnf("WARNING: Fallback Beacon client disconnected (%s)", err.Error())
				m.fallbackReady = false
				return nil, fmt.Errorf("all Beacon clients failed")
			}
			// If it's a different error, just return it
			return nil, err
		}
		// If there's no error, return the result
		return result, nil
	}

	return nil, fmt.Errorf("no Beacon clients were ready")

}

// Attempts to run a function progressively through each client until one succeeds or they all fail.
func (m *BeaconClientManager) runFunction2(function bcFunction2) (interface{}, interface{}, error) {

	// Check if we can use the primary
	if m.primaryReady {
		// Try to run the function on the primary
		result1, result2, err := function(m.primaryBc)
		if err != nil {
			if m.isDisconnected(err) {
				// If it's disconnected, log it and try the fallback
				m.logger.Printlnf("WARNING: Primary Beacon client disconnected (%s), using fallback...", err.Error())
				m.primaryReady = false
				return m.runFunction2(function)
			}
			// If it's a different error, just return it
			return nil, nil, err
		}
		// If there's no error, return the result
		return result1, result2, nil
	}

	if m.fallbackReady {
		// Try to run the function on the fallback
		result1, result2, err := function(m.fallbackBc)
		if err != nil {
			if m.isDisconnected(err) {
				// If it's disconnected, log it and try the fallback
				m.logger.Printlnf("WARNING: Fallback Beacon client disconnected (%s)", err.Error())
				m.fallbackReady = false
				return nil, nil, fmt.Errorf("all Beacon clients failed")
			}
			// If it's a different error, just return it
			return nil, nil, err
		}
		// If there's no error, return the result
		return result1, result2, nil
	}

	return nil, nil, fmt.Errorf("no Beacon clients were ready")

}

// Returns true if the error was a connection failure and a backup client is available
func (m *BeaconClientManager) isDisconnected(err error) bool {
	return strings.Contains(err.Error(), "dial tcp")
}
