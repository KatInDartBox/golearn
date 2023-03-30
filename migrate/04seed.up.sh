#! /usr/bin/bash

source "./01declairVar.sh"

# create student schema
# psql -q -d "$uDB" -f "./db/seed/student.sql"
# psql -q -d "$uDB" -f "./db/seed/teacher.sql"
# psql -q -d "$uDB" -f "./db/seed/subject.sql"
# psql -q -d "$uDB" -f "./db/seed/subject_student.sql"

for tb in "${tables[@]}"; do
  psql -q -d "$uDB" -f "./db/seed/$tb.sql"
done
