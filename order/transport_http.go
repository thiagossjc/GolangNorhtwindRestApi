package order

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIdHandler := kithttp.NewServer(makeGetOrderByIdEndpoint(s), getOrderByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIdHandler)

	getOrdersHandler := kithttp.NewServer(makeGetOrdersEndpoint(s), getOrdersRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getOrdersHandler)

	addOrderHandler := kithttp.NewServer(makeAddOrderEndpoint(s), addOrderRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addOrderHandler)

	updateOrderHandler := kithttp.NewServer(makeUpdateOrderEndpoint(s), updateOrderRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateOrderHandler)

	deleteOrderDetailHandler := kithttp.NewServer(makeDeleteOrderDetailEndpoint(s), deleteOrderDetailRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{orderId}/detail/{orderDetailId}", deleteOrderDetailHandler)

	deleteOrderHandler := kithttp.NewServer(makeDeleteOrderEndpoint(s), deleteOrderRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteOrderHandler)

	return r
}

func getOrderByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)

	return getOrderByIdRequest{
		OrderID: orderId}, nil
}

func getOrdersRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getOrdersRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func addOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func updateOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := addOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func deleteOrderDetailRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderDetailRequest{
		OrderDetailId: chi.URLParam(r, "orderDetailId"),
	}, nil
}

func deleteOrderRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	return deleteOrderRequest{
		OrderId: chi.URLParam(r, "id"),
	}, nil
}
