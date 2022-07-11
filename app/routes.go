package app

import (
	"Office-Booking/app/config"
	_userController "Office-Booking/controllers/users"
	mid "Office-Booking/delivery/http/middleware"
	userRepository "Office-Booking/repository/users"
	userUsecase "Office-Booking/usecase/users"

	_nearbyController "Office-Booking/controllers/nearby"
	nearbyRepository "Office-Booking/repository/nearby"
	nearbyUsecase "Office-Booking/usecase/nearby"

	_jenisgedungController "Office-Booking/controllers/jenisgedung"
	jenisgedungRepository "Office-Booking/repository/jenisgedung"
	jenisgedungUsecase "Office-Booking/usecase/jenisgedung"

	_gedungController "Office-Booking/controllers/gedung"
	gedungRepository "Office-Booking/repository/gedung"
	gedungUsecase "Office-Booking/usecase/gedung"

	_reviewController "Office-Booking/controllers/review"
	reviewRepository "Office-Booking/repository/review"
	reviewUsecase "Office-Booking/usecase/review"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Run() {

	db := config.InitDB()
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)

	nearbyRepository := nearbyRepository.NewNearbyRepository(db)
	nearbyUsecase := nearbyUsecase.NewNearbyUsecase(nearbyRepository)

	jenisgedungRepository := jenisgedungRepository.NewJenisgedungRepository(db)
	jenisgedungUsecase := jenisgedungUsecase.NewJenisgedungUsecase(jenisgedungRepository)

	gedungRepository := gedungRepository.NewGedungRepository(db)
	gedungUsecase := gedungUsecase.NewGedungUseCase(gedungRepository)

	reviewRepository := reviewRepository.NewReviewRepository(db)
	reviewUsecase := reviewUsecase.NewReviewUseCase(reviewRepository)

	e := echo.New()
	mid.NewGoMiddleware().LogMiddleware(e)
	_userController.NewUserController(e, userUsecase)
	_nearbyController.NewNearbyController(e, nearbyUsecase)
	_jenisgedungController.NewJenisgedungController(e, jenisgedungUsecase)
	_gedungController.NewGedungController(e, gedungUsecase)
	_reviewController.NewReviewController(e, reviewUsecase)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://officebooking-app-pn6n3.ondigitalocean.app/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	address := fmt.Sprintf(":%d", 8080)

	if err := e.Start(address); err != nil {
		log.Info("Exit The Server")
	}
}
