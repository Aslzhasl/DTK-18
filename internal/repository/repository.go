package repository

import (
	"guilt-type-service/internal/model"
)

type GuiltTypeRepository interface {
	FindAll() ([]model.GuiltType, error)
	Create(guiltType model.GuiltType) (model.GuiltType, error)
	Update(id uint, updated model.GuiltType) (model.GuiltType, error)
	Delete(id uint) error
	FindByID(id uint) (model.GuiltType, error)
	BulkInsert([]model.GuiltType) error
}
