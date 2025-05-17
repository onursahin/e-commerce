package repositories

import (
	"github.com/onursahin/e-commerce/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return NewBaseRepository[models.User](db)
}
