package product

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()
	getProductBYIDHandler := kitHttp.NewServer(MakeGetProductByIdEndPoint(s Service),GetProductoByIdRequestDecoder,kitHttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductBYIDHandler)
	return r
}

func GetProductoByIdRequestDecoder(context context.Context, r *http.Request) interface{},error {
	productId,_:= strconv.Atoi(URLParam(r,"id")
	return GetProductoByIdRequest{ProductID: productId
	},nil
}
