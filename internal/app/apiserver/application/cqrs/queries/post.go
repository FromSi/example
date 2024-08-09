package queries

import (
	"errors"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/mappers"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
)

type GetAllPostQuery struct {
	Pageable entities.Pageable
	Sortable entities.Sortable
}

type GetAllPostQueryHandler struct {
	QueryRepository *repositories.QueryRepository
}

func (handler GetAllPostQueryHandler) Handle(query Query) (any, error) {
	queryImplementation, exists := query.(GetAllPostQuery)

	if !exists {
		return nil, errors.New("invalid query type")
	}

	posts, err := handler.QueryRepository.PostRepository.GetAll(queryImplementation.Pageable, queryImplementation.Sortable)

	if err != nil {
		return nil, err
	}

	total, err := handler.QueryRepository.PostRepository.GetTotal()

	if err != nil {
		return nil, err
	}

	err = queryImplementation.Pageable.SetTotal(total)

	if err != nil {
		return nil, err
	}

	response, err := mappers.ToGetAllPostQueryResponse(posts, queryImplementation.Pageable)

	if err != nil {
		return nil, err
	}

	return response, nil
}

type FindByIdPostQuery struct {
	ID string
}

type FindByIdPostQueryHandler struct {
	QueryRepository *repositories.QueryRepository
}

func (handler FindByIdPostQueryHandler) Handle(query Query) (any, error) {
	queryImplementation, exists := query.(FindByIdPostQuery)

	if !exists {
		return nil, errors.New("invalid query type")
	}

	findPostFilter, err := filters.NewFindPostFilter(queryImplementation.ID)

	if err != nil {
		return nil, err
	}

	post, err := handler.QueryRepository.PostRepository.FindByFilterWithTrashed(*findPostFilter)

	if err != nil {
		return nil, err
	}

	response, err := mappers.ToFindByIdPostQueryResponse(post)

	if err != nil {
		return nil, err
	}

	return response, nil
}
