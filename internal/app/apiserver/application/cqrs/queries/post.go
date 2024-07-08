package queries

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/mappers"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/pkg/cqrs"
)

type GetAllQuery struct {
}

type GetAllQueryHandler struct {
	QueryRepository *repositories.QueryRepository
}

func (handler GetAllQueryHandler) Handle(query cqrs.Query) (any, error) {
	_, exists := query.(GetAllQuery)

	if !exists {
		return nil, errors.New("invalid command type")
	}

	posts, err := handler.QueryRepository.PostRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return mappers.ToGetAllQueryResponse(posts), nil
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

	return mappers.ToFindByIdQueryResponse(post), nil
}
