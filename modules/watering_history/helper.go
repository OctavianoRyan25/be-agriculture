package wateringhistory

import "github.com/OctavianoRyan25/be-agriculture/modules/plant"

func MapPlantToPlantResponse(plant *plant.Plant) *PlantResponse {
	return &PlantResponse{
		ID:               plant.ID,
		Name:             plant.Name,
		Description:      plant.Description,
		IsToxic:          plant.IsToxic,
		HarvestDuration:  plant.HarvestDuration,
		Sunlight:         plant.Sunlight,
		PlantingTime:     plant.PlantingTime,
		ClimateCondition: plant.ClimateCondition,
		CreatedAt:        plant.CreatedAt,
	}
}
