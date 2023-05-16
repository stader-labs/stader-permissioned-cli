/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [0.3.0-beta]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package node

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/stader-labs/stader-node/shared/services"
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"github.com/stader-labs/stader-node/shared/utils/crypto"
	"github.com/stader-labs/stader-node/shared/utils/eth2"
	"github.com/stader-labs/stader-node/shared/utils/log"
	"github.com/stader-labs/stader-node/shared/utils/stader"
	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// Config
var preSignedCooldown, _ = time.ParseDuration("30m")
var preSignedBatchCooldown, _ = time.ParseDuration("2s")
var preSignBatchSize = 10 // Go thru 100 keys in each pass
var merkleProofsDownloaderInterval, _ = time.ParseDuration("5m")

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
	nodeAccount, err := w.GetNodeAccount()
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

	operatorId, err := node.GetOperatorId(pnr, nodeAccount.Address, nil)
	if err != nil {
		return err
	}

	// Wait group to handle the various threads
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// validator presigned loop
	go func() {
		for {
			pnr, err := services.GetPermissionedNodeRegistry(c)
			if err != nil {
				errorLog.Printf("Could not get permissioned node registry: %s\n", err)
				continue
			}

			publicKey, err := stader.GetPublicKey()
			if err != nil {
				errorLog.Printf("Could not get public key: %s\n", err)
				continue
			}
			// make a map of all validators actually registered with stader
			// user might just move the validator keys to the directory. we don't wanna send the presigned msg of them

			infoLog.Println("Building a map of user validators registered with stader")
			registeredValidators, err := stdr.GetAllValidatorsRegisteredWithOperator(pnr, operatorId, nil)
			if err != nil {
				errorLog.Printf("Could not get all validators registered with operator %s\n", operatorId)
				continue
			}

			infoLog.Println("Starting a pass of the presign daemon!")

			currentHead, err := bc.GetBeaconHead()
			if err != nil {
				panic("not able to communicate with beacon chain!")
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

			for validatorPubKey, validatorInfo := range registeredValidators {
				infoLog.Printf("Checking validator Pub key: %s\n", validatorPubKey.String())

				if stdr.IsValidatorTerminal(validatorInfo) {
					errorLog.Printf("Validator pub key: %s is in terminal state in the stader contracts\n", validatorPubKey)
					continue
				}

				// check if validator has not yet been registered on beacon chain
				validatorStatus, err := bc.GetValidatorStatus(validatorPubKey, nil)
				if err != nil {
					errorLog.Printf("Error finding validator status for validator: %s\n", validatorPubKey)
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

				// check if the presigned message has been registered. if it has been registered, then continue
				isRegistered, err := stader.IsPresignedKeyRegistered(validatorPubKey)
				if isRegistered {
					errorLog.Printf("Validator pub key: %s pre signed key already registered\n", validatorPubKey)
					continue
				} else if err != nil {
					errorLog.Printf("Could not query presign api to check if validator: %s is registered\n", validatorPubKey)
					continue
				}

				exitEpoch := currentHead.Epoch

				hexSignature, err := w3signer.GetVoluntaryExitSignature(validatorPubKey.String(), validatorStatus.Index, exitEpoch, forkInfo, eth2Config)
				if err != nil {
					errorLog.Printf("Failed to generate the SignedExitMessage for validator with pub key: %s\n", validatorPubKey.String())
					continue
				}

				signature, err := types.HexToValidatorSignature(hexSignature[2:])
				if err != nil {
					errorLog.Printf("Failed to convert signature to validator signature: %s\n", hexSignature)
					continue
				}

				// encrypt the signature and srHash
				exitSignatureEncrypted, err := crypto.EncryptUsingPublicKey([]byte(signature.String()), publicKey)
				if err != nil {
					errorLog.Printf("Failed to encrypt exit signature for validator: %s\n", validatorPubKey)
					continue
				}
				//fmt.Printf("Encrypted exit signature: %v\n", exitSignatureEncrypted)
				exitSignatureEncryptedString := crypto.EncodeBase64(exitSignatureEncrypted)

				// send it to the presigned api
				backendRes, err := stader.SendPresignedMessageToStaderBackend(stader_backend.PreSignSendApiRequestType{
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
				if !backendRes.Success {
					errorLog.Printf("Failed to send the presigned api: %s\n", backendRes.Message)
					continue
				} else if backendRes.Success {
					errorLog.Printf("Successfully sent the presigned message for validator: %s\n", validatorPubKey)
					continue
				} else if err != nil {
					errorLog.Printf("Sending presigned message failed with %v\n", err)
					continue
				}

				time.Sleep(preSignedBatchCooldown)
			}

			infoLog.Printf("Done with the pass of presign daemon")

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
