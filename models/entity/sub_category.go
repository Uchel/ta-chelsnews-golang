package entity

import (
	"time"
)

type SubCategory struct {
	Name      string    `json:"name" gorm:"type:varchar(199);primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
