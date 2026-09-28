package main

import (
	"errors"
	"fmt"
)

var notSufficentBalnce = errors.New("No sufficent balcnace")

type Account struct {
	Id      int
	Name    string
	Balance int
}

func (a *Account) deposite(amount int) int {

	fmt.Println("Update", a.Balance, "+", amount, "--------->", a.Balance+amount)
	a.Balance = a.Balance + amount

	return a.Balance
}
func (a *Account) WithDraw(amount int) error {
	if a.Balance < amount {
		return notSufficentBalnce
	} else {
		a.Balance = a.Balance - amount
		return nil
	}
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func III() {
	acc1 := Account{
		Id:      1,
		Name:    "Ranjit",
		Balance: 0,
	}
	for i := range 100 {
		fmt.Println(i)
		go acc1.deposite(10)
	}

	fmt.Println("Currunet Balance ", acc1.Balance)
}
