package repository

import (
	"chelsnews/models/entity"
	"fmt"

	"gorm.io/gorm"
)

type NewsSubRepository interface {
	Create(category *entity.SubCategory) error
	FindAll() ([]entity.SubCategory, error)
	FindByID(name string) (*entity.SubCategory, error)
	Delete(name string) error
}

type NewsSubRepositoryImpl struct {
	DB *gorm.DB
}

func NewNewsSubRepository(db *gorm.DB) NewsSubRepository {
	return &NewsSubRepositoryImpl{DB: db}
}

func (r *NewsSubRepositoryImpl) Create(newsSub *entity.SubCategory) error {
	return r.DB.Create(newsSub).Error
}

func (r *NewsSubRepositoryImpl) FindAll() ([]entity.SubCategory, error) {
	var newsSub []entity.SubCategory
	if err := r.DB.Find(&newsSub).Error; err != nil {
		return nil, err
	}
	return newsSub, nil
}
func (r *NewsSubRepositoryImpl) FindByID(name string) (*entity.SubCategory, error) {
	var newsSub entity.SubCategory
	if err := r.DB.First(&newsSub, name).Error; err != nil {
		return nil, err
	}
	return &newsSub, nil
}

func (r *NewsSubRepositoryImpl) Delete(name string) error {
	var newsSub *entity.SubCategory
	err := r.DB.Debug().First(&newsSub, "name = ?", name).Error
	if err != nil {
		return fmt.Errorf("Gagal menghapus kategory")
	}
	return r.DB.Delete(newsSub).Error
}
