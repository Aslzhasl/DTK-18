package service

import "guilt-type-service/internal/model"

type GuiltTypeService interface {
	GetAll() ([]model.GuiltType, error)
	Create(gt model.GuiltType) (model.GuiltType, error)
	Update(id uint, gt model.GuiltType) (model.GuiltType, error)
	Delete(id uint) error
}
