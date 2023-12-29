package stader

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/stader-labs/stader-node/shared/services"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/net"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
)

func SendPresignedMessageToStaderBackend(c *cli.Context, preSignedMessage stader_backend.PreSignSendApiRequestType) (*stader_backend.PreSignSendApiResponseType, error) {
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	res, err := net.MakePostRequest(config.StaderNode.GetPresignSendApi(), preSignedMessage)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var preSignSendResponse stader_backend.PreSignSendApiResponseType
	err = json.NewDecoder(res.Body).Decode(&preSignSendResponse)
	if err != nil {
		return nil, err
	}

	return &preSignSendResponse, nil
}

func SendBulkPresignedMessageToStaderBackend(c *cli.Context, preSignedMessages []stader_backend.PreSignSendApiRequestType) (*map[string]stader_backend.PreSignSendApiResponseType, error) {
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	res, err := net.MakePostRequest(config.StaderNode.GetBulkPresignSendApi(), preSignedMessages)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var preSignSendResponse map[string]stader_backend.PreSignSendApiResponseType
	err = json.NewDecoder(res.Body).Decode(&preSignSendResponse)
	if err != nil {
		return nil, err
	}

	return &preSignSendResponse, nil
}

func IsPresignedKeyRegistered(c *cli.Context, validatorPubKey types.ValidatorPubkey) (bool, error) {
	config, err := services.GetConfig(c)
	if err != nil {
		return false, err
	}

	// check if it is already there
	preSignCheckRequest := stader_backend.PreSignCheckApiRequestType{
		ValidatorPublicKey: validatorPubKey.String(),
	}

	res, err := net.MakePostRequest(config.StaderNode.GetPresignCheckApi(), preSignCheckRequest)

	defer res.Body.Close()
	var preSignCheckResponse stader_backend.PreSignCheckApiResponseType
	err = json.NewDecoder(res.Body).Decode(&preSignCheckResponse)
	if err != nil {
		return false, err
	}
	return preSignCheckResponse.Value, nil
}

func BulkIsPresignedKeyRegistered(c *cli.Context, validatorPubKeys []types.ValidatorPubkey) (map[string]bool, error) {
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	res, err := net.MakePostRequest(config.StaderNode.GetBulkPresignCheckApi(), stader_backend.BulkPreSignCheckApiRequestType{ValidatorPubKeys: validatorPubKeys})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	//fmt.Printf("res.body is %s\n", res)
	var preSignCheckResponse map[string]bool
	err = json.NewDecoder(res.Body).Decode(&preSignCheckResponse)
	if err != nil {
		return nil, err
	}
	return preSignCheckResponse, nil
}

func GetPublicKey(c *cli.Context) (*rsa.PublicKey, error) {
	config, err := services.GetConfig(c)
	if err != nil {
		return nil, err
	}

	decodedPublicKey, err := crypto.DecodeBase64(config.StaderNode.GetPresignEncryptionKey())
	if err != nil {
		return nil, err
	}

	publicKey, err := crypto.BytesToPublicKey(decodedPublicKey)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
