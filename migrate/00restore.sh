#! /usr/bin/bash

source "./01declairVar.sh"

# need only role and db
source "./02role.down.sh"
source "./02role.up.sh"

#unzip db

gzip -d "./db/seed/db.sql.gz"
# get back db.sql

# restore db
psql -q -d "$uDB" -f "./db/seed/db.sql"
