-- name: GetAccountList :many
SELECT * FROM account
where id>=sqlc.arg(from_id)::bigint AND id<=sqlc.arg(to_id)::bigint;

-- name: GetAccount :one
SELECT * FROM account
where id=$1;

-- name: UpdateAccountBalance :one
update account set 
balance=balance + sqlc.arg(balance)::bigint 
where id=sqlc.arg(id)::bigint  returning *;

-- name: CreateAccount :one
INSERT INTO account (
  holder ,
  balance 
) VALUES (
  $1, $2
)
RETURNING *;