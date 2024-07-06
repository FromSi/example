package cqrs

import (
	"errors"
	"github.com/fromsi/example/internal/pkg/cqrs"
)

type DefaultQueryCQRS struct {
}

func (cqrs *DefaultQueryCQRS) Ask(query cqrs.Query) (any, error) {
	queryHandler, err := getQueryHandler(query)

	if err != nil {
		return nil, err
	}

	return queryHandler.Handle(query)
}

func NewQueryCQRS() cqrs.QueryCQRS {
	return &DefaultQueryCQRS{}
}

func getQueryHandler(query cqrs.Query) (cqrs.QueryHandler, error) {
	return nil, errors.New("query handler not found")
}
