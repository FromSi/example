package main

import (
	"fmt"
	"github.com/fromsi/example/cmd/apiserver/config"
	"github.com/fromsi/example/cmd/apiserver/database"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs"
	"github.com/fromsi/example/internal/app/apiserver/presentation/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
)

func main() {
	applicationConfig, err := config.NewConfig()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fxProvidersRelationDatabase, err := database.NewFXProvidersRelationDatabase(applicationConfig)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fx.
		New(
			fx.Supply(applicationConfig),
			fx.Provide(NewApplication),
			fx.Provide(fx.Annotate(cqrs.NewCommandCQRS, fx.As(new(cqrs.CommandCQRS)))),
			fx.Provide(fx.Annotate(cqrs.NewQueryCQRS, fx.As(new(cqrs.QueryCQRS)))),
			fxProvidersRelationDatabase,
			fx.Invoke(func(application *Application) {
				err = application.Run()

				if err != nil {
					log.Fatalln(err.Error())
				}
			}),
		).
		Run()
}

type Application struct {
	Config      *config.Config
	CommandCQRS cqrs.CommandCQRS
	QueryCQRS   cqrs.QueryCQRS
}

func NewApplication(config *config.Config, commandCQRS cqrs.CommandCQRS, queryCQRS cqrs.QueryCQRS) *Application {
	return &Application{
		Config:      config,
		CommandCQRS: commandCQRS,
		QueryCQRS:   queryCQRS,
	}
}

func (application Application) Run() error {
	route := gin.Default()

	authController := controllers.GinAuthController{
		Engine:      route,
		CommandCQRS: &application.CommandCQRS,
		QueryCQRS:   &application.QueryCQRS,
	}

	route.POST("/auth/login", authController.Login)

	route.GET("/auth/login", authController.ShowLogin)

	postController := controllers.GinPostController{
		Engine:      route,
		CommandCQRS: &application.CommandCQRS,
		QueryCQRS:   &application.QueryCQRS,
	}

	route.POST("/posts", postController.Create)

	route.GET("/posts", postController.Index)

	route.GET("/posts/:id", postController.Show)

	route.PATCH("/posts/:id", postController.Update)

	route.DELETE("/posts/:id", postController.Delete)

	route.POST("/posts/:id", postController.Restore)

	err := route.Run(fmt.Sprintf("%s:%d", application.Config.App.Host, application.Config.App.Port))

	if err != nil {
		return err
	}

	return nil
}
