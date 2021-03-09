package main

import (
	"fmt"

	"github.com/Ho-yeong/practiceGo/banking"
)

func main() {
	account := banking.Account{Owner: "simon"}

	fmt.Println(account)
}
