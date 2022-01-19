package employee

import "github.com/GolangNorhtwindRestApi/helper"

type Service interface {
	getEmployees(params *getEmployeesRequest) (*EmployeeList, error)
	//GetTotalEmployees() (int64, error)
	GetEmployeeById(params *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
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
