package gedung

import (
	"Office-Booking/domain/gedung/request"
	"time"

	"gorm.io/gorm"
)

type Gedung struct {
	ID          int            `json:"id" gorm:"PrimaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	Name        string         `json:"name"`
	Location    string         `json:"location"`
	Price       string         `json:"price"`
	Latitude    string         `json:"latitude"`
	Longitude   string         `json:"longitude"`
	Description string         `json:"description"`
}

type GedungRepository interface {
	Create(gedung *Gedung) (*Gedung, error)
	GetAll() ([]Gedung, error)
	GetByID(id int) (*Gedung, error)
	GetByPrice(price string) (*Gedung, error)
	Update(id int) (*Gedung, error)
	Delete(id int) (*Gedung, error)
}

type GedungUsecase interface {
	Create(request request.PostRequest) (*Gedung, error)
	GetAll() ([]Gedung, error)
	GetByID(id int) (*Gedung, error)
	GetByPrice(price string) (*Gedung, error)
	Update(id int) (*Gedung, error)
	Delete(id int) (*Gedung, error)
}
