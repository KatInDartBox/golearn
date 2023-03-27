#! /usr/bin/bash

source "./01declairVar.sh"

# drop down db first then role
# drop db
dropDB="drop database if exists $uDBname;"
dropRole="drop role if exists $uRole;"

# drop role
if [ "$os" == "win" ]; then
  psql -q -d "$rDB" -c "$dropDB"
  psql -q -d "$rDB" -c "$dropRole"
else
  sudo -u postgres psql -q -c "$dropDB"
  sudo -u postgres psql -q -c "$dropRole"
fi
