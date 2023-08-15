package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Can't withdraw you are poor")

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// String is the method that is called when printing with the %s format
// It is the same as toString() in Java
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account.\nHas: ", a.balance)
}
