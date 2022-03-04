package user

import (
	"context"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-kit/kit/endpoint"
)

type getUserByIDRequest struct {
	UserID    string
	CountryId string
}

type addUserRequest struct {
	FirstName string
	LastName  string
	DocState  string
	Login     string
	Password  string
	Email     string
	CelPhone  string
	Country   string
	DtReg     string
	IdUser    int64
	Cancel    bool
	IdCancel  int64
}

type updateUserRequest struct {
	Id        int64
	FirstName string
	LastName  string
	DocState  string
	Login     string
	Password  string
	Email     string
	CelPhone  string
	Country   string
	DtReg     string
	IdUser    int64
	Cancel    bool
	IdCancel  int64
}

type cancelUserRequest struct {
	UserId string
}

// @Summary User By User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User Id"
// @Success 200 {object} user.User "ok"
// @Router /users/{id} [get]
func makeGetUserByIdEndPoint(s Service) endpoint.Endpoint {
	getUserByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByIDRequest)
		result, err := s.GetUserById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getUserByIdEndPoint
}

// @Summary Insertar User
// @Tags Users
// @Accept json
// @Produce json
// @Param request body user.addUserRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /users/ [post]
func makeAddUserEndPoint(s Service) endpoint.Endpoint {
	addUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addUserRequest)
		userId, err := s.InsertUser(&req)
		helper.Catch(err)
		return userId, nil
	}
	return addUserEndPoint
}

// @Summary Actualizar Users
// @Tags Users
// @Accept json
// @Produce json
// @Param request body user.updateUserRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /users/ [put]
func makeUpdateUserEndPoint(s Service) endpoint.Endpoint {
	updateUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		userId, err := s.UpdateUser(&req)
		helper.Catch(err)
		return userId, nil
	}
	return updateUserEndPoint
}

// @Summary Cancelar User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {integer} int "ok"
// @Router /users/{id} [delete]
func makeCancelUserEndPoint(s Service) endpoint.Endpoint {
	cancelUserEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelUserRequest)
		userId, err := s.CancelUser(&req)
		helper.Catch(err)
		return userId, nil
	}
	return cancelUserEndPoint
}
