package notification

import (
	"time"
)

type NotificationResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserId    int       `json:"user_id"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type CustomizeWateringReminderResponse struct {
	Id        int       `json:"id"`
	MyPlantID int       `json:"my_plant_id"`
	Time      string    `json:"time"`
	Recurring bool      `json:"recurring"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
