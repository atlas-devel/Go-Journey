package main

import "fmt"

type BankAccount struct {
	owner   string
	balance float64
}

func (d BankAccount) Deposit(addMoney float64) string {
	if addMoney <= 0 {
		return "Unsupported funds"
	}
	addDeposit := d.balance + addMoney
	return fmt.Sprintf("Your account has been charged with %f", addDeposit)

}

func (a BankAccount) Withdraw(withdrawMoney float64) string {
	if withdrawMoney <= 0 {
		return "Unsupported funds"
	}
	result := a.balance - withdrawMoney
	if result <= 0 {
		return "insuffient balance"
	}
	return fmt.Sprintf("withdraw success, your new balance is %f", result)
}

func (b BankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	account := BankAccount{
		owner:   "Leon",
		balance: 12000,
	}

	fmt.Println(account.Deposit(200))
	// fmt.Print(account.Withdraw(-10))
	//fmt.Print(account.GetBalance())
	fmt.Println(account.balance)
}
