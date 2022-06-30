package app

import (
	"Office-Booking/app/config"
	_userController "Office-Booking/controllers/users"
	mid "Office-Booking/delivery/http/middleware"
	repository "Office-Booking/repository/users"
	usecase "Office-Booking/usecase/users"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Run() {

	db := config.InitDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	e := echo.New()
	mid.NewGoMiddleware().LogMiddleware(e)
	_userController.NewUserController(e, userUsecase)
	address := fmt.Sprintf(":%d", 8080)

	if err := e.Start(address); err != nil {
		log.Info("Exit The Server")
	}
}
