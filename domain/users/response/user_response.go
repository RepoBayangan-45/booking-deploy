package response

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
	Alamat   string `json:"alamat"`
	Phone    string `json:"phone"`
}
