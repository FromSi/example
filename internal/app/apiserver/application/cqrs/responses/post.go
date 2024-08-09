package responses

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"time"
)

type PostQueryResponse struct {
	ID        string
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type GetAllPostQueryResponse struct {
	Data     []PostQueryResponse
	Pageable entities.Pageable
}

type FindByIdPostQueryResponse struct {
	Data      PostQueryResponse
	IsDeleted bool
}
