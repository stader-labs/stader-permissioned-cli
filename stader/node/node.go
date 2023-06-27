package node

import (
	"fmt"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/stader-labs/stader-node/shared/services"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"
)

// Config
var preSignedCooldown, _ = time.ParseDuration("1h")
var merkleProofsDownloaderInterval, _ = time.ParseDuration("3h")

const (
	MaxConcurrentEth1Requests   = 200
	ErrorColor                  = color.FgRed
	InfoColor                   = color.FgHiGreen
	MerkleProofsDownloaderColor = color.FgHiYellow
)

// Register node command
func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Run Stader node activity daemon",
		Action: func(c *cli.Context) error {
			return run(c)
		},
	})
}

// Run daemon
func run(c *cli.Context) error {
	w, err := services.GetWallet(c)
	if err != nil {
		return err
	}
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return err
	}
	w3signer, err := services.GetWeb3SignerClient(c)
	if err != nil {
		return err
	}
	pnr, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return err
	}

	err = services.WaitEthClientSynced(c, true)
	if err != nil {
		return err
	}
	err = services.WaitNodeRegistered(c, nodeAccount.Address, true)
	if err != nil {
		return err
	}

	// Configure
	configureHTTP()

	// Initialize loggers
	errorLog := log.NewColorLogger(ErrorColor)
	infoLog := log.NewColorLogger(InfoColor)

	merkleProofsDownloader, err := NewMerkleProofsDownloader(c, log.NewColorLogger(MerkleProofsDownloaderColor))
	if err != nil {
		return err
	}

	publicKey, err := stader.GetPublicKey(c)
	if err != nil {
		return err
	}

	// Wait group to handle the various threads
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// validator presigned loop
	go func() {
		for {
			// Check the EC status
			err := services.WaitEthClientSynced(c, false) // Force refresh the primary / fallback EC status
			if err != nil {
				errorLog.Println(err)
				continue
			} else {
				// Check the BC status
				err := services.WaitBeaconClientSynced(c, false) // Force refresh the primary / fallback BC status
				if err != nil {
					errorLog.Println(err)
					continue
				}
			}

			operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
			if err != nil {
				errorLog.Printf("Failed to get operator id: %s\n", err.Error())
				continue
			}

			// make a map of all validators actually registered with stader
			// user might just move the validator keys to the directory. we don't wanna send the presigned msg of them
			infoLog.Println("Building a map of user validators registered with stader")
			registeredValidators, validatorPubKeys, err := stdr.GetAllValidatorsRegisteredWithOperator(pnr, operatorId, nodeAccount.Address, nil)
			if err != nil {
				errorLog.Printf("Could not get all validators registered with operator %s with error %s\n", operatorId, err.Error())
				continue
			}

			infoLog.Printlnf("Found %d validators registered with operator %s", len(registeredValidators), operatorId)
			infoLog.Println("Starting a pass of the presign daemon!")

			currentHead, err := bc.GetBeaconHead()
			if err != nil {
				errorLog.Printf("Could not get beacon head with error %s\n", err.Error())
				continue
			}

			forkInfo, err := bc.GetForkInfo()
			if err != nil {
				errorLog.Printf("Could not get fork info from beacon chain: %s\n", err)
				continue
			}
			eth2Config, err := bc.GetEth2Config()
			if err != nil {
				errorLog.Printf("Could not get eth2 config from beacon chain: %s\n", err)
				continue
			}

			err = w.Reload()
			if err != nil {
				errorLog.Printf("Could not reload wallet: %s\n", err.Error())
				continue
			}

			preSignRegisteredMap, err := stader.BulkIsPresignedKeyRegistered(c, validatorPubKeys)
			if err != nil {
				errorLog.Printf("Could not bulk check presigned keys with error %s\n", err.Error())
				continue
			}

			preSignSendMessages := []stader_backend.PreSignSendApiRequestType{}

			pageNumber := 0
			pageSize := 40
			for {
				startIndex := pageNumber * pageSize
				if startIndex > len(validatorPubKeys) {
					break
				}
				endIndex := (pageNumber + 1) * pageSize
				if endIndex > len(validatorPubKeys) {
					endIndex = len(validatorPubKeys)
				}

				infoLog.Printf("Starting a batch of %d validators\n", endIndex-startIndex)
				infoLog.Printf("Starting index: %d, End index: %d\n", startIndex, endIndex)

				validatorKeyBatch := validatorPubKeys[startIndex:endIndex]

				for _, validatorPubKey := range validatorKeyBatch {
					infoLog.Printf("Checking validator pubkey %s\n", validatorPubKey.String())
					validatorInfo, ok := registeredValidators[validatorPubKey]
					if !ok {
						errorLog.Printf("Validator pub key: %s not found in stader contracts\n", validatorPubKey)
						continue
					}
					if stdr.IsValidatorTerminal(validatorInfo) {
						errorLog.Printf("Validator pub key: %s is in terminal state in the stader contracts\n", validatorPubKey)
						continue
					}

					registeredPresign, ok := preSignRegisteredMap[validatorPubKey.String()]
					if !ok {
						errorLog.Printf("Could not query presign api to check if validator: %s is registered\n", validatorPubKey)
						continue
					}
					if registeredPresign {
						infoLog.Printf("Validator pub key: %s pre signed key already registered\n", validatorPubKey)
						continue
					} else {
						infoLog.Printf("Validator pub key: %s pre signed key not registered. Creating presigned message\n", validatorPubKey)
					}

					// check if validator has not yet been registered on beacon chain
					validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
					if err != nil {
						errorLog.Printf("Error finding validator status for validator: %s err: %s\n", validatorPubKey, err.Error())
						continue
					}
					if !validatorStatus.Exists {
						errorLog.Printf("Validator pub key: %s not found on beacon chain\n", validatorPubKey)
						continue
					}

					// check if validator is already in an exiting phase, then no point sending a pre-signed message
					if eth2.IsValidatorExiting(validatorStatus) {
						errorLog.Printf("Validator pub key: %s already exiting or exited with status %s", validatorPubKey, validatorStatus.Status)
						continue
					}

					exitEpoch := currentHead.Epoch

					// get the presigned msg
					hexSignature, err := w3signer.GetVoluntaryExitSignature(validatorPubKey.String(), validatorStatus.Index, exitEpoch, forkInfo, eth2Config)
					if err != nil {
						errorLog.Printf("Failed to generate the SignedExitMessage for validator with pub key: %s with err: %s\n", validatorPubKey.String(), err.Error())
						continue
					}

					signature, err := types.HexToValidatorSignature(hexSignature[2:])
					if err != nil {
						errorLog.Printf("Failed to convert signature to validator signature: %s with err: %s\n", hexSignature, err.Error())
						continue
					}

					// encrypt the signature and srHash
					exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(signature.String()), publicKey)
					if err != nil {
						errorLog.Printf("Failed to encrypt exit signature for validator: %s with err: %s\n", validatorPubKey, err.Error())
						continue
					}
					exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)

					// send it to the presigned api
					preSignSendMessages = append(preSignSendMessages, stader_backend.PreSignSendApiRequestType{
						Message: struct {
							Epoch          string `json:"epoch"`
							ValidatorIndex string `json:"validator_index"`
						}{
							Epoch:          strconv.FormatUint(exitEpoch, 10),
							ValidatorIndex: strconv.FormatUint(validatorStatus.Index, 10),
						},
						Signature:          exitSignatureEncryptedString,
						ValidatorPublicKey: validatorPubKey.String(),
					})
				}

				//fmt.Printf("Sending %d presigned messages to stader backend\n", len(preSignSendMessages))
				fmt.Printf("Sending %d presigned messages to stader backend\n", len(preSignSendMessages))
				fmt.Printf("Pre-sign messages being sent are %v\n", preSignSendMessages)
				if len(preSignSendMessages) > 0 {
					res, err := stader.SendBulkPresignedMessageToStaderBackend(c, preSignSendMessages)
					if err != nil {
						errorLog.Printf("Sending bulk presigned message failed with %v\n", err)
					}
					for pubKey, response := range *res {
						if response.Success {
							infoLog.Printf("Successfully sent the presigned message for validator: %s\n", pubKey)
						} else {
							errorLog.Printf("Failed to send the presigned api: %s\n", response.Message)
						}
					}
				}

				pageNumber += 1
			}

			infoLog.Printf("Done with the pass of presign daemon")
			// run loop every 12 hours
			time.Sleep(preSignedCooldown)
		}

		wg.Done()
	}()

	go func() {
		for {
			infoLog.Printlnf("Checking if there are any available merkle proofs to download")

			// Manage the fee recipient for the node
			if err := merkleProofsDownloader.run(); err != nil {
				errorLog.Println(err)
			}

			infoLog.Printlnf("Done checking for merkle proofs to download")
			time.Sleep(merkleProofsDownloaderInterval)
		}

		wg.Done()
	}()

	// Wait for both threads to stop
	wg.Wait()
	return nil

}

// Configure HTTP transport settings
func configureHTTP() {

	// The guardian daemon makes a large number of concurrent RPC requests to the Eth1 client
	// The HTTP transport is set to cache connections for future re-use equal to the maximum expected number of concurrent requests
	// This prevents issues related to memory consumption and address allowance from repeatedly opening and closing connections
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = MaxConcurrentEth1Requests

}
