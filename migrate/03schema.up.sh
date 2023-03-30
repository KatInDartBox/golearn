#! /usr/bin/bash

source "./01declairVar.sh"

# create student schema
psql -q -d "$uDB" -f "./db/schema/up.sql"
