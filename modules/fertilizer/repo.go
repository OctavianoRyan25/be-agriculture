package fertilizer

import "gorm.io/gorm"

type FertilizerRepository interface {
	GetFertilizer() ([]Fertilizer, error)
	GetFertilizerByID(id int) (Fertilizer, error)
	CreateFertilizer(*Fertilizer) (*Fertilizer, error)
	UpdateFertilizer(id int, fertilizer *Fertilizer) (*Fertilizer, error)
	DeleteFertilizer(int) error
}

type fertilizerRepository struct {
	db *gorm.DB
}

func NewFertilizerRepository(db *gorm.DB) FertilizerRepository {
	return &fertilizerRepository{
		db: db,
	}
}

func (r *fertilizerRepository) GetFertilizer() ([]Fertilizer, error) {
	var fertilizers []Fertilizer
	err := r.db.Preload("Plant").Find(&fertilizers).Error
	return fertilizers, err
}

func (r *fertilizerRepository) GetFertilizerByID(id int) (Fertilizer, error) {
	var fertilizer Fertilizer
	err := r.db.Preload("Plant").First(&fertilizer, id).Error
	return fertilizer, err
}

func (r *fertilizerRepository) CreateFertilizer(fertilizer *Fertilizer) (*Fertilizer, error) {
	err := r.db.Create(&fertilizer).Error
	return fertilizer, err
}

func (r *fertilizerRepository) UpdateFertilizer(id int, fertilizer *Fertilizer) (*Fertilizer, error) {
	err := r.db.Where("id = ?", id).Updates(fertilizer).Error
	return fertilizer, err
}

func (r *fertilizerRepository) DeleteFertilizer(id int) error {
	err := r.db.Where("id = ?", id).Delete(&Fertilizer{}).Error
	return err
}
