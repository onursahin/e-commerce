package repositories

import (
	"gorm.io/gorm"
)

type baseRepositoryImpl[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepositoryImpl[T]{db: db}
}

func (r *baseRepositoryImpl[T]) GetAll(filter func(db *gorm.DB) *gorm.DB, preloads ...string) ([]T, error) {
	var entities []T
	query := r.db

	if filter != nil {
		query = filter(query)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	err := query.Find(&entities).Error
	return entities, err
}

func (r *baseRepositoryImpl[T]) GetOne(filter func(db *gorm.DB) *gorm.DB, preloads ...string) (*T, error) {
	var entity T
	query := r.db

	if filter != nil {
		query = filter(query)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	err := query.First(&entity).Error

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *baseRepositoryImpl[T]) GetByID(id uint, preloads ...string) (*T, error) {
	var entity T
	query := r.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	err := query.First(&entity, id).Error

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (r *baseRepositoryImpl[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *baseRepositoryImpl[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *baseRepositoryImpl[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}
