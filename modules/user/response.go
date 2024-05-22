package user

import (
	"time"
)

type UserResponse struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Is_Active  bool      `json:"is_active"`
	Url_Image  string    `json:"url_image"`
	Created_at time.Time `json:"created_at"`
}
