package order

import (
	"context"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-kit/kit/endpoint"
)

type getOrderByIdRequest struct {
	OrderID int64
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

type addOrderRequest struct {
	OrderDate    string
	CustomerId   int64
	OrderDetails []*addOrderDetailRequest
	//Vamos utilizalas para el update acrescentando la variable ID
	Id int64
}

type addOrderDetailRequest struct {
	OrderId   int64
	ProductId int64
	Quantity  int64
	UnitPrice float64
	//Vamos utilizalas para el update acrescentando la variable ID
	Id int64
}

type deleteOrderDetailRequest struct {
	OrderDetailId string
}

type deleteOrderRequest struct {
	OrderId string
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
		req := request.(getOrderByIdRequest)
		result, err := s.GetOrderById(&req)
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
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
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
		req := request.(addOrderRequest)
		result, err := s.InsertOrder(&req)
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
		req := request.(addOrderRequest)
		result, err := s.UpdateOrder(&req)
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
		req := request.(deleteOrderDetailRequest)
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
		req := request.(deleteOrderRequest)
		result, err := s.DeleteOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteOrderEndPoint
}
