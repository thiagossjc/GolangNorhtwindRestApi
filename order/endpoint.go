package order

import (
	"context"

	"github.com/GolangNorhtwindRestApi/helper"
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

func makeGetOrderByIdEndpoint(s Service) endpoint.Endpoint {
	getOrderByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderByIdRequest)
		result, err := s.GetOrderById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrderByIdEndPoint
}

func makeGetOrdersEndpoint(s Service) endpoint.Endpoint {
	getOrdersEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
		helper.Catch(err)
		return result, nil
	}
	return getOrdersEndPoint
}

func makeAddOrderEndpoint(s Service) endpoint.Endpoint {
	addOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.InsertOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return addOrderEndPoint
}

func makeUpdateOrderEndpoint(s Service) endpoint.Endpoint {
	updateOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.UpdateOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return updateOrderEndPoint
}

func makeDeleteOrderDetailEndpoint(s Service) endpoint.Endpoint {
	deleteOrderDetailEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		result, err := s.DeleteOrderDetail(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteOrderDetailEndPoint
}

func makeDeleteOrderEndpoint(s Service) endpoint.Endpoint {
	deleteOrderEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		result, err := s.DeleteOrder(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteOrderEndPoint
}
