package plant

import "time"

type PlantCategoryClimateResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PlantResponse struct {
	ID               int                   `json:"id"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	IsToxic          bool                  `json:"is_toxic"`
	HarvestDuration  int                   `json:"harvest_duration"`
	PlantCategory    PlantCategoryClimateResponse `json:"plant_category"`
	ClimateCondition PlantCategoryClimateResponse `json:"climate_condition"`
	PlantImages      []PlantImageResponse  `json:"plant_images"`
	CreatedAt        time.Time             `json:"created_at"`
}

type PlantImageResponse struct {
	ID        int    `json:"id"`
	PlantID   int    `json:"plant_id"`
	FileName  string `json:"file_name"`
	IsPrimary int    `json:"is_primary"`
}

func NewPlantResponse(plant Plant) PlantResponse {
	return PlantResponse{
		ID:               plant.ID,
		Name:             plant.Name,
		Description:      plant.Description,
		IsToxic:          plant.IsToxic,
		HarvestDuration:  plant.HarvestDuration,
		PlantCategory:    NewPlantCategoryResponse(plant.PlantCategory),
		ClimateCondition: NewClimateConditionResponse(plant.ClimateCondition),
		PlantImages:      NewPlantImageResponses(plant.PlantImages),
		CreatedAt:        plant.CreatedAt,
	}
}

func NewPlantCategoryResponse(category PlantCategory) PlantCategoryClimateResponse {
	return PlantCategoryClimateResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}


func NewClimateConditionResponse(condition ClimateCondition) PlantCategoryClimateResponse {
	return PlantCategoryClimateResponse{
		ID:   condition.ID,
		Name: condition.Name,
	}
}

func NewPlantImageResponses(images []PlantImage) []PlantImageResponse {
	var responses []PlantImageResponse
	for _, img := range images {
		responses = append(responses, NewPlantImageResponse(img))
	}
	return responses
}

func NewPlantImageResponse(image PlantImage) PlantImageResponse {
	return PlantImageResponse{
		ID:        image.ID,
		PlantID:   image.PlantID,
		FileName:  image.FileName,
		IsPrimary: image.IsPrimary,
	}
}