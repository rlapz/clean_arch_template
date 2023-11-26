package model

/*
 * Http
 */
type HttpResponse[T any] struct {
	Success    bool                    `json:"success"`
	Message    string                  `json:"message,omitempty"`
	Data       T                       `json:"data,omitempty"`
	Pagination *HttpPaginationResponse `json:"pagination,omitempty"`
}

type HttpPaginationResponse struct {
	Total   uint64 `json:"total"`
	Current uint64 `json:"current"`
	Count   uint64 `json:"count"`
	HasNext bool   `json:"has_next"`
}
