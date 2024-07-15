package cqrs

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/queries"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
)

type QueryCQRS interface {
	Ask(queries.Query) (any, error)
}

type DefaultQueryCQRS struct {
	QueryRepository *repositories.QueryRepository
}

func NewQueryCQRS(queryRepository *repositories.QueryRepository) QueryCQRS {
	return &DefaultQueryCQRS{QueryRepository: queryRepository}
}

func (cqrs *DefaultQueryCQRS) Ask(query queries.Query) (any, error) {
	queryHandler, err := getQueryHandler(query, cqrs)

	if err != nil {
		return nil, err
	}

	return queryHandler.Handle(query)
}

func getQueryHandler(query queries.Query, cqrs *DefaultQueryCQRS) (queries.QueryHandler, error) {
	switch query.(type) {
	case queries.GetAllQuery:
		return &queries.GetAllQueryHandler{QueryRepository: cqrs.QueryRepository}, nil
	case queries.FindByIdQuery:
		return &queries.FindByIdQueryHandler{QueryRepository: cqrs.QueryRepository}, nil
	}

	return nil, errors.New("query handler not found")
}
