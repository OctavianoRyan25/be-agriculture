package plant

type PlantCategoryClimateInput struct {
	Name string `json:"name" validate:"required"`
}

type CreatePlantInput struct {
	Name               string `json:"name" validate:"required"`
	Description        string `json:"description" validate:"required"`
	IsToxic            bool   `json:"is_toxic"`
	HarvestDuration    int    `json:"harvest_duration"`
	PlantCategoryID    int    `json:"plant_category_id" validate:"required"`
	ClimateConditionID int    `json:"climate_condition_id" validate:"required"`
}
