package order

import "github.com/GoGooliveryProviderAPI/helper"

type Service interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
	InsertOrder(param *addOrderRequest) (int64, error)
	UpdateOrder(param *addOrderRequest) (int64, error)
	DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error)
	DeleteOrder(param *deleteOrderRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetOrderById(params *getOrderByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderById(params)
}

func (s *service) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(params)
	helper.Catch(err)
	totalOrders, err := s.repo.GetTotalOrders(params)
	helper.Catch(err)
	//var list *OrderList
	//list.Data = orders
	//list.TotalRecords = totalOrders
	//return list, err
	//o se puede hacer asi con menos codigo
	return &OrderList{Data: orders, TotalRecords: totalOrders}, nil
}

func (s *service) InsertOrder(param *addOrderRequest) (int64, error) {
	orderId, err := s.repo.InsertOrder(param)
	helper.Catch(err)

	for _, detail := range param.OrderDetails {
		detail.OrderId = orderId
		_, err := s.repo.InsertOrderDetal(detail)
		helper.Catch(err)
	}
	return orderId, nil
}

func (s *service) UpdateOrder(param *addOrderRequest) (int64, error) {
	orderId, err := s.repo.UpdateOrder(param)
	helper.Catch(err)

	for _, detail := range param.OrderDetails {
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

func (s service) DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error) {
	return s.repo.DeleteOrderDetail(param)
}

func (s service) DeleteOrder(param *deleteOrderRequest) (int64, error) {
	_, err := s.repo.DeleteOrderDetailbyOrderId(param)
	helper.Catch(err)
	return s.repo.DeleteOrder(param)
}
