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

func (uc *fertilizerUseCase) CreateFertilizer(f *Fertilizer) (*Fertilizer, error) {
	f, err := uc.repo.CreateFertilizer(f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (uc *fertilizerUseCase) GetFertilizer(userID uint) ([]Fertilizer, error) {
	f, err := uc.repo.GetFertilizer(userID)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (u *fertilizerUseCase) DeleteFertilizer(userID uint) error {
	return u.repo.DeleteFertilizer(userID)
}
