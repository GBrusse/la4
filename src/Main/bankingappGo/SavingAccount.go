package bankingappGo

import (
	"fmt"
	"time"
)

type savingAccount struct {
	AccountNumber int
	Balance       float64
	InterestRate  float64
}

func (sa *savingAccount) Deposit(amount float64) {
	if amount > 0 {
		sa.Balance += amount
		fmt.Println("Deposit successful")
	} else {
		fmt.Println("Invalid amount")
	}
}

func (sa *savingAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= sa.Balance {
		sa.Balance -= amount
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Invalid amount or insufficient funds")
	}
}

func (sa *savingAccount) AddInterest() {
	periods := 12 // assuming monthly compounding
	rate := sa.InterestRate / float64(periods)
	interest := sa.Balance * rate
	sa.Balance += interest
	fmt.Printf("Added $%.2f interest\n", interest)
}

func (sa *savingAccount) Accrue(interestRate float64, duration time.Duration, interestChannel chan float64) {
	defer close(interestChannel)

	periodsPerYear := 12 // assuming monthly compounding
	rate := interestRate / float64(periodsPerYear)
	interest := sa.Balance * rate * float64(duration) / (24 * time.Hour).Seconds() // assuming interest is calculated using actual number of seconds in duration
	interestChannel <- interest
}
