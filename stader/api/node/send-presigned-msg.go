package node

import (
	"encoding/hex"
	"fmt"
	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	"math/big"
	"strconv"
)

func canSendPresignedMsg(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.CanSendPresignedMsgResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	pnr, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	canSendPresignedMsgResponse := api.CanSendPresignedMsgResponse{}

	// check if validator keystore is present by querying validator index
	_, err = w.GetValidatorKeyByPubkey(validatorPubKey)
	if err != nil {
		return nil, err
	}

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Int64() == 0 {
		canSendPresignedMsgResponse.OperatorNotRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	validatorId, err := node.GetValidatorIdByPubKey(pnr, validatorPubKey.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	if validatorId.Cmp(big.NewInt(0)) == 0 {
		canSendPresignedMsgResponse.ValidatorNotRegisteredWithStader = true
		return &canSendPresignedMsgResponse, nil
	}

	// check if validator is present by querying validator index
	validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		return nil, err
	}
	if !validatorStatus.Exists {
		canSendPresignedMsgResponse.ValidatorNotRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	if eth2.IsValidatorExiting(validatorStatus) {
		canSendPresignedMsgResponse.ValidatorIsNotActive = true
		return &canSendPresignedMsgResponse, nil
	}

	// check if already registered
	isRegistered, err := stader.IsPresignedKeyRegistered(validatorPubKey)
	if err != nil {
		return nil, err
	}
	if isRegistered {
		canSendPresignedMsgResponse.ValidatorPreSignKeyAlreadyRegistered = true
		return &canSendPresignedMsgResponse, nil
	}

	return &canSendPresignedMsgResponse, nil
}

func sendPresignedMsg(c *cli.Context, validatorPubKey types.ValidatorPubkey) (*api.SendPresignedMsgResponse, error) {
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	w3Signer, err := services.GetWeb3SignerClient(c)
	if err != nil {
		return nil, err
	}

	currentHead, err := bc.GetBeaconHead()
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

	response := api.SendPresignedMsgResponse{}

	validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
	if err != nil {
		return nil, err
	}

	exitEpoch := currentHead.Epoch

	exitSignature, err := w3Signer.GetVoluntaryExitSignature(validatorPubKey.String(), validatorStatus.Index, exitEpoch, forkInfo, eth2Config)
	if err != nil {
		return nil, err
	}

	decodedExitSignature, err := hex.DecodeString(exitSignature[2:])
	if err != nil {
		return nil, err
	}
	//
	//signatureDomain, err := bc.GetDomainData(eth2types.DomainVoluntaryExit[:], exitEpoch, false)
	//if err != nil {
	//	return nil, err
	//}
	//
	//exitMsg, _, err := validator.GetSignedExitMessage(validatorPrivateKey, validatorStatus.Index, exitEpoch, signatureDomain)
	//if err != nil {
	//	return nil, err
	//}

	// get the public key
	publicKey, err := stader.GetPublicKey()
	if err != nil {
		return nil, err
	}

	exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey(decodedExitSignature, publicKey)
	if err != nil {
		return nil, err
	}
	exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)

	// encrypt the presigned exit message object
	preSignedMessageRequest := stader_backend.PreSignSendApiRequestType{
		Message: struct {
			Epoch          string `json:"epoch"`
			ValidatorIndex string `json:"validator_index"`
		}{
			Epoch:          strconv.FormatUint(exitEpoch, 10),
			ValidatorIndex: strconv.FormatUint(validatorStatus.Index, 10),
		},
		Signature:          exitSignatureEncryptedString,
		ValidatorPublicKey: validatorPubKey.String(),
	}

	res, err := stader.SendPresignedMessageToStaderBackend(preSignedMessageRequest)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, fmt.Errorf("send-presigned-message failed: %s\n", res.Message)
	}

	response.ValidatorIndex = validatorStatus.Index
	response.ValidatorPubKey = validatorPubKey
	response.ExitEpoch = exitEpoch
	response.SignedMsg = exitSignatureEncryptedString

	return &response, nil
}
