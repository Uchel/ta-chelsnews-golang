package entity

import (
	"time"

	"gorm.io/gorm"
)

type NewsCategory struct {
	Name      string         `json:"name" gorm:"type:varchar(199);primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
