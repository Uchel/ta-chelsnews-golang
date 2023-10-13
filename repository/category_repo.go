package repository

import (
	"chelsnews/models/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type NewsCategoryRepository interface {
	Create(category *entity.NewsCategory) error
	FindAll() ([]entity.NewsCategory, error)
	Update(news *entity.NewsCategory, name string) error
	FindByID(name string) (*entity.NewsCategory, error)

	Delete(name string) error
}

type NewsCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewNewsCategoryRepository(db *gorm.DB) NewsCategoryRepository {
	return &NewsCategoryRepositoryImpl{DB: db}
}

func (r *NewsCategoryRepositoryImpl) Create(news *entity.NewsCategory) error {
	return r.DB.Create(news).Error
}

func (r *NewsCategoryRepositoryImpl) FindAll() ([]entity.NewsCategory, error) {
	var newsCtg []entity.NewsCategory
	if err := r.DB.Find(&newsCtg).Error; err != nil {
		return nil, err
	}
	return newsCtg, nil
}
func (r *NewsCategoryRepositoryImpl) FindByID(name string) (*entity.NewsCategory, error) {
	var newsCtg entity.NewsCategory
	if err := r.DB.First(&newsCtg, name).Error; err != nil {
		return nil, err
	}
	return &newsCtg, nil
}

func (r *NewsCategoryRepositoryImpl) Update(news *entity.NewsCategory, name string) error {

	news.UpdatedAt = time.Now()

	errSave := r.DB.Where("name = ?", name).Updates(&news).Error
	return errSave

}

func (r *NewsCategoryRepositoryImpl) Delete(name string) error {
	var newsCtg *entity.NewsCategory
	err := r.DB.Debug().First(&newsCtg, "name = ?", name).Error
	if err != nil {
		return fmt.Errorf("Gagal menghapus kategory")
	}
	return r.DB.Delete(newsCtg).Error
}
