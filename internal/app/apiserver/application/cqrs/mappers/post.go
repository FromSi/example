package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/pkg/data"
)

func ToCqrsGetAllQueryResponse(posts *[]entities.Post, pageable data.Pageable) *responses.CqrsGetAllQueryResponse {
	response := responses.CqrsGetAllQueryResponse{
		Pageable: pageable,
	}

	for _, post := range *posts {
		response.Data = append(response.Data, responses.QueryResponse{
			ID:        post.ID.GetId(),
			Text:      post.Text.GetText(),
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return &response
}

func ToCqrsFindByIdQueryResponse(post *entities.Post) *responses.CqrsFindByIdQueryResponse {
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

	return &response
}
