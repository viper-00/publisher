package endpoints

import (
	"context"
	"errors"
	"publisher/pkg/authorization"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	LoginEndpoint         endpoint.Endpoint
	LogoutEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc authorization.Service) Set {
	return Set{
		LoginEndpoint:         MakeLoginEndpoint(svc),
		LogoutEndpoint:        MakeLogoutEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func MakeLoginEndpoint(auth authorization.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := auth.Login(ctx, req.Account, req.Password)
		if err != nil {
			return LoginResponse{Token: token, Err: err.Error()}, nil
		}
		return LoginResponse{Token: token, Err: ""}, nil
	}
}

func (s *Set) Login(ctx context.Context, account, password string) (string, error) {
	resp, err := s.LoginEndpoint(ctx, LoginRequest{Account: account, Password: password})
	if err != nil {
		return "", err
	}
	loginResp := resp.(LoginResponse)
	if loginResp.Err != "" {
		return "", errors.New(loginResp.Err)
	}
	return loginResp.Token, nil
}

func MakeLogoutEndpoint(auth authorization.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogoutRequest)
		code, err := auth.Logout(ctx, req.Account, req.Token)
		if err != nil {
			return LogoutResponse{Code: code, Err: err.Error()}, nil
		}
		return LogoutResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Logout(ctx context.Context, account, token string) (int, error) {
	resp, err := s.LogoutEndpoint(ctx, LogoutRequest{Account: account, Token: token})
	logoutResp := resp.(LogoutResponse)
	if err != nil {
		return logoutResp.Code, err
	}
	if logoutResp.Err != "" {
		return logoutResp.Code, errors.New(logoutResp.Err)
	}
	return logoutResp.Code, nil
}

func MakeServiceStatusEndpoint(auth authorization.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := auth.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
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
