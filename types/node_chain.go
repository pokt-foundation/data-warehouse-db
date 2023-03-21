package types

type NodeChain struct {
	NodeChainID  int32  `json:"nodeChainID"`
	NodeAddress  string `json:"nodeAddress"`
	CurrentStake int64  `json:"currentStake"`
	ChainID      int32  `json:"chainID"`
}
