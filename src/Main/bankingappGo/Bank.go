package bakingappGo

import (
	"fmt"
	"time"
)

type bank struct {
	Customers         []*Customer
	NumberOfCustomers int
}

func NewBank() *Bank {
	return &Bank{}
}

func (b *Bank) AddCustomer(firstName, lastName string) {
	b.NumberOfCustomers++
	customer := NewCustomer(firstName, lastName)
	b.Customers = append(b.Customers, customer)
}

func (b *bank) GetNumberOfCustomers() int {
	return b.NumberOfCustomers
}

func (b *bank) GetCustomer(index int) *Customer {
	if index < 0 || index >= b.NumberOfCustomers {
		return nil
	}
	return b.Customers[index]
}

func (b *bank) ShowCustomers() {
	for i := 0; i < b.NumberOfCustomers; i++ {
		customer := b.GetCustomer(i)
		fmt.Printf("Customer [%d] = %s %s\n", i+1, customer.GetFirstName(), customer.GetLastName())
	}
}

func (b *bank) Accrue(interestRate float64, duration time.Duration) {
	interestChannel := make(chan float64)
	var totalInterest float64

	for _, customer := range b.Customers {
		for _, account := range customer.Accounts {
			go account.Accrue(interestRate, duration, interestChannel)
		}
	}

	for range b.Customers {
		for interest := range interestChannel {
			totalInterest += interest
		}
	}

	fmt.Printf("Total interest accrued: $%.2f\n", totalInterest)
}
