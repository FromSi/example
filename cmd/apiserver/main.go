package main

import (
	"fmt"
	config2 "github.com/fromsi/example/configs"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	config, err := config2.NewBuilderConfig().Build()

	if err != nil {
		log.Fatalln(err.Error())
	}

	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = database.AutoMigrate(&models.GormPostModel{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	route := gin.Default()

	postRepository := repositories.GormPostRepository{
		Database: database,
	}

	postController := controllers.GinPostController{
		Engine:     route,
		Repository: &postRepository,
	}

	route.POST("/posts", postController.Create)

	route.GET("/posts", postController.Index)

	route.GET("/posts/:id", postController.Show)

	route.PATCH("/posts/:id", postController.Update)

	route.DELETE("/posts/:id", postController.Delete)

	route.POST("/posts/:id", postController.Reset)

	err = route.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))

	if err != nil {
		log.Println(err.Error())
	}
}
