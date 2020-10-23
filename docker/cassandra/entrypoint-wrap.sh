#!/bin/bash
if [[ ! -z "$CASSANDRA_KEYSPACE" && $1 = 'cassandra' ]]; then
  # Create default keyspace for single node cluster
  CQL="CREATE KEYSPACE $CASSANDRA_KEYSPACE WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}; CREATE TABLE $CASSANDRA_KEYSPACE.trades (id uuid PRIMARY KEY, selector text, trade_id text, price double, size double, time timestamp, side text);"
  until echo $CQL | cqlsh; do
    echo "cqlsh: Cassandra is unavailable - retry later"
    sleep 2
  done &
fi

exec /usr/local/bin/docker-entrypoint.sh "$@"
