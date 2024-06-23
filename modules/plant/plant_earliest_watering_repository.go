// repository.go
package plant

import "gorm.io/gorm"

type PlantEarliestWateringRepository interface {
	GetEarliestWatering(plantID string) ([]PlantEarliestWatering, error)
}

type plantEarliestWateringRepository struct {
	db *gorm.DB
}

func NewPlantEarliestWateringRepository(db *gorm.DB) PlantEarliestWateringRepository {
	return &plantEarliestWateringRepository{db}
}

// GetEarliestWatering implements PlantEarliestWateringRepository.
func (r *plantEarliestWateringRepository) GetEarliestWatering(plantID string) ([]PlantEarliestWatering, error) {
	var categories []PlantEarliestWatering
	err := r.db.First(&categories).Error
	return categories, err
}