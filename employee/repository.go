package employee

import (
	"database/sql"

	"github.com/GolangNorhtwindRestApi/helper"
)

type Respository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
	GetEmployeeById(params *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(params *deleteEmployeeRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Respository {
	return &repository{
		db: db,
	}
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT
				id,
				first_name,
				last_name,
				company,
				email_address,
				job_title,
				business_phone,
				home_phone,
				COALESCE(mobile_phone, ''),
				fax_number,
				address
				FROM employees
				LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)
	var employees []*Employee
	for results.Next() {
		employee := &Employee{}
		err = results.Scan(
			&employee.ID,
			&employee.FirstName,
			&employee.LastName,
			&employee.Company,
			&employee.EmailAddress,
			&employee.JobTitle,
			&employee.BusinessPhone,
			&employee.HomePhone,
			&employee.MobilePhone,
			&employee.FaxNumber,
			&employee.Address)
		helper.Catch(err)
		employees = append(employees, employee)
	}
	return employees, nil
}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = "SELECT COUNT(*) FROM employees"
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}

func (repo *repository) GetEmployeeById(params *getEmployeeByIDRequest) (*Employee, error) {
	const sql = `SELECT id,
				first_name,
				last_name,
				company,
				email_address,
				job_title,
				business_phone,
				home_phone,
				COALESCE(mobile_phone, ''),
				fax_number,
				address
				FROM employees
				WHERE id=?`
	row := repo.db.QueryRow(sql, params.EmployeeID)
	employee := &Employee{}
	err := row.Scan(&employee.ID,
		&employee.FirstName,
		&employee.LastName,
		&employee.Company,
		&employee.EmailAddress,
		&employee.JobTitle,
		&employee.BusinessPhone,
		&employee.HomePhone,
		&employee.MobilePhone,
		&employee.FaxNumber,
		&employee.Address)
	helper.Catch(err)

	return employee, nil
}

func (repo *repository) GetBestEmployee() (*BestEmployee, error) {
	const sql = `SELECT e.id,
				count(e.id) as totalVentas,
				e.first_name,
				e.last_name
				FROM orders o
				INNER JOIN employees e ON o.employee_id= e.id 
				GROUP BY o.employee_id
				ORDER BY totalVentas desc
				limit 1`

	row := repo.db.QueryRow(sql)
	employee := &BestEmployee{}
	err := row.Scan(&employee.Id, &employee.TotalVentas, &employee.FirstName, &employee.LastName)
	helper.Catch(err)
	return employee, nil

}

func (repo *repository) InsertEmployee(params *addEmployeeRequest) (int64, error) {
	const sql = `INSERT INTO employees
				(first_name,
				 last_name,
				 company,
				 address,
				 business_phone,
				 email_address,
				 fax_number,
				 home_phone,
				 job_title,
				 mobile_phone)
				 VALUES (?,?,?,?,?,?,?,?,?,?)`

	result, err := repo.db.Exec(sql,
		params.FirstName,
		params.LastName,
		params.Company,
		params.Address,
		params.BusinessPhone,
		params.EmailAddress,
		params.FaxNumber,
		params.HomePhone,
		params.JobTitle,
		params.MobilePhone)

	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil

}

func (repo *repository) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	const sql = `UPDATE employees
				SET first_name =?,
				last_name =?,
				company =?,
				address =?,
				business_phone=?,
				email_address=?,
				fax_number=?,
				home_phone=?,
				job_title=?,
				mobile_phone=?
				WHERE id=?`
	result, err := repo.db.Exec(sql,
		params.FirstName,
		params.LastName,
		params.Company,
		params.Address,
		params.BusinessPhone,
		params.EmailAddress,
		params.FaxNumber,
		params.HomePhone,
		params.JobTitle,
		params.MobilePhone)

	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) DeleteEmployee(params *deleteEmployeeRequest) (int64, error) {
	const sql = `DELETE FROM employees WHERE id = ?`
	result, err := repo.db.Exec(sql, params.EmployeeId)
	helper.Catch(err)
	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}
