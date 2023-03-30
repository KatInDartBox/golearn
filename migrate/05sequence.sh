#! /usr/bin/bash

source "./01declairVar.sh"

for tb in "${seqID[@]}"; do
  seq="SELECT setval('${tb}_id_seq', (SELECT MAX(id) FROM ${tb}));"

  psql -q -d "$uDB" -c "$seq"

done
