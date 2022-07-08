package config

import (
	_ge "Office-Booking/domain/gedung"
	_je "Office-Booking/domain/jenisgedung"
	_ne "Office-Booking/domain/nearby"
	_re "Office-Booking/domain/review"
	_us "Office-Booking/domain/users"
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
		DB_Port:     "3306",
		DB_Host:     "office-booking-do-user-11917513-0.b.db.ondigitalocean.com",
		DB_Name:     "45_office_booking",
	}

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
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
	DB.AutoMigrate(&_us.User{})
	DB.AutoMigrate(&_ge.Gedung{})
	DB.AutoMigrate(&_je.Jenisgedung{})
	DB.AutoMigrate(&_ne.Nearby{})
	DB.AutoMigrate(&_re.Review{})
}
