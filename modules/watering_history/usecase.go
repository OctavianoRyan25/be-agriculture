package wateringhistory

type WateringHistoryUseCase interface {
	StoreWateringHistory(*WateringHistory) (*WateringHistory, error)
	GetAllWateringHistories(uint) ([]WateringHistory, error)
}

type wateringHistoryUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *wateringHistoryUseCase {
	return &wateringHistoryUseCase{
		repo: repo,
	}
}

func (uc *wateringHistoryUseCase) StoreWateringHistory(wh *WateringHistory) (*WateringHistory, error) {
	wh, err := uc.repo.StoreWateringHistory(wh)
	if err != nil {
		return nil, err
	}
	return wh, nil
}

func (uc *wateringHistoryUseCase) GetAllWateringHistories(userID uint) ([]WateringHistory, error) {
	wh, err := uc.repo.GetAllWateringHistories(userID)
	if err != nil {
		return nil, err
	}
	return wh, nil
}
