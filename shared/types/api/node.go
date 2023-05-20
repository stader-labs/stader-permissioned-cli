package api

import (
	stader_backend "github.com/stader-labs/stader-node/shared/types/stader-backend"
	"math/big"
	"time"

	"github.com/stader-labs/stader-node/shared/utils/stdr"
	"github.com/stader-labs/stader-node/stader-lib/stader"
	"github.com/stader-labs/stader-node/stader-lib/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stader-labs/stader-node/stader-lib/tokens"
)

type NodeStatusResponse struct {
	Status                            string                             `json:"status"`
	Error                             string                             `json:"error"`
	NumberOfValidatorsRegistered      string                             `json:"numberOfValidatorsRegistered"`
	AccountAddress                    common.Address                     `json:"accountAddress"`
	AccountAddressFormatted           string                             `json:"accountAddressFormatted"`
	OperatorId                        *big.Int                           `json:"operatorId"`
	OperatorName                      string                             `json:"operatorName"`
	OperatorAddress                   common.Address                     `json:"operatorAddress"`
	OperatorActive                    bool                               `json:"operatorActive"`
	OperatorRewardAddress             common.Address                     `json:"operatorRewardAddress"`
	Registered                        bool                               `json:"registered"`
	AccountBalances                   tokens.Balances                    `json:"accountBalances"`
	OperatorRewardInETH               *big.Int                           `json:"operatorRewardInETH"`
	SocializingPoolRewardCycleDetails types.RewardCycleDetails           `json:"socializingPoolRewardCycleDetails"`
	SocializingPoolStartTime          time.Time                          `json:"socializingPoolStartTime"`
	TotalNonTerminalValidators        *big.Int                           `json:"totalNonTerminalValidators"`
	ValidatorInfos                    []stdr.ValidatorInfo               `json:"validatorInfos"`
	ClaimedSocializingPoolMerkles     []stader_backend.CycleMerkleProofs `json:"claimedSocializingPoolMerkles"`
	UnclaimedSocializingPoolMerkles   []stader_backend.CycleMerkleProofs `json:"unclaimedSocializingPoolMerkles"`
	Web3SignerUrl                     string                             `json:"web3SignerUrl"`
	Web3SignerConnectionSuccess       bool                               `json:"web3SignerConnectionSuccess"`
	Web3SignerConnectionError         string                             `json:"web3SignerConnectionError"`
	BeaconChainUrl                    string                             `json:"beaconChainUrl"`
	ExecutionChainUrl                 string                             `json:"executionChainUrl"`
	OperatorWhitelisted               bool                               `json:"operatorWhitelisted"`
}

type CanRegisterNodeResponse struct {
	Status                    string         `json:"status"`
	Error                     string         `json:"error"`
	AlreadyRegistered         bool           `json:"alreadyRegistered"`
	RegistrationPaused        bool           `json:"registrationPaused"`
	OperatorNameTooLong       bool           `json:"operatorNameTooLong"`
	OperatorRewardAddressZero bool           `json:"operatorRewardAddressZero"`
	OperatorNotWhitelisted    bool           `json:"operatorNotWhitelisted"`
	GasInfo                   stader.GasInfo `json:"gasInfo"`
}

type RegisterNodeResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanNodeDepositSdResponse struct {
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	CollateralContractPaused bool           `json:"collateralContractPaused"`
	InsufficientBalance      bool           `json:"insufficientBalance"`
	GasInfo                  stader.GasInfo `json:"gasInfo"`
}
type NodeDepositSdApproveGasResponse struct {
	Status  string         `json:"status"`
	Error   string         `json:"error"`
	GasInfo stader.GasInfo `json:"gasInfo"`
}
type NodeDepositSdApproveResponse struct {
	Status        string      `json:"status"`
	Error         string      `json:"error"`
	ApproveTxHash common.Hash `json:"approveTxHash"`
}
type NodeDepositSdResponse struct {
	Status        string      `json:"status"`
	Error         string      `json:"error"`
	DepositTxHash common.Hash `json:"stakeTxHash"`
}
type NodeDepositSdAllowanceResponse struct {
	Status    string   `json:"status"`
	Error     string   `json:"error"`
	Allowance *big.Int `json:"allowance"`
}

type CanRegisterValidatorsResponse struct {
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	CanDeposit               bool           `json:"CanDeposit"`
	InsufficientBalance      bool           `json:"insufficientBalance"`
	InvalidAmount            bool           `json:"invalidAmount"`
	OperatorNotRegistered    bool           `json:"operatorNotRegistered"`
	OperatorNotActive        bool           `json:"operatorNotActive"`
	DepositPaused            bool           `json:"depositPaused"`
	MaxValidatorLimitReached bool           `json:"maxValidatorLimitReached"`
	InputKeyLimitReached     bool           `json:"inputKeyLimitReached"`
	InputKeyLimit            uint16         `json:"inputKeyLimit"`
	GasInfo                  stader.GasInfo `json:"gasInfo"`
}

type ValidatorRegisterResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanNodeSendResponse struct {
	Status              string         `json:"status"`
	Error               string         `json:"error"`
	CanSend             bool           `json:"canSend"`
	InsufficientBalance bool           `json:"insufficientBalance"`
	GasInfo             stader.GasInfo `json:"gasInfo"`
}
type NodeSendResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type NodeSyncProgressResponse struct {
	Status   string              `json:"status"`
	Error    string              `json:"error"`
	EcStatus ClientManagerStatus `json:"ecStatus"`
	BcStatus ClientManagerStatus `json:"bcStatus"`
}

type ContractsInfoResponse struct {
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	Network                  uint64         `json:"network"`
	BeaconDepositContract    common.Address `json:"beaconDepositContract"`
	BeaconNetwork            uint64         `json:"beaconNetwork"`
	PermissionedNodeRegistry common.Address `json:"permissionedNodeRegistry"`
	VaultFactory             common.Address `json:"vaultFactory"`
	EthxToken                common.Address `json:"ethxToken"`
	SdToken                  common.Address `json:"sdToken"`
	SdCollateralContract     common.Address `json:"sdCollateralContract"`
	SocializingPoolContract  common.Address `json:"socializingPoolContract"`
	PermissionedPool         common.Address `json:"permissionedPool"`
	StaderOracle             common.Address `json:"staderOracle"`
	StaderConfig             common.Address `json:"staderConfig"`
	StakePoolManager         common.Address `json:"stakePoolManager"`
}

type DebugExitResponse struct {
	Status          string                   `json:"status"`
	Error           string                   `json:"error"`
	ValidatorPubKey types.ValidatorPubkey    `json:"validatorPubKey"`
	ExitEpoch       uint64                   `json:"exitEpoch"`
	CurrentEpoch    uint64                   `json:"currentEpoch"`
	ValidatorIndex  uint64                   `json:"validatorIndex"`
	SrHash          common.Hash              `json:"srHash"`
	SignedMsg       types.ValidatorSignature `json:"signedMsg"`
	SignatureDomain common.Hash              `json:"signatureDomain"`
}

type CanSendPresignedMsgResponse struct {
	Status                               string `json:"status"`
	Error                                string `json:"error"`
	OperatorNotRegistered                bool   `json:"operatorNotRegistered"`
	ValidatorNotRegisteredWithStader     bool   `json:"validatorNotRegisteredWithStader"`
	ValidatorNotRegisteredWithOperator   bool   `json:"validatorNotRegisteredWithOperator"`
	ValidatorNotRegistered               bool   `json:"validatorNotRegistered"`
	ValidatorPreSignKeyAlreadyRegistered bool   `json:"validatorPreSignKeyAlreadyRegistered"`
	ValidatorIsNotActive                 bool   `json:"validatorIsNotActive"`
}

type SendPresignedMsgResponse struct {
	Status          string                `json:"status"`
	Error           string                `json:"error"`
	ValidatorPubKey types.ValidatorPubkey `json:"validatorPubKey"`
	ExitEpoch       uint64                `json:"exitEpoch"`
	ValidatorIndex  uint64                `json:"validatorIndex"`
	SignedMsg       string                `json:"signedMsg"`
}

type CanExitValidatorResponse struct {
	Status                 string `json:"status"`
	Error                  string `json:"error"`
	OperatorNotRegistered  bool   `json:"operatorNotRegistered"`
	ValidatorNotRegistered bool   `json:"validatorNotRegistered"`
	ValidatorNotExists     bool   `json:"validatorNotExists"`
	ValidatorTooYoung      bool   `json:"validatorTooYoung"`
	ValidatorExiting       bool   `json:"validatorExiting"`
}

type ExitValidatorResponse struct {
	BeaconChainUrl string `json:"beaconChainUrl"`
	Status         string `json:"status"`
	Error          string `json:"error"`
}

type ValidatorInfoResponse struct {
	ValidatorInfo types.ValidatorContractInfo `json:"validatorInfo"`
	ClRewards     *big.Int                    `json:"clRewards"`
	DisplayStatus string                      `json:"displayStatus"`
	Status        string                      `json:"status"`
	Error         string                      `json:"error"`
}

type CanUpdateSocializeElResponse struct {
	Status                        string         `json:"status"`
	Error                         string         `json:"error"`
	OperatorNotRegistered         bool           `json:"operatorNotRegistered"`
	OperatorNotActive             bool           `json:"operatorNotActive"`
	SocializingPoolContractPaused bool           `json:"socializingPoolContractPaused"`
	AlreadyOptedIn                bool           `json:"alreadyOptedIn"`
	AlreadyOptedOut               bool           `json:"alreadyOptedOut"`
	InCooldown                    bool           `json:"inCooldown"`
	GasInfo                       stader.GasInfo `json:"gasInfo"`
}

type UpdateSocializeElResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanClaimClRewardsResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	OperatorNotRegistered bool           `json:"operatorNotRegistered"`
	VaultAlreadySettled   bool           `json:"vaultAlreadySettled"`
	NoClRewards           bool           `json:"noClRewards"`
	TooManyClRewards      bool           `json:"tooManyClRewards"`
	ValidatorNotFound     bool           `json:"validatorNotFound"`
	GasInfo               stader.GasInfo `json:"gasInfo"`
}

type ClaimClRewardsResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	ClRewardsAmount       *big.Int       `json:"clRewardsAmount"`
	OperatorRewardAddress common.Address `json:"operatorRewardAddress"`
	TxHash                common.Hash    `json:"txHash"`
}

type CanSettleExitFunds struct {
	Status                 string         `json:"status"`
	Error                  string         `json:"error"`
	OperatorNotRegistered  bool           `json:"operatorNotRegistered"`
	ValidatorNotWithdrawn  bool           `json:"validatorNotWithdrawn"`
	ValidatorNotRegistered bool           `json:"validatorNotRegistered"`
	NoEthToWithdraw        bool           `json:"notEthToWithdraw"`
	VaultAlreadySettled    bool           `json:"vaultAlreadySettled"`
	GasInfo                stader.GasInfo `json:"gasInfo"`
}

type SettleExitFunds struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	ExitAmount            *big.Int       `json:"exitShare"`
	OperatorRewardAddress common.Address `json:"operatorRewardAddress"`
	TxHash                common.Hash    `json:"txHash"`
}

type CanWithdrawElRewardsResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	NoElRewards           bool           `json:"noElRewards"`
	OperatorNotRegistered bool           `json:"operatorNotRegistered"`
	GasInfo               stader.GasInfo `json:"gasInfo"`
}

type WithdrawElRewardsResponse struct {
	Status                string         `json:"status"`
	Error                 string         `json:"error"`
	ElRewardsAmount       *big.Int       `json:"elRewardsAmount"`
	OperatorRewardAddress common.Address `json:"operatorRewardAddress"`
	TxHash                common.Hash    `json:"txHash"`
}

type CanRequestWithdrawSdResponse struct {
	Status                     string         `json:"status"`
	Error                      string         `json:"error"`
	OperatorNotRegistered      bool           `json:"operatorNotRegistered"`
	InsufficientSdCollateral   bool           `json:"insufficientSdCollateral"`
	InsufficientWithdrawableSd bool           `json:"insufficientWithdrawableSd"`
	GasInfo                    stader.GasInfo `json:"gasInfo"`
}

type RequestWithdrawSdResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanClaimSdResponse struct {
	Status                   string         `json:"status"`
	Error                    string         `json:"error"`
	NoExistingClaim          bool           `json:"noExistingClaim"`
	ClaimIsInUnbondingPeriod bool           `json:"claimIsInUnbondingPeriod"`
	GasInfo                  stader.GasInfo `json:"gasInfo"`
}

type ClaimSdResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanDownloadSpMerkleProofsResponse struct {
	Status                string  `json:"status"`
	Error                 string  `json:"error"`
	OperatorNotRegistered bool    `json:"operatorNotRegistered"`
	NoMissingCycles       bool    `json:"noMissingCycles"`
	MissingCycles         []int64 `json:"missingCycles"`
	CurrentCycle          int64   `json:"currentCycle"`
}

type DownloadSpMerkleProofsResponse struct {
	Status           string  `json:"status"`
	Error            string  `json:"error"`
	DownloadedCycles []int64 `json:"downloadedCycles"`
}

type CanClaimSpRewardsResponse struct {
	Status                        string     `json:"status"`
	Error                         string     `json:"error"`
	OperatorNotRegistered         bool       `json:"operatorNotRegistered"`
	SocializingPoolContractPaused bool       `json:"socializingPoolContractPaused"`
	IneligibleCycles              []*big.Int `json:"ineligibleCycles"`
	ClaimedCycles                 []*big.Int `json:"claimedCycles"`
	UnclaimedCycles               []*big.Int `json:"unclaimedCycles"`
	CyclesToDownload              []*big.Int `json:"cyclesToDownload"`
}

type EstimateClaimSpRewardsGasResponse struct {
	Status  string         `json:"status"`
	Error   string         `json:"error"`
	GasInfo stader.GasInfo `json:"gasInfo"`
}

type ClaimSpRewardsResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanUpdateOperatorName struct {
	Status              string         `json:"status"`
	Error               string         `json:"error"`
	OperatorNotActive   bool           `json:"operatorNotActive"`
	OperatorNameTooLong bool           `json:"operatorNameTooLong"`
	NothingToUpdate     bool           `json:"nothingToUpdate"`
	GasInfo             stader.GasInfo `json:"gasInfo"`
}

type UpdateOperatorName struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type CanUpdateOperatorRewardAddress struct {
	Status                    string         `json:"status"`
	Error                     string         `json:"error"`
	OperatorNotActive         bool           `json:"operatorNotActive"`
	OperatorRewardAddressZero bool           `json:"operatorRewardAddressZero"`
	NothingToUpdate           bool           `json:"nothingToUpdate"`
	GasInfo                   stader.GasInfo `json:"gasInfo"`
}

type UpdateOperatorRewardAddress struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	TxHash common.Hash `json:"txHash"`
}

type NodeSignResponse struct {
	Status     string `json:"status"`
	Error      string `json:"error"`
	SignedData string `json:"signedData"`
}

type CyclesDetailedInfo struct {
	Status             string                    `json:"status"`
	Error              string                    `json:"error"`
	DetailedCyclesInfo []DetailedMerkleProofInfo `json:"detailedCyclesInfo"`
}

type DetailedMerkleProofInfo struct {
	MerkleProofInfo stader_backend.CycleMerkleProofs `json:"merkleProofInfo"`
	CycleTime       time.Time                        `json:"cycleTime"`
}
