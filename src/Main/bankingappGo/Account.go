package bankingappGo

import (
	"fmt"
	"time"
)

type Account interface {
	Accrue(interestRate float64, duration time.Duration, interestChannel chan<- float64)
}

type CheckingAccount struct {
	accountNumber string
	balance       float64
	interestRate  float64
}

func (c *CheckingAccount) getAccountNumber() string {
	return c.accountNumber
}

func (c *CheckingAccount) getBalance() float64 {
	return c.balance
}

func (c *CheckingAccount) getInterestRate() float64 {
	return c.interestRate
}

func (c *CheckingAccount) setInterestRate(rate float64) {
	c.interestRate = rate
}

func (c *CheckingAccount) deposit(amount float64) {
	c.balance += amount
}

func (c *CheckingAccount) withdraw(amount float64) error {
	if amount > c.balance {
		return fmt.Errorf("Insufficient funds")
	}
	c.balance -= amount
	return nil
}

func (c *CheckingAccount) transfer(amount float64, toAccount *CheckingAccount) error {
	if amount > c.balance {
		return fmt.Errorf("Insufficient funds")
	}
	c.balance -= amount
	toAccount.balance += amount
	return nil
}

func newCheckingAccount(accountNumber string, balance float64, interestRate float64) *CheckingAccount {
	return &CheckingAccount{
		accountNumber: accountNumber,
		balance:       balance,
		interestRate:  interestRate,
	}
}

func (acc *CheckingAccount) Accrue(interestRate float64, duration time.Duration, interestChannel chan<- float64) {
	interest := acc.balance * interestRate * float64(duration) / (365.0 * 24.0 * float64(time.Hour))
	interestChannel <- interest
}

type SavingsAccount struct {
	accountNumber int
	balance       float64
	interestRate  float64
}

func newSavingsAccount(accountNumber int, balance float64, interestRate float64) *SavingsAccount {
	return &SavingsAccount{
		accountNumber: accountNumber,
		balance:       balance,
		interestRate:  interestRate,
	}
}

func (sa *SavingsAccount) Deposit(amount float64) {
	if amount > 0 {
		sa.balance += amount
		fmt.Println("Deposit successful")
	} else {
		fmt.Println("Invalid amount")
	}
}

func (sa *SavingsAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= sa.balance {
		sa.balance -= amount
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Invalid amount or insufficient funds")
	}
}

func (sa *SavingsAccount) AddInterest() {
	periods := 12 // assuming monthly compounding
	rate := sa.interestRate / float64(periods)
	interest := sa.balance * rate
	sa.balance += interest
}

func (sa *SavingsAccount) Accrue(interestRate float64, duration time.Duration, interestChannel chan<- float64) {
	interest := sa.balance * interestRate * float64(duration) / (365.0 * 24.0 * float64(time.Hour))
	interestChannel <- interest
}
