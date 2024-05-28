package plant

import "gorm.io/gorm"

type PlantRepository interface {
	FindAll() ([]Plant, error)
	FindByID(id int) (Plant, error)
	Create(plant Plant) (Plant, error)
	Update(plant Plant) (Plant, error)
	Delete(id int) error
	FindByIDWithRelations(id int) (Plant, error)
}

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) PlantRepository {
	return &plantRepository{db}
}

func (r *plantRepository) FindAll() ([]Plant, error) {
	var plants []Plant
	err := r.db.Preload("PlantCategory").Preload("PlantCharacteristic").Preload("WateringSchedule").
		Preload("PlantInstructions").Preload("PlantFAQs").Preload("PlantImages").Find(&plants).Error
	return plants, err
}

func (r *plantRepository) FindByID(id int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantCategory").Preload("PlantCharacteristic").Preload("WateringSchedule").
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

func (r *plantRepository) FindByIDWithRelations(id int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantCategory").
			Preload("PlantCharacteristic").
			Preload("WateringSchedule").
			Preload("PlantInstructions").
			Preload("PlantFAQs").
			Preload("PlantImages").
			Where("id = ?", id).
			First(&plant).Error
	if err != nil {
			return Plant{}, err
	}
	return plant, nil
}

func (r *plantRepository) Delete(id int) error {
	if err := r.db.Delete(&Plant{}, id).Error; err != nil {
			return err
	}
	return nil
}

