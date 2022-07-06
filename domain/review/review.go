package review

import "Office-Booking/domain/review/request"

type Review struct {
	ID          int     `json:"id" gorm:"PrimaryKey"`
	Rating      float64 `json:"rating"`
	Description string  `json:"description"`
}

type Reviews []Review

type ReviewRepository interface {
	Create(review *Review) (*Review, error)
	Delete(id int) (*Review, error)
}

type ReviewUsecase interface {
	Create(request request.ReviewPost) (*Review, error)
	Delete(id int) (*Review, error)
}
