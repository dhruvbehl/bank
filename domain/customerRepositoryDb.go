package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dhruvbehl/bank/errors"
	"github.com/dhruvbehl/bank/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errors.AppError) {
	sqlQuery := "select * from customers"
	rows, err := d.client.Query(sqlQuery)
	if err != nil {
		logger.Error(fmt.Sprintf("[error while querying db] %v", err.Error()))
		return nil, errors.NewInternalServerError("unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			logger.Error(fmt.Sprintf("[error while scanning results] %v",err.Error()))
			return nil, errors.NewInternalServerError("error while scanning results")
		}
		customers = append(customers, c)
	}
	logger.Info("returning results for FindAll query")
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errors.AppError) {
	sqlQuery := fmt.Sprintf("select * from customers where customer_id=%v", id)
	row := d.client.QueryRow(sqlQuery)

	c := Customer{}
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error(fmt.Sprintf("[error while scanning results] %v",err.Error()))
			return nil, errors.NewNotFoundError("customer not found")
		}
		logger.Error(fmt.Sprintf("[error while scanning results] %v",err.Error()))
		return nil, errors.NewInternalServerError("unexpected database error")
	}
	logger.Info("returning results for FindById query")
	return &c, nil
}

func (d CustomerRepositoryDb) FindByStatus(status string) ([]Customer, *errors.AppError) {
	sqlQuery := fmt.Sprintf("SELECT * FROM customers where status=%v", status)
	rows, err := d.client.Query(sqlQuery)
	if err != nil {
		logger.Error(fmt.Sprintf("[error while querying db] %v", err.Error()))
		return nil, errors.NewInternalServerError("unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			logger.Error(fmt.Sprintf("[error while scanning results] %v",err.Error()))
			return nil, errors.NewInternalServerError("error while scanning results")
		}
		customers = append(customers, c)
	}
	logger.Info("returning results for FindByStatus query")
	return customers, nil
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