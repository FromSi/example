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

func (post Post) SetText(text string) error {
	return post.Text.SetText(text)
}

type Id struct {
	id string
}

func (valueObject Id) GetId() string {
	return valueObject.id
}

func (valueObject *Id) setId(id string) error {
	err := validate.Var(id, "required,uuid")

	if err != nil {
		return err
	}

	valueObject.id = id

	return nil
}

func NewId(id string) (*Id, error) {
	valueObject := Id{}

	err := valueObject.setId(id)

	if err != nil {
		return nil, err
	}

	return &valueObject, nil
}

type Text struct {
	text string
}

func (valueObject Text) GetText() string {
	return valueObject.text
}

func (valueObject *Text) SetText(text string) error {
	err := validate.Var(text, "required,gte=3,lte=255")

	if err != nil {
		return err
	}

	valueObject.text = text

	return nil
}

func NewText(text string) (*Text, error) {
	valueObject := Text{}

	err := valueObject.SetText(text)

	if err != nil {
		return nil, err
	}

	return &valueObject, nil
}
