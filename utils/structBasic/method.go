package structBasic

type Itransaction interface {
	Widthraw(amount int) (int, error)
	Deposit(amount int) int
	GetBalance() int
}
