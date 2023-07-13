package validator

import (
	"fmt"
	"github.com/prysmaticlabs/prysm/v3/crypto/bls"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/shared/utils/validator"
	"github.com/stader-labs/stader-node/stader-lib/node"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"github.com/urfave/cli"
	_ "golang.org/x/sync/errgroup"
	"math/big"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/stader-labs/stader-node/shared/utils/eth1"
)

func VerifyDepositSignatures(validatorPubKey types.ValidatorPubkey, preDepositHash [32]byte, depositHash [32]byte, preDepositSignature string, depositSignature string) (bool, error) {
	signaturesToVerify := [][]byte{}
	signaturesToVerify = append(signaturesToVerify, []byte(preDepositSignature))
	signaturesToVerify = append(signaturesToVerify, []byte(depositSignature))

	msgsSigned := [][32]byte{}
	msgsSigned = append(msgsSigned, preDepositHash)
	msgsSigned = append(msgsSigned, depositHash)

	pubKey, err := bls.PublicKeyFromBytes(validatorPubKey[:])
	if err != nil {
		return false, err
	}
	pubKeys := []bls.PublicKey{}
	pubKeys = append(pubKeys, pubKey)
	pubKeys = append(pubKeys, pubKey)

	res, err := bls.VerifyMultipleSignatures(signaturesToVerify, msgsSigned, pubKeys)
	if err != nil {
		return false, err
	}

	return res, nil
}

func canRegisterValidators(c *cli.Context, validatorList string) (*api.CanRegisterValidatorsResponse, error) {
	if err := services.RequireNodeWallet(c); err != nil {
		return nil, err
	}
	if err := services.RequireNodeRegistered(c); err != nil {
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

	canNodeDepositResponse := api.CanRegisterValidatorsResponse{}

	registeredPublicKeys, err := web3SignerClient.GetValidatorPubKeysList()
	if err != nil {
		return nil, err
	}

	validators := string_utils.DestringifyStringArray(validatorList)
	numValidators := big.NewInt(int64(len(validators)))

	inputKeyLimit, err := node.GetInputKeyCountLimit(prn, nil)
	if err != nil {
		return nil, err
	}

	if len(validators) > int(inputKeyLimit) {
		canNodeDepositResponse.InputKeyLimitReached = true
		canNodeDepositResponse.InputKeyLimit = inputKeyLimit
		return &canNodeDepositResponse, nil
	}

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

		valPubKey, err := types.HexToValidatorPubkey(val[2:])
		if err != nil {
			return nil, err
		}
		validatorId, err := node.GetValidatorIdByPubKey(prn, valPubKey.Bytes(), nil)
		if err != nil {
			return nil, err
		}
		if validatorId.Cmp(big.NewInt(0)) != 0 {
			return nil, fmt.Errorf("validatorPubkey %s is already registered with the permissioned node registry", val)
		}

		// add check for whether validator is present in beacon chain or not
		validatorStatus, err := bc.GetValidatorStatus(valPubKey, nil)
		if err != nil {
			return nil, err
		}
		if validatorStatus.Exists {
			return nil, fmt.Errorf("validator %s is already registered with the beacon chain", val)
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

		_, preDepositRootHash, err := validator.GetDepositDataSigningRoot(validatorPubkey, withdrawCredentials, eth2Config, 1000000000)
		if err != nil {
			return nil, err
		}
		_, depositRootHash, err := validator.GetDepositDataSigningRoot(validatorPubkey, withdrawCredentials, eth2Config, 31000000000)
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

		decodedHexPubKey, err := types.HexToValidatorPubkey(validatorPubkey[2:])
		if err != nil {
			return nil, err
		}

		decodedHexPreDepositSignature, err := types.HexToValidatorSignature(preDepositSignature[2:])
		if err != nil {
			return nil, err
		}

		decodedHexDepositSignature, err := types.HexToValidatorSignature(depositSignature[2:])
		if err != nil {
			return nil, err
		}

		pubKeys[i] = decodedHexPubKey.Bytes()
		preDepositSignatures[i] = decodedHexPreDepositSignature.Bytes()
		depositSignatures[i] = decodedHexDepositSignature.Bytes()

		res, err := VerifyDepositSignatures(decodedHexPubKey, preDepositRootHash, depositRootHash, preDepositSignature, depositSignature)
		if err != nil {
			return nil, err
		}
		if !res {
			return nil, fmt.Errorf("Deposit signatures are invalid for validator: %s\n", decodedHexPubKey.String())
		}

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

func registerValidators(c *cli.Context, validatorList string) (*api.ValidatorRegisterResponse, error) {

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
	numValidators := big.NewInt(int64(len(validators)))

	// Response
	response := api.ValidatorRegisterResponse{}

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

		decodedHexPubKey, err := types.HexToValidatorPubkey(validatorPubkey[2:])
		if err != nil {
			return nil, err
		}

		decodedHexPreDepositSignature, err := types.HexToValidatorSignature(preDepositSignature[2:])
		if err != nil {
			return nil, err
		}

		decodedHexDepositSignature, err := types.HexToValidatorSignature(depositSignature[2:])
		if err != nil {
			return nil, err
		}

		pubKeys[i] = decodedHexPubKey.Bytes()
		preDepositSignatures[i] = decodedHexPreDepositSignature.Bytes()
		depositSignatures[i] = decodedHexDepositSignature.Bytes()

		newValidatorKey.Add(newValidatorKey, big.NewInt(1))
	}

	// Override the provided pending TX if requested
	err = eth1.CheckForNonceOverride(c, opts)
	if err != nil {
		return nil, fmt.Errorf("error checking for nonce override: %w", err)
	}

	tx, err := node.AddValidatorKeys(prn, pubKeys, preDepositSignatures, depositSignatures, opts)
	if err != nil {
		return nil, err
	}

	response.TxHash = tx.Hash()

	// Return response
	return &response, nil

}
