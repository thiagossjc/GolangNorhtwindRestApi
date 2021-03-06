package customer

import (
	"database/sql"

	"github.com/GolangNorhtwindRestApi/helper"
)

type Repository interface {
	GetCustomers(param *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomers() (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetCustomers(param *getCustomersRequest) ([]*Customer, error) {
	const sql = `SELECT id,
				 first_name,
				 last_name,
				 address,
				 business_phone,
				 city,
				 company
				 FROM customers
				 LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, param.Limit, param.Offset)
	helper.Catch(err)

	var customers []*Customer
	for results.Next() {
		customer := &Customer{}
		//scan captura os valores do resultado
		err = results.Scan(
			&customer.ID,
			&customer.FirstName,
			&customer.LastName,
			&customer.Address,
			&customer.BusinessPhone,
			&customer.City,
			&customer.Company)
		helper.Catch(err)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repo *repository) GetTotalCustomers() (int64, error) {
	const sql = "SELECT COUNT(*) AS total FROM customers"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}
