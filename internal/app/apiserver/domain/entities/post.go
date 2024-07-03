package entities

import (
	"time"
)

type Post struct {
	ID        string
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
