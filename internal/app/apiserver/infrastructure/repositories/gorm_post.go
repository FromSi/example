package repositories

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"gorm.io/gorm"
	"time"
)

type GormPostRepository struct {
	Database *gorm.DB
}

type GormPostModel struct {
	ID        string `gorm:"primaryKey"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (repository *GormPostRepository) Create(post *entities.Post) error {
	return repository.Database.Create(ConvertEntityToModel(post)).Error
}

func (repository *GormPostRepository) UpdateById(id string, post *entities.Post) error {
	return repository.Database.Model(&GormPostModel{ID: id}).Updates(ConvertEntityToModel(post)).Error
}

func (repository *GormPostRepository) FindByIdWithTrashed(id string) (*entities.Post, error) {
	var postEntity *entities.Post

	postModel := GormPostModel{ID: id}

	err := repository.Database.Unscoped().First(&postModel).Error

	if err == nil {
		postEntity = ConvertModelToEntity(&postModel)
	}

	return postEntity, err
}

func (repository *GormPostRepository) GetAll() (*[]entities.Post, error) {
	var postModels []GormPostModel
	var postEntities *[]entities.Post

	err := repository.Database.Find(&postModels).Error

	if err == nil {
		postEntities = ConvertArrayModelToEntity(&postModels)
	}

	return postEntities, err
}

func (repository *GormPostRepository) DeleteById(id string) error {
	return repository.Database.Delete(&GormPostModel{ID: id}).Error
}

func (repository *GormPostRepository) RestoreById(id string) error {
	return repository.Database.Unscoped().Model(&GormPostModel{ID: id}).Update("deleted_at", nil).Error
}

func ConvertEntityToModel(entity *entities.Post) *GormPostModel {
	model := GormPostModel{
		ID:   entity.ID,
		Text: entity.Text,
	}

	if entity.CreatedAt != nil {
		model.CreatedAt = *entity.CreatedAt
	}

	if entity.UpdatedAt != nil {
		model.UpdatedAt = *entity.UpdatedAt
	}

	if entity.DeletedAt != nil {
		model.DeletedAt = &gorm.DeletedAt{
			Time:  *entity.DeletedAt,
			Valid: true,
		}
	}

	return &model
}

func ConvertArrayEntityToModel(entities *[]entities.Post) *[]GormPostModel {
	posts := []GormPostModel{}

	for _, item := range *entities {
		posts = append(posts, *ConvertEntityToModel(&item))
	}

	return &posts
}

func ConvertModelToEntity(model *GormPostModel) *entities.Post {
	entity := entities.Post{
		ID:        model.ID,
		Text:      model.Text,
		CreatedAt: &model.CreatedAt,
		UpdatedAt: &model.UpdatedAt,
	}

	if model.DeletedAt != nil {
		entity.DeletedAt = &model.DeletedAt.Time
	}

	return &entity
}

func ConvertArrayModelToEntity(models *[]GormPostModel) *[]entities.Post {
	posts := []entities.Post{}

	for _, item := range *models {
		posts = append(posts, *ConvertModelToEntity(&item))
	}

	return &posts
}
