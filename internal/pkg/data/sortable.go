package data

import (
	"errors"
	"github.com/fromsi/example/internal/pkg/requests"
	"github.com/fromsi/example/internal/pkg/tools"
)

type Sortable interface {
	GetIterator() *tools.MapStringIterator
}

type EntitySortable struct {
	Data map[string]string
}

func NewEntitySortable(sortable requests.SortableRequest) (*EntitySortable, error) {
	for _, order := range sortable.GetData() {
		if order != "asc" && order != "desc" {
			return nil, errors.New("sort order must be desc or asc")
		}
	}

	return &EntitySortable{
		Data: sortable.GetData(),
	}, nil
}

func (request EntitySortable) GetIterator() *tools.MapStringIterator {
	return tools.NewMapStringIterator(request.Data)
}
