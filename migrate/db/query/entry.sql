
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

-- name: GetEntryList :many
SELECT ac.id,ac.holder,ac.balance,
en.amount,en.note
FROM "account" as ac 
INNER JOIN "entry" as en
on ac.id = en."account_id"
where ac.id = sqlc.arg(account_id)::bigint
and en.id > sqlc.arg(entry_last_id)::bigint
limit sqlc.arg(page_size);

/*
ac.id,ac.holder,ac.balance,
which is redundance 
*/