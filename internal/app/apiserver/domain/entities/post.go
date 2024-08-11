package entities

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

const (
	TextPostMinLength = 3
	TextPostMaxLength = 255
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Post struct {
	ID        IdPost
	Text      TextPost
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewPost(id string, text string, createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time) (*Post, error) {
	idValueObject, err := NewIdPost(id)

	if err != nil {
		return nil, err
	}

	textValueObject, err := NewTextPost(text)

	if err != nil {
		return nil, err
	}

	post := Post{
		ID:   *idValueObject,
		Text: *textValueObject,
	}

	if createdAt != nil {
		createdAtCopy := *createdAt
		post.CreatedAt = &createdAtCopy
	}

	if updatedAt != nil {
		updatedAtCopy := *updatedAt
		post.UpdatedAt = &updatedAtCopy
	}

	if deletedAt != nil {
		deletedAtCopy := *deletedAt
		post.DeletedAt = &deletedAtCopy
	}

	return &post, nil
}

func (post *Post) SetText(text string) error {
	textValueObject, err := NewTextPost(text)

	if err != nil {
		return err
	}

	post.Text = *textValueObject

	return nil
}

type IdPost struct {
	id string
}

func NewIdPost(id string) (*IdPost, error) {
	if id == "" {
		newUUID, err := uuid.NewRandom()

		if err != nil {
			return nil, err
		}

		id = newUUID.String()
	}

	err := validate.Var(id, "required,uuid")

	if err != nil {
		return nil, err
	}

	return &IdPost{id: id}, nil
}

func (valueObject IdPost) GetId() string {
	return valueObject.id
}

type TextPost struct {
	text string
}

func NewTextPost(text string) (*TextPost, error) {
	err := validate.Var(text, fmt.Sprintf("required,gte=%d,lte=%d", TextPostMinLength, TextPostMaxLength))

	if err != nil {
		return nil, err
	}

	return &TextPost{text: text}, nil
}

func (valueObject TextPost) GetText() string {
	return valueObject.text
}
