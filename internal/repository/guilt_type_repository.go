package repository

import (
	"gorm.io/gorm"
	"guilt-type-service/internal/model"
	"errors"
)

type guiltTypeRepository struct {
	db *gorm.DB
}

func NewGuiltTypeRepository(db *gorm.DB) GuiltTypeRepository {
	return &guiltTypeRepository{db: db}
}

func (r *guiltTypeRepository) FindAll() ([]model.GuiltType, error) {
	var types []model.GuiltType
	result := r.db.Find(&types)
	return types, result.Error
}

func (r *guiltTypeRepository) Create(gt model.GuiltType) (model.GuiltType, error) {
	result := r.db.Create(&gt)
	return gt, result.Error
}

func (r *guiltTypeRepository) Update(id uint, updated model.GuiltType) (model.GuiltType, error) {
	var existing model.GuiltType
	if err := r.db.First(&existing, id).Error; err != nil {
		return model.GuiltType{}, err
	}
	existing.Name = updated.Name
	existing.OtherInfo = updated.OtherInfo
	if err := r.db.Save(&existing).Error; err != nil {
		return model.GuiltType{}, err
	}
	return existing, nil
}

func (r *guiltTypeRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.GuiltType{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *guiltTypeRepository) FindByID(id uint) (model.GuiltType, error) {
	var gt model.GuiltType
	if err := r.db.First(&gt, id).Error; err != nil {
		return model.GuiltType{}, errors.New("not found")
	}
	return gt, nil
}

func (r *guiltTypeRepository) BulkInsert(gts []model.GuiltType) error {
	return r.db.Create(&gts).Error
}