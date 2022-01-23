package employee

import "github.com/GolangNorhtwindRestApi/helper"

type Service interface {
	getEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	//GetTotalEmployees() (int64, error)
	GetEmployeeById(params *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(params *deleteEmployeeRequest) (int64, error)
}

type service struct {
	repo Respository
}

func NewService(repo Respository) Service {
	return &service{repo: repo}
}

func (s *service) getEmployees(params *getEmployeesRequest) (*EmployeeList, error) {
	employees, err := s.repo.GetEmployees(params)
	helper.Catch(err)
	totalEmployees, err := s.repo.GetTotalEmployees()
	helper.Catch(err)
	return &EmployeeList{Data: employees, TotalRecords: totalEmployees}, nil
}

func (s *service) GetEmployeeById(params *getEmployeeByIDRequest) (*Employee, error) {
	return s.repo.GetEmployeeById(params)
}

func (s *service) GetBestEmployee() (*BestEmployee, error) {
	return s.repo.GetBestEmployee()
}

func (s *service) InsertEmployee(params *addEmployeeRequest) (int64, error) {
	return s.repo.InsertEmployee(params)
}

func (s *service) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	return s.repo.UpdateEmployee(params)
}

func (s *service) DeleteEmployee(params *deleteEmployeeRequest) (int64, error) {
	return s.repo.DeleteEmployee(params)
}
