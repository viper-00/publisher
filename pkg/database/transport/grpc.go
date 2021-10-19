package transport

import (
	"context"
	"publisher/api/v1/pb/db"
	"publisher/internal"
	"publisher/pkg/database/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	add           grpctransport.Handler
	get           grpctransport.Handler
	update        grpctransport.Handler
	remove        grpctransport.Handler
	serviceStatus grpctransport.Handler
	// forward compatible implementations.
	db.UnimplementedDatabaseServer
}

func NewGRPCServer(ep endpoints.Set) db.DatabaseServer {
	return &grpcServer{
		add: grpctransport.NewServer(
			ep.AddEndpoint,
			decodeGRPCAddRequest,
			decodeGRPCAddResponse,
		),
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		update: grpctransport.NewServer(
			ep.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			decodeGRPCUpdateResponse,
		),
		remove: grpctransport.NewServer(
			ep.RemoveEndpoint,
			decodeGRPCRemoveRequest,
			decodeGRPCRemoveResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) Add(ctx context.Context, r *db.AddRequest) (*db.AddReply, error) {
	_, rep, err := g.add.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.AddReply), nil
}

func decodeGRPCAddRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.AddRequest)
	doc := &internal.Document{
		Content:   req.Document.Content,
		Title:     req.Document.Title,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return endpoints.AddRequest{Document: doc}, nil
}

func decodeGRPCAddResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.AddReply)
	return endpoints.AddResponse{TicketID: reply.TicketID, Err: reply.Err}, nil
}

func (g *grpcServer) Get(ctx context.Context, r *db.GetRequest) (*db.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.GetReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.GetRequest)
	var filters []internal.Filter
	for _, f := range req.Filters {
		filters = append(filters, internal.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.GetRequest{Filters: filters}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.GetReply)
	var docs []internal.Document
	for _, d := range reply.Documents {
		doc := internal.Document{
			Content:   d.Content,
			Title:     d.Title,
			Author:    d.Author,
			Topic:     d.Topic,
			Watermark: d.Watermark,
		}
		docs = append(docs, doc)
	}
	return endpoints.GetResponse{Documents: docs, Err: reply.Err}, nil
}

func (g *grpcServer) Update(ctx context.Context, r *db.UpdateRequest) (*db.UpdateReply, error) {
	_, rep, err := g.update.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.UpdateReply), nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.UpdateRequest)
	doc := &internal.Document{
		Content:   req.Document.Content,
		Title:     req.Document.Title,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return endpoints.UpdateRequest{TicketID: req.TicketID, Document: doc}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.UpdateReply)
	return endpoints.UpdateResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func (g *grpcServer) Remove(ctx context.Context, r *db.RemoveRequest) (*db.RemoveReply, error) {
	_, rep, err := g.remove.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.RemoveReply), nil
}

func decodeGRPCRemoveRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.RemoveRequest)
	return endpoints.RemoveRequest{TicketID: req.TicketID}, nil
}

func decodeGRPCRemoveResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.RemoveReply)
	return endpoints.RemoveResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *db.ServiceStatusRequest) (*db.ServiceStatusReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.ServiceStatusReply), nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
