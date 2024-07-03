package controllers

import (
	"fmt"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	repositories "github.com/fromsi/example/internal/app/apiserver/domain/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type CreatePostRequestBody struct {
	Text string `from:"text" binding:"required"`
}

type ShowPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdatePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdatePostRequestBody struct {
	Text *string `from:"text"`
}

type DeletePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ResetPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func GinPostHandler(engine *gin.Engine, repository repositories.PostRepository) {
	engine.POST("/posts", func(context *gin.Context) {
		var requestBody CreatePostRequestBody

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

		post := &entities.Post{
			ID:   id.String(),
			Text: requestBody.Text,
		}

		err = repository.Create(post)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		context.Status(http.StatusAccepted)
	})

	engine.GET("/posts", func(context *gin.Context) {
		posts, err := repository.GetAll()

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		context.JSON(http.StatusOK, gin.H{"data": posts})
	})

	engine.GET("/posts/:id", func(context *gin.Context) {
		var request ShowPostRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		post, err := repository.FindByIdWithTrashed(request.ID)

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

		context.JSON(http.StatusOK, gin.H{"data": post})
	})

	engine.PATCH("/posts/:id", func(context *gin.Context) {
		var request UpdatePostRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		var requestBody UpdatePostRequestBody

		if err := context.ShouldBind(&requestBody); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		post := entities.Post{ID: request.ID}

		if requestBody.Text != nil {
			post.Text = *requestBody.Text
		}

		err := repository.UpdateById(request.ID, &post)

		if err != nil {
			log.Println(err)
		}

		context.Status(http.StatusAccepted)
	})

	engine.DELETE("/posts/:id", func(context *gin.Context) {
		var request DeletePostRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		err := repository.DeleteById(request.ID)

		if err != nil {
			log.Println(err)
		}

		context.Status(http.StatusAccepted)
	})

	engine.POST("/posts/:id", func(context *gin.Context) {
		var request ResetPostRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		err := repository.RestoreById(request.ID)

		if err != nil {
			log.Println(err.Error())
		}

		context.Status(http.StatusAccepted)
	})
}
