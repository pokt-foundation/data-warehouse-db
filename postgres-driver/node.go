package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/data-warehouse-db/types"
)

func (d *PostgresDriver) WriteNode(ctx context.Context, node types.Node) error {
	return d.Queries.InsertNode(ctx, InsertNodeParams{
		NodeAddress:          node.NodeAddress,
		NodeAverageLatency:   node.NodeAverageLatency,
		NodeUptimePercentage: node.NodeUptimePercentage,
		NodeErrorPercentage:  node.NodeErrorPercentage,
		IsAltruistNode:       node.IsAltruistNode,
		ChainID:              node.ChainID,
	})
}

func (d *PostgresDriver) ReadNode(ctx context.Context, nodeID int) (types.Node, error) {
	node, err := d.Queries.SelectNode(ctx, int32(nodeID))
	if err != nil {
		return types.Node{}, nil
	}

	return types.Node{
		NodeID:               node.NodeID,
		NodeAddress:          node.NodeAddress,
		NodeAverageLatency:   node.NodeAverageLatency,
		NodeUptimePercentage: node.NodeUptimePercentage,
		NodeErrorPercentage:  node.NodeErrorPercentage,
		IsAltruistNode:       node.IsAltruistNode,
		ChainID:              node.ChainID,
	}, nil
}
