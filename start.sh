#!/bin/sh

set -e 

echo "Run DB Migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"