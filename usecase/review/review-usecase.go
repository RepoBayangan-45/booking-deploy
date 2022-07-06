package review

import (
	domain "Office-Booking/domain/review"
	"Office-Booking/domain/review/request"
	"errors"
)

type reviewUsecase struct {
	ReviewRepo domain.ReviewRepository
}

// Create implements domain.ReviewRepository
func (u *reviewUsecase) Create(request request.ReviewPost) (*domain.Review, error) {
	if request.Rating == 0.0 {
		return nil, errors.New("rating belum isi")
	}
	if request.Description == "" {
		return nil, errors.New("description belum isi")
	}
	review := &domain.Review{
		Rating:      request.Rating,
		Description: request.Description,
	}

	postReview, err := u.ReviewRepo.Create(review)
	if err != nil {
		return nil, err
	}

	return postReview, nil
}

// Delete implements domain.ReviewRepository
func (u *reviewUsecase) Delete(id int) (*domain.Review, error) {
	review, err := u.ReviewRepo.Delete(id)
	if err != nil {
		return nil, err
	}

	return review, err
}

func NewReviewUseCase(ur domain.ReviewRepository) domain.ReviewUsecase {
	return &reviewUsecase{ReviewRepo: ur}
}
