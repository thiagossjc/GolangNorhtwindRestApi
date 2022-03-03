package request

import "github.com/GoGooliveryProviderAPI/helper"

type Service interface {
	GetRequestsOrders(params *getRequestsRequest) ([]*RequestOrder, error)
	//	GetTotalRequestsOrders() (int64, error)
	GetRequestOrderById(params *getRequestBYIDRequest) (*RequestOrder, error)
	InsertRequestOrder(params *addRequestOrderRequest) (int64, error)
	UpdateRequestOrder(params *updateRequestOrderRequest) (int64, error)
	CancelRequestOrder(params *cancelRequestOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetRequestOrderById(params *getRequestBYIDRequest) (*RequestOrder, error) {
	return s.repo.GetRequestOrderById(params)
}

func (s *service) GetRequestsOrders(params *getRequestsRequest) ([]*RequestOrder, error) {
	orders, err := s.repo.GetRequestsOrders(params)
	helper.Catch(err)
	totalOrders, err := s.repo.GetTotalRequestsOrders()
	helper.Catch(err)
	//var list *OrderList
	//list.Data = orders
	//list.TotalRecords = totalOrders
	//return list, err
	//o se puede hacer asi con menos codigo
	return &OrderList{Data: orders, TotalRecords: totalOrders}, nil
}

func (s *service) InsertRequestOrder(params *addRequestOrderRequest) (int64, error) {
	orderId, err := s.repo.InsertRequestOrder(params)
	helper.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderId = orderId
		_, err := s.repo(detail)
		helper.Catch(err)
	}
	return orderId, nil
}

func (s *service) UpdateRequestOrder(params *updateRequestOrderRequest) (int64, error) {
	orderId, err := s.repo.UpdateRequestOrder(params)
	helper.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderId = orderId
		if detail.Id == 0 {
			_, err = s.repo.InsertOrderDetal(detail)
		} else {
			_, err = s.repo.UpdateOrderDetal(detail)
		}
		helper.Catch(err)

	}
	return orderId, nil
}

func (s service) CancelRequestOrder(params *cancelRequestOrderRequest) (int64, error) {
	_, err := s.repo.CancelRequestOrder(params)
	helper.Catch(err)
	return s.repo.CancelRequestOrder(params)
}
