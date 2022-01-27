package order

import "github.com/GolangNorhtwindRestApi/helper"

type Service interface {
	GetOrderById(params *getOrderByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
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
