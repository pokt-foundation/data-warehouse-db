package postgresdriver

import (
	"context"
	"testing"

	"github.com/pokt-foundation/data-warehouse-db/types"
	"github.com/stretchr/testify/suite"
)

const (
	connectionString = "postgres://postgres:pgpassword@localhost:5432/postgres?sslmode=disable" // pragma: allowlist secret
)

type PGDriverTestSuite struct {
	suite.Suite
	connectionString string
	driver           *PostgresDriver
	firsNode         types.Node
	firstNodeChain   types.NodeChain
}

func Test_RunPGDriverSuite(t *testing.T) {
	testSuite := new(PGDriverTestSuite)
	testSuite.connectionString = connectionString

	suite.Run(t, testSuite)
}

// SetupSuite runs before each test suite run
func (ts *PGDriverTestSuite) SetupSuite() {
	ts.NoError(ts.initPostgresDriver())

	ts.NoError(ts.driver.WriteNode(context.Background(), types.Node{
		NodeAddress:          "21",
		NodeAverageLatency:   21,
		NodeUptimePercentage: 21,
		NodeErrorPercentage:  21,
		IsAltruistNode:       false,
		ChainID:              21,
	}))

	ts.NoError(ts.driver.WriteNodeChain(context.Background(), types.NodeChain{
		NodeAddress:  "21",
		CurrentStake: 21,
		ChainID:      21,
	}))

	node, err := ts.driver.ReadNode(context.Background(), 1)
	ts.NoError(err)
	ts.firsNode = node

	nodeChain, err := ts.driver.ReadNodeChain(context.Background(), 1)
	ts.NoError(err)
	ts.firstNodeChain = nodeChain
}

// Initializes a real instance of the Postgres driver that connects to the test Postgres Docker container
func (ts *PGDriverTestSuite) initPostgresDriver() error {
	driver, err := NewPostgresDriver(ts.connectionString)
	if err != nil {
		return err
	}
	ts.driver = driver

	return nil
}
