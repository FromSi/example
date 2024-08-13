package models

import (
	"gorm.io/gorm"
	"time"
)

type GormUserModel struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (GormUserModel) TableName() string {
	return "users"
}
