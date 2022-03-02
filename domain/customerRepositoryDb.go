package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dhruvbehl/bank/errors"
	"github.com/dhruvbehl/bank/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errors.AppError) {
	customers := make([]Customer, 0)
	sqlQuery := "select * from customers"
	err := d.client.Select(&customers, sqlQuery)
	if err != nil {
		logger.Error(fmt.Sprintf("[error while scanning results] %v",err.Error()))
		return nil, errors.NewInternalServerError("error while scanning results")
	}
	logger.Info("returning results for FindAll query")
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errors.AppError) {
	c := Customer{}
	sqlQuery := fmt.Sprintf("select * from customers where customer_id=%v", id)
	err := d.client.Get(&c, sqlQuery)

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
	customers := make([]Customer, 0)
	sqlQuery := fmt.Sprintf("SELECT * FROM customers where status=%v", status)
	err := d.client.Select(&customers, sqlQuery)
	if err != nil {
		logger.Error(fmt.Sprintf("[error while scanning results] %v",err.Error()))
		return nil, errors.NewInternalServerError("error while scanning results")
	}
	logger.Info("returning results for FindByStatus query")
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:qs@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}