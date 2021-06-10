package repository

import (
	"github.com/ranggarifqi/qashir-api/domain"
	"gorm.io/gorm"
)

type urlRepository struct {
	conn *gorm.DB
}

func NewUrlRepository(conn *gorm.DB) domain.IUrlRepository {
	return &urlRepository{
		conn,
	}
}

func (r *urlRepository) FindById(id string) (*domain.Url, error) {
	var res domain.Url
	err := r.conn.First(&res, id).Error
	return &res, err
}

func (r *urlRepository) Create(payload domain.CreateUrlDto) (*domain.Url, error) {
	user := domain.Url{
		ShortenedUrl: payload.ShortenedUrl,
	}
	err := r.conn.Create(&user).Error
	return &user, err
}
