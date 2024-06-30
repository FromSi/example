package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CreateRequest struct {
}

type IndexRequest struct {
}

type ShowRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type UpdateRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type DeleteRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ResetRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func main() {
	route := gin.Default()

	route.POST("/books", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "created",
		})
	})

	route.GET("/books", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "index",
		})
	})

	route.GET("/books/:id", func(context *gin.Context) {
		var request ShowRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "show",
		})
	})

	route.PATCH("/books/:id", func(context *gin.Context) {
		var request UpdateRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "updated",
		})
	})

	route.DELETE("/books/:id", func(context *gin.Context) {
		var request DeleteRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "deleted",
		})
	})

	route.POST("/books/:id", func(context *gin.Context) {
		var request ResetRequest

		if err := context.ShouldBindUri(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "rested",
		})
	})

	err := route.Run("localhost:8080")

	if err != nil {
		log.Fatal(err.Error())
	}
}
