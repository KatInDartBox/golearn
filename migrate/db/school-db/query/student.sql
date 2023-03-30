-- -- name: GetStudent :one
-- SELECT * FROM student
-- WHERE id = $1 LIMIT 1;

-- name: ListStudent :many
SELECT * FROM student
ORDER BY name;

-- name: CreateStudent :one
INSERT INTO student (
  first_name ,
  last_name ,
  email  ,
  gender  ,
  phone  ,
  balance 
) VALUES (
  $1, $2, $3,$4, $5,$6
)
RETURNING *;

-- -- name: DeleteStudent :exec
-- DELETE FROM student
-- WHERE id = $1;

-- -- name: UpdateStudent :one
-- UPDATE student
-- set
--   first_name =$2,
--   last_name =$3,
--   email =$4,
--   gender =$5,
--   phone =$6,
--   balance =$7
-- WHERE id =$1
-- RETURNING *;