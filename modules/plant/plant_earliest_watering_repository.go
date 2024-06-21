// repository.go
package plant

import "gorm.io/gorm"

type PlantEarliestWateringRepository interface {
	GetWateringSchedules(plantID string) ([]PlantEarliestWatering, error)
}

type plantEarliestWateringRepository struct {
	db *gorm.DB
}

// GetWateringSchedules implements PlantEarliestWateringRepository.
func (r *plantEarliestWateringRepository) GetWateringSchedules(plantID string) ([]PlantEarliestWatering, error) {
	panic("unimplemented")
}

func NewPlantEarliestWateringRepository(db *gorm.DB) PlantEarliestWateringRepository {
	return &plantEarliestWateringRepository{db}
}

func (r *plantEarliestWateringRepository) FindAll() ([]PlantEarliestWatering, error) {
	var categories []PlantEarliestWatering
	err := r.db.Find(&categories).Error
	return categories, err
}
