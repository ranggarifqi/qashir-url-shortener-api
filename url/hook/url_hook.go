package hook

import (
	"github.com/ranggarifqi/qashir-api/domain"
	"github.com/ranggarifqi/qashir-api/helper"
	"gorm.io/gorm"
)

func (u *domain.Url) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = helper.GenerateID(6)
	return nil
}
