package structBasic

import "errors"

type Acl struct {
	balance int
	fee     int
}

func NewIacl() Itransaction {
	var acc Itransaction
	acc = NewAcl()
	return acc
}

func NewAcl() *Acl {
	return &Acl{
		balance: 0,
		fee:     5,
	}
}

func (b *Acl) Widthraw(amount int) (int, error) {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return 0, errors.New("insuficient balance")
	}
	b.balance = newBalance
	return amount, nil
}

func (b *Acl) Deposit(amount int) int {
	b.balance += amount
	return b.balance
}

func (b *Acl) GetBalance() int {
	return b.balance
}
