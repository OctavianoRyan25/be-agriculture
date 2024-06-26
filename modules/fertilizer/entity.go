package fertilizer

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
)

type Fertilizer struct {
	Id           int         `json:"id" gorm:"primaryKey"`
	Name         string      `json:"name"`
	PlantID      int         `json:"plantId"`
	Plant        plant.Plant `json:"plant"`
	Compostition string      `json:"compostition"`
	CreateAt     time.Time   `json:"createAt"`
}
