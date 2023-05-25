package services

import (
	"math/big"

	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/services/config"
	"github.com/stader-labs/stader-node/shared/services/web3signer"
	web3signer_client "github.com/stader-labs/stader-node/shared/services/web3signer/web3signer-client"
)

type Web3SignerManager struct {
	Web3SignerEndpoint string
	Web3signerClient   web3signer.Client
}

func NewWeb3SignerManager(cfg *config.StaderConfig) (*Web3SignerManager, error) {
	endPoint := cfg.ExternalWeb3Signer.HttpUrl.Value.(string)
	allowAnyCa := cfg.ExternalWeb3Signer.AllowAnyCa.Value.(bool)

	client := web3signer_client.NewStandardHttpClient(endPoint, allowAnyCa)
	return &Web3SignerManager{
		Web3SignerEndpoint: endPoint,
		Web3signerClient:   client,
	}, nil
}

func (w *Web3SignerManager) GetValidatorPubKeysList() ([]string, error) {
	validatorKeys, err := w.Web3signerClient.GetValidatorPubKeys()
	if err != nil {
		return nil, err
	}

	return validatorKeys, nil
}

func (w *Web3SignerManager) GetDepositDataSignature(pubKey string, withdrawCredentials string, amount *big.Int, eth2Config beacon.Eth2Config) (string, error) {
	return w.Web3signerClient.GetDepositDataSignature(pubKey, withdrawCredentials, amount, eth2Config)
}

func (w *Web3SignerManager) GetVoluntaryExitSignature(pubKey string, validatorIndex uint64, epoch uint64, forkInfo beacon.ForkInfo, eth2Config beacon.Eth2Config) (string, error) {
	return w.Web3signerClient.GetVoluntaryExitMessageSignature(pubKey, validatorIndex, epoch, forkInfo, eth2Config)
}

func (w *Web3SignerManager) ReloadKeys() error {
	return w.Web3signerClient.ReloadKeys()
}

func (w *Web3SignerManager) HealthCheck() error {
	return w.Web3signerClient.HealthCheck()
}
