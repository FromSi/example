package repositories

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/pkg/data"
)

type MutablePostRepository interface {
	Create(*entities.Post) error
	UpdateById(string, *entities.Post) error
	DeleteById(string) error
	RestoreById(string) error
}

type QueryPostRepository interface {
	FindByIdWithTrashed(string) (*entities.Post, error)
	GetAll(data.Pageable) (*[]entities.Post, error)
	GetTotal() (int, error)
}
