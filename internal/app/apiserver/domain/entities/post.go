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

type IdValueObject struct {
	id string
}

func (valueObject IdValueObject) GetId() string {
	return valueObject.id
}

func (valueObject *IdValueObject) setId(id string) error {
	err := validate.Var(id, "required,uuid")

	if err != nil {
		return err
	}

	valueObject.id = id

	return nil
}

func NewIdValueObject(id string) (*IdValueObject, error) {
	valueObject := IdValueObject{}

	err := valueObject.setId(id)

	if err != nil {
		return nil, err
	}

	return &valueObject, nil
}

type TextValueObject struct {
	text string
}

func (valueObject TextValueObject) GetText() string {
	return valueObject.text
}

func (valueObject *TextValueObject) SetText(text string) error {
	err := validate.Var(text, "required,gte=3,lte=255")

	if err != nil {
		return err
	}

	valueObject.text = text

	return nil
}

func NewTextValueObject(text string) (*TextValueObject, error) {
	valueObject := TextValueObject{}

	err := valueObject.SetText(text)

	if err != nil {
		return nil, err
	}

	return &valueObject, nil
}

type Post struct {
	ID        IdValueObject
	Text      TextValueObject
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (post Post) SetText(text string) error {
	return post.Text.SetText(text)
}
