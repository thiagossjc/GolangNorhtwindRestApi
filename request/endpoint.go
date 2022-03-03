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

// @Summary Order By Id
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order Id"
// @Success 200 {object} order.OrderItem "ok"
// @Router /orders{id} [get]
func makeGetOrderByIdEndpoint(s Service) endpoint.Endpoint {
	getOrderByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequestBYIDRequest)
		result, err := s.GetRequestOrderById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrderByIdEndPoint
}

// @Summary Lista de Ordenes
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body order.getOrdersRequest true "User Data"
// @Success 200 {object} order.OrderList "ok"
// @Router /orders/paginated [post]
func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	getOrdersEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequestsRequest)
		result, err := s.GetRequestsOrders(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrdersEndPoint
}

// @Summary Insertar Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body order.addOrderRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /orders/ [post]
func makeAddOrderEndpoint(s Service) endpoint.Endpoint {
	addOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequestOrderRequest)
		result, err := s.InsertRequestOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return addOrderEndPoint
}

// @Summary Actualizar Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body order.addOrderRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /orders/ [put]
func makeUpdateOrderEndpoint(s Service) endpoint.Endpoint {
	updateOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequestOrderRequest)
		result, err := s.UpdateRequestOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return updateOrderEndPoint
}

// @Summary Eliminar Detal de la Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderDetailId path int true "Order Detail ID"
// @Success 200 {integer} int "ok"
// @Router /orders/{orderId}/detail/{orderDetailId} [delete]
func makeDeleteOrderDetailEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDetailEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelRequestOrderRequest)
		result, err := s.DeleteOrderDetail(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteOrderDetailEndPoint
}

// @Summary Eliminar Orden
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {integer} int "ok"
// @Router /orders/{id} [delete]
func makeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	deleteOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelRequestOrderRequest)
		result, err := s.CancelRequestOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteOrderEndPoint
}
