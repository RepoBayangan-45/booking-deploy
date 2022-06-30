package response

import "time"

type UserCreateResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	TanggalLahir time.Time
}
