package status

import "github.com/GoGooliveryProviderAPI/helper"

type Service interface {
	GetStatus(params *getStatusRequest) (*StatusList, error)
	GetStatusById(params *getStatusByIDRequest) (*Status, error)
	InsertStatu(params *addStatusRequest) (int32, error)
	UpdateStatu(params *updateStatusRequest) (int32, error)
	CancelaStatu(params *cancelStatusRequest) (int32, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetStatus(params *getStatusRequest) (*StatusList, error) {
	status, err := s.repo.GetStatus(params)
	helper.Catch(err)
	//totalEmployees, err := s.repo.GetTotalEmployees()
	//	helper.Catch(err)
	return &StatusList{Data: status}, nil
}

func (s *service) GetStatusById(params *getStatusByIDRequest) (*Status, error) {
	return s.repo.GetStatusById(params)
}

func (s *service) InsertStatu(params *addStatusRequest) (int32, error) {
	return s.repo.InsertStatu(params)
}

func (s *service) UpdateStatu(params *updateStatusRequest) (int32, error) {
	return s.repo.UpdateStatu(params)
}

func (s *service) CancelaStatu(params *cancelStatusRequest) (int32, error) {
	return s.repo.CancelaStatu(params)
}
