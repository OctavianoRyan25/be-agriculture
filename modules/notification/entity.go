package notification

import "time"

type Notification struct {
	Id        int `gorm:"primaryKey"`
	Title     string
	Body      string
	UserId    int `gorm:"foreignKey:UserID;references:Id"`
	IsRead    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
