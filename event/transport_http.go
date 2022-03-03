package event

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getEventsHandler := kithttp.NewServer(MakeGetEventEndPoint(s), getEventsRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getEventsHandler)

	getEventsByIdHandler := kithttp.NewServer(makeGetEventByIdEndPoint(s), getEventByIdRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/{id}", getEventsByIdHandler)

	AddEventHandler := kithttp.NewServer(makeAddStatuEndPoint(s), getaddEventRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", AddEventHandler)

	UpdateEventHandler := kithttp.NewServer(makeUpdateEventEndPoint(s), getUpdateEventRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", UpdateEventHandler)

	CancelEventHandler := kithttp.NewServer(makeCancelaEventEndPoint(s), getCancelEventRequestDecoder,
		kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", CancelEventHandler)

	return r
}

func getEventsRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	request := getEventsRequest{
		Limit:     0,
		Offset:    0,
		CountryId: 0,
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getEventByIdRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {

	return getEventByIDRequest{
		EventID: chi.URLParam(r, "id")}, nil
}

func getaddEventRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := addEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getUpdateEventRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getCancelEventRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return cancelEventRequest{
		EventId: chi.URLParam(r, "id"),
	}, nil
}
