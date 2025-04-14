#!/bin/bash

# WARNING! This script is for local development only.

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    -- Create role if it doesn't exist
    DO \$\$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = '${POSTGRES_USER}') THEN
            CREATE ROLE ${POSTGRES_USER} WITH LOGIN PASSWORD '${POSTGRES_PASSWORD}';
        END IF;
    END
    \$\$;

    -- Create database and grant access
    DO \$\$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_database WHERE datname = '${$POSTGRES_DB}') THEN
            CREATE DATABASE ${POSTGRES_DB};
        END IF;
    END
    \$\$;

    GRANT ALL PRIVILEGES ON DATABASE ${$POSTGRES_DB} TO ${POSTGRES_USER};
EOSQL
