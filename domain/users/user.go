package domain

import (
	"Office-Booking/domain/users/request"
	"Office-Booking/domain/users/response"

	"gorm.io/gorm"
)

type User struct {
	ID           int            `json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
	Email        string         `json:"email"`
	Name         string         `json:"name"`
	Phone        string         `json:"phone"`
	Password     string `json:"password"`
}

type Users []User

type UserRepository interface {
	Create(user *User) (*User, error)
	ReadByID(id int) (*User, error)
	ReadByName(name string) (*User, error)
	ReadAll() (*Users, error)
	CheckLogin(user *User) (*User, bool, error)
	RegisterUser(user *User) (*User, error)
	Delete(id int) (*User, error)
	Updates(id int) (*User, error)
}

type UserUsecase interface {
	Create(request request.UserCreateRequest) (*User, error)
	ReadByID(id int) (*User, error)
	ReadByName(name string) (*User, error)
	ReadAll() (*Users, error)
	Login(request request.LoginRequest) (*response.SuccessLogin, error)
	RegisterUser(request request.RegisterRequest) (*User, error)
	Delete(id int) (*User, error)
	Updates(id int) (*User, error)
}
