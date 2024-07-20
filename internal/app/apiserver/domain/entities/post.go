package entities

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

const (
	TextMinLength = 3
	TextMaxLength = 250
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

type Post struct {
	ID        Id
	Text      Text
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func NewPost(id string, text string, createdAt *time.Time, updatedAt *time.Time, deletedAt *time.Time) (*Post, error) {
	idValueObject, err := NewId(id)

	if err != nil {
		return nil, err
	}

	textValueObject, err := NewText(text)

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
	textValueObject, err := NewText(text)

	if err != nil {
		return err
	}

	post.Text = *textValueObject

	return nil
}

type Id struct {
	id string
}

func NewId(id string) (*Id, error) {
	err := validate.Var(id, "required,uuid")

	if err != nil {
		return nil, err
	}

	return &Id{id: id}, nil
}

func (valueObject Id) GetId() string {
	return valueObject.id
}

type Text struct {
	text string
}

func NewText(text string) (*Text, error) {
	err := validate.Var(text, fmt.Sprintf("required,gte=%d,lte=%d", TextMinLength, TextMaxLength))

	if err != nil {
		return nil, err
	}

	return &Text{text: text}, nil
}

func (valueObject Text) GetText() string {
	return valueObject.text
}
