package bankingappGo

import "fmt"

type Customer struct {
	firstName, lastName, address string
}

func NewCustomer(firstName, lastName, address string) *Customer {
	return &Customer{
		firstName: firstName,
		lastName:  lastName,
		address:   address,
	}
}

func (c *Customer) GetFirstName() string {
	return c.firstName
}

func (c *Customer) SetFirstName(firstName string) {
	c.firstName = firstName
}

func (c *Customer) GetLastName() string {
	return c.lastName
}

func (c *Customer) SetLastName(lastName string) {
	c.lastName = lastName
}

func (c *Customer) GetAddress() string {
	return c.address
}

func (c *Customer) SetAddress(address string) {
	c.address = address
}

func (c *Customer) String() string {
	return fmt.Sprintf("%s %s, %s", c.firstName, c.lastName, c.address)
}
