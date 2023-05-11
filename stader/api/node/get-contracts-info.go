package node

import (
	"fmt"
	"github.com/urfave/cli"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
)

func getContractsInfo(c *cli.Context) (*api.ContractsInfoResponse, error) {
	// Response
	response := api.ContractsInfoResponse{}

	// Get the ETH1 network ID
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, fmt.Errorf("Error getting configuration: %w", err)
	}
	response.Network = uint64(config.StaderNode.GetChainID())

	// Get the Beacon Client info
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, fmt.Errorf("Error getting beacon client: %w", err)
	}
	//eth2Config, err := bc.GetEth2Config()
	//if err != nil {
	//	return nil, err
	//}
	//forkInfo, err := bc.GetForkInfo()
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("eth2Config is: %v\n", eth2Config)
	//fmt.Printf("Genesis Fork is %s\n", common.BytesToHash(eth2Config.GenesisForkVersion))
	eth2DepositContract, err := bc.GetEth2DepositContract()
	if err != nil {
		return nil, fmt.Errorf("Error getting beacon deposit contract: %w", err)
	}
	//web3signer, err := services.GetWeb3SignerClient(c)
	//if err != nil {
	//	return nil, fmt.Errorf("Error getting web3signer: %w", err)
	//}

	//pubKey := "b1129ac24bf99c446ecb8027f29745c41fd2849e2885170d8f1f96d5b17dab346a4d8db98f22a0c667f015e506a30ced"
	////withdrawCreds := "00e6cc0e49ac7d19b6967ad5b8da1f6f9a74d4cc713f6086a4c7ca8394b49b3a"
	////amount := big.NewInt(32000000000)
	////signature, err := web3signer.GetDepositDataSignature(pubKey, withdrawCreds, amount, eth2Config)
	////if err != nil {
	////	return nil, err
	////}
	////fmt.Printf("deposit signature is: %v", signature)
	//signature, err := web3signer.GetVoluntaryExitSignature(pubKey, 467162, 169030, forkInfo, eth2Config)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("voluntary exit signature is: %v\n", signature)
	//validatorPubKeys, err := web3signer.GetValidatorPubKeysList()
	//if err != nil {
	//	fmt.Printf("failed with error: %v", err)
	//	return nil, err
	//}

	response.BeaconNetwork = eth2DepositContract.ChainID
	response.BeaconDepositContract = eth2DepositContract.Address
	response.SdCollateralContract, err = services.GetSdCollateralAddress(c)
	if err != nil {
		return nil, err
	}
	response.EthxToken, err = services.GetEthxTokenAddress(c)
	if err != nil {
		return nil, err
	}
	response.SdToken, err = services.GetSdTokenAddress(c)
	if err != nil {
		return nil, err
	}
	response.PermissionedNodeRegistry, err = services.GetPermissionedNodeRegistryAddress(c)
	if err != nil {
		return nil, err
	}
	response.VaultFactory, err = services.GetVaultFactoryAddress(c)
	if err != nil {
		return nil, err
	}
	response.SocializingPoolContract, err = services.GetSocializingPoolAddress(c)
	if err != nil {
		return nil, err
	}
	response.PermissionedPool, err = services.GetPermissionedPoolAddress(c)
	if err != nil {
		return nil, err
	}
	response.StaderOracle, err = services.GetStaderOracleAddress(c)
	if err != nil {
		return nil, err
	}
	response.StaderConfig = config.StaderNode.GetStaderConfigAddress()
	response.StakePoolManager, err = services.GetStakePoolManagerAddress(c)
	if err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
