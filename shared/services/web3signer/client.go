package web3signer

import (
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"math/big"
)

type Client interface {
	GetValidatorPubKeys() ([]string, error)
	GetDepositDataSignature(pubKey string, withdrawCredentials string, amount *big.Int, eth2Config beacon.Eth2Config) (string, error)
	GetVoluntaryExitMessageSignature(pubKey string, validatorIndex uint64, epoch uint64, ForkInfo beacon.ForkInfo, eth2Config beacon.Eth2Config) (string, error)
	ReloadKeys() error
	HealthCheck() error
}
