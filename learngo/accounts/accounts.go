package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string // 대문자시작해야 public
	balance int
}

var errNoMoney = errors.New("Can't withdraw") // error 변수명 앞에 err 붙여주기

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) /*Reciver *붙이지 않으면 Object를 복사하므로 값이 바뀌지 않는다 */ Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney // return errors.New("Can't withdraw")
	}
	a.balance -= amount
	return nil // error 가 아니다
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string{
	return fmt.Sprint(a.Owner(), "'s account.\n Has:", a.Balance())
}