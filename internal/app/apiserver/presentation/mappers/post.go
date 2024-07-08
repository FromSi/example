package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	presentationresponses "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
	"github.com/gin-gonic/gin"
)

func ToGinShowPostResponse(post *responses.FindByIdQueryResponse) *presentationresponses.SuccessResponse {
	return &presentationresponses.SuccessResponse{
		Data: gin.H{
			"id":         post.Data.ID,
			"text":       post.Data.Text,
			"created_at": post.Data.CreatedAt,
			"updated_at": post.Data.UpdatedAt,
		},
	}
}

func ToGinIndexPostResponse(posts *responses.GetAllQueryResponse) *presentationresponses.SuccessArrayResponse {
	response := []presentationresponses.PostResponse{}

	for _, post := range (*posts).Data {
		response = append(response, presentationresponses.PostResponse{
			ID:        post.ID,
			Text:      post.Text,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return &presentationresponses.SuccessArrayResponse{
		Data: response,
	}
}
