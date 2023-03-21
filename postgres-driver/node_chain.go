package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/data-warehouse-db/types"
)

func (d *PostgresDriver) WriteNodeChain(ctx context.Context, nodeChain types.NodeChain) error {
	return d.InsertNodeChain(ctx, InsertNodeChainParams{
		NodeAddress:  nodeChain.NodeAddress,
		CurrentStake: nodeChain.CurrentStake,
		ChainID:      nodeChain.ChainID,
	})
}

func (d *PostgresDriver) ReadNodeChain(ctx context.Context, nodeChainID int) (types.NodeChain, error) {
	nodeChain, err := d.SelectNodeChain(ctx, int32(nodeChainID))
	if err != nil {
		return types.NodeChain{}, err
	}

	return types.NodeChain{
		NodeChainID:  nodeChain.NodeChainID,
		NodeAddress:  nodeChain.NodeAddress,
		CurrentStake: nodeChain.CurrentStake,
		ChainID:      nodeChain.ChainID,
	}, nil
}
