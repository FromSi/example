package requests

import (
	"github.com/gin-gonic/gin"
)

type GinCreatePostRequest struct {
	Body GinCreatePostRequestBody
	Text string `from:"text" binding:"required"`
}

func NewGinCreatePostRequest(context *gin.Context) (*GinCreatePostRequest, error) {
	var request GinCreatePostRequest

	requestBody, err := NewGinCreatePostRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type GinCreatePostRequestBody struct {
	Text string `from:"text" binding:"required"`
}

func NewGinCreatePostRequestBody(context *gin.Context) (*GinCreatePostRequestBody, error) {
	var requestBody GinCreatePostRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}

type GinIndexPostRequest struct {
	Pageable GinPageableRequest
}

func NewGinIndexPostRequest(context *gin.Context) (*GinIndexPostRequest, error) {
	var request GinIndexPostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		return nil, err
	}

	pageableRequest, err := NewGinPageableRequest(context)

	if err != nil {
		return nil, err
	}

	request.Pageable = *pageableRequest

	return &request, nil
}

type GinShowPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func NewGinShowPostRequest(context *gin.Context) (*GinShowPostRequest, error) {
	var request GinShowPostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		return nil, err
	}

	return &request, nil
}

type GinUpdatePostRequest struct {
	Body GinUpdatePostRequestBody
	ID   string `uri:"id" binding:"required,uuid"`
}

func NewGinUpdatePostRequest(context *gin.Context) (*GinUpdatePostRequest, error) {
	var request GinUpdatePostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		return nil, err
	}

	requestBody, err := NewGinUpdatePostRequestBody(context)

	if err != nil {
		return nil, err
	}

	request.Body = *requestBody

	return &request, nil
}

type GinUpdatePostRequestBody struct {
	Text *string `from:"text"`
}

func NewGinUpdatePostRequestBody(context *gin.Context) (*GinUpdatePostRequestBody, error) {
	var requestBody GinUpdatePostRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		return nil, err
	}

	return &requestBody, nil
}

type GinDeletePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func NewGinDeletePostRequest(context *gin.Context) (*GinDeletePostRequest, error) {
	var request GinDeletePostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		return nil, err
	}

	return &request, nil
}

type GinRestorePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func NewGinRestorePostRequest(context *gin.Context) (*GinRestorePostRequest, error) {
	var request GinRestorePostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		return nil, err
	}

	return &request, nil
}
