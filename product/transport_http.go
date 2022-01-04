package product

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	kitHttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductBYIDHandler := kitHttp.NewServer(MakeGetProductByIdEndPoint,
		GetProductoByIdRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodGet,
		"/{id}", getProductBYIDHandler)
	return r
}

func GetProductoByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return GetProductByIDRequest{ProductID: productId}, nil
}
