package main

import "fmt"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	myAccounts := []IBankAccount{
		NewWellsFargo(),
		NewBitcointAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withdraw(70); err != nil {
			fmt.Printf("ERR: %v\n", err)
		}
		balance := account.GetBalance()
		fmt.Printf("balance = %v\n", balance)
	}
}
