package repositories

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
)

type MutableUserRepository interface {
	CreateIfNotExistsById(string) error
	DeleteById(string) error
	RestoreById(string) error
	Truncate() error
}

type QueryUserRepository interface {
	FindByFilterWithTrashed(filters.FindUserFilter) (*entities.User, error)
	GetAll(entities.Pageable, entities.Sortable) (*[]entities.Post, error)
	GetTotal() (int, error)
}
