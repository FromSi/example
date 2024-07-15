package entities

import (
	"errors"
	"github.com/fromsi/example/internal/pkg/tools"
)

type Sortable interface {
	GetIterator() *tools.MapStringIterator
}

type EntitySortable struct {
	Data map[string]string
}

func NewEntitySortable(data map[string]string) (*EntitySortable, error) {
	for _, order := range data {
		if order != "asc" && order != "desc" {
			return nil, errors.New("sort order must be desc or asc")
		}
	}

	return &EntitySortable{
		Data: data,
	}, nil
}

func (request EntitySortable) GetIterator() *tools.MapStringIterator {
	return tools.NewMapStringIterator(request.Data)
}
