-- name: GetEntryList :many
SELECT * FROM entry
where id>=sqlc.arg(from_id)::bigint AND id<=sqlc.arg(to_id)::bigint;

-- name: GetEntry :one
SELECT * FROM entry
where id=$1;

-- name: CreateEntry :one
INSERT INTO entry (
  account_id ,
  amount,
  note
) VALUES (
  $1, $2, $3
)
RETURNING *;