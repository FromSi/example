package repositories

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/mappers"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	Database *gorm.DB
}

func (repository *GormPostRepository) Create(post *entities.Post) error {
	return repository.Database.Create(mappers.EntityToGorm(post)).Error
}

func (repository *GormPostRepository) UpdateById(id string, post *entities.Post) error {
	return repository.Database.Model(&models.GormPostModel{ID: id}).Updates(mappers.EntityToGorm(post)).Error
}

func (repository *GormPostRepository) FindByIdWithTrashed(id string) (*entities.Post, error) {
	var postEntity *entities.Post

	postModel := models.GormPostModel{ID: id}

	err := repository.Database.Unscoped().First(&postModel).Error

	if err != nil {
		return nil, err
	}

	postEntity, err = mappers.GormToEntity(&postModel)

	if err != nil {
		return nil, err
	}

	return postEntity, err
}

func (repository *GormPostRepository) GetAll() (*[]entities.Post, error) {
	var postModels []models.GormPostModel
	var postEntities *[]entities.Post

	err := repository.Database.Find(&postModels).Error

	if err != nil {
		return nil, err
	}

	postEntities, err = mappers.ArrayGormToArrayEntity(&postModels)

	if err != nil {
		return nil, err
	}

	return postEntities, err
}

func (repository *GormPostRepository) DeleteById(id string) error {
	return repository.Database.Delete(&models.GormPostModel{ID: id}).Error
}

func (repository *GormPostRepository) RestoreById(id string) error {
	return repository.Database.Unscoped().Model(&models.GormPostModel{ID: id}).Update("deleted_at", nil).Error
}
