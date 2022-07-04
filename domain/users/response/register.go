package response

import "time"

type SuccessRegister struct {
	ID           int `json:"id" form:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string `json:"email" form:"email"`
	Name         string `json:"name" form:"name"`
	Phone        string `json:"phone" form:"phone"`
	Password     string `json:"password" form:"password"`
}
