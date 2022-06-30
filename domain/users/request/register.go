package request

import "time"

type RegisterRequest struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	TanggalLahir time.Time
	Phone        string `json:"phone"`
	Password     string `json:"password"`
}
