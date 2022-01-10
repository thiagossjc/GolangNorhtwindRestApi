package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	kitHttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductBYIDHandler := kitHttp.NewServer(MakeGetProductByIdEndPoint(s),
		GetProductoByIdRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductBYIDHandler)

	getProductsHandler := kitHttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)
	return r
}

func GetProductoByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductBYIDRequest{ProdutoID: productId}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}
	return request, nil
}
