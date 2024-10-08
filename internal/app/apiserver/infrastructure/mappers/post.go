package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"gorm.io/gorm"
	"time"
)

func EntityToGormPost(entity *entities.Post) (*models.GormPostModel, error) {
	if entity == nil {
		return nil, nil
	}

	model := models.GormPostModel{
		ID:        entity.ID.GetId(),
		Text:      entity.Text.GetText(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	return &model, nil
}

func ArrayEntityToArrayGormPost(entitySlice *[]entities.Post) (*[]models.GormPostModel, error) {
	if entitySlice == nil {
		return nil, nil
	}

	modelSlice := []models.GormPostModel{}

	for _, item := range *entitySlice {
		model, err := EntityToGormPost(&item)

		if err != nil {
			return nil, err
		}

		modelSlice = append(modelSlice, *model)
	}

	return &modelSlice, nil
}

func GormToEntityPost(model *models.GormPostModel) (*entities.Post, error) {
	if model == nil {
		return nil, nil
	}

	createdAtCopy := model.CreatedAt
	updatedAtCopy := model.UpdatedAt

	var deletedAt *time.Time

	if model.DeletedAt != nil {
		deletedAtCopy := model.DeletedAt.Time
		deletedAt = &deletedAtCopy
	}

	entity, err := entities.NewPost(model.ID, model.Text, &createdAtCopy, &updatedAtCopy, deletedAt)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func ArrayGormToArrayEntityPost(modelSlice *[]models.GormPostModel) (*[]entities.Post, error) {
	if modelSlice == nil {
		return nil, nil
	}

	entitySlice := []entities.Post{}

	for _, item := range *modelSlice {
		entity, err := GormToEntityPost(&item)

		if err != nil {
			return nil, err
		}

		entitySlice = append(entitySlice, *entity)
	}

	return &entitySlice, nil
}
