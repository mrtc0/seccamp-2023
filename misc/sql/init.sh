#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE TABLE books (
    id integer,
    title varchar(255)
  );
  
  INSERT INTO books VALUES (1, 'Security Camp 101');
EOSQL
