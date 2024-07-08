package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
)

func ToGetAllQueryResponse(posts *[]entities.Post) *responses.GetAllQueryResponse {
	response := responses.GetAllQueryResponse{}

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

func ToFindByIdQueryResponse(post *entities.Post) *responses.FindByIdQueryResponse {
	response := responses.FindByIdQueryResponse{
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
