CREATE KEYSPACE IF NOT EXISTS tatanka WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 3};

CREATE TABLE trades (
    id uuid PRIMARY KEY,
    selector text,
    trade_id text,
    price double,
    size double,
    time timestamp,
    side text
);
