/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.1]

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
package stader

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stader-labs/stader-node/shared/types/api"
	string_utils "github.com/stader-labs/stader-node/shared/utils/string-utils"
	"github.com/stader-labs/stader-node/stader-lib/types"
	"math/big"
)

// Get node status
func (c *Client) NodeStatus() (api.NodeStatusResponse, error) {
	responseBytes, err := c.callAPI("node status")
	if err != nil {
		return api.NodeStatusResponse{}, fmt.Errorf("could not get node status: %w", err)
	}
	var response api.NodeStatusResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeStatusResponse{}, fmt.Errorf("could not decode node status response: %w", err)
	}
	if response.Error != "" {
		return api.NodeStatusResponse{}, fmt.Errorf("could not get node status: %s", response.Error)
	}
	if response.AccountBalances.ETH == nil {
		response.AccountBalances.ETH = big.NewInt(0)
	}
	if response.AccountBalances.Sd == nil {
		response.AccountBalances.Sd = big.NewInt(0)
	}
	if response.OperatorRewardInETH == nil {
		response.OperatorRewardInETH = big.NewInt(0)
	}
	if response.TotalNonTerminalValidators == nil {
		response.TotalNonTerminalValidators = big.NewInt(0)
	}

	return response, nil
}

// Check whether the node can be registered
func (c *Client) CanRegisterNode(operatorName string, operatorRewardAddress common.Address) (api.CanRegisterNodeResponse, error) {
	responseBytes, err := c.callAPI("node can-register", operatorName, operatorRewardAddress.Hex())
	if err != nil {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("could not get can register node status: %w", err)
	}
	var response api.CanRegisterNodeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("could not decode can register node response: %w", err)
	}
	if response.Error != "" {
		return api.CanRegisterNodeResponse{}, fmt.Errorf("could not get can register node status: %s", response.Error)
	}
	return response, nil
}

// Register the node
func (c *Client) RegisterNode(operatorName string, operatorRewardAddress common.Address) (api.RegisterNodeResponse, error) {
	responseBytes, err := c.callAPI("node register", operatorName, operatorRewardAddress.Hex())
	if err != nil {
		return api.RegisterNodeResponse{}, fmt.Errorf("could not register node: %w", err)
	}
	var response api.RegisterNodeResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.RegisterNodeResponse{}, fmt.Errorf("could not decode register node response: %w", err)
	}
	if response.Error != "" {
		return api.RegisterNodeResponse{}, fmt.Errorf("could not register node: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can make a deposit
func (c *Client) CanRegisterValidators(validatorList string) (api.CanRegisterValidatorsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator can-register %s", validatorList))
	if err != nil {
		return api.CanRegisterValidatorsResponse{}, fmt.Errorf("could not get validator can-register status: %w", err)
	}
	var response api.CanRegisterValidatorsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanRegisterValidatorsResponse{}, fmt.Errorf("could not decode validator can-register response: %w", err)
	}
	if response.Error != "" {
		return api.CanRegisterValidatorsResponse{}, fmt.Errorf("could not get validator can-register status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) GetContractsInfo() (api.ContractsInfoResponse, error) {
	responseBytes, err := c.callAPI("node get-contracts-info")
	if err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not get networks contracts info: %w", err)
	}
	var response api.ContractsInfoResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not decode networks contracts info: %w", err)
	}
	if response.Error != "" {
		return api.ContractsInfoResponse{}, fmt.Errorf("could not get networks contract info: %s", response.Error)
	}
	return response, nil
}

func (c *Client) GetValidatorInfo(valPubKey types.ValidatorPubkey) (api.ValidatorInfoResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator get-validator-info %s", valPubKey))
	if err != nil {
		return api.ValidatorInfoResponse{}, fmt.Errorf("could not get validator info: %w", err)
	}
	var response api.ValidatorInfoResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ValidatorInfoResponse{}, fmt.Errorf("could not decode get-validator-info: %w", err)
	}
	if response.Error != "" {
		return api.ValidatorInfoResponse{}, fmt.Errorf("could not get-validator-info: %s", response.Error)
	}
	return response, nil
}

// Make a node deposit
func (c *Client) RegisterValidators(validatorPubKeyList string) (api.ValidatorRegisterResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator register %s", validatorPubKeyList))
	if err != nil {
		return api.ValidatorRegisterResponse{}, fmt.Errorf("could not make validator register as err: %w", err)
	}
	var response api.ValidatorRegisterResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ValidatorRegisterResponse{}, fmt.Errorf("could not decode validator register response: %w", err)
	}
	if response.Error != "" {
		return api.ValidatorRegisterResponse{}, fmt.Errorf("could not make validator register: %s", response.Error)
	}
	return response, nil
}

// Check whether the node can send tokens
func (c *Client) CanNodeSend(amountWei *big.Int, token string) (api.CanNodeSendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-send %s %s", amountWei.String(), token))
	if err != nil {
		return api.CanNodeSendResponse{}, fmt.Errorf("could not get can node send status: %w", err)
	}
	var response api.CanNodeSendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanNodeSendResponse{}, fmt.Errorf("could not decode can node send response: %w", err)
	}
	if response.Error != "" {
		return api.CanNodeSendResponse{}, fmt.Errorf("could not get can node send status: %s", response.Error)
	}
	return response, nil
}

// Send tokens from the node to an address
func (c *Client) NodeSend(amountWei *big.Int, token string, toAddress common.Address) (api.NodeSendResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node send %s %s %s", amountWei.String(), token, toAddress.Hex()))
	if err != nil {
		return api.NodeSendResponse{}, fmt.Errorf("could not send tokens from node: %w", err)
	}
	var response api.NodeSendResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.NodeSendResponse{}, fmt.Errorf("could not decode node send response: %w", err)
	}
	if response.Error != "" {
		return api.NodeSendResponse{}, fmt.Errorf("could not send tokens from node: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanSendClRewards(validatorPubKey types.ValidatorPubkey) (api.CanSendClRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator can-send-cl-rewards %s", validatorPubKey))
	if err != nil {
		return api.CanSendClRewardsResponse{}, fmt.Errorf("could not get validator can-send-cl-rewards response: %w", err)
	}
	var response api.CanSendClRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanSendClRewardsResponse{}, fmt.Errorf("could not decode validator can-send-cl-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanSendClRewardsResponse{}, fmt.Errorf("could not get validator can-send-cl-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) SendClRewards(validatorPubKey types.ValidatorPubkey) (api.SendClRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator send-cl-rewards %s", validatorPubKey))
	if err != nil {
		return api.SendClRewardsResponse{}, fmt.Errorf("could not get validator send-cl-rewards response: %w", err)
	}
	var response api.SendClRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.SendClRewardsResponse{}, fmt.Errorf("could not decode validator send-cl-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.SendClRewardsResponse{}, fmt.Errorf("could not get validator send-cl-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanDownloadSpMerkleProofs() (api.CanDownloadSpMerkleProofsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-download-sp-merkle-proofs"))
	if err != nil {
		return api.CanDownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node can-download-sp-merkle-proofs response: %w", err)
	}
	var response api.CanDownloadSpMerkleProofsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanDownloadSpMerkleProofsResponse{}, fmt.Errorf("could not decode node can-download-sp-merkle-proofs response: %w", err)
	}
	if response.Error != "" {
		return api.CanDownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node can-download-sp-merkle-proofs response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) DownloadSpMerkleProofs() (api.DownloadSpMerkleProofsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node download-sp-merkle-proofs"))
	if err != nil {
		return api.DownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node download-sp-merkle-proofs response: %w", err)
	}
	var response api.DownloadSpMerkleProofsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.DownloadSpMerkleProofsResponse{}, fmt.Errorf("could not decode node download-sp-merkle-proofs response: %w", err)
	}
	if response.Error != "" {
		return api.DownloadSpMerkleProofsResponse{}, fmt.Errorf("could not get node download-sp-merkle-proofs response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) GetDetailedCyclesInfo(cycles []*big.Int) (api.CyclesDetailedInfo, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node detailed-cycles-info %s", stringifiedCycleList))
	if err != nil {
		return api.CyclesDetailedInfo{}, fmt.Errorf("could not get node detailed-cycles-info response: %w", err)
	}
	var response api.CyclesDetailedInfo
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CyclesDetailedInfo{}, fmt.Errorf("could not decode node detailed-cycles-info response: %w", err)
	}
	if response.Error != "" {
		return api.CyclesDetailedInfo{}, fmt.Errorf("could not get node detailed-cycles-info response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanClaimRewards() (api.CanClaimRewards, error) {
	responseBytes, err := c.callAPI("node can-claim-rewards")
	if err != nil {
		return api.CanClaimRewards{}, fmt.Errorf("could not get node can-claim-rewards response: %w", err)
	}
	var response api.CanClaimRewards
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimRewards{}, fmt.Errorf("could not decode node can-claim-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimRewards{}, fmt.Errorf("could not get node can-claim-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimRewards() (api.ClaimRewards, error) {
	responseBytes, err := c.callAPI("node claim-rewards")
	if err != nil {
		return api.ClaimRewards{}, fmt.Errorf("could not get node claim-rewards response: %w", err)
	}
	var response api.ClaimRewards
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimRewards{}, fmt.Errorf("could not decode node claim-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimRewards{}, fmt.Errorf("could not get node claim-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanClaimSpRewards() (api.CanClaimSpRewardsResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("node can-claim-sp-rewards"))
	if err != nil {
		return api.CanClaimSpRewardsResponse{}, fmt.Errorf("could not get node can-claim-sp-rewards response: %w", err)
	}
	var response api.CanClaimSpRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanClaimSpRewardsResponse{}, fmt.Errorf("could not decode node can-claim-sp-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.CanClaimSpRewardsResponse{}, fmt.Errorf("could not get node can-claim-sp-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) EstimateClaimSpRewardsGas(cycles []*big.Int) (api.EstimateClaimSpRewardsGasResponse, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node estimate-claim-sp-rewards-gas %s", stringifiedCycleList))
	if err != nil {
		return api.EstimateClaimSpRewardsGasResponse{}, fmt.Errorf("could not get node estimate-claim-sp-rewards-gas response: %w", err)
	}
	var response api.EstimateClaimSpRewardsGasResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.EstimateClaimSpRewardsGasResponse{}, fmt.Errorf("could not decode node estimate-claim-sp-rewards-gas response: %w", err)
	}
	if response.Error != "" {
		return api.EstimateClaimSpRewardsGasResponse{}, fmt.Errorf("could not get node estimate-claim-sp-rewards-gas response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) ClaimSpRewards(cycles []*big.Int) (api.ClaimSpRewardsResponse, error) {
	stringifiedCycleList := string_utils.StringifyArray(cycles)
	responseBytes, err := c.callAPI(fmt.Sprintf("node claim-sp-rewards %s", stringifiedCycleList))
	if err != nil {
		return api.ClaimSpRewardsResponse{}, fmt.Errorf("could not get node claim-sp-rewards response: %w", err)
	}
	var response api.ClaimSpRewardsResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ClaimSpRewardsResponse{}, fmt.Errorf("could not decode node claim-sp-rewards response: %w", err)
	}
	if response.Error != "" {
		return api.ClaimSpRewardsResponse{}, fmt.Errorf("could not get node claim-sp-rewards response: %s", response.Error)
	}

	return response, nil
}

func (c *Client) CanUpdateOperatorName(operatorName string) (api.CanUpdateOperatorName, error) {
	responseBytes, err := c.callAPI("node can-update-operator-name", operatorName)
	if err != nil {
		return api.CanUpdateOperatorName{}, fmt.Errorf("could not get can-update-operator-name response: %w", err)
	}
	var response api.CanUpdateOperatorName
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUpdateOperatorName{}, fmt.Errorf("could not decode can-update-operator-name response: %w", err)
	}
	if response.Error != "" {
		return api.CanUpdateOperatorName{}, fmt.Errorf("could not get can-update-operator-name response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) UpdateOperatorName(operatorName string) (api.UpdateOperatorName, error) {
	responseBytes, err := c.callAPI("node update-operator-name", operatorName)
	if err != nil {
		return api.UpdateOperatorName{}, fmt.Errorf("could not get update-operator-name response: %w", err)
	}
	var response api.UpdateOperatorName
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.UpdateOperatorName{}, fmt.Errorf("could not decode update-operator-name response: %w", err)
	}
	if response.Error != "" {
		return api.UpdateOperatorName{}, fmt.Errorf("could not get update-operator-name response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanUpdateOperatorRewardAddress(operatorRewardAddress common.Address) (api.CanUpdateOperatorRewardAddress, error) {
	responseBytes, err := c.callAPI("node can-update-operator-reward-address", operatorRewardAddress.Hex())
	if err != nil {
		return api.CanUpdateOperatorRewardAddress{}, fmt.Errorf("could not get can-update-operator-reward-address response: %w", err)
	}
	var response api.CanUpdateOperatorRewardAddress
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanUpdateOperatorRewardAddress{}, fmt.Errorf("could not decode can-update-operator-reward-address response: %w", err)
	}
	if response.Error != "" {
		return api.CanUpdateOperatorRewardAddress{}, fmt.Errorf("could not get can-update-operator-reward-address response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) UpdateOperatorRewardAddress(operatorRewardAddress common.Address) (api.UpdateOperatorRewardAddress, error) {
	responseBytes, err := c.callAPI("node update-operator-reward-address", operatorRewardAddress.Hex())
	if err != nil {
		return api.UpdateOperatorRewardAddress{}, fmt.Errorf("could not get update-operator-reward-address response: %w", err)
	}
	var response api.UpdateOperatorRewardAddress
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.UpdateOperatorRewardAddress{}, fmt.Errorf("could not decode update-operator-reward-address response: %w", err)
	}
	if response.Error != "" {
		return api.UpdateOperatorRewardAddress{}, fmt.Errorf("could not get update-operator-reward-address response: %s", response.Error)
	}
	return response, nil
}

func (c *Client) CanExitValidator(validatorPubKey types.ValidatorPubkey) (api.CanExitValidatorResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator can-exit-validator %s", validatorPubKey))
	if err != nil {
		return api.CanExitValidatorResponse{}, fmt.Errorf("could not get can-exit-validator status: %w", err)
	}
	var response api.CanExitValidatorResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.CanExitValidatorResponse{}, fmt.Errorf("could not decode can-exit-validator response: %w", err)
	}
	if response.Error != "" {
		return api.CanExitValidatorResponse{}, fmt.Errorf("could not get can-exit-validator status: %s", response.Error)
	}
	return response, nil
}

func (c *Client) ExitValidator(validatorPubKey types.ValidatorPubkey) (api.ExitValidatorResponse, error) {
	responseBytes, err := c.callAPI(fmt.Sprintf("validator exit-validator %s", validatorPubKey))
	if err != nil {
		return api.ExitValidatorResponse{}, fmt.Errorf("could not get exit-validator status: %w", err)
	}
	var response api.ExitValidatorResponse
	if err := json.Unmarshal(responseBytes, &response); err != nil {
		return api.ExitValidatorResponse{}, fmt.Errorf("could not decode exit-validator response: %w", err)
	}
	if response.Error != "" {
		return api.ExitValidatorResponse{}, fmt.Errorf("could not get exit-validator status: %s", response.Error)
	}
	return response, nil
}
