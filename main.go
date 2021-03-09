package main

import (
	"fmt"

	"github.com/Ho-yeong/practiceGo/accounts"
)

func main() {
	account := accounts.NewAccount("simon")
	account.Deposit(1000)
	fmt.Println(account.Balance())

	err := account.Withdraw(20000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

}
