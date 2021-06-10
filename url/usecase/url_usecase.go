package usecase

import (
	"github.com/ranggarifqi/qashir-api/domain"
)

type urlUsecase struct {
	urlRepository domain.IUrlRepository
}

func NewUrlUsecase(urlRepository domain.IUrlRepository) domain.IUrlUsecase {
	return &urlUsecase{
		urlRepository,
	}
}

func (u *urlUsecase) FindById(id string) (*domain.Url, error) {
	result, err := u.urlRepository.FindById(id)
	return result, err
}

func (u *urlUsecase) Create(payload domain.CreateUrlDto) (*domain.Url, error) {
	result, err := u.urlRepository.Create(payload)
	return result, err
}
