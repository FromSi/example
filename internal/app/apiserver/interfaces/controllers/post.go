package controllers

import (
	"fmt"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	repositories "github.com/fromsi/example/internal/app/apiserver/domain/repositories"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/mappers"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/requests"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type GinPostController struct {
	Engine     *gin.Engine
	Repository repositories.PostRepository
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

	id, err := uuid.NewRandom()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	idValueObject, err := entities.NewIdValueObject(id.String())

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	textValueObject, err := entities.NewTextValueObject(requestBody.Text)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	post := &entities.Post{
		ID:   *idValueObject,
		Text: *textValueObject,
	}

	err = controller.Repository.Create(post)

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
	posts, err := controller.Repository.GetAll()

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

	post, err := controller.Repository.FindByIdWithTrashed(request.ID)

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

	idValueObject, err := entities.NewIdValueObject(request.ID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	post := entities.Post{ID: *idValueObject}

	if requestBody.Text != nil {
		textValueObject, err := entities.NewTextValueObject(*requestBody.Text)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		post.Text = *textValueObject
	}

	err = controller.Repository.UpdateById(request.ID, &post)

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

	err := controller.Repository.DeleteById(request.ID)

	if err != nil {
		log.Println(err)
	}

	context.Status(http.StatusAccepted)
}

func (controller GinPostController) Reset(context *gin.Context) {
	var request requests.GinResetPostRequest

	if err := context.ShouldBindUri(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})

		log.Println(err.Error())

		return
	}

	err := controller.Repository.RestoreById(request.ID)

	if err != nil {
		log.Println(err.Error())
	}

	context.Status(http.StatusAccepted)
}
