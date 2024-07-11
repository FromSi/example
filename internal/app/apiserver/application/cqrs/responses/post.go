package responses

import (
	"github.com/fromsi/example/internal/pkg/data"
	"time"
)

type QueryResponse struct {
	ID        string
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type CqrsGetAllQueryResponse struct {
	Data     []QueryResponse
	Pageable data.Pageable
}

type CqrsFindByIdQueryResponse struct {
	Data      QueryResponse
	IsDeleted bool
}
