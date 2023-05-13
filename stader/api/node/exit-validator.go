package node

import (
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	"math/big"
)

func canExitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanExitValidatorResponse, error) {

	// Get services
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	// Response
	response := api.CanExitValidatorResponse{}

	// check if validator is registered with stader
	validatorId, err := node.GetValidatorIdByPubKey(prn, validatorPubKey.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	if validatorId.Cmp(big.NewInt(0)) == 0 {
		response.ValidatorNotRegistered = true
		return &response, nil
	}

	res, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		return nil, err
	}
	if !res.Exists {
		response.ValidatorNotExists = true
		return &response, nil
	}

	if eth2.IsValidatorExiting(res) {
		response.ValidatorExiting = true
		return &response, nil
	}

	beaconHead, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}
	currentEpoch := beaconHead.Epoch

	if res.ActivationEpoch+256 > currentEpoch {
		response.ValidatorTooYoung = true
		return &response, nil
	}

	return &response, nil
}

func exitValidator(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.ExitValidatorResponse, error) {

	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	w3signer, err := services.GetWeb3SignerClient(c)
	if err != nil {
		return nil, err
	}

	// Response
	response := api.ExitValidatorResponse{}

	// Get beacon head
	head, err := bc.GetBeaconHead()
	if err != nil {
		return nil, err
	}

	// Get validator index
	validatorIndex, err := bc.GetValidatorIndex(validatorPubKey)
	if err != nil {
		return nil, err
	}

	forkInfo, err := bc.GetForkInfo()
	if err != nil {
		return nil, err
	}
	eth2Config, err := bc.GetEth2Config()
	if err != nil {
		return nil, err
	}

	hexExitSignature, err := w3signer.GetVoluntaryExitSignature(validatorPubKey.Hex(), validatorIndex, head.Epoch, forkInfo, eth2Config)
	if err != nil {
		return nil, err
	}
	exitSignature, err := types.HexToValidatorSignature(hexExitSignature)
	if err != nil {
		return nil, err
	}

	// Broadcast voluntary exit message
	if err := bc.ExitValidator(validatorIndex, head.Epoch, exitSignature); err != nil {
		return nil, err
	}

	// Return response
	return &response, nil

}
