package app

type CustomerDetails struct {
	Id		string	`json:"customer_id"`
	Name	string	`json:"full_name"`
	City	string	`json:"city"`
	Phone	string	`json:"phone"`
}