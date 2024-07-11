package cqrs

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/queries"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/pkg/cqrs"
)

type DefaultQueryCQRS struct {
	QueryRepository *repositories.QueryRepository
}

func (cqrs *DefaultQueryCQRS) Ask(query cqrs.Query) (any, error) {
	queryHandler, err := getQueryHandler(query, cqrs)

	if err != nil {
		return nil, err
	}

	return queryHandler.Handle(query)
}

func NewQueryCQRS(queryRepository *repositories.QueryRepository) cqrs.QueryCQRS {
	return &DefaultQueryCQRS{QueryRepository: queryRepository}
}

func getQueryHandler(query cqrs.Query, cqrs *DefaultQueryCQRS) (cqrs.QueryHandler, error) {
	switch query.(type) {
	case queries.GetAllQuery:
		return &queries.GetAllQueryHandler{QueryRepository: cqrs.QueryRepository}, nil
	case queries.FindByIdQuery:
		return &queries.FindByIdQueryHandler{QueryRepository: cqrs.QueryRepository}, nil
	}

	return nil, errors.New("query handler not found")
}