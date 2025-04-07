package service

import (
	"guilt-type-service/internal/model"
	"guilt-type-service/internal/repository"
	"errors"
)

type guiltTypeService struct {
	repo repository.GuiltTypeRepository
}

func NewGuiltTypeService(r repository.GuiltTypeRepository) GuiltTypeService {
	return &guiltTypeService{repo: r}
}

func (s *guiltTypeService) GetAll() ([]model.GuiltType, error) {
	return s.repo.FindAll()
}

func (s *guiltTypeService) Create(gt model.GuiltType) (model.GuiltType, error) {
	if gt.Name == "Другое" && gt.OtherInfo == "" {
		return model.GuiltType{}, errors.New("При выборе 'Другое' нужно указать дополнительную информацию")
	}
	return s.repo.Create(gt)
}

func (s *guiltTypeService) Update(id uint, gt model.GuiltType) (model.GuiltType, error) {
	return s.repo.Update(id, gt)
}

func (s *guiltTypeService) Delete(id uint) error {
	return s.repo.Delete(id)
}