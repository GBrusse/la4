package bankingappGo

import (
	"fmt"
	"time"
)

type CheckingAccount struct {
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

func (c *CAccount) transfer(amount float64, toAccount *CheckingAccount) error {
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

func (acc *CAccount) Accrue(interestRate float64, duration time.Duration, interestChannel chan float64) {
	defer close(interestChannel)

	interest := acc.balance * interestRate * float64(duration) / 365.0
	interestChannel <- interest
}
