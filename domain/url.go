package domain

import "time"

type Url struct {
	ID           string    `json:"id"`
	ShortenedUrl string    `json:"url" validate:"required" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
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
