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

	return r
}

func getEmployeesRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEmployeesRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEmployeesByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	//request := getEmployeeByIDRequest{}
	//err := json.NewDecoder(r.Body).Decode(&request)
	//helper.Catch(err)
	return getEmployeeByIDRequest{
		EmployeeID: chi.URLParam(r, "id")}, nil
}
