package plant

import (
	"time"
)

type Plant struct {
	ID                	int                `json:"id" gorm:"primaryKey"`
	Name              	string             `json:"name"`
	Description       	string             `json:"description"`
	IsToxic           	bool               `json:"is_toxic"`
	HarvestDuration   	int                `json:"harvest_duration"`
	PlantCategoryID   	int                `json:"plant_category_id"`
	PlantCategory     	PlantCategory      `json:"plant_category"`
	ClimateConditionID  int              	 `json:"climate_condition_id"`
	ClimateCondition  	ClimateCondition   `json:"climate_condition"`
	WateringSchedule  	PlantReminder  		 `json:"watering_schedule" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
  PlantInstructions 	[]PlantInstruction `json:"plant_instructions" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
  PlantFAQs         	[]PlantFAQ         `json:"plant_faqs" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
  PlantImages       	[]PlantImage       `json:"plant_images" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	CreatedAt         	time.Time          `json:"created_at"`
	UpdatedAt         	time.Time          `json:"updated_at"`
}

type PlantReminder struct {
	ID                   int    `json:"id" gorm:"primaryKey"`
	PlantID              int    `json:"plant_id"`
	WateringFrequency    int    `json:"watering_frequency"`
	Each                 string `json:"each"`
	WateringAmount       int    `json:"watering_amount"`
	Unit                 string `json:"unit"`
	WateringTime         string `json:"watering_time"`
	WeatherCondition     string `json:"weather_condition"`
	ConditionDescription string `json:"condition_description"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`

}

type PlantInstruction struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	PlantID         int    `json:"plant_id"`
	StepNumber      int    `json:"step_number"`
	StepTitle       string `json:"step_title"`
	StepDescription string `json:"step_description"`
	StepImageURL    string `json:"step_image_url"`
	AdditionalTips  string `json:"additional_tips"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`

}

type PlantFAQ struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	PlantID   int    `json:"plant_id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

type PlantImage struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	PlantID   int    `json:"plant_id"`
	FileName  string `json:"file_name"`
	IsPrimary int    `json:"is_primary"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`

}

type PlantCategory struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

type ClimateCondition struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}
