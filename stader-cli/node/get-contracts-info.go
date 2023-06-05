package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/shared/services/stader"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/urfave/cli"
)

func getContractsInfo(c *cli.Context) error {
	staderClient, err := stader.NewClientFromCtx(c)
	if err != nil {
		return err
	}
	defer staderClient.Close()

	// Get node fee
	response, err := staderClient.GetContractsInfo()
	if err != nil {
		return err
	}

	fmt.Printf("%s=== Beacon Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Beacon Network: %d\n\n", response.BeaconNetwork)
	fmt.Printf("Beacon Deposit Contract: %s\n\n", response.BeaconDepositContract)

	fmt.Printf("%s=== Stader Network Contract Details ===%s\n", log.ColorGreen, log.ColorReset)
	fmt.Printf("Network: %d\n\n", response.Network)
	fmt.Printf("Permissioned Node Registry: %s\n\n", response.PermissionedNodeRegistry)
	fmt.Printf("Vault Factory: %s\n\n", response.VaultFactory)
	fmt.Printf("EthX Token: %s\n\n", response.EthxToken)
	fmt.Printf("Sd Token: %s\n\n", response.SdToken)
	fmt.Printf("Socializing Pool: %s\n\n", response.SocializingPoolContract)
	fmt.Printf("Permissioned Pool: %s\n\n", response.PermissionedPool)
	fmt.Printf("Stader Oracle: %s\n\n", response.StaderOracle)
	fmt.Printf("Stader Config: %s\n\n", response.StaderConfig)
	fmt.Printf("Stake Pool Manager: %s\n\n", response.StakePoolManager)

	return nil
}
