package node

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stader-labs/stader-node/stader-lib/stader"
)

func IsPermissionedNodeRegistryPaused(pnr *stader.PermissionedNodeRegistryContractManager, opts *bind.CallOpts) (bool, error) {
	return pnr.PermissionedNodeRegistry.Paused(opts)
}
