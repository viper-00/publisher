package transport

import (
	"context"
	"publisher/api/v1/pb/auth"
	"publisher/pkg/authorization/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	login         grpctransport.Handler
	logout        grpctransport.Handler
	serviceStatus grpctransport.Handler

	// forward compatible implementations.
	auth.UnimplementedAuthorizationServer
}

func NewGRPCServer(ep endpoints.Set) auth.AuthorizationServer {
	return &grpcServer{
		login: grpctransport.NewServer(
			ep.LoginEndpoint,
			decodeGRPCLoginRequest,
			decodeGRPCLoginResponse,
		),
		logout: grpctransport.NewServer(
			ep.LogoutEndpoint,
			decodeGRPCLogoutRequest,
			decodeGRPCLogoutResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) Login(ctx context.Context, r *auth.LoginRequest) (*auth.LoginReply, error) {
	_, rep, err := g.login.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*auth.LoginReply), nil
}

func decodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*auth.LoginRequest)
	return endpoints.LoginRequest{Account: req.Account, Password: req.Password}, nil
}

func decodeGRPCLoginResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*auth.LoginReply)
	return endpoints.LoginResponse{Token: reply.Token, Err: reply.Err}, nil
}

func (g *grpcServer) Logout(ctx context.Context, r *auth.LogoutRequest) (*auth.LogoutReply, error) {
	_, rep, err := g.logout.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*auth.LogoutReply), nil
}

func decodeGRPCLogoutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*auth.LogoutRequest)
	return endpoints.LogoutRequest{Account: req.Account, Token: req.Token}, nil
}

func decodeGRPCLogoutResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*auth.LogoutReply)
	return endpoints.LogoutResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *auth.ServiceStatusRequest) (*auth.ServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*auth.ServiceStatusReply), nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	_ = grpcReq.(*auth.ServiceStatusRequest)
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*auth.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
