package models

import (
	"gorm.io/gorm"
	"time"
)

type GormPostModel struct {
	ID        string `gorm:"primaryKey"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (GormPostModel) TableName() string {
	return "post"
}
