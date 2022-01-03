package product

import "github.com/go-kit/kit/endpoint"

type getProductBYIDRequest struct {
	ProdutoID int
}

func MakeGetProductByIdEndPoint(s ServiceS) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductBYIDRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}
	return getProductByIdEndpoint
}
