package domain

import (
	"strconv"

	"github.com/dhruvbehl/bank/dto"
	"github.com/dhruvbehl/bank/errors"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) GetCustomerDto() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.getStatus(),
	}
}

func (c Customer) getStatus() string {
	statusString := []string{"inactive", "active"}
	statusInt, _ := strconv.Atoi(c.Status)
	return statusString[statusInt]
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errors.AppError)
	FindById(string) (*Customer, *errors.AppError)
	FindByStatus(string) ([]Customer, *errors.AppError)
}
