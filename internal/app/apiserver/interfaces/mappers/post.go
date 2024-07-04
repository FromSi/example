package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/gin-gonic/gin"
)

func ToGinShowResponse(post *entities.Post) *gin.H {
	return &gin.H{
		"id":         post.ID.GetId(),
		"text":       post.Text.GetText(),
		"created_at": post.CreatedAt,
		"updated_at": post.UpdatedAt,
	}
}

func ToGinIndexResponse(posts *[]entities.Post) *[]gin.H {
	response := []gin.H{}

	for _, post := range *posts {
		response = append(response, gin.H{
			"id":         post.ID.GetId(),
			"text":       post.Text.GetText(),
			"created_at": post.CreatedAt,
			"updated_at": post.UpdatedAt,
		})
	}

	return &response
}
