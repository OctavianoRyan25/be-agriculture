package wateringhistory

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
)

type WateringHistory struct {
	ID        int `gorm:"primaryKey"`
	PlantID   int
	Plant     plant.Plant `gorm:"foreignKey:PlantID;references:ID"`
	UserID    int
	User      user.User `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
