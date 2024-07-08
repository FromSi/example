package controllers

import (
	"fmt"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/commands"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/queries"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/presentation/mappers"
	"github.com/fromsi/example/internal/app/apiserver/presentation/requests"
	"github.com/fromsi/example/internal/pkg/cqrs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GinPostController struct {
	Engine      *gin.Engine
	CommandCQRS *cqrs.CommandCQRS
	QueryCQRS   *cqrs.QueryCQRS
}

func (controller GinPostController) Create(context *gin.Context) {
	var requestBody requests.GinCreatePostRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	err := (*controller.CommandCQRS).Dispatch(commands.CreatePostCommand{Text: requestBody.Text})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Index(context *gin.Context) {
	postQueryResponse, err := (*controller.QueryCQRS).Ask(queries.GetAllQuery{})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	postQueryResponseImplementation, exists := postQueryResponse.(*responses.GetAllQueryResponse)

	if !exists {
		context.Status(http.StatusNotFound)

		log.Println("invalid command type")

		return
	}

	context.JSON(http.StatusOK, mappers.ToGinIndexPostResponse(postQueryResponseImplementation))
}

func (controller GinPostController) Show(context *gin.Context) {
	var request requests.GinShowPostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	postQueryResponse, err := (*controller.QueryCQRS).Ask(queries.FindByIdQuery{ID: request.ID})

	if err != nil {
		context.Status(http.StatusNotFound)

		log.Println(fmt.Sprintf("%s is not found!", request.ID))

		return
	}

	postQueryResponseImplementation, exists := postQueryResponse.(*responses.FindByIdQueryResponse)

	if !exists {
		context.Status(http.StatusNotFound)

		log.Println("invalid command type")

		return
	}

	if postQueryResponseImplementation.IsDeleted {
		context.Status(http.StatusGone)

		log.Println(fmt.Sprintf("%s is deleted!", request.ID))

		return
	}

	context.JSON(http.StatusOK, mappers.ToGinShowPostResponse(postQueryResponseImplementation))
}

func (controller GinPostController) Update(context *gin.Context) {
	var request requests.GinUpdatePostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	var requestBody requests.GinUpdatePostRequestBody

	if err := context.ShouldBind(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	err := (*controller.CommandCQRS).Dispatch(commands.UpdateByIdPostCommand{
		ID:   request.ID,
		Text: requestBody.Text,
	})

	if err != nil {
		log.Println(err)
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Delete(context *gin.Context) {
	var request requests.GinDeletePostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	err := (*controller.CommandCQRS).Dispatch(commands.DeletePostCommand{ID: request.ID})

	if err != nil {
		log.Println(err)
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Restore(context *gin.Context) {
	var request requests.GinRestorePostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	err := (*controller.CommandCQRS).Dispatch(commands.RestorePostCommand{ID: request.ID})

	if err != nil {
		log.Println(err.Error())
	}

	context.Status(http.StatusAccepted)
}
