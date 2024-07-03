package repositories

import "github.com/fromsi/example/internal/app/apiserver/domain/entities"

type PostRepository interface {
	Create(post *entities.Post) error
	UpdateById(id string, post *entities.Post) error
	FindByIdWithTrashed(id string) (*entities.Post, error)
	GetAll() (*[]entities.Post, error)
	DeleteById(id string) error
	RestoreById(id string) error
}
