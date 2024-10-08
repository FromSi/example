package repositories

import (
	"errors"
	"fmt"
	"github.com/fromsi/example/cmd/apiserver/types"
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/domain/filters"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/mappers"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	Database *gorm.DB
}

func NewGormPostRepository(database *gorm.DB) *GormPostRepository {
	return &GormPostRepository{
		Database: database,
	}
}

func (repository *GormPostRepository) Create(post *entities.Post) error {
	model, err := mappers.EntityToGormPost(post)

	if err != nil {
		return err
	}

	return repository.Database.Create(model).Error
}

func (repository *GormPostRepository) UpdateById(id string, post *entities.Post) error {
	model, err := mappers.EntityToGormPost(post)

	if err != nil {
		return err
	}

	return repository.Database.Model(&models.GormPostModel{ID: id}).Updates(model).Error
}

func (repository *GormPostRepository) FindByFilterWithTrashed(filter filters.FindPostFilter) (*entities.Post, error) {
	var postEntity *entities.Post

	postModel := models.GormPostModel{ID: filter.ID}

	err := repository.Database.Unscoped().First(&postModel).Error

	if err != nil {
		return nil, err
	}

	postEntity, err = mappers.GormToEntityPost(&postModel)

	if err != nil {
		return nil, err
	}

	return postEntity, err
}

func (repository *GormPostRepository) GetAll(pageable entities.Pageable, sortable entities.Sortable) (*[]entities.Post, error) {
	var postModels []models.GormPostModel
	var postEntities *[]entities.Post

	offset := pageable.GetLimit() * (pageable.GetPage() - 1)
	query := repository.Database.Limit(pageable.GetLimit()).Offset(offset)

	for iterator := sortable.GetIterator(); iterator.HasNext(); {
		field, order := iterator.GetNext()

		query.Order(fmt.Sprintf("%s %s", field, order))
	}

	err := query.Find(&postModels).Error

	if err != nil {
		return nil, err
	}

	postEntities, err = mappers.ArrayGormToArrayEntityPost(&postModels)

	if err != nil {
		return nil, err
	}

	return postEntities, err
}

func (repository *GormPostRepository) GetTotal() (int, error) {
	var postModels []models.GormPostModel
	var total int64

	err := repository.Database.Model(&postModels).Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), err
}

func (repository *GormPostRepository) DeleteById(id string) error {
	return repository.Database.Delete(&models.GormPostModel{ID: id}).Error
}

func (repository *GormPostRepository) RestoreById(id string) error {
	return repository.Database.Unscoped().Model(&models.GormPostModel{ID: id}).Update("deleted_at", nil).Error
}

func (repository *GormPostRepository) Truncate() error {
	statement := &gorm.Statement{DB: repository.Database}
	err := statement.Parse(&models.GormPostModel{})

	if err != nil {
		return err
	}

	tableName := statement.Schema.Table

	switch repository.Database.Name() {
	case types.DatabaseSQLiteDatabaseType:
		repository.Database.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	default:
		return errors.New("database type is not found")
	}

	return nil
}
