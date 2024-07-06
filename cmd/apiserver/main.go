package main

import (
	"flag"
	"fmt"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	configDirPath, exists := os.LookupEnv("GO_EXAMPLE_CONFIG_DIR_PATH")

	if !exists {
		flag.StringVar(&configDirPath, "config_dir_path", ".", "configuration file directory path")
		flag.Parse()
	}

	config, err := NewConfig(configDirPath)

	if err != nil {
		log.Fatalln(err.Error())
	}

	database, err := gorm.Open(sqlite.Open(config.Database.Connections.Sqlite.Dsn), &gorm.Config{})

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

	err = route.Run(fmt.Sprintf("%s:%d", config.App.Host, config.App.Port))

	if err != nil {
		log.Println(err.Error())
	}
}
