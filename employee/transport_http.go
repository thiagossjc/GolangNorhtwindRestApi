package employee

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GolangNorhtwindRestApi/helper"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getEmployesHandler := kithttp.NewServer(MakeGetEmployeesEndPoint(s), getEmployeesRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEmployesHandler)

	getEmployesByIdHandler := kithttp.NewServer(makeGetEmployeeByIdEndPoint(s), getEmployeesByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEmployesByIdHandler)

	getBestEmployeeHandler := kithttp.NewServer(MakeGetBestEmployeeEndPoint(s), getBestEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/best", getBestEmployeeHandler)

	AddEmployeeHandler := kithttp.NewServer(makeAddEmployeeEndPoint(s), getaddEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", AddEmployeeHandler)

	UpdateEmployeeHandler := kithttp.NewServer(makeUpdateEmployeeEndPoint(s), getUpdateEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", UpdateEmployeeHandler)

	DeleteEmployeeHandler := kithttp.NewServer(makeDeleteEmployeeEndPoint(s), getDeleteEmployeeRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", DeleteEmployeeHandler)

	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{
		Limit:  0,
		Offset: 0,
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeesByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	return getEmployeeByIDRequest{
		EmployeeID: chi.URLParam(r, "id")}, nil
}

func getBestEmployeeRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return getBestEmployeeRequest{}, nil
}

func getaddEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getUpdateEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateEmployeeRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getDeleteEmployeeRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deleteEmployeeRequest{
		EmployeeId: chi.URLParam(r, "id"),
	}, nil
}
