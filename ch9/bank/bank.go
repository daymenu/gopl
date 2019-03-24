package bank

var balance int = 0

func Deposit(deposit int) {
	balance += deposit
}

func Balance() int {
	return balance
}
