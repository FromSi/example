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
	"time"
)

type GormUserRepository struct {
	Database *gorm.DB
}

func NewGormUserRepository(database *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		Database: database,
	}
}

func (repository *GormUserRepository) CreateIfNotExistsById(id string) error {
	model := models.GormUserModel{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return repository.Database.Where(models.GormUserModel{ID: id}).FirstOrCreate(&model).Error
}

func (repository *GormUserRepository) FindByFilterWithTrashed(filter filters.FindUserFilter) (*entities.User, error) {
	var userEntity *entities.User

	userModel := models.GormUserModel{ID: filter.ID}

	err := repository.Database.Unscoped().First(&userModel).Error

	if err != nil {
		return nil, err
	}

	userEntity, err = mappers.GormToEntityUser(&userModel)

	if err != nil {
		return nil, err
	}

	return userEntity, err
}

func (repository *GormUserRepository) GetAll(pageable entities.Pageable, sortable entities.Sortable) (*[]entities.User, error) {
	var userModels []models.GormUserModel
	var userEntities *[]entities.User

	offset := pageable.GetLimit() * (pageable.GetPage() - 1)
	query := repository.Database.Limit(pageable.GetLimit()).Offset(offset)

	for iterator := sortable.GetIterator(); iterator.HasNext(); {
		field, order := iterator.GetNext()

		query.Order(fmt.Sprintf("%s %s", field, order))
	}

	err := query.Find(&userModels).Error

	if err != nil {
		return nil, err
	}

	userEntities, err = mappers.ArrayGormToArrayEntityUser(&userModels)

	if err != nil {
		return nil, err
	}

	return userEntities, err
}

func (repository *GormUserRepository) GetTotal() (int, error) {
	var userModels []models.GormUserModel
	var total int64

	err := repository.Database.Model(&userModels).Count(&total).Error

	if err != nil {
		return 0, err
	}

	return int(total), err
}

func (repository *GormUserRepository) DeleteById(id string) error {
	return repository.Database.Delete(&models.GormUserModel{ID: id}).Error
}

func (repository *GormUserRepository) RestoreById(id string) error {
	return repository.Database.Unscoped().Model(&models.GormUserModel{ID: id}).Update("deleted_at", nil).Error
}

func (repository *GormUserRepository) Truncate() error {
	statement := &gorm.Statement{DB: repository.Database}
	err := statement.Parse(&models.GormUserModel{})

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
