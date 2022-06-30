package config

import (
	domain "Office-Booking/domain/users"
	"fmt"

	"gorm.io/driver/sqlserver"
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
		DB_Username: "postgres",
		DB_Password: "root123",
		DB_Port:     "5432",
		DB_Host:     "localhost",
		DB_Name:     "45_office_booking",
	}

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",

		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitialMigration()

	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&domain.User{})
}
