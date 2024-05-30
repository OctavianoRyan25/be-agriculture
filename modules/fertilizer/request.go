package fertilizer

import (
	"time"
)

type FertilizerRequest struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Compostition string `json:"compostition"`
	CreateAt time.Time `json:"createAt"`
}