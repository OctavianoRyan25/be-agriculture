package plant

import "gorm.io/gorm"

type PlantRepository interface {
	FindAll() ([]Plant, error)
	FindByID(id int) (Plant, error)
	Create(plant Plant) (Plant, error)
	Update(plant Plant) (Plant, error)
	Delete(id int) (Plant, error)
}

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) PlantRepository {
	return &plantRepository{db}
}

func (r *plantRepository) FindAll() ([]Plant, error) {
	var plants []Plant
	err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("WateringSchedule").
		Preload("PlantInstructions").Preload("PlantFAQs").Preload("PlantImages").Find(&plants).Error
	return plants, err
}

func (r *plantRepository) FindByID(id int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("WateringSchedule").
		Preload("PlantInstructions").Preload("PlantFAQs").Preload("PlantImages").First(&plant, id).Error
	return plant, err
}


func (r *plantRepository) Create(plant Plant) (Plant, error) {
	err := r.db.Create(&plant).Error
	return plant, err
}

func (r *plantRepository) Update(plant Plant) (Plant, error) {
	err := r.db.Save(&plant).Error
	return plant, err
}

func (r *plantRepository) Delete(id int) (Plant, error) {
	var plant Plant
	err := r.db.First(&plant, id).Error
	if err != nil {
		return plant, err
	}
	err = r.db.Delete(&plant).Error
	return plant, err
}
