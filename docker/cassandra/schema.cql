CREATE KEYSPACE IF NOT EXISTS tatanka WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

CREATE TABLE tatanka.trades (
    id uuid PRIMARY KEY,
    selector text,
    trade_id text,
    price double,
    size double,
    time timestamp,
    side text
);
