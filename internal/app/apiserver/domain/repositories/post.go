package repositories

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
)

type MutablePostRepository interface {
	Create(*entities.Post) error
	UpdateById(string, *entities.Post) error
	DeleteById(string) error
	RestoreById(string) error
	Truncate() error
}

type QueryPostRepository interface {
	FindByFilterWithTrashed(filters.FindPostFilter) (*entities.Post, error)
	GetAll(entities.Pageable, entities.Sortable) (*[]entities.Post, error)
	GetTotal() (int, error)
}
