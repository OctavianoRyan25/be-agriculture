package plant

import "time"

type Plant struct {
	ID                   int                `gorm:"column:id;primaryKey"`
	Name                 string             `gorm:"column:name"`
	Description          string             `gorm:"column:description;type:TEXT"`
	IsToxic              bool               `gorm:"column:is_toxic"`
	HarvestDuration      int                `gorm:"column:harvest_duration"`
	// WateringSchedule     string             `gorm:"column:watering_schedule"`
	PlantCategoryID      int                `gorm:"column:plant_category_id"`
	ClimateConditionID   int                `gorm:"column:climate_condition_id"`
	CreatedAt            time.Time          `gorm:"column:created_at"`
	UpdatedAt            time.Time          `gorm:"column:updated_at"`

	PlantCategory        PlantCategory      `gorm:"foreignkey:PlantCategoryID;references:ID"`
	ClimateCondition     ClimateCondition   `gorm:"foreignkey:ClimateConditionID;references:ID"`
	PlantImages          []PlantImage       `gorm:"foreignkey:PlantID;references:ID"`
}

type PlantCategory struct {
	ID   								 int    						`gorm:"column:id;primaryKey"`
	Name 								 string 						`gorm:"column:name"`
	CreatedAt            time.Time          `gorm:"column:created_at"`
	UpdatedAt            time.Time          `gorm:"column:updated_at"`
}

type ClimateCondition struct {
	ID   								 int    						`gorm:"column:id;primaryKey"`
	Name 								 string 						`gorm:"column:name"`
	CreatedAt            time.Time          `gorm:"column:created_at"`
	UpdatedAt            time.Time          `gorm:"column:updated_at"`
}

type PlantImage struct {
	ID         int       `gorm:"column:id;primaryKey"`
	PlantID 	 int       `gorm:"column:plant_id"`
	FileName   string    `gorm:"column:file_name"`
	IsPrimary  int       `gorm:"column:is_primary"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

