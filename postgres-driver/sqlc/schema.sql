CREATE  TABLE node (
	node_id     integer  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
    node_address varchar NOT NULL  UNIQUE,
    node_average_latency int      NOT NULL  ,
    node_uptime_percentage int NOT NULL ,
    node_error_percentage int NOT NULL ,
    is_altruist_node boolean NOT NULL ,
    chain_id int NOT NULL ,
	CONSTRAINT pk_node PRIMARY KEY ( node_id )
 );

CREATE  TABLE node_chain (
	node_chain_id     integer  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
    node_address varchar NOT NULL  UNIQUE,
    current_stake bigint NOT NULL ,
    chain_id int NOT NULL ,
	CONSTRAINT pk_node_chain PRIMARY KEY ( node_chain_id )
 );
