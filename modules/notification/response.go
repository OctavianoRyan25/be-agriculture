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
