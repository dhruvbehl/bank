package service

import "github.com/dhruvbehl/bank/domain"

type CustomerService interface{
	getAllCustomer() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return d.repository.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}