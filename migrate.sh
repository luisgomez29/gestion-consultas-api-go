#!/bin/sh
# Database URL
DSN=postgres://postgres:$DBPWD@localhost:5432/gestion_citas_api?sslmode=disable

# Up or down migrations
UP_OR_DOWN=$1

# Apply N up or down migrations
STEPS=$2

echo "Running migrations"
c:/go-migrate/migrate.windows-amd64.exe -source file://migrations/versions -database $DSN $UP_OR_DOWN $STEPS
