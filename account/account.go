package account

type Account struct {
	balance float64
	owner   string
}

func NewAccount(initialAmount float64, owner string) *Account {
	return &Account{initialAmount, owner}
}

func (account *Account) Deposit(amount float64) {
	account.balance += amount

}

func (account *Account) Withdraw(amount float64) {
  account.balance -= amount
}

func (account Account) Balance() float64 {
	return account.balance
}
