#!/bin/sh
# wait-for.sh

host="$1"
shift
port=5432

until nc -z "$host" $port; do
  echo "Waiting for PostgreSQL at $host:$port..."
  sleep 2
done

exec "$@"

