package wateringhistory

import "gorm.io/gorm"

type Repository interface {
	StoreWateringHistory(*WateringHistory) (*WateringHistory, error)
	GetAllWateringHistories(uint) ([]WateringHistory, error)
}

type wateringHistoryRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *wateringHistoryRepository {
	return &wateringHistoryRepository{
		db: db,
	}
}

func (r *wateringHistoryRepository) StoreWateringHistory(wh *WateringHistory) (*WateringHistory, error) {
	err := r.db.Create(wh).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Preload("User").Preload("Plant").First(wh, wh.ID).Error
	if err != nil {
		return nil, err
	}
	return wh, nil
}

func (r *wateringHistoryRepository) GetAllWateringHistories(userID uint) ([]WateringHistory, error) {
	var wh []WateringHistory
	err := r.db.Preload("User").Preload("Plant").Order("created_at desc").Where("user_id = ?", userID).Find(&wh).Error
	if err != nil {
		return nil, err
	}
	return wh, nil
}