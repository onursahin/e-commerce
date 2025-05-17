package repositories

import (
	"github.com/onursahin/e-commerce/models"
	"gorm.io/gorm"
)

type PersonRepository interface {
	BaseRepository[models.Person]
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	return NewBaseRepository[models.Person](db)
}
