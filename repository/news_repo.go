package repository

import (
	"chelsnews/models/entity"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type NewsRepository interface {
	Create(news *entity.News) error
	FindAll() ([]entity.News, error)
	FindByID(id uint) (*entity.News, error)
	Update(news *entity.News) error
	Delete(id uint) error
}

type NewsRepositoryImpl struct {
	DB *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return &NewsRepositoryImpl{DB: db}
}

func (r *NewsRepositoryImpl) Create(news *entity.News) error {
	return r.DB.Create(news).Error
}

func (r *NewsRepositoryImpl) FindAll() ([]entity.News, error) {
	var news []entity.News
	if err := r.DB.Preload("NewsCategory").Find(&news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (r *NewsRepositoryImpl) FindByID(id uint) (*entity.News, error) {
	var news entity.News
	if err := r.DB.Preload("NewsCategory").First(&news, id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}
func (r *NewsRepositoryImpl) Update(news *entity.News) error {
	fmt.Println(news.ID)
	var newsO *entity.News
	err := r.DB.Debug().First(&newsO, "id = ?", news.ID).Error
	if err != nil {
		return err
	}

	if news.Image != "" {
		newsO.Image = news.Image
	}
	if news.Title != "" {
		newsO.Title = news.Title
	}
	if news.Content != "" {
		newsO.Content = news.Content
	}
	if news.NewsCategoryName != "" {
		newsO.NewsCategoryName = news.NewsCategoryName
	}
	if news.SubCategoryName != "" {
		newsO.SubCategoryName = news.SubCategoryName
	}
	newsO.UpdatedAt = time.Now()

	errSave := r.DB.Save(&newsO).Error
	return errSave

}

func (r *NewsRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&entity.News{}, id).Error
}
