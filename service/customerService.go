package service

import (
	"strings"

	"github.com/dhruvbehl/bank/domain"
	"github.com/dhruvbehl/bank/dto"
	"github.com/dhruvbehl/bank/errors"
)

type CustomerService interface{
	GetAllCustomer() ([]dto.CustomerResponse, error)
	GetCustomerById(string) (*dto.CustomerResponse, error)
	GetCustomerByStatus(string) ([]dto.CustomerResponse, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errors.AppError) {
	customers, err := d.repository.FindAll()
	if err != nil {
		return nil, err
	}
	customerResponse := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		customerResponse = append(customerResponse, dto.CustomerResponse(*c.GetCustomerDto()))
	}
	return customerResponse, nil
}

func (d DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errors.AppError) {
	c, err := d.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return c.GetCustomerDto(), nil
}

func (d DefaultCustomerService) GetCustomerByStatus(status string) ([]dto.CustomerResponse, *errors.AppError) {
	if strings.EqualFold(status, "active") {
		status = "1"
	} else {
		status = "0"
	}
	customers, err := d.repository.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	customerResponse := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		customerResponse = append(customerResponse, dto.CustomerResponse(*c.GetCustomerDto()))
	}
	return customerResponse, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}