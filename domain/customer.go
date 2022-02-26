package domain

type Customer struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Statusa     string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
