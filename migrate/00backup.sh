#! /usr/bin/bash

source "./01declairVar.sh"

rm -rf "./db/seed/db.sql.gz"
rm -rf "./db/seed/db.sql"

pg_dump -O "$uDB" -f "./db/seed/db.sql"
gzip "./db/seed/db.sql"

# pg_dump -O -Fc "$uDB" -f "./db/seed/db.dump"
