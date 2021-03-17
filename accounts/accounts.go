package accounts

import "errors"

// account struct
type Account struct {
	owner   string
	balance int
	Address string
}

// NewAccount creates Account
func NewAccount(owner string, address string) *Account {
	account := Account{owner: owner, balance: 0, Address: address}
	return &account
}

// Deposit X amount on your account
// go는 복사본을 만드려는 성격이 있기때문에, 아스타(*)의 존재 유무가 굉장히 중요하다
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your balance
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't withdraw you are poor")

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("Can't withdraw you are poor")
		return errNoMoney
	}
	a.balance -= amount

	return nil // nil is null or none
}
