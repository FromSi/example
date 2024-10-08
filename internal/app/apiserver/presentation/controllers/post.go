package controllers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/commands"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/queries"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/presentation/mappers"
	"github.com/fromsi/example/internal/app/apiserver/presentation/requests"
	presentationresponses "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
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
	request, err := requests.NewGinCreatePostRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	err = (*controller.CommandCQRS).Dispatch(commands.CreatePostCommand{Text: request.Body.Text})

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Index(context *gin.Context) {
	request, err := requests.NewGinIndexPostRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	pageable, err := entities.NewEntityPageable(request.Pageable.GetPage(), request.Pageable.GetLimit(), entities.MinTotalItems)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	sortable, err := entities.NewEntitySortable(request.Sortable.GetData())

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	postQueryResponse, err := (*controller.QueryCQRS).Ask(queries.GetAllPostQuery{Pageable: pageable, Sortable: sortable})

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	postQueryResponseImplementation, exists := postQueryResponse.(*responses.GetAllPostQueryResponse)

	if !exists {
		context.Status(http.StatusNotFound)

		log.Println("invalid query type")

		return
	}

	response, err := mappers.ToGinIndexPostResponse(postQueryResponseImplementation)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	context.JSON(http.StatusOK, response)
}

func (controller GinPostController) Show(context *gin.Context) {
	request, err := requests.NewGinShowPostRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	postQueryResponse, err := (*controller.QueryCQRS).Ask(queries.FindByIdPostQuery{ID: request.ID})

	if err != nil {
		context.Status(http.StatusNotFound)

		log.Printf("%s is not found!\n", request.ID)

		return
	}

	postQueryResponseImplementation, exists := postQueryResponse.(*responses.FindByIdPostQueryResponse)

	if !exists {
		context.Status(http.StatusNotFound)

		log.Println("invalid query type")

		return
	}

	if postQueryResponseImplementation.IsDeleted {
		context.Status(http.StatusGone)

		log.Printf("%s is deleted!\n", request.ID)

		return
	}

	response, err := mappers.ToGinShowPostResponse(postQueryResponseImplementation)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	context.JSON(http.StatusOK, response)
}

func (controller GinPostController) Update(context *gin.Context) {
	request, err := requests.NewGinUpdatePostRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	err = (*controller.CommandCQRS).Dispatch(commands.UpdateByIdPostCommand{
		ID:   request.ID,
		Text: request.Body.Text,
	})

	if err != nil {
		log.Println(err)
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Delete(context *gin.Context) {
	request, err := requests.NewGinDeletePostRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	err = (*controller.CommandCQRS).Dispatch(commands.DeletePostCommand{ID: request.ID})

	if err != nil {
		log.Println(err)
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Restore(context *gin.Context) {
	request, err := requests.NewGinRestorePostRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	err = (*controller.CommandCQRS).Dispatch(commands.RestorePostCommand{ID: request.ID})

	if err != nil {
		log.Println(err.Error())
	}

	context.Status(http.StatusAccepted)
}
