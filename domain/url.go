package domain

import (
	"time"

	"github.com/ranggarifqi/qashir-api/helper"
	"gorm.io/gorm"
)

type Url struct {
	ID           string    `json:"id"`
	ShortenedUrl string    `json:"url" validate:"required" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
}

func (u *Url) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = helper.GenerateID(6)
	return nil
}

type CreateUrlDto struct {
	ShortenedUrl string `json:"url" validate:"required"`
}

type IUrlRepository interface {
	FindById(id string) (*Url, error)
	Create(payload CreateUrlDto) (*Url, error)
}

type IUrlUsecase interface {
	FindById(id string) (*Url, error)
	Create(payload CreateUrlDto) (*Url, error)
}
