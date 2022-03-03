package employee

import (
	"context"

	"github.com/GoGooliveryProviderAPI/helper"
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

// @Summary Lista de Empleados
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body employee.getEmployeesRequest true "User Data"
// @Success 200 {object} employee.EmployeeList "ok"
// @Router /employees/paginated [post]
func MakeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployessEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.getEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployessEndPoint
}

// @Summary Empleado By Id
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "ok"
// @Router /employees/{id} [get]
func makeGetEmployeeByIdEndPoint(s Service) endpoint.Endpoint {
	getEmployeeByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIDRequest)
		result, err := s.GetEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployeeByIdEndPoint
}

// @Summary Mejor Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Success 200 {object} employee.BestEmployee "ok"
// @Router /employees/best [get]
func MakeGetBestEmployeeEndPoint(s Service) endpoint.Endpoint {
	getBestEmployeeEndPoint := func(_ context.Context, _ interface{}) (interface{}, error) {
		result, err := s.GetBestEmployee()
		helper.Catch(err)
		return result, nil
	}
	return getBestEmployeeEndPoint
}

// @Summary Insertar Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body employee.addEmployeeRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [post]
func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {
	addEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		employeeId, err := s.InsertEmployee(&req)
		helper.Catch(err)
		return employeeId, nil
	}
	return addEmployeeEndPoint
}

// @Summary Actualizar Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Param request body employee.updateEmployeeRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /employees/ [put]
func makeUpdateEmployeeEndPoint(s Service) endpoint.Endpoint {
	updateEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		employeeId, err := s.UpdateEmployee(&req)
		helper.Catch(err)
		return employeeId, nil
	}
	return updateEmployeeEndPoint
}

// @Summary Eliminar Empleado
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {integer} int "ok"
// @Router /employees/{id} [delete]
func makeDeleteEmployeeEndPoint(s Service) endpoint.Endpoint {
	deleteEmployeeEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		employeeId, err := s.DeleteEmployee(&req)
		helper.Catch(err)
		return employeeId, nil
	}
	return deleteEmployeeEndPoint
}
