package event

import (
	"context"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-kit/kit/endpoint"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type getEventsRequest struct {
	Limit     int
	Offset    int
	CountryId int32
}

type getEventByIDRequest struct {
	EventID   string
	CountryId string
}

type addEventRequest struct {
	Description         string
	IdStatus            int32
	UserReg             int64
	DtReg               timestamp.Timestamp
	Cancel              bool
	Success             bool
	Traffic             bool
	InCustommer         bool
	InSupplier          bool
	OnPicking           bool
	OnPacking           bool
	OnExpedition        bool
	Finisher            bool
	Going_to            string
	IdCountry           int32
	InAttend            bool
	InLogisticOperation bool
}

type updateEventRequest struct {
	Id                  int32
	IdStatus            int32
	Description         string
	UserReg             int64
	DtReg               timestamp.Timestamp
	Cancel              bool
	Success             bool
	Traffic             bool
	InCustommer         bool
	InSupplier          bool
	OnPicking           bool
	OnPacking           bool
	OnExpedition        bool
	Finisher            bool
	Going_to            string
	IdPais              int32
	InAttend            bool
	InLogisticOperation bool
}

type cancelEventRequest struct {
	EventId string
}

// @Summary Lista de Eventos
// @Tags Events
// @Accept json
// @Produce json
// @Param request body event.getEventsRequest true "User Data"
// @Success 200 {object} event.EventList "ok"
// @Router /events/paginated [post]
func MakeGetEventEndPoint(s Service) endpoint.Endpoint {
	getEventsEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEventsRequest)
		result, err := s.GetEvents(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEventsEndPoint
}

// @Summary Event By Id
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Event Id"
// @Success 200 {object} event.Event "ok"
// @Router /events/{id} [get]
func makeGetEventByIdEndPoint(s Service) endpoint.Endpoint {
	getEventByIdEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEventByIDRequest)
		result, err := s.GetEventById(&req)
		helper.Catch(err)
		return result, nil
	}
	return getEventByIdEndPoint
}

// @Summary Insertar Event
// @Tags Events
// @Accept json
// @Produce json
// @Param request body event.addEventRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /events/ [post]
func makeAddStatuEndPoint(s Service) endpoint.Endpoint {
	addEventEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEventRequest)
		eventId, err := s.InsertEvent(&req)
		helper.Catch(err)
		return eventId, nil
	}
	return addEventEndPoint
}

// @Summary Actualizar Events
// @Tags Events
// @Accept json
// @Produce json
// @Param request body event.updateEventRequest true "User data"
// @Success 200 {integer} int "ok"
// @Router /events/ [put]
func makeUpdateEventEndPoint(s Service) endpoint.Endpoint {
	updateEventEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEventRequest)
		eventId, err := s.UpdateEvent(&req)
		helper.Catch(err)
		return eventId, nil
	}
	return updateEventEndPoint
}

// @Summary Cancelar Events
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Events ID"
// @Success 200 {integer} int "ok"
// @Router /events/{id} [delete]
func makeCancelaEventEndPoint(s Service) endpoint.Endpoint {
	cancelaEventEndPoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cancelEventRequest)
		eventId, err := s.CancelEvent(&req)
		helper.Catch(err)
		return eventId, nil
	}
	return cancelaEventEndPoint
}
