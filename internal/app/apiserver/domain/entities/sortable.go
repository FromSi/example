package entities

import (
	"fmt"
	"github.com/fromsi/example/internal/pkg/tools"
)

const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

type Sortable interface {
	GetIterator() *tools.MapStringIterator
}

type EntitySortable struct {
	Data map[string]string
}

func NewEntitySortable(data map[string]string) (*EntitySortable, error) {
	for _, order := range data {
		if order != OrderAsc && order != OrderDesc {
			return nil, fmt.Errorf("sort order must be %s or %s", OrderAsc, OrderDesc)
		}
	}

	return &EntitySortable{
		Data: data,
	}, nil
}

func (request EntitySortable) GetIterator() *tools.MapStringIterator {
	return tools.NewMapStringIterator(request.Data)
}
