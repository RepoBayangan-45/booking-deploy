package config

import (
	domain "Office-Booking/domain/users"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	config := Config{
		DB_Username: "officebooking",
		DB_Password: "AVNS_3PvNAqGO5pQkRw4rtHK",
		DB_Port:     "25060",
		DB_Host:     "office-booking-do-user-11917513-0.b.db.ondigitalocean.com",
		DB_Name:     "45_office_booking",

		// DB_Username: "root",
		// DB_Password: "root123",
		// DB_Port:     "3306",
		// DB_Host:     "127.0.0.1",
		// DB_Name:     "45_office_booking",
	}

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,
		config.DB_Password,
		config.DB_Port,
		config.DB_Host,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitialMigration()

	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&domain.User{})
}
