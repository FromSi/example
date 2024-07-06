package controllers

import (
	"fmt"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/commands"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/mappers"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/requests"
	"github.com/fromsi/example/internal/pkg/cqrs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GinPostController struct {
	Engine          *gin.Engine
	CommandCQRS     *cqrs.CommandCQRS
	QueryCQRS       *cqrs.QueryCQRS
	QueryRepository *repositories.QueryRepository
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
	posts, err := controller.QueryRepository.PostRepository.GetAll()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	response := mappers.ToGinIndexResponse(posts)

	context.JSON(http.StatusOK, gin.H{"data": response})
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

	post, err := controller.QueryRepository.PostRepository.FindByIdWithTrashed(request.ID)

	if err != nil {
		context.Status(http.StatusNotFound)

		log.Println(fmt.Sprintf("%s is not found!", request.ID))

		return
	}

	if post.DeletedAt != nil {
		context.Status(http.StatusGone)

		log.Println(fmt.Sprintf("%s is deleted!", request.ID))

		return
	}

	response := mappers.ToGinShowResponse(post)

	context.JSON(http.StatusOK, gin.H{"data": response})
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
