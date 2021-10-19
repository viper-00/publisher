package endpoints

import (
	"context"
	"errors"
	"go-kit-example/internal"
	"go-kit-example/pkg/database"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Set struct {
	AddEndpoint           endpoint.Endpoint
	GetEndpoint           endpoint.Endpoint
	UpdateEndpoint        endpoint.Endpoint
	RemoveEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc database.Service) Set {
	return Set{
		AddEndpoint:           MakeAddEndpoint(svc),
		GetEndpoint:           MakeGetEndpoint(svc),
		UpdateEndpoint:        MakeUpdateEndpoint(svc),
		RemoveEndpoint:        MakeRemoveEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func (s *Set) Add(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddEndpoint(ctx, AddRequest{Document: doc})
	if err != nil {
		return "", err
	}
	addResp := resp.(AddResponse)
	if addResp.Err != "" {
		return "", errors.New(addResp.Err)
	}
	return addResp.TicketID, nil
}

func MakeAddEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		ticketID, err := svc.Add(ctx, req.Document)
		if err != nil {
			return AddResponse{TicketID: ticketID, Err: err.Error()}, nil
		}
		return AddResponse{TicketID: ticketID, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{Filters: filters})
	if err != nil {
		return []internal.Document{}, err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return []internal.Document{}, errors.New(getResp.Err)
	}
	return getResp.Documents, nil
}

func MakeGetEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		docs, err := svc.Get(ctx, req.Filters...)
		if err != nil {
			return GetResponse{Documents: docs, Err: err.Error()}, nil
		}
		return GetResponse{docs, ""}, nil
	}
}

func (s *Set) Update(ctx context.Context, ticketID string, doc *internal.Document) (int, error) {
	resp, err := s.UpdateEndpoint(ctx, UpdateRequest{TicketID: ticketID, Document: doc})
	if err != nil {
		return http.StatusBadRequest, err
	}
	updateResp := resp.(UpdateResponse)
	if updateResp.Err != "" {
		return http.StatusConflict, errors.New(updateResp.Err)
	}
	return http.StatusOK, nil
}

func MakeUpdateEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		code, err := svc.Update(ctx, req.TicketID, req.Document)
		if err != nil {
			return UpdateResponse{Code: code, Err: err.Error()}, nil
		}
		return UpdateResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Remove(ctx context.Context, ticketID string) (int, error) {
	resp, err := s.RemoveEndpoint(ctx, RemoveRequest{TicketID: ticketID})
	removeResp := resp.(RemoveResponse)
	if err != nil {
		return removeResp.Code, err
	}
	if removeResp.Err != "" {
		return removeResp.Code, errors.New(removeResp.Err)
	}
	return removeResp.Code, nil
}

func MakeRemoveEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveRequest)
		code, err := svc.Remove(ctx, req.TicketID)
		if err != nil {
			return UpdateResponse{Code: code, Err: err.Error()}, nil
		}
		return UpdateResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	serviceStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return serviceStatusResp.Code, err
	}
	if serviceStatusResp.Err != "" {
		return serviceStatusResp.Code, errors.New(serviceStatusResp.Err)
	}
	return serviceStatusResp.Code, nil
}

func MakeServiceStatusEndpoint(svc database.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
