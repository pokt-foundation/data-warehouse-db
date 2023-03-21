package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/data-warehouse-db/types"
)

func (ts *PGDriverTestSuite) TestPostgresDriver_WriteNodeChain() {
	tests := []struct {
		name      string
		nodeChain types.NodeChain
		err       error
	}{
		{
			name: "Success",
			nodeChain: types.NodeChain{
				NodeAddress:  "22",
				CurrentStake: 21,
				ChainID:      21,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		ts.Equal(ts.driver.WriteNodeChain(context.Background(), tt.nodeChain), tt.err)
	}
}

func (ts *PGDriverTestSuite) TestPostgresDriver_ReadNodeChain() {
	tests := []struct {
		name         string
		nodeChainID  int
		expNodeChain types.NodeChain
		err          error
	}{
		{
			name:         "Success",
			nodeChainID:  int(ts.firstNodeChain.NodeChainID),
			expNodeChain: ts.firstNodeChain,
			err:          nil,
		},
	}
	for _, tt := range tests {
		nodeChain, err := ts.driver.ReadNodeChain(context.Background(), tt.nodeChainID)
		ts.Equal(err, tt.err)
		ts.Equal(nodeChain, tt.expNodeChain)
	}
}
