package employee

import (
	"context"

	"github.com/GolangNorhtwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

type getEmployeeByIDRequest struct {
	EmployeeID string
}

type getBestEmployeeRequest struct {
}

type addEmployeeRequest struct {
	Address       string
	Business      string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type updateEmployeeRequest struct {
	Id            int64
	Address       string
	Business      string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type deleteEmployeeRequest struct {
	EmployeeId string
}

func MakeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployessEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.getEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployessEndPoint
}

func makeGetEmployeeByIdEndPoint(s Service) endpoint.Endpoint {
	getEmployeeByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		result, err := s.GetEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeeByIdEndPoint
}

func MakeGetBestEmployeeEndPoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {

		result, err := s.GetBestEmployee()
		helper.Catch(err)
		return result, nil
	}
	return getBestEmployeeEndPoint
}

func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {
	addEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		employeeId, err := s.InsertEmployee(&req)
		helper.Catch(err)
		return employeeId, nil
	}
	return addEmployeeEndPoint
}

func makeUpdateEmployeeEndPoint(s Service) endpoint.Endpoint {
	updateEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		employeeId, err := s.UpdateEmployee(&req)
		helper.Catch(err)
		return employeeId, nil
	}
	return updateEmployeeEndPoint
}

func makeDeleteEmployeeEndPoint(s Service) endpoint.Endpoint {
	deleteEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		employeeId, err := s.DeleteEmployee(&req)
		helper.Catch(err)
		return employeeId, nil
	}
	return deleteEmployeeEndPoint
}
