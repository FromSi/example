package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	presentationresponses "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
)

func ToGinShowPostResponse(post *responses.CqrsFindByIdQueryResponse) *presentationresponses.Response {
	return &presentationresponses.Response{
		Data: presentationresponses.PostResponse{
			ID:        post.Data.ID,
			Text:      post.Data.Text,
			CreatedAt: post.Data.CreatedAt,
			UpdatedAt: post.Data.UpdatedAt,
		},
	}
}

func ToGinIndexPostResponse(posts *responses.CqrsGetAllQueryResponse) *presentationresponses.ListResponse {
	response := []presentationresponses.PostResponse{}

	for _, post := range (*posts).Data {
		response = append(response, presentationresponses.PostResponse{
			ID:        post.ID,
			Text:      post.Text,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return &presentationresponses.ListResponse{
		Data:     response,
		Pageable: presentationresponses.NewPageableResponse(posts.Pageable),
	}
}
