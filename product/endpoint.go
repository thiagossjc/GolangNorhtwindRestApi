package product

import (
	"context"

	"github.com/GolangNorhtwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getProductBYIDRequest struct {
	ProdutoID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandartCost string
	ProductCode  string
	ProductName  string
}

type updateProductRequest struct {
	ID           int
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID string
}

type getBestellersRequest struct {
}

func MakeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductBYIDRequest)
		product, err := s.GetProductoById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		helper.Catch(err)
		return result, nil

	}
	return getProductsEndPoint
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	addProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		productId, err := s.InsertProduct(&req)
		helper.Catch(err)
		return productId, nil
	}
	return addProductEndPoint
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	updateProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		productId, err := s.UpdateProduct(&req)
		helper.Catch(err)
		return productId, nil
	}
	return updateProductEndPoint
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	deleteProductEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProduct(&req)
		helper.Catch(err)
		return result, nil
	}
	return deleteProductEndPoint
}

func makeBestellersEndPoint(s Service) endpoint.Endpoint {
	getBestSellersEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellers()
		helper.Catch(err)
		return result, nil
	}
	return getBestSellersEndPoint
}
