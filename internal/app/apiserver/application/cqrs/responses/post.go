package responses

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
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
	Pageable entities.Pageable
}

type CqrsFindByIdQueryResponse struct {
	Data      QueryResponse
	IsDeleted bool
}
