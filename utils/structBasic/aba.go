package structBasic

import "errors"

type Aba struct {
	balance int
}

func NewIaba() Itransaction {
	var acc Itransaction
	acc = NewAba()
	return acc
}

func NewAba() *Aba {
	return &Aba{
		balance: 0,
	}
}

func (b *Aba) Widthraw(amount int) (int, error) {
	newBalance := b.balance - amount
	if newBalance < 0 {
		return 0, errors.New("insuficient balance")
	}
	b.balance = newBalance
	return amount, nil
}

func (b *Aba) Deposit(amount int) int {
	b.balance += amount
	return b.balance
}

func (b *Aba) GetBalance() int {
	return b.balance
}
