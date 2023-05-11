package web3signer_client

type DepositInfo struct {
	PublicKey           string `json:"pubkey"`
	WithdrawCredentials string `json:"withdrawal_credentials"`
	Amount              string `json:"amount"`
	GenesisForkVersion  string `json:"genesis_fork_version"`
}

type DepositDataSignatureRequest struct {
	Type    string      `json:"type"`
	Deposit DepositInfo `json:"deposit"`
}

type DepositDataSignatureResponse struct {
	Signature string `json:"signature"`
}

type SigningForkInfo struct {
	Fork struct {
		PreviousVersion string `json:"previous_version"`
		CurrentVersion  string `json:"current_version"`
		Epoch           string `json:"epoch"`
	} `json:"fork"`
	GenesisValidatorRoot string `json:"genesis_validators_root"`
}

type SigningVoluntaryExit struct {
	Epoch          string `json:"epoch"`
	ValidatorIndex string `json:"validator_index"`
}

type VoluntaryExitMessageSignatureRequest struct {
	Type              string               `json:"type"`
	ForkInfoData      SigningForkInfo      `json:"fork_info"`
	VoluntaryExitData SigningVoluntaryExit `json:"voluntary_exit"`
}
