package node

import (
	"encoding/hex"
	"fmt"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
)

func canNodeDeposit(c *cli.Context, validatorList string) (*api.CanNodeDepositResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	vfc, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	web3SignerClient, err := services.GetWeb3SignerClient(c)
	if err != nil {
		return nil, err
	}

	registeredPublicKeys, err := web3SignerClient.GetValidatorPubKeysList()
	if err != nil {
		return nil, err
	}

	validators := string_utils.DestringifyStringArray(validatorList)
	numValidators := big.NewInt(int64(len(validators)))

	if len(validators) > len(registeredPublicKeys) {
		return nil, fmt.Errorf("number of validators to register is greater than the number of validators registered with the web3signer")
	}
	for _, val := range validators {
		found := false
		for _, regVal := range registeredPublicKeys {
			if val == regVal {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("validatorPubkey %s is not registered with the web3signer", val)
		}

		validatorId, err := node.GetValidatorIdByPubKey(prn, []byte(val), nil)
		if err != nil {
			return nil, err
		}
		if validatorId.Cmp(big.NewInt(0)) != 0 {
			return nil, fmt.Errorf("validatorPubkey %s is already registered with the permissioned node registry", val)
		}
	}

	// Get eth2 config
	eth2Config, err := bc.GetEth2Config()
	if err != nil {
		return nil, err
	}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	canNodeDepositResponse := api.CanNodeDepositResponse{}

	isPermissionedNodeRegistryPaused, err := node.IsPermissionedNodeRegistryPaused(prn, nil)
	if err != nil {
		return nil, err
	}
	if isPermissionedNodeRegistryPaused {
		canNodeDepositResponse.DepositPaused = true
		return &canNodeDepositResponse, nil
	}

	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}
	if operatorId.Cmp(big.NewInt(0)) == 0 {
		canNodeDepositResponse.OperatorNotRegistered = true
		return &canNodeDepositResponse, nil
	}

	operatorInfo, err := node.GetOperatorInfo(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}
	if !operatorInfo.Active {
		canNodeDepositResponse.OperatorNotActive = true
		return &canNodeDepositResponse, nil
	}

	totalValidatorKeys, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}
	totalValidatorNonTerminalKeys, err := node.GetTotalNonTerminalValidatorKeys(prn, nodeAccount.Address, totalValidatorKeys, nil)
	if err != nil {
		return nil, err
	}
	maxKeysPerOperator, err := node.GetMaxValidatorKeysPerOperator(prn, nil)
	if err != nil {
		return nil, err
	}

	if totalValidatorNonTerminalKeys+numValidators.Uint64() > maxKeysPerOperator {
		canNodeDepositResponse.MaxValidatorLimitReached = true
		return &canNodeDepositResponse, nil
	}

	pubKeys := make([][]byte, numValidators.Int64())
	preDepositSignatures := make([][]byte, numValidators.Int64())
	depositSignatures := make([][]byte, numValidators.Int64())

	operatorKeyCount, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}

	newValidatorKey := operatorKeyCount

	for i, validatorPubkey := range validators {
		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(vfc, 2, operatorId, newValidatorKey, nil)
		if err != nil {
			return nil, err
		}

		withdrawCredentials, err := node.GetValidatorWithdrawalCredential(vfc, rewardWithdrawVault, nil)
		if err != nil {
			return nil, err
		}

		// TODO - bchain - add check with sr hash
		preDepositSignature, err := web3SignerClient.GetDepositDataSignature(validatorPubkey, withdrawCredentials.String(), big.NewInt(1000000000), eth2Config)
		if err != nil {
			return nil, err
		}
		depositSignature, err := web3SignerClient.GetDepositDataSignature(validatorPubkey, withdrawCredentials.String(), big.NewInt(31000000000), eth2Config)
		if err != nil {
			return nil, err
		}

		decodedHexPubKey, err := hex.DecodeString(validatorPubkey[2:])
		if err != nil {
			return nil, err
		}
		decodedHexPreDeposit, err := hex.DecodeString(preDepositSignature[2:])
		if err != nil {
			return nil, err
		}
		decodedHexDeposit, err := hex.DecodeString(depositSignature[2:])
		if err != nil {
			return nil, err
		}

		pubKeys[i] = decodedHexPubKey[:]
		preDepositSignatures[i] = decodedHexPreDeposit[:]
		depositSignatures[i] = decodedHexDeposit[:]

		newValidatorKey.Add(newValidatorKey, big.NewInt(1))
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	//// Do not send transaction unless requested
	gasInfo, err := node.EstimateAddValidatorKeys(prn, pubKeys, preDepositSignatures, depositSignatures, opts)
	if err != nil {
		return nil, err
	}

	canNodeDepositResponse.CanDeposit = true
	canNodeDepositResponse.GasInfo = gasInfo

	return &canNodeDepositResponse, nil
}

func nodeDeposit(c *cli.Context, validatorList string, submit bool) (*api.NodeDepositResponse, error) {

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}
	prn, err := services.GetPermissionedNodeRegistry(c)
	if err != nil {
		return nil, err
	}
	bc, err := services.GetBeaconClient(c)
	if err != nil {
		return nil, err
	}
	vfc, err := services.GetVaultFactory(c)
	if err != nil {
		return nil, err
	}
	web3SignerClient, err := services.GetWeb3SignerClient(c)
	if err != nil {
		return nil, err
	}

	// Get eth2 config
	eth2Config, err := bc.GetEth2Config()
	if err != nil {
		return nil, err
	}

	// Get node account
	nodeAccount, err := w.GetNodeAccount()
	if err != nil {
		return nil, err
	}

	validators := string_utils.DestringifyStringArray(validatorList)
	//fmt.Printf("validators: %v\n", validators)
	numValidators := big.NewInt(int64(len(validators)))
	//fmt.Printf("numValidators: %d\n", numValidators)

	// Response
	response := api.NodeDepositResponse{}

	// get the vault address and vault credential
	operatorId, err := node.GetOperatorId(prn, nodeAccount.Address, nil)
	if err != nil {
		return nil, err
	}

	// Get transactor
	opts, err := w.GetNodeAccountTransactor()
	if err != nil {
		return nil, err
	}

	pubKeys := make([][]byte, numValidators.Int64())
	preDepositSignatures := make([][]byte, numValidators.Int64())
	depositSignatures := make([][]byte, numValidators.Int64())

	validatorKeyCount, err := node.GetTotalValidatorKeys(prn, operatorId, nil)
	if err != nil {
		return nil, err
	}

	newValidatorKey := validatorKeyCount

	for i, validatorPubkey := range validators {

		rewardWithdrawVault, err := node.ComputeWithdrawVaultAddress(vfc, 2, operatorId, newValidatorKey, nil)
		if err != nil {
			return nil, err
		}

		withdrawCredentials, err := node.GetValidatorWithdrawalCredential(vfc, rewardWithdrawVault, nil)
		if err != nil {
			return nil, err
		}

		preDepositSignature, err := web3SignerClient.GetDepositDataSignature(validatorPubkey, withdrawCredentials.String(), big.NewInt(1000000000), eth2Config)
		if err != nil {
			return nil, err
		}
		depositSignature, err := web3SignerClient.GetDepositDataSignature(validatorPubkey, withdrawCredentials.String(), big.NewInt(31000000000), eth2Config)
		if err != nil {
			return nil, err
		}

		decodedHexPubKey, err := hex.DecodeString(validatorPubkey[2:])
		if err != nil {
			return nil, err
		}
		decodedHexPreDeposit, err := hex.DecodeString(preDepositSignature[2:])
		if err != nil {
			return nil, err
		}
		decodedHexDeposit, err := hex.DecodeString(depositSignature[2:])
		if err != nil {
			return nil, err
		}

		pubKeys[i] = decodedHexPubKey
		preDepositSignatures[i] = decodedHexPreDeposit
		depositSignatures[i] = decodedHexDeposit

		newValidatorKey.Add(newValidatorKey, big.NewInt(1))
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	// Do not send transaction unless requested
	opts.NoSend = !submit

	tx, err := node.AddValidatorKeys(prn, pubKeys, preDepositSignatures, depositSignatures, opts)
	if err != nil {
		return nil, err
	}

	// Print transaction if requested
	if !submit {
		b, err := tx.MarshalBinary()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%x\n", b)
	}

	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
