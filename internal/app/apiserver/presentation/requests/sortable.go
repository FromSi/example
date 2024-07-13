package requests

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type GinSortableRequest struct {
	Data map[string]string
}

func NewGinSortableRequest(context *gin.Context) (*GinSortableRequest, error) {
	sortableRequest := GinSortableRequest{
		Data: context.QueryMap("sort"),
	}

	for _, order := range sortableRequest.Data {
		if order != "asc" && order != "desc" {
			return nil, errors.New("sort order must be desc or asc")
		}
	}

	return &sortableRequest, nil
}

func (request GinSortableRequest) GetData() map[string]string {
	return request.Data
}
