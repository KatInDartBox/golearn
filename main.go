package main

import (
	"api-test/db/sqlc"
	"database/sql"
	"fmt"
)

var test sql.DB
var qty sqlc.Queries

func main() {
	fmt.Println("start programming")
	x := false

	if !x {
		fmt.Println("ok is false")
	}
}
