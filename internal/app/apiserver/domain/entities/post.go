package entities

import (
	"github.com/go-playground/validator/v10"
	"time"
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
	err := validate.Var(text, "required,gte=3,lte=255")

	if err != nil {
		return nil, err
	}

	return &Text{text: text}, nil
}

func (valueObject Text) GetText() string {
	return valueObject.text
}
