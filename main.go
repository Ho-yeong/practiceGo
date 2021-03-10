package main

import (
	"fmt"

	"github.com/Ho-yeong/practiceGo/accounts"
	"github.com/Ho-yeong/practiceGo/mydict"
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

	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

	// dictionary
	dictionary := mydict.Dictionary{"first": "first word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	word := "hello"
	def := "Greeting"

	err2 := dictionary.Add(word, def)

	if err2 != nil {
		fmt.Println(err)
	}
	definition2, err := dictionary.Search(word)
	fmt.Println(definition2)

	err3 := dictionary.Update(word, "modified")
	if err3 != nil {
		fmt.Println(err3)
	}
	result, _ := dictionary.Search(word)
	fmt.Println(result)

	dictionary.Delete(word)
	result2, err4 := dictionary.Search(word)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(result2)
	}
}
