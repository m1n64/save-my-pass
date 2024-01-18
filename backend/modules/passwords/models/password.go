package models

import (
	models2 "backend/modules/categories/models"
	"backend/modules/users/models"
	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	UserID     uint             `gorm:"not null"`
	User       models.User      `gorm:"foreignKey:UserID"`
	CategoryID uint             `gorm:"nullable"`
	Category   models2.Category `gorm:"foreignKey:CategoryID"`
	Name       string           `gorm:"not null"`
	Login      string           `gorm:"not null"`
	Password   string           `gorm:"not null"`
	Additional string           `gorm:"not null"`
}
