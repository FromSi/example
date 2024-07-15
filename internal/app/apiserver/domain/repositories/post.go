package repositories

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
)

type MutablePostRepository interface {
	Create(*entities.Post) error
	UpdateById(string, *entities.Post) error
	DeleteById(string) error
	RestoreById(string) error
}

type QueryPostRepository interface {
	FindByIdWithTrashed(string) (*entities.Post, error)
	GetAll(entities.Pageable, entities.Sortable) (*[]entities.Post, error)
	GetTotal() (int, error)
}
