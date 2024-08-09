package requests

import (
	"github.com/gin-gonic/gin"
)

type GinShowLoginAuthRequest struct{}

func NewGinShowLoginAuthRequest(context *gin.Context) (*GinShowLoginAuthRequest, error) {
	var request GinShowLoginAuthRequest

	if err := context.ShouldBindUri(&request); err != nil {
		return nil, err
	}

	return &request, nil
}
