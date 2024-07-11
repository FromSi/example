package queries

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/mappers"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/pkg/cqrs"
	"github.com/fromsi/example/internal/pkg/data"
)

type GetAllQuery struct {
	Pageable data.Pageable
}

type GetAllQueryHandler struct {
	QueryRepository *repositories.QueryRepository
}

func (handler GetAllQueryHandler) Handle(query cqrs.Query) (any, error) {
	queryImplementation, exists := query.(GetAllQuery)

	if !exists {
		return nil, errors.New("invalid command type")
	}

	posts, err := handler.QueryRepository.PostRepository.GetAll(queryImplementation.Pageable)

	if err != nil {
		return nil, err
	}

	total, err := handler.QueryRepository.PostRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	queryImplementation.Pageable.SetTotal(total)

	return mappers.ToCqrsGetAllQueryResponse(posts, queryImplementation.Pageable), nil
}

type FindByIdQuery struct {
	ID string
}

type FindByIdQueryHandler struct {
	QueryRepository *repositories.QueryRepository
}

func (handler FindByIdQueryHandler) Handle(query cqrs.Query) (any, error) {
	queryImplementation, exists := query.(FindByIdQuery)

	if !exists {
		return nil, errors.New("invalid command type")
	}

	post, err := handler.QueryRepository.PostRepository.FindByIdWithTrashed(queryImplementation.ID)

	if err != nil {
		return nil, err
	}

	return mappers.ToCqrsFindByIdQueryResponse(post), nil
}