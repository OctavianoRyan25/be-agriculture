package wateringhistory

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/user"
)

type WateringHistoryResponse struct {
	Id        int               `json:"id"`
	Plant     PlantResponse     `json:"plant"`
	User      user.UserResponse `json:"user"`
	CreatedAt time.Time         `json:"created_at"`
}

type PlantResponse struct {
	ID               int       `json:"id" gorm:"primaryKey"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	IsToxic          bool      `json:"is_toxic"`
	HarvestDuration  int       `json:"harvest_duration"`
	Sunlight         string    `json:"sunlight"`
	PlantingTime     string    `json:"planting_time"`
	ClimateCondition string    `json:"climate_condition"`
	CreatedAt        time.Time `json:"created_at"`
}
