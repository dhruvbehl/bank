package service

import (
	"strings"

	"github.com/dhruvbehl/bank/domain"
	"github.com/dhruvbehl/bank/errors"
)

type CustomerService interface{
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, error)
	GetCustomerByStatus(string) ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errors.AppError) {
	return d.repository.FindAll()
}

func (d DefaultCustomerService) GetCustomerById(id string) (*domain.Customer, *errors.AppError) {
	return d.repository.FindById(id)
}

func (d DefaultCustomerService) GetCustomerByStatus(status string) ([]domain.Customer, *errors.AppError) {
	if strings.EqualFold(status, "active") {
		status = "1"
	} else {
		status = "0"
	}
	return d.repository.FindByStatus(status)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}