package request

import (
	"context"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-kit/kit/endpoint"
)

type getRequestBYIDRequest struct {
	RequestID int
	Country   string
}

type getRequestsRequest struct {
	Limit   int
	Offset  int
	Country string
}

type addRequestOrderRequest struct {
	IdCustomer         int64
	IdSupplier         int64
	TotalQuantity      int64
	TotalPrice         float32
	TotPriceWidhoutTax float32
	Tax                float32
	Other_tax          float32
	Country            string
	DocState           string
	Email              string
	Cellphone          string
	Description        string
	IdUser             int64
	DeliverySupplier   bool
	IdStatus           int32
	idEvent            int32
}

type updateRequestOrderRequest struct {
	Id                 int64
	IdCustomer         int64
	IdSupplier         int64
	TotalQuantity      int64
	TotalPrice         float32
	TotPriceWidhoutTax float32
	Tax                float32
	Other_tax          float32
	Country            string
	DocState           string
	Email              string
	Cellphone          string
	Description        string
	IdCancel           int64
	DeliverySupplier   bool
	IdStatus           int32
	idEvent            int32
}

type cancelRequestOrderRequest struct {
	Id       int64
	IdCancel int64
	Cancel   bool
	IdStatus int32
	idEvent  int32
}

type cancelRequestDetailRequest struct {
	RequestDetailId string
}

// @Summary Requests By Id
// @Tags Requests
// @Accept json
// @Produce json
// @Param id path int true "Requests Id"
// @Success 200 {object} request.RequestOrder "ok"
// @Router /requests{id} [get]
func makeGetRequestByIdEndpoint(s Service) endpoint.Endpoint {
	getRequestByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequestBYIDRequest)
		result, err := s.GetRequestOrderById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getRequestByIdEndPoint
}

// @Summary Lista de Requests
// @Tags Requests
// @Accept json
// @Produce json
// @Param request body request.getRequestsRequest true "User Data"
// @Success 200 {object} request. "ok"
// @Router /requests/paginated [post]
func makeGetRequestsEndpoint(s Service) endpoint.Endpoint {
	getRequestsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequestsRequest)
		result, err := s.GetRequestsOrders(&req)
		helper.Catch(err)
		return result, nil
	}
	return getRequestsEndPoint
}

// @Summary Insertar Request
// @Tags Requests
// @Accept json
// @Produce json
// @Param request body request.addRequestOrderRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /requests/ [post]
func makeAddOrderEndpoint(s Service) endpoint.Endpoint {
	addRequestEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequestOrderRequest)
		result, err := s.InsertRequestOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return addRequestEndPoint
}

// @Summary Actualizar Request
// @Tags Requests
// @Accept json
// @Produce json
// @Param request body request.addRequestRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /requests/ [put]
func makeUpdateRequestEndpoint(s Service) endpoint.Endpoint {
	updateRequestEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequestOrderRequest)
		result, err := s.UpdateRequestOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return updateRequestEndPoint
}

// @Summary Cancelar Detal de la Orden
// @Tags Requests
// @Accept json
// @Produce json
// @Param orderDetailId path int true "Order Detail ID"
// @Success 200 {integer} int "ok"
// @Router /orders/{orderId}/detail/{orderDetailId} [delete]
func makeCancelOrderDetailEndpoint(s Service) endpoint.Endpoint {
	cancelRequestDetailEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelRequestOrderRequest)
		result, err := s.CancelRequestDetail(&req)
		helper.Catch(err)
		return result, nil
	}
	return cancelRequestDetailEndPoint
}

// @Summary Cancelar Requisición
// @Tags Requests
// @Accept json
// @Produce json
// @Param id path int true "Request ID"
// @Success 200 {integer} int "ok"
// @Router /requests/{id} [delete]
func makeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	deleteOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelRequestOrderRequest)
		result, err := s.CancelRequestOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteOrderEndPoint
}
