package usecase

import (
	"chelsnews/models/entity"
	"chelsnews/repository"
)

type NewsSubUseCase interface {
	CreateNewsCategory(news *entity.SubCategory) error
	GetNewsCategoryList() ([]entity.SubCategory, error)
	GetNewsCategoryByName(name string) (*entity.SubCategory, error)
	DeleteNewsCategory(name string) error
}

type NewsSubUseCaseImpl struct {
	NewsSubRepo repository.NewsSubRepository
}

func NewNewsSubUseCase(newsCtgRepo repository.NewsSubRepository) NewsSubUseCase {
	return &NewsSubUseCaseImpl{NewsSubRepo: newsCtgRepo}
}

func (uc *NewsSubUseCaseImpl) CreateNewsCategory(news *entity.SubCategory) error {
	return uc.NewsSubRepo.Create(news)
}

func (uc *NewsSubUseCaseImpl) GetNewsCategoryList() ([]entity.SubCategory, error) {
	return uc.NewsSubRepo.FindAll()
}
func (uc *NewsSubUseCaseImpl) GetNewsCategoryByName(name string) (*entity.SubCategory, error) {
	return uc.NewsSubRepo.FindByID(name)
}

func (uc *NewsSubUseCaseImpl) DeleteNewsCategory(name string) error {
	return uc.NewsSubRepo.Delete(name)
}
