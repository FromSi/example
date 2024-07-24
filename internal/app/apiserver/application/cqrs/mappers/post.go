package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
)

func ToCqrsGetAllQueryResponse(posts *[]entities.Post, pageable entities.Pageable) (*responses.CqrsGetAllQueryResponse, error) {
	if pageable == nil {
		entityPageable, err := entities.NewEntityPageable(
			entities.MinPageOrder,
			entities.MaxLimitItems,
			entities.MinTotalItems,
		)

		if err == nil {
			return nil, err
		}

		pageable = entityPageable
	}

	response := responses.CqrsGetAllQueryResponse{
		Pageable: pageable,
	}

	if posts == nil {
		return &response, nil
	}

	for _, post := range *posts {
		response.Data = append(response.Data, responses.QueryResponse{
			ID:        post.ID.GetId(),
			Text:      post.Text.GetText(),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return &response, nil
}

func ToCqrsFindByIdQueryResponse(post *entities.Post) (*responses.CqrsFindByIdQueryResponse, error) {
	if post == nil {
		return nil, nil
	}

	response := responses.CqrsFindByIdQueryResponse{
		Data: responses.QueryResponse{
			ID:        post.ID.GetId(),
			Text:      post.Text.GetText(),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		},
	}

	if post.DeletedAt != nil {
		response.IsDeleted = true
	}

	return &response, nil
}
