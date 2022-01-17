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

func MakeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	getEmployessEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.getEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEmployessEndPoint
}
