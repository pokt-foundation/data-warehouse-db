-- name: InsertNode :exec
INSERT INTO node (
    node_address,
    node_average_latency,
    node_uptime_percentage,
    node_error_percentage,
    is_altruist_node,
    chain_id
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);
-- name: InsertNodeChain :exec
INSERT INTO node_chain (
    node_address,
    current_stake,
    chain_id
)
VALUES (
    $1,
    $2,
    $3
);
-- name: SelectNode :one
SELECT
    node_id,
    node_address,
    node_average_latency,
    node_uptime_percentage,
    node_error_percentage,
    is_altruist_node,
    chain_id
FROM node
WHERE node_id = $1;
-- name: SelectNodeChain :one
SELECT
    node_chain_id,
    node_address,
    current_stake,
    chain_id
FROM node_chain
WHERE node_chain_id = $1;

