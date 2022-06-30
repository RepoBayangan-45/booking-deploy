package response

import "time"

type SuccessLogin struct {
	ID        int `json:"id" form:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Phone     string `json:"phone" form:"phone"`
	Token     string `json:"token" form:"token"`
}
