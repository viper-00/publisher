package watermark

import (
	"context"
	"net/http"
	"os"
	"publisher/internal"

	"github.com/go-kit/log"
	"github.com/google/uuid"
)

var logger log.Logger

type watermarkService struct{}

func NewService() Service {
	return &watermarkService{}
}

func (w *watermarkService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	// query the database using the filters and return the list of docuemtns
	// return errror if the filter (key) is invalid and also return error if no item found
	doc := internal.Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rowling",
		Topic:   "Fiction and Magic",
	}

	return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(_ context.Context, ticketID string) (internal.Status, error) {
	// query database using the ticketID and return the document info
	// return err if the ticketID is invalid or no Document exists for that ticketID
	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketID, mark string) (int, error) {
	// update the database entry with watermark field as non empty
	// first check if the watermark status is not already in InProgress, Started or Finished state
	// If yes, then return invalid request
	// return error if not item found using the ticketID
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *internal.Document) (string, error) {
	// add the document entry in the database by calling the database service
	// return error if the doc is invalid and/or the database invalid entry error
	newTicketID := uuid.New()
	return newTicketID.String(), nil
}

func (w *watermarkService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
