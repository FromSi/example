package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"gorm.io/gorm"
	"time"
)

func EntityToGorm(entity *entities.Post) *models.GormPostModel {
	model := models.GormPostModel{
		ID:   entity.ID.GetId(),
		Text: entity.Text.GetText(),
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

func ArrayEntityToArrayGorm(entities *[]entities.Post) *[]models.GormPostModel {
	posts := []models.GormPostModel{}

	for _, item := range *entities {
		posts = append(posts, *EntityToGorm(&item))
	}

	return &posts
}

func GormToEntity(model *models.GormPostModel) (*entities.Post, error) {
	idValueObject, err := entities.NewId(model.ID)

	if err != nil {
		return nil, err
	}

	textValueObject, err := entities.NewText(model.Text)

	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time

	if model.DeletedAt != nil {
		deletedAtCopy := model.DeletedAt.Time
		deletedAt = &deletedAtCopy
	}

	createdAt := model.CreatedAt
	updatedAt := model.UpdatedAt

	entity := entities.Post{
		ID:        *idValueObject,
		Text:      *textValueObject,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		DeletedAt: deletedAt,
	}

	return &entity, nil
}

func ArrayGormToArrayEntity(models *[]models.GormPostModel) (*[]entities.Post, error) {
	posts := []entities.Post{}

	for _, item := range *models {
		post, err := GormToEntity(&item)

		if err != nil {
			return nil, err
		}

		posts = append(posts, *post)
	}

	return &posts, nil
}
