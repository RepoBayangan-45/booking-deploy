package request

import "time"

type RegisterRequest struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
}
