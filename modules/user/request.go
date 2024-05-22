package user

import (
	"time"
)

type UserRequest struct {
	Name       string    `json:"name" valid:"required~Name is required"`
	Email      string    `json:"email" valid:"required~Email is required,email~Email is not valid"`
	Password   string    `json:"password" valid:"required~Password is required"`
	Is_Active  bool      `json:"is_active"`
	Url_Image  string    `json:"url_image"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
