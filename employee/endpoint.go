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
