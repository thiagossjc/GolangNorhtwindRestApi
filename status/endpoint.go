package status

import (
	"context"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-kit/kit/endpoint"
)

type getStatusRequest struct {
	Limit     int
	Offset    int
	CountryId int32
}

type getStatusByIDRequest struct {
	StatusID  string
	CountryId string
}

type addStatusRequest struct {
	Description         string
	UsereReg            int64
	Success             bool
	Traffic             bool
	InCustomer          bool
	InSupplier          bool
	OnPicking           bool
	OnPacking           bool
	OnExpedition        bool
	InAttend            bool
	GoingTo             string
	InLogisticOperation bool
	IdCountry           int32
}

type updateStatusRequest struct {
	Id                  int32
	Description         string
	UsereReg            int64
	Success             bool
	Traffic             bool
	InCustomer          bool
	InSupplier          bool
	OnPicking           bool
	OnPacking           bool
	OnExpedition        bool
	InAttend            bool
	GoingTo             string
	InLogisticOperation bool
	IdCountry           int32
}

type cancelStatusRequest struct {
	StatusId string
}

// @Summary Lista de Status
// @Tags Status
// @Accept json
// @Produce json
// @Param request body status.getStatusRequest true "User Data"
// @Success 200 {object} status.StatusList "ok"
// @Router /status/paginated [post]
func MakeGetStatusEndPoint(s Service) endpoint.Endpoint {
	getStatusEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getStatusRequest)
		result, err := s.GetStatus(&req)
		helper.Catch(err)
		return result, nil
	}
	return getStatusEndPoint
}

// @Summary Status By Id
// @Tags Status
// @Accept json
// @Produce json
// @Param id path int true "Status Id"
// @Success 200 {object} status.Status "ok"
// @Router /status/{id} [get]
func makeGetStatusByIdEndPoint(s Service) endpoint.Endpoint {
	getStatusByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getStatusByIDRequest)
		result, err := s.GetStatusById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getStatusByIdEndPoint
}

// @Summary Insertar Status
// @Tags Status
// @Accept json
// @Produce json
// @Param request body employee.addStatusRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /status/ [post]
func makeAddStatuEndPoint(s Service) endpoint.Endpoint {
	addStatuEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addStatusRequest)
		statusId, err := s.InsertStatu(&req)
		helper.Catch(err)
		return statusId, nil
	}
	return addStatuEndPoint
}

// @Summary Actualizar Status
// @Tags Status
// @Accept json
// @Produce json
// @Param request body status.updateStatusRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /status/ [put]
func makeUpdateStatuEndPoint(s Service) endpoint.Endpoint {
	updateStatuEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateStatusRequest)
		statuId, err := s.UpdateStatu(&req)
		helper.Catch(err)
		return statuId, nil
	}
	return updateStatuEndPoint
}

// @Summary Cancelar Status
// @Tags Status
// @Accept json
// @Produce json
// @Param id path int true "Status ID"
// @Success 200 {integer} int "ok"
// @Router /status/{id} [delete]
func makeCancelaStatusEndPoint(s Service) endpoint.Endpoint {
	cancelaStatuEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelStatusRequest)
		statuId, err := s.CancelaStatu(&req)
		helper.Catch(err)
		return statuId, nil
	}
	return cancelaStatuEndPoint
}
