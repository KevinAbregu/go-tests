package main

import "errors"

// Creacion del objeto y su estructura
type BitcointAccount struct {
	balance int
	fee     int
}

// Constructor para crear instancias de BitcointAccount. Retorna un puntero al nuevo objeto inicializado con los valores indicados.
func NewBitcointAccount() *BitcointAccount {
	return &BitcointAccount{
		balance: 0,
		fee:     300,
	}
}

// Las funciones se implementan con la lógica deseada (Importante que se envía un puntero del objeto (ya que es lo que recibimos del constructor))
func (b *BitcointAccount) GetBalance() int {
	return b.balance
}

func (b *BitcointAccount) Deposit(amount int) {
	b.balance += amount

}

func (b *BitcointAccount) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}
