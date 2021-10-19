package database

import (
	"context"
	"net/http"
	"os"
	"publisher/internal"

	"github.com/go-kit/log"
)

type dbService struct{}

func NewService() Service {
	return &dbService{}
}

// implement interface;

func (d *dbService) Add(_ context.Context, doc *internal.Document) (string, error) {
	return "", nil
}

func (d *dbService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	return []internal.Document{}, nil
}

func (d *dbService) Update(_ context.Context, ticketID string, doc *internal.Document) (int, error) {
	return http.StatusOK, nil
}

func (d *dbService) Remove(_ context.Context, ticketID string) (int, error) {
	return http.StatusOK, nil
}

func (d *dbService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
