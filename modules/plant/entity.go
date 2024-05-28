package plant

import (
	"time"
)

type Plant struct {
	ID                	 int                `json:"id" gorm:"primaryKey"`
	Name              	 string             `json:"name"`
	Description       	 string             `json:"description"`
	IsToxic           	 bool               `json:"is_toxic"`
	HarvestDuration   	 int                `json:"harvest_duration"`
	Sunlight          	 string             `json:"sunlight"`
	PlantingTime      	 string             `json:"planting_time"`
	PlantCategoryID   	 int                `json:"plant_category_id"`
	PlantCategory     	 PlantCategory      `json:"plant_category"`
	ClimateCondition		 string             `json:"climate_condition"`
	PlantCharateristicID int                `json:"plant_charateristic_id"`
	PlantCharateristic   PlantCharateristic `json:"plant_charateristic"`
	WateringSchedule  	 PlantReminder  		`json:"watering_schedule" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
  PlantInstructions 	 []PlantInstruction `json:"plant_instructions" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
  PlantFAQs         	 []PlantFAQ         `json:"plant_faqs" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
  PlantImages       	 []PlantImage       `json:"plant_images" gorm:"foreignKey:PlantID;constraint:OnDelete:CASCADE"`
	CreatedAt         	 time.Time          `json:"created_at"`
	UpdatedAt         	 time.Time          `json:"updated_at"`
}

type PlantCharateristic struct {
	ID         					 int    						`json:"id" gorm:"primaryKey"`
	PlantID    					 int    						`json:"plant_id"`
	Height     					 int    						`json:"height"`
	HeightUnit 					 string 						`json:"height_unit"`
	Wide       					 int    						`json:"wide"`
	WideUnit   					 string 						`json:"wide_unit"`
	LeafColor  					 string 						`json:"leaf_color"`
}

type PlantReminder struct {
	ID                   int    					 	`json:"id" gorm:"primaryKey"`
	PlantID              int    					 	`json:"plant_id"`
	WateringFrequency    int    					 	`json:"watering_frequency"`
	Each                 string 					 	`json:"each"`
	WateringAmount       int    					 	`json:"watering_amount"`
	Unit                 string 					 	`json:"unit"`
	WateringTime         string 					 	`json:"watering_time"`
	WeatherCondition     string 					 	`json:"weather_condition"`
	ConditionDescription string 					 	`json:"condition_description"`
	CreatedAt         	 time.Time         	`json:"created_at"`
	UpdatedAt         	 time.Time         	`json:"updated_at"`

}

type PlantInstruction struct {
	ID             	 		 int    						`json:"id" gorm:"primaryKey"`
	PlantID         		 int    						`json:"plant_id"`
	StepNumber      		 int    						`json:"step_number"`
	StepTitle       		 string 						`json:"step_title"`
	StepDescription 		 string 						`json:"step_description"`
	StepImageURL    		 string 						`json:"step_image_url"`
	AdditionalTips  		 string 						`json:"additional_tips"`
	CreatedAt         	 time.Time          `json:"created_at"`
	UpdatedAt         	 time.Time          `json:"updated_at"`

}

type PlantFAQ struct {
	ID        					 int    						`json:"id" gorm:"primaryKey"`
	PlantID   					 int    						`json:"plant_id"`
	Question  					 string 						`json:"question"`
	Answer    					 string 						`json:"answer"`
	CreatedAt         	 time.Time          `json:"created_at"`
	UpdatedAt         	 time.Time          `json:"updated_at"`
}

type PlantImage struct {
	ID        					 int    						`json:"id" gorm:"primaryKey"`
	PlantID   					 int    						`json:"plant_id"`
	FileName  					 string 						`json:"file_name"`
	IsPrimary 					 int    						`json:"is_primary"`
	CreatedAt         	 time.Time          `json:"created_at"`
	UpdatedAt         	 time.Time          `json:"updated_at"`

}

type PlantCategory struct {
	ID   								 int    						`json:"id" gorm:"primaryKey"`
	Name 								 string 						`json:"name"`
	ImageURL 						 string 						`json:"image_url"`
	CreatedAt         	 time.Time          `json:"created_at"`
	UpdatedAt         	 time.Time          `json:"updated_at"`
}