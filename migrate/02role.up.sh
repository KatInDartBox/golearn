#! /usr/bin/bash

source "./01declairVar.sh"
source "./02role.down.sh"

# create role
createRole="CREATE ROLE $uRole WITH CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD '$uPass';"

if [ "$os" == "win" ]; then
  psql -q -d "$rDB" -c "$createRole"
else
  # make sure linux user "postgres" is create
  sudo -u postgres psql -q -c "$createRole"
fi

# create db
createDB="create database $uDBname"
uDB="postgresql://$uRole:$uPass@localhost:5432/postgres"
psql -q -d "$uDB" -c "$createDB"

# reset uDB
uDB="postgresql://$uRole:$uPass@localhost:5432/$uDBname"
