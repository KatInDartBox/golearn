package initer

import (
	"database/sql"
	"learn/db/sqlc"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	driver = "pgx"
	dns    = "postgresql://tester:admin@localhost:5432/dbaccounting"
)

var DB *sql.DB
var Qry *sqlc.Queries

func InitDB() (*sql.DB, *sqlc.Queries) {
	var err error
	DB, err = sql.Open(driver, dns)
	if err != nil {
		log.Fatalf("cant connect to db %v\n", err)
	}
	Qry = sqlc.New(DB)

	return DB, Qry
}
