package bankingappGo

import (
	"fmt"
	"time"
)

type Account interface {
	Accrue(interestRate float64, duration time.Duration, interestChannel chan<- float64)
}

type CAccount struct {
	accountNumber string
	balance       float64
	interestRate  float64
}

func (c *CAccount) getAccountNumber() string {
	return c.accountNumber
}

func (c *CAccount) getBalance() float64 {
	return c.balance
}

func (c *CAccount) getInterestRate() float64 {
	return c.interestRate
}

func (c *CAccount) setInterestRate(rate float64) {
	c.interestRate = rate
}

func (c *CAccount) deposit(amount float64) {
	c.balance += amount
}

func (c *CAccount) withdraw(amount float64) error {
	if amount > c.balance {
		return fmt.Errorf("Insufficient funds")
	}
	c.balance -= amount
	return nil
}

func (c *CAccount) transfer(amount float64, toAccount *Account) error {
	if amount > c.balance {
		return fmt.Errorf("Insufficient funds")
	}
	c.balance -= amount
	CAccount.balance += amount
	return nil
}

func newCAccount(accountNumber string, balance float64, interestRate float64) *CAccount {
	return &CAccount{
		accountNumber: accountNumber,
		balance:       balance,
		interestRate:  interestRate,
	}
}

func (acc *CAccount) Accrue(interestRate float64, duration time.Duration, interestChannel chan<- float64) {
	interest := acc.balance * interestRate * float64(duration) / (365.0 * 24.0 * float64(time.Hour))
	interestChannel <- interest
}

type savingsAccount struct {
	accountNumber int
	balance       float64
	interestRate  float64
}

func newsavingsAccount(accountNumber int, balance float64, interestRate float64) *savingsAccount {
	return &savingsAccount{
		accountNumber: accountNumber,
		balance:       balance,
		interestRate:  interestRate,
	}
}

func (sa *savingsAccount) Deposit(amount float64) {
	if amount > 0 {
		sa.balance += amount
		fmt.Println("Deposit successful")
	} else {
		fmt.Println("Invalid amount")
	}
}

func (sa *savingsAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= sa.balance {
		sa.balance -= amount
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Invalid amount or insufficient funds")
	}
}

func (sa *savingsAccount) AddInterest() {
	periods := 12 // assuming monthly compounding
	rate := sa.interestRate / float64(periods)
	interest := sa.balance * rate
	sa.balance += interest
}

func (sa *savingsAccount) Accrue(interestRate float64, duration time.Duration, interestChannel chan<- float64) {
	interest := sa.balance * interestRate * float64(duration) / (365.0 * 24.0 * float64(time.Hour))
	interestChannel <- interest
}
