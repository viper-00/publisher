package transport

import (
	"context"
	"publisher/internal"
	"publisher/pkg/watermark/endpoints"

	"publisher/api/v1/pb/watermark"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	addDocument   grpctransport.Handler
	watermark     grpctransport.Handler
	serviceStatus grpctransport.Handler
	// forward compatible implementations.
	watermark.UnimplementedWatermarkServer
}

func NewGRPCServer(ep endpoints.Set) watermark.WatermarkServer {
	return &grpcServer{
		get:           grpctransport.NewServer(ep.GetEndpoint, decodeGRPCGetRequest, decodeGRPCGetResponse),
		status:        grpctransport.NewServer(ep.StatusEndpoint, decodeGRPCStatusRequest, decodeGRPCStatusResponse),
		addDocument:   grpctransport.NewServer(ep.AddDocumentEndpoint, decodeGRPCAddDocumentRequest, decodeGRPCAddDocumentResponse),
		watermark:     grpctransport.NewServer(ep.WatermarkEndpoint, decodeGRPCWatermarkRequest, decodeGRPCWatermarkResponse),
		serviceStatus: grpctransport.NewServer(ep.ServiceStatusEndpoint, decodeGRPCServiceStatusRequest, decodeGRPCServiceStatusResponse),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *watermark.GetRequest) (*watermark.GetReply, error) {
	_, req, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return req.(*watermark.GetReply), nil
}

func (g *grpcServer) Status(ctx context.Context, r *watermark.StatusRequest) (*watermark.StatusReply, error) {
	_, req, err := g.status.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return req.(*watermark.StatusReply), nil
}

func (g *grpcServer) AddDocument(ctx context.Context, r *watermark.AddDocumentRequest) (*watermark.AddDocumentReply, error) {
	_, req, err := g.addDocument.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return req.(*watermark.AddDocumentReply), nil
}

func (g *grpcServer) Watermark(ctx context.Context, r *watermark.WatermarkRequest) (*watermark.WatermarkReply, error) {
	_, req, err := g.watermark.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return req.(*watermark.WatermarkReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *watermark.ServiceStatusRequest) (*watermark.ServiceStatusReply, error) {
	_, req, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return req.(*watermark.ServiceStatusReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcRequest interface{}) (interface{}, error) {
	req := grpcRequest.(*watermark.GetRequest)
	var filters []internal.Filter
	for _, f := range req.Filters {
		filters = append(filters, internal.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.GetRequest{Filters: filters}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.GetReply)
	var docs []internal.Document
	for _, d := range reply.Documents {
		doc := internal.Document{
			Title:     d.Title,
			Content:   d.Content,
			Author:    d.Author,
			Topic:     d.Topic,
			Watermark: d.Watermark,
		}
		docs = append(docs, doc)
	}
	return endpoints.GetResponse{Documents: docs, Err: reply.Err}, nil
}

func decodeGRPCStatusRequest(_ context.Context, grpcRequest interface{}) (interface{}, error) {
	req := grpcRequest.(*watermark.StatusRequest)
	return endpoints.StatusRequest{TicketID: req.TicketID}, nil
}

func decodeGRPCStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.StatusReply)
	return endpoints.StatusResponse{Status: internal.Status(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCAddDocumentRequest(_ context.Context, grpcRequest interface{}) (interface{}, error) {
	req := grpcRequest.(*watermark.AddDocumentRequest)
	doc := &internal.Document{
		Title:     req.Document.Title,
		Content:   req.Document.Content,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return endpoints.AddDocumentRequest{Document: doc}, nil
}

func decodeGRPCAddDocumentResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.AddDocumentReply)
	return endpoints.AddDocumentResponse{TicketID: reply.TicketID, Err: reply.Err}, nil
}

func decodeGRPCWatermarkRequest(_ context.Context, grpcRequest interface{}) (interface{}, error) {
	req := grpcRequest.(*watermark.WatermarkRequest)
	return endpoints.WatermarkRequest{TicketID: req.TicketID, Mark: req.Mark}, nil
}

func decodeGRPCWatermarkResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	req := grpcReply.(*watermark.WatermarkReply)
	return endpoints.WatermarkResponse{Code: int(req.Code), Err: req.Err}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcRequest interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	req := grpcReply.(*watermark.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(req.Code), Err: req.Err}, nil
}
