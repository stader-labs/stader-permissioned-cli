package web3signer_client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/services/beacon"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
)

const (
	RequestUrlFormat      = "%s%s"
	RequestContentType    = "application/json"
	PublicKeyListEndpoint = "/api/v1/eth2/publicKeys"
	SignatureEndpoint     = "/api/v1/eth2/sign/%s" // %s = validator public key
	Reload                = "/reload"
	HealthCheck           = "/healthcheck"
)

// TODO - bchain - abstract this out to a generic HTTP client

// Beacon client using the standard Beacon HTTP REST API (https://ethereum.github.io/beacon-APIs/)
type StandardHttpClient struct {
	providerAddress string
	httpClient      *http.Client
}

// Create a new client instance
func NewStandardHttpClient(providerAddress string, skipTlsVerification bool) *StandardHttpClient {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipTlsVerification,
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	httpClient := &http.Client{
		Transport: tr,
	}

	return &StandardHttpClient{
		providerAddress: providerAddress,
		httpClient:      httpClient,
	}
}

func (c *StandardHttpClient) GetValidatorPubKeys() ([]string, error) {

	// Get public keys
	body, status, err := c.getRequest(PublicKeyListEndpoint)
	if err != nil {
		return []string{}, err
	}
	if status != http.StatusOK {
		return []string{}, fmt.Errorf("Error getting validator public keys. Status code: %d", status)
	}

	// Parse response
	response := []string{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []string{}, err
	}

	// Return
	return response, nil

}

func (c *StandardHttpClient) GetDepositDataSignature(pubKey string, withdrawCredentials string, amount *big.Int, eth2Config beacon.Eth2Config) (string, error) {

	genesisForkVersion := common.BytesToHash(eth2Config.GenesisForkVersion)
	last8Bytes := genesisForkVersion[len(genesisForkVersion)-4:]

	// Get signature
	body, status, err := c.postRequest(fmt.Sprintf(SignatureEndpoint, pubKey), DepositDataSignatureRequest{
		Type: "DEPOSIT",
		Deposit: DepositInfo{
			pubKey,
			withdrawCredentials,
			amount.String(),
			common.Bytes2Hex(last8Bytes),
		},
	})
	if err != nil {
		return "", err
	}
	if status != http.StatusOK {
		return "", fmt.Errorf("error getting deposit data signature. Status code: %d", status)
	}

	// Return
	return string(body), nil

}

func (c *StandardHttpClient) ReloadKeys() error {
	_, status, err := c.postRequest(Reload, nil)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return fmt.Errorf("error getting voluntary exit message signature. Status code: %d", status)
	}

	return nil
}

func (c *StandardHttpClient) GetVoluntaryExitMessageSignature(validatorPubKey string, validatorIndex uint64, epoch uint64, forkInfo beacon.ForkInfo, eth2Config beacon.Eth2Config) (string, error) {
	genesisValidatorRoot := common.BytesToHash(eth2Config.GenesisValidatorsRoot)
	epochString := strconv.FormatUint(forkInfo.Epoch, 10)

	// TODO - we currently only support mainnet  We will have to update this as we support testnet
	previousVersion := eth2.MainnetBellatrixForkVersion
	currentVersion := eth2.MainnetCapellaForkVersion

	body, status, err := c.postRequest(fmt.Sprintf(SignatureEndpoint, validatorPubKey), VoluntaryExitMessageSignatureRequest{
		Type: "VOLUNTARY_EXIT",
		ForkInfoData: SigningForkInfo{
			Fork: struct {
				PreviousVersion string `json:"previous_version"`
				CurrentVersion  string `json:"current_version"`
				Epoch           string `json:"epoch"`
			}{
				PreviousVersion: previousVersion,
				CurrentVersion:  currentVersion,
				Epoch:           epochString,
			},
			GenesisValidatorRoot: genesisValidatorRoot.String(),
		},
		VoluntaryExitData: SigningVoluntaryExit{
			ValidatorIndex: strconv.FormatUint(validatorIndex, 10),
			Epoch:          strconv.FormatUint(epoch, 10),
		},
	})

	if err != nil {
		return "", err
	}
	if status != http.StatusOK {
		return "", fmt.Errorf("error getting voluntary exit message signature. Status code: %d", status)
	}

	// Return
	return string(body), nil

}

func (c *StandardHttpClient) HealthCheck() error {
	_, status, err := c.getRequest(HealthCheck)
	if status != http.StatusOK || err != nil {
		return err
	}

	return nil
}

//func (c *StandardHttpClient)

// Make a GET request to the beacon node
func (c *StandardHttpClient) getRequest(requestPath string) ([]byte, int, error) {
	// Send request
	response, err := c.httpClient.Get(fmt.Sprintf(RequestUrlFormat, c.providerAddress, requestPath))
	if err != nil {
		return []byte{}, 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Get response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	// Return
	return body, response.StatusCode, nil

}

// Make a POST request to the beacon node
func (c *StandardHttpClient) postRequest(requestPath string, requestBody interface{}) ([]byte, int, error) {

	// Get request body
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return []byte{}, 0, err
	}
	requestBodyReader := bytes.NewReader(requestBodyBytes)

	// Send request
	response, err := c.httpClient.Post(fmt.Sprintf(RequestUrlFormat, c.providerAddress, requestPath), RequestContentType, requestBodyReader)
	if err != nil {
		return []byte{}, 0, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Get response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, 0, err
	}

	// Return
	return body, response.StatusCode, nil

}
