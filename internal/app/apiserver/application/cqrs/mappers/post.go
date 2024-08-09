package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
)

func ToGetAllPostQueryResponse(posts *[]entities.Post, pageable entities.Pageable) (*responses.GetAllPostQueryResponse, error) {
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

	response := responses.GetAllPostQueryResponse{
		Pageable: pageable,
	}

	if posts == nil {
		return &response, nil
	}

	for _, post := range *posts {
		response.Data = append(response.Data, responses.PostQueryResponse{
			ID:        post.ID.GetId(),
			Text:      post.Text.GetText(),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return &response, nil
}

func ToFindByIdPostQueryResponse(post *entities.Post) (*responses.FindByIdPostQueryResponse, error) {
	if post == nil {
		return nil, nil
	}

	response := responses.FindByIdPostQueryResponse{
		Data: responses.PostQueryResponse{
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
