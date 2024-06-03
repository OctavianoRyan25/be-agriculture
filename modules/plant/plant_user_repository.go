package plant

import "gorm.io/gorm"

type UserPlantRepository interface {
	AddUserPlant(userPlant UserPlant) (UserPlant, error)
	GetUserPlantsByUserID(userID int, limit int, offset int) ([]UserPlant, error)
	DeleteUserPlantByID(userPlantID int) error
	GetUserPlantByID(userPlantID int) (UserPlant, error)
	CountByUserID(userID int, count *int64) error
}

type userPlantRepository struct {
	db *gorm.DB
}

func NewUserPlantRepository(db *gorm.DB) UserPlantRepository {
	return &userPlantRepository{db}
}

func (r *userPlantRepository) AddUserPlant(userPlant UserPlant) (UserPlant, error) {
	err := r.db.Create(&userPlant).Error
	if err != nil {
		return userPlant, err
	}

	err = r.db.Preload("Plant").
		Preload("Plant.PlantCategory").
		Preload("Plant.PlantCharacteristic").
		Preload("Plant.WateringSchedule").
		Preload("Plant.PlantInstructions").
		Preload("Plant.PlantFAQs").
		Preload("Plant.PlantImages").
		First(&userPlant, userPlant.ID).Error

	return userPlant, err
}

func (r *userPlantRepository) GetUserPlantsByUserID(userID int, limit int, offset int) ([]UserPlant, error) {
	var userPlants []UserPlant
	query := r.db.Preload("Plant").
		Preload("Plant.PlantCategory").
		Preload("Plant.PlantCharacteristic").
		Preload("Plant.WateringSchedule").
		Preload("Plant.PlantInstructions").
		Preload("Plant.PlantFAQs").
		Preload("Plant.PlantImages").
		Where("user_id = ?", userID)

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset >= 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&userPlants).Error
	return userPlants, err
}



func (r *userPlantRepository) DeleteUserPlantByID(userPlantID int) error {
	return r.db.Where("id = ?", userPlantID).Delete(&UserPlant{}).Error
}


func (r *userPlantRepository) GetUserPlantByID(userPlantID int) (UserPlant, error) {
	var userPlant UserPlant
	err := r.db.Preload("Plant").
			Preload("Plant.PlantCategory").
			Preload("Plant.PlantCharacteristic").
			Preload("Plant.WateringSchedule").
			Preload("Plant.PlantInstructions").
			Preload("Plant.PlantFAQs").
			Preload("Plant.PlantImages").
			Where("id = ?", userPlantID).
			First(&userPlant).Error
	if err != nil {
			return UserPlant{}, err
	}
	return userPlant, nil
}

func (r *userPlantRepository) CountByUserID(userID int, count *int64) error {
	return r.db.Model(&UserPlant{}).Where("user_id = ?", userID).Count(count).Error
}
