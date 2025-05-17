package repositories

import "gorm.io/gorm"

type BaseRepository[T any] interface {
	GetAll(filter func(db *gorm.DB) *gorm.DB, preloads ...string) ([]T, error)
	GetOne(filter func(db *gorm.DB) *gorm.DB, preloads ...string) (*T, error)
	GetByID(id uint, preloads ...string) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}
