#! /usr/bin/bash

#env
isProd=false
os="lin"

#rootDB
rDB="postgresql://postgres:12345@localhost:5432/postgres"

#userDB
uRole="tester"
uPass="admin"
uDBname="dbaccounting"

if [ "$isProd" == "true" ]; then
  #rootDB

  #userDB
  uRole="tester"
  uPass="admin"
  uDBname="uvschool"
fi

uDB="postgresql://$uRole:$uPass@localhost:5432/$uDBname"

tables=("account" "entry")
seqID=("account")

# script var
scrPath=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)
