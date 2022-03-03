package event

import "github.com/GoGooliveryProviderAPI/helper"

type Service interface {
	GetEvents(params *getEventsRequest) (*EventList, error)
	GetEventById(params *getEventByIDRequest) (*Event, error)
	InsertEvent(params *addEventRequest) (int32, error)
	UpdateEvent(params *updateEventRequest) (int32, error)
	CancelEvent(params *cancelEventRequest) (int32, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetEvents(params *getEventsRequest) (*EventList, error) {
	events, err := s.repo.GetEvents(params)
	helper.Catch(err)
	//totalEmployees, err := s.repo.GetTotalEmployees()
	//	helper.Catch(err)
	return &EventList{Data: events}, nil
}

func (s *service) GetEventById(params *getEventByIDRequest) (*Event, error) {
	return s.repo.GetEventById(params)
}

func (s *service) InsertEvent(params *addEventRequest) (int32, error) {
	return s.repo.InsertEvent(params)
}

func (s *service) UpdateEvent(params *updateEventRequest) (int32, error) {
	return s.repo.UpdateEvent(params)
}

func (s *service) CancelEvent(params *cancelEventRequest) (int32, error) {
	return s.repo.CancelaEvent(params)
}
