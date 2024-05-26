package plant

type PlantCategoryClimateInput struct {
	Name     string `json:"name" validate:"required"`
	ImageURL string `json:"image_url"`
}

type UpdatePlantInput struct {
	Name               string `json:"name" binding:"required"`
	Description        string `json:"description" binding:"required"`
	IsToxic            bool   `json:"is_toxic"`
	HarvestDuration    int    `json:"harvest_duration"`
	PlantCategoryID    int    `json:"plant_category_id"`
	ClimateConditionID int    `json:"climate_condition_id"`
	WateringScheduleID int    `json:"watering_schedule_id"`
	PlantInstructionID int    `json:"plant_instruction_id"`
	PlantFAQID         int    `json:"plant_faq_id"`
}

type CreatePlantInput struct {
	Name               string                        `json:"name" validate:"required"`
	Description        string                        `json:"description" validate:"required"`
	IsToxic            bool                          `json:"is_toxic"`
	HarvestDuration    int                           `json:"harvest_duration" validate:"required"`
	PlantCategoryID    int                           `json:"plant_category_id" validate:"required"`
	ClimateConditionID int                           `json:"climate_condition_id" validate:"required"`
	WateringSchedule   CreateWateringScheduleInput   `json:"watering_schedule"`
	PlantInstructions  []CreatePlantInstructionInput `json:"plant_instructions"`
	PlantFAQs          []CreatePlantFAQInput         `json:"plant_faqs"`
	PlantImages        []CreatePlantImageInput       `json:"plant_images" validate:"required,dive"`
}

type CreateWateringScheduleInput struct {
	WateringFrequency    int    `json:"watering_frequency" validate:"required"`
	Each                 string `json:"each" validate:"required"`
	WateringAmount       int    `json:"watering_amount" validate:"required"`
	Unit                 string `json:"unit" validate:"required"`
	WateringTime         string `json:"watering_time" validate:"required"`
	WeatherCondition     string `json:"weather_condition"`
	ConditionDescription string `json:"condition_description"`
}

type CreatePlantInstructionInput struct {
	StepNumber      int    `json:"step_number" validate:"required"`
	StepTitle       string `json:"step_title" validate:"required"`
	StepDescription string `json:"step_description" validate:"required"`
	StepImageURL    string `json:"step_image_url"`
	AdditionalTips  string `json:"additional_tips"`
}

type CreatePlantFAQInput struct {
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}

type CreatePlantImageInput struct {
	FileName  string `json:"file_name" validate:"required"`
	IsPrimary int    `json:"is_primary"`
}
