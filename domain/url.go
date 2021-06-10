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

type UrlRepository interface {
	FindById(id string) (*Url, error)
	Create(payload Url) (*Url, error)
}

type UrlUsecase interface {
	FindById(id string) (*Url, error)
	Create(payload CreateUrlDto) (*Url, error)
}
