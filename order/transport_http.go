package order

import (
	"context"
	"net/http"
	"strconv"

	"github.com/GolangNorhtwindRestApi/helper"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getOrderByIdHandler := kithttp.NewServer(makeGetOrderByIdEndpoint(s), getOrderByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getOrderByIdHandler)

	return r
}

func getOrderByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	orderId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	helper.Catch(err)

	return getOrderByIdRequest{
		OrderID: orderId}, nil
}
