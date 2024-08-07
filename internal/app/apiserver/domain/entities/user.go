package entities

import (
	"github.com/fromsi/example/internal/pkg/tools"
	"time"
)

type User struct {
	ID        IdUser
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewUser(id string, createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time) (*User, error) {
	idValueObject, err := NewIdUser(id)

	if err != nil {
		return nil, err
	}

	user := User{
		ID: *idValueObject,
	}

	if createdAt != nil {
		createdAtCopy := *createdAt
		user.CreatedAt = &createdAtCopy
	}

	if updatedAt != nil {
		updatedAtCopy := *updatedAt
		user.UpdatedAt = &updatedAtCopy
	}

	if deletedAt != nil {
		deletedAtCopy := *deletedAt
		user.DeletedAt = &deletedAtCopy
	}

	return &user, nil
}

type IdUser struct {
	id string
}

func NewIdUser(id string) (*IdUser, error) {
	if id == "" {
		addressBTC := tools.NewAddressBTC()
		id = addressBTC.GetAddress()
	}

	return &IdUser{id: id}, nil
}

func (valueObject IdUser) GetId() string {
	return valueObject.id
}
