package fertilizer

import "time"

type Fertilizer struct {
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Compostition string `json:"compostition"`
	CreateAt time.Time `json:"createAt"`
}