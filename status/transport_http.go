package status

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

	getStatusHandler := kithttp.NewServer(MakeGetStatusEndPoint(s), getStatusRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getStatusHandler)

	getStatusByIdHandler := kithttp.NewServer(makeGetStatusByIdEndPoint(s), getStatusByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getStatusByIdHandler)

	AddStatusHandler := kithttp.NewServer(makeAddStatuEndPoint(s), getaddStatuRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", AddStatusHandler)

	UpdateStatuHandler := kithttp.NewServer(makeUpdateStatuEndPoint(s), getUpdateStatuRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", UpdateStatuHandler)

	CancelaStatuHandler := kithttp.NewServer(makeCancelaStatusEndPoint(s), getCancelStatuRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", CancelaStatuHandler)

	return r
}

func getStatusRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getStatusRequest{
		Limit:     0,
		Offset:    0,
		CountryId: 0,
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getStatusByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	return getStatusByIDRequest{
		StatusID: chi.URLParam(r, "id")}, nil
}

func getaddStatuRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addStatusRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getUpdateStatuRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateStatusRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getCancelStatuRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return cancelStatusRequest{
		StatusId: chi.URLParam(r, "id"),
	}, nil
}
