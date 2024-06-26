package fertilizer

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
)

type FertilizerResponse struct {
	Id           int                 `json:"id" gorm:"primaryKey"`
	Name         string              `json:"name"`
	Plant        plant.PlantResponse `json:"plant"`
	Compostition string              `json:"compostition"`
	CreateAt     time.Time           `json:"createAt"`
}

func NewFertilizerResponse(fertilizer Fertilizer) *FertilizerResponse {
	return &FertilizerResponse{
		Id:           fertilizer.Id,
		Name:         fertilizer.Name,
		Plant:        plant.NewPlantResponse(fertilizer.Plant),
		Compostition: fertilizer.Compostition,
		CreateAt:     fertilizer.CreateAt,
	}
}
