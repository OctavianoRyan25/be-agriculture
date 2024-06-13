package fertilizer

type FertilizerUseCase interface {
	CreateFertilizer(*Fertilizer) (*Fertilizer, error)
	GetFertilizer(uint) ([]Fertilizer, error)
	DeleteFertilizer(uint) error
	UpdateFertilizer(uint) error
}

type fertilizerUseCase struct {
	repo Repository
}

// UpdateFertilizer implements FertilizerUseCase.
func (uc *fertilizerUseCase) UpdateFertilizer(uint) error {
	panic("unimplemented")
}

func NewUseCase(repo Repository) *fertilizerUseCase {
	return &fertilizerUseCase{
		repo: repo,
	}
}

func (uc *fertilizerUseCase) CreateFertilizer(wh *Fertilizer) (*Fertilizer, error) {
	wh, err := uc.repo.CreateFertilizer(wh)
	if err != nil {
		return nil, err
	}
	return wh, nil
}

func (uc *fertilizerUseCase) GetFertilizer(userID uint) ([]Fertilizer, error) {
	wh, err := uc.repo.GetFertilizer(userID)
	if err != nil {
		return nil, err
	}
	return wh, nil
}

func (u *fertilizerUseCase) DeleteFertilizer(userID uint) error {
	return u.repo.DeleteFertilizer(userID)
}
