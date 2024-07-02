package main

import (
	"fmt"
	config2 "github.com/fromsi/example/configs"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/gin/controllers"
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

	err = database.AutoMigrate(&controllers.PostModel{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	route := gin.Default()

	controllers.InitBookController(route, database)

	err = route.Run(fmt.Sprintf("%s:%d", config.Host, config.Port))

	if err != nil {
		log.Println(err.Error())
	}
}
