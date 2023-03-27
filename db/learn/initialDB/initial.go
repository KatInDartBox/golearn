package initialDB

import (
	"api-test/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	driver = "pgx"
	dns    = "postgresql://tester:admin@localhost:5432/dbaccounting"
)

func Initial() (DB *sql.DB, Qty *sqlc.Queries) {
	var err error
	DB, err = sql.Open(driver, dns)
	if err != nil {
		log.Fatalf("cant connect to db %v\n", err)
	}
	Qty = sqlc.New(DB)

	return DB, Qty
}
