package entity

import (
	// "chelsnews/models/response"

	"time"

	"gorm.io/gorm"
)

type News struct {
	ID               uint           `json:"id" form:"id" gorm:"primaryKey"`
	Image            string         `json:"image" gorm:"not null"`
	Title            string         `json:"title" form:"title" gorm:"not null"`
	Content          string         `json:"content" form:"content" gorm:"not null"`
	NewsCategoryName string         `json:"news_category_name" form:"news_category_name" gorm:"type:varchar(199);not null"`
	NewsCategory     NewsCategory   `json:"news_category"`
	SubCategoryName  string         `json:"sub_category_name" form:"sub_category_name" gorm:"type:varchar(199)"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"-" gorm:"index"`
}

// AdminId      string       `json:"-"`
// Admin        Admin
