package notification

import (
	"time"

	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
)

type Notification struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Body      string
	UserId    int `gorm:"foreignKey:UserID;references:Id"`
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomizeWateringReminder struct {
	Id        int `gorm:"primaryKey"`
	MyPlantId int `gorm:"foreignKey:MyPlantId;references:Id"`
	MyPlant   plant.UserPlant
	Time      string
	Recurring bool
	Type      string // "daily", "weekly", "monthly", "yearly
	CreatedAt time.Time
	UpdatedAt time.Time
}
