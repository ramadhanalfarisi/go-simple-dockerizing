#!/bin/bash
set -e
export PGPASSWORD=postgres123;
psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "mangastore" <<-EOSQL
  CREATE DATABASE mangastore;
  GRANT ALL PRIVILEGES ON DATABASE mangastore TO "postgres";
EOSQL