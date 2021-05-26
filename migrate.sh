#!/bin/sh
# Database URL
DSN=postgres://postgres:$DBPWD@localhost:5432/gestion_citas_api?sslmode=disable

echo "Running migrations"
c:/go-migrate/migrate.windows-amd64.exe -source file://migrations/versions -database $DSN $1 $2
