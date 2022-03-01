package domain

import "github.com/dhruvbehl/bank/errors"

type Customer struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errors.AppError)
	FindById(string) (*Customer, *errors.AppError)
	FindByStatus(string) ([]Customer, *errors.AppError)
}
