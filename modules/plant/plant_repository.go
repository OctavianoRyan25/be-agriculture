package plant

import "gorm.io/gorm"

type PlantRepository interface {
	FindAll() ([]Plant, error)
	FindByID(id int) (Plant, error)
	Create(plant CreatePlantInput) (Plant, error)
	Update(id int, plant CreatePlantInput) (Plant, error)
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
	err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("PlantImages").Find(&plants).Error
	return plants, err
}

func (r *plantRepository) FindByID(id int) (Plant, error) {
	var plant Plant
	err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("PlantImages").First(&plant, id).Error
	return plant, err
}

func (r *plantRepository) Create(plant CreatePlantInput) (Plant, error) {
	newPlant := Plant{
			Name:               plant.Name,
			Description:        plant.Description,
			IsToxic:            plant.IsToxic,
			HarvestDuration:    plant.HarvestDuration,
			PlantCategoryID:    plant.PlantCategoryID,
			ClimateConditionID: plant.ClimateConditionID,
	}

	err := r.db.Create(&newPlant).Error
	if err != nil {
			return Plant{}, err
	}

	if err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("PlantImages").First(&newPlant, newPlant.ID).Error; err != nil {
			return Plant{}, err
	}

	return newPlant, nil
}

func (r *plantRepository) Update(id int, plant CreatePlantInput) (Plant, error) {
	var existingPlant Plant
	if err := r.db.First(&existingPlant, id).Error; err != nil {
			return Plant{}, err
	}

	existingPlant.Name = plant.Name
	existingPlant.Description = plant.Description
	existingPlant.IsToxic = plant.IsToxic
	existingPlant.HarvestDuration = plant.HarvestDuration
	existingPlant.PlantCategoryID = plant.PlantCategoryID
	existingPlant.ClimateConditionID = plant.ClimateConditionID

	if err := r.db.Save(&existingPlant).Error; err != nil {
			return Plant{}, err
	}

	if err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("PlantImages").First(&existingPlant, existingPlant.ID).Error; err != nil {
			return Plant{}, err
	}

	return existingPlant, nil
}

func (r *plantRepository) Delete(id int) (Plant, error) {
	var plant Plant
	if err := r.db.Preload("PlantCategory").Preload("ClimateCondition").Preload("PlantImages").First(&plant, id).Error; err != nil {
			return Plant{}, err
	}

	if err := r.db.Delete(&plant).Error; err != nil {
			return Plant{}, err
	}

	return plant, nil
}