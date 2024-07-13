package requests

import "github.com/gin-gonic/gin"

type GinPageableRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func NewGinPageableRequest(context *gin.Context) (*GinPageableRequest, error) {
	var pageableRequest GinPageableRequest

	if err := context.ShouldBindQuery(&pageableRequest); err != nil {
		return nil, err
	}

	return &pageableRequest, nil
}

func (request GinPageableRequest) GetPage() int {
	return request.Page
}

func (request GinPageableRequest) GetLimit() int {
	return request.Limit
}
