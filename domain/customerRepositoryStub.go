package domain

import "fmt"

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func (s CustomerRepositoryStub) FinById(id string) (Customer, error) {
	for _, c := range s.customer {
		if c.Id == id {
			return c, nil
		}
	}
	return Customer{}, fmt.Errorf("no customer found with id [%v]", id)
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	newCustomer := []Customer{
		{"1", "Dhruv Behl", "Lucknow", "226006", "1992-12-20", "1"},
		{"2", "Aditi Behl", "Lucknow", "226006", "1998-08-01", "0"},
	}

	return CustomerRepositoryStub{newCustomer}
}