package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withraw")

// NewAccount constructor ( creates Account )
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// () 안의 친구는 reciever 라고 부름
// reciever의 이름은 보통 struct의 첫 소문자를 따름
// * 은 method를 call 한 struct의 주소를 가져와서 사용하라는 뜻
// * 을 붙이지 않으면 복사본을 만들어서 처리한다 (반영이 안됨)

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance 함수는 값을 보여주기만 할 뿐이라서 반영이 안되도 상관없음

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw of your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
