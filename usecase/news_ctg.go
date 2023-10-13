package usecase

import (
	"chelsnews/models/entity"
	"chelsnews/repository"
)

type NewsCategoryUseCase interface {
	CreateNewsCategory(news *entity.NewsCategory) error
	GetNewsCategoryList() ([]entity.NewsCategory, error)
	GetNewsCategoryByName(name string) (*entity.NewsCategory, error)
	UpdateNewsCategoryByName(news *entity.NewsCategory, name string) error
	DeleteNewsCategory(name string) error
}

type NewsCategoryUseCaseImpl struct {
	NewsCategoryRepo repository.NewsCategoryRepository
}

func NewNewsCategoryUseCase(newsCtgRepo repository.NewsCategoryRepository) NewsCategoryUseCase {
	return &NewsCategoryUseCaseImpl{NewsCategoryRepo: newsCtgRepo}
}

func (uc *NewsCategoryUseCaseImpl) CreateNewsCategory(news *entity.NewsCategory) error {
	return uc.NewsCategoryRepo.Create(news)
}

func (uc *NewsCategoryUseCaseImpl) GetNewsCategoryList() ([]entity.NewsCategory, error) {
	return uc.NewsCategoryRepo.FindAll()
}
func (uc *NewsCategoryUseCaseImpl) UpdateNewsCategoryByName(news *entity.NewsCategory, name string) error {
	return uc.NewsCategoryRepo.Update(news, name)
}
func (uc *NewsCategoryUseCaseImpl) GetNewsCategoryByName(name string) (*entity.NewsCategory, error) {
	return uc.NewsCategoryRepo.FindByID(name)
}

func (uc *NewsCategoryUseCaseImpl) DeleteNewsCategory(name string) error {
	return uc.NewsCategoryRepo.Delete(name)
}
