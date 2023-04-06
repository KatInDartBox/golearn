package main

import (
	"context"
	"database/sql"
	"learn/db/learn/initialDB"
	"learn/db/sqlc"

	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB
var Qry *sqlc.Queries

func init() {
	DB, Qry = initialDB.Initial()

}

func main() {

	args := []sTransfer{
		{
			frId:   3,
			toId:   1,
			amount: 50,
		},
		{
			frId:   3,
			toId:   1,
			amount: 50,
		},
		{
			frId:   3,
			toId:   1,
			amount: 50,
		},
		{
			frId:   3,
			toId:   1,
			amount: 50,
		},
		{
			frId:   3,
			toId:   1,
			amount: 50,
		},
		{
			frId:   3,
			toId:   1,
			amount: 50,
		},

		{
			frId:   1,
			toId:   3,
			amount: 50,
		},
		{
			frId:   1,
			toId:   3,
			amount: 50,
		},
		{
			frId:   1,
			toId:   3,
			amount: 50,
		},
		{
			frId:   1,
			toId:   3,
			amount: 50,
		},
		{
			frId:   1,
			toId:   3,
			amount: 50,
		},
		{
			frId:   1,
			toId:   3,
			amount: 50,
		},
	}

	chErr := make(chan error)

	for i := range args {
		arg := args[i]
		go func() {
			err := accTransfer(arg)
			chErr <- err
		}()
	}

	for _, arg := range args {
		err := <-chErr
		if err != nil {
			panicErr(err)
			fmt.Println(arg)
		}
	}

	/*
		a1:550
		a3:850
		row id: 55
		exp:
			a1:550
			a3:850

	*/

	// argEnty := sEntry{
	// 	id:     1,
	// 	amount: 500,
	// 	note:   "add 500",
	// }
	// addEntry(argEnty, Qry)
}

type sTransfer struct {
	frId   int64
	toId   int64
	amount int64
}

func accTransfer(arg sTransfer) error {
	var err error
	tx, err := DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	qtx := Qry.WithTx(tx)
	frId, toId, amount := arg.frId, arg.toId, arg.amount
	// var frId,toId,amount int64 = 1,3,100
	note := fmt.Sprintf("transfer %v", amount)

	frArg := sEntry{
		id:     frId,
		amount: -amount,
		note:   note,
	}
	toArg := sEntry{
		id:     toId,
		amount: amount,
		note:   note,
	}

	/*
		to avoid db dead lock
		we need to execute db query
		in the same order, meaning that
		in our case we try to update balance by where id=$1 clause,
		so we need to execute it inorder of id asc or desc.

		we check if toId > than frId then
		we need to reverse the execution order
	*/
	if toId > frId {
		m := frArg
		frArg = toArg
		toArg = m
	}

	//exe frId first as it smaller
	err = addEntry(frArg, qtx)
	if err != nil {
		return err
	}

	err = addEntry(toArg, qtx)
	if err != nil {
		return err
	}

	// if frId < toId {
	// 	//exe frId first as it smaller
	// 	err = addEntry(frArg, qtx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = addEntry(toArg, qtx)
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	// 	//exe toId first as it smaller
	// 	err = addEntry(toArg, qtx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = addEntry(frArg, qtx)
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	return tx.Commit()
}

type sEntry struct {
	id     int64
	amount int64
	note   string
}

func addEntry(arg sEntry, qtx *sqlc.Queries) error {
	var err error
	var accId int64 = arg.id
	var amount int64 = arg.amount

	// tx, err := DB.Begin()
	// if err != nil {return err}

	// qtx := Qry.WithTx(tx)

	argEty := sqlc.CreateEntryParams{
		AccountID: accId,
		Note:      arg.note,
		Amount:    amount,
	}

	_, err = qtx.CreateEntry(context.Background(), argEty)
	if err != nil {
		return err
	}

	// err = errors.New("entry error: ")
	if err != nil {
		return err
	}

	argUpd := sqlc.UpdateAccountBalanceParams{
		ID:      accId,
		Balance: amount,
	}
	_, err = qtx.UpdateAccountBalance(context.Background(), argUpd)
	if err != nil {
		return err
	}

	return err

}

func exeTxtest() {
	err := exeTransaction(context.Background(), DB, func(q *sqlc.Queries) error {
		var err error
		var accId int64 = 1
		var amount int64 = 100

		acc1, err := q.GetAccount(context.Background(), accId)
		if err != nil {
			return err
		}
		logDB("acc1 before: ", acc1)

		argEty := sqlc.CreateEntryParams{
			AccountID: accId,
			Note:      "add 100",
			Amount:    amount,
		}

		ety, err := q.CreateEntry(context.Background(), argEty)

		logDB("create entry", ety)

		argUpd := sqlc.UpdateAccountBalanceParams{
			ID:      accId,
			Balance: amount,
		}
		acc2, err := q.UpdateAccountBalance(context.Background(), argUpd)
		if err != nil {
			return err
		}
		logDB("acc1 after: ", acc2)
		return err
	})

	panicErr(err)
}

func exeTransaction(ctx context.Context, db *sql.DB, fn func(q *sqlc.Queries) error) error {

	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := Qry.WithTx(tx)

	err = fn(qtx)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func createEntry() error {
	var err error
	var accId int64 = 1
	var amount int64 = 100

	acc1, err := Qry.GetAccount(context.Background(), accId)
	if err != nil {
		return err
	}
	logDB("acc1 before: ", acc1)

	tx, err := DB.Begin()

	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := Qry.WithTx(tx)

	argEty := sqlc.CreateEntryParams{
		AccountID: 1,
		Note:      "add 100",
		Amount:    amount,
	}

	ety, err := qtx.CreateEntry(context.Background(), argEty)

	logDB("create entry", ety)

	argUpd := sqlc.UpdateAccountBalanceParams{
		ID:      1,
		Balance: 100,
	}
	acc, err := qtx.UpdateAccountBalance(context.Background(), argUpd)
	if err != nil {
		return err
	}
	logDB("update account", acc)

	return tx.Commit()

}

func getAcc(id int64) sqlc.Account {
	acc, err := Qry.GetAccount(context.Background(), id)
	panicErr(err)
	return acc
}

func createAccount() {
	arg := sqlc.CreateAccountParams{
		Holder:  "holder13",
		Balance: 400}
	acc, err := Qry.CreateAccount(context.Background(), arg)
	panicErr(err)
	logDB("create account", acc)

	argList := sqlc.GetAccountListParams{
		LastID:    0,
		PageLimit: 10}
	accList, err := Qry.GetAccountList(context.Background(), argList)
	panicErr(err)
	logDB("get account list", accList)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
func logDB(msg string, v any) {
	fmt.Println(msg)
	fmt.Println("---------------")
	fmt.Println(v)
	fmt.Println("---------------")
}
