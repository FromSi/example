package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	presentationresponses "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
)

func ToGinShowPostResponse(post *responses.FindByIdPostQueryResponse) (*presentationresponses.Response, error) {
	if post == nil {
		return nil, nil
	}

	return &presentationresponses.Response{
		Data: presentationresponses.PostResponse{
			ID:        post.Data.ID,
			Text:      post.Data.Text,
			CreatedAt: post.Data.CreatedAt,
			UpdatedAt: post.Data.UpdatedAt,
		},
	}, nil
}

func ToGinIndexPostResponse(posts *responses.GetAllPostQueryResponse) (*presentationresponses.ListResponse, error) {
	if posts == nil {
		return nil, nil
	}

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
	}, nil
}
