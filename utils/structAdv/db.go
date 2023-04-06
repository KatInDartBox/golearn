package structAdv

import "errors"

type Method interface {
	Widthraw(amount int) error
	Deposit(amount int)
	GetBalance() int
}

type DB struct {
	balance int
	fee     int
}

func New(balance int, fee int) *DB {
	return &DB{
		balance: balance,
		fee:     fee,
	}
}

func (db *DB) Widthraw(amount int) error {
	newBalance := db.balance - amount - db.fee

	if newBalance < 0 {
		return errors.New("sufficiant balance")
	}
	db.balance = newBalance
	return nil
}

func (db *DB) Deposit(amount int) {
	db.balance += amount
}
func (db *DB) GetBalance() int {
	return db.balance
}
