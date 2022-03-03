package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-chi/chi/v5"

	kitHttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	//ConsultaporId de Produto
	r := chi.NewRouter()
	getProductBYIDHandler := kitHttp.NewServer(MakeGetProductByIdEndPoint(s),
		GetProductoByIdRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getProductBYIDHandler)

	//ConsultaPaginadaProducto
	getProductsHandler := kitHttp.NewServer(makeGetProductsEndPoint(s), getProductsRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	//InsertProduct
	addProductHandler := kitHttp.NewServer(makeAddProductEndPoint(s), addProductRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	//updateProduct
	updateProductHandler := kitHttp.NewServer(makeUpdateProductEndPoint(s), updateProductRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	//deleteProduct
	deleteProductHandler := kitHttp.NewServer(makeDeleteProductEndPoint(s), getdeleteProductRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	//BestSellers
	getBestSellersHandler := kitHttp.NewServer(makeBestSellersEndPoint(s), getBestSellersRequestDecoder, kitHttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/bestSellers", getBestSellersHandler)

	return r

}

func GetProductoByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductBYIDRequest{ProdutoID: productId}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func updateProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getdeleteProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return deleteProductRequest{
		ProductID: chi.URLParam(r, "id"),
	}, nil
}

func getBestSellersRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getBestSellersRequest{}, nil
}
