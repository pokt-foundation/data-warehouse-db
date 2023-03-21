package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/data-warehouse-db/types"
)

func (ts *PGDriverTestSuite) TestPostgresDriver_WriteNode() {
	tests := []struct {
		name string
		node types.Node
		err  error
	}{
		{
			name: "Success",
			node: types.Node{
				NodeAddress:          "22",
				NodeAverageLatency:   21,
				NodeUptimePercentage: 21,
				NodeErrorPercentage:  21,
				IsAltruistNode:       false,
				ChainID:              21,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		ts.Equal(ts.driver.WriteNode(context.Background(), tt.node), tt.err)
	}
}

func (ts *PGDriverTestSuite) TestPostgresDriver_ReadNode() {
	tests := []struct {
		name    string
		nodeID  int
		expNode types.Node
		err     error
	}{
		{
			name:    "Success",
			nodeID:  int(ts.firsNode.NodeID),
			expNode: ts.firsNode,
			err:     nil,
		},
	}
	for _, tt := range tests {
		node, err := ts.driver.ReadNode(context.Background(), tt.nodeID)
		ts.Equal(err, tt.err)
		ts.Equal(node, tt.expNode)
	}
}
