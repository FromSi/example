package main

import (
	"fmt"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs"
	domairepository "github.com/fromsi/example/internal/app/apiserver/domain/repositories"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/repositories"
	"github.com/fromsi/example/internal/app/apiserver/interfaces/controllers"
	pkgcqrs "github.com/fromsi/example/internal/pkg/cqrs"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	fx.
		New(
			fx.Provide(NewApplication),
			fx.Provide(NewConfig),
			fx.Provide(repositories.NewMutableRepository),
			fx.Provide(repositories.NewQueryRepository),
			fx.Provide(fx.Annotate(cqrs.NewCommandCQRS, fx.As(new(pkgcqrs.CommandCQRS)))),
			fx.Provide(fx.Annotate(cqrs.NewQueryCQRS, fx.As(new(pkgcqrs.QueryCQRS)))),
			fx.Provide(func(config *Config) (*gorm.DB, error) {
				return gorm.Open(sqlite.Open(config.Database.Connections.Sqlite.Dsn), &gorm.Config{})
			}),
			fx.Provide(
				fx.Annotate(repositories.NewGormPostRepository,
					fx.As(new(domairepository.QueryPostRepository)),
					fx.As(new(domairepository.MutablePostRepository)),
				),
			),
			fx.Invoke(func(application *Application) {
				err := application.Run()

				if err != nil {
					log.Fatalln(err.Error())
				}
			}),
		).
		Run()
}

type Application struct {
	Config          *Config
	Database        *gorm.DB
	CommandCQRS     pkgcqrs.CommandCQRS
	QueryCQRS       pkgcqrs.QueryCQRS
	QueryRepository *repositories.QueryRepository
}

func NewApplication(config *Config, database *gorm.DB, commandCQRS pkgcqrs.CommandCQRS, queryCQRS pkgcqrs.QueryCQRS, queryRepository *repositories.QueryRepository) *Application {
	return &Application{
		Config:          config,
		Database:        database,
		CommandCQRS:     commandCQRS,
		QueryCQRS:       queryCQRS,
		QueryRepository: queryRepository,
	}
}

func (application Application) Run() error {
	err := application.Database.AutoMigrate(&models.GormPostModel{})

	if err != nil {
		return err
	}

	route := gin.Default()

	postController := controllers.GinPostController{
		Engine:          route,
		CommandCQRS:     &application.CommandCQRS,
		QueryCQRS:       &application.QueryCQRS,
		QueryRepository: application.QueryRepository,
	}

	route.POST("/posts", postController.Create)

	route.GET("/posts", postController.Index)

	route.GET("/posts/:id", postController.Show)

	route.PATCH("/posts/:id", postController.Update)

	route.DELETE("/posts/:id", postController.Delete)

	route.POST("/posts/:id", postController.Restore)

	err = route.Run(fmt.Sprintf("%s:%d", application.Config.App.Host, application.Config.App.Port))

	if err != nil {
		return err
	}

	return nil
}
