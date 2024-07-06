package repositories

import "github.com/fromsi/example/internal/app/apiserver/domain/entities"

type MutablePostRepository interface {
	Create(post *entities.Post) error
	UpdateById(id string, post *entities.Post) error
	DeleteById(id string) error
	RestoreById(id string) error
}

type QueryPostRepository interface {
	FindByIdWithTrashed(id string) (*entities.Post, error)
	GetAll() (*[]entities.Post, error)
}
