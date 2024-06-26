package fertilizer

import "time"

type FertilizerService interface {
	CreateFertilizer(*Fertilizer) (*Fertilizer, error)
	GetFertilizer() ([]Fertilizer, error)
	GetFertilizerByID(id int) (*Fertilizer, error)
	DeleteFertilizer(id int) error
	UpdateFertilizer(id int, fertilizer *Fertilizer) error
}

type fertilizerService struct {
	repository FertilizerRepository
}

func NewFertilizerService(repository FertilizerRepository) FertilizerService {
	return &fertilizerService{repository}
}

func (s *fertilizerService) GetFertilizer() ([]Fertilizer, error) {
	fertilizer, err := s.repository.GetFertilizer()
	if err != nil {
		return nil, err
	}
	return fertilizer, nil
}

func (s *fertilizerService) GetFertilizerByID(id int) (*Fertilizer, error) {
	fertilizer, err := s.repository.GetFertilizerByID(id)
	if err != nil {
		return nil, err
	}
	return &fertilizer, nil
}

func (s *fertilizerService) CreateFertilizer(fertilizer *Fertilizer) (*Fertilizer, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	fertilizer.CreateAt = time.Now().In(location)
	fertilizer.UpdatedAt = time.Now().In(location)
	fertilizer, err := s.repository.CreateFertilizer(fertilizer)
	if err != nil {
		return nil, err
	}

	return fertilizer, nil
}

func (s *fertilizerService) UpdateFertilizer(id int, fertilizer *Fertilizer) error {
	location, _ := time.LoadLocation("Asia/Jakarta")
	fertilizer.UpdatedAt = time.Now().In(location)
	_, err := s.repository.UpdateFertilizer(id, fertilizer)
	if err != nil {
		return err
	}
	return nil
}

func (s *fertilizerService) DeleteFertilizer(id int) error {
	return s.repository.DeleteFertilizer(id)
}
