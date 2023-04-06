-- name: GetAccountList :many
select * from account 
where id > sqlc.arg(last_id)::bigint 
limit sqlc.arg(page_limit)::bigint ;


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




-- array list work with "github.com/lib/pq" by default
-- select * from account where id=ANY($1::int[]);