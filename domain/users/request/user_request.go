package request

import "time"

type UserCreateRequest struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
}
