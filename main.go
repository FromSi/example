package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type CreateBookRequestBody struct {
	Text string `from:"text" binding:"required"`
}

type ShowBookRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBookRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateBookRequestBody struct {
	Text *string `from:"text"`
}

type DeleteBookRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ResetBookRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type PostModel struct {
	ID        string `gorm:"primaryKey"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func main() {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = database.AutoMigrate(&PostModel{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	route := gin.Default()

	route.POST("/books", func(context *gin.Context) {
		var requestBody CreateBookRequestBody

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

		postModel := &PostModel{
			ID:   id.String(),
			Text: requestBody.Text,
		}

		database.Create(postModel)

		context.Status(http.StatusAccepted)
	})

	route.GET("/books", func(context *gin.Context) {
		var postModels []PostModel

		database.Find(&postModels)

		context.JSON(http.StatusOK, gin.H{"data": &postModels})
	})

	route.GET("/books/:id", func(context *gin.Context) {
		var request ShowBookRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		postModel := PostModel{ID: request.ID}

		databaseResult := database.Unscoped().First(&postModel)

		if databaseResult.Error != nil {
			context.Status(http.StatusNotFound)

			log.Println(fmt.Sprintf("%s is not found!", request.ID))

			return
		}

		if postModel.DeletedAt != nil {
			context.Status(http.StatusGone)

			log.Println(fmt.Sprintf("%s is deleted!", request.ID))

			return
		}

		context.JSON(http.StatusOK, gin.H{"data": &postModel})
	})

	route.PATCH("/books/:id", func(context *gin.Context) {
		var request UpdateBookRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		var requestBody UpdateBookRequestBody

		if err := context.ShouldBind(&requestBody); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		var postModel PostModel

		if requestBody.Text != nil {
			postModel.Text = *requestBody.Text
		}

		databaseResult := database.Model(&PostModel{ID: request.ID}).Updates(&postModel)

		if databaseResult.Error != nil {
			log.Println(err.Error())
		}

		context.Status(http.StatusAccepted)
	})

	route.DELETE("/books/:id", func(context *gin.Context) {
		var request DeleteBookRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		databaseResult := database.Delete(&PostModel{ID: request.ID})

		if databaseResult.Error != nil {
			log.Println(err.Error())
		}

		context.Status(http.StatusAccepted)
	})

	route.POST("/books/:id", func(context *gin.Context) {
		var request ResetBookRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"data": err.Error(),
			})

			log.Println(err.Error())

			return
		}

		err = database.Unscoped().Model(&PostModel{ID: request.ID}).Update("deleted_at", nil).Error

		if err != nil {
			log.Println(err.Error())
		}

		context.Status(http.StatusAccepted)
	})

	err = route.Run("localhost:8080")

	if err != nil {
		log.Println(err.Error())
	}
}
