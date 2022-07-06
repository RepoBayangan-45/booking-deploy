package review

import (
	domain "Office-Booking/domain/review"

	"gorm.io/gorm"
)

type reviewRepository struct {
	Conn *gorm.DB
}

// Create implements domain.ReviewRepository
func (u *reviewRepository) Create(review *domain.Review) (*domain.Review, error) {
	if err := u.Conn.Create(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

// Delete implements domain.ReviewRepository
func (u *reviewRepository) Delete(id int) (*domain.Review, error) {
	review := &domain.Review{ID: id}
	if err := u.Conn.Delete(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func NewReviewRepository(Conn *gorm.DB) domain.ReviewRepository {
	return &reviewRepository{Conn: Conn}
}
