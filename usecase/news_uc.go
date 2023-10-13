package usecase

import (
	"chelsnews/models/entity"
	"chelsnews/repository"
)

type NewsUseCase interface {
	CreateNews(news *entity.News) error
	GetNewsList() ([]entity.News, error)
	GetNewsByID(id uint) (*entity.News, error)
	UpdateNews(news *entity.News) error
	DeleteNews(id uint) error
}

type NewsUseCaseImpl struct {
	NewsRepo repository.NewsRepository
}

func NewNewsUseCase(newsRepo repository.NewsRepository) NewsUseCase {
	return &NewsUseCaseImpl{NewsRepo: newsRepo}
}

func (uc *NewsUseCaseImpl) CreateNews(news *entity.News) error {
	return uc.NewsRepo.Create(news)
}

func (uc *NewsUseCaseImpl) GetNewsList() ([]entity.News, error) {
	return uc.NewsRepo.FindAll()
}

func (uc *NewsUseCaseImpl) GetNewsByID(id uint) (*entity.News, error) {
	return uc.NewsRepo.FindByID(id)
}

func (uc *NewsUseCaseImpl) UpdateNews(news *entity.News) error {
	return uc.NewsRepo.Update(news)
}

func (uc *NewsUseCaseImpl) DeleteNews(id uint) error {
	return uc.NewsRepo.Delete(id)
}
