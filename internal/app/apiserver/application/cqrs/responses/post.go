package responses

import "time"

type QueryResponse struct {
	ID        string
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type GetAllQueryResponse struct {
	Data []QueryResponse
}

type FindByIdQueryResponse struct {
	Data      QueryResponse
	IsDeleted bool
}
