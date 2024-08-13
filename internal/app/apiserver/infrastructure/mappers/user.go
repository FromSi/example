package mappers

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
	"github.com/fromsi/example/internal/app/apiserver/infrastructure/models"
	"gorm.io/gorm"
	"time"
)

func EntityToGormUser(entity *entities.User) (*models.GormUserModel, error) {
	if entity == nil {
		return nil, nil
	}

	model := models.GormUserModel{
		ID:        entity.ID.GetId(),
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

func ArrayEntityToArrayGormUser(entitySlice *[]entities.User) (*[]models.GormUserModel, error) {
	if entitySlice == nil {
		return nil, nil
	}

	modelSlice := []models.GormUserModel{}

	for _, item := range *entitySlice {
		model, err := EntityToGormUser(&item)

		if err != nil {
			return nil, err
		}

		modelSlice = append(modelSlice, *model)
	}

	return &modelSlice, nil
}

func GormToEntityUser(model *models.GormUserModel) (*entities.User, error) {
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

	entity, err := entities.NewUser(model.ID, &createdAtCopy, &updatedAtCopy, deletedAt)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func ArrayGormToArrayEntityUser(modelSlice *[]models.GormUserModel) (*[]entities.User, error) {
	if modelSlice == nil {
		return nil, nil
	}

	entitySlice := []entities.User{}

	for _, item := range *modelSlice {
		entity, err := GormToEntityUser(&item)

		if err != nil {
			return nil, err
		}

		entitySlice = append(entitySlice, *entity)
	}

	return &entitySlice, nil
}
