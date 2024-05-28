package plant

import (
	"gorm.io/gorm"
)

type ClimateConditionRepository interface {
	FindAll() ([]ClimateCondition, error)
	FindByID(id int) (ClimateCondition, error)
	Create(climateCondition ClimateCondition) (ClimateCondition, error)
	Update(climateCondition ClimateCondition) (ClimateCondition, error)
	Delete(climateCondition ClimateCondition) error
}

type climateConditionRepository struct {
	db *gorm.DB
}

func NewClimateConditionRepository(db *gorm.DB) ClimateConditionRepository {
	return &climateConditionRepository{db}
}

func (r *climateConditionRepository) FindAll() ([]ClimateCondition, error) {
	var climateConditions []ClimateCondition
	err := r.db.Find(&climateConditions).Error
	return climateConditions, err
}

func (r *climateConditionRepository) FindByID(id int) (ClimateCondition, error) {
	var climateCondition ClimateCondition
	err := r.db.First(&climateCondition, id).Error
	return climateCondition, err
}

func (r *climateConditionRepository) Create(climateCondition ClimateCondition) (ClimateCondition, error) {
	err := r.db.Create(&climateCondition).Error
	return climateCondition, err
}

func (r *climateConditionRepository) Update(climateCondition ClimateCondition) (ClimateCondition, error) {
	err := r.db.Save(&climateCondition).Error
	return climateCondition, err
}

func (r *climateConditionRepository) Delete(climateCondition ClimateCondition) error {
	err := r.db.Delete(&climateCondition).Error
	return err
}
