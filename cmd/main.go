package main

import (
	"fmt"

	"github.com/DilshodK0592/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}

	account, err := svc.RegisterAccount("+992000000001")
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(account)
	err = svc.Deposit(account.ID, -10)
	if err != nil {
		switch err{
		case wallet.ErrAccountNotFound:
			fmt.Println("аккаунт не найден")
		case wallet.ErrAmountMustBePositive:
			fmt.Println("сумма должна быть положительной")
		}
		return 
	}
	fmt.Println(account.Balance) //10
}

