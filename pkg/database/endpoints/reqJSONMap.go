package endpoints

import "publisher/internal"

type AddRequest struct {
	Document *internal.Document `json:"document"`
}

type AddResponse struct {
	TicketID string `json:"ticketID"`
	Err      string `json:"err"`
}

type GetRequest struct {
	Filters []internal.Filter `json:"filters,omitempty"`
}

type GetResponse struct {
	Documents []internal.Document `json:"documents"`
	Err       string              `json:"err,omitempty"`
}

type UpdateRequest struct {
	TicketID string             `json:"ticketID"`
	Document *internal.Document `json:"document"`
}

type UpdateResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}

type RemoveRequest struct {
	TicketID string `json:"ticketID"`
}

type RemoveResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}
