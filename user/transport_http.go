package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getUserByIdHandler := kithttp.NewServer(makeGetUserByIdEndPoint(s), getUserByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getUserByIdHandler)

	AddUserHandler := kithttp.NewServer(makeAddUserEndPoint(s), getaddUserRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", AddUserHandler)

	UpdateUserHandler := kithttp.NewServer(makeUpdateUserEndPoint(s), getUpdateUserRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", UpdateUserHandler)

	CancelUserHandler := kithttp.NewServer(makeCancelUserEndPoint(s), getCancelUserRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", CancelUserHandler)

	return r
}

func getUserByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	return getUserByIDRequest{
		UserID: chi.URLParam(r, "id")}, nil
}

func getaddUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getUpdateUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getCancelUserRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return cancelUserRequest{
		UserId: chi.URLParam(r, "id"),
	}, nil
}
