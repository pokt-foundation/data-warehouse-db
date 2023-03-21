package types

type Node struct {
	NodeID               int32  `json:"nodeID"`
	NodeAddress          string `json:"nodeAddress"`
	NodeAverageLatency   int32  `json:"nodeAverageLatency"`
	NodeUptimePercentage int32  `json:"nodeUptimePercentage"`
	NodeErrorPercentage  int32  `json:"nodeErrorPercentage"`
	IsAltruistNode       bool   `json:"isAltruistNode"`
	ChainID              int32  `json:"chainID"`
}
