package fertilizer

import "gorm.io/gorm"

type Repository interface {
	CreateFertilizer(*Fertilizer) (*Fertilizer, error)
	GetFertilizer(uint) ([]Fertilizer, error)
	GetFertilizerByID(uint) ([]Fertilizer, error)
	DeleteFertilizer(uint) error
	UpdateFertilizer(uint) error
}

type FertilizerRepository struct {
	db *gorm.DB
}

// UpdateFertilizer implements Repository.
func (r *FertilizerRepository) UpdateFertilizer(uint) error {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) *FertilizerRepository {
	return &FertilizerRepository{
		db: db,
	}
}

func (r *FertilizerRepository) CreateFertilizer(f *Fertilizer) (*Fertilizer, error) {
	err := r.db.Create(f).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Preload("User").Preload("Plant").First(f, f.Id).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FertilizerRepository) GetFertilizer(userID uint) ([]Fertilizer, error) {
	var f []Fertilizer
	err := r.db.Preload("Id").Preload("Name").Preload("Plant").Order("create_at").Where("user_id = ?", userID).Find(&f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FertilizerRepository) GetFertilizerByID(userID uint) ([]Fertilizer, error) {
	var f []Fertilizer
	err := r.db.Preload("Id").Preload("Name").Preload("Plant").Order("create_at").Where("user_id = ?", userID).Find(&f).Error
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *FertilizerRepository) DeleteFertilizer(userID uint) error {
	err := r.db.Where("user_id = ?", userID).Delete(&Fertilizer{}).Error
	if err != nil {
		return err
	}

	return nil
}
