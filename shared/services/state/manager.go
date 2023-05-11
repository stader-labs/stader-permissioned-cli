package state

import (
	"context"
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/urfave/cli"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type NetworkStateManager struct {
	c            *cli.Context
	cfg          *config.StaderConfig
	ec           stader.ExecutionClient
	bc           beacon.Client
	log          *log.ColorLogger
	Config       *config.StaderConfig
	Network      cfgtypes.Network
	ChainID      uint
	BeaconConfig beacon.Eth2Config
}

// Create a new manager for the network state
func NewNetworkStateManager(c *cli.Context, cfg *config.StaderConfig, ec stader.ExecutionClient, bc beacon.Client, log *log.ColorLogger) (*NetworkStateManager, error) {

	// Create the manager
	m := &NetworkStateManager{
		c:       c,
		cfg:     cfg,
		ec:      ec,
		bc:      bc,
		log:     log,
		Config:  cfg,
		Network: cfg.StaderNode.Network.Value.(cfgtypes.Network),
		ChainID: cfg.StaderNode.GetChainID(),
	}

	// Get the Beacon config info
	var err error
	m.BeaconConfig, err = m.bc.GetEth2Config()
	if err != nil {
		return nil, err
	}

	return m, nil

}

// Get the state of the network using the latest Execution layer block
func (m *NetworkStateManager) GetHeadState(nodeAddress common.Address) (*NetworkStateCache, error) {
	targetSlot, err := m.GetHeadSlot()
	if err != nil {
		return nil, fmt.Errorf("error getting latest Beacon slot: %w", err)
	}
	return m.getStateForNode(nodeAddress, targetSlot)
}

// Get the state of the network for a single node using the latest Execution layer block, along with the total effective RPL stake for the network
func (m *NetworkStateManager) GetHeadStateForNode(nodeAddress common.Address) (*NetworkStateCache, error) {
	targetSlot, err := m.GetHeadSlot()
	if err != nil {
		return nil, fmt.Errorf("error getting latest Beacon slot: %w", err)
	}
	return m.getStateForNode(nodeAddress, targetSlot)
}

// Get the state of the network at the provided Beacon slot
func (m *NetworkStateManager) GetStateForSlot(nodeAddress common.Address, slotNumber uint64) (*NetworkStateCache, error) {
	return m.getStateForNode(nodeAddress, slotNumber)
}

// Gets the latest valid block
func (m *NetworkStateManager) GetLatestBeaconBlock() (beacon.BeaconBlock, error) {
	targetSlot, err := m.GetHeadSlot()
	if err != nil {
		return beacon.BeaconBlock{}, fmt.Errorf("error getting head slot: %w", err)
	}
	return m.getLatestProposedBeaconBlock(targetSlot)
}

// Gets the latest valid finalized block
func (m *NetworkStateManager) GetLatestFinalizedBeaconBlock() (beacon.BeaconBlock, error) {
	head, err := m.bc.GetBeaconHead()
	if err != nil {
		return beacon.BeaconBlock{}, fmt.Errorf("error getting Beacon chain head: %w", err)
	}
	targetSlot := head.FinalizedEpoch*m.BeaconConfig.SlotsPerEpoch + (m.BeaconConfig.SlotsPerEpoch - 1)
	return m.getLatestProposedBeaconBlock(targetSlot)
}

// Gets the Beacon slot for the latest execution layer block
func (m *NetworkStateManager) GetHeadSlot() (uint64, error) {
	// Get the latest EL block
	latestBlockHeader, err := m.ec.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, fmt.Errorf("error getting latest EL block: %w", err)
	}

	// Get the corresponding Beacon slot based on the timestamp
	latestBlockTime := time.Unix(int64(latestBlockHeader.Time), 0)
	genesisTime := time.Unix(int64(m.BeaconConfig.GenesisTime), 0)
	secondsSinceGenesis := uint64(latestBlockTime.Sub(genesisTime).Seconds())
	targetSlot := secondsSinceGenesis / m.BeaconConfig.SecondsPerSlot
	return targetSlot, nil
}

// Gets the target Beacon block, or if it was missing, the first one under it that wasn't missing
func (m *NetworkStateManager) getLatestProposedBeaconBlock(targetSlot uint64) (beacon.BeaconBlock, error) {
	for {
		// Try to get the current block
		block, exists, err := m.bc.GetBeaconBlock(fmt.Sprint(targetSlot))
		if err != nil {
			return beacon.BeaconBlock{}, fmt.Errorf("error getting Beacon block %d: %w", targetSlot, err)
		}

		// If the block was missing, try the previous one
		if !exists {
			m.logLine("Slot %d was missing, trying the previous one...", targetSlot)
			targetSlot--
		} else {
			return block, nil
		}
	}
}

// Get the state of the network for a specific node only at the provided Beacon slot
func (m *NetworkStateManager) getStateForNode(nodeAddress common.Address, slotNumber uint64) (*NetworkStateCache, error) {
	state, err := CreateNetworkStateCache(m.c, m.cfg.StaderNode, m.ec, m.bc, m.log, slotNumber, m.BeaconConfig, nodeAddress)
	if err != nil {
		return nil, err
	}
	return state, nil
}

// Logs a line if the logger is specified
func (m *NetworkStateManager) logLine(format string, v ...interface{}) {
	if m.log != nil {
		m.log.Printlnf(format, v)
	}
}
