package domain

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/dhruvbehl/bank/errors"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errors.AppError) {
	sqlQuery := "select * from customers"
	rows, err := d.client.Query(sqlQuery)
	if err != nil {
		log.Default().Printf("er ror while querying db: %v\n", err.Error())
		return nil, errors.NewInternalServerError("unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Default().Printf("error while scanning results: %v\n", err.Error())
			return nil, errors.NewInternalServerError("error while scanning results")
		}
		customers = append(customers, c)
	}
	log.Default().Println("returning results for FindAll query")
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errors.AppError) {
	sqlQuery := fmt.Sprintf("select * from customers where customer_id=%v", id)
	row := d.client.QueryRow(sqlQuery)

	c := Customer{}
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Default().Printf("error while scanning results: %v\n", err.Error())
			return nil, errors.NewNotFoundError("customer not found")
		}
		log.Default().Printf("error while scanning results: %v\n", err.Error())
		return nil, errors.NewInternalServerError("unexpected database error")
	}
	log.Default().Println("returning results for FindById query")
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:qs@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}